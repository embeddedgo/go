// Copyright (c) 1995 Cygnus Support
//
// The authors hereby grant permission to use, copy, modify, distribute, and
// license this software and its documentation for any purpose, provided that
// existing copyright notices are retained in all copies and that this notice is
// included verbatim in any distributions. No written agreement, license, or
// royalty fee is required for any of the authorized uses. Modifications to this
// software may be copyrighted by their authors and need not follow the
// licensing terms described here, provided that the new terms are clearly
// indicated on the first page of each file where they apply.


// Useful memory constants
#define KSEG0  0xFFFFFFFF80000000
#define KSEG1  0xFFFFFFFFA0000000

// Standard Co-Processor 0 registers
#define C0_INDEX     0   // Index of TLB Entry
#define C0_ENTRYLO0  2   // TLB entry's first PFN
#define C0_ENTRYLO1  3   // TLB entry's second PFN
#define C0_PAGEMASK  5   // Size of TLB Entries
#define C0_COUNT     9   // Timer Count Register
#define C0_ENTRYHI   10  // VPN and ASID of two TLB entry
#define C0_COMPARE   11  // Timer Compare Register
#define C0_SR        12  // Status Register
#define C0_CAUSE     13  // last exception description
#define C0_EPC       14  // Exception error address
#define C0_PRID      15  // Processor Revision ID
#define C0_CONFIG    16  // CPU configuration
#define C0_WATCHLO   18  // Watchpoint

// Standard Processor Revision ID Register field offsets
#define PR_IMP  8

// Standard Config Register field offsets
#define CR_DB  4
#define CR_IB  5
#define CR_DC  6
#define CR_IC  9
#define CR_SC  17
#define CR_SS  20
#define CR_SB  22

// Standard Status Register bitmasks
#define SR_CU1  0x20000000  // Mark CP1 as usable
#define SR_FR   0x04000000  // Enable MIPS III FP registers
#define SR_BEV  0x00400000  // Controls location of exception vectors
#define SR_PE   0x00100000  // Mark soft reset (clear parity error)

#define SR_KX   0x00000080  // Kernel extended addressing enabled
#define SR_SX   0x00000040  // Supervisor extended addressing enabled
#define SR_UX   0x00000020  // User extended addressing enabled

#define SR_ERL  0x00000004  // Error level
#define SR_EXL  0x00000002  // Exception level
#define SR_IE   0x00000001  // Interrupts enabled

// Standard Cause Register bitmasks
#define CAUSE_EXC_MASK         (0x1F << 2)
#define CAUSE_EXC_INTERRUPT    (0    << 2)
#define CAUSE_EXC_SYSCALL      (8    << 2)
#define CAUSE_EXC_BREAKPOINT   (9    << 2)
#define CAUSE_EXC_COPROCESSOR  (11   << 2)

// Masks for Status Registers IM bits and Cause Registers IP bits
#define INTR_MASK   (0xFF << 8)
#define INTR_SW0    (0x01 << 8)
#define INTR_SW1    (0x02 << 8)
#define INTR_SW     (0x03 << 8)
#define INTR_EXT    (0x7C << 8)
#define INTR_TIMER  (0x80 << 8)

// BREAK is overloaded CACHE opcode. Register number specifies the cache op.
#define CACHE  BREAK
// Prepend NOOP to ERET to avert CP0 hazards
#define ERET   NOOP; NOOP; WORD $0x42000018

// Individual cache operations
#define INDEX_INVALIDATE_I             R0
#define INDEX_WRITEBACK_INVALIDATE_D   R1
#define INDEX_INVALIDATE_SI            R2
#define INDEX_WRITEBACK_INVALIDATE_SD  R3
#define INDEX_LOAD_TAG_I               R4
#define INDEX_LOAD_TAG_D               R5
#define INDEX_LOAD_TAG_SI              R6
#define INDEX_LOAD_TAG_SD              R7
#define INDEX_STORE_TAG_I              R8
#define INDEX_STORE_TAG_D              R9
#define INDEX_STORE_TAG_SI             R10
#define INDEX_STORE_TAG_SD             R11
#define CREATE_DIRTY_EXCLUSIVE_D       R13
#define CREATE_DIRTY_EXCLUSIVE_SD      R15
#define HIT_INVALIDATE_I               R16
#define HIT_INVALIDATE_D               R17
#define HIT_INVALIDATE_SI              R18
#define HIT_INVALIDATE_SD              R19
#define CACHE_FILL_I                   R20
#define HIT_WRITEBACK_INVALIDATE_D     R21
#define HIT_WRITEBACK_INVALIDATE_SD    R23
#define HIT_WRITEBACK_I                R24
#define HIT_WRITEBACK_D                R25
#define HIT_WRITEBACK_SD               R27
#define HIT_SET_VIRTUAL_SI             R30
#define HIT_SET_VIRTUAL_SD             R31
