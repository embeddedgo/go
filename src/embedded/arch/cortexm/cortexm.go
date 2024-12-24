// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package cortexm

import "unsafe"

// CPUID partno constants
const (
	StarMC1    int16 = 0x132
	CortexM0   int16 = 0xc20
	CortexM1   int16 = 0xc21
	CortexM3   int16 = 0xc23
	CortexM4   int16 = 0xc24
	CortexM7   int16 = 0xc27
	CortexM0p  int16 = 0xc60
	CortexM23  int16 = 0xd20
	CortexM33  int16 = 0xd21
	CortexM35P int16 = 0xd31
	CortexM55  int16 = 0xd22
	CortexM85  int16 = 0xd23
	SLx2       int16 = 0xdB0
	RealM200   int16 = 0xd20
	RealM300   int16 = 0xd22
)

const (
	ARM      uint8 = 0x41
	Infineon uint8 = 0x49
	ARMChina uint8 = 0x63
	Realtek  uint8 = 0x72
)

func CPUID() (partno int16, variant, revision int8, implementer uint8) {
	cpuid := *(*uint32)(unsafe.Pointer(uintptr(0x0e000ed00)))
	revision = int8(cpuid & 0xf)
	partno = int16(cpuid >> 4 & 0x3f)
	//arch = cpuid >> 16 & 0xf - always 0xf
	variant = int8(cpuid >> 20 & 0xf)
	implementer = uint8(cpuid >> 24)
	return
}

func IsARMv8M(partno int16) bool {
	return partno&0xf00 == 0xd00
}
