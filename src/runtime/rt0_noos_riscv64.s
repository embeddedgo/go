// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"
#include "asm_riscv64.h"


// Prefer X8-X15 registers (S0, S1, A0-A5) to allow compressed instructions if
// compiler will support C extension. By the convention An registers are
// preffered for addresses and addressing related things, Sn for numbers and
// other things..
//
// The gdb and objdump use C ABI names for registers (Tn, An, Sn, ...). In most
// cases there is imposible to make them print Xn names so we use C ABI names
// in RISC-V assembly except LR (X1), X2 (stack pointer), g (X27) and TMP (X31).


#define handlerStackSize 4*1024 // size of stack usesd by trap handlers
#define persistAllocMin 128*1024

#define intEnabled (MTI+MSI+SEI+MEI)

DATA runtime·waitInit+0(SB)/4, $-1
GLOBL runtime·waitInit(SB), NOPTR, $4

// _rt0_riscv64_noos initializs all running cores
TEXT _rt0_riscv64_noos(SB),NOSPLIT|NOFRAME,$0
	// Disable interrupts globally, enable FPU (Kendryte K210 supports only
	// FS=0(off)/3(dirty), this is a weakness of the Rocket Chip Generator used
	// to generate K210 cores).
	MOV   $0x7FFF, S0
	CSRC  (s0, mstatus)
	MOV   $(1<<FSn), S0  // set FS to init
	CSRS  (s0, mstatus)

	// enable interrupts locally
	MOV   $intEnabled, S0
	CSRW  (s0, mie)

	// Set a temporary trap handler.
	MOV   $·defaultExceptionHandler(SB), S0
	CSRW  (s0, mtvec)

	// Clear some other CSRs
	CSRWI  (0, mideleg)
	CSRWI  (0, medeleg)
	CSRWI  (0, mscratch)
	CSRWI  (0, fcsr)

	//MOV  ZERO, X1
	//...
	//MOV  ZERO, X32
	//
	//FCVTDL  ZERO, F0
	//...
	//FCVTDL  ZERO, F31

	// park excess harts
	CSRR  (mhartid, s0)  // from now S0 used only to provide hart id
	MOV   $const_maxHarts, S1
	BGEU  S0, S1, parkHart

	// ensure handler stacks are 16 byte aligned (may be required in the future)
	MOV  $runtime·end(SB), A0
	ADD  $(handlerStackSize+15), A0
	AND  $~15, A0

	// set handler SP for this hart
	MOV  $handlerStackSize, A1
	MUL  S0, A1, A2
	ADD  A2, A0, X2

	// clear msip register
	MOV   $msip, A0
	SLL   $2, S0, A1
	ADD   A1, A0
	MOVW  ZERO, (A0)

	// set mtimecmp to maximum value
	MOV  $mtimecmp, A0
	SLL  $3, S0, A1
	ADD  A1, A0
	MOV  $-1, A1
	MOV  A1, (A0)

	BEQ  ZERO, S0, clear

	// Other harts have to wait for hart0 to initaialize all shared components
	MOV   $·waitInit(SB), A0
	MOVW  (A0), S1
	BNE   ZERO, S1, -1(PC)
	JMP   cleared

clear:
	// clear the mtime register
	MOV  $mtime, A0
	MOV  ZERO, (A0)

	// clear the BSS and the whole unallocated memory
	ADD   $-24, X2
	MOV   $runtime·bss(SB), A0
	MOV   $runtime·ramend(SB), A1
	SUB   A0, A1
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
	ADD   $24, X2

cleared:
	// setup handler stack in harts[mhartid].gh
	MOV  $runtime·harts(SB), g
	MOV  $cpuctx__size, A0
	MUL  S0, A0
	ADD  A0, g  // gh is the first field of the cpuctx struct
	ADD  $-handlerStackSize, X2, A0
	MOV  A0, (g_stack+stack_lo)(g)
	MOV  X2, (g_stack+stack_hi)(g)

	// setup gh and mh
	ADD  $cpuctx_mh, g, A0
	MOV  g, m_g0(A0)       // harts[mhartid].mh.g0 = harts[mhartid].gh
	MOV  g, m_gsignal(A0)  // harts[mhartid].mh.gsignal = harts[mhartid].gh (to easily check for handler mode)
	MOV  A0, g_m(g)        // harts[mhartid].gh.m = harts[mhartid].mh

	// as we have SP and g set we can set the real trap handler in mtvec
	MOV   $·trapHandler(SB), A0
	CSRW  (a0, mtvec)

	// hart0 runs the m0
	BNE  ZERO, S0, 2(PC)
	JMP  runtime·rt0_go(SB)

	MOV  S0, g_goid(g)  // set cpuctx.gh.goid to mhartid

	// set thetasker.allcpu.len to hartid+1 if lower
	MOV  cpuctx_t(g), A0         // &thetasker
	ADD  $(tasker_allcpu+8), A0  // &thetasker.allcpu.len
	ADD  $1, S0
again:
	LRD   (a0, s1)
	BGEU  S1, S0, runScheduler
	SCD   (s0, s1, a0)
	BNE   ZERO, S1, again

runScheduler:
	// prepare trap context (only mstatus and mie are required on the stack)
	ADD   $-trapCtxSize, X2
	CSRR  (mstatus, s0)
	OR    $(1<<MPIEn), S0
	MOV   S0, _mstatus(X2)
	MOV   $intEnabled, S0
	MOV   S0, _mie(X2)
	MOV   $0xFF, S0
	CSRC  (s0, mie)  // disable MTI+MSI before enter the scheduler

	// enter scheduler
	MOV   $1, S0
	MOVB  S0, (cpuctx_schedule)(g)
	JMP   ·enterScheduler(SB)

parkHart:
	WFI
	JMP  -1(PC)


// rt0_go is known as top level function
TEXT runtime·rt0_go(SB),NOSPLIT|NOFRAME,$0

	// set up m0 (bootstrap thread), temporarily use harts[0].gh as g
	MOV  $runtime·m0(SB), A0
	MOV  g, m_g0(A0)  // m0.g0 = harts[0].gh
	MOV  A0, g_m(g)   // harts[0].gh.m = m0

	CALL  runtime·check(SB)
	CALL  runtime·osinit(SB)

	// initialize sysMem

	// calculate the beginning of free memory (just after handler stacks)
	MOV  $runtime·end(SB), A0
	ADD  $(const_maxHarts*handlerStackSize+15), A0
	AND  $~15, A0
	MOV  $runtime·ramend(SB), A1
	SUB  A0, A1, A5  // size of available memory (DMA capable)

	// estimate the space need for non-heap allocations
	SRL  $(const__PageShift+1), A5, A4
	MOV  $mspan__size, A2
	MUL  A2, A4
	ADD  $persistAllocMin, A4

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

	// run other harts
	MOV   $·waitInit(SB), A0
	MOVW  ZERO, (A0)

	// allocate g0 for m0
	MOV   $(2*const__StackMin), A0
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

	// newg to g
	MOV  g, A1
	MOV  A0, g

	// fix harts[0].gh
	ADD   $cpuctx_mh, A1, A0
	MOV   A0, g_m(A1)  // harts[0].gh.m = harts[0].mh
	CSRW  (a1, mscratch)

	// switch to the user mode
	MOV    $(1<<MPIEn), S0
	CSRS   (s0, mstatus)
	AUIPC  $0, A0
	ADD    $16, A0  // A0 must point just after MRET
	CSRW   (a0, mepc)
	MRET

	// K210: Don't be surprised the MIE is cleared though we set the MPIE before
	// MRET. It's described in RISC-V 1.9.1 spec (changed in 1.10 and fixed in
	// Rocket Chip by 29414f3a239174201938a345ac8565726892fdbb commit). Despite
	// this, machine mode interrupts are globally enabled because we are now in
	// user mode.

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
