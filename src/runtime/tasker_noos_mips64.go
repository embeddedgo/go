// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Implements tasker support for mips64.
//
// Relies on external implementations of the following target specific
// functions, that must be provided by the target module via go:linkname pragma
// or in assembly:
//
//     _rt0_mips64_noos()
//     runtime.unhandledException()
//     runtime.defaultWrite(fd int, p []byte) int

import (
	"internal/abi"
	"internal/cpu/r4000/creg"
	"unsafe"
)

// See saveGPRS and saveFPRS
const (
	numGPRS = 28
	numFPRS = 33
)

type mOS struct {
	// thread context
	gprs            [numGPRS]uintptr
	fprs            [numFPRS]float64
	sp, fp, ra, epc uintptr
}

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

//go:nosplit
func curcpuSleep() {}

//go:nosplit
func (cpu *cpuctx) newwork() {
	creg.CAUSE.SetBits(creg.IP_SW0)
}

//go:nosplit
func curcpuWakeup() {}

//go:nosplit
func curcpuSavectxSched() {}

//go:nosplit
func curcpuSavectxCall() {}

//go:nosplit
func cpuid() int {
	// for now only single CPU is supported
	return 0
}

//go:nosplit
func archnewm(m *m) {
	m.epc = abi.FuncPCABI0(mstart)
	m.sp = m.g0.stack.hi
	m.fp = uintptr(unsafe.Pointer(m.g0))
	m.ra = 1 // smallCtx flag
}

//go:nosplit
func curcpuSchedule() {
	// syscall still needs to duffcopy it's result back.  mark for schedule
	// and call scheduler after duffcopy in ISR.
	curcpu().schedule = true
}

//go:nowritebarrierrec
//go:nosplit
func defaultWrite(fd int, p []byte) int

// only called from ISR, do not use
func enterScheduler()
func saveGPRs()
func restoreGPRs()
func saveFPRs()
func restoreFPRs()
func unhandledExternalInterrupt()

// syscalls unsupported by mips64
func syssetprivlevel(newlevel int) (oldlevel, errno int)       { return }
func sysirqctl(irq, ctl, ctxid int) (enabled, prio, errno int) { return }
