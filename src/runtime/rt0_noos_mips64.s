// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"


TEXT _rt0_mips64_noos(SB),NOSPLIT|NOFRAME,$0
	// TODO implement
	RET


#define PALLOC_MIN 20*1024

#define SCB_BASE  0xE000ED00
#define SCB_VTOR  0x008
#define SCB_CPACR 0x088
#define SCB_FPCCR 0x234

TEXT runtime·rt0_go(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	// TODO implement
	RET
