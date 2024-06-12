// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"
#include "asm_mips64.h"

// BREAK is overloaded CACHE opcode. Register number specifies the cache op.
#define CACHE BREAK

TEXT _rt0_mips64_noos(SB),NOSPLIT|NOFRAME,$0
	JMP ·rt0_target(SB)

TEXT runtime·_rt0_mips64_noos1(SB),NOSPLIT|NOFRAME,$0
	// Clear .bss, .noptrbss and unallocated memory.
	SUBU $16, R29

	MOVW $runtime·bss(SB), R4
	MOVW $runtime·ebss(SB), R5
	SUB  R4, R5
	MOVV R4, 8(R29)
	MOVV R5, 16(R29)
	JAL  runtime·memclrNoHeapPointers(SB)

	MOVW $runtime·noptrbss(SB), R4
	MOVW $runtime·enoptrbss(SB), R5
	SUB  R4, R5
	MOVV R4, 8(R29)
	MOVV R5, 16(R29)
	JAL  runtime·memclrNoHeapPointers(SB)

	MOVW $runtime·end(SB), R4
	MOVW $runtime·ramend(SB), R5
	SUB  R4, R5
	MOVV R4, 8(R29)
	MOVV R5, 16(R29)
	JAL  runtime·memclrNoHeapPointers(SB)

	ADDU $16, R29

	// Load interrupt vector
	MOVW $runtime·intvector(SB), R8
	MOVW $0xa0000000, R9
	MOVW $4, R10
loop:
	MOVW (R8), R11
	MOVW R11, 0(R9)
	MOVW R11, 0x80(R9)
	MOVW R11, 0x100(R9)
	MOVW R11, 0x180(R9)
	CACHE R16, 0(R9)
	CACHE R16, 0x80(R9)
	CACHE R16, 0x100(R9)
	CACHE R16, 0x180(R9)
	ADD $4, R8
	ADD $4, R9
	ADDU $-1, R10
	BGTZ R10,loop

	JMP runtime·rt0_go(SB)


#define PALLOC_MIN 20*1024

TEXT runtime·rt0_go(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	// setup main stack in cpu0.gh
	MOVV  $runtime·cpu0(SB), t0  // gh is the first field of the cpuctx struct
	MOVV  $runtime·ramend(SB), sp  // main stack starts at the end of memory
	SUB   $16, sp
	MOVV  sp, (g_stack+stack_hi)(t0)
	SUB   $4096, sp, t1
	MOVV  t1, (g_stack+stack_lo)(t0)
	ADD   $const_stackGuard, t1
	MOVV  t1, g_stackguard0(t0)
	MOVV  t1, g_stackguard1(t0)

	// set up m0 (bootstrap thread), temporarily use cpu0.gh as g
	MOVV  $runtime·m0(SB), t1
	MOVV  t0, m_g0(t1)  // m0.g0 = cpu0.gh
	MOVV  t1, g_m(t0)   // cpu0.gh.m = m0

	MOVV  t0, g  // set g to gh

	JAL runtime·check(SB)
	JAL runtime·osinit(SB)

	// initialize noosMem

	MOVV  $runtime·end(SB), t0
	MOVV  $runtime·ramend(SB), t1
	SUB   $4096, t1
	SUB   t0, t1, t5  // size of available memory (DMA capable)

	// estimate the space need for non-heap allocations
	SRL   $(const__PageShift), t5, t4
	MOVV  $mspan__size, t2
	MUL   t2, t4
	MOVV  LO, t4
	ADD   $PALLOC_MIN, t4

	MOVV  $runtime·nodmastart(SB), t2
	MOVV  $runtime·nodmaend(SB), t3
	SUB   t2, t3, t7  // size of non-DMA memory
	ADD   t5, t7, t6  // size of the whole free memory

	// we prefer the non-DMA memory for non-heap objects to preserve as much as
	// possible of the DMA capable memory for heap allocations
	SUB   t7, t4

	// reduce the arena by the remain of the non-heap space that did not fit in
	// the non-DMA memory, properly align the arena
	BLTZ  t4, 2(PC)
	SUB   t4, t5
	AND   $~(const_heapArenaBytes-1), t5
	SUB   t5, t1
	MOVV  t1, t4

	// save {free.start,free.end,nodma.start,nodma.end,arenaStart,arenaSize,size}
	MOVV  $runtime·noosMem(SB), t7
	MOVV  t0, 0(t7)
	MOVV  t1, 8(t7)
	MOVV  t2, 16(t7)
	MOVV  t3, 24(t7)
	MOVV  t4, 32(t7)
	MOVV  t5, 40(t7)
	MOVV  t6, 48(t7)

	// initialize noos tasker and Go scheduler
	JAL  runtime·taskerinit(SB)
	JAL  runtime·schedinit(SB)

	// allocate g0 for m0 and leave gh
	SUB   $16, sp
	MOVW  $(2*const_stackMin), a0
	MOVW  a0, 8(sp)
	JAL   runtime·malg(SB)
	MOVV  16(sp), t0  // newg in t0
	ADD   $16, sp

	// stackguard check during newproc requires valid stackguard1 but
	// malg sets it to 0xFFFFFFFF (mstart fixes this but is called later)
	MOVV  g_stackguard0(t0), t1
	MOVV  t1, g_stackguard1(t0)

	MOVV  $runtime·m0(SB), t1
	MOVV  t0, m_g0(t1)  // m0.g0 = newg
	MOVV  t1, g_m(t0)   // newg.m = m0

	// newg stack pointer to sp
	MOVV  (g_stack+stack_hi)(t0), sp

	// newg to g
	MOVV  g, t2
	MOVV  t0, g

	// fix cpu0.gh, cpu0.mh
	ADD   $cpuctx_mh, t2, t1  // t2 points to cpu0 (and to cpu0.gh at the same time)
	MOVV  t2, m_g0(t1)        // cpu0.mh.g0 = cpu0.gh
	MOVV  t2, m_gsignal(t1)   // cpu0.mh.gsignal = cpu0.gh (to easily check for handler mode)
	MOVV  t1, g_m(t2)         // cpu0.gh.m = cpu0.mh

	// TODO switch to the user mode?

	// create a new goroutine to start program
	SUB	$16, sp
	MOVV	$runtime·mainPC(SB), t0
	MOVV	t0, 8(sp)	// arg 1: fn
	MOVV	$0, t0
	MOVV	t0, 0(sp)	// dummy LR
	JAL	runtime·newproc(SB)
	ADD	$16, sp	// pop args and LR

	// enable interrupts
	// TODO where to enable interupts correctly?
	MOVW z0, M(C0_COMPARE)
	MOVW M(C0_SR), t0
	OR  $(SR_IE|INTR_SW|INTR_EXT), t0
	MOVW t0, M(C0_SR)

	// start this M
	JAL  runtime·mstart(SB)

	UNDEF
