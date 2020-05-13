// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

func cpuid() int // consider remove this function, don't use it in isr(), taskerinit()

const maxHarts = 2

var (
	harts  [maxHarts]cpuctx
	pharts [maxHarts]*cpuctx
)

//go:nowritebarrierrec
//go:nosplit
func taskerinit() {
	if cpuid() == 0 {
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
		getcpuctx().exe.set(getg().m)
	}
}

// m.tls fields

const (
	mstatus = 4 // privilege level, small context flag
	mepc    = 5

	thrSmallCtx = 1 << 4 // context saved in mOS contains only LR, SP, g registers

	numGPRS = 27 // must be in sync with saveRegs function

)

type mOS struct {
	x    [numGPRS]uintptr
	f    [32]float64
	fcsr uint64
}

func archnewm(m *m) {
	m.tls[mstatus] = thrSmallCtx
	m.tls[mepc] = funcPC(mstart)
	m.x[1] = m.g0.stack.hi                 // SP
	m.x[2] = uintptr(unsafe.Pointer(m.g0)) // g
}

func curcpuSleep()                                         { breakpoint() }
func curcpuSavectxSched()                                  { breakpoint() }
func curcpuWakeup()                                        { breakpoint() }
func curcpuSchedule()                                      { breakpoint() }
func curcpuSavectxCall()                                   { breakpoint() }
func (cpu *cpuctx) wakeup()                                { breakpoint() }
func sysirqctl(irq, ctl int) (enabled, prio, errno int)    { breakpoint(); return }
func syssetprivlevel(newlevel int) (oldlevel, errno int)   { breakpoint(); return }
func syswrite(fd uintptr, p unsafe.Pointer, n int32) int32 { breakpoint(); return -1 }
