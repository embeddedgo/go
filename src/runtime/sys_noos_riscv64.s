// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "syscall_noos.h"

// if you add new syscall you must check SYS_MAX_ARGS in tasker_noos_thumb.s

// syscalls allowed for low priority interrupt handlers
DATA runtime·syscalls+(SYS_nanotime*8)(SB)/8, $·sysnanotime(SB)
DATA runtime·syscalls+(SYS_walltime*8)(SB)/8, $·syswalltime(SB)
DATA runtime·syscalls+(SYS_setwalltime*8)(SB)/8, $·syssetwalltime(SB)
DATA runtime·syscalls+(SYS_irqctl*8)(SB)/8, $·sysirqctl(SB)
DATA runtime·syscalls+(SYS_setprivlevel*8)(SB)/8, $·syssetprivlevel(SB)
DATA runtime·syscalls+(SYS_write*8)(SB)/8, $·syswrite(SB)

// syscalls disallowed for low priority interrupt handlers
DATA runtime·syscalls+(SYS_setsystim1*8)(SB)/8, $·syssetsystim1(SB)
DATA runtime·syscalls+(SYS_newosproc*8)(SB)/8, $·sysnewosproc(SB)
DATA runtime·syscalls+(SYS_exitThread*8)(SB)/8, $·sysexitThread(SB)
DATA runtime·syscalls+(SYS_futexsleep*8)(SB)/8, $·sysfutexsleep(SB)
DATA runtime·syscalls+(SYS_futexwakeup*8)(SB)/8, $·sysfutexwakeup(SB)
DATA runtime·syscalls+(SYS_osyield*8)(SB)/8, $·curcpuSchedule(SB)
DATA runtime·syscalls+(SYS_nanosleep*8)(SB)/8, $·sysnanosleep(SB)

GLOBL runtime·syscalls(SB), RODATA, $(SYS_NUM*8)

// func nanotime1() int64
TEXT ·nanotime1(SB),NOSPLIT|NOFRAME,$0
	MOV  $SYS_nanotime, A0
	MOV  $(0+8), A1
	MOV  $8, A2
	ECALL
	RET

// func walltime1() (sec int64, nsec int32)
TEXT ·walltime1(SB),NOSPLIT|NOFRAME,$0
	MOV  $SYS_walltime, A0
	MOV  $(0+8), A1
	MOV  $16, A2
	ECALL
	RET

// func setwalltime(sec int64, nsec int32)
TEXT ·setwalltime(SB),NOSPLIT|NOFRAME,$0-12
	EBREAK
	RET

// func irqctl(irq, ctl int) (enabled, prio, errno int)
TEXT ·irqctl(SB),NOSPLIT|NOFRAME,$0-20
	EBREAK
	RET

// func setprivlevel(newlevel int) (oldlevel, errno int)
TEXT ·setprivlevel(SB),NOSPLIT|NOFRAME,$0-12
	EBREAK
	RET

// func write1(fd uintptr, p unsafe.Pointer, n int32) int32
TEXT ·write1(SB),NOSPLIT|NOFRAME,$0-16
	EBREAK
	RET

// func setsystim1()
TEXT ·setsystim1(SB),NOSPLIT|NOFRAME,$0-0
	EBREAK
	RET


// func newosproc(mp *m)
TEXT ·newosproc(SB),NOSPLIT|NOFRAME,$0-4
	EBREAK
	RET

// func exitThread(wait *uint32)
TEXT ·exitThread(SB),NOSPLIT|NOFRAME,$0-4
	EBREAK
	RET

// func futexsleep(addr *uint32, val uint32, ns int64)
TEXT ·futexsleep(SB),NOSPLIT|NOFRAME,$0-16
	EBREAK
	RET

// func futexwakeup(addr *uint32, cnt uint32)
TEXT ·futexwakeup(SB),NOSPLIT|NOFRAME,$0-8
	EBREAK
	RET

// func osyield()
TEXT ·osyield(SB),NOSPLIT|NOFRAME,$0-0
	EBREAK
	RET

// func nanosleep(ns int64)
TEXT ·nanosleep(SB),NOSPLIT|NOFRAME,$0-8
	EBREAK
	RET

// unsupported syscalls

// func exit(r int32)
TEXT ·exit(SB),NOSPLIT|NOFRAME,$0-4
	EBREAK
	JMP  -1(PC)
