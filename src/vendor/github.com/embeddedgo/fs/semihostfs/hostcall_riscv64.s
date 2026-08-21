// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// https://github.com/riscv-non-isa/riscv-semihosting/blob/main/riscv-semihosting.adoc

// func hostCall(cmd int, arg uintptr, avoidGC unsafe.Pointer) int
TEXT ·hostCall(SB),NOSPLIT|NOFRAME,$0-32
	MOV   cmd+0(FP), A0
	MOV   arg+8(FP), A1

	// Make sure the following instruction sequence is recognizable
	// by emulators (word align, avoid compressed instructions).
	PCALIGN $4
	WORD    $0x01f01013  // SLLI  $0x1f, ZERO, ZERO
	WORD    $0x00100073  // EBREAK
	WORD    $0x40705013  // SRAI  $0x7, ZERO, ZERO

	MOV   A0, ret+24(FP)
	RET
