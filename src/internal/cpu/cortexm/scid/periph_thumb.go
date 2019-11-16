// Instances:
//  SCID  0xE000E000  -  -  System Control and ID registers
// Registers:
//  0x00  32  MCR    Master Control register, Reserved
//  0x04  32  ICTR   Interrupt Controller Type Register
//  0x08  32  ACTLR  Auxiliary Control Register
package scid

const (
	INTLINESNUM ICTR = 15 << 0 //+ The number of IRQs = 32*(INTLINESNUM+1)
)

const (
	DISMCYCINT ACTLR = 1 << 0 //+
	DISDEFWBUF ACTLR = 1 << 1 //+
	DISFOLD    ACTLR = 1 << 2 //+
	DISFPCA    ACTLR = 1 << 8 //+
	DISOOFP    ACTLR = 1 << 9 //+
)
