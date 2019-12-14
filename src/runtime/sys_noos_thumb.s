// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "syscall_noos.h"

// if you add new syscall you must check SYS_MAX_ARGS in tasker_noos_thumb.s

// non-blocking syscalls
DATA runtime·syscalls+(SYS_nanotime*4)(SB)/4, $·sysnanotime(SB)
DATA runtime·syscalls+(SYS_setsystim1*4)(SB)/4, $·syssetsystim1(SB)
DATA runtime·syscalls+(SYS_walltime*4)(SB)/4, $·syswalltime(SB)
DATA runtime·syscalls+(SYS_setwalltime*4)(SB)/4, $·syssetwalltime(SB)
DATA runtime·syscalls+(SYS_irqctl*4)(SB)/4, $·sysirqctl(SB)
DATA runtime·syscalls+(SYS_setprivlevel*4)(SB)/4, $·syssetprivlevel(SB)
DATA runtime·syscalls+(SYS_write*4)(SB)/4, $·syswrite(SB)

// blocking syscalls
DATA runtime·syscalls+(SYS_newosproc*4)(SB)/4, $·sysnewosproc(SB)
DATA runtime·syscalls+(SYS_exitThread*4)(SB)/4, $·sysexitThread(SB)
DATA runtime·syscalls+(SYS_futexsleep*4)(SB)/4, $·sysfutexsleep(SB)
DATA runtime·syscalls+(SYS_futexwakeup*4)(SB)/4, $·sysfutexwakeup(SB)
DATA runtime·syscalls+(SYS_osyield*4)(SB)/4, $·curcpuSchedule(SB)
DATA runtime·syscalls+(SYS_nanosleep*4)(SB)/4, $·sysnanosleep(SB)

GLOBL runtime·syscalls(SB), RODATA, $(SYS_NUM*4)

// non-blocking syscalls

// func nanotime() int64
TEXT ·nanotime(SB),NOSPLIT|NOFRAME,$0-8
	MOVW  $SYS_nanotime, R4
	MOVW  $(0+4), R5
	MOVW  $8, R6
	SWI
	RET

// func setsystim1()
TEXT ·setsystim1(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  $SYS_setsystim1, R4
	MOVW  $(0+4), R5
	MOVW  $0, R6
	SWI
	RET

// func walltime() (sec int64, nsec int32)
TEXT ·walltime(SB),NOSPLIT|NOFRAME,$0-12
	MOVW  $SYS_walltime, R4
	MOVW  $(0+4), R5
	MOVW  $12, R6
	SWI
	RET

// func setwalltime(sec int64, nsec int32)
TEXT ·setwalltime(SB),NOSPLIT|NOFRAME,$0-12
	MOVW  $SYS_setwalltime, R4
	MOVW  $(12+4), R5
	MOVW  $0, R6
	SWI
	RET

// func irqctl(irq, ctl int) (enabled, prio, errno int)
TEXT ·irqctl(SB),NOSPLIT|NOFRAME,$0-20
	MOVW  $SYS_irqctl, R4
	MOVW  $(8+4), R5
	MOVW  $12, R6
	SWI
	RET

// func setprivlevel(newlevel int) (oldlevel, errno int)
TEXT ·setprivlevel(SB),NOSPLIT|NOFRAME,$0-12
	MOVW  $SYS_setprivlevel, R4
	MOVW  $(4+4), R5
	MOVW  $8, R6
	SWI
	RET

// func write(fd uintptr, p unsafe.Pointer, n int32) int32
TEXT ·write(SB),NOSPLIT|NOFRAME,$0-16
	MOVW  $SYS_write, R4
	MOVW  $(12+4), R5
	MOVW  $4, R6
	SWI
	RET

// blocking syscalls

// func newosproc(mp *m)
TEXT ·newosproc(SB),NOSPLIT|NOFRAME,$0-4
	MOVW  $SYS_newosproc, R4
	MOVW  $(4+4), R5
	MOVW  $0, R6
	SWI
	RET

// func exitThread(wait *uint32)
TEXT ·exitThread(SB),NOSPLIT|NOFRAME,$0-4
	MOVW  $SYS_exitThread, R4
	MOVW  $(4+4), R5
	MOVW  $0, R6
	SWI
	RET

// func futexsleep(addr *uint32, val uint32, ns int64)
TEXT ·futexsleep(SB),NOSPLIT|NOFRAME,$0-16
	MOVW  $SYS_futexsleep, R4
	MOVW  $(16+4), R5
	MOVW  $0, R6
	SWI
	RET

// func futexwakeup(addr *uint32, cnt uint32)
TEXT ·futexwakeup(SB),NOSPLIT|NOFRAME,$0-8
	MOVW  $SYS_futexwakeup, R4
	MOVW  $(8+4), R5
	MOVW  $0, R6
	SWI
	RET

// func osyield()
TEXT ·osyield(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  $SYS_osyield, R4
	MOVW  $(0+4), R5
	MOVW  $0, R6
	SWI
	RET

// func publicationBarrier()
TEXT ·publicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	DMB  MB_ST
	RET

// func nanosleep(ns int64)
TEXT ·nanosleep(SB),NOSPLIT|NOFRAME,$0-8
	MOVW  $SYS_nanosleep, R4
	MOVW  $(8+4), R5
	MOVW  $0, R6
	SWI
	RET

// unsupported syscalls

// func exit(r int32)
TEXT ·exit(SB),NOSPLIT|NOFRAME,$0-4
	BKPT
	B   -1(PC)
