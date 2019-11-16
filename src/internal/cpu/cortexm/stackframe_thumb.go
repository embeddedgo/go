// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cortexm

// Masks to select PSR subregisters/bits
const (
	APSR  = 0xF80F0000 // Application Program Status Register:
	GE    = 0x000F0000 // - greater than or equal flags (SIMD instructions)
	Q     = 0x08000000 // - saturation flag (DSP instructions)
	V     = 0x10000000 // - overflow condition flag
	C     = 0x20000000 // - carry condition flag
	Z     = 0x40000000 // - zero condition flag
	N     = 0x80000000 // - negative condition flag
	IPSR  = 0x000001FF // Interrupt Program Status Register (exception number)
	EPSR  = 0x0700FC00 // Execution Program Status Register:
	SR    = 0x00000200 // - stack realigned (in stacked PSR only)
	ITICI = 0x0600FC00 // - interrupt-continue for LDM/STM or IT block
	T     = 0x01000000 // - execution mode (0: ARM, 1: Thumb)
)

// StackFrame represents the ARMv7-M stack frame
type StackFrame struct {
	R   [4]uintptr
	R12 uintptr
	LR  uintptr
	PC  uintptr
	PSR uint32
}

// StackFrameF represents the ARMv7-M extended stack frame with
// single-precision floating-point registers
type StackFrameF struct {
	StackFrame
	S     [16]float32
	FPSCR uint32
}

// StackFrameD represents the ARMv7-M extended stack frame with
// double-precision floating-point registers
type StackFrameD struct {
	StackFrame
	D     [8]float64
	FPSCR uint32
}
