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
#define MEI (1<<11)

// core peripherals

#define CLINT_BASE 0x2000000 // true for QEMU, K210 but not standardized
#define msip     (CLINT_BASE + 0x0000)
#define mtimecmp (CLINT_BASE + 0x4000)
#define mtime    (CLINT_BASE + 0xBFF8)

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
#define G    4
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
#define s11 27
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

#define SAVE_GPRS(mp) \
\ // LR saved separately
\ // SP saved separately
MOV  GP, (0*8+m_mOS)(mp) \
\ // g saved separately
MOV  T0, (1*8+m_mOS)(mp) \
MOV  T1, (2*8+m_mOS)(mp) \
MOV  T2, (3*8+m_mOS)(mp) \
MOV  S0, (4*8+m_mOS)(mp) \
MOV  S1, (5*8+m_mOS)(mp) \
\ // A0 saved separately
MOV  A1, (6*8+m_mOS)(mp) \
MOV  A2, (7*8+m_mOS)(mp) \
MOV  A3, (8*8+m_mOS)(mp) \
MOV  A4, (9*8+m_mOS)(mp) \
MOV  A5, (10*8+m_mOS)(mp) \
MOV  A6, (11*8+m_mOS)(mp) \
MOV  A7, (12*8+m_mOS)(mp) \
MOV  S2, (13*8+m_mOS)(mp) \
MOV  S3, (14*8+m_mOS)(mp) \
MOV  S4, (15*8+m_mOS)(mp) \
MOV  S5, (16*8+m_mOS)(mp) \
MOV  S6, (17*8+m_mOS)(mp) \
MOV  S7, (18*8+m_mOS)(mp) \
MOV  S8, (19*8+m_mOS)(mp) \
MOV  S9, (20*8+m_mOS)(mp) \
MOV  S10, (21*8+m_mOS)(mp) \
MOV  S11, (22*8+m_mOS)(mp) \
MOV  T3, (23*8+m_mOS)(mp) \
MOV  T4, (24*8+m_mOS)(mp) \
MOV  T5, (25*8+m_mOS)(mp) \
MOV  TMP, (26*8+m_mOS)(mp)

#define RESTORE_GPRS(mp) \
\ // LR loaded separately
\ // SP loaded separately
MOV  (0*8+m_mOS)(mp), GP \
\ // g loaded separately
MOV  (1*8+m_mOS)(mp), T0 \
MOV  (2*8+m_mOS)(mp), T1 \
MOV  (3*8+m_mOS)(mp), T2 \
MOV  (4*8+m_mOS)(mp), S0 \
MOV  (5*8+m_mOS)(mp), S1 \
\ // A0 loaded separately
MOV  (6*8+m_mOS)(mp), A1 \
MOV  (7*8+m_mOS)(mp), A2 \
MOV  (8*8+m_mOS)(mp), A3 \
MOV  (9*8+m_mOS)(mp), A4 \
MOV  (10*8+m_mOS)(mp), A5 \
MOV  (11*8+m_mOS)(mp), A6 \
MOV  (12*8+m_mOS)(mp), A7 \
MOV  (13*8+m_mOS)(mp), S2 \
MOV  (14*8+m_mOS)(mp), S3 \
MOV  (15*8+m_mOS)(mp), S4 \
MOV  (16*8+m_mOS)(mp), S5 \
MOV  (17*8+m_mOS)(mp), S6 \
MOV  (18*8+m_mOS)(mp), S7 \
MOV  (19*8+m_mOS)(mp), S8 \
MOV  (20*8+m_mOS)(mp), S9 \
MOV  (21*8+m_mOS)(mp), S10 \
MOV  (22*8+m_mOS)(mp), S11 \
MOV  (23*8+m_mOS)(mp), T3 \
MOV  (24*8+m_mOS)(mp), T4 \
MOV  (25*8+m_mOS)(mp), T5 \
MOV  (26*8+m_mOS)(mp), TMP
