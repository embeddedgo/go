// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/abi"
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
	m.tls[mepc] = abi.FuncPCABI0(mstart)
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

	// reset PLIC to known state
	// BUG: tries to reset all possible interrupts and all possible contexts,
	// can raise Store Access Fault exception
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

const (
	numGPRS = 31
	numFPRS = 32 + 1 // include fcsr
)

type mOS struct {
	x    [numGPRS]uintptr
	fcsr uintptr
	f    [numFPRS - 1]float64
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

var exceptionNames = [...]string{
	0:  "instruction address misaligned",
	1:  "instruction access fault",
	2:  "illegal instruction",
	3:  "breakpoint",
	4:  "load address misaligned",
	5:  "load access fault",
	6:  "store/AMO address misaligned",
	7:  "store/AMO access fault",
	8:  "environment call from U-mode",
	9:  "environment call from S-mode",
	10: "reserved",
	11: "environment call from M-mode",
	12: "instruction page fault",
	13: "load page fault",
	14: "reserved",
	15: "store/AMO page fault",
}

func printReg(name string, u uintptr, end string) {
	const dig = "0123456789abcdef"
	var buf [19]byte
	buf[0] = ' '
	buf[1] = '0'
	buf[2] = 'x'
	for i := 0; i < 16; i++ {
		buf[18-i] = dig[u&15]
		u >>= 4
	}
	print(name, string(buf[:]), end)
}

// keep lr, a0, mstatus, mepc, mie and rs order in sync with asm_riscv64.h
//go:nosplit
func fatalException(rs [numGPRS - 4]uintptr, mcause int, sp, lr, a0, mstatus, mepc, mie uintptr) {
	mode := "handler"
	gp := getg()
	g := gp
	if mepc&1 != 0 {
		mode = "thread"
		g = gp.sched.g.ptr()
		sp = gp.sched.sp
		mepc &^= 1
	}
	print("\n\nfatal exception in ", mode, " mode at ", hex(mepc), ": ")
	if mcause < len(exceptionNames) {
		print(exceptionNames[mcause])
	} else {
		print("reserved (", mcause, ")")
	}
	print("\n\n")

	/*
		print("ZERO                  \t")
		printReg("LR ", lr, "\n")
		printReg("SP ", sp, "\t")
		printReg("GP ", rs[0], "\n")
		printReg("g  ", uintptr(unsafe.Pointer(g)), "\t")
		printReg("T0 ", rs[1], "\n")
		printReg("T1 ", rs[2], "\t")
		printReg("T2 ", rs[3], "\n")
		printReg("S0 ", rs[4], "\t")
		printReg("S1 ", rs[5], "\n")
		printReg("A0 ", a0, "\t")
		printReg("A1 ", rs[6], "\n")
		printReg("A2 ", rs[7], "\t")
		printReg("A3 ", rs[8], "\n")
		printReg("A4 ", rs[9], "\t")
		printReg("A5 ", rs[10], "\n")
		printReg("A6 ", rs[11], "\t")
		printReg("A7 ", rs[12], "\n")
		printReg("S2 ", rs[13], "\t")
		printReg("S3 ", rs[14], "\n")
		printReg("S4 ", rs[15], "\t")
		printReg("S5 ", rs[16], "\n")
		printReg("S6 ", rs[17], "\t")
		printReg("S7 ", rs[18], "\n")
		printReg("S8 ", rs[19], "\t")
		printReg("S9 ", rs[20], "\n")
		printReg("S10", rs[21], "\t")
		printReg("S11", rs[22], "\n")
		printReg("T3 ", rs[23], "\t")
		printReg("T4 ", rs[24], "\n")
		printReg("T5 ", rs[25], "\t")
		printReg("tmp", rs[26], "\n\n")
	*/

	traceback1(mepc, sp, lr, g, _TraceTrap)
	crash()
}
