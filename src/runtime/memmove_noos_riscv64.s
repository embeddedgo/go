// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// See memmove Go doc for important implementation constraints.

// void runtime·memmove(void*, void*, uintptr)
TEXT runtime·memmove<ABIInternal>(SB),NOSPLIT,$-0-24
	// A0 = to
	// A1 = from
	// A2 = n

	BEQ  A2, ZERO, done

	// check alignment
	// The code below requires all of the "to", "from", "n" are aligned
	// to use multi-byte load/stores. TODO: relax this requirement.
	OR   A0, A1, S0
	OR   A2, S0
	AND  $7, S0
	AND  $3, S0, S1
	AND  $1, S0, S2

	// check the relative position of the "to" and "from"
	BLTU  A1, A0, backward

	ADD  A0, A2
	BEQ  S0, ZERO, forwardLoop8
	BEQ  S1, ZERO, forwardLoop4
	BEQ  S2, ZERO, forwardLoop2

forwardLoop1:
	MOVBU  (A1), S0
	ADD    $1, A1
	MOVB   S0, (A0)
	ADD    $1, A0
	BNE    A0, A2, forwardLoop1
done:
	RET

forwardLoop8:
	MOV  (A1), S0
	ADD  $8, A1
	MOV  S0, (A0)
	ADD  $8, A0
	BNE  A0, A2, forwardLoop8
	RET

forwardLoop4:
	MOVWU  (A1), S0
	ADD    $4, A1
	MOVW   S0, (A0)
	ADD    $4, A0
	BNE    A0, A2, forwardLoop4
	RET

forwardLoop2:
	MOVHU  (A1), S0
	ADD    $2, A1
	MOVH   S0, (A0)
	ADD    $2, A0
	BNE    A0, A2, forwardLoop2
	RET

backward:
	ADD  A2, A1
	ADD  A0, A2
	BEQ  S0, ZERO, backwardLoop8
	BEQ  S1, ZERO, backwardLoop4
	BEQ  S2, ZERO, backwardLoop2

backwardLoop1:
	ADD    $-1, A1
	MOVBU  (A1), S0
	ADD    $-1, A2
	MOVB   S0, (A2)
	BNE    A0, A2, backwardLoop1
	RET

backwardLoop8:
	ADD  $-8, A1
	MOV  (A1), S0
	ADD  $-8, A2
	MOV  S0, (A2)
	BNE  A0, A2, backwardLoop8
	RET

backwardLoop4:
	ADD    $-4, A1
	MOVWU  (A1), S0
	ADD    $-4, A2
	MOVW   S0, (A2)
	BNE    A0, A2, backwardLoop4
	RET

backwardLoop2:
	ADD    $-2, A1
	MOVHU  (A1), S0
	ADD    $-2, A2
	MOVH   S0, (A2)
	BNE    A0, A2, backwardLoop2
	RET
