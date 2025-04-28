// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"embedded/arch/cortexm/mpu"
	"embedded/arch/cortexm/mpu/mpu7"
	"embedded/arch/cortexm/mpu/mpu8"
	"embedded/mmio"
	"internal/abi"
	"internal/cpu/armm"
	"internal/cpu/armm/cmt"
	"internal/cpu/armm/debug/itm"
	"internal/cpu/armm/fpu"
	"internal/cpu/armm/nvic"
	"internal/cpu/armm/pft"
	"internal/cpu/armm/scb"
	"internal/cpu/armm/scid"
	"internal/goarch"
	"internal/runtime/atomic"
	"unsafe"
)

func sev()
func isb()
func curcpuSleep()
func curcpuSavectxSched()
func curcpuSavectxCall() {} // all registars saved on caller's stack
func preemptOrWakeup(cpuid int)

//go:nosplit
func curcpuWakeup() {
	preemptOrWakeup(-1)
}

//go:nosplit
func (cpu *cpuctx) newwork() {
	preemptOrWakeup(int(cpu.gh.goid))
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
//
//go:nosplit
func archnewm(m *m) {
	sp := m.g0.stack.hi - unsafe.Sizeof(armm.StackFrame{})
	sf := (*armm.StackFrame)(unsafe.Pointer(sp))
	sf.PSR = armm.T
	sf.PC = abi.FuncPCABI0(mstart)
	m.tls[msp] = sp | thrSmallCtx | thrPrivLevel
	m.tls[mer] = armm.ExcReturnSecureThreadPSP
	m.libcall.fn = uintptr(unsafe.Pointer(m.g0))
}

//go:nosplit
func leadingZeros32(x uint32) uint {
	var n uint
	for x != 0 {
		x >>= 1
		n++
	}
	return 32 - n
}

var runOtherCPUs atomic.Bool

//go:nowritebarrierrec
//go:nosplit
func taskerinit(stackStart, stackEnd uintptr) {
	allcpu := (*notInHeapSlice)(unsafe.Pointer(&thetasker.allcpu))
	allcpu.len = int(ncpu)
	allcpu.cap = int(ncpu)
	allcpu.array = (*notInHeap)(noosRawAlloc(goarch.PtrSize*uintptr(ncpu), goarch.PtrSize))
	cpus := noosRawAlloc(unsafe.Sizeof(cpuctx{})*uintptr(ncpu), unsafe.Alignof(cpuctx{}))
	stackSize := (stackEnd - stackStart) / uintptr(ncpu)

	// In case of multiple CPUs, the CPU0 main stack is at the top of the stack
	// space. The CPU1 stack is just below the CPU0 stack and so on.
	for i := range thetasker.allcpu {
		stackStart = stackEnd - stackSize

		cpu := (*cpuctx)(cpus)
		cpu.t = &thetasker
		cpu.gh.goid = uint64(i) // cpuid
		cpu.gh.stack.lo = stackStart
		cpu.gh.stack.hi = stackEnd
		cpu.gh.stackguard0 = stackStart + stackGuard
		cpu.gh.stackguard1 = stackStart + stackGuard
		setMNoWB(&cpu.gh.m, &cpu.mh)
		setGNoWB(&cpu.mh.g0, &cpu.gh)
		setGNoWB(&cpu.mh.gsignal, &cpu.gh)

		thetasker.allcpu[i] = cpu
		cpus = unsafe.Add(cpus, unsafe.Sizeof(cpuctx{}))
		stackEnd = stackStart
	}

	// Now the target identcurcpu should work so other CPUs can enter tasker.
	runOtherCPUs.Store(true)
	sev()
}

const debugBusFault = false

// initCPU is called by every CPU in the system, very early, even before BSS and
// data segments are initialized.
//
//go:nowritebarrierrec
//go:nosplit
func initCPU(vectors uintptr) {
	SCB := scb.SCB()

	// Set VTOR (required mainly if the boot process is based on a bootloader)
	SCB.VTOR.Store(scb.VTOR(vectors))
	// Enable fault handlers
	SCB.SHCSR.SetBits(scb.MEMFAULTENA | scb.BUSFAULTENA | scb.USGFAULTENA)
	// Division by zero will causes the UsageFault.
	SCB.CCR.SetBits(scb.DIV_0_TRP)
	// Set PendSV and SVCall priorities according to description in rtos package
	SCB.SHPR2.StoreBits(scb.PRI_SVCall, (4<<5)<<scb.PRI_SVCalln)
	SCB.SHPR3.StoreBits(scb.PRI_PendSV, 255<<scb.PRI_PendSVn)

	// All other exceptions/interrupts by default have the highest priority.

	// Enable FPU.
	FPU := fpu.FPU()
	FPU.CPACR.Store(fpu.CP10 | fpu.CP11)
	if FPU.CPACR.LoadBits(fpu.CP10|fpu.CP11) == fpu.CP10|fpu.CP11 {
		FPU.FPCCR.Store(fpu.LSPEN | fpu.ASPEN)
	}

	if debugBusFault {
		// Disable buffering to make bus faults synchronous.
		scid.SCID().ACTLR.SetBits(scid.DISDEFWBUF)
	} else {
		// Enable L1 cache if present and not enabled before (may already be
		// enabled in case of the soft reboot)
		PFT := pft.PFT()
		CMT := cmt.CMT()
		clidr := PFT.CLIDR.Load()
		cc := SCB.CCR.Load()
		if clidr&pft.CL1I != 0 && cc&scb.IC == 0 {
			// L1 instruction cache implemented. Must invalidate it before use.
			CMT.ICIALLU.Store(0)
			mmio.MB()
			isb()
			SCB.CCR.SetBits(scb.IC)
			mmio.MB()
			isb()
		}
		if clidr&pft.CL1D != 0 && cc&scb.DC == 0 {
			// L1 data cache implemented. Must invalidate it before use.
			PFT.CSSELR.Store(0) // select L1 cache size info
			mmio.MB()
			csi := PFT.CCSIDR.Load() // load L1 cache size info

			maxset := uint32(csi&pft.NumSets) >> pft.NumSetsn
			maxway := uint32(csi&pft.Associativity) >> pft.Associativityn
			log2bpl := uint(csi&pft.LineSize)>>pft.LineSizen + 4
			wayshift := leadingZeros32(maxway)

			for set := uint32(0); set <= maxset; set++ {
				for way := uint32(0); way <= maxway; way++ {
					CMT.DCISW.U32.Store(way<<wayshift | set<<log2bpl)
				}
			}
			mmio.MB()
			isb()
			SCB.CCR.SetBits(scb.DC)
			mmio.MB()
			isb()
		}
	}

	// Use MPU if available to catch the bad pointer dereferences. We configure
	// the MPU to mimic the default CPU behavior, without the MPU enabled. The
	// first 64 bytes of the memory are configured inaccessible in the user mode
	// (or both modes for ARMv7-M) to catch the bad pointer dereferences.

	partno := SCB.CPUID.LoadBits(scb.PartNo) >> scb.PartNon
	_, dregn, _ := mpu.Type()
	switch {
	case mpu.State()&mpu.ENABLE != 0:
		// enabled before
		goto skipMPU
	case partno&0xf00 == 0xd00: // ARMv8-M
		if dregn < 6 {
			goto skipMPU
		}
		const (
			Periph   = 0
			ExtDev   = 1
			NormalWT = 2
			NormalWB = 3
		)
		mpu8.SetAttr03(
			mpu8.Device, mpu8.DnGnRE, // Periph
			mpu8.Device, mpu8.DnGnRnE, // ExtDev
			mpu8.NormalWT, mpu8.NormalWT, // NormalWT
			mpu8.NormalWB, mpu8.NormalWB, // NormalWB
		)

		// The first 64 bytes of the memory are inaccessible in the user mode.
		mpu.Select(0)
		mpu8.SetBas(0x0000_0000, mpu8.Arw__)
		mpu8.SetLim(0x0000_003f, NormalWT, true)

		// The code region occupies the first 512 MiB.
		mpu.Select(1)
		mpu8.SetBas(0x0000_0040, mpu8.Arwrw)
		mpu8.SetLim(0x1fff_ffff, NormalWT, true)

		// First RAM region, 512 MiB.
		mpu.Select(2)
		mpu8.SetBas(0x2000_0000, mpu8.Arwrw|mpu8.SI)
		mpu8.SetLim(0x3fff_ffff, NormalWB, true)

		// Peripherals.
		mpu.Select(3)
		mpu8.SetBas(0x4000_0000, mpu8.Arwrw|mpu8.XN)
		mpu8.SetLim(0x5fff_ffff, Periph, true)

		// Second RAM region, 1 GiB.
		mpu.Select(4)
		mpu8.SetBas(0x6000_0000, mpu8.Arwrw|mpu8.SI)
		mpu8.SetLim(0x9fff_ffff, NormalWB, true)

		// External device region, 1 GiB
		mpu.Select(5)
		mpu8.SetBas(0xa000_0000, mpu8.Arwrw|mpu8.XN)
		mpu8.SetLim(0xdfff_ffff, ExtDev, true)
	default: // ARMv7-M
		if dregn < 4 {
			goto skipMPU
		}
		const (
			noacc  = mpu7.A____
			code   = mpu7.Arwrw | mpu7.C                      // normal WT
			ram    = mpu7.Arwrw | mpu7.TEX1 | mpu7.C | mpu7.B // normal WB+WA
			periph = mpu7.Arwrw | mpu7.B | mpu7.XN            // device
		)

		// The peripheral region covers adresses not covered by other regions.
		mpu7.SetRegion(mpu7.VALID|0, mpu7.ENA|mpu7.SIZE(32)|periph)

		// The code region occupies the first 512 MiB.
		mpu7.SetRegion(mpu7.VALID|1, mpu7.ENA|mpu7.SIZE(29)|code)

		// The first 64 bytes of the code region are inaccessible.
		mpu7.SetRegion(mpu7.VALID|2, mpu7.ENA|mpu7.SIZE(6)|noacc)

		// RAM region occupies 512 MiB @ 0x2000_0000 and 1 GiB @ 0x6000_0000.
		mpu7.SetRegion(mpu7.VALID|3, mpu7.ENA|mpu7.SIZE(32)|mpu7.SRD(0b11100101)|ram)

	}
	mmio.MB()
	mpu.Set(mpu.ENABLE | mpu.PRIVDEFENA)
	mmio.MB()
skipMPU:
}

//go:nowritebarrierrec
//go:nosplit
func defaultWrite(fd int, p []byte) int {
	ITM := itm.ITM()
	port := &ITM.STIM[fd]
	portena := mmio.UM32{&ITM.TER[fd>>5].U32, 1 << (fd & 31)}
	for i := 0; i < len(p); {
		for port.LoadBits(1<<0) == 0 {
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

//go:noescape
func syscachemaint(op int, p unsafe.Pointer, size int)

//go:noescape
func sysreset(level int, addr unsafe.Pointer) bool

//go:nowritebarrierrec
//go:nosplit
func syssetprivlevel(newlevel int) (oldlevel, errno int) {
	// this code requires thrPrivLevel == 1
	const check byte = (thrPrivLevel - 1) * 256

	ctrl := cpucontrol()
	oldlevel = int(ctrl & 1)
	if uint(newlevel) <= 1 {
		setcpucontrol(ctrl&^1 | uint32(newlevel))
	} else if newlevel > 0 {
		errno = 2 // rtos.ErrBadPrivLevel
	}
	return
}

//go:nowritebarrierrec
//go:nosplit
func sysirqctl(irq, ctl, ctxid int) (enabled, prio, errno int) {
	if uint(irq) >= irqNum() {
		errno = 4 // rtos.ErrBadIntNumber
		return
	}
	if ctxid != -1 {
		errno = 6 // rtos.ErrBadIntCtx
		return
	}
	NVIC := nvic.NVIC()
	rn, bn := irq>>5, uint(irq&31)
	// rtos package ensures valid ctl
	switch {
	case ctl >= -1: // enable IRQ
		if ctl >= 0 {
			// ctl contains new priority
			NVIC.IPR[irq].Store(nvic.IPR(255 - ctl))
		}
		NVIC.ISER[rn].Store(1 << bn)
	case ctl == -2: // disable IRQ
		NVIC.ICER[rn].Store(1 << bn)
	default: // -3, IRQ status
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

// Exceptions handled in runtime
func svcallHandler()
func pendsvHandler()
func reservedHandler()
func nmiHandler()
func hardfaultHandler()
func memmanageHandler()
func busfaultHandler()
func usagefaultHandler()
func securefaultHandler()
func debugmonHandler()
