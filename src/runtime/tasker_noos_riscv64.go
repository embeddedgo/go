// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/cpu/riscv/clint"
	"internal/cpu/riscv/plic"
	"unsafe"
)

func cpuid() int
func curcpuSavectxSched()
func curcpuSleep()
func curcpuWakeup()      {}
func curcpuSavectxCall() {} // all registars saved on caller's stack

//go:nosplit
func (cpu *cpuctx) newwork() {
	clint.CLINT().MSIP[cpu.id()].Store(1)
}

//go:nosplit
func curcpuSchedule() { curcpu().schedule = true }

const thrSmallCtx = 1 // context saved in mOS contains only LR, SP, g

//go:nosplit
func archnewm(m *m) {
	m.tls[mstatus] = thrSmallCtx
	m.tls[mepc] = funcPC(mstart)
	m.x[numGPRS-3] = m.g0.stack.hi                 // SP
	m.x[numGPRS-2] = uintptr(unsafe.Pointer(m.g0)) // g
}

const maxHarts = 2

var (
	harts  [maxHarts]cpuctx
	pharts [maxHarts]*cpuctx
)

//go:nowritebarrierrec
//go:nosplit
func taskerinit() {
	// only hart0 runs this function

	uharts := (*[maxHarts]uintptr)(unsafe.Pointer(&pharts))
	for i := range harts {
		hart := &harts[i]
		*(*uintptr)(unsafe.Pointer(&hart.t)) = uintptr(unsafe.Pointer(&thetasker))
		uharts[i] = uintptr(unsafe.Pointer(hart))
	}
	allcpu := (*slice)(unsafe.Pointer(&thetasker.allcpu))
	*(*uintptr)(unsafe.Pointer(&allcpu.array)) = uintptr(unsafe.Pointer(uharts))
	allcpu.len = 1
	allcpu.cap = maxHarts
	curcpu().exe.set(getg().m)

	// disable interrupts in PLIC
	// BUG: tries to disable all possible interrupts in all possible contexts,
	// accesses reseved address space, can raise Store Access Fault exception
	PLIC := plic.PLIC()
	for ctxid := range PLIC.EN {
		for i := range PLIC.EN[ctxid] {
			PLIC.EN[ctxid][i].Store(0)
		}
		PLIC.TC[ctxid].THR.Store(0)
	}
}

//go:nowritebarrierrec
//go:nosplit
func defaultWrite(fd int, p []byte) int { return len(p) }

const (
	// m.tls fields
	mstatus = 4 // privilege level (MPP>>7), thrSmallCtx
	mepc    = 5 // exception return address
)

const numGPRS = 31

type mOS struct {
	x    [numGPRS]uintptr
	fcsr uintptr
	f    [32]float64
}

func syssetprivlevel(newlevel int) (oldlevel, errno int)

var plicmx cpumtx

//go:nowritebarrierrec
//go:nosplit
func sysirqctl(irq, ctl, ctxid int) (enabled, prio, errno int) {
	PLIC := plic.PLIC()
	if uint(irq) > uint(len(PLIC.PRIO)) {
		errno = 4 // rtos.ErrBadIntNumber
		return
	}
	if uint(ctxid) > uint(len(PLIC.EN)) {
		errno = 6 // rtos.ErrBadIntCtx
	}
	// rtos package ensures valid ctl
	if ctl >= 0 {
		PLIC.PRIO[irq].Store(uint32(ctl))
	}
	rn, bn := irq>>5, irq&31
	switch {
	case ctl >= -1:
		plicmx.lock()
		PLIC.EN[ctxid][rn].SetBit(bn)
		plicmx.unlock()
	case ctl == -2:
		plicmx.lock()
		PLIC.EN[ctxid][rn].ClearBit(bn)
		plicmx.unlock()
	default:
		enabled = int(PLIC.EN[ctxid][rn].Load()) >> bn & 1
		prio = int(PLIC.PRIO[irq].Load())
	}
	return
}
