// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"runtime/internal/atomic"
)

// A mq represents a queue of threads.
type mq struct {
	_mqprivate
}

type _mqprivate struct {
	first muintptr
	last  muintptr
	n     uint
	mx    cpumtx
}

//go:nosplit
func (q *mq) lock() { q.mx.lock() }

//go:nosplit
func (q *mq) unlock() { q.mx.unlock() }

// atomicLen returns the approximate number of elements in the q. It returns an
// exact value if called by the only mutator of q.
//go:nosplit
func (q *mq) atomicLen() int { return int(atomic.Loaduint(&q.n)) }

// push inserts m at the end of q
//go:nosplit
func (q *mq) push(m *m) {
	if q.n == 0 {
		q.first.set(m)
	} else {
		msetnext1(q.last.ptr(), m)
	}
	q.last.set(m)
	q.n++
}

// pop removes the first m from the beginning of q and returns it
//go:nosplit
func (q *mq) pop() *m {
	var ret *m
	if q.n != 0 {
		q.n--
		ret = q.first.ptr()
		q.first.set(mnext1(ret))
	}
	return ret
}

// A mcl represents a circular list of threads.
type mcl struct {
	_mclprivate
	//_ [(cpu.CacheLinePadSize - unsafe.Sizeof(_mclprivate{}))]byte
}

type _mclprivate struct {
	cur muintptr
	n   uint
	mx  cpumtx
}

//go:nosplit
func (q *mcl) lock() { q.mx.lock() }

//go:nosplit
func (q *mcl) unlock() { q.mx.unlock() }

// push inserts m just before the current m.
//go:nosplit
func (q *mcl) push(m *m) {
	if q.n == 0 {
		msetnext1(m, m)
		msetprev1(m, m)
		q.cur.set(m)
	} else {
		cur := q.cur.ptr()
		prev := mprev1(cur)
		msetprev1(m, prev)
		msetnext1(m, cur)
		msetnext1(prev, m)
		msetprev1(cur, m)
	}
	q.n++
}

// find finds and returns the pointer to the first m in q that matches the
// provided key. As a side effect it rotates the q so the m.next becomes the
// current element.
//go:nosplit
func (q *mcl) find(key uintptr) *m {
	if q.n == 0 {
		return nil
	}
	cur := q.cur.ptr()
	first := cur
	for {
		if mkey(cur) == key {
			q.cur.set(mnext1(cur))
			return cur
		}
		cur = mnext1(cur)
		if cur == first {
			return nil
		}
	}
}

// remove removes m from q. Remove is fast (O(1)) but the caller must ensure
// that m belongs to q, othervise the effect of remove is unpredictable.
//go:nosplit
func (q *mcl) remove(m *m) {
	q.n--
	if q.n == 0 {
		return
	}
	prev := mprev1(m)
	next := mnext1(m)
	if m == q.cur.ptr() {
		q.cur.set(next)
	}
	msetnext1(prev, next)
	msetprev1(next, prev)
}

// A msl represents a sorted list of threads.
type msl struct {
	_mslprivate
}

type _mslprivate struct {
	head muintptr
	n    uint
	mx   cpumtx
}

//go:nosplit
func (q *msl) lock() { q.mx.lock() }

//go:nosplit
func (q *msl) unlock() { q.mx.unlock() }

// insertbyval inserts m into q just before the first item with greater value
//go:nosplit
func (q *msl) insertbyval(m *m) {
	if q.n == 0 {
		msetnext2(m, m)
		msetprev2(m, m)
		q.head.set(m)
	} else {
		val := mval(m)
		// search from the last to the first because mval(m) is strongly
		// corelated with the current (monotonic) time
		last := mprev2(q.head.ptr())
		cur := last
		for {
			if mval(cur) <= val {
				break
			}
			cur = mprev2(cur)
			if cur == last {
				q.head.set(m)
				break
			}
		}
		next := mnext2(cur)
		msetprev2(m, cur)
		msetnext2(m, next)
		msetnext2(cur, m)
		msetprev2(next, m)

	}
	q.n++
}

// first returns the pointer to the first m in q.
//go:nosplit
func (q *msl) first() *m {
	if q.n == 0 {
		return nil
	}
	return q.head.ptr()
}

// remove removes m from q. Remove is fast (O(1)) but the caller must ensure
// that m belongs to q, othervise the effect of remove is unpredictable.
//go:nosplit
func (q *msl) remove(m *m) {
	q.n--
	if q.n == 0 {
		return
	}
	prev := mprev2(m)
	next := mnext2(m)
	if m == q.head.ptr() {
		q.head.set(next)
	}
	msetnext2(prev, next)
	msetprev2(next, prev)
}

// m fields used

//go:nosplit
func mprev1(m *m) *m { return muintptr(m.tls[0]).ptr() }

//go:nosplit
func mnext1(m *m) *m { return muintptr(m.tls[1]).ptr() }

//go:nosplit
func mprev2(m *m) *m { return muintptr(m.tls[2]).ptr() }

//go:nosplit
func mnext2(m *m) *m { return muintptr(m.tls[3]).ptr() }

//go:nosplit
func mval(m *m) int64 { return int64(m.ncgocall) }

//go:nosplit
func mkey(m *m) uintptr { return m.thread }

//go:nosplit
func msetprev1(m, prev *m) { (*muintptr)(&m.tls[0]).set(prev) }

//go:nosplit
func msetnext1(m, next *m) { (*muintptr)(&m.tls[1]).set(next) }

//go:nosplit
func msetprev2(m, prev *m) { (*muintptr)(&m.tls[2]).set(prev) }

//go:nosplit
func msetnext2(m, next *m) { (*muintptr)(&m.tls[3]).set(next) }

//go:nosplit
func msetval(m *m, val int64) { m.ncgocall = uint64(val) }

//go:nosplit
func msetkey(m *m, key uintptr) { m.thread = key }
