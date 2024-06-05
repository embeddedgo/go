// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// https://github.com/ARM-software/abi-aa/blob/main/semihosting/semihosting.rst

// func hostCall(cmd int, arg uintptr, avoidGC unsafe.Pointer) int
TEXT ·hostCall(SB),NOSPLIT|NOFRAME,$0-16
	MOVW  cmd+0(FP), R0
	MOVW  arg+4(FP), R1
	BKPT  $0xab
	MOVW  R0, ret+12(FP)
	RET
