// Inferno utils/5c/5.out.h
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/5c/5.out.h
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package thumb

import "cmd/internal/obj"

const (
	// must be 16-aligned
	REG_R0 = obj.RBaseThumb + iota
	REG_R1
	REG_R2
	REG_R3
	REG_R4
	REG_R5
	REG_R6
	REG_R7
	REG_R8
	REG_R9
	REG_R10
	REG_R11
	REG_R12
	REG_R13
	REG_R14
	REG_R15

	// must be 16-aligned
	REG_F0
	REG_F1
	REG_F2
	REG_F3
	REG_F4
	REG_F5
	REG_F6
	REG_F7
	REG_F8
	REG_F9
	REG_F10
	REG_F11
	REG_F12
	REG_F13
	REG_F14
	REG_F15

	// must be 32-aligned
	REG_APSR        // 0
	REG_IAPSR       // 1
	REG_EAPSR       // 2
	REG_XPSR        // 3
	_               // 4
	REG_IPSR        // 5
	REG_EPSR        // 6
	REG_IEPSR       // 7
	REG_MSP         // 8
	REG_PSP         // 9
	_               // 10
	_               // 11
	_               // 12
	_               // 13
	_               // 14
	_               // 15
	REG_PRIMASK     // 16
	REG_BASEPRI     // 17
	REG_BASEPRI_MAX // 18
	REG_FAULTMASK   // 19
	REG_CONTROL     // 20

	REG_FPSCR

	// Use R7 as REGTMP insetad of R11 to raise the chance to generate 16-bit
	// instructions. R11 seems to be twice as often used as R7 in arm code:
	// for i in 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15; do
	//   echo -n "R$i "
	//   go tool objdump compile |grep -v 00000000 |grep "R$i[ ,]" |wc -l
	// done
	// R0 461185
	// R1 315395
	// R2 209245
	// R3 164614
	// R4 111672
	// R5 77637
	// R6 58855
	// R7 58084
	// R8 40492
	// R9 35562
	// R10 6508
	// R11 100948
	// R12 29360
	// R13 23314
	// R14 44902
	// R15 6227
	// Drawback: incompatible with arm assembly that use R7.

	REGTMP  = REG_R7  // reserved for compiler and linker
	REGG    = REG_R10 // pointer to goroutine structure (g)
	REGCTXT = REG_R11 // closure pointer
	REGSP   = REG_R13
	REGLINK = REG_R14
	REGPC   = REG_R15

	FREGTMP = REG_F15 // reserved for compiler and linker
)

const pseudoBase = obj.RBaseThumb + 1<<8

const (
	// Pseudo-registers that encode options for DMB, DSB, ISB (must be 16-aligned).
	REG_MB_OSHST = pseudoBase + 2
	REG_MB_OSH   = pseudoBase + 3
	REG_MB_NSHST = pseudoBase + 6
	REG_MB_NSH   = pseudoBase + 7
	REG_MB_ISHST = pseudoBase + 10
	REG_MB_ISH   = pseudoBase + 11
	REG_MB_ST    = pseudoBase + 14
	REG_MB_SY    = pseudoBase + 15

	// Pseudo-registers that encode firstcond for IT (must be 16-aligned).
	REG_EQ = pseudoBase + 16 // equal (Z==1)
	REG_NE = pseudoBase + 17 // not equal (Z == 0)
	REG_HS = pseudoBase + 18 // unsigned higher or same (C == 1)
	REG_CS                   // carry set (C == 1)
	REG_LO = pseudoBase + 19 // unsigned lower (C == 0)
	REG_CC                   // carry clear (C == 0)
	REG_MI = pseudoBase + 20 // minus/negative (N == 1)
	REG_PL = pseudoBase + 21 // plus/positive or zero (N == 0)
	REG_VS = pseudoBase + 22 // overflow (V==1)
	REG_VC = pseudoBase + 23 // no overflow (V==0)
	REG_HI = pseudoBase + 24 // unsigned higher (C==1 && Z==0)
	REG_LS = pseudoBase + 25 // unsigned lower or same (C==0 || Z==1)
	REG_GE = pseudoBase + 26 // signed greater than or equal (N == V)
	REG_LT = pseudoBase + 27 // signed less than (N != V)
	REG_GT = pseudoBase + 28 // signed greater than (Z==0 && N==V)
	REG_LE = pseudoBase + 29 // signed less than or equal (Z==1 || N!=V)
	REG_AL = pseudoBase + 30 // always

	MAXREG
)

type Aclass uint8

//go:generate stringer -type Aclass

const (
	C_NONE Aclass = iota

	// Do not fragment or reorder these register/shift classes, or match/aclass/oplook will break.
	C_RLO      // R0-R7
	C_SP       // R13
	C_PC       // R15
	C_REG      // R0-R15
	C_SHIFTILO // register shift R>>x (R0-R7)
	C_SHIFTI   // register shift R>>x

	C_REGREG
	C_REGREG2 // multiply accumulate dest regs

	// Do not fragment or reorder these registerlist classes, or match/aclass/oplook will break.
	C_LISTLO   // R0-R7
	C_LISTLOLR // R0-R7,LR
	C_LISTLOPC // R0-R7,PC
	C_LIST

	C_SHIFTR   // register shift R>>R
	C_SHIFTRLO // register shift R>>R (R0-R7)
	C_FREG
	C_FCR
	C_SPEC // special register
	C_MB   // memory barier option (pseudo-register)
	C_IT   // firstcond for IT

	// Do not fragment or reorder these C_*CON classes, or match/aclass/oplook will break.

	C_ZCON
	C_U1CON2
	C_U6CON2
	C_U8CON1_4 // not contain C_U8CON2
	C_U8CON5_8
	C_U7CON2 // T16 ADD  u7<<2, R13
	C_U8CON2 // T16 ADD  u8<<2, R13, Rd
	C_U3CON  // T16 ADD  u3, Rn, Rd
	C_U8CON  // T16 ADD  u8, Rdn
	C_U12CON // T32 ADD  u12, Rn, Rd
	C_U16CON // T32 MOVW u16, Rd
	C_E32CON // T32 ADD  e32, Rn, Rd
	C_LCON

	C_ZFCON
	C_SFCON // T32 MOVF f8, Fd
	C_LFCON

	// Do not fragment or reorder these C_*BRA classes, or match/aclass/brlook will break.
	C_U6BRA  // T16 CBZ   Rn, u6<<1
	C_S8BRA  // T16 Bcond i8<<1
	C_S11BRA // T16 B     i11<<1
	C_S20BRA // T32 Bcond ji20<<1
	C_S24BRA // T32 B     ji24<<1

	// Do not fragment or reorder these C_*OR* classes, or match/aclass/oplook will break.
	C_BORLO   // T16 MOVB       u5(Rn), Rt
	C_HORLO   // T16 MOVH       u5<<1(Rn), Rt
	C_WORLO   // T16 MOVW       u5<<2(Rn), Rt
	C_WOSP    // T16 MOVW       u8<<2(R13), Rt
	C_WOPC    // T16 MOVW       u8<<2(PC), Rt
	C_UOREG   // T32 MOV{B,H,W} u12(Rn), Rt
	C_SOREG   // T32 MOV{B,H,W} ±u8(Rn), Rt
	C_SOPC    // T32 MOVW       ±u12(R15), Rt
	C_FOREG   // T32 MOV{F,D}   ±u8<<2(Rn), Rt (MOVW ±u8<<2(Rn), (Rta, Rtb))
	C_ZORLO   // T16 MOVM.IA.W  (Rn), reglist
	C_ZOSP    // T16 MOVM.IA.W  (R13), reglist
	C_ZOREG   // T16 JMP        (Rn)
	C_LOREG   // long offset
	C_LORLO   // long offset, Rn <= R7
	C_U0ORLO  // (Rn)        C_BORLO, C_HORLO, C_WORLO, C_UOREG, C_SOREG, C_FOREG, ZORLO, ZOREG
	C_U3ORLO2 // u3<<2(Rn)   C_BORLO, C_HORLO, C_WORLO, C_UOREG, C_SOREG, C_FOREG
	C_U4ORLO2 // u4<<2(Rn)   C_HORLO, C_WORLO, C_UOREG, C_SOREG, C_FOREG
	C_U5ORLO2 // u5<<2(Rn)   C_WORLO, C_UOREG, C_SOREG, C_FOREG
	C_U4ORLO1 // u4<<1(Rn)   C_BORLO, C_HORLO, C_UOREG, C_SOREG
	C_U5ORLO1 // u5<<1(Rn)   C_HORLO, C_UOREG, C_SOREG
	C_U5ORLO  // u5(Rn)      C_BORLO, C_UOREG, C_SOREG
	C_U0OPC   // (R15)       C_WOPC,  C_SOPC,  C_UOREG, C_SOREG, C_FOREG, ZOREG
	C_U6OPC2  // u6<<2(R15)  C_WOPC,  C_SOPC,  C_UOREG, C_SOREG, C_FOREG
	C_U8OPC2  // u8<<2(R15)  C_WOPC,  C_SOPC,  C_UOREG, C_FOREG
	C_U12OPC  // u12(R15)    C_SOPC,  C_UOREG
	C_S8OPC   // -u8(R15)    C_SOPC,  C_SOREG
	C_S8OPC2  // -u8<<2(R15) C_SOPC,  C_FOREG
	C_S12OPC  // -u12(R15)   C_SOPC
	C_U0OSP   // (R13)       C_WOSP,  C_UOREG, C_SOREG, C_FOREG, ZOSP, ZOREG
	C_U6OSP2  // u6<<2(R13)  C_WOSP,  C_UOREG, C_SOREG, C_FOREG
	C_U8OSP2  // u8<<2(R13)  C_WOSP,  C_UOREG, C_FOREG
	C_U0OREG  // (Rn)        C_UOREG, C_SOREG, C_FOREG, ZOREG
	C_U6OREG2 // u6<<2(Rn)   C_UOREG, C_SOREG, C_FOREG
	C_U8OREG2 // u8<<2(Rn)   C_UOREG, C_FOREG
	C_U8OREG  // u8(Rn)      C_UOREG, C_SOREG
	C_U12OREG // u12(Rn)     C_UOREG
	C_S6OREG2 // -u6<<2(Rn)  C_SOREG, C_FOREG
	C_S8OREG2 // -u8<<2(Rn)  C_FOREG
	C_S8OREG  // -u8(Rn)     C_SOREG

	C_ROREG // (Rn)(Rm*x)
	C_RORLO // (Rn)(Rm)

	C_TEXTSIZE

	C_GOK
)

//go:generate go run ../stringer.go -i $GOFILE -o anames.go -p thumb

const (
	AAND = obj.ABaseThumb + obj.A_ARCHSPECIFIC + iota
	AEOR
	ASUB
	ARSB
	AADD
	AADC
	ASBC
	ATST
	ATEQ
	ACMP
	ACMN
	AORR
	ABIC

	AMVN
	AORN

	AMUL
	ADIV
	ADIVU

	AMULL
	AMULLU
	AMULAL
	AMULALU
	AMULA
	AMULS
	AMULAWB
	AMULAWT

	ACLZ
	AREV
	AREV16
	ARBIT
	AREVSH
	ASEL
	ABFX
	ABFXU
	ABFC
	ABFI

	// Do not reorder or fragment the conditional branch
	// opcodes, or the predication code will break.
	ABEQ
	ABNE
	ABCS
	ABHS
	ABCC
	ABLO
	ABMI
	ABPL
	ABVS
	ABVC
	ABHI
	ABLS
	ABGE
	ABLT
	ABGT
	ABLE
	ACBZ
	ACBNZ

	ATBB
	ATBH

	// do not reorder or split AITxyz
	AITTTT
	AITTT
	AITTTE
	AITT
	AITTET
	AITTE
	AITTEE
	AIT
	AITETT
	AITET
	AITETE
	AITE
	AITEET
	AITEE
	AITEEE

	// do not reorder or split NOP-compatible opcodes
	ANOP2
	AYIELD
	AWFE
	AWFI
	ASEV

	ACPSID
	ACPSIE

	// do not reorder or split memory barrier opcodes
	ADSB
	ADMB
	AISB

	AMOVB
	AMOVBU
	AMOVH
	AMOVHU
	AMOVW
	AMOVM
	AMOVT

	ACMPF
	ACMPD

	ASQRTF
	ASQRTD
	AMOVF
	AMOVD

	ASWI
	ABKPT

	ALDREX
	ALDREXB
	ALDREXH
	ASTREX
	ASTREXB
	ASTREXH
	ACLREX

	ANOP4

	// do not reorder shift opcodes
	ASLL
	ASRL
	ASRA

	// not implemented
	AMMULA
	AMMULS
	AMULABB

	AADDF
	AADDD
	ASUBF
	ASUBD
	AMULF
	AMULD
	AMULAF
	AMULAD
	AMULSF
	AMULSD
	ANEGF
	ANEGD
	ANMULF
	ANMULD
	ADIVF
	ADIVD
	AMOVWF
	AMOVWD
	AMOVFW
	AMOVDW
	AMOVFD
	AMOVDF

	AWORD
	AHWORD

	ALAST

	// aliases
	AB     = obj.AJMP  // B, BX
	ABL    = obj.ACALL // BL, BLX
	AMOVBS = AMOVB
	AMOVHS = AMOVH
)

// scond byte
const (
	C_SCOND = (1 << 4) - 1
	C_SBIT  = 1 << 4
	C_PBIT  = 1 << 5
	C_WBIT  = 1 << 6
	C_FBIT  = 1 << 7 /* psr flags-only */
	C_UBIT  = 1 << 7 /* up bit, unsigned bit */
	C_IA    = C_UBIT
	C_DB    = C_PBIT

	// These constants are the ARM condition codes encodings,
	// XORed with 14 so that C_SCOND_NONE has value 0,
	// so that a zeroed Prog.scond means "always execute".
	C_SCOND_XOR = 14

	C_SCOND_EQ   = 0 ^ C_SCOND_XOR
	C_SCOND_NE   = 1 ^ C_SCOND_XOR
	C_SCOND_HS   = 2 ^ C_SCOND_XOR // CS
	C_SCOND_LO   = 3 ^ C_SCOND_XOR // CC
	C_SCOND_MI   = 4 ^ C_SCOND_XOR
	C_SCOND_PL   = 5 ^ C_SCOND_XOR
	C_SCOND_VS   = 6 ^ C_SCOND_XOR
	C_SCOND_VC   = 7 ^ C_SCOND_XOR
	C_SCOND_HI   = 8 ^ C_SCOND_XOR
	C_SCOND_LS   = 9 ^ C_SCOND_XOR
	C_SCOND_GE   = 10 ^ C_SCOND_XOR
	C_SCOND_LT   = 11 ^ C_SCOND_XOR
	C_SCOND_GT   = 12 ^ C_SCOND_XOR
	C_SCOND_LE   = 13 ^ C_SCOND_XOR
	C_SCOND_NONE = 14 ^ C_SCOND_XOR
	C_SCOND_NV   = 15 ^ C_SCOND_XOR

	/* D_SHIFT type */
	SHIFT_LL = 0 << 5
	SHIFT_LR = 1 << 5
	SHIFT_AR = 2 << 5
	SHIFT_RR = 3 << 5
)

// http://infocenter.arm.com/help/topic/com.arm.doc.ihi0040b/IHI0040B_aadwarf.pdf
var ARMDWARFRegisters = map[int16]int16{}

func init() {
	// f assigns dwarfregisters[from:to] = (base):(step*(to-from)+base)
	f := func(from, to, base, step int16) {
		for r := int16(from); r <= to; r++ {
			ARMDWARFRegisters[r] = step*(r-from) + base
		}
	}
	f(REG_R0, REG_R15, 0, 1)
	f(REG_F0, REG_F15, 64, 2) // Use d0 through D15, aka S0, S2, ..., S30
}
