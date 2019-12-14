// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package systim iplements ticking system timer using the ARMv7-M SysTick
// peripheral. The SysTick limitations do not allow to implement precise
// tickless system timer.
//
// This implementation uses non-atomic loads of 64-bit counter and for this
// reason it can be used only in uniprocessor system and only if the Nanotime
// function is called with priority lower or same as the SysTick exception.
package systim

import (
	"embedded/rtos"
	"internal/cpu/cortexm/scb"
	"internal/cpu/cortexm/systick"
	_ "unsafe"
)

var systim struct {
	reloadcnt int64 // must be the first field
	periodns  int64
}

// Nanotime returns
//go:nosplit
func Nanotime() int64 {
	st := systick.SYSTICK()
	var downtick uint32
	reloadcnt := systim.reloadcnt // non-atomic!
	for {
		downtick = uint32(st.CURRENT().Load())
		reloadcnt1 := systim.reloadcnt // non-atomic!
		if reloadcnt1 == reloadcnt {
			break
		}
		reloadcnt = reloadcnt1
	}
	// the SysTick asserts an exception when the CURRENT changes from 1 to 0
	periodtick := uint32(st.RELOAD().Load() + 1)
	tick := uint32(0)
	if downtick != 0 {
		tick = periodtick - downtick
	}
	return systim.periodns*reloadcnt +
		systim.periodns*int64(tick)/int64(periodtick)
}

// Setup setups Cortex-M SysTick timer to work as sytem timer.
//  periodns - number of nanoseconds between SysTick interrupts,
//  clkhz    - frequency of SysTick clock source,
//  external - clock source (true: external clock, false: CPU clock).
// Setup must be run in privileged mode.
func Setup(periodns, clkhz int64, external bool) {
	st := systick.SYSTICK()
	if periodns <= 0 || clkhz <= 0 {
		st.CSR.ClearBits(systick.ENABLE | systick.TICKINT)
		systim.periodns = 0
		return
	}
	// Set SysTick exception priority accortding to rtos package.
	scb.SCB().PRI_SysTick().Store(255 - rtos.IntPrioSysTimer)
	systim.periodns = periodns
	systim.reloadcnt = 1 // ensure that nanotime never return zero
	periodtick := uint32((periodns*clkhz + 5e8) / 1e9)
	st.RELOAD().Store(systick.RVR(periodtick - 1))
	st.CURRENT().Store(0)

	cfg := systick.ENABLE | systick.TICKINT
	if !external {
		cfg |= systick.CLKSOURCE
	}
	st.CSR.Store(systick.ENABLE | systick.TICKINT)
}
