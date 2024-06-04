// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"runtime/internal/atomic"
	"unsafe"
)

const (
	blockprofilerate = 0
	mutexprofilerate = 0
)

const (
	goroutineProfileAbsent = iota
	goroutineProfileInProgress
	goroutineProfileSatisfied
)

var (
	disableMemoryProfiling bool
	profInsertLock         mutex
	profBlockLock          mutex
	profMemActiveLock      mutex
	profMemFutureLock      [len(memRecord{}.future)]mutex
	goroutineProfile       struct{ active bool }
)

type bucket struct{}
type memRecordCycle struct{}

type memRecord struct {
	future [0]memRecordCycle
}

type goroutineProfileStateHolder struct{}

func (p *goroutineProfileStateHolder) Store(x int) {}

func blockevent(cycles int64, skip int)                     {}
func tracealloc(p unsafe.Pointer, size uintptr, typ *_type) {}
func mProf_Malloc(p unsafe.Pointer, size uintptr)           {}
func mProf_PostSweep()                                      {}
func mProf_NextCycle()                                      {}
func mProf_Flush()                                          {}
func tracegc()                                              {}
func tracefree(p unsafe.Pointer, size uintptr)              {}
func mProf_Free(b *bucket, size uintptr)                    {}
func tryRecordGoroutineProfileWB(gp1 *g)                    {}
func tryRecordGoroutineProfile(gp1 *g, yield func())        {}

//go:linkname mutexevent sync.event
func mutexevent(cycles int64, skip int) {}

// Stack formats a stack trace of the calling goroutine into buf
// and returns the number of bytes written to buf.
// If all is true, Stack formats stack traces of all other goroutines
// into buf after the trace for the current goroutine.
func Stack(buf []byte, all bool) int {
	var stw worldStop
	if all {
		stw = stopTheWorld(stwAllGoroutinesStack)
	}

	n := 0
	if len(buf) > 0 {
		gp := getg()
		sp := getcallersp()
		pc := getcallerpc()
		systemstack(func() {
			g0 := getg()
			// Force traceback=1 to override GOTRACEBACK setting,
			// so that Stack's results are consistent.
			// GOTRACEBACK is only about crash dumps.
			g0.m.traceback = 1
			g0.writebuf = buf[0:0:len(buf)]
			goroutineheader(gp)
			traceback(pc, sp, 0, gp)
			if all {
				tracebackothers(gp)
			}
			g0.m.traceback = 0
			n = len(g0.writebuf)
			g0.writebuf = nil
		})
	}

	if all {
		startTheWorld(stw)
	}
	return n
}

type lockTimer struct {
	lock *mutex
}

func (lt *lockTimer) begin() {}
func (lt *lockTimer) end()   {}

type mLockProfile struct {
	waitTime atomic.Int64
}

func (prof *mLockProfile) recordUnlock(l *mutex) {}

type BlockProfileRecord struct {
	Count  int64
	Cycles int64
	StackRecord
}

type StackRecord struct {
	Stack0 [0]uintptr
}

func (r *StackRecord) Stack() []uintptr { return nil }

type MemProfileRecord struct {
	AllocBytes, FreeBytes     int64
	AllocObjects, FreeObjects int64
	Stack0                    [0]uintptr
}

func (r *MemProfileRecord) InUseBytes() int64   { return 0 }
func (r *MemProfileRecord) InUseObjects() int64 { return 0 }
func (r *MemProfileRecord) Stack() []uintptr    { return nil }

var MemProfileRate int

func ThreadCreateProfile(p []StackRecord) (n int, ok bool)             { return 0, true }
func MemProfile(p []MemProfileRecord, inuseZero bool) (n int, ok bool) { return 0, true }
func SetBlockProfileRate(rate int)                                     {}
func SetMutexProfileFraction(rate int) int                             { return 0 }
func BlockProfile(p []BlockProfileRecord) (n int, ok bool)             { return 0, true }
func MutexProfile(p []BlockProfileRecord) (n int, ok bool)             { return 0, true }
