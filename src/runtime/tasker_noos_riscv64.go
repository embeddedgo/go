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
			uharts[i] = uintptr(unsafe.Pointer(&harts[i]))
		}
		allcpu := (*slice)(unsafe.Pointer(&thetasker.allcpu))
		*(*uintptr)(unsafe.Pointer(&allcpu.array)) = uintptr(unsafe.Pointer(uharts))
		allcpu.len = 1
		allcpu.cap = maxHarts
		getcpuctx().exe.set(getg().m)
	}
}

// m.tls fields

const mer = 5

type mOS struct {
	x    [31]uint64 // x1-x31
	fcsr uint64
	f    [32]float64
}

func curcpuSleep()          { breakpoint() }
func curcpuSavectxSched()   { breakpoint() }
func curcpuWakeup()         { breakpoint() }
func archnewm(m *m)         { breakpoint() }
func curcpuSchedule()       { breakpoint() }
func curcpuSavectxCall()    { breakpoint() }
func (cpu *cpuctx) wakeup() { breakpoint() }
