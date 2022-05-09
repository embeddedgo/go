// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"


TEXT _rt0_thumb_noos(SB),NOSPLIT|NOFRAME,$0
	// initialize data and BSS

	MOVW       R13, R0  // R13 points to the top of ISR stack and the beggining of DATA segment
	MOVW       $runtime·romdata(SB), R1
	MOVW       $runtime·bss(SB), R3
	MOVW       $runtime·ramend(SB), R4
	SUB        R0, R3, R2
	SUB        R3, R4
	MOVM.DB.W  [R0-R4], (R13)  // push: to,from,n for memmove, ptr,n for memclr
	SUB        $4, R13
	BL         runtime·memmove(SB)  // copy data to RAM
	ADD        $12, R13
	BL         runtime·memclrNoHeapPointers(SB)  // clear BSS and unallocated memory
	MOVW       $runtime·nodmastart(SB), R0
	MOVW       $runtime·nodmaend(SB), R1
	SUB        R0, R1
	MOVW       R0, 4(R13)
	MOVW       R1, 8(R13)
	BL         runtime·memclrNoHeapPointers(SB)  // clear non-DMA memory
	ADD        $12, R13

	B   runtime·rt0_go(SB)  // rt0_go is known as top of a goroutine stack


#define PALLOC_MIN 20*1024
#define FPU_CTRL_BASE 0xE000ED88
#define FPU_CPACR 0x000
#define FPU_FPCCR 0x1AC

TEXT runtime·rt0_go(SB),NOSPLIT|NOFRAME|TOPFRAME,$0

	// setup main stack in cpu0.gh
	MOVW  $runtime·cpu0(SB), R0      // gh is the first field of the cpuctx struct
	MOVW  $runtime·ramstart(SB), R1  // main stack starts at the beggining of memory
	MOVW  R1, (g_stack+stack_lo)(R0)
	MOVW  R13, (g_stack+stack_hi)(R0)
	ADD   $const__StackGuard, R1
	MOVW  R1, g_stackguard0(R0)
	MOVW  R1, g_stackguard1(R0)

	// set up m0 (bootstrap thread), temporarily use cpu0.gh as g
	MOVW  $runtime·m0(SB), R1
	MOVW  R0, m_g0(R1)  // m0.g0 = cpu0.gh
	MOVW  R1, g_m(R0)   // cpu0.gh.m = m0

	MOVW  R0, g  // set g to gh

	// enable FPU if GOARM is xF or xD
	MOVB  runtime·goarm(SB), R0
	AND   $0xD, R0
	CMP   $0xD, R0
	BNE   skipFPU
	MOVW  $FPU_CTRL_BASE, R0  // address of CPACR
	MOVW  $0xF<<20, R1        // full access to CP10 and CP11 instructions
	MOVW  R1, FPU_CPACR(R0)
	SLL   $10, R1
	MOVW  R1, FPU_FPCCR(R0)  // set LSPEN and ASPEN (lazy auto save FP context)
skipFPU:

	//BL  runtime·emptyfunc(SB)  // fault if stack check is wrong
	BL  runtime·check(SB)
	BL  runtime·osinit(SB)

	// initialize sysMem

	MOVW  $runtime·end(SB), R0
	MOVW  $runtime·ramend(SB), R1
	SUB   R0, R1, R5  // size of available memory (DMA capable)

	// estimate the space need for non-heap allocations
	MOVW  R5>>(const__PageShift+2), R4
	MOVW  $mspan__size, R2
	MUL   R2, R4
	ADD   $PALLOC_MIN, R4

	MOVW  $runtime·nodmastart(SB), R2
	MOVW  $runtime·nodmaend(SB), R3
	SUB   R2, R3, R6  // size of non-DMA memory

	// we prefer the non-DMA memory for non-heap objects to preserve as much as
	// possible of the DMA capable memory for heap allocations
	SUB.S  R6, R4

	// reduce the arena by the remain of the non-heap space that did not fit in
	// the non-DMA memory, properly align the arena
	SUB.HI  R4, R5
	BIC     $(const_heapArenaBytes-1), R5
	SUB     R5, R1
	MOVW    R1, R4

	// save {free.start,free.end,nodma.start,nodma.end,arenaStart,arenaSize}
	MOVW     $runtime·sysMem(SB), R6
	MOVM.IA  [R0-R5], (R6)

	// initialize noos tasker and Go scheduler

	BL  runtime·taskerinit(SB)
	BL  runtime·schedinit(SB)

	// allocate g0 for m0 and leave gh

	SUB        $4, R13
	MOVW       $0, R0
	MOVW       $(2*const__StackMin), R1
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

	MOVW  g, R2
	MOVW  R0, g

	// fix cpu0.gh, cpu0.mh

	ADD   $cpuctx_mh, R2, R1  // R2 points to cpu0 (and to cpu0.gh at the same time)
	MOVW  R2, m_g0(R1)        // cpu0.mh.g0 = cpu0.gh
	MOVW  R2, m_gsignal(R1)   // cpu0.mh.gsignal = cpu0.gh (to easily check for handler mode)
	MOVW  R1, g_m(R2)         // cpu0.gh.m = cpu0.mh

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
	SUB	$8, R13
	MOVW	$runtime·mainPC(SB), R0
	MOVW	R0, 4(R13)	// arg 1: fn
	MOVW	$0, R0
	MOVW	R0, 0(R13)	// dummy LR
	BL	runtime·newproc(SB)
	ADD	$8, R13	// pop args and LR

	// start this M
	BL  runtime·mstart(SB)

	UNDEF  // fail
