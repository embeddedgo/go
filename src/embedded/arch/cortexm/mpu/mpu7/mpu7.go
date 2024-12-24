// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mpu7 provides an interface to the ARMv7-M Memory Protection Unit.
package mpu7

import (
	"embedded/mmio"
	"unsafe"
)

type regs struct {
	_    [3]uint32
	rbar mmio.R32[uint32]
	rasr mmio.R32[uint32]
}

func p() *regs { return (*regs)(unsafe.Pointer(uintptr(0xE000ED90))) }

type Attr uint32

const (
	ENA Attr = 1 << 0 // Enables region

	B Attr = 1 << 16 // Bufferable
	C Attr = 1 << 17 // Cacheable
	S Attr = 1 << 18 // Shareable

	TEX0 Attr = 0 << 19
	TEX1 Attr = 1 << 19
	TEX2 Attr = 2 << 19

	// Access permissons.
	Amask Attr = 7 << 24 // Use to extract access permission bits
	A____ Attr = 0 << 24 // No access
	Ar___ Attr = 5 << 24 // Priv-RO
	Arw__ Attr = 1 << 24 // Priv-RW
	Ar_r_ Attr = 6 << 24 // Priv-RO, Unpriv-RO
	Arwr_ Attr = 2 << 24 // Priv-RW, Unpriv-RO
	Arwrw Attr = 3 << 24 // Priv-RW, Unpriv-RW

	XN Attr = 1 << 28 // Instruction access disable
)

func SIZE(exp int) Attr {
	return Attr(exp-1) & 0x1f << 1
}

func (a Attr) SIZE() (exp int) {
	return int(a>>1)&0x1f + 1
}

func SRD(srd uint8) Attr {
	return Attr(srd) << 8
}

func (a Attr) SRD() uint8 {
	return uint8(a >> 8)
}

func TEX(tex int) Attr {
	return Attr(tex&7) << 19
}

func (a Attr) TEX() int {
	return int(a>>19) & 7
}

const VALID = 1 << 4

func SetRegion(base uintptr, attr Attr) {
	p().rbar.Store(uint32(base))
	p().rasr.Store(uint32(attr))
}

func Region() (base uintptr, attr Attr) {
	return uintptr(p().rbar.Load()), Attr(p().rasr.Load())
}

/*
TODO: Implement SetRegions using STM instrucion.
type BaseAttr struct {
	RBAR uintptr
	RASR Attr
}

func SetRegions(bas [4]BaseAttr)
*/
