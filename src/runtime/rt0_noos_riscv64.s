// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"


// There is prefered to use X8-X15 registers (S0, S1, A0-A5) to promote
// short instruction encoding if C extension will be supported.


#define PALLOC_MIN 20*1024


TEXT _rt0_riscv64_noos(SB),NOSPLIT|NOFRAME,$0

	// initialize data and BSS

	MOV  $runtime·ramstart(SB), A0
	MOV  $runtime·romdata(SB), A1
	MOV  $runtime·bss(SB), A3
	MOV  $runtime·ramend(SB), A4

	JMP  runtime·rt0_go(SB)  // rt0_go is known as top of a goroutine stack


TEXT runtime·rt0_go(SB),NOSPLIT|NOFRAME,$0

	JMP  main·main(SB)
