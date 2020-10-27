// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package systim iplements tickless system timer using the RISC-V CLINT timer.
package systim

import (
	"embedded/rtos"
	"internal/cpu/riscv/clint"
	_ "unsafe" // for linkname
)

var timerHz uint64

// Setup setups CLINT timer to work as sytem timer.
//  clkhz - frequency of CLINT timer clock source
func Setup(clkhz int64) {
	timerHz = uint64(clkhz)
	rtos.SetSystemTimer(nanotime, setAlarm)
}

//go:nosplit
func nanotime() int64 {
	return int64(mulDiv(clint.CLINT().MTIME.U64.Load(), 1e9, timerHz))
}

//go:nosplit
func setAlarm(ns int64) {
	timecmp := uint64(1<<64 - 1)
	if ns >= 0 {
		timecmp = mulDivUp(uint64(ns), timerHz, 1e9)
	}
	clint.CLINT().MTIMECMP[cpuid()].U64.Store(timecmp)
}

//go:nosplit
func mulDiv(x, m, d uint64) uint64 {
	divx := x / d
	modx := x - divx*d
	divm := m / d
	modm := m - divm*d
	return divx*m + modx*divm + modx*modm/d
}

//go:nosplit
func mulDivUp(x, m, d uint64) uint64 {
	o := d - 1
	divx := (x + o) / d
	modx := x - divx*d
	divm := (m + o) / d
	modm := m - divm*d
	return divx*m + modx*divm + (modx*modm+o)/d
}

//go:linkname cpuid runtime.cpuid
func cpuid() int
