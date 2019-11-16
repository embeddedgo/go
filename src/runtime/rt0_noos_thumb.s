// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"


#define PALLOC_MIN 20*1024


TEXT _rt0_thumb_noos(SB),NOSPLIT|NOFRAME,$0
	// TODO: check does we have to disable interrupts for initialization.

	// initialize data and BSS

	MOVW       R13, R0
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

	MOVW  R0, g  // using R0 instead of g above gives shorter encoding

	BL  runtime·emptyfunc(SB)  // fault if stack check is wrong
	//BL  runtime·check(SB)
	BL  runtime·osinit(SB)

	// initialize sysMem

	MOVW  $runtime·end(SB), R0
	MOVW  $runtime·ramend(SB), R1
	SUB   R0, R1, R5  // size of available memory (DMA capable)

	// estimate the space need for non-heap allocations
	MOVW  R5>>(const__PageShift+2), R4
	MOVW  $const_mspanSize, R2
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

	BL  runtime·taskerpreinit(SB)
	BL  runtime·schedinit(SB)
	BL  runtime·taskerinit(SB)

	// allocate g0 for m0 and leave gh

	SUB        $4, R13
	MOVW       $0, R0
	MOVW       $(2*const__FixedStack), R1
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
	MOVW       $0, R0
	MOVW       $8, R1
	MOVW       $runtime·mainPC(SB), R2
	MOVM.DB.W  [R0-R2], (R13)
	BL         runtime·newproc(SB)
	MOVW       $12(R13), R13

	// start this M
	BL  runtime·mstart(SB)

	UNDEF  // fail

DATA runtime·mainPC+0(SB)/4,$runtime·main(SB)
GLOBL runtime·mainPC(SB),RODATA,$4

TEXT runtime·breakpoint(SB),NOSPLIT,$0-0
	BKPT
	RET

// asminit is called by mstart1 for every new thread
TEXT runtime·asminit(SB),NOSPLIT,$0-0
	RET

//
// go-routine
//

// void gosave(Gobuf*)
// save state in Gobuf; setjmp
TEXT runtime·gosave(SB),NOSPLIT|NOFRAME,$0-4
	MOVW  buf+0(FP), R0
	MOVW  R13, gobuf_sp(R0)
	MOVW  LR, gobuf_pc(R0)
	MOVW  g, gobuf_g(R0)
	MOVW  $0, REGTMP
	MOVW  REGTMP, gobuf_lr(R0)
	MOVW  REGTMP, gobuf_ret(R0)
	// Assert ctxt is zero. See func save.
	MOVW  gobuf_ctxt(R0), R0
	CMP   R0, REGTMP
	B.EQ  2(PC)
	CALL  runtime·badctxt(SB)
	RET

// void gogo(Gobuf*)
// restore state from Gobuf; longjmp
TEXT runtime·gogo(SB),NOSPLIT,$8-4
	MOVW  buf+0(FP), R1
	MOVW  gobuf_g(R1), R0
	BL    setg<>(SB)

	// NOTE: We updated g above, and we are about to update SP.
	// Until LR and PC are also updated, the g/SP/LR/PC quadruple
	// are out of sync and must not be used as the basis of a traceback.
	// Sigprof skips the traceback when SP is not within g's bounds,
	// and when the PC is inside this function, runtime.gogo.
	// Since we are about to update SP, until we complete runtime.gogo
	// we must not leave this function. In particular, no calls
	// after this point: it must be straight-line code until the
	// final B instruction.
	// See large comment in sigprof for more details.
	MOVW  gobuf_sp(R1), R13  // restore SP==R13
	MOVW  gobuf_lr(R1), LR
	MOVW  gobuf_ret(R1), R0
	MOVW  gobuf_ctxt(R1), REGCTXT
	MOVW  $0, REGTMP
	MOVW  REGTMP, gobuf_sp(R1)  // clear to help garbage collector
	MOVW  REGTMP, gobuf_ret(R1)
	MOVW  REGTMP, gobuf_lr(R1)
	MOVW  REGTMP, gobuf_ctxt(R1)
	MOVW  gobuf_pc(R1), REGTMP
	CMP   REGTMP, REGTMP  // set condition codes for == test, needed by stack split
	B     (REGTMP)

// func mcall(fn func(*g))
// Switch to m->g0's stack, call fn(g).
// Fn must never return. It should gogo(&g->sched)
// to keep running g.
TEXT runtime·mcall(SB),NOSPLIT|NOFRAME,$0-4
	// Save caller state in g->sched.
	MOVW  R13, (g_sched+gobuf_sp)(g)
	MOVW  LR, (g_sched+gobuf_pc)(g)
	MOVW  $0, REGTMP
	MOVW  REGTMP, (g_sched+gobuf_lr)(g)
	MOVW  g, (g_sched+gobuf_g)(g)

	// Switch to m->g0 & its stack, call fn.
	MOVW  g, R1
	MOVW  g_m(g), R8
	MOVW  m_g0(R8), R0
	BL    setg<>(SB)
	CMP   g, R1
	B.NE  2(PC)
	B     runtime·badmcall(SB)
	MOVW  fn+0(FP), R0
	MOVW  (g_sched+gobuf_sp)(g), R13
	SUB   $8, R13
	MOVW  R1, 4(R13)
	MOVW  R0, REGCTXT
	MOVW  0(R0), R0
	BL    (R0)
	B     runtime·badmcall2(SB)
	RET

// systemstack_switch is a dummy routine that systemstack leaves at the bottom
// of the G stack. We need to distinguish the routine that
// lives at the bottom of the G stack from the one that lives
// at the top of the system stack because the one at the top of
// the system stack terminates the stack walk (see topofstack()).
TEXT runtime·systemstack_switch(SB),NOSPLIT,$0-0
	MOVW  $0, R0
	BL    (R0)  // clobber lr to ensure push {lr} is kept
	RET

// func systemstack(fn func())
TEXT runtime·systemstack(SB),NOSPLIT,$0-4
	MOVW  fn+0(FP), R0  // R0 = fn
	MOVW  g_m(g), R1    // R1 = m

	MOVW  m_g0(R1), R2  // R2 = g0
	CMP   g, R2
	B.EQ  noswitch

	MOVW  m_curg(R1), R3
	CMP   g, R3
	B.EQ  switch

	// Bad: g is not g0, not curg. What is it?
	// Hide call from linker nosplit analysis.
	MOVW  $runtime·badsystemstack(SB), R0
	BL    (R0)
	B     runtime·abort(SB)

switch:
	// save our state in g->sched. Pretend to
	// be systemstack_switch if the G stack is scanned.
	MOVW  $runtime·systemstack_switch(SB), R3
	ADD   $4, R3, R3  // get past push {lr}
	MOVW  R3, (g_sched+gobuf_pc)(g)
	MOVW  R13, (g_sched+gobuf_sp)(g)
	MOVW  LR, (g_sched+gobuf_lr)(g)
	MOVW  g, (g_sched+gobuf_g)(g)

	// switch to g0
	MOVW  R0, R5
	MOVW  R2, R0
	BL    setg<>(SB)
	MOVW  R5, R0
	MOVW  (g_sched+gobuf_sp)(R2), R3
	// make it look like mstart called systemstack on g0, to stop traceback
	SUB   $4, R3, R3
	MOVW  $runtime·mstart(SB), R4
	MOVW  R4, 0(R3)
	MOVW  R3, R13

	// call target function
	MOVW  R0, REGCTXT
	MOVW  0(R0), R0
	BL    (R0)

	// switch back to g
	MOVW  g_m(g), R1
	MOVW  m_curg(R1), R0
	BL    setg<>(SB)
	MOVW  (g_sched+gobuf_sp)(g), R13
	MOVW  $0, R3
	MOVW  R3, (g_sched+gobuf_sp)(g)
	RET

noswitch:
	// Using a tail call here cleans up tracebacks since we won't stop
	// at an intermediate systemstack.
	MOVW    R0, REGCTXT
	MOVW    0(R0), R0
	MOVW.P  4(R13), R14  // restore LR
	B       (R0)

/*
	*   support for morestack
	*/

// Called during function prolog when more stack is needed.
// R3 prolog's LR
// using NOFRAME means do not save LR on stack.

// The traceback routines see morestack on a g0 as being
// the top of a stack (for example, morestack calling newstack
// calling the scheduler calling newm calling gc), so we must
// record an argument size. For that purpose, it has no arguments.
TEXT runtime·morestack(SB),NOSPLIT|NOFRAME,$0-0
	// Cannot grow scheduler stack (m->g0).
	MOVW  g_m(g), R8
	MOVW  m_g0(R8), R4
	CMP   g, R4
	BNE   3(PC)
	BL    runtime·badmorestackg0(SB)
	B     runtime·abort(SB)

	// Called from f.
	// Set g->sched to context in f.
	MOVW  R13, (g_sched+gobuf_sp)(g)
	MOVW  LR, (g_sched+gobuf_pc)(g)
	MOVW  R3, (g_sched+gobuf_lr)(g)
	MOVW  REGCTXT, (g_sched+gobuf_ctxt)(g)

	// Called from f.
	// Set m->morebuf to f's caller.
	MOVW  R3, (m_morebuf+gobuf_pc)(R8)   // f's caller's PC
	MOVW  R13, (m_morebuf+gobuf_sp)(R8)  // f's caller's SP
	MOVW  g, (m_morebuf+gobuf_g)(R8)

	// Call newstack on m->g0's stack.
	MOVW    m_g0(R8), R0
	BL      setg<>(SB)
	MOVW    (g_sched+gobuf_sp)(g), R13
	MOVW    $0, R0
	MOVW.W  R0, -4(R13)  // create a call frame on g0 (saved LR)
	BL      runtime·newstack(SB)

	// Not reached, but make sure the return PC from the call to newstack
	// is still in this function, and not the beginning of the next.
	RET

TEXT runtime·morestack_noctxt(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  $0, REGCTXT
	B     runtime·morestack(SB)

// reflectcall: call a function with the given argument list
// func call(argtype *_type, f *FuncVal, arg *byte, argsize, retoffset uint32).
// we don't have variable-sized frames, so we use a small number
// of constant-sized-frame functions to encode a few bits of size in the pc.
// Caution: ugly multiline assembly macros in your future!

#define DISPATCH(NAME,MAXSIZE) \
	CMP   $MAXSIZE, R0; \
	B.HI  3(PC); \
	MOVW  $NAME(SB), R1; \
	B     (R1)

TEXT ·reflectcall(SB),NOSPLIT|NOFRAME,$0-20
	MOVW                              argsize+12(FP), R0
	DISPATCH(runtime·call16,          16)
	DISPATCH(runtime·call32,          32)
	DISPATCH(runtime·call64,          64)
	DISPATCH(runtime·call128,         128)
	DISPATCH(runtime·call256,         256)
	DISPATCH(runtime·call512,         512)
	DISPATCH(runtime·call1024,        1024)
	DISPATCH(runtime·call2048,        2048)
	DISPATCH(runtime·call4096,        4096)
	DISPATCH(runtime·call8192,        8192)
	DISPATCH(runtime·call16384,       16384)
	DISPATCH(runtime·call32768,       32768)
	DISPATCH(runtime·call65536,       65536)
	DISPATCH(runtime·call131072,      131072)
	DISPATCH(runtime·call262144,      262144)
	DISPATCH(runtime·call524288,      524288)
	DISPATCH(runtime·call1048576,     1048576)
	DISPATCH(runtime·call2097152,     2097152)
	DISPATCH(runtime·call4194304,     4194304)
	DISPATCH(runtime·call8388608,     8388608)
	DISPATCH(runtime·call16777216,    16777216)
	DISPATCH(runtime·call33554432,    33554432)
	DISPATCH(runtime·call67108864,    67108864)
	DISPATCH(runtime·call134217728,   134217728)
	DISPATCH(runtime·call268435456,   268435456)
	DISPATCH(runtime·call536870912,   536870912)
	DISPATCH(runtime·call1073741824,  1073741824)
	MOVW                              $runtime·badreflectcall(SB), R1
	B                                 (R1)

#define CALLFN(NAME,MAXSIZE) \
TEXT NAME(SB), WRAPPER, $MAXSIZE-20; \
	NO_LOCAL_POINTERS;  \
/*                  copy arguments to stack */ \
	MOVW     argptr+8(FP), R0; \
	MOVW     argsize+12(FP), R2; \
	ADD      $4, R13, R1; \
	CMP      $0, R2; \
	B.EQ     5(PC); \
	MOVBU.P  1(R0), R5; \
	MOVB.P   R5, 1(R1); \
	SUB      $1, R2, R2; \
	B        -5(PC); \
/*                  call function */ \
	MOVW    f+4(FP), REGCTXT; \
	MOVW    (REGCTXT), R0; \
	PCDATA  $PCDATA_StackMapIndex, $0; \
	BL      (R0); \
/*                  copy return values back */ \
	MOVW  argtype+0(FP), R4; \
	MOVW  argptr+8(FP), R0; \
	MOVW  argsize+12(FP), R2; \
	MOVW  retoffset+16(FP), R3; \
	ADD   $4, R13, R1; \
	ADD   R3, R1; \
	ADD   R3, R0; \
	SUB   R3, R2; \
	BL    callRet<>(SB); \
	RET

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers.
TEXT callRet<>(SB), NOSPLIT, $16-0
	MOVW  R4, 4(R13)
	MOVW  R0, 8(R13)
	MOVW  R1, 12(R13)
	MOVW  R2, 16(R13)
	BL    runtime·reflectcallmove(SB)
	RET

	CALLFN(·call16,          16)
	CALLFN(·call32,          32)
	CALLFN(·call64,          64)
	CALLFN(·call128,         128)
	CALLFN(·call256,         256)
	CALLFN(·call512,         512)
	CALLFN(·call1024,        1024)
	CALLFN(·call2048,        2048)
	CALLFN(·call4096,        4096)
	CALLFN(·call8192,        8192)
	CALLFN(·call16384,       16384)
	CALLFN(·call32768,       32768)
	CALLFN(·call65536,       65536)
	CALLFN(·call131072,      131072)
	CALLFN(·call262144,      262144)
	CALLFN(·call524288,      524288)
	CALLFN(·call1048576,     1048576)
	CALLFN(·call2097152,     2097152)
	CALLFN(·call4194304,     4194304)
	CALLFN(·call8388608,     8388608)
	CALLFN(·call16777216,    16777216)
	CALLFN(·call33554432,    33554432)
	CALLFN(·call67108864,    67108864)
	CALLFN(·call134217728,   134217728)
	CALLFN(·call268435456,   268435456)
	CALLFN(·call536870912,   536870912)
	CALLFN(·call1073741824,  1073741824)

// void jmpdefer(fn, sp);
// called from deferreturn.
// 1. grab stored LR for caller
// 2. sub 4 bytes to get back to BL deferreturn
// 3. B to fn
// TODO(rsc): Push things on stack and then use pop
// to load all registers simultaneously, so that a profiling
// interrupt can never see mismatched SP/LR/PC.
// (And double-check that pop is atomic in that way.)
TEXT runtime·jmpdefer(SB),NOSPLIT,$0-8
	MOVW  0(R13), LR
	MOVW  $-4(LR), LR  // BL deferreturn
	MOVW  fv+0(FP), REGCTXT
	MOVW  argp+4(FP), R13
	MOVW  $-4(R13), R13  // SP is 4 below argp, due to saved LR
	MOVW  0(REGCTXT), R1
	B     (R1)

// Save state of caller into g->sched. Smashes REGTMP.
TEXT gosave<>(SB),NOSPLIT|NOFRAME,$0
	MOVW  LR, (g_sched+gobuf_pc)(g)
	MOVW  R13, (g_sched+gobuf_sp)(g)
	MOVW  $0, REGTMP
	MOVW  REGTMP, (g_sched+gobuf_lr)(g)
	MOVW  REGTMP, (g_sched+gobuf_ret)(g)
	MOVW  REGTMP, (g_sched+gobuf_ctxt)(g)
	// Assert ctxt is zero. See func save.
	MOVW  (g_sched+gobuf_ctxt)(g), REGTMP
	CMP   $0, REGTMP
	B.EQ  2(PC)
	CALL  runtime·badctxt(SB)
	RET

TEXT ·asmcgocall(SB),NOSPLIT|NOFRAME,$0-12
	BKPT
	B   -1(PC)

// void setg(G*); set g. for use by needm.
TEXT runtime·setg(SB),NOSPLIT|NOFRAME,$0-4
	MOVW  gg+0(FP), R0
	B     setg<>(SB)

TEXT setg<>(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  R0, g
	RET

TEXT runtime·emptyfunc(SB),0,$0-0
	RET

TEXT runtime·abort(SB),NOSPLIT|NOFRAME,$0-0
	BKPT
	B   -1(PC)

// AES hashing not implemented for ARM
TEXT runtime·aeshash(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  $0, R0
	MOVW  (R0), R1
TEXT runtime·aeshash32(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  $0, R0
	MOVW  (R0), R1
TEXT runtime·aeshash64(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  $0, R0
	MOVW  (R0), R1
TEXT runtime·aeshashstr(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  $0, R0
	MOVW  (R0), R1

TEXT runtime·return0(SB),NOSPLIT,$0
	MOVW  $0, R0
	RET

TEXT runtime·procyield(SB),NOSPLIT|NOFRAME,$0
	MOVW  cycles+0(FP), R1
yieldloop:
	YIELD
	CMP  $0, R1
	RET.EQ
	SUB  $1, R1
	B    yieldloop

// Called from cgo wrappers, this function returns g->m->curg.stack.hi.
// Must obey the gcc calling convention.
TEXT _cgo_topofstack(SB),NOSPLIT,$8
	// R11 (REGCTXT) and g register are clobbered by load_g. They are
	// callee-save in the gcc calling convention, so save them here.
	MOVW  REGCTXT, saveR11-4(SP)
	MOVW  g, saveG-8(SP)

	BL    runtime·load_g(SB)
	MOVW  g_m(g), R0
	MOVW  m_curg(R0), R0
	MOVW  (g_stack+stack_hi)(R0), R0

	MOVW  saveG-8(SP), g
	MOVW  saveR11-4(SP), REGCTXT
	RET

// The top-most function running on a goroutine
// returns to goexit+PCQuantum.
TEXT runtime·goexit(SB),NOSPLIT|NOFRAME,$0-0
	NOP2
	BL  runtime·goexit1(SB)  // does not return
	// traceback from goexit1 must hit code range of goexit
	NOP2

// x -> x/1000000, x%1000000, called from Go with args, results on stack.
TEXT runtime·usplit(SB),NOSPLIT,$0-12
	MOVW  x+0(FP), R0
	CALL  runtime·usplitR0(SB)
	MOVW  R0, q+4(FP)
	MOVW  R1, r+8(FP)
	RET

// R0, R1 = R0/1000000, R0%1000000
TEXT runtime·usplitR0(SB),NOSPLIT,$0
	// magic multiply to avoid software divide without available m.
	// see output of go tool compile -S for x/1000000.
	MOVW   R0, R3
	MOVW   $1125899907, R1
	MULLU  R1, R0, (R0, R1)
	MOVW   R0>>18, R0
	MOVW   $1000000, R1
	MUL    R0, R1
	SUB    R1, R3, R1
	RET

TEXT runtime·sigreturn(SB),NOSPLIT,$0-0
	RET

#ifndef GOOS_nacl
// This is called from .init_array and follows the platform, not Go, ABI.
TEXT runtime·addmoduledata(SB),NOSPLIT,$0-8
	MOVW  R9, saver9-4(SP)        // The access to global variables below implicitly uses R9, which is callee-save
	MOVW  REGCTXT, saver11-8(SP)  // Likewise, R11 is the REGCTXT register, but callee-save in C ABI
	MOVW  runtime·lastmoduledatap(SB), R1
	MOVW  R0, moduledata_next(R1)
	MOVW  R0, runtime·lastmoduledatap(SB)
	MOVW  saver11-8(SP), REGCTXT
	MOVW  saver9-4(SP), R9
	RET
#endif

TEXT ·checkASM(SB),NOSPLIT,$0-1
	MOVW  $1, R3
	MOVB  R3, ret+0(FP)
	RET

// gcWriteBarrier performs a heap pointer write and informs the GC.

// gcWriteBarrier does NOT follow the Go ABI. It takes two arguments:
// - R2 is the destination of the write
// - R3 is the value being written at R2
// It clobbers condition codes.
// It does not clobber any other general-purpose registers,
// but may clobber others (e.g., floating point registers).
// The act of CALLing gcWriteBarrier will clobber R14 (LR).
TEXT runtime·gcWriteBarrier(SB),NOSPLIT|NOFRAME,$0
	// Save the registers clobbered by the fast path.
	MOVM.DB.W  [R0,R1], (R13)
	MOVW       g_m(g), R0
	MOVW       m_p(R0), R0
	MOVW       (p_wbBuf+wbBuf_next)(R0), R1
	// Increment wbBuf.next position.
	ADD   $8, R1
	MOVW  R1, (p_wbBuf+wbBuf_next)(R0)
	MOVW  (p_wbBuf+wbBuf_end)(R0), R0
	CMP   R1, R0
	// Record the write.
	MOVW  R3, -8(R1)  // Record value
	MOVW  (R2), R0    // TODO: This turns bad writes into bad reads.
	MOVW  R0, -4(R1)  // Record *slot
	// Is the buffer full? (flags set in CMP above)
	B.EQ  flush
ret:
	MOVM.IA.W  (R13), [R0,R1]
	// Do the write.
	MOVW  R3, (R2)
	// Normally RET on nacl clobbers R12, but because this
	// function has no frame it doesn't have to usual epilogue.
	RET

flush:
	// Save all general purpose registers since these could be
	// clobbered by wbBufFlush and were not saved by the caller.

	// R0 and R1 were saved at entry.
	// R7 is linker temp, so no need to save.
	// R10 is g, so preserved.
	// R13 is stack pointer.
	// R15 is PC.

	// This also sets up R2 and R3 as the arguments to wbBufFlush.
	MOVM.DB.W  [R2-R6,R8-R9,R11-R12], (R13)
	// Save R14 (LR) because the fast path above doesn't save it,
	// but needs it to RET. This is after the MOVM so it appears below
	// the arguments in the stack frame.
	MOVW.W  LR, -4(R13)

	// This takes arguments R2 and R3.
	CALL  runtime·wbBufFlush(SB)

	MOVW.P     4(R13), LR
	MOVM.IA.W  (R13), [R2-R6,R8-R9,R11-R12]
	JMP        ret

// Note: these functions use a special calling convention to save generated code space.
// Arguments are passed in registers, but the space for those arguments are allocated
// in the caller's stack frame. These stubs write the args into that stack space and
// then tail call to the corresponding runtime handler.
// The tail call makes these stubs disappear in backtraces.
TEXT runtime·panicIndex(SB),NOSPLIT,$0-8
	MOVW  R0, x+0(FP)
	MOVW  R1, y+4(FP)
	JMP   runtime·goPanicIndex(SB)
TEXT runtime·panicIndexU(SB),NOSPLIT,$0-8
	MOVW  R0, x+0(FP)
	MOVW  R1, y+4(FP)
	JMP   runtime·goPanicIndexU(SB)
TEXT runtime·panicSliceAlen(SB),NOSPLIT,$0-8
	MOVW  R1, x+0(FP)
	MOVW  R2, y+4(FP)
	JMP   runtime·goPanicSliceAlen(SB)
TEXT runtime·panicSliceAlenU(SB),NOSPLIT,$0-8
	MOVW  R1, x+0(FP)
	MOVW  R2, y+4(FP)
	JMP   runtime·goPanicSliceAlenU(SB)
TEXT runtime·panicSliceAcap(SB),NOSPLIT,$0-8
	MOVW  R1, x+0(FP)
	MOVW  R2, y+4(FP)
	JMP   runtime·goPanicSliceAcap(SB)
TEXT runtime·panicSliceAcapU(SB),NOSPLIT,$0-8
	MOVW  R1, x+0(FP)
	MOVW  R2, y+4(FP)
	JMP   runtime·goPanicSliceAcapU(SB)
TEXT runtime·panicSliceB(SB),NOSPLIT,$0-8
	MOVW  R0, x+0(FP)
	MOVW  R1, y+4(FP)
	JMP   runtime·goPanicSliceB(SB)
TEXT runtime·panicSliceBU(SB),NOSPLIT,$0-8
	MOVW  R0, x+0(FP)
	MOVW  R1, y+4(FP)
	JMP   runtime·goPanicSliceBU(SB)
TEXT runtime·panicSlice3Alen(SB),NOSPLIT,$0-8
	MOVW  R2, x+0(FP)
	MOVW  R3, y+4(FP)
	JMP   runtime·goPanicSlice3Alen(SB)
TEXT runtime·panicSlice3AlenU(SB),NOSPLIT,$0-8
	MOVW  R2, x+0(FP)
	MOVW  R3, y+4(FP)
	JMP   runtime·goPanicSlice3AlenU(SB)
TEXT runtime·panicSlice3Acap(SB),NOSPLIT,$0-8
	MOVW  R2, x+0(FP)
	MOVW  R3, y+4(FP)
	JMP   runtime·goPanicSlice3Acap(SB)
TEXT runtime·panicSlice3AcapU(SB),NOSPLIT,$0-8
	MOVW  R2, x+0(FP)
	MOVW  R3, y+4(FP)
	JMP   runtime·goPanicSlice3AcapU(SB)
TEXT runtime·panicSlice3B(SB),NOSPLIT,$0-8
	MOVW  R1, x+0(FP)
	MOVW  R2, y+4(FP)
	JMP   runtime·goPanicSlice3B(SB)
TEXT runtime·panicSlice3BU(SB),NOSPLIT,$0-8
	MOVW  R1, x+0(FP)
	MOVW  R2, y+4(FP)
	JMP   runtime·goPanicSlice3BU(SB)
TEXT runtime·panicSlice3C(SB),NOSPLIT,$0-8
	MOVW  R0, x+0(FP)
	MOVW  R1, y+4(FP)
	JMP   runtime·goPanicSlice3C(SB)
TEXT runtime·panicSlice3CU(SB),NOSPLIT,$0-8
	MOVW  R0, x+0(FP)
	MOVW  R1, y+4(FP)
	JMP   runtime·goPanicSlice3CU(SB)

// Extended versions for 64-bit indexes.
TEXT runtime·panicExtendIndex(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R0, lo+4(FP)
	MOVW  R1, y+8(FP)
	JMP   runtime·goPanicExtendIndex(SB)
TEXT runtime·panicExtendIndexU(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R0, lo+4(FP)
	MOVW  R1, y+8(FP)
	JMP   runtime·goPanicExtendIndexU(SB)
TEXT runtime·panicExtendSliceAlen(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R1, lo+4(FP)
	MOVW  R2, y+8(FP)
	JMP   runtime·goPanicExtendSliceAlen(SB)
TEXT runtime·panicExtendSliceAlenU(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R1, lo+4(FP)
	MOVW  R2, y+8(FP)
	JMP   runtime·goPanicExtendSliceAlenU(SB)
TEXT runtime·panicExtendSliceAcap(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R1, lo+4(FP)
	MOVW  R2, y+8(FP)
	JMP   runtime·goPanicExtendSliceAcap(SB)
TEXT runtime·panicExtendSliceAcapU(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R1, lo+4(FP)
	MOVW  R2, y+8(FP)
	JMP   runtime·goPanicExtendSliceAcapU(SB)
TEXT runtime·panicExtendSliceB(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R0, lo+4(FP)
	MOVW  R1, y+8(FP)
	JMP   runtime·goPanicExtendSliceB(SB)
TEXT runtime·panicExtendSliceBU(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R0, lo+4(FP)
	MOVW  R1, y+8(FP)
	JMP   runtime·goPanicExtendSliceBU(SB)
TEXT runtime·panicExtendSlice3Alen(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R2, lo+4(FP)
	MOVW  R3, y+8(FP)
	JMP   runtime·goPanicExtendSlice3Alen(SB)
TEXT runtime·panicExtendSlice3AlenU(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R2, lo+4(FP)
	MOVW  R3, y+8(FP)
	JMP   runtime·goPanicExtendSlice3AlenU(SB)
TEXT runtime·panicExtendSlice3Acap(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R2, lo+4(FP)
	MOVW  R3, y+8(FP)
	JMP   runtime·goPanicExtendSlice3Acap(SB)
TEXT runtime·panicExtendSlice3AcapU(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R2, lo+4(FP)
	MOVW  R3, y+8(FP)
	JMP   runtime·goPanicExtendSlice3AcapU(SB)
TEXT runtime·panicExtendSlice3B(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R1, lo+4(FP)
	MOVW  R2, y+8(FP)
	JMP   runtime·goPanicExtendSlice3B(SB)
TEXT runtime·panicExtendSlice3BU(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R1, lo+4(FP)
	MOVW  R2, y+8(FP)
	JMP   runtime·goPanicExtendSlice3BU(SB)
TEXT runtime·panicExtendSlice3C(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R0, lo+4(FP)
	MOVW  R1, y+8(FP)
	JMP   runtime·goPanicExtendSlice3C(SB)
TEXT runtime·panicExtendSlice3CU(SB),NOSPLIT,$0-12
	MOVW  R4, hi+0(FP)
	MOVW  R0, lo+4(FP)
	MOVW  R1, y+8(FP)
	JMP   runtime·goPanicExtendSlice3CU(SB)
