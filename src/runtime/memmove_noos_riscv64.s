// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// See memmove Go doc for important implementation constraints.

// void runtime·memmove(void*, void*, uintptr)
TEXT runtime·memmove(SB),NOSPLIT,$-0-24
	MOV  to+0(FP), A0
	MOV  from+8(FP), A1
	MOV  n+16(FP), A2
	ADD  A0, A2

	// In case od noos we can't use unaligned load/stores.
	// For now we just copy bytes.
	// TODO: optimize this

loop:
	BEQ    A0, A2, done
	MOVBU  (A1), S0
	ADD    $1, A1
	MOVB   S0, (A0)
	ADD    $1, A0
	JMP    loop

done:
	RET
