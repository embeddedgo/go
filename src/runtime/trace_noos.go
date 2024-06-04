// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const (
	traceBlockGeneric traceBlockReason = iota
	traceBlockForever
	traceBlockNet
	traceBlockSelect
	traceBlockCondWait
	traceBlockSync
	traceBlockChanSend
	traceBlockChanRecv
	traceBlockGCMarkAssist
	traceBlockGCSweep
	traceBlockSystemGoroutine
	traceBlockPreempted
	traceBlockDebugCall
	traceBlockUntilGCEnds
	traceBlockSleep
)

const defaultTraceAdvancePeriod = 0

var trace struct {
	lock mutex
}

type traceBlockReason uint8

type mTraceState struct{}
type pTraceState struct{ reclaimed uintptr }

type gTraceState struct{}

func (_ *gTraceState) reset() {}

type traceLocker struct{}

func (_ traceLocker) GCSweepStart()                            {}
func (_ traceLocker) GCSweepDone()                             {}
func (_ traceLocker) GCStart()                                 {}
func (_ traceLocker) GCDone()                                  {}
func (_ traceLocker) GCMarkAssistDone()                        {}
func (_ traceLocker) GCMarkAssistStart()                       {}
func (_ traceLocker) GoUnpark(gp *g, skip int)                 {}
func (_ traceLocker) HeapAlloc(live uint64)                    {}
func (_ traceLocker) HeapGoal()                                {}
func (_ traceLocker) GCSweepSpan(bytesSwept uintptr)           {}
func (_ traceLocker) STWStart(reason stwReason)                {}
func (_ traceLocker) GoSysBlock(pp *p)                         {}
func (_ traceLocker) ProcSteal(pp *p, forMe bool)              {}
func (_ traceLocker) STWDone()                                 {}
func (_ traceLocker) GoCreateSyscall(gp *g)                    {}
func (_ traceLocker) OneNewExtraM(gp *g)                       {}
func (_ traceLocker) GoDestroySyscall()                        {}
func (_ traceLocker) GoSysExit(lostP bool)                     {}
func (_ traceLocker) GoStart()                                 {}
func (_ traceLocker) GoPark(reason traceBlockReason, skip int) {}
func (_ traceLocker) GoPreempt()                               {}
func (_ traceLocker) GoSched()                                 {}
func (_ traceLocker) GoEnd()                                   {}
func (_ traceLocker) GoSysCall()                               {}
func (_ traceLocker) ProcStop(pp *p)                           {}
func (_ traceLocker) ProcStart()                               {}
func (_ traceLocker) GoCreate(newg *g, pc uintptr)             {}
func (_ traceLocker) Gomaxprocs(procs int32)                   {}

//go:nosplit
func (_ traceLocker) ok() bool { return false }

//go:nosplit
func (_ traceLocker) RecordSyscallExitedTime(gp *g, oldp *p) {}

func traceLockInit()                                   {}
func traceThreadDestroy(_ *m)                          {}
func traceReaderAvailable() *g                         { return nil }
func traceExitingSyscall()                             {}
func traceExitedSyscall()                              {}
func traceCPUSample(gp *g, _ *m, pp *p, stk []uintptr) {}

//go:nosplit
func traceAcquire() traceLocker { return traceLocker{} }

//go:nosplit
func traceRelease(tl traceLocker) {}

//go:nosplit
func traceEnabled() bool { return false }

//go:nosplit
func traceShuttingDown() bool { return true }

//go:systemstack
func traceReader() *g { return nil }

//go:systemstack
func traceProcFree(pp *p) {}

func StartTrace() error { return nil }
func ReadTrace() []byte { return nil }
func StopTrace()        {}
