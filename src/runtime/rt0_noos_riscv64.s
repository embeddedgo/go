// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"
#include "asm_riscv64.h"


// Prefer S0-X15 registers (S0, S1, A0-A5) to allow 16-bit instructions if C
// extension (compressed instruction-set) will be supported. By the convention
// we use An for addresses and addressing related things, Sn for numbers.
//
// The gdb and objdump use C ABI names (Tn, An, Sn, ...) by default. In most
// cases there is imposible to make them print Xn names so we use C ABI names
// in RISC-V assembly except X2 (stack pointer) and g (X4).


#define HSTACK_SIZE 4*1024 // size of stack usesd by trap handlers
#define PALLOC_MIN 64*1024

TEXT _rt0_riscv64_noos(SB),NOSPLIT|NOFRAME,$0

	// initialize all running cores

	CSRWI  (0, MIE)
	MOV   $·trapHandler(SB), S0
	CSRW  (s0, MTVEC)
	
	// Disable interrupts and enable FPU (Kendryte K210 supports only
	// FS=0(off)/3(dirty), this is a weakness of the Rocket Chip Generator used
	// to generate K210 cores).
	MOV   $0x7FFF, S0
	CSRC  (s0, MSTATUS)
	MOV   $(1<<FSn + 1<<MIEn), S0
	CSRS  (s0, MSTATUS)

	CSRWI  (0, MIDELEG)
	CSRWI  (0, MEDELEG)
	CSRWI  (0, MSCRATCH)
	CSRWI  (0, FCSR)

	//MOV  ZERO, X1
	//...
	//MOV  ZERO, X32
	//
	//FCVTDL  ZERO, F0
	//...
	//FCVTDL  ZERO, F31

	// park excess harts
	CSRR  (MHARTID, s0)
	MOV   $const_maxHarts, S1
	BGE   S0, S1, parkHart

	// allocate handler stack for this hart

	// ensure stacks are 16 byte aligned (may be required in the future)
	MOV  $runtime·end(SB), A0
	ADD  $(HSTACK_SIZE+15), A0
	AND  $~15, A0

	// set handler SP for this hart
	MOV  $HSTACK_SIZE, A1
	MUL  S0, A1, A2
	ADD  A2, A0, X2

	// only hart0 does further initialisation
	BNE  ZERO, S0, parkHart

	// clear the BSS and the whole unallocated memory

	MOV   $runtime·bss(SB), A0
	MOV   $runtime·ramend(SB), A1
	SUB   A0, A1
	ADD   $-24, X2
	MOV   ZERO, 0(X2)
	MOV   A0, 8(X2)
	MOV   A1, 16(X2)
	CALL  runtime·memclrNoHeapPointers(SB)

	MOV   $runtime·nodmastart(SB), A0
	MOV   $runtime·nodmaend(SB), A1
	SUB   A0, A1
	MOV   A0, 8(X2)
	MOV   A1, 16(X2)
	CALL  runtime·memclrNoHeapPointers(SB)

	ADD  $24, X2

	JMP  runtime·rt0_go(SB)  // rt0_go is known as top of a goroutine stack

parkHart:
	WFI
	JMP  parkHart


TEXT runtime·rt0_go(SB),NOSPLIT|NOFRAME,$0

	// setup handler stack in harts[mhartid].gh
	MOV   $runtime·harts(SB), g
	MOV   $cpuctx__size, A1
	CSRR  (MHARTID, s0)
	MUL   S0, A1
	ADD   A1, g                  // gh is the first field of the cpuctx struct
	ADD   $-HSTACK_SIZE, X2, A1  // stack_lo
	MOVW  A1, (g_stack+stack_lo)(g)
	MOVW  X2, (g_stack+stack_hi)(g)

	// **** TODO: other harts should go the scheduler here (rearange this) ****

	// set up m0 (bootstrap thread), temporarily use harts[0].gh as g
	MOV  $runtime·m0(SB), A1
	MOV  g, m_g0(A1)  // m0.g0 = harts[mhartid].gh
	MOV  A1, g_m(g)   // harts[mhartid].gh.m = m0

	CALL  runtime·check(SB)
	CALL  runtime·osinit(SB)

	// initialize sysMem

	// calculate the beginning of free memory (just after handler stacks)
	MOV  $runtime·end(SB), A0
	ADD  $(const_maxHarts*HSTACK_SIZE+15), A0
	AND  $~15, A0
	MOV  $runtime·ramend(SB), A1
	SUB  A0, A1, A5  // size of available memory (DMA capable)

	// estimate the space need for non-heap allocations
	SRL  $(const__PageShift+2), A5, A4
	MOV  $mspan__size, A2
	MUL  A2, A4
	ADD  $PALLOC_MIN, A4

	MOV  $runtime·nodmastart(SB), A2
	MOV  $runtime·nodmaend(SB), A3
	SUB  A2, A3, S0  // size of non-DMA memory

	// we prefer the non-DMA memory for non-heap objects to preserve as much as
	// possible of the DMA capable memory for heap allocations
	SUB  S0, A4

	// reduce the arena by the remain of the non-heap space that did not fit in
	// the non-DMA memory, properly align the arena
	BLT  A4, ZERO, 2(PC)
	SUB  A4, A5
	AND  $~(const_heapArenaBytes-1), A5
	SUB  A5, A1

	// save {free.start,free.end,nodma.start,nodma.end,arenaStart,arenaSize}
	MOV  $runtime·sysMem(SB), S0
	MOV  A0, 0(S0)
	MOV  A1, 8(S0)
	MOV  A2, 16(S0)
	MOV  A3, 24(S0)
	MOV  A1, 32(S0)
	MOV  A5, 40(S0)

	// initialize noos tasker and Go scheduler

	CALL  runtime·taskerinit(SB)
	CALL  runtime·schedinit(SB)

	// allocate g0 for m0 and leave gh

	MOV   $(2*const__FixedStack), A0
	ADD   $-24, X2
	MOV   ZERO, 0(X2)
	MOV   A0, 8(X2)
	CALL  runtime·malg(SB)
	MOV   16(X2), A0  // newg in A0
	ADD   $24, X2

	// stackguard check during newproc requires valid stackguard1 but malg
	// sets it to 0xFFFFFFFFFFFFFFFF (mstart fixes this but is called later)
	MOV  g_stackguard0(A0), A1
	MOV  A1, g_stackguard1(A0)

	MOV  $runtime·m0(SB), A1
	MOV  A0, m_g0(A1)  // m0.g0 = newg
	MOV  A1, g_m(A0)   // newg.m = m0

	// newg stack pointer to X2
	MOV  (g_stack+stack_hi)(A0), X2

	MOV  g, A1
	MOV  A0, g

	// fix harts[0].gh, harts[0].mh

	ADD   $cpuctx_mh, A1, A0
	MOV   A1, m_g0(A0)  // harts[0].mh.g0 = harts[0].gh
	MOV   A0, g_m(A1)   // harts[0].gh.m = harts[0].mh
	CSRW  (a1, MSCRATCH)

	// enable interrupts and switch to user mode
	MOV    $0x888, A0
	CSRW   (a0, MIE)
	AUIPC  $0, A0
	ADD    $16, A0  // A0 must point just after MRET
	CSRW   (a0, MEPC)
	MRET

	// create a new goroutine to start program
	MOV   $runtime·mainPC(SB), A0  // entry
	ADD   $-24, X2
	MOV   A0, 16(X2)
	MOV   ZERO, 8(X2)
	MOV   ZERO, 0(X2)
	CALL  runtime·newproc(SB)
	ADD   $24, X2

	// start this M
	CALL  runtime·mstart(SB)

	UNDEF  // fail
