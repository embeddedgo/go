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

#define MSI 1<<3
#define MTI 1<<7
#define MEI 1<<11

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


