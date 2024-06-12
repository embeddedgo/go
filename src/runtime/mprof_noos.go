// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

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
	MemProfileRate         int
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
func mutexevent(cycles int64, skip int)                     {}
func tryRecordGoroutineProfileWB(gp1 *g)                    {}
func tryRecordGoroutineProfile(gp1 *g, yield func())        {}

// Stack formats a stack trace of the calling goroutine into buf
// and returns the number of bytes written to buf.
// If all is true, Stack formats stack traces of all other goroutines
// into buf after the trace for the current goroutine.
func Stack(buf []byte, all bool) int {
	if all {
		stopTheWorld("stack trace")
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
		startTheWorld()
	}
	return n
}
