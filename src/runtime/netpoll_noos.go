// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build noos

package runtime

import (
	"runtime/internal/atomic"
	"runtime/internal/sys"
	"unsafe"
)

var netpollInited atomic.Uint32
var netpollWaiters atomic.Uint32

var netpollStubLock mutex
var netpollNote note

var wakerq pollList

func netpollGenericInit() {
	netpollInited.Store(1)
}

func netpollBreak() {
	// do not use notewakeup, we allow multiple wakeups for this note
	if !atomic.Cas(key32(&netpollNote.key), 0, 1) {
		return
	}
	futexwakeup(key32(&netpollNote.key), 1)
	return
}

// Polls for goroutines waiting on interrupts.
// Returns list of goroutines that become runnable.
func netpoll(delay int64) (toRun gList, delta int32) {
	// This lock ensures that only one goroutine tries to use
	// the note. It should normally be completely uncontended.
	lock(&netpollStubLock)

	if !atomic.Cas(key32(&netpollNote.key), 1, 0) { // try noteclear
		notetsleep(&netpollNote, delay)
	}

	n := wakerq.free()
	for n != nil {
		delta += netpollready(&toRun, n)
		n = n.pop()
	}

	unlock(&netpollStubLock)

	return
}

func netpollinited() bool {
	return netpollInited.Load() != 0
}

// netpollAnyWaiters reports whether any goroutines are waiting for I/O.
func netpollAnyWaiters() bool {
	return netpollWaiters.Load() > 0
}

// netpollAdjustWaiters adds delta to netpollWaiters.
func netpollAdjustWaiters(delta int32) {
	if delta != 0 {
		netpollWaiters.Add(delta)
	}
}

const (
	pdNil   uintptr = 0
	pdReady uintptr = 1
	pdWait  uintptr = 2
)

// netpollblock parks the goroutine on pd.  It returns whether the note was
// woken up in the timeout specified by ns.
func netpollblock(pd *pollDesc, ns int64) bool {
	var gp *g
	var t *timer
	gpp := &pd.g

	if ns == 0 {
		goto clear
	}

	lock(&pd.lock)
	pd.seq++
	unlock(&pd.lock)

	// configure deadline timer
	gp = getg()
	t = gp.timer
	if t == nil {
		t = new(timer)
		gp.timer = t
	}
	t.f = netpolldeadline
	t.arg = pd
	t.seq = pd.seq
	if ns < 0 {
		t.nextwhen = maxWhen
	} else {
		t.nextwhen = nanotime() + ns
		if t.nextwhen < 0 { // check for overflow.
			t.nextwhen = maxWhen
		}
	}

	// set the gpp semaphore to pdWait
	for {
		// Consume notification if already ready.
		if gpp.CompareAndSwap(pdReady, pdNil) {
			return true
		}
		if gpp.CompareAndSwap(pdNil, pdWait) {
			break
		}

		// Double check that this isn't corrupt; otherwise we'd loop
		// forever.
		if v := gpp.Load(); v != pdReady && v != pdNil {
			throw("runtime: double wait")
		}
	}

	gopark(netpollblockcommit, unsafe.Pointer(gpp), waitReasonIOWait, traceBlockNet, 5)

clear:
	// be careful to not lose concurrent pdReady notification
	old := gpp.Swap(pdNil)
	if old > pdWait {
		throw("runtime: corrupted polldesc")
	}
	return old == pdReady
}

func netpollblockcommit(gp *g, gpp unsafe.Pointer) bool {
	r := atomic.Casuintptr((*uintptr)(gpp), pdWait, uintptr(unsafe.Pointer(gp)))
	resettimer(gp.timer, gp.timer.nextwhen)
	if r {
		netpollAdjustWaiters(1)
	}
	return r
}

// netpollunblock moves pd.g depending on ioready into the pdNil or pdReady
// state. This returns any goroutine blocked on pd.g. It adds any adjustment to
// netpollWaiters to *delta; this adjustment should be applied after the
// goroutine has been marked ready.
func netpollunblock(pd *pollDesc, ioready bool, delta *int32) *g {
	gpp := &pd.g

	for {
		old := gpp.Load()
		if old == pdReady {
			return nil
		}
		if old == pdNil && !ioready {
			return nil
		}
		new := pdNil
		if ioready {
			new = pdReady
		}
		if gpp.CompareAndSwap(old, new) {
			if old == pdWait {
				old = pdNil
			} else if old != pdNil {
				*delta -= 1
			}
			return (*g)(unsafe.Pointer(old))
		}
	}
}

// netpollready declares that the g associated with pd is ready to run. The
// toRun argument is used to build a list of goroutines to return from netpoll.
//
// This returns a delta to apply to netpollWaiters.
//
// This may run while the world is stopped, so write barriers are not allowed.
//
//go:nowritebarrier
func netpollready(toRun *gList, pd *pollDesc) (delta int32) {
	gp := netpollunblock(pd, true, &delta)
	if gp != nil {
		toRun.push(gp)
	}
	return
}

// netpolldeadline is the deadline timers callback.
func netpolldeadline(arg any, seq uintptr) {
	pd := arg.(*pollDesc)

	lock(&pd.lock)

	if pd.seq != seq {
		unlock(&pd.lock)
		return // timer is stale, ignore
	}

	delta := int32(0)
	gp := netpollunblock(pd, false, &delta)
	unlock(&pd.lock)
	if gp != nil {
		goready(gp, 1)
	}
	netpollAdjustWaiters(delta)
}

func rtos_condwait(n *pollDesc, timeout int64) bool {
	return netpollblock(n, timeout)
}

// rtos_condsignal wakes up the netpoller if a goroutine is waiting or in
// pdWait.  Otherwise it only sets the event to pdReady.
//
//go:nowritebarrierrec
//go:nosplit
func rtos_condsignal(n *pollDesc) {
	for {
		old := n.g.Load()
		if old < pdWait {
			if n.g.CompareAndSwap(old, pdReady) {
				return
			}
		} else {
			break
		}
	}

	if wakerq.insert(n) {
		if isr() {
			if !atomic.Cas(key32(&netpollNote.key), 0, 1) {
				return
			}
			curcpu().wakeNetpoller = true
			curcpuWakeup()
		} else {
			netpollBreak()
		}
	}
}

type pollDesc struct {
	_    sys.NotInHeap
	g    atomic.Uintptr
	seq  uintptr
	lock mutex // protects seq
	link atomic.Uintptr
}

//go:nosplit
func (n *pollDesc) pop() *pollDesc {
	next := n.link.Swap(0)
	return (*pollDesc)(unsafe.Pointer(next))
}

type pollList struct {
	head atomic.Uintptr
}

// insert inserts n at the beginning of l if it isn't already
//
//go:nosplit
func (l *pollList) insert(n *pollDesc) bool {
	if !n.link.CompareAndSwap(0, 1) {
		return false
	}
	for {
		head := l.head.Load()
		n.link.Store(head)
		if l.head.CompareAndSwap(head, uintptr(unsafe.Pointer(n))) {
			break
		}
	}
	return true
}

//go:nosplit
func (l *pollList) free() *pollDesc {
	return (*pollDesc)(unsafe.Pointer(l.head.Swap(0)))
}
