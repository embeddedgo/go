/*
 * standard MIPS register names.
 *
 * Copyright (c) 1995 Cygnus Support
 *
 * The authors hereby grant permission to use, copy, modify, distribute,
 * and license this software and its documentation for any purpose, provided
 * that existing copyright notices are retained in all copies and that this
 * notice is included verbatim in any distributions. No written agreement,
 * license, or royalty fee is required for any of the authorized uses.
 * Modifications to this software may be copyrighted by their authors
 * and need not follow the licensing terms described here, provided that
 * the new terms are clearly indicated on the first page of each file where
 * they apply.
 */

/* Standard MIPS register names: */
#define zero	R0
#define z0	R0
#define at	R1
#define v0	R2
#define v1	R3
#define a0	R4
#define a1	R5
#define a2	R6
#define a3	R7
#define t0	R8
#define t1	R9
#define t2	R10
#define t3	R11
#define t4	R12
#define t5	R13
#define t6	R14
#define t7	R15
#define s0	R16
#define s1	R17
#define s2	R18
#define s3	R19
#define s4	R20
#define s5	R21
#define s6	R22
#define s7	R23
#define t8	R24
#define t9	R25
#define k0	R26	/* kernel private register 0 */
#define k1	R27	/* kernel private register 1 */
#define gp	R28	/* global data pointer */
#define sp 	R29	/* stack-pointer */
#define fp	R30	/* frame-pointer */
#define ra	R31	/* return address */

#define fp0	F0
#define fp1	F1

/* Useful memory constants: */
#define K0BASE		0x80000000
#ifndef __mips64
#define K1BASE		0xA0000000
#define K0BASE_ADDR	((char *)K0BASE)
#define K1BASE_ADDR	((char *)K1BASE)
#else
#define K1BASE		0xFFFFFFFFA0000000LL
#define K0BASE_ADDR	((char *)0xFFFFFFFF80000000LL)
#define K1BASE_ADDR	((char *)K1BASE)
#endif

#define PHYS_TO_K1(a)   ((unsigned)(a) | K1BASE)

/* Standard Co-Processor 0 registers */
#define C0_INDEX        0               /* Index of TLB Entry */
#define C0_ENTRYLO0     2               /* TLB entry's first PFN */
#define C0_ENTRYLO1     3               /* TLB entry's second PFN */
#define C0_PAGEMASK     5               /* Size of TLB Entries */
#define C0_COUNT	9		/* Timer Count Register */
#define C0_ENTRYHI      10              /* VPN and ASID of two TLB entry */
#define C0_COMPARE	11		/* Timer Compare Register */
#define C0_SR		12		/* Status Register */
#define C0_CAUSE	13		/* last exception description */
#define C0_EPC		14		/* Exception error address */
#define C0_PRID		15		/* Processor Revision ID */
#define C0_CONFIG	16		/* CPU configuration */
#define C0_WATCHLO	18		/* Watchpoint */

/* Standard Processor Revision ID Register field offsets */
#define PR_IMP	8

/* Standard Config Register field offsets */
#define CR_DB	4
#define CR_IB	5
#define CR_DC	6	/* NOTE v4121 semantics != 43,5xxx semantics */
#define CR_IC	9	/* NOTE v4121 semantics != 43,5xxx semantics */
#define CR_SC	17
#define CR_SS	20
#define CR_SB	22


/* Standard Status Register bitmasks: */
#define SR_CU1		0x20000000	/* Mark CP1 as usable */
#define SR_FR		0x04000000	/* Enable MIPS III FP registers */
#define SR_BEV		0x00400000	/* Controls location of exception vectors */
#define SR_PE		0x00100000	/* Mark soft reset (clear parity error) */

#define SR_KX		0x00000080	/* Kernel extended addressing enabled */
#define SR_SX		0x00000040	/* Supervisor extended addressing enabled */
#define SR_UX		0x00000020	/* User extended addressing enabled */

#define SR_ERL      0x00000004  /* Error level */
#define SR_EXL      0x00000002  /* Exception level */
#define SR_IE       0x00000001  /* Interrupts enabled */

/* Standard Cause Register bitmasks: */
#define CAUSE_EXC_MASK             (0x1F << 2)
#define CAUSE_EXC_INTERRUPT        (0    << 2)
#define CAUSE_EXC_SYSCALL          (8    << 2)
#define CAUSE_EXC_BREAKPOINT       (9    << 2)
#define CAUSE_EXC_COPROCESSOR      (11   << 2)

// Masks for Status Registers IM bits and Cause Registers IP bits
#define INTR_MASK                  (0xFF << 8)
#define INTR_SW                    (0x03 << 8)
#define INTR_EXT                   (0x7C << 8)
#define INTR_TIMER                 (0x80 << 8)

/* Standard (R4000) cache operations. Taken from "MIPS R4000
   Microprocessor User's Manual" 2nd edition: */

#define CACHE_I		(0)	/* primary instruction */
#define CACHE_D		(1)	/* primary data */
#define CACHE_SI	(2)	/* secondary instruction */
#define CACHE_SD	(3)	/* secondary data (or combined instruction/data) */

#define INDEX_INVALIDATE		(0)	/* also encodes WRITEBACK if CACHE_D or CACHE_SD */
#define INDEX_LOAD_TAG			(1)
#define INDEX_STORE_TAG			(2)
#define CREATE_DIRTY_EXCLUSIVE		(3)	/* CACHE_D and CACHE_SD only */
#define HIT_INVALIDATE			(4)
#define CACHE_FILL			(5)	/* CACHE_I only */
#define HIT_WRITEBACK_INVALIDATE	(5)	/* CACHE_D and CACHE_SD only */
#define HIT_WRITEBACK			(6)	/* CACHE_I, CACHE_D and CACHE_SD only */
#define HIT_SET_VIRTUAL			(7)	/* CACHE_SI and CACHE_SD only */

#define BUILD_CACHE_OP(o,c)		(((o) << 2) | (c))

/* Individual cache operations: */
#define INDEX_INVALIDATE_I		BUILD_CACHE_OP(INDEX_INVALIDATE,CACHE_I)
#define INDEX_WRITEBACK_INVALIDATE_D	BUILD_CACHE_OP(INDEX_INVALIDATE,CACHE_D)
#define INDEX_INVALIDATE_SI             BUILD_CACHE_OP(INDEX_INVALIDATE,CACHE_SI)
#define INDEX_WRITEBACK_INVALIDATE_SD	BUILD_CACHE_OP(INDEX_INVALIDATE,CACHE_SD)

#define INDEX_LOAD_TAG_I		BUILD_CACHE_OP(INDEX_LOAD_TAG,CACHE_I)
#define INDEX_LOAD_TAG_D                BUILD_CACHE_OP(INDEX_LOAD_TAG,CACHE_D)
#define INDEX_LOAD_TAG_SI               BUILD_CACHE_OP(INDEX_LOAD_TAG,CACHE_SI)
#define INDEX_LOAD_TAG_SD               BUILD_CACHE_OP(INDEX_LOAD_TAG,CACHE_SD)

#define INDEX_STORE_TAG_I              	BUILD_CACHE_OP(INDEX_STORE_TAG,CACHE_I)
#define INDEX_STORE_TAG_D               BUILD_CACHE_OP(INDEX_STORE_TAG,CACHE_D)
#define INDEX_STORE_TAG_SI              BUILD_CACHE_OP(INDEX_STORE_TAG,CACHE_SI)
#define INDEX_STORE_TAG_SD              BUILD_CACHE_OP(INDEX_STORE_TAG,CACHE_SD)

#define CREATE_DIRTY_EXCLUSIVE_D        BUILD_CACHE_OP(CREATE_DIRTY_EXCLUSIVE,CACHE_D)
#define CREATE_DIRTY_EXCLUSIVE_SD       BUILD_CACHE_OP(CREATE_DIRTY_EXCLUSIVE,CACHE_SD)

#define HIT_INVALIDATE_I                BUILD_CACHE_OP(HIT_INVALIDATE,CACHE_I)
#define HIT_INVALIDATE_D                BUILD_CACHE_OP(HIT_INVALIDATE,CACHE_D)
#define HIT_INVALIDATE_SI               BUILD_CACHE_OP(HIT_INVALIDATE,CACHE_SI)
#define HIT_INVALIDATE_SD               BUILD_CACHE_OP(HIT_INVALIDATE,CACHE_SD)

#define CACHE_FILL_I                    BUILD_CACHE_OP(CACHE_FILL,CACHE_I)
#define HIT_WRITEBACK_INVALIDATE_D      BUILD_CACHE_OP(HIT_WRITEBACK_INVALIDATE,CACHE_D)
#define HIT_WRITEBACK_INVALIDATE_SD     BUILD_CACHE_OP(HIT_WRITEBACK_INVALIDATE,CACHE_SD)

#define HIT_WRITEBACK_I                 BUILD_CACHE_OP(HIT_WRITEBACK,CACHE_I)
#define HIT_WRITEBACK_D                 BUILD_CACHE_OP(HIT_WRITEBACK,CACHE_D)
#define HIT_WRITEBACK_SD                BUILD_CACHE_OP(HIT_WRITEBACK,CACHE_SD)

#define HIT_SET_VIRTUAL_SI		BUILD_CACHE_OP(HIT_SET_VIRTUAL,CACHE_SI)
#define HIT_SET_VIRTUAL_SD              BUILD_CACHE_OP(HIT_SET_VIRTUAL,CACHE_SD)
