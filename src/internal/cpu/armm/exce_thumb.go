// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package armm

// Cortex-M exception numbers.
const (
	Reset       = 1  // IRQ number:
	NMI         = 2  // -14
	HardFault   = 3  // -13
	MemManage   = 4  // -12
	BusFault    = 5  // -11
	UsageFault  = 6  // -10
	SecureFault = 7  //  -9
	_           = 8  //  -8
	_           = 9  //  -7
	_           = 10 //  -6
	SVCall      = 11 //  -5
	DebugMon    = 12 //  -4
	_           = 13 //  -3
	PendSV      = 14 //  -2
	SysTick     = 15 //  -1
)

// Exception number for the first external interrupt.
const IRQ0 = 16

// EXC_RETURN fields for ARMv8-M with the Security Extension
const (
	ExcReturnPrefix  = 0xff << 24   // Indicates that this is EXC_RETURN value
	ExcReturnRes23_7 = 0x1ffff << 7 // Reserved

	// The least sginificant bits listed below indicate the required return
	// stack, processor mode, security state, and stack frame. The descriptions
	// refer to the state when the bit is set.

	ExcReturnS     = 1 << 6 // Registers stacked to secure stack.
	ExcReturnDCRS  = 1 << 5 // Default rules for stacking calle registers.
	ExcReturnFType = 1 << 4 // No FPU context on the stack.
	ExcReturnMode  = 1 << 3 // Return to thread mode.
	ExcReturnSPSEL = 1 << 2 // Registers stacked to PSP, use PSP after return.
	ExcReturnES    = 1 << 0 // Return to secure mode / ARMv7-M thumb mode bit.

	ExcReturnBase            = ExcReturnPrefix | ExcReturnRes23_7
	ExcReturnSecureThreadPSP = ExcReturnBase |
		ExcReturnS |
		ExcReturnDCRS |
		ExcReturnFType |
		ExcReturnMode |
		ExcReturnSPSEL |
		ExcReturnES
)
