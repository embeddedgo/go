// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/abi"
	"internal/cpu/r4000/creg"
	"unsafe"
)

// see saveGPRS and saveFPRS
const (
	numGPRS = 27
	numFPRS = 33
)

type mOS struct {
	gprs    [numGPRS]uintptr
	fprs    [numFPRS]float64
	sp, fp  uintptr // goroutine context
	ra, epc uintptr // exception context
}

var bbplayer bool // TODO move to n64 board support package

var (
	cpu0  cpuctx
	pcpu0 = &cpu0
)

//go:nowritebarrierrec
//go:nosplit
func taskerinit() {
	*(*uintptr)(unsafe.Pointer(&cpu0.t)) = uintptr(unsafe.Pointer(&thetasker))
	cpu0.exe.set(getg().m)
	allcpu := (*slice)(unsafe.Pointer(&thetasker.allcpu))
	*(*uintptr)(unsafe.Pointer(&allcpu.array)) = uintptr(unsafe.Pointer(&pcpu0))
	allcpu.len = 1
	allcpu.cap = 1
}

// This functions is called to put the CPU to sleep. It is allowed it does
// nothing.
//
//go:nosplit
func curcpuSleep() {}

// This function is used to inform the another CPU that there is a new thread
// added to its runnable queue. It should wake up the sleeping CPU or preempt
// the currently running thread to run the scheduler. The thread preemption can
// be set as delayed to allow a running thread to run for a minimum period of
// time.
//
//go:nosplit
func (cpu *cpuctx) newwork() {
	// TODO use only one of the two interrupts
	creg.CAUSE.SetBits(creg.IP_SW)
}

//go:nosplit
func curcpuWakeup() {
	// TODO implement
}

// This function is called to save the remaining context, not saved at syscall
// entry (eg. it can save FPU state).
//
//go:nosplit
func curcpuSavectxSched() {
	// TODO implement
}

// This function is called to save the remaining context, not saved at syscall
// entry (eg. it can save FPU state).
//
//go:nosplit
func curcpuSavectxCall() {
	// TODO implement
}

//go:nosplit
func cpuid() int {
	// for now only single CPU is supported
	return 0
}

// This function is called to create the inintial state of the new thread and
// save it in provided m.
//
//go:nosplit
func archnewm(m *m) {
	m.epc = abi.FuncPCABI0(mstart)
	m.sp = m.g0.stack.hi
	m.fp = uintptr(unsafe.Pointer(m.g0))
	m.ra = 1 // smallCtx flag
}

// Run scheduler immediately or at syscall exit. It's called only just before
// syscall exit.
//
// The actual context switch is performed by architecture specific code at
// curcpuRunScheduler exit. It should check the cpuctx.newexe variable and if
// true switch the context to the new thread specified in cpuctx.exe.
//
// Tasker code does not use FPU so the architecture specific context switch
// code can avoid saving/restoring FPU context if not need.
//
//go:nosplit
func curcpuSchedule() {
	// syscall still needs to duffcopy it's result back.  mark for schedule
	// and call scheduler after duffcopy in ISR.
	curcpu().schedule = true
}

// only called from ISR, do not use
func enterScheduler()
func saveGPRs()
func restoreGPRs()
func saveFPRs()
func restoreFPRs()

var lastlog [1024 * 16]byte
var lastidx = 0

// Just write into some ringbuffer for now, which can be inspected in a memory
// dump.
//
//go:nowritebarrierrec
//go:nosplit
func defaultWrite(fd int, p []byte) int {
	for i := 0; i < len(p); i++ {
		lastlog[lastidx] = p[i]
		lastidx += 1
		if lastidx >= len(lastlog) {
			lastidx = 0
		}
	}
	return len(p)
}

// syscalls not used by runtime
func syssetprivlevel(newlevel int) (oldlevel, errno int)       { return } // TODO
func sysirqctl(irq, ctl, ctxid int) (enabled, prio, errno int) { return } // TODO
