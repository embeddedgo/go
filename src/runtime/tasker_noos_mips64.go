// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

type mOS [1]uint32

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
	// for now only single CPU is supported
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
	// TODO implement
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
	// TODO implement
}

//go:nowritebarrierrec
//go:nosplit
func defaultWrite(fd int, p []byte) int {
	return 0
}
