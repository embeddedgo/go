#define LR  RA // link register
#define TMP T6 // REG_TMP

// CSRs

#define FFLAGS   0x001
#define FRM      0x002
#define FCSR     0x003 // FFLAGS + FRM

#define MSTATUS  0x300
#define MEDELEG  0x302
#define MIDELEG  0x303
#define MIE      0x304
#define MTVEC    0x305

#define MSCRATCH 0x340
#define MEPC     0x341
#define MCAUSE   0x342
#define MTVAL    0x343
#define MIP      0x344

#define MHARTID  0xF14

// mstatus field offsets

#define MIEn  3
#define MPIEn 7
#define FSn  13

// instructinos not implemented by assembly

#define WFI  WORD $0x10500073
#define MRET WORD $0x30200073

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

// register numbers for CSR* macros

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


