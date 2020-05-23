// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "syscall_noos.h"
#include "asm_riscv64.h"


// func cpuid() int
TEXT ·cpuid(SB),NOSPLIT|NOFRAME,$0
	CSRR  (mhartid, a0)
	MOV   A0, ret+0(FP)
	RET


DATA runtime·interruptHandlers+(0*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·interruptHandlers+(1*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·interruptHandlers+(2*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·interruptHandlers+(3*8)(SB)/8, $·softwareInterruptHandler(SB)
DATA runtime·interruptHandlers+(4*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·interruptHandlers+(5*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·interruptHandlers+(6*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·interruptHandlers+(7*8)(SB)/8, $·timerInterruptHandler(SB)
DATA runtime·interruptHandlers+(8*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·interruptHandlers+(9*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·interruptHandlers+(10*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·interruptHandlers+(11*8)(SB)/8, $·externalInterruptHandler(SB)
#define interruptHandlersSize (12*8)
GLOBL runtime·interruptHandlers(SB), RODATA, $interruptHandlersSize

DATA runtime·exceptionHandlers+(0*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·exceptionHandlers+(1*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·exceptionHandlers+(2*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·exceptionHandlers+(3*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·exceptionHandlers+(4*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·exceptionHandlers+(5*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·exceptionHandlers+(6*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·exceptionHandlers+(7*8)(SB)/8, $·defaultHandler(SB)
DATA runtime·exceptionHandlers+(8*8)(SB)/8, $·environmentCallHandler(SB)
DATA runtime·exceptionHandlers+(9*8)(SB)/8, $·environmentCallHandler(SB)
DATA runtime·exceptionHandlers+(10*8)(SB)/8, $·environmentCallHandler(SB)
DATA runtime·exceptionHandlers+(11*8)(SB)/8, $·environmentCallHandler(SB)
#define exceptionHandlersSize (12*8)
GLOBL runtime·exceptionHandlers(SB), RODATA, $exceptionHandlersSize


#define sysMaxArgs (24+8)
#define envCallFrameSize (sysMaxArgs+3*8)

#define _LR (0*8)
#define _A0 (1*8)
#define _mstatus (2*8)
#define _mepc (3*8)
#define _mie (4*8)
#define trapCtxSize (5*8)


// The RISC-V Instruction Set Manual Volume II: Privileged Architecture defines
// the following decreasing interrupt priority order:
//
// MEI, MSI, MTI, SEI, SSI, STI, UEI, USI, UTI
//
// That's a hardware priority order conclusive in case of multiple simultaneous
// pending interrupts provided that all are enabled in the mie register.
//
// The trapHandler supports nested interrupts and implements slightly different
// software order using mie register to mask lower priority interrupts:
//
// MEI, MTI, MSI
//
// We don't support supervisor or user mode interrupts.
TEXT runtime·trapHandler(SB),NOSPLIT|NOFRAME,$0
	// At this point the interrupts are globaly disabled (mstatus.MIE=0).
	// We want to enable higher priority interrupts as soon as possible.
	// Be carefult to don't clobber T6 (A0) and A3-A5 (syscall args).

	// mscratch contains &cpuctx if trap from thread mode, 0 for nested trap
	CSRRW  (a0, mscratch, a0)  // swap A0 with cpuctx in mscratch

	// setup g and SP for handler mode, save thread ones to cpuctx.gh.sched
	BEQ  ZERO, A0, nestedTrap
	MOV  X2, (g_sched+gobuf_sp)(A0)
	MOV  g, (g_sched+gobuf_g)(A0)
	MOV  (g_stack+stack_hi)(A0), X2
	MOV  A0, g
nestedTrap:

	// save trap context, free another register (LR)
	ADD     $-trapCtxSize, X2
	MOV     LR, _LR(X2)
	SLTU    A0, ZERO, LR       // calculate fromThread flag
	CSRRWI  (0, mscratch, a0)  // set mscratch=0
	MOV     A0, _A0(X2)        // save original A0 content
	CSRR    (mstatus, a0)
	OR      $(1<<MPIEn), A0  // fix RISC-V <1.10 behavior if trap from user mode
	MOV     A0, _mstatus(X2)
	CSRR    (mepc, a0)
	OR      LR, A0
	MOV     A0, _mepc(X2)
	// mie will be saved below

	// mask same or lower priority interrupts (always mask software interrupts)
	CSRR  (mcause, a0)
	SRA   $63, A0, LR
	AND   LR, A0  // interrupt: A0=mcause, exception: A0=0
	MOV   $~1, LR
	SLL   A0, LR     // only 6 lower bits of A0 are used as shift amount
	AND   $~0xF, LR  // always mask software interrupts
	CSRR  (mie, a0)
	MOV   A0, _mie(X2)
	AND   LR, A0
	CSRW  (a0, mie)

	// enable interrupts
	CSRR   (mcause, a0)  // read mcause before enable
	CSRSI  (8, mstatus)  // set mstatus.MIE=1

	// jump to the exception/interrupt handler passing mcause*8 in A0
	BGE  A0, ZERO, handleException
handleInterrupt:
	SLL  $3, A0
	MOV  $interruptHandlersSize, LR
	BGE  A0, LR, unsupported
	MOV  $·interruptHandlers(SB), LR
	ADD  A0, LR
	MOV  (LR), LR
	JMP  (LR)
handleException:
	SLL  $3, A0
	MOV  $exceptionHandlersSize, LR
	BGE  A0, LR, unsupported
	MOV  $·exceptionHandlers(SB), LR
	ADD  A0, LR
	MOV  (LR), LR
	JMP  (LR)

unsupported:
	EBREAK
	JMP  -1(PC)


TEXT runtime·interruptReturn(SB),NOSPLIT|NOFRAME,$0
	MOV  _LR(X2), LR  // restore LR

	// restore CSRs
	MOV   _mstatus(X2), A0
	CSRW  (a0, mstatus)  // disables interrupts
	MOV   _mie(X2), A0
	CSRW  (a0, mie)
	MOV   _mepc(X2), A0
	CSRW  (a0, mepc)
	AND   $1, A0
	BEQ   ZERO, A0, returnToHandler

	// return to thread
	MOV   _A0(X2), A0                // restore A0
	MOV   (g_sched+gobuf_sp)(g), X2  // restore thread SP
	CSRW  (G, mscratch)              // cpuctx to mscratch
	MOV   (g_sched+gobuf_g)(g), g    // restore thread g
	MRET

returnToHandler:
	MOV  _A0(X2), A0  // restore A0
	ADD  $trapCtxSize, X2
	MRET


TEXT runtime·softwareInterruptHandler(SB),NOSPLIT|NOFRAME,$0

	// if cpuctx.exe == nil then context saved by environmentCallHandler
	MOVW  (cpuctx_exe)(g), A0
	BEQ   ZERO, A0, contextSaved

	// save most of the GPRs
	ADD   $m_mOS, A0
	CALL  ·saveGPRs(SB)
	// save the remaining ones: LR, SP, g, status, mepc
	MOV  _LR(X2), A1
	MOV  _A0(X2), A2
	MOV  _mstatus(X2), S0
	SRL  $7, S0  // MPP field is in a very unfortunate place
	AND  $(3<<(MPPn-7)), S0
	MOV  _mepc(X2), S1
	MOV  (g_sched+gobuf_sp)(g), A3
	MOV  (g_sched+gobuf_g)(g), A4
	MOV  A1, (const_numGPRS*8-32)(A0)  // LR
	MOV  A3, (const_numGPRS*8-24)(A0)  // SP
	MOV  A4, (const_numGPRS*8-16)(A0)  // g
	MOV  A2, (const_numGPRS*8-8)(A0)   // A0
	MOV  S0, (m_tls+const_mstatus*8-m_mOS)(A0)
	MOV  S1, (m_tls+const_mepc*8-m_mOS)(A0)

contextSaved:
	// clear software interrupt
	MOV    $msip, A0
	CSRR   (mhartid, a1)
	SLL    $2, A1  // msip registers are 32-bit
	ADD    A1, A0
	MOVW   ZERO, (A0)
	FENCE  // ensure clearing happens before checking nanotime and futexes

	// enter scheduler
	CALL  ·curcpuRunScheduler(SB)

	// load cpuctx.exe
	MOV  (cpuctx_exe)(g), A0
	ADD  $m_mOS, A0

	// check context size
	MOV  (m_tls+const_mstatus*8-m_mOS)(A0), S0
	AND  $const_thrSmallCtx, S0
	BNE  ZERO, S0, smallCtx
	// no need to restore FPRs if exe didn't changed (cpuctx.newexe == false)
	MOVB  (cpuctx_newexe)(g), A1
	BEQ   ZERO, A1, oldexe
	CALL  ·restoreFPRs(SB)
oldexe:
	CALL  ·restoreGPRs(SB)  // restore most of the GPRs
smallCtx:
	MOVB  ZERO, (cpuctx_newexe)(g)  // clear cpuctx.newexe

	// thread scheduler works at lowest interrupt priority level so it always
	// returns to the thread mode

	// load thread status
	MOV  (m_tls+const_mstatus*8-m_mOS)(A0), LR
	AND  $(3<<(MPPn-7)), LR  // select priority field

	// disable interrupts
	CSRCI  ((1<<MIEn), mstatus)

	// restore CSRs
	CSRW  (G, mscratch)
	MOV   _mstatus(X2), g
	SRL   $7, g
	OR    LR, g
	SLL   $7, g
	CSRW  (G, mstatus)
	MOV   (m_tls+const_mepc*8-m_mOS)(A0), g
	CSRW  (G, mepc)
	MOV   _mie(X2), g
	CSRW  (G, mie)

	// restore remaining GPRs
	MOV  (const_numGPRS*8-32)(A0), LR
	MOV  (const_numGPRS*8-24)(A0), X2
	MOV  (const_numGPRS*8-16)(A0), g
	MOV  (const_numGPRS*8-8)(A0), A0

	MRET


TEXT runtime·timerInterruptHandler(SB),NOSPLIT|NOFRAME,$0

	// clear timer interrupt
	MOV   $mtimecmp, A0
	CSRR  (mhartid, lr)
	SLL   $3, LR
	ADD   LR, A0
	MOV   $-1, LR
	MOV   LR, (A0)

	// rise software interrupt
	MOV   $msip, A0
	CSRR  (mhartid, lr)
	SLL   $2, LR  // msip registers are 32-bit
	ADD   LR, A0
	MOV   $1, LR
	MOVW  LR, (A0)

	JMP  ·interruptReturn(SB)


TEXT runtime·externalInterruptHandler(SB),NOSPLIT|NOFRAME,$0
	EBREAK
	JMP  -1(PC)
	JMP  ·interruptReturn(SB)


// System call is like oridnary function call so all registers except LR are
// caller save (Go ABI0). The tiny wrapper over ECALL instruction add
// additional parameters in A3-A5 registers:
//
// A3: syscall number
// A4: argument data size on the stack (+8 for frame-pointer)
// A5: return data size on the stack
TEXT runtime·environmentCallHandler(SB),NOSPLIT|NOFRAME,$0

	// check the syscall number
	MOV   $SYS_NUM, A0
	BGEU  A3, A0, badSyscall
	MOV   $SYS_LAST_FAST, S0

	// determine the caller stack
	MOV  _mepc(X2), S1
	AND  $1, S1, A0  // fromThread flag
	BEQ  ZERO, A0, currentStack

	// saved stack (called from thread)
	MOV   (g_sched+gobuf_sp)(g), A0
	BGEU  S0, A3, continue  // fast syscall
	// save thread context (small): LR, SP, g, thrSmallCtx+prio, mepc
	MOV  (g_sched+gobuf_g)(g), A1
	MOV  _LR(X2), A2
	MOV  (cpuctx_exe)(g), S0
	MOV  A2, (m_mOS+const_numGPRS*8-32)(S0)  // LR
	MOV  A0, (m_mOS+const_numGPRS*8-24)(S0)  // SP
	MOV  A1, (m_mOS+const_numGPRS*8-16)(S0)  // g
	MOV  _mstatus(X2), A1
	SRL  $7, A1  // MPP field is in a very unfortunate place
	AND  $(3<<(MPPn-7)), A1
	OR   $const_thrSmallCtx, A1
	MOV  A1, (m_tls+const_mstatus*8)(S0)
	MOV  S1, (m_tls+const_mepc*8)(S0)
	JMP  continue

currentStack: // called from handler
	BLTU  S0, A3, slowSyscallFromHandler  // handlers can use fast syscalls only
	ADD   $trapCtxSize, X2, A0

continue:
	// make a space on the stack for arguments + 3 registers
	ADD  $-envCallFrameSize, X2

	// copy arguments from the caller's stack
	MOV   $·duffcopy+2048(SB), A2
	SLL   $1, A4
	SUB   A4, A2
	MOV   X2, A1
	CALL  A2

	// save data needed to copy the return values back to the caller's stack
	MOV  A0, (sysMaxArgs+0*8)(X2)
	MOV  A1, (sysMaxArgs+1*8)(X2)
	MOV  A5, (sysMaxArgs+2*8)(X2)

	// call the service routine
	MOV   $·syscalls(SB), A0
	SLL   $3, A3
	ADD   A3, A0
	MOV   (A0), A0
	CALL  A0

	// copy the return values back to the caller's stack
	MOV   (sysMaxArgs+2*8)(X2), A4
	BEQ   ZERO, A4, nothingToCopy
	MOV   (sysMaxArgs+0*8)(X2), A1
	MOV   (sysMaxArgs+1*8)(X2), A0
	MOV   $·duffcopy+2048(SB), A2
	SLL   $1, A4
	SUB   A4, A2
	CALL  A2
nothingToCopy:

	// pop everything from the stack
	ADD  $envCallFrameSize, X2
	MOV  _LR(X2), LR
	MOV  _mstatus(X2), A0
	MOV  _mepc(X2), A1
	ADD  $4, A1  // points back to ECALL, adjust it
	MOV  _mie(X2), A2
	ADD  $trapCtxSize, X2

	// disable interrupts and restore trap context
	CSRW  (a0, mstatus)
	CSRW  (a1, mepc)
	CSRW  (a2, mie)

	// restore thread g and SP
	AND   $1, A1  // fromThread flag
	BEQ   ZERO, A1, fromHandler
	MOV   (g_sched+gobuf_sp)(g), X2
	CSRW  (G, mscratch)
	MOV   (g_sched+gobuf_g)(g), g
fromHandler:

	MRET

badSyscall:
	EBREAK  // bad syscall number
	JMP     -1(PC)

slowSyscallFromHandler:
	EBREAK  // syscall not allowed in handler mode
	JMP     -1(PC)


TEXT runtime·defaultHandler(SB),NOSPLIT|NOFRAME,$0
	EBREAK
	JMP  -1(PC)


// saveGPRs saves general purpose registers to the memory pointed by A0
TEXT runtime·saveGPRs(SB),NOSPLIT|NOFRAME,$0
	// LR saved separately
	// SP saved separately
	MOV  GP, (0*8)(A0)
	// g saved separately
	MOV  T0, (1*8)(A0)
	MOV  T1, (2*8)(A0)
	MOV  T2, (3*8)(A0)
	MOV  S0, (4*8)(A0)
	MOV  S1, (5*8)(A0)
	// A0 saved separately
	MOV  A1, (6*8)(A0)
	MOV  A2, (7*8)(A0)
	MOV  A3, (8*8)(A0)
	MOV  A4, (9*8)(A0)
	MOV  A5, (10*8)(A0)
	MOV  A6, (11*8)(A0)
	MOV  A7, (12*8)(A0)
	MOV  S2, (13*8)(A0)
	MOV  S3, (14*8)(A0)
	MOV  S4, (15*8)(A0)
	MOV  S5, (16*8)(A0)
	MOV  S6, (17*8)(A0)
	MOV  S7, (18*8)(A0)
	MOV  S8, (19*8)(A0)
	MOV  S9, (20*8)(A0)
	MOV  S10, (21*8)(A0)
	MOV  S11, (22*8)(A0)
	MOV  T3, (23*8)(A0)
	MOV  T4, (24*8)(A0)
	MOV  T5, (25*8)(A0)
	MOV  TMP, (26*8)(A0)
	RET


TEXT runtime·restoreGPRs(SB),NOSPLIT|NOFRAME,$0
	// LR loaded separately
	// SP loaded separately
	MOV  (0*8)(A0), GP
	// g loaded separately
	MOV  (1*8)(A0), T0
	MOV  (2*8)(A0), T1
	MOV  (3*8)(A0), T2
	MOV  (4*8)(A0), S0
	MOV  (5*8)(A0), S1
	// A0 loaded separately
	MOV  (6*8)(A0), A1
	MOV  (7*8)(A0), A2
	MOV  (8*8)(A0), A3
	MOV  (9*8)(A0), A4
	MOV  (10*8)(A0), A5
	MOV  (11*8)(A0), A6
	MOV  (12*8)(A0), A7
	MOV  (13*8)(A0), S2
	MOV  (14*8)(A0), S3
	MOV  (15*8)(A0), S4
	MOV  (16*8)(A0), S5
	MOV  (17*8)(A0), S6
	MOV  (18*8)(A0), S7
	MOV  (19*8)(A0), S8
	MOV  (20*8)(A0), S9
	MOV  (21*8)(A0), S10
	MOV  (22*8)(A0), S11
	MOV  (23*8)(A0), T3
	MOV  (24*8)(A0), T4
	MOV  (25*8)(A0), T5
	MOV  (26*8)(A0), TMP
	RET


// curcpuSavectxSched saves floating-point registers to the memory at
// A0 + numGPRS*8 using S0 as scratch register
TEXT ·curcpuSavectxSched(SB),NOSPLIT|NOFRAME,$0
	MOVW  (cpuctx_exe)(g), A0
	ADD   $m_mOS, A0
saveFPRs: // calls from assembly functions can enter here
	CSRR  (fcsr, s0)
	MOV   S0, ((0+const_numGPRS)*8)(A0)
	MOVD  F0, ((1+const_numGPRS)*8)(A0)
	MOVD  F1, ((2+const_numGPRS)*8)(A0)
	MOVD  F2, ((3+const_numGPRS)*8)(A0)
	MOVD  F3, ((4+const_numGPRS)*8)(A0)
	MOVD  F4, ((5+const_numGPRS)*8)(A0)
	MOVD  F5, ((6+const_numGPRS)*8)(A0)
	MOVD  F6, ((7+const_numGPRS)*8)(A0)
	MOVD  F7, ((8+const_numGPRS)*8)(A0)
	MOVD  F8, ((9+const_numGPRS)*8)(A0)
	MOVD  F9, ((10+const_numGPRS)*8)(A0)
	MOVD  F10, ((11+const_numGPRS)*8)(A0)
	MOVD  F11, ((12+const_numGPRS)*8)(A0)
	MOVD  F12, ((13+const_numGPRS)*8)(A0)
	MOVD  F13, ((14+const_numGPRS)*8)(A0)
	MOVD  F14, ((15+const_numGPRS)*8)(A0)
	MOVD  F15, ((16+const_numGPRS)*8)(A0)
	MOVD  F16, ((17+const_numGPRS)*8)(A0)
	MOVD  F17, ((18+const_numGPRS)*8)(A0)
	MOVD  F18, ((19+const_numGPRS)*8)(A0)
	MOVD  F19, ((20+const_numGPRS)*8)(A0)
	MOVD  F20, ((21+const_numGPRS)*8)(A0)
	MOVD  F21, ((22+const_numGPRS)*8)(A0)
	MOVD  F22, ((23+const_numGPRS)*8)(A0)
	MOVD  F23, ((24+const_numGPRS)*8)(A0)
	MOVD  F24, ((25+const_numGPRS)*8)(A0)
	MOVD  F25, ((26+const_numGPRS)*8)(A0)
	MOVD  F26, ((27+const_numGPRS)*8)(A0)
	MOVD  F27, ((28+const_numGPRS)*8)(A0)
	MOVD  F28, ((29+const_numGPRS)*8)(A0)
	MOVD  F29, ((30+const_numGPRS)*8)(A0)
	MOVD  F30, ((31+const_numGPRS)*8)(A0)
	MOVD  F31, ((32+const_numGPRS)*8)(A0)
	RET


TEXT runtime·restoreFPRs(SB),NOSPLIT|NOFRAME,$0
	MOV   ((0+const_numGPRS)*8)(A0), S0
	CSRW  (s0, fcsr)
	MOVD  ((1+const_numGPRS)*8)(A0), F0
	MOVD  ((2+const_numGPRS)*8)(A0), F1
	MOVD  ((3+const_numGPRS)*8)(A0), F2
	MOVD  ((4+const_numGPRS)*8)(A0), F3
	MOVD  ((5+const_numGPRS)*8)(A0), F4
	MOVD  ((6+const_numGPRS)*8)(A0), F5
	MOVD  ((7+const_numGPRS)*8)(A0), F6
	MOVD  ((8+const_numGPRS)*8)(A0), F7
	MOVD  ((9+const_numGPRS)*8)(A0), F8
	MOVD  ((10+const_numGPRS)*8)(A0), F9
	MOVD  ((11+const_numGPRS)*8)(A0), F10
	MOVD  ((12+const_numGPRS)*8)(A0), F11
	MOVD  ((13+const_numGPRS)*8)(A0), F12
	MOVD  ((14+const_numGPRS)*8)(A0), F13
	MOVD  ((15+const_numGPRS)*8)(A0), F14
	MOVD  ((16+const_numGPRS)*8)(A0), F15
	MOVD  ((17+const_numGPRS)*8)(A0), F16
	MOVD  ((18+const_numGPRS)*8)(A0), F17
	MOVD  ((19+const_numGPRS)*8)(A0), F18
	MOVD  ((20+const_numGPRS)*8)(A0), F19
	MOVD  ((21+const_numGPRS)*8)(A0), F20
	MOVD  ((22+const_numGPRS)*8)(A0), F21
	MOVD  ((23+const_numGPRS)*8)(A0), F22
	MOVD  ((24+const_numGPRS)*8)(A0), F23
	MOVD  ((25+const_numGPRS)*8)(A0), F24
	MOVD  ((26+const_numGPRS)*8)(A0), F25
	MOVD  ((27+const_numGPRS)*8)(A0), F26
	MOVD  ((28+const_numGPRS)*8)(A0), F27
	MOVD  ((29+const_numGPRS)*8)(A0), F26
	MOVD  ((30+const_numGPRS)*8)(A0), F29
	MOVD  ((31+const_numGPRS)*8)(A0), F30
	MOVD  ((32+const_numGPRS)*8)(A0), F31
	RET


// func curcpuSleep()
TEXT ·curcpuSleep(SB),NOSPLIT|NOFRAME,$0-0
	WFI
	RET


// func syssetprivlevel(newlevel int) (oldlevel, errno int)
TEXT ·syssetprivlevel(SB),NOSPLIT|NOFRAME,$0-24
	MOV  newlevel+0(FP), A0
	MOV  (envCallFrameSize+_mstatus)(X2), S0
	SRL  $MPPn, S0, S1
	AND  $3, S1
	MOV  $3, A1
	SUB  S1, A1, S1
	MOV  S1, oldlevel+8(FP)

	BLTU  A1, A0, badPrivLevel
	SUB   A0, A1, S1
	SLL   $MPPn, S1
	SLL   $MPPn, A1
	XOR   $-1, A1
	AND   A1, S0
	OR    S1, S0
	MOV   S0, (envCallFrameSize+_mstatus)(X2)
	MOV   ZERO, errno+16(FP)
	RET
badPrivLevel:
	MOV  $0, S0
	BLT  A0, ZERO, 2(PC)
	MOV  $2, S0  // rtos.ErrBadPrivLevel
	MOV  S0, errno+16(FP)
	RET
