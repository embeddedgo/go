// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mpu8 provides an interface to the ARMv8-M Memory Protection Unit.
package mpu8

import (
	"embedded/mmio"
	"unsafe"
)

type regs struct {
	_ [3]uint32
	r [4]struct {
		bar mmio.R32[uint32]
		lar mmio.R32[uint32]
	}
	_    uint32
	mair [2]mmio.R32[uint32]
}

func p() *regs { return (*regs)(unsafe.Pointer(uintptr(0xE000ED90))) }

// Region permission and shareability attributes for SetBase function.
const (
	XN uint8 = 1 << 0 // Execution not permitted

	Arw__ uint8 = 0 << 1 // Read/write by privileged code only
	Arwrw uint8 = 1 << 1 // Read/write by any privilege level
	Ar___ uint8 = 2 << 1 // Read-only by privileged code only
	Ar_r_ uint8 = 3 << 1 // Read-only by any privilege level

	SN uint8 = 0 << 3 // Non-shareable
	SO uint8 = 2 << 3 // Outer shareable
	SI uint8 = 3 << 3 // Inner Shareable
)

// SetBas sets the base address and the permissions of the selected region.
func SetBas(base uintptr, permsh uint8) {
	p().r[0].bar.Store(uint32(base)&^0x1f | uint32(permsh)&0x1f)
}

// Bas returns the base address and the permissions of the selected region.
func Bas() (base uintptr, permsh int8) {
	v := p().r[0].bar.Load()
	base = uintptr(v &^ 0x1f)
	permsh = int8(v & 0x1f)
	return
}

// SetLim sets the limit address and the attribute id of the selected region.
func SetLim(limit uintptr, attrid int, en bool) {
	enbit := uint32(*(*uint8)(unsafe.Pointer(&en)))
	p().r[0].lar.Store(uint32(limit)&^0x1f | uint32(attrid&7)<<1 | enbit)
}

// Lim returns the limit address and the attribute id of the selected region.
func Lim() (limit uintptr, attrid int, en bool) {
	la := p().r[0].lar.Load()
	limit = uintptr(la | 0x1f)
	attrid = int(la >> 1 & 7)
	en = la&1 != 0
	return
}

// Region memory cacheability attributes for SetAttr function.
const (
	// Bit definitions.
	WA int8 = 0b0001 // Normal memory, write allocate
	RA int8 = 0b0010 // Normal memory, read allocate
	WB int8 = 0b0100 // Normal memory, write-back
	NT int8 = 0b1000 // Normal memory, non-transient (cache for long period)

	// Common types of memory.
	Device   int8 = 0                 // Device (SetAttr outer argument only)
	NormalNC int8 = WB                // Normal non-cacheable
	NormalWT int8 = NT | RA           // Normal write-through, no write allocate
	NormalWB int8 = NT | WB | RA | WA // Normal write-back, write allocate
)

// Device attributes (SetAttr inner argument).
const (
	DnGnRnE int8 = 0 // Device-nGnRnE
	DnGnRE  int8 = 1 // Device-nGnRE
	DnGRE   int8 = 2 // Device-nGRE
	DGRE    int8 = 3 // Device-GRE
)

// SetAttr03 sets the values of the 0,1,2,3 cacheability attributes.
func SetAttr03(out0, inn0, out1, inn1, out2, inn2, out3, inn3 int8) {
	p().mair[0].Store(uint32(out0&15)<<4 | uint32(inn0&15) |
		uint32(out1&15)<<12 | uint32(inn1&15)<<8 |
		uint32(out2&15)<<20 | uint32(inn2&15)<<16 |
		uint32(out3&15)<<28 | uint32(inn3&15)<<24)
}

// SetAttr47 sets the values of the 4,5,6,7 cacheability attributes.
func SetAttr47(out4, inn4, out5, inn5, out6, inn6, out7, inn7 int8) {
	p().mair[1].Store(uint32(out4&15)<<4 | uint32(inn4&15) |
		uint32(out5)<<12 | uint32(inn5)<<8 |
		uint32(out6)<<20 | uint32(inn6)<<16 |
		uint32(out7)<<28 | uint32(inn7)<<24)
}
