#define LR  RA // link register
#define TMP T6 // REG_TMP

// CSRs

#define fflags   0x001
#define frm      0x002
#define fcsr     0x003 // fflags + frm

#define mstatus  0x300
#define medeleg  0x302
#define mideleg  0x303
#define mie      0x304
#define mtvec    0x305

#define mscratch 0x340
#define mepc     0x341
#define mcause   0x342
#define mtval    0x343
#define mip      0x344

#define mhartid  0xF14

// mstatus field offsets

#define MIEn  3
#define MPIEn 7
#define MPPn 11
#define FSn  13

// mip/mie bits

#define MSI (1<<3)
#define MTI (1<<7)
#define SEI (1<<9)
#define MEI (1<<11)

// core peripherals

#define CLINT_BASE 0x2000000 // true for Qemu Virt, K210, some SiFive cores
#define msip     (CLINT_BASE + 0x0000)
#define mtimecmp (CLINT_BASE + 0x4000)
#define mtime    (CLINT_BASE + 0xBFF8)

#define PLIC_BASE 0x0C000000 // true for Qemu Virt, K210, some SiFive cores
#define PLIC_EN   (PLIC_BASE + 0x002000)
#define PLIC_TC   (PLIC_BASE + 0x200000)

// instructinos not implemented by assembly

#define WFI   WORD $0x10500073
#define MRET  WORD $0x30200073
#define FENCE WORD $0x0ff0000f

#define CSRW(RS,CSR)     WORD $(0x1073 + RS<<15 + CSR<<20)
#define CSRR(CSR,RD)     WORD $(0x2073 + RD<<7 + CSR<<20)
#define CSRS(RS,CSR)     WORD $(0x2073 + RS<<15 + CSR<<20)
#define CSRC(RS,CSR)     WORD $(0x3073 + RS<<15 + CSR<<20)
#define CSRRW(RS,CSR,RD) WORD $(0x1073 + RD<<7 + RS<<15 + CSR<<20)

#define CSRWI(U5,CSR)     WORD $(0x5073 + U5<<15 + CSR<<20)
#define CSRSI(U5,CSR)     WORD $(0x6073 + U5<<15 + CSR<<20)
#define CSRCI(U5,CSR)     WORD $(0x7073 + U5<<15 + CSR<<20)
#define CSRRWI(U5,CSR,RD) WORD $(0x5073 + RD<<7 + U5<<15 + CSR<<20)
#define CSRRCI(U5,CSR,RD) WORD $(0x7073 + RD<<7 + U5<<15 + CSR<<20)

#define LRW(RA, RD)     WORD $(0x1600202F + RD<<7 + RA<<15)
#define LRD(RA, RD)     WORD $(0x1600302F + RD<<7 + RA<<15)
#define SCW(RS, RD, RA) WORD $(0x1E00202F + RD<<7 + RA<<15 + RS<<20)
#define SCD(RS, RD, RA) WORD $(0x1E00302F + RD<<7 + RA<<15 + RS<<20)

// register numbers for above CSR* macros

#define zero 0
#define lr   1
#define x2   2
#define gp   3
#define tp   4
#define t0   5
#define t1   6
#define t2   7
#define s0   8
#define s1   9
#define a0  10
#define a1  11
#define a2  12
#define a3  13
#define a4  14
#define a5  15
#define a6  16
#define a7  17
#define s2  18
#define s3  19
#define s4  20
#define s5  21
#define s6  22
#define s7  23
#define s8  24
#define s9  25
#define s10 26
#define G   27
#define t3  28
#define t4  29
#define t5  30
#define tmp 31

// trap context on the stack
#define _LR (0*8)
#define _A0 (1*8)
#define _mstatus (2*8) // cannot be zero, see: tasker_noos_riscv64.s:/SCW
#define _mepc (3*8)
#define _mie (4*8) // cannot be zero, see: tasker_noos_riscv64.s:/SCW
#define trapCtxSize (5*8)

#define SAVE_GPRS(base,offset) \
\ // LR saved separately
\ // SP saved separately
MOV  GP, (0*8+offset)(base) \
MOV  TP, (1*8+offset)(base) \
MOV  T0, (2*8+offset)(base) \
MOV  T1, (3*8+offset)(base) \
MOV  T2, (4*8+offset)(base) \
MOV  S0, (5*8+offset)(base) \
MOV  S1, (6*8+offset)(base) \
\ // A0 saved separately
MOV  A1, (7*8+offset)(base) \
MOV  A2, (8*8+offset)(base) \
MOV  A3, (9*8+offset)(base) \
MOV  A4, (10*8+offset)(base) \
MOV  A5, (11*8+offset)(base) \
MOV  A6, (12*8+offset)(base) \
MOV  A7, (13*8+offset)(base) \
MOV  S2, (14*8+offset)(base) \
MOV  S3, (15*8+offset)(base) \
MOV  S4, (16*8+offset)(base) \
MOV  S5, (17*8+offset)(base) \
MOV  S6, (18*8+offset)(base) \
MOV  S7, (19*8+offset)(base) \
MOV  S8, (20*8+offset)(base) \
MOV  S9, (21*8+offset)(base) \
MOV  S10, (22*8+offset)(base) \
\ // g saved separately
MOV  T3, (23*8+offset)(base) \
MOV  T4, (24*8+offset)(base) \
MOV  T5, (25*8+offset)(base) \
MOV  TMP, (26*8+offset)(base)

#define RESTORE_GPRS(base,offset) \
\ // LR loaded separately
\ // SP loaded separately
MOV  (0*8+offset)(base), GP \
MOV  (1*8+offset)(base), TP \
MOV  (2*8+offset)(base), T0 \
MOV  (3*8+offset)(base), T1 \
MOV  (4*8+offset)(base), T2 \
MOV  (5*8+offset)(base), S0 \
MOV  (6*8+offset)(base), S1 \
\ // A0 loaded separately
MOV  (7*8+offset)(base), A1 \
MOV  (8*8+offset)(base), A2 \
MOV  (9*8+offset)(base), A3 \
MOV  (10*8+offset)(base), A4 \
MOV  (11*8+offset)(base), A5 \
MOV  (12*8+offset)(base), A6 \
MOV  (13*8+offset)(base), A7 \
MOV  (14*8+offset)(base), S2 \
MOV  (15*8+offset)(base), S3 \
MOV  (16*8+offset)(base), S4 \
MOV  (17*8+offset)(base), S5 \
MOV  (18*8+offset)(base), S6 \
MOV  (19*8+offset)(base), S7 \
MOV  (20*8+offset)(base), S8 \
MOV  (21*8+offset)(base), S9 \
MOV  (22*8+offset)(base), S10 \
\ // g loaded separately
MOV  (23*8+offset)(base), T3 \
MOV  (24*8+offset)(base), T4 \
MOV  (25*8+offset)(base), T5 \
MOV  (26*8+offset)(base), TMP

#define SAVE_FPRS(base,offset) \
CSRR  (fcsr, tmp) \
MOV   TMP, (0*8+offset)(base) \
MOVD  F0, (1*8+offset)(base) \
MOVD  F1, (2*8+offset)(base) \
MOVD  F2, (3*8+offset)(base) \
MOVD  F3, (4*8+offset)(base) \
MOVD  F4, (5*8+offset)(base) \
MOVD  F5, (6*8+offset)(base) \
MOVD  F6, (7*8+offset)(base) \
MOVD  F7, (8*8+offset)(base) \
MOVD  F8, (9*8+offset)(base) \
MOVD  F9, (10*8+offset)(base) \
MOVD  F10, (11*8+offset)(base) \
MOVD  F11, (12*8+offset)(base) \
MOVD  F12, (13*8+offset)(base) \
MOVD  F13, (14*8+offset)(base) \
MOVD  F14, (15*8+offset)(base) \
MOVD  F15, (16*8+offset)(base) \
MOVD  F16, (17*8+offset)(base) \
MOVD  F17, (18*8+offset)(base) \
MOVD  F18, (19*8+offset)(base) \
MOVD  F19, (20*8+offset)(base) \
MOVD  F20, (21*8+offset)(base) \
MOVD  F21, (22*8+offset)(base) \
MOVD  F22, (23*8+offset)(base) \
MOVD  F23, (24*8+offset)(base) \
MOVD  F24, (25*8+offset)(base) \
MOVD  F25, (26*8+offset)(base) \
MOVD  F26, (27*8+offset)(base) \
MOVD  F27, (28*8+offset)(base) \
MOVD  F28, (29*8+offset)(base) \
MOVD  F29, (30*8+offset)(base) \
MOVD  F30, (31*8+offset)(base) \
MOVD  F31, (32*8+offset)(base)

#define RESTORE_FPRS(base,offset) \
MOV   (0*8+offset)(base), TMP \
CSRW  (tmp, fcsr) \
MOVD  (1*8+offset)(base), F0 \
MOVD  (2*8+offset)(base), F1 \
MOVD  (3*8+offset)(base), F2 \
MOVD  (4*8+offset)(base), F3 \
MOVD  (5*8+offset)(base), F4 \
MOVD  (6*8+offset)(base), F5 \
MOVD  (7*8+offset)(base), F6 \
MOVD  (8*8+offset)(base), F7 \
MOVD  (9*8+offset)(base), F8 \
MOVD  (10*8+offset)(base), F9 \
MOVD  (11*8+offset)(base), F10 \
MOVD  (12*8+offset)(base), F11 \
MOVD  (13*8+offset)(base), F12 \
MOVD  (14*8+offset)(base), F13 \
MOVD  (15*8+offset)(base), F14 \
MOVD  (16*8+offset)(base), F15 \
MOVD  (17*8+offset)(base), F16 \
MOVD  (18*8+offset)(base), F17 \
MOVD  (19*8+offset)(base), F18 \
MOVD  (20*8+offset)(base), F19 \
MOVD  (21*8+offset)(base), F20 \
MOVD  (22*8+offset)(base), F21 \
MOVD  (23*8+offset)(base), F22 \
MOVD  (24*8+offset)(base), F23 \
MOVD  (25*8+offset)(base), F24 \
MOVD  (26*8+offset)(base), F25 \
MOVD  (27*8+offset)(base), F26 \
MOVD  (28*8+offset)(base), F27 \
MOVD  (29*8+offset)(base), F26 \
MOVD  (30*8+offset)(base), F29 \
MOVD  (31*8+offset)(base), F30 \
MOVD  (32*8+offset)(base), F31
