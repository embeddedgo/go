// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "syscall_noos.h"

#include "asm_mips64.h"

#define sysMaxArgs (48+8)

TEXT runtime·intvector(SB),NOSPLIT|NOFRAME,$0
	JMP ·inthandler(SB)

// main exception handler
//
// interrupt exceptions are disabled by EXL=1 but other exceptions can
// still occur.
TEXT runtime·inthandler(SB),NOSPLIT|NOFRAME,$0
	// R26 and R27 are free, see runtime/asm_mips64x.s
	// switch to special ISR goroutine
	MOVV  $·cpu0(SB), R26
	BNE   R26, g, fromThread
	MOVV  $1, R27  // fromHandler flag
	JMP   fromHandler

fromThread:
	MOVV  $0, R27
	MOVV  R29, (g_sched+gobuf_sp)(R26)
	MOVV  g, (g_sched+gobuf_g)(R26)
	MOVV  (g_stack+stack_hi)(R26), R29
	MOVV  R26, g

fromHandler:
	// save exception context
	SUB   $excCtxSize, R29
	MOVV  R31, _LR(R29)
	MOVV  M(C0_CAUSE), R26
	MOVV  R26, _mcause(R29)
	MOVV  M(C0_EPC), R26
	OR    R27, R26 // encode fromHandler flag in EPC
	MOVV  R26, _mepc(R29)

	// TODO reenable interrupts?  they might clobber R26, R27

	// branch depending on exception cause
	MOVV  _mcause(R29), R26
	AND   $CAUSE_EXC_MASK, R26

	MOVV  $CAUSE_EXC_SYSCALL, R27
	BEQ   R26, R27, syscall

	MOVV  $CAUSE_EXC_INTERRUPT, R27
	BEQ   R26, R27, interrupt

	JMP  fatal

syscall:
	// R8 - syscall number,
	// R9 - argument data size on the stack (+8 for frame-pointer),
	// R10 - return data size on the stack.
	//
	// Now that we know we are in a syscall, we are free to use any
	// registers.  The caller should expect them to be clobbered after a
	// function call.

	// determine caller stack
	MOVV  _mepc(R29), R26
	AND   $1, R26, R27  // fromHandler flag
	BNE   R27, R0, currentStack

	MOVV  (g_sched+gobuf_sp)(g), R1 // duffcopy src thread
	JMP   continue

currentStack:
	ADD   $excCtxSize, R29, R1 // duffcopy src handler

continue:
	// 3 extra registers to preserve src, dst and size of result
	SUB   $sysMaxArgs+3*8, R29

	// copy arguments from the caller's stack
	MOVV   $·duffcopy<ABIInternal>+2048(SB), R26
	SLL    $1, R9
	SUB    R9, R26
	MOVV   R29, R2 // duffcopy dst
	JAL    (R26)

	// save data needed to copy the return values back to the caller's stack
	MOVV   R1, (sysMaxArgs+0*8)(R29)
	MOVV   R2, (sysMaxArgs+1*8)(R29)
	MOVV   R10, (sysMaxArgs+2*8)(R29)

	// call the service routine
	MOVV  $·syscalls(SB), R26
	SLL   $3, R8
	ADD   R8, R26
	MOVV  (R26), R26
	JAL   (R26)

	// copy the return values back to the caller's stack
	MOVV  (sysMaxArgs+2*8)(R29), R10
	BEQ   R0, R10, nothingToCopy
	MOVV  (sysMaxArgs+0*8)(R29), R2 // duffcopy dst
	MOVV  (sysMaxArgs+1*8)(R29), R1 // duffcopy src
	MOVV  $·duffcopy<ABIInternal>+2048(SB), R26
	SLL   $1, R10
	SUB   R10, R26
	JAL   (R26)

nothingToCopy:
	ADD $sysMaxArgs+3*8, R29

	JMP  restore

interrupt:
	// for now only clear IE
	MOVV  _mcause(R29), R26
	AND  $~CAUSE_IP_MASK, R26
	MOVV R26, M(C0_CAUSE)
	JMP  restore

restore:
	// TODO disable interrupts again
	MOVV  _LR(R29), R31
	MOVV  _mcause(R29), R26
	MOVV  R26, M(C0_CAUSE)
	MOVV  _mepc(R29), R26
	AND   $1, R26, R27
	AND   $~1, R26  // remove fromHandler flag from epc
	ADD   $4, R26  // don't execute exception instruction again
	MOVV  R26, M(C0_EPC)
	ADD   $excCtxSize, R29

	BNE   R27, R0, return

	MOVV  $·cpu0(SB), R26
	MOVV  (g_sched+gobuf_sp)(R26), R29
	MOVV  (g_sched+gobuf_g)(R26), g

return:
	WORD $0x42000018 // ERET

fatal:
	BREAK
	JMP -1(PC)
	

// func sysreset(level int, addr unsafe.Pointer) bool
TEXT ·sysreset(SB),NOSPLIT|NOFRAME,$0-12
	NOP // TODO

// func syscachemaint(op int, p unsafe.Pointer, size int)
TEXT ·syscachemaint(SB),NOSPLIT,$0-12
	NOP // TODO
