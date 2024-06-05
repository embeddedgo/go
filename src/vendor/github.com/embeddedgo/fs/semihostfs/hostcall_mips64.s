// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// https://lore.kernel.org/lkml/202311010521.7YJZiVJm-lkp@intel.com/T/

#define SDBBP(code) WORD $(0x7000003f + code<<6)

#define SDBBP_R6(code) WORD $(0x0000000e + code<<6)

// func hostCall(cmd int, a0, a1, a2 uintptr, avoidGC *byte) int
TEXT ·hostCall(SB),NOSPLIT|NOFRAME,$0-48
	MOVV  cmd+0(FP), R25
	MOVV  a0+8(FP), R4
	MOVV  a1+16(FP), R5
	MOVV  a2+24(FP), R6
	SDBBP(1)
	MOVV  $-1, R1
	BNE   R1, R2, 2(PC)
	SUB   R3, R0, R2
	MOVV  R2, ret+40(FP)
	RET
