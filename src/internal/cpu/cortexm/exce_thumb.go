// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cortexm

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

// EXC_RETURN fields
const (
	ExcReturnBase    = 0xFFFFFFE0 // EXC_RETURN base value
	ExcReturnMode    = 0xF        // Selects bits responsible for return mode:
	ExcReturnHandler = 0x01       // - return to handler mode, use MSP
	ExcReturnMSP     = 0x09       // - return to thread mode, use MSP
	ExcReturnPSP     = 0x0D       // - return to thread mode, use PSP
	ExcReturnNoFPU   = 0x10       // Basic frame on the stack (no FPU state)
)
