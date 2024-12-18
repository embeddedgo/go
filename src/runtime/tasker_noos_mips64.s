// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implementation of the exception handler for the tasker.
//
// There are different levels of context that we save on preemption:
// - goroutine (sp, fp):  Needed to return to the preempted goroutine
//   from the "signal" goroutine, if we didn't schedule.  Saved in g.sched.
// - exception (lr, status, epc):  Registers changed by the exception, must be
//   restored at the end of the exception.  Saved on "signal" stack, i.e. might
//   be nested.
// - thread (sp, fp, lr, epc):  Needed to schedule a goroutine.  If we are
//   going to call the scheduler, the goroutine context is moved into the thread
//   context.  Saved in m.mOS.
// - full (gprs, fprs):  Full context saves all registers.  Needed on arbitrary
//   preemption, i.e. interrupts.  These will be saved on the "signal" stack or
//   in m.mOS.  External interrupts will choose the former, to allow nested
//   interrupts.  Software interrupts will choose the latter, to allow
//   scheduling.
//
// There are some conditions that affect how the exception is handled:
// - fromHandler:  This flag is set when the exception preempted the "signal"
//   goroutine, i.e. we are in a nested exception.  A nested exception must not
//   enter the scheduler, but always return to the caller stack.  Therefore it's
//   also only allowed to call fast syscalls, which are guaranteed to not call
//   the scheduler.  This flag is encoded in bit 0 of the exception context's
//   _mepc value.
// - smallCtx:  This flag indicates that we didn't save the gprs/fprs.  We don't
//   need to save them if preemption happened at function entry, i.e. syscalls.
//   The caller won't care about these because of the Go ABI0.  This flag is
//   encoded in bit 0 of the exception context's _lr value.
// - fast syscall:  Fast syscalls are guaranteed to not enter the scheduler.  We
//   need only to save the goroutine's and exception's context.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "syscall_noos.h"

#include "asm_mips64.h"

#define sysMaxArgs (48+8)

// Exception Context
#define _lr        (0*8)
#define _mstatus   (1*8)
#define _mepc      (2*8)
#define excCtxSize (3*8)


// This will be copied into the processor's general exception vector.  Since the
// general exception vector is limited in size, it just jumps to the actual
// implementation.
TEXT runtime·intvector(SB),NOSPLIT|NOFRAME,$0
	JMP ·exceptionHandler(SB)


// Main exception handler.  If necessary saves goroutine context and switches
// stach to "signal" goroutine.  Saves exceptions context.  Masks pending
// interrupts in preparation for nested interrupts. Lastly evaluates the
// exception cause and calls the appropriate function.
//
// R26 and R27 are free, see runtime/asm_mips64x.s
//
// Interrupt exceptions are disabled by EXL=1 but other exceptions can
// still occur.
//
// Only syscalls and interrupts are handled at the moment, all other exceptions
// are fatal.
TEXT runtime·exceptionHandler(SB),NOSPLIT|NOFRAME,$0
	// Determine caller stack
	MOVV  $·cpu0(SB), R26
	BNE   R26, g, fromThread
	MOVV  $1, R27  // fromHandler flag
	JMP   fromHandler

fromThread:
	MOVV  $0, R27

	// Save goroutine context in g.sched
	MOVV  R29, (g_sched+gobuf_sp)(R26)
	MOVV  g, (g_sched+gobuf_g)(R26)
	MOVV  (g_stack+stack_hi)(R26), R29

	// Switch to "signal" goroutine
	MOVV  R26, g

fromHandler:
	// Save exception context on ISR stack
	SUB   $excCtxSize, R29
	OR    $1, R31, R26 // Encode smallCtx flag in lr
	MOVV  R26, _lr(R29) // R29 is now free for use
	MOVV  M(C0_SR), R26
	MOVV  R26, _mstatus(R29)
	MOVV  M(C0_EPC), R26
	OR    R27, R26 // Encode fromHandler flag in EPC
	MOVV  R26, _mepc(R29)

	// Branch depending on exception cause
	MOVV  M(C0_CAUSE), R26
	AND   $CAUSE_EXC_MASK, R26

	MOVV  $CAUSE_EXC_SYSCALL, R27
	BNE   R26, R27, 2(PC)
	JMP   ·syscallHandler(SB)

	MOVV  $CAUSE_EXC_INTERRUPT, R27
	BNE   R26, R27, fatal

	MOVV  M(C0_CAUSE), R26
	AND   $INTR_MASK, R26

	AND   $INTR_EXT, R26, R27
	BEQ   R27, R0, 2(PC)
	JMP   ·externalInterruptHandler(SB)

	AND   $INTR_SW0, R26, R27
	BEQ   R27, R0, 2(PC)
	JMP   ·softwareInterruptHandler(SB)

	// no interrupt is pending, just return
	JMP  ·exceptionReturn(SB)

fatal:
	JMP   ·unhandledException(SB)


// Execute a system call.  First unsets exception level to allow nested
// interrupts.  Forwards epc by one instruction, otherwise the syscall is
// triggered again on ERET.  Saves thread context if we might schedule, i.e. not
// a fast syscall and not nested exception.  Duffcopies the syscall arguments
// from caller stack to "signal" stack.  Jump to syscall handler.  Duffcopy
// results back.  It the syscall called curcpuSchedule, jump to enterScheduler.
// Otherwise restore exception and goroutine context.
//
// System call is like an oridnary function call, so all registers are caller
// save (Go ABI0).  Meaning we don't need to save GPRs and are free to use them.
// The tiny wrapper over SYSCALL instruction adds additional parameters in
// R8-R10 registers:
//
// R8: syscall number
// R9: argument data size on the stack (+8 for caller return address)
// R10: return data size on the stack
TEXT runtime·syscallHandler(SB),NOSPLIT|NOFRAME,$0
	// Allow nested interrupts.  Reminder: Don't use R26 or R27 when
	// exceptions are enabled.
	MOVV  M(C0_SR), R26
	AND   $~SR_EXL, R26
	MOVV  R26, M(C0_SR)

	MOVV  _mepc(R29), R2
	ADD   $4, R2  // Don't execute syscall instruction again
	MOVV  R2, _mepc(R29)

	MOVV  $SYS_LAST_FAST, R4
	SUB   R4, R8, R4 // R4 <= 0: fast syscall

	// If fromHandler skip saving thread context
	AND   $1, R2  // fromHandler flag
	BNE   R2, R0, currentStack

	MOVV  (g_sched+gobuf_sp)(g), R1 // duffcopy src thread
	BLEZ  R4, duffcopy // fast syscall

	// Save thread context in mOS
	// Needs to be done before the syscall, cpuctx.exe might be nil
	// afterwards (e.g. in nanosleep)
	MOVV  (cpuctx_exe)(g), R2

	MOVV  _lr(R29), R3
	MOVV  R3, (m_mOS+mOS_ra)(R2)

	MOVV  _mepc(R29), R3
	AND   $~1, R3  // Remove fromHandler flag from epc
	MOVV  R3, (m_mOS+mOS_epc)(R2)

	MOVV  (g_sched+gobuf_sp)(g), R3
	MOVV  R3, (m_mOS+mOS_sp)(R2)
	MOVV  (g_sched+gobuf_g)(g), R3
	MOVV  R3, (m_mOS+mOS_fp)(R2)

	JMP   duffcopy

currentStack:
	BGTZ  R4, badSyscall // slow syscall from handler
	ADD   $excCtxSize, R29, R1 // duffcopy src handler

duffcopy:
	// 3 extra registers to preserve src, dst and size of result
	SUB   $sysMaxArgs+3*8, R29

	// Copy arguments from the caller's stack
	MOVV   $·duffcopy<ABIInternal>+2048(SB), R3
	SLL    $1, R9
	SUB    R9, R3
	MOVV   R29, R2 // duffcopy dst
	JAL    (R3)

	// Save data needed to copy the return values back to the caller's stack
	MOVV   R1, (sysMaxArgs+0*8)(R29)
	MOVV   R2, (sysMaxArgs+1*8)(R29)
	MOVV   R10, (sysMaxArgs+2*8)(R29)

	// Call the service routine
	MOVV  $·syscalls(SB), R9
	SLL   $3, R8
	ADD   R8, R9
	MOVV  (R9), R9
	JAL   (R9)

	// Copy the return values back to the caller's stack
	MOVV  (sysMaxArgs+2*8)(R29), R10
	BEQ   R0, R10, nothingToCopy
	MOVV  (sysMaxArgs+0*8)(R29), R2 // duffcopy dst
	MOVV  (sysMaxArgs+1*8)(R29), R1 // duffcopy src
	MOVV  $·duffcopy<ABIInternal>+2048(SB), R3
	SLL   $1, R10
	SUB   R10, R3
	JAL   (R3)

nothingToCopy:
	ADD $sysMaxArgs+3*8, R29

	// Run the scheduler if the syscall wants it
	MOVB  cpuctx_schedule(g), R8  // Always false if nested or fast syscall
	BEQ   R8, R0, 2(PC)
	JMP  ·enterScheduler(SB)

	// Disable nested interrupts
	MOVV  M(C0_SR), R8
	OR    $SR_EXL, R8
	MOVV  R8, M(C0_SR)

	JMP  ·exceptionReturn(SB)

badSyscall:
	BREAK


// An interrupt was caused by software, particularly SW0.  This can happen on
// any instruction.  Saves full context in the thread context, enables nested
// interrupts and jumps to enterScheduler.
TEXT runtime·softwareInterruptHandler(SB),NOSPLIT|NOFRAME,$0
	MOVV  (cpuctx_exe)(g), R26
	MOVV  $(m_mOS+mOS_gprs)(R26), R26
	JAL   ·saveGPRs(SB)

	// Save thread context in mOS
	MOVV  (cpuctx_exe)(g), R27

	MOVV  _lr(R29), R26
	AND   $~1, R26  // Remove smallCtx flag from lr
	MOVV  R26, (m_mOS+mOS_ra)(R27)

	MOVV  _mepc(R29), R26
	AND   $~1, R26  // Remove fromHandler flag from epc
	MOVV  R26, (m_mOS+mOS_epc)(R27)

	MOVV  (g_sched+gobuf_sp)(g), R26
	MOVV  R26, (m_mOS+mOS_sp)(R27)
	MOVV  (g_sched+gobuf_g)(g), R26
	MOVV  R26, (m_mOS+mOS_fp)(R27)

	// Enable nested interrupts
	MOVV  M(C0_SR), R26
	AND   $~SR_EXL, R26
	MOVV  R26, M(C0_SR)

	// Clear pending bit and enter the scheduler.
	MOVV  M(C0_CAUSE), R26
	AND   $~INTR_SW0, R26
	MOVV  R26, M(C0_CAUSE)

	JMP  ·enterScheduler(SB)


// Calls the scheduler and switches to the newly scheduled goroutine.  Assumes
// nested interrupts are enabled.
TEXT runtime·enterScheduler(SB),NOSPLIT|NOFRAME,$0
	// Enter scheduler
	MOVB  R0, cpuctx_schedule(g)
	JAL   ·curcpuRunScheduler(SB)

	// Clear cpuctx.newexe
	MOVB  R0, (cpuctx_newexe)(g)

	// Restore mstatus from exception context
	MOVV  _mstatus(R29), R1
	MOVV  R1, M(C0_SR)
	ADD   $excCtxSize, R29

	// Disable nested interrupts
	MOVV  M(C0_SR), R8
	OR    $SR_EXL, R8
	MOVV  R8, M(C0_SR)

	MOVV  (cpuctx_exe)(g), R27
	MOVV  (m_mOS+mOS_ra)(R27), R9

	AND   $1, R9 // smallCtx flag
	BNE   R0, R9, smallCtx

	MOVV  $(m_mOS+mOS_gprs)(R27), R26
	JAL   ·restoreGPRs(SB)
	// Only use R26, R27 from here

smallCtx:
	MOVV  (m_mOS+mOS_sp)(R27), R29
	MOVV  (m_mOS+mOS_fp)(R27), g
	MOVV  (m_mOS+mOS_ra)(R27), R31
	MOVV  (m_mOS+mOS_epc)(R27), R26
	MOVV  R26, M(C0_EPC)
	MOVV  $~1, R27
	AND   R27, R31 // Remove smallCtx flag

	ERET


// Calls a user implementation for a pending external interrupt.  External
// interrupt can happen anytime, saves full context.  Context will be saved on
// the ISR stack, to allow nested interrupts be supported.  Calculates the
// runtime.vector offset from the interrupt number and calls it.  Restores
// context and returns.
TEXT runtime·externalInterruptHandler(SB),NOSPLIT|NOFRAME,$0
	SUB   $const_numGPRS*8, R29
	MOVV  R29, R26
	JAL   ·saveGPRs(SB)
	SUB   $const_numFPRS*8, R29
	MOVV  R29, R26
	JAL   ·saveFPRs(SB)

	// Mask interrupts
	MOVV  $~INTR_EXT, R10
	MOVV  M(C0_CAUSE), R8 // read cause before enabling nested interrupts
	SYNC
	MOVW  ·highPrioIRQMask(SB), R9
	SYNC
	AND   R8, R9, R12
	BNE   R12, R0, highPrioInterrupt

	// low prio interrrupt: mask all except high priority interrupts
	OR    R9, R10
	MOVV  $~0, R9
highPrioInterrupt: // mask all interrupts
	MOVV  M(C0_SR), R27
	AND   R10, R27

	// Context is saved.  Enable nested interrupts.  Don't use R26, R27.
	AND   $~SR_EXL, R27
	MOVV  R27, M(C0_SR)

	// External interrupts are handled by the application.  We need to call
	// one of the registered handlers.
	AND   $INTR_EXT, R8
	AND   R9, R8 // remove low prio interrupts if high prio is pending

	SRL   $8, R8 // INTR_EXT in lsb
	MOVV  $1, R9

findVector:
	// Find and handle only the first pending interrupt.  If there are
	// multiple interrupts pending, the interrupt handler will run again.
	AND   $1, R8, R10
	BNE   R10, R0, callVector
	ADD   $1, R9
	SRL   $1, R8
	JMP   findVector

callVector:
	SLL   $3, R9, R8 // irq vector offset

	// Get interrupt vector
	MOVV  $runtime·vectors(SB), R10
	MOVV  (R10), R11
	SUB   R9, R11
	BLEZ  R11, fatal
	ADD   R8, R10
	MOVV  (R10), R10

	JAL   (R10)

	// External interrupts can't be cleared by clearing the bits in cause
	// register.  The user handler must have cleared the interrupt at this
	// point in the external source.

	// Disable nested interrupts
	MOVV  M(C0_SR), R8
	OR    $SR_EXL, R8
	MOVV  R8, M(C0_SR)

	// Restore context of caller
	MOVV  R29, R26
	JAL   ·restoreFPRs(SB)
	ADD   $const_numFPRS*8, R29
	MOVV  R29, R26
	JAL   ·restoreGPRs(SB)
	// Only use R26, R27 from here
	ADD   $const_numGPRS*8, R29

	JMP  ·exceptionReturn(SB)

fatal:
	JMP  ·unhandledExternalInterrupt(SB)


TEXT runtime·exceptionReturn(SB),NOSPLIT|NOFRAME,$0
	MOVV  _mstatus(R29), R26
	MOVV  R26, M(C0_SR)
	MOVV  _lr(R29), R26
	MOVV  $~1, R27
	AND   R27, R26, R31 // Remove smallCtx flag from lr
	MOVV  _mepc(R29), R26
	MOVV  $~1, R27
	AND   R26, R27  // Remove fromHandler flag from EPC
	MOVV  R27, M(C0_EPC)

	MOVV  $1, R27
	AND   R26, R27

	ADD   $excCtxSize, R29

	// Don't switch stacks yet if we were called from handler
	BNE   R27, R0, return

	MOVV  $·cpu0(SB), R26
	MOVV  (g_sched+gobuf_sp)(R26), R29
	MOVV  (g_sched+gobuf_g)(R26), g

return:
	ERET


// Do not remove.  Required by the linker, runtime.vectors defaults to this.
TEXT ·unhandledExternalInterrupt(SB),NOSPLIT|NOFRAME,$0
	BREAK


// R26 must point to where gprs will be stored.
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
	MOVV R23, 176(R26)
	MOVV R24, 184(R26)
	MOVV R25, 192(R26)
	MOVV RSB, 200(R26)
	MOVV HI, R1
	MOVV R1, 208(R26)
	MOVV LO, R1
	MOVV R1, 216(R26)
	RET


// R26 must point to stored gprs.  Only use R26, R27 after restoring.  Be
// especially careful and look at the disassembly;  The assembler might decide
// to use R16-R23 for you.
TEXT ·restoreGPRs(SB),NOSPLIT|NOFRAME,$0
	MOVV 216(R26), R1
	MOVV R1, LO
	MOVV 208(R26), R1
	MOVV R1, HI
	MOVV 200(R26), RSB
	MOVV 192(R26), R25
	MOVV 184(R26), R24
	MOVV 176(R26), R23
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


// Might clobber some gprs!
// R26 must point to where fprs will be stored.
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


// Might clobber some gprs!
// R26 must point to stored fprs.
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


// func writeback(addr uintptr, length int)
TEXT ·writeback(SB),NOSPLIT|NOFRAME,$0-16
	MOVV  addr+0(FP), R4
	MOVV  length+8(FP), R5
	ADDU  R5, R4, R8
	AND   $const_cacheLineMask, R4

loop:
	SUB   R4, R8, R9
	BLEZ  R9, done
	CACHE HIT_WRITEBACK_D, 0(R4)
	ADDU  $const_cacheLineSize, R4
	JMP   loop

done:
	RET


// func invalidate(addr uintptr, length int)
TEXT ·invalidate(SB),NOSPLIT|NOFRAME,$0-16
	MOVV  addr+0(FP), R4
	MOVV  length+8(FP), R5
	ADDU  R5, R4, R8
	AND   $const_cacheLineMask, R4

loop:
	SUB   R4, R8, R9
	BLEZ  R9, done
	CACHE HIT_INVALIDATE_D, 0(R4)
	ADDU  $const_cacheLineSize, R4
	JMP   loop

done:
	RET


// func writebackInvalidate(addr uintptr, length int)
TEXT ·writebackInvalidate(SB),NOSPLIT|NOFRAME,$0-16
	MOVV  addr+0(FP), R4
	MOVV  length+8(FP), R5
	ADDU  R5, R4, R8
	AND   $const_cacheLineMask, R4

loop:
	SUB   R4, R8, R9
	BLEZ  R9, done
	CACHE HIT_WRITEBACK_INVALIDATE_D, 0(R4)
	ADDU  $const_cacheLineSize, R4
	JMP   loop

done:
	RET


// syscalls not supported by mips64

// func sysreset(level int, addr unsafe.Pointer) bool
TEXT ·sysreset(SB),NOSPLIT|NOFRAME,$0-12
	NOP
