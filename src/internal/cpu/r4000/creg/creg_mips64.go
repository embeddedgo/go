// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package creg provides access to the control registers.
// Page references refer to Vr4300 User' Manual (7th edition)
package creg

type CP0Reg uint

// System Control Processor (CP0) registers (p.46)
const (
	// memory management registers
	INDEX        CP0Reg = 0
	RANDOM       CP0Reg = 1
	WIRED        CP0Reg = 6
	PRID         CP0Reg = 15
	CONFIG       CP0Reg = 16
	LLADDR       CP0Reg = 17
	ENTRY_LO_0   CP0Reg = 2
	ENTRY_LO_1   CP0Reg = 3
	ENTRY_HI     CP0Reg = 10
	PAGE_MASK    CP0Reg = 5
	CACHE_TAG_LO CP0Reg = 28
	CACHE_TAG_HI CP0Reg = 29

	// exception processing registers
	CONTEXT    CP0Reg = 4
	BAD_VADDR  CP0Reg = 8
	COUNT      CP0Reg = 9
	COMPARE    CP0Reg = 11
	STATUS     CP0Reg = 12
	CAUSE      CP0Reg = 13
	EPC        CP0Reg = 14
	WATCH_LO   CP0Reg = 18
	WATCH_HI   CP0Reg = 19
	XCONTEXT   CP0Reg = 20
	PARITY_ERR CP0Reg = 26
	CACHE_ERR  CP0Reg = 27
	ERROR_EPC  CP0Reg = 30
)

// Status register bits (p.166)
const (
	IE  = 0x001 << 0  // interrupt enable
	EXL = 0x001 << 1  // exception level
	ERL = 0x001 << 2  // error level
	KSU = 0x003 << 3  // kernel/supervisor/user mode
	UX  = 0x001 << 5  // user mode 64-bit addressing and operations
	SX  = 0x001 << 6  // supervisor mode 64-bit addressing and operations
	KX  = 0x001 << 7  // kernel mode 64-bit addressing
	IM  = 0x0ff << 8  // interrupt mask
	DS  = 0x1ff << 16 // diagnostic status
	RE  = 0x001 << 25 // reverse-endian
	FR  = 0x001 << 26 // floating-point registers
	RP  = 0x001 << 27 // reduced power
	CU  = 0x00f << 28 // coprocessor usability
)

// Cause register bits (p.171)
const (
	IP_SW0     = 0x01 << 8
	IP_SW1     = 0x02 << 8
	IP_SWMASK  = 0x03 << 8
	IP_EXTMASK = 0x7C << 8
	IP_TIMER   = 0x80 << 8
)

func (r CP0Reg) Load() uint32     { return mfc0(r) }
func (r CP0Reg) Store(val uint32) { mtc0(r, val) }

func (r CP0Reg) LoadBits(mask uint32) uint32 {
	return r.Load() & mask
}
func (r CP0Reg) StoreBits(mask uint32, val uint32) {
	r.Store((r.Load() & ^mask) | (val & mask))
}
func (r CP0Reg) SetBits(mask uint32) {
	r.Store((r.Load() | mask))
}
func (r CP0Reg) ClearBits(mask uint32) {
	r.Store((r.Load() & ^mask))
}

func mtc0(creg CP0Reg, val uint32)
func mfc0(creg CP0Reg) uint32
