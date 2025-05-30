// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

// _rt0_thumb_noos is the first function of the Embedded Go program
TEXT _rt0_thumb_noos(SB),NOSPLIT|NOFRAME,$0
	// Uncomment in case of truble to connect the debugger to the running program.
	//B  0(PC) // gdb `set $pc += 4` to exit this loop

	// Disable exceptions until the CPU is ready for SVCall and PendSV.
	CPSID

	// Cortex-M settings.
	MOVW       $0, R0                    // dummy RA
	MOVW       $runtime·vectors(SB), R1  // arg
	MOVM.DB.W  [R0-R1], (R13)
	BL         runtime·initCPU(SB)
	ADD        $8, R13

	// Clear memory and load the data segment from Flash.
	BL  runtime·initRAMfromROM(SB)

	// Set the numer of available CPUs.
	MOVW  $1, R0
	MOVW  $runtime·ncpu(SB), R1
	MOVW  R0, (R1)

	MOVW  R0, R1  // inform rt0_go that we use default stack arrangement
	B     runtime·rt0_go(SB)


// identcurcpu indetifies the current CPU and returns the pointer to its cpuctx
// in R0. It can clobber R0-R4,LR registers (other registers must be preserved).
TEXT runtime·identcurcpu(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  $runtime·thetasker(SB), R0
	MOVW  (R0), R0  // allcpu is the first field of the runtime.tasker struct
	MOVW  (R0), R0  // R0 = thetasker.allcpu[0] (single CPU system)
	RET


// func preepmtOrWakeup(cpuid int)
TEXT runtime·preemptOrWakeup(SB),NOSPLIT|NOFRAME,$0-4
	// Only single CPU is supported so only cpuid=-1 is expected (wakeup).
	SEV  // see ARM Errata 563915, STM32F10xx Errata 1.1.2
	RET
