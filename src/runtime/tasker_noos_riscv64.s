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
	CSRR  (mhartid, s0)
	MOV   S0, ret+0(FP)
	RET

#define userPriv 0
#define supePriv 1
#define resePriv 2
#define machPriv 3

#define softInt 0
#define timeInt 4
#define exteInt 8

DATA runtime·interruptHandlers+(softInt+userPriv)*8(SB)/8, $·defaultInterruptHandler(SB)
DATA runtime·interruptHandlers+(softInt+supePriv)*8(SB)/8, $·defaultInterruptHandler(SB)
DATA runtime·interruptHandlers+(softInt+resePriv)*8(SB)/8, $·defaultInterruptHandler(SB)
DATA runtime·interruptHandlers+(softInt+machPriv)*8(SB)/8, $·enterScheduler(SB)
DATA runtime·interruptHandlers+(timeInt+userPriv)*8(SB)/8, $·defaultInterruptHandler(SB)
DATA runtime·interruptHandlers+(timeInt+supePriv)*8(SB)/8, $·defaultInterruptHandler(SB)
DATA runtime·interruptHandlers+(timeInt+resePriv)*8(SB)/8, $·defaultInterruptHandler(SB)
DATA runtime·interruptHandlers+(timeInt+machPriv)*8(SB)/8, $·enterScheduler(SB)
DATA runtime·interruptHandlers+(exteInt+userPriv)*8(SB)/8, $·defaultInterruptHandler(SB)
DATA runtime·interruptHandlers+(exteInt+supePriv)*8(SB)/8, $·externalInterruptHandler(SB)
DATA runtime·interruptHandlers+(exteInt+resePriv)*8(SB)/8, $·defaultInterruptHandler(SB)
DATA runtime·interruptHandlers+(exteInt+machPriv)*8(SB)/8, $·externalInterruptHandler(SB)
#define interruptHandlersSize ((1+exteInt+machPriv)*8)
GLOBL runtime·interruptHandlers(SB), RODATA, $interruptHandlersSize

DATA runtime·exceptionHandlers+(0*8)(SB)/8, $·defaultExceptionHandler(SB)
DATA runtime·exceptionHandlers+(1*8)(SB)/8, $·defaultExceptionHandler(SB)
DATA runtime·exceptionHandlers+(2*8)(SB)/8, $·defaultExceptionHandler(SB)
DATA runtime·exceptionHandlers+(3*8)(SB)/8, $·defaultExceptionHandler(SB)
DATA runtime·exceptionHandlers+(4*8)(SB)/8, $·defaultExceptionHandler(SB)
DATA runtime·exceptionHandlers+(5*8)(SB)/8, $·defaultExceptionHandler(SB)
DATA runtime·exceptionHandlers+(6*8)(SB)/8, $·defaultExceptionHandler(SB)
DATA runtime·exceptionHandlers+(7*8)(SB)/8, $·defaultExceptionHandler(SB)
DATA runtime·exceptionHandlers+(8*8)(SB)/8, $·environmentCallHandler(SB)
DATA runtime·exceptionHandlers+(9*8)(SB)/8, $·environmentCallHandler(SB)
DATA runtime·exceptionHandlers+(10*8)(SB)/8, $·environmentCallHandler(SB)
DATA runtime·exceptionHandlers+(11*8)(SB)/8, $·environmentCallHandler(SB)
#define exceptionHandlersSize (12*8)
GLOBL runtime·exceptionHandlers(SB), RODATA, $exceptionHandlersSize


#define sysMaxArgs (48+8)
#define envCallFrameSize (sysMaxArgs+3*8)


// The RISC-V Instruction Set Manual Volume II: Privileged Architecture defines
// the following increasing interrupt priority order:
//
// UTI, USI, UEI, STI, SSI, SEI, MTI, MSI, MEI
//
// That's a hardware priority order conclusive in case of multiple simultaneous
// pending interrupts provided that all are enabled in the mie register.
//
// The trapHandler supports nested interrupts and implements different
// software order using mie register to mask lower priority interrupts:
//
// MTI, MSI - timer and software interrupts have the same, lowest priority,
//            both used to enter or wakeup the scheduler,
//
// SEI, MEI - external interrupts have higher priority than MTI and MSI, they
//            can preempt and wakeup the scheduler, MEI has higher priority
//            than SEI.
//
// We don't support supervisor or user mode interrupts. The platform-specific
// interrupts with id >= 16 (local interrupts) are probably supported (with
// higher priority than MEI) but not tested.
TEXT runtime·trapHandler(SB),NOSPLIT|NOFRAME,$0
	// At this point the interrupts are globaly disabled (mstatus.MIE=0).
	// We want to enable higher priority interrupts as soon as possible.
	// Be carefult to don't clobber T6 (TMP) and A3-A5 (syscall args).

	// mscratch contains &cpuctx if trap from thread mode, 0 for nested trap
	CSRRW  (a0, mscratch, a0)  // swap A0 with cpuctx in mscratch

	// setup g and SP for handler mode, save thread ones to the cpuctx.gh.sched
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

	// mask same or lower priority interrupts (always mask MSI and MTI)
	CSRR  (mcause, lr)
	SRA   $63, LR, A0
	AND   A0, LR  // interrupt: LR=mcause, exception: LR=0
	MOV   $~1, A0
	SLL   LR, A0      // only 6 lower bits of LR are used as shift amount
	AND   $~0xFF, A0  // always mask software and timer interrupts
	CSRR  (mie, lr)
	AND   LR, A0
	CSRW  (a0, mie)

	// enable interrupts
	CSRR   (mcause, a0)  // read mcause before enable interrupts
	CSRSI  ((1<<MIEn), mstatus)

	MOV  LR, _mie(X2)

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


// enterScheduler is caled by timer interrupt or environment call trap
TEXT runtime·enterScheduler(SB),NOSPLIT|NOFRAME,$0

	// if cpuctx.schedule then context saved by environmentCallHandler
	MOVB  (cpuctx_schedule)(g), A0
	BNE   ZERO, A0, contextSaved

	MOV        (cpuctx_exe)(g), A0
	SAVE_GPRS  (A0, m_mOS)  // save most of GPRs

	// save the remaining registers: LR, SP, g, status, mepc
	MOV  _LR(X2), A1
	MOV  _A0(X2), A2
	MOV  _mstatus(X2), S0
	SRL  $7, S0  // MPP field is in a very unfortunate place
	AND  $(3<<(MPPn-7)), S0
	MOV  _mepc(X2), S1
	MOV  (g_sched+gobuf_sp)(g), A3
	MOV  (g_sched+gobuf_g)(g), A4
	MOV  A1, (m_mOS+(const_numGPRS-4)*8)(A0)  // LR
	MOV  A3, (m_mOS+(const_numGPRS-3)*8)(A0)  // SP
	MOV  A4, (m_mOS+(const_numGPRS-2)*8)(A0)  // g
	MOV  A2, (m_mOS+(const_numGPRS-1)*8)(A0)  // A0
	MOV  S0, (m_tls+const_mstatus*8)(A0)
	MOV  S1, (m_tls+const_mepc*8)(A0)

contextSaved:
	MOVB  ZERO, (cpuctx_schedule)(g)

	// clear MSI and MTI
	MOV   $msip, A0
	CSRR  (mhartid, s0)
	SLL   $2, S0
	ADD   S0, A0
	MOVW  ZERO, (A0)
	MOV   $mtimecmp, A0
	SLL   $1, S0
	ADD   S0, A0
	MOV   $-1, S0
	MOV   S0, (A0)

	FENCE  // ensure clearing happens before checking nanotime and futexes

	// enter scheduler
	CALL  ·curcpuRunScheduler(SB)

	MOV  (cpuctx_exe)(g), A0  // load cpuctx.exe

	// check context size
	MOV  (m_tls+const_mstatus*8)(A0), S0
	AND  $const_thrSmallCtx, S0
	BNE  ZERO, S0, smallCtx
	// no need to restore FPRs if exe didn't changed
	MOVB          (cpuctx_newexe)(g), A1
	BEQ           ZERO, A1, 2(PC)
	RESTORE_FPRS  (A0, m_mOS+const_numGPRS*8)  // clobbers TMP
	RESTORE_GPRS  (A0, m_mOS)                  // restore most of GPRs
smallCtx:
	MOVB  ZERO, (cpuctx_newexe)(g)  // clear cpuctx.newexe
	SCW   (zero, zero, x2)          // invalidate possible dangling LR.x instruction by SC to the free word on top of the stack

	// scheduler always returns to the thread mode

	// restore mstatus
	MOV   _mstatus(X2), LR
	SRL   $7, LR
	AND   $~(3<<(MPPn-7)), LR  // clear MPP field
	SLL   $7, LR
	CSRW  (lr, mstatus)  // disables interrupts

	// restore remaining CSRs
	CSRW  (G, mscratch)
	MOV   (m_tls+const_mstatus*8)(A0), g  // load thread status
	AND   $(3<<(MPPn-7)), g
	SLL   $7, g
	CSRS  (G, mstatus)  // set priority field
	MOV   (m_tls+const_mepc*8)(A0), g
	CSRW  (G, mepc)
	MOV   _mie(X2), g
	CSRW  (G, mie)

	// restore remaining GPRs
	MOV  (m_mOS+(const_numGPRS-4)*8)(A0), LR
	MOV  (m_mOS+(const_numGPRS-3)*8)(A0), X2
	MOV  (m_mOS+(const_numGPRS-2)*8)(A0), g
	MOV  (m_mOS+(const_numGPRS-1)*8)(A0), A0

	MRET


TEXT runtime·externalInterruptHandler(SB),NOSPLIT|NOFRAME,$0
	ADD        $-(const_numGPRS-4+const_numFPRS+3)*8, X2
	SAVE_GPRS  (X2, 3*8)
	SAVE_FPRS  (X2, (const_numGPRS-4+3)*8)

	// BUG: the following code assumes two (M,S) PLIC contexts per hart
	// ctxid = mhartid*2 + (11-mcause)/2
	SRL   $3, A0  // mcause
	CSRR  (mhartid, a1)
	SLL   $2, A1
	SUB   A0, A1   // mhartid*4 - mcause
	SLL   $11, A1  // 0x1000*mhartid*2 - 0x1000*mcause/2
	ADD   $(PLIC_TC+11*0x1000/2), A1
	MOV   A1, (X2)
	MOV   $1, S1
	SLL   A0, S1  // 1<<mcause
	MOVW  S1, 12(X2)
	MOVW  (A1), A3    // PLIC.TC[ctxid].THR
	MOVW  A3, 16(X2)  // save current priority threshold

loop:
	MOVW  4(A1), S0  // claim
	BEQ   ZERO, S0, done
	MOVW  S0, 8(X2)

	// allow nested interrupts
	MOV   $PLIC_BASE, A0
	SLL   $2, S0, A2
	ADD   A2, A0
	MOVW  (A0), A2
	MOVW  A2, (A1)  // PLIC.TC[ctxid].THR = PLIC.PRIO[claim]
	FENCE
	CSRS  (s1, mie)

	// get interrupt vector
	MOV  $runtime·vectors(SB), A0
	MOV  (A0), S1
	BGE  S0, S1, noHandler
	SLL  $3, S0
	ADD  S0, A0
	MOV  (A0), A0

	CALL  A0  // call user handler (IRQn_Handler)

	MOV   (X2), A1
	MOV   8(X2), S0
	SRL   $32, S0, S1
	CSRC  (s1, mie)  // disallow nested interrupts
	FENCE
	MOVW  S0, 4(A1)  // complete

	MOVW  16(X2), A3      // saved priority threshold
	BEQ   ZERO, A3, loop  // nested handler can handle only one interrupt

done:
	MOVW  A3, (A1)  // restore priority threshold

	RESTORE_FPRS  (X2, (const_numGPRS-4+3)*8)
	RESTORE_GPRS  (X2, 3*8)
	ADD           $(const_numGPRS-4+const_numFPRS+3)*8, X2

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

noHandler:
	JMP  ·unhandledExternalInterrupt(SB)


// System call is like oridnary function call so all registers are
// caller save (Go ABI0). The tiny wrapper over ECALL instruction add
// additional parameters in A3-A5 registers:
//
// A3: syscall number
// A4: argument data size on the stack (+8 for caller return address)
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
	MOV  A2, (m_mOS+(const_numGPRS-4)*8)(S0)  // LR
	MOV  A0, (m_mOS+(const_numGPRS-3)*8)(S0)  // SP
	MOV  A1, (m_mOS+(const_numGPRS-2)*8)(S0)  // g
	MOV  _mstatus(X2), A1
	SRL  $7, A1  // MPP field is in a very unfortunate place
	AND  $(3<<(MPPn-7)), A1
	OR   $const_thrSmallCtx, A1
	MOV  A1, (m_tls+const_mstatus*8)(S0)
	ADD  $4, S1  // mepc points back to ECALL, adjust it
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
	ADD  $envCallFrameSize, X2

	// run the scheduler if the syscall wants it
	MOVB  cpuctx_schedule(g), S0
	BEQ   ZERO, S0, 2(PC)
	JMP   ·enterScheduler(SB)

	// pop everything from the stack
	MOV  _LR(X2), LR
	MOV  _mstatus(X2), A0
	MOV  _mepc(X2), A1
	ADD  $4, A1  // mepc points back to ECALL, adjust it
	MOV  _mie(X2), A2
	ADD  $trapCtxSize, X2

	// disable interrupts and restore trap context
	CSRW  (a0, mstatus)
	CSRW  (a1, mepc)
	CSRW  (a2, mie)

	AND  $1, A1  // fromThread flag
	BEQ  ZERO, A1, fromHandler

	// restore thread g and SP
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


TEXT runtime·defaultInterruptHandler(SB),NOSPLIT|NOFRAME,$0
	EBREAK
	JMP  -1(PC)


TEXT runtime·defaultExceptionHandler(SB),NOSPLIT|NOFRAME,$0
	ADD  $-16, X2
	SRA  $3, A0
	MOV  A0, (X2)
	ADD  $16+trapCtxSize, X2, A0
	MOV  A0, 8(X2)

	ADD        $-(const_numGPRS-4+1)*8, X2
	SAVE_GPRS  (X2, 8)
	MOV        ZERO, (X2)

	JMP  ·fatalException(SB)


TEXT runtime·unhandledExternalInterrupt(SB),NOSPLIT|NOFRAME,$0
	EBREAK
	JMP  -1(PC)


// curcpuSavectxSched saves floating-point registers to m.mOS
TEXT ·curcpuSavectxSched(SB),NOSPLIT|NOFRAME,$0
	MOV        (cpuctx_exe)(g), A0
	SAVE_FPRS  (A0, m_mOS+const_numGPRS*8)
	RET


//#define WFI CSRR(mie, s1); CSRR(mip, a0); AND S1, A0; BEQ ZERO, A0, -2(PC)


// func curcpuSleep()
TEXT ·curcpuSleep(SB),NOSPLIT|NOFRAME,$0-0

	// We want MSI, MTI to wake up the hart from WFI so they must be enabled in
	// the mie register. Unfortunately, it means that we need to disable
	// interrupts globally which adds a few cycles to the interrupt latency.
	MOV    $(MSI+MTI), S0
	CSRCI  ((1<<MIEn), mstatus)
	CSRS   (s0, mie)
	WFI
	CSRC   (s0, mie)
	CSRSI  ((1<<MIEn), mstatus)

	// clear MSI, MTI
	MOV   $msip, A0
	CSRR  (mhartid, s0)
	SLL   $2, S0
	ADD   S0, A0
	MOVW  ZERO, (A0)
	MOV   $mtimecmp, A0
	SLL   $1, S0
	ADD   S0, A0
	MOV   $-1, S0
	MOV   S0, (A0)

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

// func syscachemaint(op int, p unsafe.Pointer, size int)
TEXT ·syscachemaint(SB),NOSPLIT,$0-24
	//MOVW  op+0(FP), R0
	//MOVW  p+8(FP), R1
	//MOVW  size+16(FP), R2
	EBREAK
	JMP  -1(PC)

// func sysreset(level int, addr unsafe.Pointer) bool
TEXT ·sysreset(SB),NOSPLIT|NOFRAME,$0-24
	EBREAK
	JMP  -1(PC)
