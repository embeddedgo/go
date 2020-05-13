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
	CSRR  (MHARTID, a0)
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

#define trapLR (0*8)
#define trapA0 (1*8)
#define trapMSTATUS (2*8)
#define trapMEPC (3*8)
#define trapMIE (4*8)
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

	// MSCRATCH contains &cpuctx if trap from thread mode, 0 for nested trap
	CSRRW  (a0, MSCRATCH, a0)  // swap A0 with cpuctx in MSCRATCH

	// setup g and SP for handler mode, save thread ones to cpuctx.gh.sched.sp
	BEQ  ZERO, A0, nested
	MOV  X2, (g_sched+gobuf_sp)(A0)
	MOV  g, (g_sched+gobuf_g)(A0)
	MOV  (g_stack+stack_hi)(A0), X2
	MOV  A0, g
nested:

	// save trap context, free another register (LR)
	ADD     $-trapCtxSize, X2
	MOV     LR, trapLR(X2)
	CSRRWI  (0, MSCRATCH, lr)  // set MSCRATCH=0
	MOV     LR, trapA0(X2)     // save original A0 content
	CSRR    (MSTATUS, lr)
	MOV     LR, trapMSTATUS(X2)
	SLTU    A0, ZERO, LR  // calculate FROM_THREAD flag
	CSRR    (MEPC, a0)
	OR      LR, A0
	MOV     A0, trapMEPC(X2)
	// MIE will be saved below

	// mask same or lower priority interrupts, mask <=MSI in case of exception
	CSRR  (MCAUSE, a0)
	SRA   $63, A0, LR
	AND   LR, A0  // interrupt: A0=MCAUSE, exception: A0=0
	MOV   $~1, LR
	SLL   A0, LR     // only 6 lower bits of A0 are used as shift amount
	AND   $~0xF, LR  // always disable software interrupts
	CSRR  (MIE, a0)
	MOV   A0, trapMIE(X2)
	AND   LR, A0
	CSRW  (a0, MIE)

	// enable interrupts
	CSRR   (MCAUSE, a0)  // read MCAUSE before enable
	CSRSI  (8, MSTATUS)  // set MSTATUS.MIE=1

	// jump to exception/interrupt handler passing MCAUSE*8 in A0
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


TEXT runtime·softwareInterruptHandler(SB),NOSPLIT|NOFRAME,$0
	EBREAK
	JMP  -1(PC)


TEXT runtime·timerInterruptHandler(SB),NOSPLIT|NOFRAME,$0
	EBREAK
	JMP  -1(PC)


TEXT runtime·externalInterruptHandler(SB),NOSPLIT|NOFRAME,$0
	EBREAK
	JMP  -1(PC)


// System call is like oridnary function call so all registers except LR are
// caller save (Go ABI0). The tiny wrappers over ECALL instruction add
// additional parameters in A3-A5 registers:
//
// A3: syscall number
// A4: argument data size on the stack (+8 for frame-pointer)
// A5: return data size on the stack

TEXT runtime·environmentCallHandler(SB),NOSPLIT|NOFRAME,$0

	// check the syscall number
	MOV   $SYS_NUM, A0
	BGEU  A3, A0, badSyscall

	// determine the caller stack
	MOV  trapMEPC(X2), A0
	AND  $1, A0  // FROM_THREAD flag
	BEQ  ZERO, A0, currentStack
	MOV  (g_sched+gobuf_sp)(g), A0
	JMP  savedStack
currentStack:
	MOV   $SYS_LAST_FAST, A0
	BLTU  A0, A3, sysCallFromHandler
	ADD   $trapCtxSize, X2, A0
savedStack:

	// make a space on the stack for arguments + 3 registers
	ADD  $-(sysMaxArgs+3*8), X2

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
	ADD  $(sysMaxArgs+3*8), X2
	MOV  trapLR(X2), LR
	MOV  trapMSTATUS(X2), A0
	SLL  $(64-12), A0
	SRL  $(64-12), A0
	MOV  trapMEPC(X2), A1
	ADD  $4, A1  // points back to ECALL, adjust it
	MOV  trapMIE(X2), A2
	ADD  $trapCtxSize, X2

	// disable interrupts
	CSRRCI  ((1<<MIEn), MSTATUS, a3)  // A3=MSTATUS

	// restore trap context
	SRL   $12, A3
	SLL   $12, A3
	OR    A0, A3
	CSRW  (a3, MSTATUS)
	CSRW  (a1, MEPC)
	CSRW  (a2, MIE)

	// restore thread g and SP
	AND   $1, A1  // FROM_THREAD flag
	BEQ   ZERO, A1, fromHandler
	MOV   (g_sched+gobuf_sp)(g), X2
	CSRW  (G, MSCRATCH)
	MOV   (g_sched+gobuf_g)(g), g
fromHandler:

	MRET


badSyscall:
	EBREAK  // bad syscall number
	JMP     -1(PC)

sysCallFromHandler:
	EBREAK  // syscall not allowed in handler mode
	JMP     -1(PC)


TEXT runtime·defaultHandler(SB),NOSPLIT|NOFRAME,$0
	EBREAK
	JMP  -1(PC)


TEXT runtime·saveRegs(SB),NOSPLIT|NOFRAME,$0
	// LR saved before
	MOV  GP, (0*8)(A0)
	MOV  T0, (1*8)(A0)
	MOV  T1, (2*8)(A0)
	MOV  T2, (3*8)(A0)
	MOV  S0, (4*8)(A0)
	MOV  S1, (5*8)(A0)
	// A0 saved before
	MOV   A1, (6*8)(A0)
	MOV   A2, (7*8)(A0)
	MOV   A3, (8*8)(A0)
	MOV   A4, (9*8)(A0)
	MOV   A5, (10*8)(A0)
	MOV   A6, (11*8)(A0)
	MOV   A7, (12*8)(A0)
	MOV   S2, (13*8)(A0)
	MOV   S3, (14*8)(A0)
	MOV   S4, (15*8)(A0)
	MOV   S5, (16*8)(A0)
	MOV   S6, (17*8)(A0)
	MOV   S7, (18*8)(A0)
	MOV   S8, (19*8)(A0)
	MOV   S9, (20*8)(A0)
	MOV   S10, (21*8)(A0)
	MOV   S11, (22*8)(A0)
	MOV   T3, (23*8)(A0)
	MOV   T4, (24*8)(A0)
	MOV   T5, (25*8)(A0)
	MOV   TMP, (26*8)(A0)  // must be in sync with numGPRS
	MOVD  F0, ((0+const_numGPRS)*8)(A0)
	MOVD  F1, ((1+const_numGPRS)*8)(A0)
	MOVD  F2, ((2+const_numGPRS)*8)(A0)
	MOVD  F3, ((3+const_numGPRS)*8)(A0)
	MOVD  F4, ((4+const_numGPRS)*8)(A0)
	MOVD  F5, ((5+const_numGPRS)*8)(A0)
	MOVD  F6, ((6+const_numGPRS)*8)(A0)
	MOVD  F7, ((7+const_numGPRS)*8)(A0)
	MOVD  F8, ((8+const_numGPRS)*8)(A0)
	MOVD  F9, ((9+const_numGPRS)*8)(A0)
	MOVD  F10, ((10+const_numGPRS)*8)(A0)
	MOVD  F11, ((11+const_numGPRS)*8)(A0)
	MOVD  F12, ((12+const_numGPRS)*8)(A0)
	MOVD  F13, ((13+const_numGPRS)*8)(A0)
	MOVD  F14, ((14+const_numGPRS)*8)(A0)
	MOVD  F15, ((15+const_numGPRS)*8)(A0)
	MOVD  F16, ((16+const_numGPRS)*8)(A0)
	MOVD  F17, ((17+const_numGPRS)*8)(A0)
	MOVD  F18, ((18+const_numGPRS)*8)(A0)
	MOVD  F19, ((19+const_numGPRS)*8)(A0)
	MOVD  F20, ((20+const_numGPRS)*8)(A0)
	MOVD  F21, ((21+const_numGPRS)*8)(A0)
	MOVD  F22, ((22+const_numGPRS)*8)(A0)
	MOVD  F23, ((23+const_numGPRS)*8)(A0)
	MOVD  F24, ((24+const_numGPRS)*8)(A0)
	MOVD  F25, ((25+const_numGPRS)*8)(A0)
	MOVD  F26, ((26+const_numGPRS)*8)(A0)
	MOVD  F27, ((27+const_numGPRS)*8)(A0)
	MOVD  F28, ((28+const_numGPRS)*8)(A0)
	MOVD  F29, ((29+const_numGPRS)*8)(A0)
	MOVD  F30, ((30+const_numGPRS)*8)(A0)
	MOVD  F31, ((31+const_numGPRS)*8)(A0)
	CSRR  (FCSR, a1)
	MOV   A1, ((32+const_numGPRS)*8)(A0)

	RET
