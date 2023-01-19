// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "syscall_noos.h"

// if you add new syscall you must check sysMaxArgs in tasker_noos_riscv64.s

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

GLOBL runtime·syscalls(SB), RODATA, $(SYS_NUM*8)

// func nanotime() int64
TEXT ·nanotime(SB),NOSPLIT|NOFRAME,$0-8
	MOV  $SYS_nanotime, A3
	MOV  $(0+8), A4
	MOV  $8, A5
	ECALL
	RET

// func cputicks() int64
TEXT ·cputicks(SB),NOSPLIT|NOFRAME,$0-8
	MOV  $SYS_nanotime, A3
	MOV  $(0+8), A4
	MOV  $8, A5
	ECALL
	RET

// func irqctl(irq, ctl, ctxid int) (enabled, prio, errno int)
TEXT ·irqctl(SB),NOSPLIT|NOFRAME,$0-48
	MOV  $SYS_irqctl, A3
	MOV  $(24+8), A4
	MOV  $24, A5
	ECALL
	RET

// func setprivlevel(newlevel int) (oldlevel, errno int)
TEXT ·setprivlevel(SB),NOSPLIT|NOFRAME,$0-24
	MOV  $SYS_setprivlevel, A3
	MOV  $(8+8), A4
	MOV  $16, A5
	ECALL
	RET

// func write(fd uintptr, p unsafe.Pointer, n int32) int32
TEXT ·write(SB),NOSPLIT|NOFRAME,$0-32
	MOV  $SYS_write, A3
	MOV  $(24+8), A4
	MOV  $8, A5
	ECALL
	RET

// func cachemaint(op int, p unsafe.Pointer, size int)
TEXT ·cachemaint(SB),NOSPLIT|NOFRAME,$0-24
	MOV  $SYS_cachemaint, A3
	MOV  $(24+8), A4
	MOV  $0, A5
	ECALL
	RET

// func setsystim1()
TEXT ·setsystim1(SB),NOSPLIT|NOFRAME,$0-0
	MOV  $SYS_setsystim1, A3
	MOV  $(0+8), A4
	MOV  $0, A5
	ECALL
	RET

// func setsyswriter1()
TEXT ·setsyswriter1(SB),NOSPLIT|NOFRAME,$0-0
	MOV  $SYS_setsyswriter1, A3
	MOV  $(0+8), A4
	MOV  $0, A5
	ECALL
	RET

// func newosproc(mp *m)
TEXT ·newosproc(SB),NOSPLIT|NOFRAME,$0-8
	MOV  $SYS_newosproc, A3
	MOV  $(8+8), A4
	MOV  $0, A5
	ECALL
	RET

// func exitThread(wait *atomic.Uint32)
TEXT ·exitThread(SB),NOSPLIT|NOFRAME,$0-8
	MOV  $SYS_exitThread, A3
	MOV  $(8+8), A4
	MOV  $0, A5
	ECALL
	RET

// func futexsleep(addr *uint32, val uint32, ns int64)
TEXT ·futexsleep(SB),NOSPLIT|NOFRAME,$0-24
	MOV  $SYS_futexsleep, A3
	MOV  $(24+8), A4
	MOV  $0, A5
	ECALL
	RET

// func futexwakeup(addr *uint32, cnt uint32)
TEXT ·futexwakeup(SB),NOSPLIT|NOFRAME,$0-16
	MOV  $SYS_futexwakeup, A3
	MOV  $(16+8), A4
	MOV  $0, A5
	ECALL
	RET

// func osyield()
TEXT ·osyield(SB),NOSPLIT|NOFRAME,$0-0
	MOV  $SYS_osyield, A3
	MOV  $(0+8), A4
	MOV  $0, A5
	ECALL
	RET

// func nanosleep(ns int64)
TEXT ·nanosleep(SB),NOSPLIT|NOFRAME,$0-8
	MOV  $SYS_nanosleep, A3
	MOV  $(8+8), A4
	MOV  $0, A5
	ECALL
	RET

// unsupported syscalls

// func exit(r int32)
TEXT ·exit(SB),NOSPLIT|NOFRAME,$0-8
	EBREAK
	JMP  -1(PC)
