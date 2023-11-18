// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"
#include "asm_mips64.h"

#define DEFAULT_C0_SR SR_CU1|SR_PE|SR_FR

#define TLBWI WORD $0x42000002

TEXT _rt0_mips64_noos(SB),NOSPLIT|NOFRAME,$0
	// Watchpoints have been proven to persist across resets and even with
	// the console being off. Zero it as early as possible, to avoid it
	// triggering during boot. This should really be done at the start IPL3.
	MOVW z0, M(C0_WATCHLO)

	JAL  runtime·rt0_tlb(SB)

	// Check whether we are running on iQue or N64. Use the MI version
	// register which has LSB set to 0xB0 on iQue. We assume 0xBn was meant
	// for BBPlayer. Notice that we want this test to be hard for emulators
	// to pass by mistake, so checking for a specific value while reading
	// seems solid enough.
	MOVW (0x80000318), t0 // memory size
	MOVW (0xA4300004), t1
	MOVW $0xB0, t2
	AND  $0xF0, t1
	MOVW $0, t7 // t7=0 -> vanilla N64
	BNE  t1, t2, set_sp

	// In iQue player, memory allocated to game can be configured and it
	// appears in 0x80000318. On the other hand, the top 8 MiB of RDRAM is
	// reserved to savegames. So avoid putting the stack there, capping the
	// size to 0x7C0000. See also get_memory_size.
	MOVW $1, t7 // t7=1 -> iQue player
	MOVW 0x800000, t1
	SGT  t2, t1, t0
	MOVW 0x0, t1
	BNE  t2, t1, set_sp
	MOVW 0x7C0000, t0

set_sp:
	MOVV $0x10, t1
	SUBV t1, t0, sp // init stack pointer
	MOVV $0, RSB // init data pointer
	MOVW $8, v0
	MOVW v0, (0xbfc007fc) // magic N64 hardware init

	// a bit from libgloss so we start at a known state
	MOVW $DEFAULT_C0_SR, v0
	MOVW v0, M(C0_SR)
	MOVW $0, M(C0_CAUSE)

	// Check if PI DMA transfer is required, knowing that IPL3 loads 1 MiB
	// of ROM to RAM, and __libdragon_text_start is located right after the
	// ROM header where this 1 MiB starts.
	// TODO test again with TLB mappings configured
	MOVW $_rt0_mips64_noos(SB), a0
	MOVW $runtime·edata(SB), a1
	MOVW $0x100000, t0 // stock IPL3 load size (1 MiB)
	SUBU a0, a1, a2	// calculate data size
	SUB  t0, a2, a2 // calculate remaining data size
	BLEZ a2, skip_dma // skip PI DMA if data is already loaded

	// Copy code and data via DMA
	MOVW $0x10001000, a1 // address in rom
	ADDU t0, a0, a0	// skip over loaded data
	ADDU t0, a1, a1				

	// Start PI DMA transfer
	MOVW $0xA4600000, t0
	MOVW a0, 0x00(t0) // PI_DRAM_ADDR
	MOVW a1, 0x04(t0) // PI_CART_ADDR
	ADD  $-1, a2
	MOVW a2, 0x0C(t0) // PI_WR_LEN

skip_dma:
	// fill .bss with 0s
	SUBU $16, sp
	MOVW $runtime·bss(SB), a0
	OR   $0x20000000, a0 // convert address to KSEG1 (uncached)
	MOVW $runtime·ebss(SB), a1
	OR   $0x20000000, a1
	SUB  a0, a1
	MOVV a0, 8(sp)
	MOVV a1, 16(sp)
	JAL  runtime·memclrNoHeapPointers(SB)  // clear BSS

	// fill .noptrbss with 0s
	MOVW $runtime·noptrbss(SB), a0
	OR   $0x20000000, a0 // convert address to KSEG1 (uncached)
	MOVW $runtime·enoptrbss(SB), a1
	OR   $0x20000000, a1
	SUB  a0, a1
	MOVV a0, 8(sp)
	MOVV a1, 16(sp)
	JAL  runtime·memclrNoHeapPointers(SB)  // clear noptrBSS

	// clear unallocated memory
	MOVW $runtime·end(SB), a0
	OR   $0x20000000, a0 // convert address to KSEG1 (uncached)
	MOVW $runtime·ramend(SB), a1
	OR   $0x20000000, a1
	SUB  a0, a1
	MOVV a0, 8(sp)
	MOVV a1, 16(sp)
	JAL  runtime·memclrNoHeapPointers(SB)  // clear unallocated memory
	ADDU $16, sp

	// Wait for DMA transfer to be finished
	MOVW $0xA4600000, t0

wait_dma_end:
	MOVW 0x10(t0), t1 // PI_STATUS
	AND  $3, t1 // PI_STATUS_DMA_BUSY | PI_STATUS_IO_BUSY
	BGTZ t1, wait_dma_end

	// Store the bbplayer flag now that BSS has been cleared
	MOVW t7, runtime·bbplayer(SB)

	// load interrupt vector
	MOVW $·intvector(SB), t0
	MOVW $0xa0000000, t1
	MOVW $4, t2

loadintvectorloop:
	MOVW (t0), t3
	MOVW t3, 0(t1)
	MOVW t3, 0x80(t1)
	MOVW t3, 0x100(t1)
	MOVW t3, 0x180(t1)
	// sync - BREAK is overloaded CACHE opcode
	BREAK R16, 0(t1)
	BREAK R16, 0x80(t1)
	BREAK R16, 0x100(t1)
	BREAK R16, 0x180(t1)
	ADD $4, t0
	ADD $4, t1
	ADDU $-1, t2
	BGTZ t2,loadintvectorloop

	JAL ·initEverdrive64USB(SB) // get ready for logging

	JMP runtime·rt0_go(SB)


// Maps KSEG0, KSEG1 and ROM to the beginning of the virtual address space.
// This saves us from sign-extending pointers correctly, as we avoid pointers
// upwards of 0x80000000.
// The only problems encountered with falsely sign-extended pointers were symbol
// addresses loaded from the ELF.  Otherwise running code in KSEG0 seems to be
// no problem in general.
// The correct way of doing this would probably involve defining a new 64-bit
// architecture with PtrSize = 4, but I have ran into some issues that weren't
// worth fixing at the moment.
// TODO currently only 16 MiB of cartridge is mapped
TEXT runtime·rt0_tlb(SB),NOSPLIT|NOFRAME,$0
	MOVV $0, R8
	MOVV R8, M(C0_INDEX)
	MOVV $0xfff << 13, R8
	MOVV R8, M(C0_PAGEMASK)
	MOVV $(0x00000000 >> 6) | 0x7, R8
	MOVV R8, M(C0_ENTRYLO0)
	MOVV $(0x01000000 >> 6) | 0x7, R8
	MOVV R8, M(C0_ENTRYLO1)
	MOVV $0x00000000, R8
	MOVV R8, M(C0_ENTRYHI)
	NOOP // avert CP0 hazards
	TLBWI

	MOVV $1, R8
	MOVV R8, M(C0_INDEX)
	MOVV $0xfff << 13, R8
	MOVV R8, M(C0_PAGEMASK)
	MOVV $(0x10000000 >> 6) | (2<<3) |  0x3, R8
	MOVV R8, M(C0_ENTRYLO0)
	MOVV $(0x11000000 >> 6) | (2<<3) |  0x3, R8
	MOVV R8, M(C0_ENTRYLO1)
	MOVV $0x10000000, R8
	MOVV R8, M(C0_ENTRYHI)
	NOOP // avert CP0 hazards
	TLBWI

	MOVV $2, R8
	MOVV R8, M(C0_INDEX)
	MOVV $0xfff << 13, R8
	MOVV R8, M(C0_PAGEMASK)
	MOVV $(0x00000000 >> 6) | (2<<3) | 0x7, R8
	MOVV R8, M(C0_ENTRYLO0)
	MOVV $(0x01000000 >> 6) | (2<<3) | 0x7, R8
	MOVV R8, M(C0_ENTRYLO1)
	MOVV $0x20000000, R8
	MOVV R8, M(C0_ENTRYHI)
	NOOP // avert CP0 hazards
	TLBWI

	MOVV $0x7fffffff, R8
	AND  R8, ra // return to the tlb mapped address
	RET


#define PALLOC_MIN 20*1024

TEXT runtime·rt0_go(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	// setup main stack in cpu0.gh
	MOVV  $runtime·cpu0(SB), t0  // gh is the first field of the cpuctx struct
	MOVV  $runtime·ramend(SB), sp  // main stack starts at the end of memory
	SUB   $16, sp
	MOVV  sp, (g_stack+stack_hi)(t0)
	SUB   $4096, sp, t1
	MOVV  t1, (g_stack+stack_lo)(t0)
	ADD   $const_stackGuard, t1
	MOVV  t1, g_stackguard0(t0)
	MOVV  t1, g_stackguard1(t0)

	// set up m0 (bootstrap thread), temporarily use cpu0.gh as g
	MOVV  $runtime·m0(SB), t1
	MOVV  t0, m_g0(t1)  // m0.g0 = cpu0.gh
	MOVV  t1, g_m(t0)   // cpu0.gh.m = m0

	MOVV  t0, g  // set g to gh

	JAL runtime·check(SB)
	JAL runtime·osinit(SB)

	// initialize noosMem

	MOVV  $runtime·end(SB), t0
	MOVV  $runtime·ramend(SB), t1
	SUB   $4096, t1
	SUB   t0, t1, t5  // size of available memory (DMA capable)

	// estimate the space need for non-heap allocations
	SRL   $(const__PageShift+1), t5, t4
	MOVV  $mspan__size, t2
	MUL   t2, t4
	MOVV  LO, t4
	ADD   $PALLOC_MIN, t4

	MOVV  $runtime·nodmastart(SB), t2
	MOVV  $runtime·nodmaend(SB), t3
	SUB   t2, t3, t7  // size of non-DMA memory
	ADD   t5, t7, t6  // size of the whole free memory

	// we prefer the non-DMA memory for non-heap objects to preserve as much as
	// possible of the DMA capable memory for heap allocations
	SUB   t7, t4

	// reduce the arena by the remain of the non-heap space that did not fit in
	// the non-DMA memory, properly align the arena
	BLTZ  t4, 2(PC)
	SUB   t4, t5
	AND   $~(const_heapArenaBytes-1), t5
	SUB   t5, t1
	MOVV  t1, t4

	// save {free.start,free.end,nodma.start,nodma.end,arenaStart,arenaSize,size}
	MOVV  $runtime·noosMem(SB), t7
	MOVV  t0, 0(t7)
	MOVV  t1, 8(t7)
	MOVV  t2, 16(t7)
	MOVV  t3, 24(t7)
	MOVV  t4, 32(t7)
	MOVV  t5, 40(t7)
	MOVV  t6, 48(t7)

	// initialize noos tasker and Go scheduler
	JAL  runtime·taskerinit(SB)
	JAL  runtime·schedinit(SB)

	// allocate g0 for m0 and leave gh
	SUB   $16, sp
	MOVW  $(2*const_stackMin), a0
	MOVW  a0, 8(sp)
	JAL   runtime·malg(SB)
	MOVV  16(sp), t0  // newg in t0
	ADD   $16, sp

	// stackguard check during newproc requires valid stackguard1 but
	// malg sets it to 0xFFFFFFFF (mstart fixes this but is called later)
	MOVV  g_stackguard0(t0), t1
	MOVV  t1, g_stackguard1(t0)

	MOVV  $runtime·m0(SB), t1
	MOVV  t0, m_g0(t1)  // m0.g0 = newg
	MOVV  t1, g_m(t0)   // newg.m = m0

	// newg stack pointer to sp
	MOVV  (g_stack+stack_hi)(t0), sp

	// newg to g
	MOVV  g, t2
	MOVV  t0, g

	// fix cpu0.gh, cpu0.mh
	ADD   $cpuctx_mh, t2, t1  // t2 points to cpu0 (and to cpu0.gh at the same time)
	MOVV  t2, m_g0(t1)        // cpu0.mh.g0 = cpu0.gh
	MOVV  t2, m_gsignal(t1)   // cpu0.mh.gsignal = cpu0.gh (to easily check for handler mode)
	MOVV  t1, g_m(t2)         // cpu0.gh.m = cpu0.mh

	// TODO switch to the user mode?

	// create a new goroutine to start program
	SUB	$16, sp
	MOVV	$runtime·mainPC(SB), t0
	MOVV	t0, 8(sp)	// arg 1: fn
	MOVV	$0, t0
	MOVV	t0, 0(sp)	// dummy LR
	JAL	runtime·newproc(SB)
	ADD	$16, sp	// pop args and LR

	// enable interrupts
	// TODO where to enable interupts correctly?
	MOVW z0, M(C0_COMPARE)
	MOVW M(C0_SR), t0
	OR  $(SR_IE|INTR_SW|INTR_EXT), t0
	MOVW t0, M(C0_SR)

	// start this M
	JAL  runtime·mstart(SB)

	UNDEF
