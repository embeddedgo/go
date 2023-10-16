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
// R26 and R27 are free, see runtime/asm_mips64x.s
//
// interrupt exceptions are disabled by EXL=1 but other exceptions can
// still occur.
//
// Only syscalls and interrupts are handled at the moment, all other exceptions
// are fatal.
//
// Called from thread:
// 1. save goroutine context in g.sched (sp, fp)
// 2. switch to ISR stack
// 3. save exception context on ISR stack (ra, epc, cause)
// 4. save thread context in m.mOS (sp, fp, ra, epc)
// 5. if interrupt goto 8, if syscall continue
// 6. do syscall
//    - copy args from caller thread stack to ISR stack
//    - call service routine in ISR context
//    - copy result from ISR stack to caller thread stack
// 7. if the syscall called curcpuSchedule, make a call to curcpuRunScheduler
// 8. if newexe, restore context from that thread (sp, fp, ra, epc)
// 9. return execution at epc
//
// Called from handler:
// same, but skip 1, 2, 4
TEXT runtime·inthandler(SB),NOSPLIT|NOFRAME,$0
start:
	// determine caller stack
	MOVV  $·cpu0(SB), R26
	BNE   R26, g, fromThread
	MOVV  $1, R27  // fromHandler flag
	JMP   fromHandler

fromThread:
	MOVV  $0, R27

	// save goroutine context in g.sched
	MOVV  R29, (g_sched+gobuf_sp)(R26)
	MOVV  g, (g_sched+gobuf_g)(R26)
	MOVV  (g_stack+stack_hi)(R26), R29
	// switch to special ISR goroutine
	MOVV  R26, g

fromHandler:
	// save exception context on ISR stack
	SUB   $excCtxSize, R29
	OR    $1, R31, R26 // encode smallCtx flag in ra
	MOVV  R26, _LR(R29)
	MOVV  M(C0_CAUSE), R26
	MOVV  R26, _mcause(R29)
	MOVV  M(C0_EPC), R26
	OR    R27, R26 // encode fromHandler flag in EPC
	MOVV  R26, _mepc(R29)

	// mask interrupts
	MOVV  M(C0_SR), R26
	MOVV  $~(INTR_SW|INTR_TIMER), R27
	AND   R27, R26
	MOVV  R26, M(C0_SR)

	// branch depending on exception cause
	MOVV  M(C0_CAUSE), R26
	AND   $CAUSE_EXC_MASK, R26

	MOVV  $CAUSE_EXC_SYSCALL, R27
	BEQ   R26, R27, syscall

	MOVV  $CAUSE_EXC_INTERRUPT, R27
	BEQ   R26, R27, interrupt

	JMP  fatal

// System call is like oridnary function call so all registers are caller save
// (Go ABI0).  Meaning we don't need to save GPRs and are free to use them.  The
// tiny wrapper over SYSCALL instruction adds additional parameters in A3-A5
// registers:
//
// R8: syscall number
// R9: argument data size on the stack (+8 for caller return address)
// R10: return data size on the stack
syscall:
	// if fromHandler skip saving thread context
	MOVV  _mepc(R29), R27
	ADD   $4, R27  // don't execute syscall instruction again
	MOVV  R27, _mepc(R29)
	AND   $1, R27  // fromHandler flag
	BNE   R27, R0, currentStack

	// save thread context in mOS
	// needs to be done before the syscall, cpuctx.exe might be nil
	// afterwards (e.g. in nanosleep)
	// TODO skip for fast syscall
	MOVV  (cpuctx_exe)(g), R27

	MOVV  _LR(R29), R26
	MOVV  R26, (m_mOS+mOS_ra)(R27)

	MOVV  _mepc(R29), R26
	AND   $~1, R26  // remove fromHandler flag from epc
	MOVV  R26, (m_mOS+mOS_epc)(R27)

	MOVV  (g_sched+gobuf_sp)(g), R26
	MOVV  R26, (m_mOS+mOS_sp)(R27)
	MOVV  (g_sched+gobuf_g)(g), R26
	MOVV  R26, (m_mOS+mOS_fp)(R27)

	MOVV  (g_sched+gobuf_sp)(g), R1 // duffcopy src thread
	JMP   duffcopy

currentStack:
	ADD   $excCtxSize, R29, R1 // duffcopy src handler

duffcopy:
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

	// reenable exceptions
	MOVV  M(C0_SR), R26
	AND   $~SR_EXL, R26
	MOVV  R26, M(C0_SR)

	// reminder: don't use R23, R26 or R27 when interrupts enabled

	// call the service routine
	MOVV  $·syscalls(SB), R9
	SLL   $3, R8
	ADD   R8, R9
	MOVV  (R9), R9
	JAL   (R9)

	// disable exceptions again
	MOVV  M(C0_SR), R8
	OR    $SR_EXL, R8
	MOVV  R8, M(C0_SR)

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

	// run the scheduler if the syscall wants it
	MOVB  cpuctx_schedule(g), R8  // handlers will not set this
	BEQ   R8, R0, restore
	JMP   enterScheduler

// An interrupt was caused by timer, software or externally.  This can happen on
// any instruction.  We need to save all GPRs in our context.
interrupt:
	MOVV  (cpuctx_exe)(g), R26
	MOVV  $(m_mOS+mOS_gprs)(R26), R26
	JAL   ·saveGPRs(SB)

	// save thread context in mOS
	MOVV  (cpuctx_exe)(g), R27

	MOVV  _LR(R29), R26
	AND   $~1, R26  // remove smallCtx flag from ra
	MOVV  R26, (m_mOS+mOS_ra)(R27)

	MOVV  _mepc(R29), R26
	AND   $~1, R26  // remove fromHandler flag from epc
	MOVV  R26, (m_mOS+mOS_epc)(R27)

	MOVV  (g_sched+gobuf_sp)(g), R26
	MOVV  R26, (m_mOS+mOS_sp)(R27)
	MOVV  (g_sched+gobuf_g)(g), R26
	MOVV  R26, (m_mOS+mOS_fp)(R27)

	// clear all IP (interrupt pending) bits
	// TODO only timer supported at the moment
	MOVV  M(C0_COMPARE), R26
	MOVV  R26, M(C0_COMPARE)

	JMP   enterScheduler

enterScheduler:
	// pop everthing from stack
	ADD   $excCtxSize, R29

	// reenable interrupts
	MOVV  M(C0_SR), R26
	AND   $~SR_EXL, R26
	MOVV  R26, M(C0_SR)

	// reminder: don't use R23, R26 or R27 when interrupts enabled

	// enter scheduler
	MOVB  R0, cpuctx_schedule(g)
	JAL   ·curcpuRunScheduler(SB)

	// disable exceptions again
	MOVV  M(C0_SR), R8
	OR    $SR_EXL, R8
	MOVV  R8, M(C0_SR)

	// clear cpuctx.newexe
	MOVB  R0, (cpuctx_newexe)(g)

	MOVV  (cpuctx_exe)(g), R27
	MOVV  (m_mOS+mOS_ra)(R27), R9
	AND   $~1, R9, R23 // remove smallCtx flag
	AND   $1, R9, R10  // smallCtx flag
	BNE   R0, R10, smallCtx

	MOVV  $(m_mOS+mOS_gprs)(R27), R26
	JAL   ·restoreGPRs(SB)

smallCtx:
	// unmask interrupts
	MOVV  M(C0_SR), R26
	OR    $(INTR_SW|INTR_TIMER), R26
	MOVV  R26, M(C0_SR)

	MOVV  (m_mOS+mOS_sp)(R27), R29
	MOVV  (m_mOS+mOS_fp)(R27), g
	MOVV  R23, R31
	MOVV  (m_mOS+mOS_epc)(R27), R26
	MOVV  R26, M(C0_EPC)

	JMP return
	
// restore ctx of caller
restore:
	MOVV  _LR(R29), R26
	AND   $~1, R26, R31 // remove smallCtx flag from ra
//	MOVV  _mcause(R29), R26
//	MOVV  R26, M(C0_CAUSE)
	MOVV  _mepc(R29), R26
	AND   $1, R26, R27
	AND   $~1, R26  // remove fromHandler flag from epc
	MOVV  R26, M(C0_EPC)

	ADD   $excCtxSize, R29

	// unmask interrupts
	MOVV  M(C0_SR), R26
	OR    $(INTR_SW|INTR_TIMER), R26
	MOVV  R26, M(C0_SR)

	BNE   R27, R0, return

	MOVV  $·cpu0(SB), R26
	MOVV  (g_sched+gobuf_sp)(R26), R29
	MOVV  (g_sched+gobuf_g)(R26), g

return:
	WORD $0x42000018 // ERET

fatal:
	BREAK
	JMP -1(PC)
	
// stolen from runtime/preempt_mips64x.s
// R26 must point to where gprs will be stored
TEXT ·saveGPRs(SB),NOSPLIT|NOFRAME,$0
	MOVV R1, 0(R26)
	MOVV R2, 8(R26)
	MOVV R3, 16(R26)
	MOVV R4, 24(R26)
	MOVV R5, 32(R26)
	MOVV R6, 40(R26)
	MOVV R7, 48(R26)
	MOVV R8, 56(R26)
	MOVV R9, 64(R26)
	MOVV R10, 72(R26)
	MOVV R11, 80(R26)
	MOVV R12, 88(R26)
	MOVV R13, 96(R26)
	MOVV R14, 104(R26)
	MOVV R15, 112(R26)
	MOVV R16, 120(R26)
	MOVV R17, 128(R26)
	MOVV R18, 136(R26)
	MOVV R19, 144(R26)
	MOVV R20, 152(R26)
	MOVV R21, 160(R26)
	MOVV R22, 168(R26)
	MOVV R24, 176(R26)
	MOVV R25, 184(R26)
	MOVV RSB, 192(R26)
	MOVV HI, R1
	MOVV R1, 200(R26)
	MOVV LO, R1
	MOVV R1, 208(R26)
	RET

// stolen from runtime/preempt_mips64x.s
// might clobber some gprs!
// R26 must point to where fprs will be stored
TEXT ·saveFPRs(SB),NOSPLIT|NOFRAME,$0
	#ifndef GOMIPS64_softfloat
	MOVV FCR31, R1
	MOVV R1, 0(R26)
	MOVD F0, 8(R26)
	MOVD F1, 16(R26)
	MOVD F2, 24(R26)
	MOVD F3, 32(R26)
	MOVD F4, 40(R26)
	MOVD F5, 48(R26)
	MOVD F6, 56(R26)
	MOVD F7, 64(R26)
	MOVD F8, 72(R26)
	MOVD F9, 80(R26)
	MOVD F10, 88(R26)
	MOVD F11, 96(R26)
	MOVD F12, 104(R26)
	MOVD F13, 112(R26)
	MOVD F14, 120(R26)
	MOVD F15, 128(R26)
	MOVD F16, 136(R26)
	MOVD F17, 144(R26)
	MOVD F18, 152(R26)
	MOVD F19, 160(R26)
	MOVD F20, 168(R26)
	MOVD F21, 176(R26)
	MOVD F22, 184(R26)
	MOVD F23, 192(R26)
	MOVD F24, 200(R26)
	MOVD F25, 208(R26)
	MOVD F26, 216(R26)
	MOVD F27, 224(R26)
	MOVD F28, 232(R26)
	MOVD F29, 240(R26)
	MOVD F30, 248(R26)
	MOVD F31, 256(R26)
	#endif
	RET


// stolen from runtime/preempt_mips64x.s
// might clobber some gprs!
// R26 must point to stored fprs
TEXT ·restoreFPRs(SB),NOSPLIT|NOFRAME,$0
	#ifndef GOMIPS64_softfloat
	MOVD 256(R26), F31
	MOVD 248(R26), F30
	MOVD 240(R26), F29
	MOVD 232(R26), F28
	MOVD 224(R26), F27
	MOVD 216(R26), F26
	MOVD 208(R26), F25
	MOVD 200(R26), F24
	MOVD 192(R26), F23
	MOVD 184(R26), F22
	MOVD 176(R26), F21
	MOVD 168(R26), F20
	MOVD 160(R26), F19
	MOVD 152(R26), F18
	MOVD 144(R26), F17
	MOVD 136(R26), F16
	MOVD 128(R26), F15
	MOVD 120(R26), F14
	MOVD 112(R26), F13
	MOVD 104(R26), F12
	MOVD 96(R26), F11
	MOVD 88(R26), F10
	MOVD 80(R26), F9
	MOVD 72(R26), F8
	MOVD 64(R26), F7
	MOVD 56(R26), F6
	MOVD 48(R26), F5
	MOVD 40(R26), F4
	MOVD 32(R26), F3
	MOVD 24(R26), F2
	MOVD 16(R26), F1
	MOVD 8(R26), F0
	MOVV 0(R26), R1
	MOVV R1, FCR31
	#endif
	RET

// stolen from runtime/preempt_mips64x.s
// R26 must point to stored gprs
TEXT ·restoreGPRs(SB),NOSPLIT|NOFRAME,$0
	MOVV 208(R26), R1
	MOVV R1, LO
	MOVV 200(R26), R1
	MOVV R1, HI
	MOVV 192(R26), RSB
	MOVV 184(R26), R25
	MOVV 176(R26), R24
	MOVV 168(R26), R22
	MOVV 160(R26), R21
	MOVV 152(R26), R20
	MOVV 144(R26), R19
	MOVV 136(R26), R18
	MOVV 128(R26), R17
	MOVV 120(R26), R16
	MOVV 112(R26), R15
	MOVV 104(R26), R14
	MOVV 96(R26), R13
	MOVV 88(R26), R12
	MOVV 80(R26), R11
	MOVV 72(R26), R10
	MOVV 64(R26), R9
	MOVV 56(R26), R8
	MOVV 48(R26), R7
	MOVV 40(R26), R6
	MOVV 32(R26), R5
	MOVV 24(R26), R4
	MOVV 16(R26), R3
	MOVV 8(R26), R2
	MOVV 0(R26), R1
	RET


// func sysreset(level int, addr unsafe.Pointer) bool
TEXT ·sysreset(SB),NOSPLIT|NOFRAME,$0-12
	NOP // TODO

// func syscachemaint(op int, p unsafe.Pointer, size int)
TEXT ·syscachemaint(SB),NOSPLIT,$0-12
	NOP // TODO

