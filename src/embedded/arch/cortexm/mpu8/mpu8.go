// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mpu8 provides interface to the ARMv7-M Memory Protection Unit.
package mpu8

import (
	"embedded/mmio"
	"unsafe"
)

type regs struct {
	typ  mmio.R32[uint32]
	ctrl mmio.R32[uint32]
	rnr  mmio.R32[uint32]
	r    [4]struct {
		bar mmio.R32[uint32]
		lar mmio.R32[uint32]
	}
	mair [2]mmio.R32[uint32]
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

// Region attributes for SetBase function.
const (
	XN int8 = 1 << 0 // Execution not permitted

	Arw__ int8 = 0 << 1 // Read/write by privileged code only
	Arwrw int8 = 1 << 1 // Read/write by any privilege level
	Ar___ int8 = 2 << 1 // Read-only by privileged code only
	Ar_r_ int8 = 3 << 1 // Read-only by any privilege level

	SHNONE  int8 = 0 << 3 // Non-shareable
	SHOUTER int8 = 2 << 3 // Outer shareable
	SHINNER int8 = 3 << 3 // Inner Shareable
)

// SetBase sets the base address and the attributes of the selected region.
func SetBase(base uintptr, attr int8) {
	p().r[0].bar.Store(uint32(base)&^0x1f | uint32(attr)&0x1f)
}

func Base() (base uintptr, attr int8) {
	v := p().r[0].bar.Load()
	base = uintptr(v &^ 0x1f)
	attr = int8(v & 0x1f)
	return
}

func SetLimit(limit uintptr, attrid int, en bool) {
	enbit := *(*uint32)(unsafe.Pointer(&en))
	p().r[0].lar.Store(uint32(limit)&^0x1f | uint32(attrid&7)<<1 | enbit)
}
