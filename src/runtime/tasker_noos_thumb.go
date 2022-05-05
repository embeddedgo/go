// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"embedded/mmio"
	"internal/abi"
	"internal/cpu/cortexm"
	"internal/cpu/cortexm/debug/itm"
	"internal/cpu/cortexm/mpu"
	"internal/cpu/cortexm/nvic"
	"internal/cpu/cortexm/scb"
	"internal/cpu/cortexm/scid"
	"unsafe"
)

// for now noos/thumb supports only single CPU

func sev()
func curcpuSleep()
func curcpuSavectxSched()
func curcpuSavectxCall() {} // all registars saved on caller's stack

//go:nosplit
func cpuid() int {
	// for now only single CPU is supported (see also identcurcpu, osinit)
	return 0
}

//go:nosplit
func curcpuWakeup() { sev() } // see ARM Errata 563915, STM32F10xx Errata 1.1.2

//go:nosplit
func (cpu *cpuctx) newwork() {
	scb.SCB().ICSR.Store(scb.PENDSVSET)
	sev()
}

//go:nosplit
func curcpuSchedule() {
	// Can't run the scheduler directly in the system call because the SVCall
	// has higher priority than some interrupts. If there is no any thread to
	// run the scheduler will sleep in a loop (there is no idle threads) at
	// SVCall priority level thereby blocking the lower priority interrupts.
	//
	// Instead, we trigger PendSV (which has priority lower that any other
	// interrupt) to run the scheduler and rely on the exception tail-chaining
	// to don't execute any instruction after SWI until the scheduler does its
	// job.
	//
	// Caution! You can't rely on tail-chaining in case of debuging.
	curcpu().schedule = true
	scb.SCB().ICSR.Store(scb.PENDSVSET)
}

// ARMv7-M requires at least 4 byte stack alignment so there are two bits
// in saved stack pointer that can be used by tasker. It uses it for:
const (
	thrPrivLevel = 1 << 0 // thread privilege level
	thrSmallCtx  = 1 << 1 // context saved in m contains only g (R10) register
)

// archnewm setups m's stack
//go:nosplit
func archnewm(m *m) {
	sp := m.g0.stack.hi - unsafe.Sizeof(cortexm.StackFrame{})
	sf := (*cortexm.StackFrame)(unsafe.Pointer(sp))
	sf.PSR = cortexm.T
	sf.PC = abi.FuncPCABI0(mstart)
	m.tls[msp] = sp | thrSmallCtx // small ctx
	m.tls[mer] = cortexm.ExcReturnBase | cortexm.ExcReturnNoFPU |
		cortexm.ExcReturnPSP
	m.libcall.fn = uintptr(unsafe.Pointer(m.g0))
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

	// setup exception priority levels

	SCB := scb.SCB()

	// enable fault handlers
	SCB.SHCSR.SetBits(scb.MEMFAULTENA | scb.BUSFAULTENA | scb.USGFAULTENA)

	// division by zero will cause the UsageFault
	SCB.CCR.SetBits(scb.DIV_0_TRP)

	// set PendSV and SVCall priorities according to description in rtos package
	SCB.SHPR2.StoreBits(scb.PRI_SVCall, (4<<5)<<scb.PRI_SVCalln)
	SCB.SHPR3.StoreBits(scb.PRI_PendSV, 255<<scb.PRI_PendSVn)

	// All other exceptions/interrupts by default have the highest priority.

	// use MPU if available to catch nil pointer dereferences (need 4 regions)
	if _, d, _ := mpu.Type(); d >= 4 {
		// Bellow there is the MPU configuration that more or less corresponds
		// to the default CPU behavior, without MPU enabled.
		//
		// All regions starts at address 0x00000000. The SIZE and SRD
		// (sub-region disabled) fields are used to set the region address
		// ranges.
		//
		// The first 64 bytes of the code region that corresponds to the first
		// 16 exception vectors are set inaccessible to catch nil pointer
		// dereferences. The code region is declared read/write because some
		// MCUs use normal memory access to program Flash.
		//
		// Tha RAM region is configured as shareable (usually shared with DMA).
		// Shareable regions are by default not cacheable. If you enable L1
		// cache in Cortex-M7 set the acc.SIWT bit so the RAM will be cacheable
		// in write-through mode. WT mode degrades performance (not as much as
		// you may think) but allows to avoid cache maintenance operations which
		// are problematic in case of Cortex-M7.
		var (
			noacc  = mpu.A____
			code   = mpu.Arwrw | mpu.C
			ram    = mpu.Arwrw | mpu.TEX(1) | mpu.C | mpu.B | mpu.S
			periph = mpu.Arwrw | mpu.B | mpu.S | mpu.XN
		)
		mpu.SetRegion(0x00000000|mpu.VALID|0, mpu.ENA|mpu.SIZE(29)|code)
		mpu.SetRegion(0x00000000|mpu.VALID|1, mpu.ENA|mpu.SIZE(6)|noacc)
		mpu.SetRegion(
			0x00000000|mpu.VALID|2,
			mpu.ENA|mpu.SIZE(32)|mpu.SRD(0b10011011)|periph,
		)
		mpu.SetRegion(
			0x00000000|mpu.VALID|3,
			mpu.ENA|mpu.SIZE(32)|mpu.SRD(0b11100101)|ram,
		)
		mmio.MB() // ensure any previous memory access is done before enable MPU
		mpu.Set(mpu.ENABLE | mpu.PRIVDEFENA)
	}

	// ensure everything is set before any subsequent memory access
	mmio.MB()
}

//go:nowritebarrierrec
//go:nosplit
func defaultWrite(fd int, p []byte) int {
	ITM := itm.ITM()
	port := &ITM.STIM[fd]
	portena := mmio.UM32{&ITM.TER[fd>>5].U32, 1 << (fd & 31)}
	for i := 0; i < len(p); {
		for port.LoadBit(0) == 0 {
			if portena.Load() == 0 || ITM.TCR.LoadBits(itm.ITMENA) == 0 {
				return len(p) // do not block on disabled port/ITM
			}
		}
		switch m := len(p) - i; {
		case m >= 4:
			port.U32.Store(uint32(p[i]) + uint32(p[i+1])<<8 |
				uint32(p[i+2])<<16 | uint32(p[i+3])<<24)
			i += 4
		case m >= 2:
			port16 := (*mmio.U16)(unsafe.Pointer(&port.U32))
			port16.Store(uint16(p[i]) | uint16(p[i+1])<<8)
			i += 2
		default:
			port8 := (*mmio.U8)(unsafe.Pointer(&port.U32))
			port8.Store(p[i])
			i++
		}
	}
	return len(p)
}

// syscalls not used by runtime

//go:nowritebarrierrec
//go:nosplit
func syssetprivlevel(newlevel int) (oldlevel, errno int) {
	// this code requires thrPrivLevel == 1
	const check byte = (thrPrivLevel - 1) * 256

	ctrl := cpucontrol()
	oldlevel = int(ctrl | 1)
	if uint(newlevel) <= 1 {
		setcpucontrol(ctrl&^1 | uint32(newlevel))
	} else if newlevel > 0 {
		errno = 2 // rtos.ErrBadPrivLevel
	}
	return
}

//go:nowritebarrierrec
//go:nosplit
func sysirqenabled(irq int) (enabled, errno int) {
	if uint(irq) >= irqNum() {
		return 0, 4 // rtos.ErrBadIRQNumber
	}
	return int(nvic.NVIC().ISER[irq>>5].Load()) >> uint(irq&31), 0
}

//go:nowritebarrierrec
//go:nosplit
func sysirqctl(irq, ctl, ctxid int) (enabled, prio, errno int) {
	if uint(irq) >= irqNum() {
		errno = 4 // rtos.ErrBadIntNumber
		return
	}
	if uint(ctxid) > 0 {
		errno = 6 // rtos.ErrBadIntCtx
	}
	NVIC := nvic.NVIC()
	// rtos package ensures valid ctl
	if ctl >= 0 {
		NVIC.IPR[irq].Store(nvic.IPR(255 - ctl))
	}
	rn, bn := irq>>5, uint(irq&31)
	switch {
	case ctl >= -1:
		NVIC.ISER[rn].Store(1 << bn)
	case ctl == -2:
		NVIC.ICER[rn].Store(1 << bn)
	default:
		enabled = int(NVIC.ISER[irq>>5].Load()) >> bn & 1
		prio = 255 - int(NVIC.IPR[irq].Load())
	}
	return
}

// utils

func cpucontrol() uint32
func setcpucontrol(ctrl uint32)

//go:nosplit
func irqNum() uint {
	n := uint(scid.SCID().ICTR.LoadBits(scid.INTLINESNUM)+1) * 32
	if n > 496 {
		n = 496
	}
	return n
}

// m.tls fields

const msp = 4
const mer = 5

// Use libcall, libcallpc, libcallsp, libcallg, syscall, vdsoSP, vdsoPC and mOS
// to save the second part of thread context. We do not save it on the gorutine
// stack to avoid waste of memory (need to increasing stack guard for any
// gorutine stack and there are much more gorutines than threads).
//
// We realy do not want to add another 96 bytes to the m so we take the trouble
// to use these unused fields for our needs. The following constant declarations
// are compile-time tests to ensure the fields were not changed or splitted.
const (
	_mLibcallAlign       = -(unsafe.Offsetof(m{}.libcall) & 3)
	_mLibcallSize        = int8((unsafe.Sizeof(m{}.libcall) - 6*4) * 129)
	_mLibcallFn          = int8(unsafe.Offsetof(m{}.libcall.fn) * 129)
	_mLibcallLibcallpc   = int8((unsafe.Offsetof(m{}.libcallpc) - unsafe.Offsetof(m{}.libcall) - 6*4) * 129)
	_mLibcallpcSize      = int8((unsafe.Sizeof(m{}.libcallpc) - 4) * 129)
	_mLibcallpcLibcallsp = int8((unsafe.Offsetof(m{}.libcallsp) - unsafe.Offsetof(m{}.libcallpc) - 4) * 129)
	_mLibcallspSize      = int8((unsafe.Sizeof(m{}.libcallsp) - 4) * 129)
	_mLibcallspLibcallg  = int8((unsafe.Offsetof(m{}.libcallg) - unsafe.Offsetof(m{}.libcallsp) - 4) * 129)
	_mLibcallgSize       = int8((unsafe.Sizeof(m{}.libcallg) - 4) * 129)
	_mLibcallgSyscall    = int8((unsafe.Offsetof(m{}.syscall) - unsafe.Offsetof(m{}.libcallg) - 4) * 129)
	_mSyscallSize        = int8((unsafe.Sizeof(m{}.syscall) - 6*4) * 129)
	_mSyscallVdsosp      = int8((unsafe.Offsetof(m{}.vdsoSP) - unsafe.Offsetof(m{}.syscall) - 6*4) * 129)
	_mVdsospSize         = int8((unsafe.Sizeof(m{}.vdsoSP) - 4) * 129)
	_mVdsospVdsopc       = int8((unsafe.Offsetof(m{}.vdsoPC) - unsafe.Offsetof(m{}.vdsoSP) - 4) * 129)
	_mVdsopcSize         = int8((unsafe.Sizeof(m{}.vdsoPC) - 4) * 129)
	_mVdsopcMos          = int8((unsafe.Offsetof(m{}.mOS) - unsafe.Offsetof(m{}.vdsoPC) - 4) * 129)
	_mSize               = int8((unsafe.Offsetof(m{}.mOS) - unsafe.Offsetof(m{}.libcall) + unsafe.Sizeof(m{}.mOS) - 24*4) * 129)
)

type mOS [7]uint32
