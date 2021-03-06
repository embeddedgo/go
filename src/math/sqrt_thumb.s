// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

#define REGTMP R7

// func Sqrt(x float64) float64
TEXT ·Sqrt(SB),NOSPLIT,$0
	MOVB	runtime·goarm(SB), REGTMP
	CMP	$0x7D, REGTMP
	BNE	softfloat
	MOVD	x+0(FP),F0
	SQRTD	F0,F0
	MOVD	F0,ret+8(FP)
	RET
softfloat:
	// Tail call to Go implementation.
	// Can't use JMP, as in softfloat mode SQRTD is rewritten
	// to a CALL, which makes this function have a frame.
	RET	·sqrt(SB)
