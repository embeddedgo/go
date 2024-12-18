// Copyright 2023 The Go Authors. All rights reserved.
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
	"internal/cpu"
	"internal/cpu/r4000/creg"
	"runtime/internal/atomic"
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

const cacheLineSize = cpu.CacheLineMinSize
const cacheLineMask = ^(cacheLineSize - 1)

//go:noescape
func writeback(addr uintptr, length int)

//go:noescape
func invalidate(addr uintptr, length int)

//go:noescape
func writebackInvalidate(addr uintptr, length int)

//go:nosplit
func syscachemaint(op int, p unsafe.Pointer, size int) {
	switch op {
	case 0: // rtos.DCacheInval
		invalidate(uintptr(p), size)
	case 1: // rtos.DCacheFlush
		writeback(uintptr(p), size)
	case 2: // rtos.DCacheFlushInval
		writebackInvalidate(uintptr(p), size)
	}
}

var highPrioIRQMask uint32

func sysirqctl(irq, ctl, ctxid int) (enabled, prio, errno int) {
	if uint(irq) > 8 { // TODO PIC dependent
		errno = 4 // rtos.ErrBadIntNumber
		return
	}
	if uint(ctxid) != 0 {
		errno = 6 // rtos.ErrBadIntCtx
	}

	irqMask := uint32(1 << (irq + 7))
	switch ctl { // Values defined in embedded/rtos/irq_noos_mips64.go
	case 1:
		atomic.Or32(&highPrioIRQMask, irqMask)
	case 0:
		atomic.And32(&highPrioIRQMask, ^irqMask)
	case -1: // IRQ.Enable()
		creg.STATUS.SetBits(irqMask)
	case -2: // IRQ.Disable()
		creg.STATUS.ClearBits(irqMask)
	case -3: // IRQ.Status()
		if irqMask&atomic.Load(&highPrioIRQMask) != 0 {
			prio = 1
		}
		if creg.STATUS.LoadBits(irqMask) != 0 {
			enabled = 1
		}
	}

	return
}

// syscalls unsupported by mips64
func syssetprivlevel(newlevel int) (oldlevel, errno int) { return }

// only called from ISR, do not use
func enterScheduler()
func saveGPRs()
func restoreGPRs()
func saveFPRs()
func restoreFPRs()
func unhandledExternalInterrupt()
