// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "syscall_noos.h"

// if you add new syscall you must check sysMaxArgs in tasker_noos_*.s

// syscalls allowed for low priority interrupt handlers
DATA runtime·syscalls+(SYS_nanotime*8)(SB)/8, $·sysnanotime(SB)
DATA runtime·syscalls+(SYS_irqctl*8)(SB)/8, $·sysirqctl(SB)
DATA runtime·syscalls+(SYS_setprivlevel*8)(SB)/8, $·syssetprivlevel(SB)
DATA runtime·syscalls+(SYS_write*8)(SB)/8, $·syswrite(SB)
DATA runtime·syscalls+(SYS_cachemaint*8)(SB)/8, $·syscachemaint(SB)

// syscalls disallowed for low priority interrupt handlers
DATA runtime·syscalls+(SYS_setsystim1*8)(SB)/8, $·syssetsystim1(SB)
DATA runtime·syscalls+(SYS_setsyswriter1*8)(SB)/8, $·syssetsyswriter1(SB)
DATA runtime·syscalls+(SYS_newosproc*8)(SB)/8, $·sysnewosproc(SB)
DATA runtime·syscalls+(SYS_exitThread*8)(SB)/8, $·sysexitThread(SB)
DATA runtime·syscalls+(SYS_futexsleep*8)(SB)/8, $·sysfutexsleep(SB)
DATA runtime·syscalls+(SYS_futexwakeup*8)(SB)/8, $·sysfutexwakeup(SB)
DATA runtime·syscalls+(SYS_osyield*8)(SB)/8, $·curcpuSchedule(SB)
DATA runtime·syscalls+(SYS_nanosleep*8)(SB)/8, $·sysnanosleep(SB)
DATA runtime·syscalls+(SYS_reset*8)(SB)/8, $·sysreset(SB)

GLOBL runtime·syscalls(SB), RODATA, $(SYS_NUM*8)

// func nanotime() int64
TEXT ·nanotime(SB),NOSPLIT|NOFRAME,$0-8
	MOVV  $SYS_nanotime, R8
	MOVV  $(0+8), R9
	MOVV  $8, R10
	SYSCALL
	RET

// func irqctl(irq, ctl, ctxid int) (enabled, prio, errno int)
TEXT ·irqctl(SB),NOSPLIT|NOFRAME,$0-48
	MOVV  $SYS_irqctl, R8
	MOVV  $(24+8), R9
	MOVV  $24, R10
	SYSCALL
	RET

// func setprivlevel(newlevel int) (oldlevel, errno int)
TEXT ·setprivlevel(SB),NOSPLIT|NOFRAME,$0-24
	MOVV  $SYS_setprivlevel, R8
	MOVV  $(8+8), R9
	MOVV  $16, R10
	SYSCALL
	RET

// func write(fd uintptr, p unsafe.Pointer, n int32) int32
TEXT ·write(SB),NOSPLIT|NOFRAME,$0-32
	MOVV  $SYS_write, R8
	MOVV  $(24+8), R9
	MOVV  $8, R10
	SYSCALL
	RET

// func cachemaint(op int, p unsafe.Pointer, size int)
TEXT ·cachemaint(SB),NOSPLIT|NOFRAME,$0-24
	MOVV  $SYS_cachemaint, R8
	MOVV  $(24+8), R9
	MOVV  $0, R10
	SYSCALL
	RET

// func setsystim1()
TEXT ·setsystim1(SB),NOSPLIT|NOFRAME,$0-0
	MOVV  $SYS_setsystim1, R8
	MOVV  $(0+8), R9
	MOVV  $0, R10
	SYSCALL
	RET


// func setsyswriter1()
TEXT ·setsyswriter1(SB),NOSPLIT|NOFRAME,$0-0
	MOVV  $SYS_setsyswriter1, R8
	MOVV  $(0+8), R9
	MOVV  $0, R10
	SYSCALL
	RET


// func newosproc(mp *m)
TEXT ·newosproc(SB),NOSPLIT|NOFRAME,$0-8
	MOVV  $SYS_newosproc, R8
	MOVV  $(8+8), R9
	MOVV  $0, R10
	SYSCALL
	RET

// func exitThread(wait *atomic.Uint32)
TEXT ·exitThread(SB),NOSPLIT|NOFRAME,$0-8
	MOVV  $SYS_exitThread, R8
	MOVV  $(8+8), R9
	MOVV  $0, R10
	SYSCALL
	RET

// func futexsleep(addr *uint32, val uint32, ns int64)
TEXT ·futexsleep(SB),NOSPLIT|NOFRAME,$0-24
	MOVV  $SYS_futexsleep, R8
	MOVV  $(24+8), R9
	MOVV  $0, R10
	SYSCALL
	RET

// func futexwakeup(addr *uint32, cnt uint32)
TEXT ·futexwakeup(SB),NOSPLIT|NOFRAME,$0-16
	MOVV  $SYS_futexwakeup, R8
	MOVV  $(16+8), R9
	MOVV  $0, R10
	SYSCALL
	RET

// func osyield()
TEXT ·osyield(SB),NOSPLIT|NOFRAME,$0-0
	MOVV  $SYS_osyield, R8
	MOVV  $(0+8), R9
	MOVV  $0, R10
	SYSCALL
	RET

// func nanosleep(ns int64)
TEXT ·nanosleep(SB),NOSPLIT|NOFRAME,$0-8
	MOVV  $SYS_nanosleep, R8
	MOVV  $(8+8), R9
	MOVV  $0, R10
	SYSCALL
	RET

// func reset(level int, addr unsafe.Pointer) bool
TEXT ·reset(SB),NOSPLIT|NOFRAME,$0-16
	MOVV  $SYS_reset, R8
	MOVV  $(16+8), R9
	MOVV  $0, R10
	SYSCALL
	RET

// func exit(r int32)
TEXT ·exit(SB),NOSPLIT|NOFRAME,$0-8
	NOOP
	JMP  -1(PC)
