// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"unsafe"
)

const (
	_NSIG             = 0
	preemptMSupported = false
)

type gsignalStack struct{}
type sigset struct{}

//go:nosplit
func setThreadCPUProfiler(hz int32) {
	for {
		breakpoint()
	}
}

//go:nosplit
func unminit() {
	for {
		breakpoint()
	}
}

func clearSignalHandlers()           {}
func sigdisable(uint32)              {}
func sigenable(uint32)               {}
func sigignore(uint32)               {}
func signame(sig uint32) string      { return "" }
func sigblock(exiting bool)          {}
func msigrestore(sigmask sigset)     {}
func initsig(preinit bool)           {}
func setProcessCPUProfiler(hz int32) {}
func mpreinit(mp *m)                 {}
func sigsave(p *sigset)              {}
func goenvs()                        {}
func minit()                         {}
func mdestroy(mp *m)                 {}
func preemptM(mp *m)                 {}

//go:nosplit
func crash() {
	for {
		breakpoint()
	}
}

//go:nosplit
func setsystim(nanotime func() int64, setalarm func(ns int64)) {
	if nanotime != nil {
		thetasker.newnanotime = nanotime
	} else {
		thetasker.newnanotime = dummyNanotime
	}
	if setalarm != nil {
		thetasker.newsetalarm = setalarm
	} else {
		thetasker.newsetalarm = dummySetalarm

	}
	setsystim1()
}

//go:nosplit
func setsyswriter(w func(fd int, p []byte) int) {
	if w != nil {
		thetasker.newwrite = w
	} else {
		thetasker.newwrite = defaultWrite
	}
	setsyswriter1()
}

//go:nosplit
func usleep(usec uint32) {
	nanosleep(int64(usec) * 1000)
}

//go:nosplit
func usleep_no_g(usec uint32) {
	nanosleep(int64(usec) * 1000)
}

//go:nosplit
func getRandomData(r []byte) {
	// BUG: true random data required
	extendRandom(r, 0)
}

//go:nosplit
func osinit() {
	ncpu = 1 // for now only single CPU is supported (see identcurcpu, cpuid)
	physPageSize = _PageSize
}

//go:nosplit
func isr() bool {
	gp := getg()
	return gp == gp.m.gsignal
	/*
		allcpu := thetasker.allcpu
		for i := 0; i < len(allcpu); i++ {
			if getg() == &allcpu[i].gh {
				return true
			}
		}
		return false
	*/
}

var timestart struct {
	sec  int64
	nsec int32
	lock rwmutex
}

// The noos has separate implementation of time_now for better performance (only
// one syscall) and to ensure that the monotonic and the wall components both
// point to the same instant in time (the RTC setting code can rely on this).

//go:linkname time_now time.now
func time_now() (sec int64, nsec int32, mono int64) {
	timestart.lock.rlock()
	sec = timestart.sec
	nsec = timestart.nsec
	timestart.lock.runlock()
	mono = nanotime()
	s := mono / 1e9
	ns := int32(mono - s*1e9)
	sec += s
	nsec += ns
	if nsec >= 1e9 {
		sec++
		nsec -= 1e9
	}
	return
}

//go:linkname time_move time.move
func time_move(sec int64, nsec int32) (sec0 int64, nsec0 int32) {
	timestart.lock.lock()
	sec0 = timestart.sec + sec
	nsec0 = timestart.nsec + nsec
	if nsec0 < 0 {
		nsec0 += 1e9
		sec0--
	} else if nsec0 >= 1e9 {
		nsec0 -= 1e9
		sec0++
	}
	timestart.sec = sec0
	timestart.nsec = nsec0
	timestart.lock.unlock()
	return
}

//go:nosplit
func osyield_no_g() {
	osyield()
}

func setsystim1()
func setsyswriter1()
func newosproc(mp *m)
func exit(code int32)
func osyield()

//go:noescape
func write(fd uintptr, p unsafe.Pointer, n int32) int32

//go:noescape
func futexsleep(addr *uint32, val uint32, ns int64)

//go:noescape
func futexwakeup(addr *uint32, cnt uint32)

//go:noescape
func exitThread(wait *uint32)

// syscalls not used by runtime

func setprivlevel(newlevel int) (oldlevel, errno int)
func irqenabled(irq int) (enabled, errno int)
func setirqenabled(irq, enabled int) (errno int)
func irqctl(irq, ctl, ctxid int) (enabled, prio, errno int)
func nanosleep(ns int64)
func nanotime() int64

// faketime is the simulated time in nanoseconds since 1970 for the
// playground.
//
// Zero means not to use faketime.
var faketime int64
