// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

#define ICSR_ADDR 0xE000ED04
#define ICSR_PENDSVSET (1<<28)

TEXT SysTick_Handler(SB),NOSPLIT|NOFRAME,$0-0
	// set PendSV bit first to avoid DSB but ensure exception tail-chaining
	MOVW  $ICSR_ADDR, R0
	MOVW  $ICSR_PENDSVSET, R1
	MOVW  R1, (R0)

	// increment systim.reloadcnt (64-bit counter)
	MOVW   $·systim(SB), R0
	MOVW   (R0), R1
	ADD.S  $1, R1
	MOVW   R1, (R0)
	RET.CC
	MOVW  4(R0), R1
	ADD   $1, R1
	MOVW  R1, 4(R0)

	// ensure that the change in PendSV pending bit and both halves of
	// streloadcnt are all observable for pendsvHandler and svcallHandler
	DMB

	RET
