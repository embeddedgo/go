// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"


TEXT ·load64(SB),NOSPLIT,$0-16
	MOVV  addr+0(FP), R4
	MOVV  (R4), R5
	MOVV  R5, ret+8(FP)
	RET

TEXT ·load32(SB),NOSPLIT,$0-12
	MOVV   addr+0(FP), R4
	MOVWU  (R4), R5
	MOVW   R5, ret+8(FP)
	RET

TEXT ·load16(SB),NOSPLIT,$0-10
	MOVV   addr+0(FP), R4
	MOVHU  (R4), R5
	MOVH   R5, ret+8(FP)
	RET

TEXT ·load8(SB),NOSPLIT,$0-9
	MOVV   addr+0(FP), R4
	MOVBU  (R4), R5
	MOVB   R5, ret+8(FP)
	RET

TEXT ·store64(SB),NOSPLIT,$0-16
	MOVV  addr+0(FP), R4
	MOVV  v+8(FP), R5
	MOVV  R5, (R4)
	RET

TEXT ·store32(SB),NOSPLIT,$0-12
	MOVV   addr+0(FP), R4
	MOVWU  v+8(FP), R5
	MOVW   R5, (R4)
	RET

TEXT ·store16(SB),NOSPLIT,$0-10
	MOVV   addr+0(FP), R4
	MOVHU  v+8(FP), R5
	MOVH   R5, (R4)
	RET

TEXT ·store8(SB),NOSPLIT,$0-9
	MOVV   addr+0(FP), R4
	MOVBU  v+8(FP), R5
	MOVB   R5, (R4)
	RET

TEXT ·MB(SB),NOSPLIT,$0
	SYNC
	RET
