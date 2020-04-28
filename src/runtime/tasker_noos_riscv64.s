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


TEXT runtime·trapHandler(SB),NOSPLIT|NOFRAME,$0
	// At this point the interrupts are globaly disabled (mstatus.MIE=0).
	// We want to enable higher priority interrupts as soon as possible.
	// Be carefult to don't clober REGTMP (T6/X31).

	// The MSCRATCH contains *cpuctx if trap from thread mode or 0 in case of
	// nested trap.

	CSRRW  (a0, MSCRATCH, a0)  // swap A0 with cpuctx in MSCRATCH
	BEQ    ZERO, A0, nestedEntry

	// setup g and SP for handler mode
	MOV  X2, (g_sched+gobuf_sp)(A0)
	MOV  g, (g_sched+gobuf_g)(A0)
	MOV  (g_stack+stack_hi)(X9), A0
	MOV  A0, g

nestedEntry:
	// Now we have handler SP and g (cpuctr) in X2 and X4 regardless or nested
	// or non-nested trap. The thread g and SP are in cpuctx.gh.sched. There is
	// one free register (A0) for further processing with its previous content
	// in MSCRATCH (must set MSCRATCH=0 before enable exceptions).

	CSRR  (MCAUSE, a0)


	CSRRWI  (0, MSCRATCH, a0)  // sestore A0 content and set MSCRATCH=0

loop:
	JMP  loop
