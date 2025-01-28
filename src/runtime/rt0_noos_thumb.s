// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"

// The runtime package dosn't implement required _rt0_thumb_noos and
// runtime·identcurcpu functions because they are generally target specific.
// See the hal/system package in the https://github.com/embeddedgo/stm32 or
// https://github.com/embeddedgo/pico repositories for example implementation.

// initRAMfromROM copies the Data segment from ROM to RAM and clears the
// remaining RAM. As it clears the whole free RAM and doesn't know about CPU
// stacks it may be called only when stacks are empty.
TEXT runtime·initRAMfromROM(SB),NOSPLIT|NOFRAME,$0
	MOVW    $0, R0        // dummy RA
	MOVW.W  R0, -16(R13)  // make a "frame"

	MOVW    $runtime·bss(SB), R0
	MOVW    $runtime·ramend(SB), R1
	MOVW.W  LR, -4(R1)  // save LR at the end of RAM
	SUB     R0, R1      // R1 = freeSize = ramEnd-4 - bssStart
	MOVW    R0, 4(R13)
	MOVW    R1, 8(R13)
	BL      runtime·memclrNoHeapPointers(SB)  // clear free RAM

	MOVW  $runtime·nodmastart(SB), R0
	MOVW  $runtime·nodmaend(SB), R1
	SUB   R0, R1  // R1 = nondmaSize = nodmaEnd - nodmaStart
	MOVW  R0, 4(R13)
	MOVW  R1, 8(R13)
	BL    runtime·memclrNoHeapPointers(SB)  // clear non-DMA RAM

	MOVW  $runtime·noptrdata(SB), R0
	MOVW  $runtime·romdata(SB), R1
	MOVW  $runtime·edata(SB), R2
	SUB   R0, R2  // R2 = dataSize = dataStart - dataEnd
	MOVW  R0, 4(R13)
	MOVW  R1, 8(R13)
	MOVW  R2, 12(R13)
	BL    runtime·memmove(SB)  // copy data to RAM

	// Restore SP, LR and return
	ADD     $16, R13
	MOVW    $runtime·ramend(SB), R0
	MOVW.W  -4(R0), LR
	MOVW    $0, R1
	MOVW    R1, (R0)  // clear the last word in RAM
	RET


// rt0_go initializes the noos tasker, Go scheduler and continues as the first
// thread that runs the first goroutine. If the system has multiple CPUs
// only one CPU can run this function (init CPU, usually CPU0). Other CPUs must
// wait until the tasker is ready.
TEXT runtime·rt0_go(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	// _rt0_thumb_noos may provide stackStart and stackEnd in R0, R1
	CMP        R0, R1
	MOVW.EQ    $runtime·ramstart(SB), R0  // default stackStart
	MOVW.EQ    R13, R1                    // default stackEnd
	MOVM.DB.W  [R0, R1], (R13)            // save stackStart,stackEnd

	// Initialize the memory allocator
	MOVW       $0, R0                       // dummy RA
	MOVW       $runtime·end(SB), R1         // freeStart
	MOVW       $runtime·ramend(SB), R2      // freeEnd
	MOVW       $runtime·nodmastart(SB), R3  // nodmaStart
	MOVW       $runtime·nodmaend(SB), R4    // nodmaEnd
	ADD        $8, R13, R5                  // stackTop
	MOVW       $0, R6                       // return value
	MOVM.DB.W  [R0-R6], (R13)
	BL         runtime·meminit(SB)
	ADD        $24, R13  // SP points to the nodmaStack (return value)

	// Initialize tasker
	MOVW  (R13), R0  // load nodmaStack
	CBZ   R0, argsReady
	MOVW  $0, R0
	MOVW  R0, (R13)                    // dummy RA
	MOVW  $runtime·nodmastart(SB), R0  // bottom of the stack space
	MOVW  R0, 4(R13)                   // update stackStart
argsReady:
	BL   runtime·taskerinit(SB)
	ADD  $12, R13

	// Enable exceptions
	CPSIE

	// set up m0 (bootstrap thread), temporarily use gh as g
	BL    ·identcurcpu(SB)  // R0 = cpuctx for current cpu
	MOVW  $runtime·m0(SB), R1
	MOVW  R0, m_g0(R1)        // m0.g0 = curcpu.gh
	MOVW  R1, g_m(R0)         // curcpu.gh.m = &m0
	MOVW  R0, g               // set g to gh
	MOVW  R1, cpuctx_exe(R0)  // curcpu.exe = &m0

	//BL  runtime·emptyfunc(SB)  // fault if stack check is wrong
	BL  runtime·check(SB)
	BL  runtime·schedinit(SB)

	// allocate g0 for m0 and leave gh

	SUB        $4, R13                  // space for return value
	MOVW       $0, R0                   // dummy RA
	MOVW       $(2*const_stackMin), R1  // arg
	MOVM.DB.W  [R0-R1], (R13)
	BL         runtime·malg(SB)
	MOVW       8(R13), R0  // newg in R0
	ADD        $12, R13

	// stackguard check during newproc requires valid stackguard1 but
	// malg sets it to 0xFFFFFFFF (mstart fixes this but is called later)
	MOVW  g_stackguard0(R0), R1
	MOVW  R1, g_stackguard1(R0)

	MOVW  $runtime·m0(SB), R1
	MOVW  R0, m_g0(R1)  // m0.g0 = newg
	MOVW  R1, g_m(R0)   // newg.m = m0

	MOVW  (g_stack+stack_hi)(R0), R1
	MOVW  R1, PSP

	MOVW  g, R2  // R2 = &curcpu.gh (R2 = curcpu)
	MOVW  R0, g  // g = newg

	// fix curcpu.gh
	ADD   $cpuctx_mh, R2, R1  // R1 = &curcpu.mh
	MOVW  R1, g_m(R2)         // curcpu.gh.m = &curcpu.gh

	// leave the main stack and the privileged mode
	DSB
	MOVW  CONTROL, R0
	ORR   $2, R0  // use PSP as stack pointer
	MOVW  R0, CONTROL
	ISB
	ORR   $1, R0  // go to unprivileged mode
	MOVW  R0, CONTROL
	ISB

	// create a new goroutine to start program
	SUB   $8, R13
	MOVW  $runtime·mainPC(SB), R0
	MOVW  R0, 4(R13)  // arg 1: fn
	MOVW  $0, R0
	MOVW  R0, 0(R13)  // dummy RA sloot
	BL    runtime·newproc(SB)
	ADD   $8, R13  // pop args and LR

	// start this M
	BL  runtime·mstart(SB)

	UNDEF  // fail
