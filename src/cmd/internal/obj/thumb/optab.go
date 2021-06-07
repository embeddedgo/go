// Inferno utils/5l/span.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/5l/span.c
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

import (
	"cmd/internal/obj"
	"cmd/internal/objabi"
	"log"
)

type asmoutFunc func(c *Ctx, p *obj.Prog, out []uint16) int

type Optab struct {
	as     obj.As
	a1     Aclass
	a2     Aclass
	a3     Aclass
	size   uint8 // instr_num<<4 | code_size
	flag   uint8
	rscond uint8 // flags required by the instruction
	oscond uint8 // optional flags accepted by the instruction
	asmout asmoutFunc
}

const (
	LFROM = 1 << iota
	LTO
	VALID
	NOIT
)

const pcoff = 4 // in Thumb mode PC points 4 bytes forward

// optab contains description of Thumb2 instruction set in slightly compressed form. buildop
// decompress it to oprange keeping the order of instruction variants unchanged. oplook
// returns the first matched description so the order matters. Current order ensures ISA
// requirements and prefers 16-bit variants.
//
// optab is manual translation of instr_group.txt - please keep them in sync.
// TODO: automatic translation of instr_group.txt.
var optab = [...]Optab{
	// as, a1, a2, a3, size, flag, rscond, oscond, asmout

	{obj.ATEXT, C_LOREG, C_NONE, C_TEXTSIZE, 0, 0, 0, 0, nil},
	{obj.AFUNCDATA, C_LCON, C_NONE, C_LOREG, 0, 0, 0, 0, nil},
	{obj.APCDATA, C_LCON, C_NONE, C_LCON, 0, 0, 0, 0, nil},
	{obj.ANOP, C_NONE, C_NONE, C_NONE, 0, VALID, 0, 0, nil},
	{obj.ANOP, C_LCON, C_NONE, C_NONE, 0, VALID, 0, 0, nil},
	{obj.ANOP, C_REG, C_NONE, C_NONE, 0, VALID, 0, 0, nil},
	{obj.ANOP, C_FREG, C_NONE, C_NONE, 0, VALID, 0, 0, nil},

	{AWORD, C_NONE, C_NONE, C_LCON, 0x14, 0, 0, 0, _WORD__u32},
	{AWORD, C_NONE, C_NONE, C_LOREG, 0x14, 0, 0, 0, _WORD__u32},

	{AHWORD, C_U16CON, C_NONE, C_NONE, 0x12, 0, 0, 0, _HWORD__u16},

	{AADD, C_REG, C_NONE, C_REG, 0x12, 0, 0, C_PBIT, _ADD__Rm__Rdn},        // ADD Rm, Rdn
	{AADD, C_U8CON2, C_SP, C_RLO, 0x12, 0, C_SBIT, 0, _ADD__u8_2__R13__Rd}, // ADD u8<<2, R13, Rd
	{AADD, C_U8CON2, C_PC, C_RLO, 0x12, 0, C_SBIT, 0, _ADD__u8_2__R13__Rd}, // ADD u8<<2, R15, Rd
	{as: obj.AXXX},

	{AADD, C_RLO, C_RLO, C_RLO, 0x12, 0, C_SBIT, 0, _ADD__Rm__Rn__Rd},                    // ADD   Rm, Rn, Rd
	{AADD, C_RLO, C_NONE, C_RLO, 0x12, 0, C_SBIT, 0, _ADD__Rm__Rn__Rd},                   // ADD   Rm, Rdn
	{AADD, C_SHIFTI, C_REG, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _ADDs__Rm_v_u5__Rn__Rd},  // ADD.s Rm<v>u5, Rn, Rd
	{AADD, C_SHIFTI, C_NONE, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _ADDs__Rm_v_u5__Rn__Rd}, // ADD.s Rm<v>u5, Rdn
	{AADD, C_U3CON, C_RLO, C_RLO, 0x12, 0, C_SBIT, 0, _ADD__u3__Rn__Rd},                  // ADD   u3, Rn, Rd
	{AADD, C_U8CON, C_NONE, C_RLO, 0x12, 0, C_SBIT, 0, _ADD__u8__Rdn},                    // ADD   u8, Rdn
	{AADD, C_U7CON2, C_NONE, C_SP, 0x12, 0, C_SBIT, 0, _ADD__u7_2__SP},                   // ADD   u7<<2, R13
	{AADD, C_U12CON, C_REG, C_REG, 0x14, 0, 0, C_PBIT, _ADD__u12__Rn__Rd},                // ADD   u12, Rn, Rd
	{AADD, C_U12CON, C_NONE, C_REG, 0x14, 0, 0, C_PBIT, _ADD__u12__Rn__Rd},               // ADD   u12, Rdn
	{AADD, C_E32CON, C_REG, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _ADDs__e32__Rn__Rd},      // ADD.s e32, Rn, Rd
	{AADD, C_E32CON, C_NONE, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _ADDs__e32__Rn__Rd},     // ADD.s e32, Rdn
	{as: ASUB},

	{AADD, C_LCON, C_NONE, C_REG, 0x26, LFROM, 0, C_PBIT, _ADD__lit__Rdn}, // ADD   lit, Rdn
	{as: obj.AXXX},

	{AADD, C_LCON, C_RLO, C_RLO, 0x26, LFROM, C_SBIT, 0, _ADD__lit__Rn__Rd},           // ADD   lit, Rn, Rd
	{AADD, C_LCON, C_NONE, C_RLO, 0x26, LFROM, C_SBIT, 0, _ADD__lit__Rn__Rd},          // ADD   lit, Rdn
	{AADD, C_LCON, C_REG, C_REG, 0x28, LFROM, 0, C_SBIT | C_PBIT, _ADD__lit__Rn__Rd},  // ADD.s lit, Rn, Rd
	{AADD, C_LCON, C_NONE, C_REG, 0x28, LFROM, 0, C_SBIT | C_PBIT, _ADD__lit__Rn__Rd}, // ADD.s lit, Rdn
	{as: ASUB},

	{AAND, C_RLO, C_NONE, C_RLO, 0x12, 0, C_SBIT, 0, _AND__Rm__Rdn}, // AND Rm, Rdn
	{as: AADC},
	{as: ABIC},
	{as: AEOR},
	{as: AORR},
	{as: ASBC},
	{as: AMUL},
	{as: AMVN},

	{ARSB, C_ZCON, C_RLO, C_RLO, 0x12, 0, C_SBIT, 0, _AND__Rm__Rdn}, // RSB $0, Rn, Rd ; old NEG

	{AAND, C_SHIFTI, C_REG, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _ADDs__Rm_v_u5__Rn__Rd},  // AND.s Rm<v>u5, Rn, Rd
	{AAND, C_SHIFTI, C_NONE, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _ADDs__Rm_v_u5__Rn__Rd}, // AND.s Rm<v>u5, Rdn
	{AAND, C_E32CON, C_REG, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _ADDs__e32__Rn__Rd},      // AND.s e32, Rn, Rd
	{AAND, C_E32CON, C_NONE, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _ADDs__e32__Rn__Rd},     // AND.s e32, Rdn
	{as: AADC},
	{as: ABIC},
	{as: AEOR},
	{as: AORR},
	{as: ARSB},
	{as: ASBC},
	{as: AORN},

	{AAND, C_LCON, C_NONE, C_RLO, 0x26, LFROM, C_SBIT, 0, _ADD__lit__Rdn}, // AND lit, Rdn
	{as: AADC},
	{as: ABIC},
	{as: AEOR},
	{as: AORR},
	{as: ASBC},

	{AAND, C_LCON, C_REG, C_REG, 0x28, LFROM, 0, C_SBIT | C_PBIT, _ADD__lit__Rn__Rd},  // AND.s lit, Rn, Rd
	{AAND, C_LCON, C_NONE, C_REG, 0x28, LFROM, 0, C_SBIT | C_PBIT, _ADD__lit__Rn__Rd}, // AND.s lit, Rdn
	{as: AADC},
	{as: ABIC},
	{as: AEOR},
	{as: AORR},
	{as: ARSB},
	{as: ASBC},
	{as: AORN},

	{AMOVW, C_REG, C_NONE, C_REG, 0x12, 0, 0, C_PBIT, _MOVW__Rm__Rd},                    // MOVW Rm, Rd
	{AMOVW, C_RLO, C_NONE, C_RLO, 0x12, NOIT, C_SBIT, 0, _MOVW__Rm__Rd},                 // MOVW Rm, Rd
	{AMOVW, C_SHIFTILO, C_NONE, C_RLO, 0x12, 0, C_SBIT, 0, _MOVW__Rm_v_u5__Rd},          // MOVW Rm<v>u5, Rd ; vv != 11
	{AMOVW, C_SHIFTRLO, C_NONE, C_NONE, 0x12, 0, C_SBIT, 0, _MOVW__Rdn_v_Rm__Rdn},       // MOVW Rdn<v>Rm, Rdn
	{AMOVW, C_SHIFTR, C_NONE, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _MOVWs__Rn_v_Rm__Rd},  // MOVW.s Rn<v>Rm, Rd
	{AMOVW, C_SHIFTR, C_NONE, C_NONE, 0x14, 0, 0, C_SBIT | C_PBIT, _MOVWs__Rn_v_Rm__Rd}, // MOVW.s Rn<v>Rm, Rd

	{AMOVW, C_U8CON, C_NONE, C_RLO, 0x12, 0, C_SBIT, 0, _ADD__u8__Rdn},     // MOVW u8, Rd
	{AMOVW, C_U16CON, C_NONE, C_REG, 0x14, 0, 0, C_PBIT, _MOVW__uyz16__Rd}, // MOVW uyz16, Rd
	{AMOVT, C_U16CON, C_NONE, C_REG, 0x14, 0, 0, C_PBIT, _MOVW__uyz16__Rd}, // MOVT uyz16, Rd

	{AMOVW, C_SHIFTI, C_NONE, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _ADDs__Rm_v_u5__Rn__Rd}, // MOVW.s Rm<v>u5, Rd
	{AMOVW, C_E32CON, C_NONE, C_REG, 0x14, 0, 0, C_SBIT | C_PBIT, _ADDs__e32__Rn__Rd},     // MOVW.s e32, Rd
	{AMOVW, C_LCON, C_NONE, C_REG, 0x14, LFROM, 0, C_PBIT, _MOVW__lit__Rd},                // MOVW lit, Rd
	{as: AMVN}, // MVN lit, Rd is converted to MOVW -lit, Rd

	{AMOVH, C_RLO, C_NONE, C_RLO, 0x12, 0, 0, 0, _MOVH__Rm__Rd},        // MOVH Rm, Rd
	{AMOVH, C_SHIFTI, C_NONE, C_REG, 0x14, 0, 0, 0, _MOVH__Rm_rot__Rd}, // MOVH Rm@>rot, Rd
	{as: AMOVB},
	{as: AMOVHU},
	{as: AMOVBU},

	{AMOVW, C_SPEC, C_NONE, C_REG, 0x14, 0, 0, 0, _MOVW__SYSm__Rd}, // MOVW SYSm, Rd
	{AMOVW, C_REG, C_NONE, C_SPEC, 0x14, 0, 0, 0, _MOVW__SYSm__Rd}, // MOVW Rn, SYSm

	{AMOVW, C_FCR, C_NONE, C_REG, 0x14, 0, 0, 0, _MOVW__FPSCR__Rt}, // MOVW FPSCR, Rt
	{AMOVW, C_REG, C_NONE, C_FCR, 0x14, 0, 0, 0, _MOVW__FPSCR__Rt}, // MOVW Rt, FPSCR

	{AMUL, C_REG, C_REG, C_REG, 0x14, 0, 0, C_PBIT, _MUL__Rm__Rn__Rd},  // MUL Rm, Rn, Rd
	{AMUL, C_REG, C_NONE, C_REG, 0x14, 0, 0, C_PBIT, _MUL__Rm__Rn__Rd}, // MUL Rm, Rdn
	{as: ADIV},
	{as: ADIVU},

	{AMULL, C_REG, C_REG, C_REGREG, 0x14, 0, 0, 0, _MULL__Rm__Rn__Rdh_Rdl}, // MULL Rm, Rn, (Rdh, Rdl)
	{as: AMULLU},
	{as: AMULAL},
	{as: AMULALU},

	{AMULA, C_REG, C_REG, C_REGREG2, 0x14, 0, 0, 0, _MULA__Rm__Rn__Ra__Rd}, // MULA Rm, Rn, Ra, Rd
	{as: AMULS},
	{as: AMULAWB},
	{as: AMULAWT},

	{AREV, C_RLO, C_NONE, C_RLO, 0x12, 0, 0, 0, _REV__Rm__Rd}, // REV Rm, Rd
	{as: AREV16},
	{as: AREVSH},

	{ACLZ, C_REG, C_NONE, C_REG, 0x14, 0, 0, 0, _CLZ__Rm__Rd}, // CLZ Rm, Rd
	{as: AREV},
	{as: AREV16},
	{as: ARBIT},
	{as: AREVSH},

	{ASEL, C_REG, C_REG, C_REG, 0x14, 0, 0, 0, _CLZ__Rm__Rd},  // SEL Rm, Rn, Rd ; DSP extension
	{ASEL, C_REG, C_NONE, C_REG, 0x14, 0, 0, 0, _CLZ__Rm__Rd}, // SEL Rm, Rd     ; DSP extension

	{ABFX, C_LCON, C_REG, C_REG, 0x14, 0, 0, 0, _BFX__width__ulsb__Rn__Rd},  // BFX width, ulsb, Rn, Rd
	{ABFX, C_LCON, C_NONE, C_REG, 0x14, 0, 0, 0, _BFX__width__ulsb__Rn__Rd}, // BFX width, ulsb, Rdn
	{as: ABFXU},
	{as: ABFI},

	{ABFC, C_LCON, C_NONE, C_REG, 0x14, 0, 0, 0, _BFX__width__ulsb__Rn__Rd}, // BFC width, ulsb, Rd

	{ATST, C_RLO, C_RLO, C_NONE, 0x12, 0, 0, 0, _TST__Rm__Rn}, // TST Rm, Rn
	{as: ACMP},
	{as: ACMN},

	{ACMP, C_REG, C_REG, C_NONE, 0x12, 0, 0, 0, _CMP__Rm__Rn},   // CMP Rm, Rn
	{ACMP, C_U8CON, C_RLO, C_NONE, 0x12, 0, 0, 0, _CMP__u8__Rn}, // CMP u8, Rn

	{ATST, C_SHIFTI, C_REG, C_NONE, 0x14, 0, 0, 0, _TST__Rm_v_u5__Rn}, // TST Rm<v>u5, Rn
	{ATST, C_E32CON, C_REG, C_NONE, 0x14, 0, 0, 0, _TST__e32__Rn},     // TST e32, Rn
	{as: ATEQ},
	{as: ACMN},
	{as: ACMP},

	{ACMN, C_LCON, C_RLO, C_NONE, 0x26, LFROM, 0, 0, _CMN__lit__Rn}, // CMN lit, Rn
	{as: ATST},
	{as: ACMP},

	{ACMP, C_LCON, C_REG, C_NONE, 0x26, LFROM, 0, 0, _CMP__lit__Rn}, // CMP lit, Rn

	{ATST, C_LCON, C_REG, C_NONE, 0x28, LFROM, 0, 0, _TST__lit__Rn}, // TST lit, Rn
	{as: ATEQ},
	{as: ACMN},

	{AMOVW, C_WORLO, C_NONE, C_RLO, 0x12, 0, 0, 0, _MOVW__u5_2_Rn__Rt},  // MOVW  u5<<2(Rn), Rt
	{AMOVW, C_WOSP, C_NONE, C_RLO, 0x12, 0, 0, 0, _MOVW__u8_2_R13__Rt},  // MOVW  u8<<2(R13), Rt
	{AMOVW, C_WOPC, C_NONE, C_RLO, 0x12, 0, 0, 0, _MOVW__u8_2_R13__Rt},  // MOVW  u8<<2(R15), Rt
	{AMOVHU, C_HORLO, C_NONE, C_RLO, 0x12, 0, 0, 0, _MOVW__u5_2_Rn__Rt}, // MOVHU u5<<1(Rn), Rt
	{AMOVBU, C_BORLO, C_NONE, C_RLO, 0x12, 0, 0, 0, _MOVW__u5_2_Rn__Rt}, // MOVBU u5(Rn), Rt

	{AMOVW, C_RLO, C_NONE, C_WORLO, 0x12, 0, 0, 0, _MOVW__u5_2_Rn__Rt}, // MOVW Rt, u5<<2(Rn)
	{AMOVW, C_RLO, C_NONE, C_WOSP, 0x12, 0, 0, 0, _MOVW__u8_2_R13__Rt}, // MOVW Rt, u8<<2(R13)
	{AMOVH, C_RLO, C_NONE, C_HORLO, 0x12, 0, 0, 0, _MOVW__u5_2_Rn__Rt}, // MOVH Rt, u5<<1(Rn)
	{as: AMOVHU},

	{AMOVB, C_RLO, C_NONE, C_BORLO, 0x12, 0, 0, 0, _MOVW__u5_2_Rn__Rt}, // MOVB Rt, u5(Rn)
	{as: AMOVBU},

	{AMOVW, C_RORLO, C_NONE, C_RLO, 0x12, 0, 0, 0, _MOVW__Rn_Rm__Rt},                 // MOVW (Rn)(Rm), Rt
	{AMOVW, C_ROREG, C_NONE, C_REG, 0x14, 0, 0, 0, _MOVW__Rn_Rm_1_u2__Rt},            // MOVW (Rn)(Rm*1<<u2), Rt
	{AMOVW, C_SOPC, C_NONE, C_REG, 0x14, 0, 0, 0, _MOVW__s12_Rn__Rt},                 // MOVW ±u12(R15), Rt
	{AMOVW, C_UOREG, C_NONE, C_REG, 0x14, 0, 0, 0, _MOVW__s12_Rn__Rt},                // MOVW u12(Rn), Rt
	{AMOVW, C_SOREG, C_NONE, C_REG, 0x14, 0, 0, C_PBIT | C_WBIT, _MOVWpw__s8_Rn__Rt}, // MOVW.p.w ±u8(Rn), Rt
	{as: AMOVH},
	{as: AMOVHU},
	{as: AMOVB},
	{as: AMOVBU},

	{AMOVW, C_LORLO, C_NONE, C_RLO, 0x26, LFROM, 0, 0, _MOVW__lit_Rn__Rt}, // MOVW lit(Rn), Rt
	{as: AMOVHU},
	{as: AMOVBU},

	{AMOVW, C_LOREG, C_NONE, C_REG, 0x28, LFROM, 0, 0, _MOVW__lit_Rn__Rt}, // MOVW lit(Rn), Rt
	{as: AMOVH},
	{as: AMOVHU},
	{as: AMOVB},
	{as: AMOVBU},

	{AMOVW, C_RLO, C_NONE, C_RORLO, 0x12, 0, 0, 0, _MOVW__Rn_Rm__Rt},                 // MOVW Rt, (Rn)(Rm)
	{AMOVW, C_REG, C_NONE, C_ROREG, 0x14, 0, 0, 0, _MOVW__Rn_Rm_1_u2__Rt},            // MOVW Rt, (Rn)(Rm*1<<u2)
	{AMOVW, C_REG, C_NONE, C_UOREG, 0x14, 0, 0, 0, _MOVW__s12_Rn__Rt},                // MOVW Rt, u12(Rn)
	{AMOVW, C_REG, C_NONE, C_SOREG, 0x14, 0, 0, C_PBIT | C_WBIT, _MOVWpw__s8_Rn__Rt}, // MOVW.p.w Rt, ±u8(Rn)
	{AMOVW, C_RLO, C_NONE, C_LORLO, 0x26, LTO, 0, 0, _MOVW__lit_Rn__Rt},              // MOVW Rt, lit(Rn)
	{AMOVW, C_REG, C_NONE, C_LOREG, 0x28, LTO, 0, 0, _MOVW__lit_Rn__Rt},              // MOVW Rt, lit(Rn)
	{as: AMOVH},
	{as: AMOVHU},
	{as: AMOVB},
	{as: AMOVBU},

	{AADDF, C_FREG, C_FREG, C_FREG, 0x14, 0, 0, 0, _ADDF__Fm__Fn__Fd}, // ADDF Fm, Fn, Fd
	{AADDF, C_FREG, C_NONE, C_FREG, 0x14, 0, 0, 0, _ADDF__Fm__Fn__Fd}, // ADDF Fm, Fdn
	{as: AADDD},
	{as: ASUBF},
	{as: ASUBD},
	{as: AMULF},
	{as: AMULD},
	{as: ADIVF},
	{as: ADIVD},
	{as: AMULAF},
	{as: AMULAD},
	{as: AMULSF},
	{as: AMULSD},
	{as: ANMULF},
	{as: ANMULD},

	{ASQRTF, C_FREG, C_NONE, C_FREG, 0x14, 0, 0, 0, _SQRTF__Fm__Fd}, // SQRTF Fm, Fd
	{as: ASQRTD},
	{as: ANEGF},
	{as: ANEGD},
	{as: AMOVF},
	{as: AMOVD},
	{as: AABSF},
	{as: AABSD},
	{as: AMOVFD},
	{as: AMOVDF},

	{AMOVW, C_FREG, C_NONE, C_REG, 0x14, 0, 0, 0, _MOVW__Fm__Rd}, // MOVW Fm, Rd},
	{AMOVW, C_REG, C_NONE, C_FREG, 0x14, 0, 0, 0, _MOVW__Fm__Rd}, // MOVW Rt, Fd},

	{AMOVFW, C_FREG, C_NONE, C_REG, 0x28, 0, 0, C_UBIT, _MOVFW__Fm__Rd}, // MOVFW Fm, Rd
	{as: AMOVDW},
	{AMOVWF, C_REG, C_NONE, C_FREG, 0x28, 0, 0, C_UBIT, _MOVFW__Fm__Rd}, // MOVWF Rm, Fd
	{as: AMOVWD},

	{ACMPF, C_FREG, C_FREG, C_NONE, 0x28, 0, 0, 0, _CMPF__Fm__Fd}, // CMPF Fm, Fd
	{ACMPF, C_FREG, C_NONE, C_NONE, 0x28, 0, 0, 0, _CMPF__Fm__Fd}, // CMPF Fd
	{as: ACMPD},

	{AMOVF, C_FOREG, C_NONE, C_FREG, 0x14, 0, 0, 0, _MOVF__s8_2_Rn__Fd},    // MOVF ±u8<<2(Rn), Fd
	{AMOVF, C_FREG, C_NONE, C_FOREG, 0x14, 0, 0, 0, _MOVF__s8_2_Rn__Fd},    // MOVF Fd, ±u8<<2(Rn)
	{AMOVF, C_LOREG, C_NONE, C_FREG, 0x3A, LFROM, 0, 0, _MOVF__lit_Rn__Rt}, // MOVF lit(Rn), Rt
	{AMOVF, C_FREG, C_NONE, C_LOREG, 0x3A, LTO, 0, 0, _MOVF__lit_Rn__Rt},   // MOVF Rt, lit(Rn)
	{AMOVF, C_SFCON, C_NONE, C_FREG, 0x14, 0, 0, 0, _MOVF__f8__Fd},         // MOVF f8, Fd
	{AMOVF, C_ZFCON, C_NONE, C_FREG, 0x28, 0, 0, 0, _MOVF__0__Fd},          // MOVF 0, Fd
	{as: AMOVD},

	{AIT, C_IT, C_NONE, C_NONE, 0x12, 0, 0, 0, _ITmask__firstcond}, // ITmask firstcond
	{as: AITT},
	{as: AITE},
	{as: AITTT},
	{as: AITET},
	{as: AITTE},
	{as: AITEE},
	{as: AITTTT},
	{as: AITETT},
	{as: AITTET},
	{as: AITEET},
	{as: AITTTE},
	{as: AITETE},
	{as: AITTEE},
	{as: AITEEE},

	{AB, C_NONE, C_NONE, C_S11BRA, 0x12, 0, 0, 0, _B__i11_1}, // B i11<<1

	{ACBZ, C_RLO, C_NONE, C_U6BRA, 0x12, 0, 0, 0, _CBZ__Rn__u6_1}, // CBZ Rn, u6<<1
	{as: ACBNZ},

	{AB, C_NONE, C_NONE, C_ZOREG, 0x12, 0, 0, 0, _B__Rm},      // B (Rm)  ; ARM: BX Rm
	{AB, C_NONE, C_NONE, C_S24BRA, 0x14, 0, 0, 0, _B__ji24_1}, // B ji24<<1
	{as: ABL},

	{ABEQ, C_NONE, C_NONE, C_S8BRA, 0x12, 0, 0, 0, _Bcond__i8_1},    // Bcond i8<<1
	{ABEQ, C_NONE, C_NONE, C_S20BRA, 0x14, 0, 0, 0, _Bcond__ji20_1}, // Bcond ji20<<1
	{as: ABNE},
	{as: ABCS},
	{as: ABHS},
	{as: ABCC},
	{as: ABLO},
	{as: ABMI},
	{as: ABPL},
	{as: ABVS},
	{as: ABVC},
	{as: ABHI},
	{as: ABLS},
	{as: ABGE},
	{as: ABLT},
	{as: ABGT},
	{as: ABLE},

	{ATBB, C_REG, C_REG, C_NONE, 0x14, 0, 0, 0, _TBB__Rm__Rn}, // TBB Rm, Rn
	{as: ATBH},

	{ASWI, C_NONE, C_NONE, C_U8CON, 0x12, 0, 0, 0, _SWI__u8}, // SWI u8
	{ASWI, C_NONE, C_NONE, C_NONE, 0x12, 0, 0, 0, _SWI__u8},  // SWI
	{as: ABKPT},
	{as: obj.AUNDEF},

	{AMOVM, C_LISTLOLR, C_NONE, C_ZOSP, 0x12, 0, C_DB | C_WBIT, 0, _PUSH__reglist}, // PUSH reglist (MOVM.DB.W)
	{AMOVM, C_LISTLO, C_NONE, C_ZORLO, 0x12, 0, C_IA | C_WBIT, 0, _MOVM_IAW},       // MOVM.IA.W reglist, (Rn)
	{AMOVM, C_LIST, C_NONE, C_LOREG, 0x14, 0, C_IA, C_WBIT, _MOVM_IAw},             // MOVM.IA.w reglist, (Rn)
	{AMOVM, C_LIST, C_NONE, C_LOREG, 0x14, 0, C_DB, C_WBIT, _MOVM_IAw},             // MOVM.DB.w reglist, (Rn)

	{AMOVM, C_ZOSP, C_NONE, C_LISTLOPC, 0x12, 0, C_IA | C_WBIT, 0, _PUSH__reglist}, // POP reglist (MOVM.IA.W)
	{AMOVM, C_ZORLO, C_NONE, C_LISTLO, 0x12, 0, C_IA, C_WBIT, _MOVM_IAW},           // MOVM.IA.W (Rn), reglist
	{AMOVM, C_LOREG, C_NONE, C_LIST, 0x14, 0, C_IA, C_WBIT, _MOVM_IAw},             // MOVM.IA.w (Rn), reglist
	{AMOVM, C_LOREG, C_NONE, C_LIST, 0x14, 0, C_DB, C_WBIT, _MOVM_IAw},             // MOVM.DB.w (Rn), reglist

	{ALDREX, C_LOREG, C_NONE, C_REG, 0x14, 0, 0, 0, _LDREX__u8_2_Rn__Rt}, // LDREX u8<<2(Rn), Rt
	{ASTREX, C_LOREG, C_REG, C_REG, 0x14, 0, 0, 0, _LDREX__u8_2_Rn__Rt},  // STREX Rt, u8<<2(Rn), Rd

	{ALDREXB, C_LOREG, C_NONE, C_REG, 0x14, 0, 0, 0, _LDREXB__Rn__Rt}, // LDREXB (Rn), Rt
	{as: ALDREXH},

	{ASTREXB, C_LOREG, C_REG, C_REG, 0x14, 0, 0, 0, _LDREXB__Rn__Rt}, // STREXB Rt, (Rn), Rd
	{as: ASTREXH},

	{ANOP2, C_NONE, C_NONE, C_NONE, 0x12, 0, 0, 0, _NOP2},
	{as: AYIELD},
	{as: AWFE},
	{as: AWFI},
	{as: ASEV},

	{ACPSID, C_NONE, C_NONE, C_NONE, 0x12, 0, 0, 0, _CPSID},
	{as: ACPSIE},

	{ANOP4, C_NONE, C_NONE, C_NONE, 0x14, 0, 0, 0, _NOP4},
	{as: ACLREX},

	{ADSB, C_NONE, C_NONE, C_NONE, 0x14, 0, 0, 0, _DSB}, // DSB
	{ADSB, C_MB, C_NONE, C_NONE, 0x14, 0, 0, 0, _DSB},   // DSB opt
	{as: ADMB},
	{as: AISB},
}

// 16-bit instructions

// 0100 0100 dmmm mddd
func _ADD__Rm__Rdn(c *Ctx, p *obj.Prog, out []uint16) int {
	Rm := int(p.From.Reg)
	Rd := int(p.To.Reg)
	out[0] = uint16(0x44<<8 | Rd&8<<4 | Rm&15<<3 | Rd&7)
	return 2
}

// 0001 10xm mmnn nddd
func _ADD__Rm__Rn__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0x1800
	if p.As == ASUB {
		o1 |= 0x0200
	}
	Rm := int(p.From.Reg)
	Rd := int(p.To.Reg)
	Rn := Rd
	if p.Reg != 0 {
		Rn = int(p.Reg)
	}
	out[0] = uint16(o1 | Rm&7<<6 | Rn&7<<3 | Rd&7)
	return 2
}

// 0001 11xu uunn nddd
func _ADD__u3__Rn__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0x1C00
	if p.As == ASUB {
		o1 |= 0x0200
	}
	u3 := int(p.From.Offset)
	Rn := int(p.Reg)
	Rd := int(p.To.Reg)
	out[0] = uint16(o1 | u3<<6 | Rn&7<<3 | Rd&7)
	return 2
}

// 1011 0000 xuuu uuuu
func _ADD__u7_2__SP(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xB000
	if p.As == ASUB {
		o1 |= 0x0080
	}
	u7 := int(p.From.Offset) >> 2
	out[0] = uint16(o1 | u7)
	return 2
}

// 001x xddd uuuu uuuu
func _ADD__u8__Rdn(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0x3000 // AADD
	switch p.As {
	case AMOVW:
		o1 = 0x2000
	case ASUB:
		o1 = 0x3800
	}
	Rd := int(p.To.Reg)
	u8 := int(p.From.Offset)
	out[0] = uint16(o1 | Rd&7<<8 | u8)
	return 2
}

// 1010 xddd uuuu uuuu
func _ADD__u8_2__R13__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xA000
	if p.Reg == REGSP {
		o1 |= 0x0800
	}
	Rd := int(p.To.Reg)
	u8 := int(p.From.Offset) >> 2
	out[0] = uint16(o1 | Rd&7<<8 | u8)
	return 2
}

// 0100 00xx xxmm mddd
func _AND__Rm__Rdn(c *Ctx, p *obj.Prog, out []uint16) int {
	Rm := int(p.From.Reg)
	Rd := int(p.To.Reg)
	var o1 int
	switch p.As {
	case AAND:
		o1 = 0x4000
	case AADC:
		o1 = 0x4140
	case ABIC:
		o1 = 0x4380
	case AEOR:
		o1 = 0x4040
	case AORR:
		o1 = 0x4300
	case ARSB:
		o1 = 0x4240
		Rm = int(p.Reg)
	case ASBC:
		o1 = 0x4180
	case AMUL:
		o1 = 0x4340
	default: // AMVN
		o1 = 0x43C0
	}
	out[0] = uint16(o1 | Rm&7<<3 | Rd&7)
	return 2
}

// 1110 0iii iiii iiii
func _B__i11_1(c *Ctx, p *obj.Prog, out []uint16) int {
	v := int(p.To.Target().Pc-p.Pc-pcoff) >> 1
	out[0] = uint16(0xE000 | v&0x7FF)
	return 2
}

// 0100 0111 xmmm m000
func _B__Rm(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0x4700 | int(p.To.Reg)&15<<3
	if p.As == ABL {
		o1 |= 0x0080
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(p.Pc)
		rel.Siz = 0
		rel.Type = objabi.R_CALLIND
	}
	out[0] = uint16(o1)
	return 2
}

// 1101 cccc iiii iiii
func _Bcond__i8_1(c *Ctx, p *obj.Prog, out []uint16) int {
	v := int(p.To.Target().Pc-p.Pc-pcoff) >> 1
	out[0] = uint16(0xD000 | obcond(p.As)<<8 | v&0xFF)
	return 2
}

// 1011 x0u1 uuuu unnn
func _CBZ__Rn__u6_1(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xB100
	if p.As == ACBNZ {
		o1 |= 0x0800
	}
	Rn := int(p.From.Reg & 7)
	v := int(p.To.Target().Pc-p.Pc-4) >> 1
	out[0] = uint16(o1 | v&0x20<<4 | v&0x1F<<3 | Rn)
	return 2
}

// 0100 0101 nmmm mnnn
func _CMP__Rm__Rn(c *Ctx, p *obj.Prog, out []uint16) int {
	Rm := int(p.From.Reg)
	Rn := int(p.Reg)
	out[0] = uint16(0x4500 | Rn&8<<4 | Rm&15<<3 | Rn&7)
	return 2
}

// 0010 1nnn uuuu uuuu
func _CMP__u8__Rn(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0x2800
	Rn := int(p.Reg)
	u8 := int(p.From.Offset)
	out[0] = uint16(o1 | Rn&7<<8 | u8)
	return 2
}

// 1011 1111 cccc mmmm
func _ITmask__firstcond(c *Ctx, p *obj.Prog, out []uint16) int {
	out[0] = uint16(0xBF00 | int(p.Scond)<<4 | int(p.Mark))
	return 2
}

// 1011 0010 xxmm mddd
func _MOVH__Rm__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xB200 // AMOVH
	switch p.As {
	case AMOVB:
		o1 |= 0x0040
	case AMOVHU:
		o1 |= 0x0080
	case AMOVBU:
		o1 |= 0x00C0
	}
	Rm := int(p.From.Reg)
	Rd := int(p.To.Reg)
	out[0] = uint16(o1 | Rm&7<<3 | Rd&7)
	return 2
}

// 1100 xnnn rrrr rrrr
func _MOVM_IAW(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xC000
	rlist, Rn := int(p.From.Offset), int(p.To.Reg)
	if p.To.Type == obj.TYPE_REGLIST {
		if !c.checkldm(p) {
			return 0
		}
		rlist, Rn = int(p.To.Offset), int(p.From.Reg)
		o1 |= 0x0800
	}
	out[0] = uint16(o1 | Rn&15<<8 | rlist)
	return 2
}

// 0100 000x xxmm mddd (xxx => vv)
func _MOVW__Rdn_v_Rm__Rdn(c *Ctx, p *obj.Prog, out []uint16) int {
	Rdn, typ, Rm := shiftr(int(p.From.Offset))
	var o1 int
	switch typ {
	case 0: // <<
		o1 = 0x4080
	case 1: // >>
		o1 = 0x40C0
	case 2: // ->
		o1 = 0x4100
	default: // @>
		o1 = 0x41C0
	}
	out[0] = uint16(o1 | Rm<<3 | Rdn)
	return 2
}

// 0100 0110 dmmm mddd
// 0000 0000 00mm mddd
func _MOVW__Rm__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	Rm := int(p.From.Reg)
	Rd := int(p.To.Reg)
	if p.Scond&C_SBIT == 0 {
		out[0] = uint16(0x4600 | Rd&8<<4 | Rm&15<<3 | Rd&7)
	} else {
		out[0] = uint16(0x0000 | Rm&7<<3 | Rd&7)
	}
	return 2
}

// 000v vuuu uumm mddd (vv != 11)
func _MOVW__Rm_v_u5__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	Rm, typ, count := shifti(int(p.From.Offset))
	Rd := int(p.To.Reg) & 7
	out[0] = uint16(typ<<11 | count<<6 | Rm<<3 | Rd)
	return 2
}

// 0101 xxxm mmnn nttt
func _MOVW__Rn_Rm__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	var o1 int
	mem, r := &p.From, &p.To
	if r.Type == obj.TYPE_REG {
		switch p.As {
		case AMOVW: // LDR
			o1 = 0x5800
		case AMOVB: // LDRSB
			o1 = 0x5600
		case AMOVBU: // LDRB
			o1 = 0x5C00
		case AMOVH: // LDRSH
			o1 = 0x5E00
		default: // AMOVHU (LDRH)
			o1 = 0x5A00
		}
	} else {
		mem, r = r, mem
		switch p.As {
		case AMOVW: // STR
			o1 = 0x5000
		case AMOVB, AMOVBU: // STRB
			o1 = 0x5400
		default: // AMOVH, AMOVHU (STRH)
			o1 = 0x5200
		}
	}
	Rn := int(mem.Reg)
	Rm := int(mem.Index)
	Rt := int(r.Reg)
	out[0] = uint16(o1 | Rm&7<<6 | Rn&7<<3 | Rt&7)
	return 2
}

// xxxx xuuu uunn nttt
func _MOVW__u5_2_Rn__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	var o1 int
	var shift uint
	mem, r := &p.From, &p.To
	if r.Type == obj.TYPE_REG {
		switch p.As {
		case AMOVW: // LDR
			o1 = 0x6800
			shift = 2
		case AMOVBU: // LDRB
			o1 = 0x7800
		default: // AMOVHU (LDRH)
			o1 = 0x8800
			shift = 1
		}
	} else {
		mem, r = r, mem
		switch p.As {
		case AMOVW: // STR
			o1 = 0x6000
			shift = 2
		case AMOVB, AMOVBU: // STRB
			o1 = 0x7000
		default: // AMOVH, AMOVHU (STRH)
			o1 = 0x8000
			shift = 1
		}
	}
	Rn := int(mem.Reg)
	u5 := int(c.offset(mem)) >> shift
	Rt := int(r.Reg)
	out[0] = uint16(o1 | u5<<6 | Rn&7<<3 | Rt&7)
	return 2
}

// xx0x xttt uuuu uuuu
func _MOVW__u8_2_R13__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	var (
		x uint16
		r int16
		a *obj.Addr
	)
	fr, to := &p.From, &p.To
	switch {
	case fr.Type == obj.TYPE_REG: // MOVW Rt, u8<<2(R13)
		x = 0x9000
		r = fr.Reg
		a = to
	case fr.Reg == REGSP: // MOVW u8<<2(R13), Rt
		x = 0x9800
		r = to.Reg
		a = fr
	default: // MOVW u8<<2(R15), Rt
		x = 0x4800
		r = to.Reg
		a = fr
	}
	out[0] = x | uint16(r&7)<<8 | uint16(c.offset(a)>>2)
	return 2
}

// 1011 1111 0xxx 0000
func _NOP2(c *Ctx, p *obj.Prog, out []uint16) int {
	out[0] = uint16(0xBF00 + int(p.As-ANOP2)<<4)
	return 2
}

// 1011 0110 011x 0010
func _CPSID(c *Ctx, p *obj.Prog, out []uint16) int {
	o := 0xB662
	if p.As == ACPSID {
		o |= 0x10
	}
	out[0] = uint16(o)
	return 2
}

// 1011 x10r rrrr rrrr
func _PUSH__reglist(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xB400
	rlist := int(p.From.Offset)
	if p.To.Type == obj.TYPE_REGLIST {
		rlist = int(p.To.Offset)
		o1 |= 0x0800 | rlist>>7&0x100 // T16 POP supports PC
	} else {
		o1 |= rlist >> 6 & 0x100 // T16 PUSH supports LR
	}
	out[0] = uint16(o1 | rlist&0xFF)
	return 2
}

// 1011 1010 xxmm mddd
func _REV__Rm__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xBA00 // AREV
	switch p.As {
	case AREV16:
		o1 |= 0x0040
	case AREVSH:
		o1 |= 0x00C0
	}
	Rm := int(p.From.Reg)
	Rd := int(p.To.Reg)
	out[0] = uint16(o1 | Rm&7<<3 | Rd&7)
	return 2
}

// 1xx1 111x uuuu uuuu
func _SWI__u8(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xDF00 // ASWI
	switch p.As {
	case obj.AUNDEF:
		o1 = 0xDE00
	case ABKPT:
		o1 = 0xBE00
	}
	u8 := int(p.To.Offset)
	out[0] = uint16(o1 | u8)
	return 2
}

// 0100 0010 xxmm mnnn
func _TST__Rm__Rn(c *Ctx, p *obj.Prog, out []uint16) int {
	var o1 int
	switch p.As {
	case ATST:
		o1 = 0x4200
	case ACMP:
		o1 = 0x4280
	default: // ACMN
		o1 = 0x42C0
	}
	Rm := int(p.From.Reg)
	Rn := int(p.Reg)
	out[0] = uint16(o1 | Rm&7<<3 | Rn&7)
	return 2
}

// 32-bit instructions

// 1111 0exx xxxs nnnn  0eee dddd eeee eeee
func _ADDs__e32__Rn__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2 := oadd32(p)
	e1, e2 := encodeMIC(uint32(p.From.Offset))
	out[0] = 0xF000 | o1 | e1
	out[1] = o2 | e2
	return 4
}

// 1110 101x xxxs nnnn  0uuu dddd uuvv mmmm
func _ADDs__Rm_v_u5__Rn__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2 := oadd32(p)
	out[0] = 0xEA00 | o1
	out[1] = o2 | oshifti32(&p.From)
	return 4
}

// 1111 0u10 x0x0 nnnn  0uuu dddd uuuu uuuu
func _ADD__u12__Rn__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	imm := int(p.From.Offset)
	Rd := int(p.To.Reg) & 15
	Rn := Rd
	if p.Reg != 0 {
		Rn = int(p.Reg) & 15
	}
	o1 := imm>>1&0x400 | Rn
	o2 := imm&0x700<<4 | Rd<<8 | imm&0xFF
	switch p.As {
	case AADD:
		o1 |= 0xF200
	default: // ASUB
		o1 |= 0xF2A0
	}
	out[0] = uint16(o1)
	out[1] = uint16(o2)
	return 4
}

// 1111 0jii iiii iiii  1xj1 jiii iiii iiii
func _B__ji24_1(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xF000
	o2 := 0x9000
	if p.As == ABL {
		o2 |= 0x4000
	}
	v := c.boffsetrel(p, o1, o2)
	s := v >> 23 & 1
	j1 := ^(v>>22 ^ s) & 1
	j2 := ^(v>>21 ^ s) & 1
	imm10 := v >> 11 & 0x3FF
	imm11 := v & 0x7FF
	out[0] = uint16(o1 | s<<10 | imm10)
	out[1] = uint16(o2 | j1<<13 | j2<<11 | imm11)
	return 4
}

// 1111 0jcc ccii iiii  10j0 jiii iiii iiii
func _Bcond__ji20_1(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xF000 | obcond(p.As)<<6
	o2 := 0x8000
	v := c.boffsetrel(p, o1, o2)
	out[0] = uint16(o1 | v>>9&0x400 | v>>11&0x3F)
	out[1] = uint16(o2 | v>>4&0x2000 | v>>7&0x800 | v&0x7FF)
	return 4
}

//1110 1000 1101 nnnn  1111 0000 000h mmmm
func _TBB__Rm__Rn(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2 := 0xE8D0, 0xF000
	if p.As == ATBH {
		o1 |= 0x0010
	}
	Rm := int(p.From.Reg)
	Rn := int(p.Reg)
	out[0] = uint16(o1 | Rn&15)
	out[1] = uint16(o2 | Rm&15)
	return 4
}

// 1111 0011 x1x0 nnnn  0uuu dddd uu0w wwww
func _BFX__width__ulsb__Rn__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	width := int(p.From.Offset)
	if len(p.RestArgs) == 0 {
		c.ctxt.Diag("missing LSB: %v", p)
		return 0
	}
	lsb := int(p.RestArgs[0].Offset)
	if uint(lsb) > 31 || width <= 0 || lsb+width > 32 {
		c.ctxt.Diag("wrong width or LSB: %v", p)
		return 0
	}
	Rd := int(p.To.Reg)
	Rn := Rd
	if p.Reg != 0 {
		Rn = int(p.Reg)
	}
	var o1, o2 int
	switch p.As {
	case ABFX, ABFXU:
		o1 = 0xF340
		if p.As == ABFXU {
			o1 |= 0x0080
		}
		o2 = width - 1
	default: // ABFI, ABFC
		o1 = 0xF360
		if p.As == ABFC {
			o1 |= 0x000F
		}
		o2 = lsb + width - 1
	}
	out[0] = uint16(o1 | Rn&15)
	out[1] = uint16(o2 | lsb&0x1C<<10 | Rd&15<<8 | lsb&3<<6)
	return 4
}

// 1111 1010 10xx nnnn  1111 dddd 10xx mmmm
func _CLZ__Rm__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2 := 0xFA90, 0xF080 // AREV
	Rm := int(p.From.Reg)
	Rn := Rm
	Rd := int(p.To.Reg)
	switch p.As {
	case ACLZ:
		o1 += 0x0020
	case AREV16:
		o2 |= 0x0010
	case ARBIT:
		o2 |= 0x0020
	case AREVSH:
		o2 |= 0x0030
	case ASEL:
		o1 += 0x0010
		if p.Reg != 0 {
			Rn = int(p.Reg)
		} else {
			Rn = Rd
		}
	}
	out[0] = uint16(o1 | Rn&15)
	out[1] = uint16(o2 | Rd&15<<8 | Rm&15)
	return 4
}

func _DSB(c *Ctx, p *obj.Prog, out []uint16) int {
	opt := int(REG_MB_SY)
	if p.From.Reg != 0 {
		opt = int(p.From.Reg)
	}
	out[0] = 0xF3BF
	out[1] = uint16(0x8F40 + int(p.As-ADSB)<<4 + opt&15)
	return 4
}

// 1110 1000 010x nnnn  tttt dddd uuuu uuuu
func _LDREX__u8_2_Rn__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xE840
	if p.As == ALDREX {
		o1 |= 0x0010
	}
	offset := uint64(c.offset(&p.From))
	if offset >= 1<<10 || offset&3 != 0 {
		c.ctxt.Diag("bad offset: %v", p)
		return 0
	}
	u8 := int(offset) >> 2
	Rn := int(p.From.Reg)
	Rt := int(p.To.Reg)
	Rd := int(REGPC)
	if p.Reg != 0 {
		Rd = Rt
		Rt = int(p.Reg)
	}
	out[0] = uint16(o1 | Rn&15)
	out[1] = uint16(Rt&15<<12 | Rd&15<<8 | u8)
	return 4
}

// 1110 1000 110x nnnn  tttt 1111 010x dddd
func _LDREXB__Rn__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	if c.offset(&p.From) != 0 {
		c.ctxt.Diag("offset not supported: %v", p)
		return 0
	}
	o1, o2 := 0xE8C0, 0x0F40 // STREXB
	switch p.As {
	case ALDREXB:
		o1 |= 0x0010
	case ASTREXH:
		o2 |= 0x0010
	case ALDREXH:
		o1 |= 0x0010
		o2 |= 0x0010
	}
	Rn := int(p.From.Reg)
	Rt := int(p.To.Reg)
	Rd := int(REGPC)
	if p.Reg != 0 {
		Rd = Rt
		Rt = int(p.Reg)
	}
	out[0] = uint16(o1 | Rn&15)
	out[1] = uint16(o2 | Rt&15<<12 | Rd&15)
	return 4
}

// 1111 1010 0x0x 1111  1111 dddd 10rr mmmm
func _MOVH__Rm_rot__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2 := 0xFA0F, 0xF080 // AMOVH
	switch p.As {
	case AMOVB:
		o1 |= 0x0040
	case AMOVHU:
		o1 |= 0x0010
	case AMOVBU:
		o1 |= 0x0050
	}
	Rm := int(p.From.Reg) & 15
	rot := 0
	if p.From.Type == obj.TYPE_SHIFT {
		var typ, count int
		Rm, typ, count = shifti(int(p.From.Offset))
		switch count {
		case 0:
			rot = 0
			typ = 3
		case 8:
			rot = 1
		case 16:
			rot = 2
		case 24:
			rot = 3
		default:
			rot = -1
		}
		if typ != 3 || rot < 0 {
			c.ctxt.Diag("only right rotation by 0,8,16,24 is supported: %v", p)
			return 0
		}
	}
	Rd := int(p.To.Reg) & 15
	out[0] = uint16(o1)
	out[1] = uint16(o2 | Rd<<8 | rot<<4 | Rm)
	return 4
}

// 1110 100x x0wx nnnn  rr0r rrrr rrrr rrrr
func _MOVM_IAw(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xE800
	rlist, mem := uint(p.From.Offset), &p.To
	if mem.Type == obj.TYPE_REGLIST {
		if !c.checkldm(p) {
			return 0
		}
		rlist, mem = uint(p.To.Offset), &p.From
		if rlist&0xC000 == 0xC000 {
			c.ctxt.Diag("both R14 and R15 in reglist: %v", p)
			return 0
		}
		o1 |= 0x0010
	} else if rlist&0x8000 != 0 {
		c.ctxt.Diag("R15 in reglist: %v", p)
		return 0
	}
	if rlist&0x2000 != 0 {
		c.ctxt.Diag("R13 in reglist: %v", p)
		return 0
	}
	if c.offset(mem) != 0 {
		c.ctxt.Diag("offset not supported: %v", p)
		return 0
	}
	if onesCount(rlist) < 2 {
		c.ctxt.Diag("to few registers in reglist: %v", p)
		return 0
	}
	if p.Scond&C_DB != 0 {
		o1 |= 0x0100
	} else {
		o1 |= 0x0080
	}
	if p.Scond&C_WBIT != 0 {
		o1 |= 0x0020
	}
	out[0] = uint16(o1 | int(mem.Reg)&15)
	out[1] = uint16(rlist)
	return 4
}

// 1111 100x 0xxx nnnn  tttt 0000 00uu mmmm
func _MOVW__Rn_Rm_1_u2__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2, a := oldrstr32(p)
	var u2 int
	switch a.Scale {
	case 0, 1:
		break
	case 2:
		u2 = 1
	case 4:
		u2 = 2
	case 8:
		u2 = 3
	default:
		log.Fatal("bad scale")
	}
	Rm := int(a.Index)
	out[0] = uint16(o1)
	out[1] = uint16(o2 | u2<<4 | Rm&15)
	return 4
}

// 1111 100x 1xxx nnnn  tttt uuuu uuuu uuuu
// 1111 100x ±xxx 1111  tttt uuuu uuuu uuuu
func _MOVW__s12_Rn__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2, a := oldrstr32(p)
	imm12 := int(c.offset(a))
	if imm12 < 0 {
		imm12 = -imm12
	} else {
		o1 |= 0x0080
	}
	out[0] = uint16(o1)
	out[1] = uint16(o2 | imm12&0xFFF)
	return 4
}

// 1111 0011 1110 1111  1000 dddd mmmm mmmm
// 1111 0011 1000 nnnn  1000 kk00 mmmm mmmm
func _MOVW__SYSm__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	if p.From.Reg >= REG_APSR && p.From.Reg <= REG_CONTROL {
		SYSm := int(p.From.Reg)
		Rd := int(p.To.Reg)
		out[0] = 0xF3EF
		out[1] = uint16(0x8000 | Rd&15<<8 | SYSm&31)
	} else {
		// TODO: support kk != 0b10 (nzcvq)
		Rn := int(p.From.Reg)
		SYSm := int(p.To.Reg)
		out[0] = uint16(0xF380 | Rn&15)
		out[1] = uint16(0x8800 | SYSm&31)
	}
	return 4
}

// 1110 1110 111x 0001  tttt 1010 0001 0000
func _MOVW__FPSCR__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2 := 0xEEE1, 0x0A10
	var Rt int
	if p.From.Reg == REG_FPSCR {
		o1 |= 1 << 4
		Rt = int(p.To.Reg)
	} else {
		Rt = int(p.From.Reg)
	}
	out[0] = uint16(o1)
	out[1] = uint16(o2 | Rt&15<<12)
	return 4
}

// 1111 0y10 x100 uuuu  0zzz dddd zzzz zzzz
func _MOVW__uyz16__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := uint(0xF240)
	if p.As == AMOVT {
		o1 |= 0x80
	}
	imm := uint(p.From.Offset)
	out[0] = uint16(o1 | imm>>12 | imm>>1&0x400)
	out[1] = uint16(imm&0x700<<4 | uint(p.To.Reg)&15<<8 | imm&0xFF)
	return 4
}

// 1111 100x 0xxx nnnn  tttt 1p±w uuuu uuuu
func _MOVWpw__s8_Rn__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2, a := oldrstr32(p)
	imm8 := int(c.offset(a))
	if imm8 < 0 {
		o2 |= 0x0800
		imm8 = -imm8
	} else {
		o2 |= 0x0A00
	}
	switch p.Scond & (C_PBIT | C_WBIT) {
	case 0:
		o2 |= 1 << 10 // offset addressing (set p)
	case C_PBIT:
		o2 |= 1 << 8 // post-indexed addressing (set w)
	case C_WBIT:
		o2 |= 1<<10 | 1<<8 // pre-indexed addressing (set p,w)
	default: // C_PBIT|C_WBIT
		c.ctxt.Diag("invalid .P/.W suffix: %v", p)
	}
	out[0] = uint16(o1)
	out[1] = uint16(o2 | imm8&0xFF)
	return 4
}

// 1111 1010 0vvs nnnn  1111 dddd 0000 mmmm
func _MOVWs__Rn_v_Rm__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xFA00
	if p.Scond&C_SBIT != 0 {
		o1 |= 0x0010
	}
	Rn, typ, Rm := shiftr(int(p.From.Offset))
	Rd := int(p.To.Reg) & 15
	out[0] = uint16(o1 | typ<<5 | Rn)
	out[1] = uint16(0xF000 | Rd<<8 | Rm)
	return 4
}

// 1111 1011 xxxx nnnn  1111 dddd xxxx mmmm
func _MUL__Rm__Rn__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2 := 0xFB00, 0xF000 // MUL
	switch p.As {
	case ADIV:
		o1 |= 0x0090
		o2 |= 0x00F0
	case ADIVU:
		o1 |= 0x00B0
		o2 |= 0x00F0
	}
	Rm := int(p.From.Reg)
	Rd := int(p.To.Reg)
	Rn := Rd
	if p.Reg != 0 {
		Rn = int(p.Reg)
	}
	out[0] = uint16(o1 | Rn&15)
	out[1] = uint16(o2 | Rd&15<<8 | Rm&15)
	return 4
}

// 1111 1011 00xx nnnn  aaaa dddd 000x mmmm
func _MULA__Rm__Rn__Ra__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2 := 0xFB00, 0x0000 // AMULA
	switch p.As {
	case AMULS:
		o2 |= 0x0010
	case AMULAWB:
		o1 |= 0x0030
	case AMULAWT:
		o1 |= 0x0030
		o2 |= 0x0010
	}
	Rm := int(p.From.Reg)
	Rn := int(p.Reg)
	Rd := int(p.To.Reg)
	Ra := int(p.To.Offset)
	out[0] = uint16(o1 | Rn&15)
	out[1] = uint16(o2 | Ra&15<<12 | Rd&15<<8 | Rm&15)
	return 4
}

// 1111 1011 10x0 nnnn  llll hhhh 0000 mmmm
func _MULL__Rm__Rn__Rdh_Rdl(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xFB00
	switch p.As {
	case AMULL:
		o1 |= 0x0080
	case AMULLU:
		o1 |= 0x00A0
	case AMULAL:
		o1 |= 0x00C0
	default: // AMULALU
		o1 |= 0x00E0
	}
	Rm := int(p.From.Reg)
	Rn := int(p.Reg)
	Rdh := int(p.To.Reg)
	Rdl := int(p.To.Offset)
	out[0] = uint16(o1 | Rn&15)
	out[1] = uint16(Rdl&15<<12 | Rdh&15<<8 | Rm&15)
	return 4
}

// 1111 0011 101x 1111  1000 xxxx 00x0 xxxx
func _NOP4(c *Ctx, p *obj.Prog, out []uint16) int {
	var o1, o2 uint16
	switch p.As {
	case ANOP4:
		o1, o2 = 0xF3AF, 0x8000
	default: // ACLREX
		o1, o2 = 0xF3BF, 0x8F2F
	}
	out[0] = o1
	out[1] = o2
	return 4
}

// 1111 0e0x x0x1 nnnn  0eee 1111 eeee eeee
func _TST__e32__Rn(c *Ctx, p *obj.Prog, out []uint16) int {
	o1, o2 := encodeMIC(uint32(p.From.Offset))
	switch p.As {
	case ACMP:
		o1 |= 0xF1B0
	case ATST:
		o1 |= 0xF010
	case ACMN:
		o1 |= 0xF110
	default: // ATEQ
		o1 |= 0xF090
	}
	out[0] = o1 | uint16(p.Reg&15)
	out[1] = o2 | 15<<8
	return 4
}

// 1110 101x x0x1 nnnn  0uuu 1111 uuvv mmmm
func _TST__Rm_v_u5__Rn(c *Ctx, p *obj.Prog, out []uint16) int {
	var o1 int
	switch p.As {
	case ACMP:
		o1 = 0xEBB0
	case ATST:
		o1 = 0xEA10
	case ACMN:
		o1 = 0xEB10
	default: // ATEQ
		o1 = 0xEA90
	}
	out[0] = uint16(o1 | int(p.Reg&15))
	out[1] = 0x0F00 | oshifti32(&p.From)
	return 4
}

// Pseudoinstructions (they must not change the number of instructions declared in optab)

func _ADD__lit__Rdn(c *Ctx, p *obj.Prog, out []uint16) int {
	q := *p
	q.As = AMOVW
	q.To.Reg = REGTMP
	n := _MOVW__lit__Rd(c, &q, out)
	out = out[n/2:]
	q.As = p.As
	q.From = q.To
	q.To = p.To
	if p.As == AADD {
		return n + _ADD__Rm__Rdn(c, &q, out)
	}
	return n + _AND__Rm__Rdn(c, &q, out)
}

func _ADD__lit__Rn__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	q := *p
	q.As = AMOVW
	q.To.Reg = REGTMP
	n := _MOVW__lit__Rd(c, &q, out)
	out = out[n/2:]
	q.As = p.As
	q.From = q.To
	q.To = p.To
	switch q.As {
	case AADD, ASUB:
		if q.Reg <= REG_R7 && q.To.Reg <= REG_R7 {
			return n + _ADD__Rm__Rn__Rd(c, &q, out)
		}
	}
	return n + _ADDs__Rm_v_u5__Rn__Rd(c, &q, out)
}

func _CMN__lit__Rn(c *Ctx, p *obj.Prog, out []uint16) int {
	return c.cmp__lit__Rn(_TST__Rm__Rn, p, out)
}

func _CMP__lit__Rn(c *Ctx, p *obj.Prog, out []uint16) int {
	return c.cmp__lit__Rn(_CMP__Rm__Rn, p, out)
}

func _MOVW__lit__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	v, short := c.litoffset(p, p.To.Reg <= REG_R7, 0)
	var q obj.Prog
	q.As = AMOVW
	q.To = p.To
	q.From.Type = obj.TYPE_MEM
	q.From.Name = obj.NAME_NONE
	q.From.Reg = REGPC
	q.From.Offset = int64(v)
	if short {
		return _MOVW__u8_2_R13__Rt(c, &q, out)
	}
	return _MOVW__s12_Rn__Rt(c, &q, out)
}

func _MOVW__lit_Rn__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	q := *p
	if q.To.Type == obj.TYPE_MEM {
		q.From, q.To.Type = q.To, q.From.Type
	}
	q.As = AMOVW
	q.To.Reg = REGTMP
	n := _MOVW__lit__Rd(c, &q, out)
	out = out[n/2:]
	q = *p
	mem, reg := &q.From, &q.To
	if q.To.Type == obj.TYPE_MEM {
		mem, reg = &q.To, &q.From
	}
	mem.Offset = 0
	switch mem.Name {
	case obj.NAME_STATIC, obj.NAME_EXTERN:
		mem.Reg = REGTMP
		if (q.To.Type == obj.TYPE_MEM || q.As != AMOVB && q.As != AMOVH) && reg.Reg <= REG_R7 {
			return n + _MOVW__u5_2_Rn__Rt(c, &q, out)
		}
		return n + _MOVW__s12_Rn__Rt(c, &q, out)
	default:
		mem.Index = REGTMP
		if mem.Reg <= REG_R7 && reg.Reg <= REG_R7 {
			return n + _MOVW__Rn_Rm__Rt(c, &q, out)
		}
		return n + _MOVW__Rn_Rm_1_u2__Rt(c, &q, out)
	}
}

func _TST__lit__Rn(c *Ctx, p *obj.Prog, out []uint16) int {
	return c.cmp__lit__Rn(_TST__Rm_v_u5__Rn, p, out)
}

func _WORD__u32(c *Ctx, p *obj.Prog, out []uint16) int {
	if p.To.Sym != nil {
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(p.Pc)
		rel.Siz = 4
		rel.Sym = p.To.Sym
		rel.Add = p.To.Offset
		rel.Type = objabi.R_ADDR
		out[0] = 0
		out[1] = 0
	} else {
		offset := int(p.To.Offset)
		out[0] = uint16(offset)
		out[1] = uint16(offset >> 16)
	}
	return 4
}

func _HWORD__u16(c *Ctx, p *obj.Prog, out []uint16) int {
	out[0] = uint16(p.From.Offset)
	return 2
}

// 1110 1101 ±d0x nnnn  dddd 101x uuuu uuuu
func _MOVF__s8_2_Rn__Fd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xED00
	reg, mem := &p.From, &p.To
	if mem.Type == obj.TYPE_REG {
		reg, mem = mem, reg
		o1 |= 0x0010
	}
	o2 := 0x0A00 | int(reg.Reg)&15<<12
	if p.As == AMOVD {
		o2 |= 0x100
	}
	o1 |= int(mem.Reg & 15)
	offset := int(c.offset(mem)) >> 2
	if offset >= 0 {
		o1 |= 0x0080
	} else {
		offset = -offset
	}
	o2 |= offset
	out[0] = uint16(o1)
	out[1] = uint16(o2)
	return 4
}

func _MOVF__lit_Rn__Rt(c *Ctx, p *obj.Prog, out []uint16) int {
	q := *p
	if q.To.Type == obj.TYPE_MEM {
		q.From, q.To.Type = q.To, q.From.Type
	}
	q.As = AMOVW
	tr := q.To.Reg
	q.To.Reg = REGTMP
	n := _MOVW__lit__Rd(c, &q, out)
	if q.From.Name != obj.NAME_STATIC && q.From.Name != obj.NAME_EXTERN {
		q.As = AADD
		q.From.Reg = tr
		n += _ADD__Rm__Rdn(c, &q, out[n/2:])
	}
	q = *p
	mem := &q.From
	if q.To.Type == obj.TYPE_MEM {
		mem = &q.To
	}
	mem.Name = obj.NAME_NONE
	mem.Reg = REGTMP
	mem.Offset = 0
	return n + _MOVF__s8_2_Rn__Fd(c, &q, out[n/2:])
}

// 1110 1110 xdxx nnnn  dddd 101x nxm0 mmmm
func _ADDF__Fm__Fn__Fd(c *Ctx, p *obj.Prog, out []uint16) int {
	Fm := int(p.From.Reg) & 15
	Fd := int(p.To.Reg) & 15
	Fn := Fd
	if p.Reg != 0 {
		Fn = int(p.Reg) & 15
	}
	o1 := 0xEE00 | Fn
	o2 := 0x0A00 | Fd<<12 | Fm
	switch p.As {
	case AADDF:
		o1 |= 0x030
	case AADDD:
		o1 |= 0x030
		o2 |= 0x100
	case ASUBF:
		o1 |= 0x030
		o2 |= 0x040
	case ASUBD:
		o1 |= 0x030
		o2 |= 0x140
	case AMULF:
		o1 |= 0x020
	case AMULD:
		o1 |= 0x020
		o2 |= 0x100
	case ADIVF:
		o1 |= 0x080
	case ADIVD:
		o1 |= 0x080
		o2 |= 0x100
	case AMULAF:
		// nothing
	case AMULAD:
		o2 |= 0x100
	case AMULSF:
		o2 |= 0x040
	case AMULSD:
		o2 |= 0x140
	case ANMULF:
		o1 |= 0x020
		o2 |= 0x040
	default: // ANMULD
		o1 |= 0x020
		o2 |= 0x140
	}
	out[0] = uint16(o1)
	out[1] = uint16(o2)
	return 4
}

// 1110 1110 1d11 0xxx  dddd 101x x1m0 mmmm
func _SQRTF__Fm__Fd(c *Ctx, p *obj.Prog, out []uint16) int {
	Fm := int(p.From.Reg) & 15
	Fd := int(p.To.Reg) & 15
	o1 := 0xEEB0
	o2 := 0x0A40 | Fd<<12 | Fm
	switch p.As {
	case ASQRTF:
		o1 |= 0x001
		o2 |= 0x080
	case ASQRTD:
		o1 |= 0x001
		o2 |= 0x180
	case ANEGF:
		o1 |= 0x001
	case ANEGD:
		o1 |= 0x001
		o2 |= 0x100
	case AMOVF:
		// nothing
	case AMOVD:
		o2 |= 0x100
	case AABSF:
		o2 |= 0x080
	case AABSD:
		o2 |= 0x180
	case AMOVFD:
		o1 |= 0x007
		o2 |= 0x080
	default: // AMOVDF
		o1 |= 0x007
		o2 |= 0x180
	}
	out[0] = uint16(o1)
	out[1] = uint16(o2)
	return 4
}

// 1110 1110 00h1 nnnn  tttt 1011 n001 0000
// 1110 1110 00h0 dddd  tttt 1011 d001 0000
func _MOVW__Fm__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	from := int(p.From.Reg) & 15
	to := int(p.To.Reg) & 15
	o1 := 0xEE00
	o2 := 0x0B10
	if Aclass(p.From.Class-1) == C_FREG {
		o1 |= 0x10 | from
		o2 |= to << 12
	} else {
		o1 |= to
		o2 |= from << 12
	}
	out[0] = uint16(o1)
	out[1] = uint16(o2)
	return 4
}

// 1111 1110 1d11 1xxx  dddd 101x x1m0 mmmm
func _MOVFW__Fm__Rd(c *Ctx, p *obj.Prog, out []uint16) int {
	q := *p
	Xm := int(p.From.Reg) & 15
	Xd := int(p.To.Reg) & 15
	Ftmp := int(FREGTMP) & 15
	switch p.As {
	case AMOVFW, AMOVDW:
		o1 := 0xEEBC // from ARMv7-A RM (the ARMv7-M RM says 0xFEBC)
		if p.Scond&C_UBIT == 0 {
			o1 |= 1
		}
		o2 := 0x0AC0 | Ftmp<<12 | Xm
		if p.As == AMOVDW {
			o2 |= 0x100
		}
		out[0] = uint16(o1)
		out[1] = uint16(o2)
		q.From.Reg = FREGTMP
		_MOVW__Fm__Rd(c, &q, out[2:4])
	default: // AMOVWF, AMOVWD
		q.To.Reg = FREGTMP
		_MOVW__Fm__Rd(c, &q, out[0:2])
		o1 := 0xEEB8 // from ARMv7-A RM (the ARMv7-M RM says 0xFEB8)
		o2 := 0x0A40 | Xd<<12 | Ftmp
		if p.As == AMOVWD {
			o2 |= 0x100
		}
		if p.Scond&C_UBIT == 0 {
			o2 |= 0x80
		}
		out[2] = uint16(o1)
		out[3] = uint16(o2)
	}
	return 8
}

// 1110 1110 1d11 0100  dddd 101x e1m0 mmmm
// 1110 1110 1d11 0101  dddd 101x e100 0000
func _CMPF__Fm__Fd(c *Ctx, p *obj.Prog, out []uint16) int {
	o1 := 0xEEB4
	o2 := 0x0AC0 // e=1
	if p.As == ACMPD {
		o2 |= 0x100
	}
	if p.Reg != 0 {
		o2 |= int(p.Reg&15)<<12 | int(p.From.Reg&15)
	} else {
		o1 |= 1
		o2 |= int(p.From.Reg&15) << 12
	}
	out[0] = uint16(o1)
	out[1] = uint16(o2)
	// move flags from FPSCR to APSR (VMRS R15)
	out[2] = 0xEEF1
	out[3] = 0xFA10
	return 8
}

// 1110 1110 1d11 ffff  dddd 101x 0000 ffff
func _MOVF__f8__Fd(c *Ctx, p *obj.Prog, out []uint16) int {
	v := c.chipfloat(p.From.Val.(float64))
	o1 := 0xEEB0 | v>>4&0xF
	o2 := 0x0A00 | int(p.To.Reg)&15<<12 | v&0xF
	if p.As == AMOVD {
		o2 |= 0x100
	}
	out[0] = uint16(o1)
	out[1] = uint16(o2)
	return 4
}

// 1110 1110 1d11 ffff  dddd 101x 0000 ffff
// 1110 1110 0d11 nnnn  dddd 101x n1m0 mmmm
func _MOVF__0__Fd(c *Ctx, p *obj.Prog, out []uint16) int {
	Rd := int(p.To.Reg) & 15

	// MOVF $1.0, Fd
	o1 := 0xEEB7
	o2 := 0x0A00 | Rd<<12
	if p.As == AMOVD {
		o2 |= 0x100
	}
	out[0] = uint16(o1)
	out[1] = uint16(o2)

	// SUBF Fd, Fd, Fd
	o1 = 0xEE30 | Rd
	o2 = 0x0A40 | Rd<<12 | Rd
	if p.As == AMOVD {
		o2 |= 1 << 8
	}
	out[2] = uint16(o1)
	out[3] = uint16(o2)

	return 8
}

// ********

func oldrstr32(p *obj.Prog) (o1, o2 int, mem *obj.Addr) {
	var r *obj.Addr
	mem, r = &p.From, &p.To
	if r.Type == obj.TYPE_REG {
		switch p.As {
		case AMOVW: // LDR
			o1 = 0xF850
		case AMOVB: // LDRSB
			o1 = 0xF910
		case AMOVBU: // LDRB
			o1 = 0xF810
		case AMOVH: // LDRSH
			o1 = 0xF930
		default: // AMOVHU (LDRH)
			o1 = 0xF830
		}
	} else {
		mem, r = r, mem
		switch p.As {
		case AMOVW: // STR
			o1 = 0xF840
		case AMOVB, AMOVBU: // STRB
			o1 = 0xF800
		default: // AMOVH, AMOVHU (STRH)
			o1 = 0xF820
		}
	}
	o1 |= int(mem.Reg) & 15
	o2 = int(r.Reg) & 15 << 12
	return
}

// u must be representable as Modified Immediate Constant (checked by mic)
func encodeMIC(u uint32) (o1, o2 uint16) {
	if u>>8 == 0 {
		// fast path for 00000000_00000000_00000000_abcdefgh
		return 0, uint16(u)
	}
	for n := 31; n >= 8; n-- {
		if v := rol(u, n); v&^0xFF == 0 {
			o1 = uint16(n & 0x10 << 6)
			o2 = uint16(n&0x0E<<11) | uint16(n&1<<7) | uint16(v&0x7F)
			return
		}
	}
	switch uint32(0) {
	case u >> 24:
		// 00000000_abcdefgh_00000000_abcdefgh
		return 0, 0x1000 | uint16(u&0xFF)
	case u << 24:
		// abcdefgh_00000000_abcdefgh_00000000
		return 0, 0x2000 | uint16(u>>8&0xFF)
	}
	// abcdefgh_abcdefgh_abcdefgh_abcdefgh
	return 0, 0x3000 | uint16(u&0xFF)
}

func oadd32(p *obj.Prog) (o1, o2 uint16) {
	switch p.As {
	case AADD:
		o1 = 0x100
	case ASUB:
		o1 = 0x1A0
	case AAND:
		o1 = 0x000
	case AORR:
		o1 = 0x040
	case ABIC:
		o1 = 0x020
	case AORN:
		o1 = 0x060
	case AEOR:
		o1 = 0x080
	case AADC:
		o1 = 0x140
	case ARSB:
		o1 = 0x1C0
	case ASBC:
		o1 = 0x160
	case AMOVW:
		o1 = 0x04F
	case AMVN:
		o1 = 0x06F
	default:
		log.Fatalf("not a data processing instruction")
	}
	if p.Scond&C_SBIT != 0 {
		o1 |= 0x10
	}
	Rd := int(p.To.Reg) & 15
	Rn := Rd
	if p.Reg != 0 {
		Rn = int(p.Reg) & 15
	}
	return o1 | uint16(Rn), uint16(Rd) << 8
}

func shiftr(offset int) (Rn, typ, Rm int) {
	Rn = offset & 15
	typ = offset >> 5 & 3
	Rm = offset >> 8 & 15
	return
}

func shifti(offset int) (Rm, typ, count int) {
	Rm = offset & 15
	typ = offset >> 5 & 3
	count = offset >> 7 & 31
	return
}

func oshifti32(a *obj.Addr) uint16 {
	Rm := int(a.Reg) & 15
	var typ, count int
	if a.Type == obj.TYPE_SHIFT {
		Rm, typ, count = shifti(int(a.Offset))
	}
	return uint16(count&0x1C<<10 | count&3<<6 | typ<<4 | Rm)
}

func (c *Ctx) checkldm(p *obj.Prog) bool {
	rinlist := 1<<uint(p.From.Reg&15)&p.To.Offset != 0 // true if Rn in reglist
	w := p.Scond&C_WBIT != 0
	if rinlist && w {
		c.ctxt.Diag("Rn in reglist: %v", p)
		return false
	}
	return true
}

func obcond(as obj.As) int {
	switch {
	case as <= ABCS:
		return int(as - ABEQ)
	case as == ABHS:
		return 2
	case as <= ABLO:
		return 3
	default: //  ABMI <= as:
		return 4 + int(as-ABMI)
	}
}

func (c *Ctx) boffsetrel(p *obj.Prog, o1, o2 int) int {
	v := -pcoff
	if p.To.Sym == nil {
		if p.To.Target() != nil {
			v += int(p.To.Target().Pc - p.Pc)
		}
		return v >> 1
	}
	rel := obj.Addrel(c.cursym)
	rel.Off = int32(p.Pc)
	rel.Siz = 4
	rel.Sym = p.To.Sym
	v += int(p.To.Offset)
	rel.Add = int64(o2)<<48 | int64(o1)<<32 | int64(uint32(v))
	rel.Type = objabi.R_CALLARM
	return 0
}

func (c *Ctx) cmp__lit__Rn(cmp asmoutFunc, p *obj.Prog, out []uint16) int {
	q := *p
	q.As = AMOVW
	q.To.Reg = REGTMP
	q.To.Type = obj.TYPE_REG
	n := _MOVW__lit__Rd(c, &q, out)
	out = out[n/2:]
	q.As = p.As
	q.From.Type = obj.TYPE_REG
	q.From.Reg = REGTMP
	return n + cmp(c, &q, out)
}
