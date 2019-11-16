// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"embedded/mmio"
	"internal/cpu/cortexm"
	"internal/cpu/cortexm/debug/itm"
	"internal/cpu/cortexm/nvic"
	"internal/cpu/cortexm/scb"
	"internal/cpu/cortexm/scid"
	"unsafe"
)

// for now noos/thumb supports only single CPU

func sev()
func curcpuSleep()
func curcpuSavectx()

//go:nosplit
func (cpu *cpuctx) wakeup() { sev() }

//go:nosplit
func curcpuWakeup() { sev() } // see ARM Errata 563915, STM32F10xx Errata 1.1.2

// ARMv7-M requires at least 4 byte stack alignment so there are two bits
// in saved stack pointer that can be used by tasker. It uses it for:
const (
	thrPrivLevel = 1 << 0 // thread privilege level
	thrSmallCtx  = 1 << 1 // manualy saved context contains only g (R10) register
)

// archnewm setups m's stack
//go:nosplit
func archnewm(m *m) {
	sp := m.g0.stack.hi - unsafe.Sizeof(cortexm.StackFrame{})
	sf := (*cortexm.StackFrame)(unsafe.Pointer(sp))
	sf.PSR = cortexm.T
	sf.PC = funcPC(mstart)
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
func taskerpreinit() {
	*(*uintptr)(unsafe.Pointer(&cpu0.t)) = uintptr(unsafe.Pointer(&thetasker))
	cpu0.exe.set(getg().m)
	allcpu := (*slice)(unsafe.Pointer(&thetasker.allcpu))
	*(*uintptr)(unsafe.Pointer(&allcpu.array)) = uintptr(unsafe.Pointer(&pcpu0))
	allcpu.len = 1
	allcpu.cap = 1

	// setup exception priority levels

	sc := scb.SCB()

	// SVCall and PendSV exceptions should have the lowest priority level. They
	// are both logically part of the thread mode so other exceptions must be
	// able to preempt them. Additionaly, they must not preempt each other.
	sc.PRI_SVCall().Store(cortexm.PrioLo << scb.PRI_SVCalln)
	sc.PRI_PendSV().Store(cortexm.PrioLo << scb.PRI_PendSVn)

	// All other exceptions/interrupts by default have the highest priority.

	mmio.MB() // ensure the exception priorities are set
}

func taskerinit() {}

// syscall handlers (architecture specific code)

//go:nowritebarrierrec
//go:nosplit
func syswrite(fd uintptr, p unsafe.Pointer, n int32) int32 {
	ITM := itm.ITM()
	itmena := ITM.ITMENA()
	port := &ITM.STIM[fd]
	portena := mmio.UM32{&ITM.TER[fd>>5].U32, 1 << (fd & 31)}
	s := (*[1 << 30]byte)(p)
	for i := 0; i < int(n); {
		for port.Bit(0) == 0 {
			if portena.Load() == 0 || itmena.Load() == 0 {
				return n // do not block on disabled port/ITM
			}
		}
		switch m := int(n) - i; {
		case m >= 4:
			port.U32.Store(uint32(s[i]) + uint32(s[i+1])<<8 |
				uint32(s[i+2])<<16 | uint32(s[i+3])<<24)
			i += 4
		case m >= 2:
			p := (*mmio.U16)(unsafe.Pointer(&port.U32))
			p.Store(uint16(s[i]) | uint16(s[i+1])<<8)
			i += 2
		default:
			p := (*mmio.U8)(unsafe.Pointer(&port.U32))
			p.Store(s[i])
			i++
		}
	}
	return n
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
	} else {
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
func sysirqctl(irq, ctl int) (enabled, prio, errno int) {
	if uint(irq) >= irqNum() {
		errno = 4 // rtos.ErrBadIntNumber
		return
	}
	NVIC := nvic.NVIC()
	// rtos package ensures valid ctl
	if ctl >= 0 {
		NVIC.IPR[irq].Store(nvic.IPR(255 - ctl))
	}
	rn, bo := irq>>5, uint(irq&31)
	switch {
	case ctl >= -1:
		NVIC.ISER[rn].Store(1 << bo)
	case ctl == -2:
		NVIC.ICER[rn].Store(1 << bo)
	default:
		enabled = int(NVIC.ISER[irq>>5].Load()) >> bo
	}
	return
}

//go:nowritebarrierrec
//go:nosplit
func sysirqprio(irq int) (prio, errno int) {
	if uint(irq) >= irqNum() {
		return 0, 4 // rtos.ErrBadIRQNumber
	}
	return int(nvic.NVIC().IPR[irq].Load()), 0
}

//go:nowritebarrierrec
//go:nosplit
func syssetirqprio(irq, prio int) (errno int) {
	if uint(irq) >= irqNum() {
		return 4 // rtos.ErrBadIRQNumber
	}
	if uint(prio) > 255 {
		return 5 // rtos.ErrBadIRQPrio
	}
	nvic.NVIC().IPR[irq].Store(nvic.IPR(prio))
	return 0
}

// utils

func cpucontrol() uint32
func setcpucontrol(ctrl uint32)

//go:nosplit
func irqNum() uint {
	n := uint(scid.SCID().INTLINESNUM().Load()+1) * 32
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
