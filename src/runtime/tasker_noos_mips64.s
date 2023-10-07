// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "syscall_noos.h"

#include "asm_mips64.h"

TEXT runtime·intvector(SB),NOSPLIT|NOFRAME,$0
	JMP ·inthandler(SB)

TEXT runtime·inthandler(SB),NOSPLIT|NOFRAME,$0
	RET // TODO implement

// func sysreset(level int, addr unsafe.Pointer) bool
TEXT ·sysreset(SB),NOSPLIT|NOFRAME,$0-12
	NOP // TODO

// func syscachemaint(op int, p unsafe.Pointer, size int)
TEXT ·syscachemaint(SB),NOSPLIT,$0-12
	NOP // TODO
