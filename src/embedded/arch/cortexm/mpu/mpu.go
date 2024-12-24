// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mpu provides an interface to the common functionalities of ARMv7-M
// and ARMv8-M Memory Protection Unit.
package mpu

import (
	"embedded/mmio"
	"unsafe"
)

type regs struct {
	typ  mmio.R32[uint32]
	ctrl mmio.R32[uint32]
	rnr  mmio.R32[uint32]
}

func p() *regs { return (*regs)(unsafe.Pointer(uintptr(0xE000ED90))) }

// Type returns information about the MPU unit:
// i - number of supported instruction regions,
// d - number of supported data regions.
// s - true if separate instruction and data regions are supported.
func Type() (i, d int, s bool) {
	typ := p().typ.Load()
	i = int(typ>>16) & 0xff
	d = int(typ>>8) & 0xff
	s = (typ&1 != 0)
	return
}

// MPU control flags
const (
	// If ENABLE is set MPU is enabled.
	ENABLE uint32 = 1 << 0
	// If HFNMIENA is not set the MPU will be disabled during HardFault, NMI
	// and FAULTMASK handlers.
	HFNMIENA uint32 = 1 << 1
	// If PRIVDEF is set the default memory map is used as background region for
	// privileged software access.
	PRIVDEFENA uint32 = 1 << 2
)

// Set sets the flags specified by fl.
func Set(fl uint32) { p().ctrl.SetBits(uint32(fl)) }

// Clear clears the flags specified by fl.
func Clear(fl uint32) { p().ctrl.ClearBits(uint32(fl)) }

// State returns the current state.
func State() uint32 { return p().ctrl.Load() }

// Select selects the region number n.
func Select(n int) { p().rnr.Store(uint32(n)) }

// Current returns the current region number.
func Current() int { return int(p().rnr.Load()) }
