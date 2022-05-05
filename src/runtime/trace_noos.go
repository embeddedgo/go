package runtime

// Event types in the trace, args are given in square brackets.
const (
	traceEvNone              = 0  // unused
	traceEvBatch             = 1  // start of per-P batch of events [pid, timestamp]
	traceEvFrequency         = 2  // contains tracer timer frequency [frequency (ticks per second)]
	traceEvStack             = 3  // stack [stack id, number of PCs, array of {PC, func string ID, file string ID, line}]
	traceEvGomaxprocs        = 4  // current value of GOMAXPROCS [timestamp, GOMAXPROCS, stack id]
	traceEvProcStart         = 5  // start of P [timestamp, thread id]
	traceEvProcStop          = 6  // stop of P [timestamp]
	traceEvGCStart           = 7  // GC start [timestamp, seq, stack id]
	traceEvGCDone            = 8  // GC done [timestamp]
	traceEvGCSTWStart        = 9  // GC STW start [timestamp, kind]
	traceEvGCSTWDone         = 10 // GC STW done [timestamp]
	traceEvGCSweepStart      = 11 // GC sweep start [timestamp, stack id]
	traceEvGCSweepDone       = 12 // GC sweep done [timestamp, swept, reclaimed]
	traceEvGoCreate          = 13 // goroutine creation [timestamp, new goroutine id, new stack id, stack id]
	traceEvGoStart           = 14 // goroutine starts running [timestamp, goroutine id, seq]
	traceEvGoEnd             = 15 // goroutine ends [timestamp]
	traceEvGoStop            = 16 // goroutine stops (like in select{}) [timestamp, stack]
	traceEvGoSched           = 17 // goroutine calls Gosched [timestamp, stack]
	traceEvGoPreempt         = 18 // goroutine is preempted [timestamp, stack]
	traceEvGoSleep           = 19 // goroutine calls Sleep [timestamp, stack]
	traceEvGoBlock           = 20 // goroutine blocks [timestamp, stack]
	traceEvGoUnblock         = 21 // goroutine is unblocked [timestamp, goroutine id, seq, stack]
	traceEvGoBlockSend       = 22 // goroutine blocks on chan send [timestamp, stack]
	traceEvGoBlockRecv       = 23 // goroutine blocks on chan recv [timestamp, stack]
	traceEvGoBlockSelect     = 24 // goroutine blocks on select [timestamp, stack]
	traceEvGoBlockSync       = 25 // goroutine blocks on Mutex/RWMutex [timestamp, stack]
	traceEvGoBlockCond       = 26 // goroutine blocks on Cond [timestamp, stack]
	traceEvGoBlockNet        = 27 // goroutine blocks on network [timestamp, stack]
	traceEvGoSysCall         = 28 // syscall enter [timestamp, stack]
	traceEvGoSysExit         = 29 // syscall exit [timestamp, goroutine id, seq, real timestamp]
	traceEvGoSysBlock        = 30 // syscall blocks [timestamp]
	traceEvGoWaiting         = 31 // denotes that goroutine is blocked when tracing starts [timestamp, goroutine id]
	traceEvGoInSyscall       = 32 // denotes that goroutine is in syscall when tracing starts [timestamp, goroutine id]
	traceEvHeapAlloc         = 33 // gcController.heapLive change [timestamp, heap_alloc]
	traceEvHeapGoal          = 34 // gcController.heapGoal (formerly next_gc) change [timestamp, heap goal in bytes]
	traceEvTimerGoroutine    = 35 // not currently used; previously denoted timer goroutine [timer goroutine id]
	traceEvFutileWakeup      = 36 // denotes that the previous wakeup of this goroutine was futile [timestamp]
	traceEvString            = 37 // string dictionary entry [ID, length, string]
	traceEvGoStartLocal      = 38 // goroutine starts running on the same P as the last event [timestamp, goroutine id]
	traceEvGoUnblockLocal    = 39 // goroutine is unblocked on the same P as the last event [timestamp, goroutine id, stack]
	traceEvGoSysExitLocal    = 40 // syscall exit on the same P as the last event [timestamp, goroutine id, real timestamp]
	traceEvGoStartLabel      = 41 // goroutine starts running with label [timestamp, goroutine id, seq, label string id]
	traceEvGoBlockGC         = 42 // goroutine blocks on GC assist [timestamp, stack]
	traceEvGCMarkAssistStart = 43 // GC mark assist start [timestamp, stack]
	traceEvGCMarkAssistDone  = 44 // GC mark assist done [timestamp]
	traceEvUserTaskCreate    = 45 // trace.NewContext [timestamp, internal task id, internal parent task id, stack, name string]
	traceEvUserTaskEnd       = 46 // end of a task [timestamp, internal task id, stack]
	traceEvUserRegion        = 47 // trace.WithRegion [timestamp, internal task id, mode(0:start, 1:end), stack, name string]
	traceEvUserLog           = 48 // trace.Log [timestamp, internal task id, key string id, stack, value string]
	traceEvCount             = 49
	// Byte is used but only 6 bits are available for event type.
	// The remaining 2 bits are used to specify the number of arguments.
	// That means, the max event type value is 63.
)

type traceBufPtr uintptr

var trace struct {
	lock        mutex
	stringsLock mutex
	bufLock     mutex
	stackTab    struct{ lock mutex }
	enabled     bool
	shutdown    bool
}

func traceGCSweepStart()                  {}
func traceGCSweepDone()                   {}
func traceGCMarkAssistStart()             {}
func traceGCMarkAssistDone()              {}
func traceHeapAlloc()                     {}
func traceHeapGoal()                     {}
func traceGoUnpark(gp *g, skip int)       {}
func traceGCStart()                       {}
func traceGCDone()                        {}
func traceGCSTWStart(kind int)            {}
func traceGCSTWDone()                     {}
func traceGCSweepSpan(bytesSwept uintptr) {}
func traceProcStart()                     {}
func traceProcStop(pp *p)                 {}
func traceProcFree(pp *p)                 {}
func traceGoCreate(newg *g, pc uintptr)   {}
func traceGoStart()                       {}
func traceGoSched()                       {}
func traceGoPreempt()                     {}
func traceGoEnd()                         {}
func traceGoPark(traceEv byte, skip int)  {}
func traceGomaxprocs(procs int32)         {}
func traceGoSysCall()                     {}
func traceGoSysBlock(pp *p)               {}
func traceGoSysExit(ts int64)             {}
func traceReader() *g                     { return nil }
