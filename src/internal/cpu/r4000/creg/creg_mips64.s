// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·mtc0(SB),NOSPLIT,$0-12
	MOVV   creg+0(FP), R4
	MOVWU  val+8(FP), R5
	SLL    $4, R4
	MOVV   $_mtc0(SB), R6
	ADDU   R6, R4
	JMP    (R4)


TEXT _mtc0(SB),NOSPLIT|NOFRAME,$0
	MOVW   R5, M(0)
	RET
	WORD   $0x0
	MOVW   R5, M(1)
	RET
	WORD   $0x0
	MOVW   R5, M(2)
	RET
	WORD   $0x0
	MOVW   R5, M(3)
	RET
	WORD   $0x0
	MOVW   R5, M(4)
	RET
	WORD   $0x0
	MOVW   R5, M(5)
	RET
	WORD   $0x0
	MOVW   R5, M(6)
	RET
	WORD   $0x0
	MOVW   R5, M(7)
	RET
	WORD   $0x0
	MOVW   R5, M(8)
	RET
	WORD   $0x0
	MOVW   R5, M(9)
	RET
	WORD   $0x0
	MOVW   R5, M(10)
	RET
	WORD   $0x0
	MOVW   R5, M(11)
	RET
	WORD   $0x0
	MOVW   R5, M(12)
	RET
	WORD   $0x0
	MOVW   R5, M(13)
	RET
	WORD   $0x0
	MOVW   R5, M(14)
	RET
	WORD   $0x0
	MOVW   R5, M(15)
	RET
	WORD   $0x0
	MOVW   R5, M(16)
	RET
	WORD   $0x0
	MOVW   R5, M(17)
	RET
	WORD   $0x0
	MOVW   R5, M(18)
	RET
	WORD   $0x0
	MOVW   R5, M(19)
	RET
	WORD   $0x0
	MOVW   R5, M(20)
	RET
	WORD   $0x0
	MOVW   R5, M(21)
	RET
	WORD   $0x0
	MOVW   R5, M(22)
	RET
	WORD   $0x0
	MOVW   R5, M(23)
	RET
	WORD   $0x0
	MOVW   R5, M(24)
	RET
	WORD   $0x0
	MOVW   R5, M(25)
	RET
	WORD   $0x0
	MOVW   R5, M(26)
	RET
	WORD   $0x0
	MOVW   R5, M(27)
	RET
	WORD   $0x0
	MOVW   R5, M(28)
	RET
	WORD   $0x0
	MOVW   R5, M(29)
	RET
	WORD   $0x0
	MOVW   R5, M(30)
	RET
	WORD   $0x0
	MOVW   R5, M(31)
	RET
	WORD   $0x0


TEXT ·mfc0(SB),NOSPLIT,$0-12
	MOVV   creg+0(FP), R4
	SLL    $4, R4
	MOVV   $_mfc0(SB), R6
	ADDU   R6, R4
	JMP    (R4)


TEXT _mfc0(SB),NOSPLIT|NOFRAME,$0
	MOVW   M(0), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(1), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(2), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(3), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(4), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(5), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(6), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(7), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(8), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(9), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(10), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(11), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(12), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(13), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(14), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(15), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(16), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(17), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(18), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(19), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(20), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(21), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(22), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(23), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(24), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(25), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(26), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(27), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(28), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(29), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(30), R5
	MOVW   R5, ret+8(FP)
	RET
	MOVW   M(31), R5
	MOVW   R5, ret+8(FP)
	RET
