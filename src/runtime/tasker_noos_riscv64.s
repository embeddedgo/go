// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "syscall_noos.h"


TEXT runtime·interruptHandler(SB),NOSPLIT|NOFRAME,$0-0
loop:
	EBREAK
	JMP  loop
