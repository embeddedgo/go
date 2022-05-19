// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package thumb

import (
	"fmt"
	"internal/buildcfg"
	"math"
	"math/bits"

	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/logopt"
	"cmd/compile/internal/ssa"
	"cmd/compile/internal/ssagen"
	"cmd/compile/internal/types"
	"cmd/internal/obj"
	"cmd/internal/obj/thumb"
)

// loadByType returns the load instruction of the given type.
func loadByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size() {
		case 4:
			return thumb.AMOVF
		case 8:
			return thumb.AMOVD
		}
	} else {
		switch t.Size() {
		case 1:
			if t.IsSigned() {
				return thumb.AMOVB
			} else {
				return thumb.AMOVBU
			}
		case 2:
			if t.IsSigned() {
				return thumb.AMOVH
			} else {
				return thumb.AMOVHU
			}
		case 4:
			return thumb.AMOVW
		}
	}
	panic("bad load type")
}

// storeByType returns the store instruction of the given type.
func storeByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size() {
		case 4:
			return thumb.AMOVF
		case 8:
			return thumb.AMOVD
		}
	} else {
		switch t.Size() {
		case 1:
			return thumb.AMOVB
		case 2:
			return thumb.AMOVH
		case 4:
			return thumb.AMOVW
		}
	}
	panic("bad store type")
}

// shift type is used as Offset in obj.TYPE_SHIFT operands to encode shifted register operands
type shift int64

// copied from ../../../internal/obj/util.go:/TYPE_SHIFT
func (v shift) String() string {
	op := "<<>>->@>"[((v>>5)&3)<<1:]
	if v&(1<<4) != 0 {
		// register shift
		return fmt.Sprintf("R%d%c%cR%d", v&15, op[0], op[1], (v>>8)&15)
	} else {
		// constant shift
		return fmt.Sprintf("R%d%c%c%d", v&15, op[0], op[1], (v>>7)&31)
	}
}

// makeshift encodes a register shifted by a constant
func makeshift(v *ssa.Value, reg int16, typ int64, s int64) shift {
	if s < 0 || s >= 32 {
		v.Fatalf("shift out of range: %d", s)
	}
	return shift(int64(reg&0xf) | typ | (s&31)<<7)
}

// genshift generates a Prog for r = r0 op (r1 shifted by n)
func genshift(s *ssagen.State, v *ssa.Value, as obj.As, r0, r1, r int16, typ int64, n int64) *obj.Prog {
	p := s.Prog(as)
	p.From.Type = obj.TYPE_SHIFT
	p.From.Offset = int64(makeshift(v, r1, typ, n))
	p.Reg = r0
	if r != 0 {
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	}
	return p
}

// find a (lsb, width) pair for BFC
// lsb must be in [0, 31], width must be in [1, 32 - lsb]
// return (0xffffffff, 0) if v is not a binary like 0...01...10...0
func getBFC(v uint32) (uint32, uint32) {
	var m, l uint32
	// BFC is not applicable with zero
	if v == 0 {
		return 0xffffffff, 0
	}
	// find the lowest set bit, for example l=2 for 0x3ffffffc
	l = uint32(bits.TrailingZeros32(v))
	// m-1 represents the highest set bit index, for example m=30 for 0x3ffffffc
	m = 32 - uint32(bits.LeadingZeros32(v))
	// check if v is a binary like 0...01...10...0
	if (1<<m)-(1<<l) == v {
		// it must be m > l for non-zero v
		return l, m - l
	}
	// invalid
	return 0xffffffff, 0
}

func ssaGenValue(s *ssagen.State, v *ssa.Value) {
	switch v.Op {
	case ssa.OpCopy, ssa.OpThumbMOVWreg:
		if v.Type.IsMemory() {
			return
		}
		x := v.Args[0].Reg()
		y := v.Reg()
		if x == y {
			return
		}
		as := thumb.AMOVW
		if v.Type.IsFloat() {
			switch v.Type.Size() {
			case 4:
				as = thumb.AMOVF
			case 8:
				as = thumb.AMOVD
			default:
				panic("bad float size")
			}
		}
		p := s.Prog(as)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x
		p.To.Type = obj.TYPE_REG
		p.To.Reg = y
	case ssa.OpThumbMOVWnop:
		// nothing to do
	case ssa.OpLoadReg:
		if v.Type.IsFlags() {
			v.Fatalf("load flags not implemented: %v", v.LongString())
			return
		}
		p := s.Prog(loadByType(v.Type))
		ssagen.AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpStoreReg:
		if v.Type.IsFlags() {
			v.Fatalf("store flags not implemented: %v", v.LongString())
			return
		}
		p := s.Prog(storeByType(v.Type))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		ssagen.AddrAuto(&p.To, v)
	case ssa.OpThumbADD,
		ssa.OpThumbADC,
		ssa.OpThumbSUB,
		ssa.OpThumbSBC,
		ssa.OpThumbRSB,
		ssa.OpThumbAND,
		ssa.OpThumbOR,
		ssa.OpThumbORN,
		ssa.OpThumbXOR,
		ssa.OpThumbBIC,
		ssa.OpThumbMUL,
		ssa.OpThumbDIV,
		ssa.OpThumbDIVU,
		ssa.OpThumbADDF,
		ssa.OpThumbADDD,
		ssa.OpThumbSUBF,
		ssa.OpThumbSUBD,
		ssa.OpThumbSLL,
		ssa.OpThumbSRL,
		ssa.OpThumbSRA,
		ssa.OpThumbSRR,
		ssa.OpThumbMULF,
		ssa.OpThumbMULD,
		ssa.OpThumbNMULF,
		ssa.OpThumbNMULD,
		ssa.OpThumbDIVF,
		ssa.OpThumbDIVD:
		r := v.Reg()
		r1 := v.Args[0].Reg()
		r2 := v.Args[1].Reg()
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r2
		p.Reg = r1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpThumbMULAF, ssa.OpThumbMULAD, ssa.OpThumbMULSF, ssa.OpThumbMULSD, ssa.OpThumbFMULAD:
		r := v.Reg()
		r0 := v.Args[0].Reg()
		r1 := v.Args[1].Reg()
		r2 := v.Args[2].Reg()
		if r != r0 {
			v.Fatalf("result and addend are not in the same register: %v", v.LongString())
		}
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r2
		p.Reg = r1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpThumbADDS,
		ssa.OpThumbSUBS:
		r := v.Reg0()
		r1 := v.Args[0].Reg()
		r2 := v.Args[1].Reg()
		p := s.Prog(v.Op.Asm())
		p.Scond = thumb.C_SBIT
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r2
		p.Reg = r1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpThumbSRAcond:
		// Thumb shift instructions uses only the low-order byte of the shift amount
		// generate conditional instructions to deal with large shifts
		// flag is already set
		// SRA.HS	$31, Rarg0, Rdst // shift 31 bits to get the sign bit
		// SRA.LO	Rarg1, Rarg0, Rdst
		r := v.Reg()
		r1 := v.Args[0].Reg()
		r2 := v.Args[1].Reg()
		p := s.Prog(thumb.ASRA)
		p.Scond = thumb.C_SCOND_HS
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 31
		p.Reg = r1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p = s.Prog(thumb.ASRA)
		p.Scond = thumb.C_SCOND_LO
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r2
		p.Reg = r1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpThumbBFX, ssa.OpThumbBFXU:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt >> 8
		p.SetFrom3Const(v.AuxInt & 0xff)
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbANDconst, ssa.OpThumbBICconst:
		// try to optimize ANDconst and BICconst to BFC, which saves bytes and ticks
		// BFC is only available on ARMv7, and its result and source are in the same register
		if v.Reg() == v.Args[0].Reg() {
			var val uint32
			if v.Op == ssa.OpThumbANDconst {
				val = ^uint32(v.AuxInt)
			} else { // BICconst
				val = uint32(v.AuxInt)
			}
			lsb, width := getBFC(val)
			// omit BFC for ARM's imm12
			if 8 < width && width < 24 {
				p := s.Prog(thumb.ABFC)
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = int64(width)
				p.SetFrom3Const(int64(lsb))
				p.To.Type = obj.TYPE_REG
				p.To.Reg = v.Reg()
				break
			}
		}
		// fall back to ordinary form
		fallthrough
	case ssa.OpThumbADDconst,
		ssa.OpThumbADCconst,
		ssa.OpThumbSUBconst,
		ssa.OpThumbSBCconst,
		ssa.OpThumbRSBconst,
		ssa.OpThumbORconst,
		ssa.OpThumbORNconst,
		ssa.OpThumbXORconst,
		ssa.OpThumbSLLconst,
		ssa.OpThumbSRLconst,
		ssa.OpThumbSRAconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbADDSconst,
		ssa.OpThumbSUBSconst,
		ssa.OpThumbRSBSconst:
		p := s.Prog(v.Op.Asm())
		p.Scond = thumb.C_SBIT
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0()
	case ssa.OpThumbSRRconst:
		genshift(s, v, thumb.AMOVW, 0, v.Args[0].Reg(), v.Reg(), thumb.SHIFT_RR, v.AuxInt)
	case ssa.OpThumbADDshiftLL,
		ssa.OpThumbADCshiftLL,
		ssa.OpThumbSUBshiftLL,
		ssa.OpThumbSBCshiftLL,
		ssa.OpThumbRSBshiftLL,
		ssa.OpThumbANDshiftLL,
		ssa.OpThumbORshiftLL,
		ssa.OpThumbORNshiftLL,
		ssa.OpThumbXORshiftLL,
		ssa.OpThumbBICshiftLL:
		genshift(s, v, v.Op.Asm(), v.Args[0].Reg(), v.Args[1].Reg(), v.Reg(), thumb.SHIFT_LL, v.AuxInt)
	case ssa.OpThumbADDSshiftLL,
		ssa.OpThumbSUBSshiftLL,
		ssa.OpThumbRSBSshiftLL:
		p := genshift(s, v, v.Op.Asm(), v.Args[0].Reg(), v.Args[1].Reg(), v.Reg0(), thumb.SHIFT_LL, v.AuxInt)
		p.Scond = thumb.C_SBIT
	case ssa.OpThumbADDshiftRL,
		ssa.OpThumbADCshiftRL,
		ssa.OpThumbSUBshiftRL,
		ssa.OpThumbSBCshiftRL,
		ssa.OpThumbRSBshiftRL,
		ssa.OpThumbANDshiftRL,
		ssa.OpThumbORshiftRL,
		ssa.OpThumbORNshiftRL,
		ssa.OpThumbXORshiftRL,
		ssa.OpThumbBICshiftRL:
		genshift(s, v, v.Op.Asm(), v.Args[0].Reg(), v.Args[1].Reg(), v.Reg(), thumb.SHIFT_LR, v.AuxInt)
	case ssa.OpThumbADDSshiftRL,
		ssa.OpThumbSUBSshiftRL,
		ssa.OpThumbRSBSshiftRL:
		p := genshift(s, v, v.Op.Asm(), v.Args[0].Reg(), v.Args[1].Reg(), v.Reg0(), thumb.SHIFT_LR, v.AuxInt)
		p.Scond = thumb.C_SBIT
	case ssa.OpThumbADDshiftRA,
		ssa.OpThumbADCshiftRA,
		ssa.OpThumbSUBshiftRA,
		ssa.OpThumbSBCshiftRA,
		ssa.OpThumbRSBshiftRA,
		ssa.OpThumbANDshiftRA,
		ssa.OpThumbORshiftRA,
		ssa.OpThumbORNshiftRA,
		ssa.OpThumbXORshiftRA,
		ssa.OpThumbBICshiftRA:
		genshift(s, v, v.Op.Asm(), v.Args[0].Reg(), v.Args[1].Reg(), v.Reg(), thumb.SHIFT_AR, v.AuxInt)
	case ssa.OpThumbADDSshiftRA,
		ssa.OpThumbSUBSshiftRA,
		ssa.OpThumbRSBSshiftRA:
		p := genshift(s, v, v.Op.Asm(), v.Args[0].Reg(), v.Args[1].Reg(), v.Reg0(), thumb.SHIFT_AR, v.AuxInt)
		p.Scond = thumb.C_SBIT
	case ssa.OpThumbXORshiftRR:
		genshift(s, v, v.Op.Asm(), v.Args[0].Reg(), v.Args[1].Reg(), v.Reg(), thumb.SHIFT_RR, v.AuxInt)
	case ssa.OpThumbMVNshiftLL:
		genshift(s, v, v.Op.Asm(), 0, v.Args[0].Reg(), v.Reg(), thumb.SHIFT_LL, v.AuxInt)
	case ssa.OpThumbMVNshiftRL:
		genshift(s, v, v.Op.Asm(), 0, v.Args[0].Reg(), v.Reg(), thumb.SHIFT_LR, v.AuxInt)
	case ssa.OpThumbMVNshiftRA:
		genshift(s, v, v.Op.Asm(), 0, v.Args[0].Reg(), v.Reg(), thumb.SHIFT_AR, v.AuxInt)
	case ssa.OpThumbHMUL,
		ssa.OpThumbHMULU:
		// 32-bit high multiplication
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_REGREG
		p.To.Reg = v.Reg()
		p.To.Offset = thumb.REGTMP // throw away low 32-bit into tmp register
	case ssa.OpThumbMULLU:
		// 32-bit multiplication, results 64-bit, high 32-bit in out0, low 32-bit in out1
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_REGREG
		p.To.Reg = v.Reg0()           // high 32-bit
		p.To.Offset = int64(v.Reg1()) // low 32-bit
	case ssa.OpThumbMULA, ssa.OpThumbMULS:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_REGREG2
		p.To.Reg = v.Reg()                   // result
		p.To.Offset = int64(v.Args[2].Reg()) // addend
	case ssa.OpThumbMOVWconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbMOVFconst,
		ssa.OpThumbMOVDconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbCMP,
		ssa.OpThumbCMN,
		ssa.OpThumbTST,
		ssa.OpThumbTEQ,
		ssa.OpThumbCMPF,
		ssa.OpThumbCMPD:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		// Special layout in ARM assembly
		// Comparing to x86, the operands of ARM's CMP are reversed.
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()
	case ssa.OpThumbCMPconst,
		ssa.OpThumbCMNconst,
		ssa.OpThumbTSTconst,
		ssa.OpThumbTEQconst:
		// Special layout in ARM assembly
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg()
	case ssa.OpThumbCMPF0,
		ssa.OpThumbCMPD0:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
	case ssa.OpThumbCMPshiftLL, ssa.OpThumbCMNshiftLL, ssa.OpThumbTSTshiftLL, ssa.OpThumbTEQshiftLL:
		genshift(s, v, v.Op.Asm(), v.Args[0].Reg(), v.Args[1].Reg(), 0, thumb.SHIFT_LL, v.AuxInt)
	case ssa.OpThumbCMPshiftRL, ssa.OpThumbCMNshiftRL, ssa.OpThumbTSTshiftRL, ssa.OpThumbTEQshiftRL:
		genshift(s, v, v.Op.Asm(), v.Args[0].Reg(), v.Args[1].Reg(), 0, thumb.SHIFT_LR, v.AuxInt)
	case ssa.OpThumbCMPshiftRA, ssa.OpThumbCMNshiftRA, ssa.OpThumbTSTshiftRA, ssa.OpThumbTEQshiftRA:
		genshift(s, v, v.Op.Asm(), v.Args[0].Reg(), v.Args[1].Reg(), 0, thumb.SHIFT_AR, v.AuxInt)
	case ssa.OpThumbMOVWaddr:
		p := s.Prog(thumb.AMOVW)
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

		var wantreg string
		// MOVW $sym+off(base), R
		// the assembler expands it as the following:
		// - base is SP: add constant offset to SP (R13)
		//               when constant is large, tmp register (R11) may be used
		// - base is SB: load external address from constant pool (use relocation)
		switch v.Aux.(type) {
		default:
			v.Fatalf("aux is of unknown type %T", v.Aux)
		case *obj.LSym:
			wantreg = "SB"
			ssagen.AddAux(&p.From, v)
		case *ir.Name:
			wantreg = "SP"
			ssagen.AddAux(&p.From, v)
		case nil:
			// No sym, just MOVW $off(SP), R
			wantreg = "SP"
			p.From.Offset = v.AuxInt
		}
		if reg := v.Args[0].RegName(); reg != wantreg {
			v.Fatalf("bad reg %s for symbol type %T, want %s", reg, v.Aux, wantreg)
		}
	case ssa.OpThumbMOVWload,
		ssa.OpThumbMOVHload,
		ssa.OpThumbMOVHUload,
		ssa.OpThumbMOVBload,
		ssa.OpThumbMOVBUload,
		ssa.OpThumbMOVDload,
		ssa.OpThumbMOVFload,
		ssa.OpThumbLoadOnce32,
		ssa.OpThumbLoadOnce16,
		ssa.OpThumbLoadOnce8:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		ssagen.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		if _, ok := v.Block.Func.RegAlloc[v.ID].(ssa.LocPair); ok {
			p.To.Reg = v.Reg0() // LoadOnce*
		} else {
			p.To.Reg = v.Reg()
		}
	case ssa.OpThumbMOVWstore,
		ssa.OpThumbMOVHstore,
		ssa.OpThumbMOVBstore,
		ssa.OpThumbMOVDstore,
		ssa.OpThumbMOVFstore,
		ssa.OpThumbStoreOnce32,
		ssa.OpThumbStoreOnce16,
		ssa.OpThumbStoreOnce8:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		ssagen.AddAux(&p.To, v)
	case ssa.OpThumbMOVWloadidx,
		ssa.OpThumbMOVHUloadidx,
		ssa.OpThumbMOVHloadidx,
		ssa.OpThumbMOVBUloadidx,
		ssa.OpThumbMOVBloadidx,
		ssa.OpThumbLoadOnce32idx,
		ssa.OpThumbLoadOnce16idx,
		ssa.OpThumbLoadOnce8idx:
		// this is just shift 0 bits
		fallthrough
	case ssa.OpThumbMOVWloadshiftLL,
		ssa.OpThumbMOVHUloadshiftLL,
		ssa.OpThumbMOVHloadshiftLL,
		ssa.OpThumbMOVBUloadshiftLL,
		ssa.OpThumbMOVBloadshiftLL,
		ssa.OpThumbLoadOnce32shiftLL,
		ssa.OpThumbLoadOnce16shiftLL,
		ssa.OpThumbLoadOnce8shiftLL:
		var toreg int16
		if _, ok := v.Block.Func.RegAlloc[v.ID].(ssa.LocPair); ok {
			toreg = v.Reg0() // LoadOnce*
		} else {
			toreg = v.Reg()
		}
		p := genshift(s, v, v.Op.Asm(), 0, v.Args[1].Reg(), toreg, thumb.SHIFT_LL, v.AuxInt)
		p.From.Reg = v.Args[0].Reg()
	case ssa.OpThumbMOVWstoreidx,
		ssa.OpThumbMOVHstoreidx,
		ssa.OpThumbMOVBstoreidx,
		ssa.OpThumbStoreOnce32idx,
		ssa.OpThumbStoreOnce16idx,
		ssa.OpThumbStoreOnce8idx:
		// this is just shift 0 bits
		fallthrough
	case ssa.OpThumbMOVWstoreshiftLL,
		ssa.OpThumbMOVHstoreshiftLL,
		ssa.OpThumbMOVBstoreshiftLL,
		ssa.OpThumbStoreOnce32shiftLL,
		ssa.OpThumbStoreOnce16shiftLL,
		ssa.OpThumbStoreOnce8shiftLL:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg()
		p.To.Type = obj.TYPE_SHIFT
		p.To.Reg = v.Args[0].Reg()
		p.To.Offset = int64(makeshift(v, v.Args[1].Reg(), thumb.SHIFT_LL, v.AuxInt))
	case ssa.OpThumbMOVBreg,
		ssa.OpThumbMOVBUreg,
		ssa.OpThumbMOVHreg,
		ssa.OpThumbMOVHUreg:
		a := v.Args[0]
		for a.Op == ssa.OpCopy || a.Op == ssa.OpThumbMOVWreg || a.Op == ssa.OpThumbMOVWnop {
			a = a.Args[0]
		}
		if a.Op == ssa.OpLoadReg {
			t := a.Type
			switch {
			case v.Op == ssa.OpThumbMOVBreg && t.Size() == 1 && t.IsSigned(),
				v.Op == ssa.OpThumbMOVBUreg && t.Size() == 1 && !t.IsSigned(),
				v.Op == ssa.OpThumbMOVHreg && t.Size() == 2 && t.IsSigned(),
				v.Op == ssa.OpThumbMOVHUreg && t.Size() == 2 && !t.IsSigned():
				// arg is a proper-typed load, already zero/sign-extended, don't extend again
				if v.Reg() == v.Args[0].Reg() {
					return
				}
				p := s.Prog(thumb.AMOVW)
				p.From.Type = obj.TYPE_REG
				p.From.Reg = v.Args[0].Reg()
				p.To.Type = obj.TYPE_REG
				p.To.Reg = v.Reg()
				return
			default:
			}
		}
		fallthrough
	case ssa.OpThumbMVN,
		ssa.OpThumbCLZ,
		ssa.OpThumbREV,
		ssa.OpThumbREV16,
		ssa.OpThumbRBIT,
		ssa.OpThumbSQRTF,
		ssa.OpThumbSQRTD,
		ssa.OpThumbNEGF,
		ssa.OpThumbNEGD,
		ssa.OpThumbABSD,
		ssa.OpThumbMOVWF,
		ssa.OpThumbMOVWD,
		ssa.OpThumbMOVFW,
		ssa.OpThumbMOVDW,
		ssa.OpThumbMOVFD,
		ssa.OpThumbMOVDF:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbMOVWUF,
		ssa.OpThumbMOVWUD,
		ssa.OpThumbMOVFWU,
		ssa.OpThumbMOVDWU:
		p := s.Prog(v.Op.Asm())
		p.Scond = thumb.C_UBIT
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbCMOVWHSconst:
		p := s.Prog(thumb.AMOVW)
		p.Scond = thumb.C_SCOND_HS
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbCMOVWLSconst:
		p := s.Prog(thumb.AMOVW)
		p.Scond = thumb.C_SCOND_LS
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbCALLstatic, ssa.OpThumbCALLclosure, ssa.OpThumbCALLinter:
		s.Call(v)
	case ssa.OpThumbCALLtail:
		s.TailCall(v)
	case ssa.OpThumbLoweredWB:
		p := s.Prog(obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = v.Aux.(*obj.LSym)
	case ssa.OpThumbLoweredPanicBoundsA, ssa.OpThumbLoweredPanicBoundsB, ssa.OpThumbLoweredPanicBoundsC:
		p := s.Prog(obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = ssagen.BoundsCheckFunc[v.AuxInt]
		s.UseArgs(8) // space used in callee args area by assembly stubs
	case ssa.OpThumbLoweredPanicExtendA, ssa.OpThumbLoweredPanicExtendB, ssa.OpThumbLoweredPanicExtendC:
		p := s.Prog(obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = ssagen.ExtendCheckFunc[v.AuxInt]
		s.UseArgs(12) // space used in callee args area by assembly stubs
	case ssa.OpThumbDUFFZERO:
		p := s.Prog(obj.ADUFFZERO)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = ir.Syms.Duffzero
		p.To.Offset = v.AuxInt
	case ssa.OpThumbDUFFCOPY:
		p := s.Prog(obj.ADUFFCOPY)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = ir.Syms.Duffcopy
		p.To.Offset = v.AuxInt
	case ssa.OpThumbLoweredNilCheck:
		if buildcfg.GOOS == "noos" {
			// BUG: avoid nil check because of MMIO
		} else {
			// Issue a load which will fault if arg is nil.
			p := s.Prog(thumb.AMOVBU)
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = v.Args[0].Reg()
			ssagen.AddAux(&p.From, v)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = thumb.REGTMP
			if logopt.Enabled() {
				logopt.LogOpt(v.Pos, "nilcheck", "genssa", v.Block.Func.Name)
			}
			if base.Debug.Nil != 0 && v.Pos.Line() > 1 { // v.Pos.Line()==1 in generated wrappers
				base.WarnfAt(v.Pos, "generated nil check")
			}
		}
	case ssa.OpThumbLoweredZero:
		// MOVW.P	Rarg2, 4(R1)
		// CMP	Rarg1, R1
		// BLE	-2(PC)
		// arg1 is the address of the last element to zero
		// arg2 is known to be zero
		// auxint is alignment
		var sz int64
		var mov obj.As
		switch {
		case v.AuxInt%4 == 0:
			sz = 4
			mov = thumb.AMOVW
		case v.AuxInt%2 == 0:
			sz = 2
			mov = thumb.AMOVH
		default:
			sz = 1
			mov = thumb.AMOVB
		}
		p := s.Prog(mov)
		p.Scond = thumb.C_PBIT
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = thumb.REG_R1
		p.To.Offset = sz
		p2 := s.Prog(thumb.ACMP)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = v.Args[1].Reg()
		p2.Reg = thumb.REG_R1
		p3 := s.Prog(thumb.ABLE)
		p3.To.Type = obj.TYPE_BRANCH
		p3.To.SetTarget(p)
	case ssa.OpThumbLoweredMove:
		// MOVW.P	4(R1), Rtmp
		// MOVW.P	Rtmp, 4(R2)
		// CMP	Rarg2, R1
		// BLE	-3(PC)
		// arg2 is the address of the last element of src
		// auxint is alignment
		var sz int64
		var mov obj.As
		switch {
		case v.AuxInt%4 == 0:
			sz = 4
			mov = thumb.AMOVW
		case v.AuxInt%2 == 0:
			sz = 2
			mov = thumb.AMOVH
		default:
			sz = 1
			mov = thumb.AMOVB
		}
		p := s.Prog(mov)
		p.Scond = thumb.C_PBIT
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = thumb.REG_R1
		p.From.Offset = sz
		p.To.Type = obj.TYPE_REG
		p.To.Reg = thumb.REGTMP
		p2 := s.Prog(mov)
		p2.Scond = thumb.C_PBIT
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = thumb.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = thumb.REG_R2
		p2.To.Offset = sz
		p3 := s.Prog(thumb.ACMP)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = v.Args[2].Reg()
		p3.Reg = thumb.REG_R1
		p4 := s.Prog(thumb.ABLE)
		p4.To.Type = obj.TYPE_BRANCH
		p4.To.SetTarget(p)
	case ssa.OpThumbEqual,
		ssa.OpThumbNotEqual,
		ssa.OpThumbLessThan,
		ssa.OpThumbLessEqual,
		ssa.OpThumbGreaterThan,
		ssa.OpThumbGreaterEqual,
		ssa.OpThumbLessThanU,
		ssa.OpThumbLessEqualU,
		ssa.OpThumbGreaterThanU,
		ssa.OpThumbGreaterEqualU:
		// generate boolean values
		// use conditional move, preserve flags
		scond := condBits[v.Op] | thumb.C_PBIT
		p := s.Prog(thumb.AMOVW)
		p.Scond = scond ^ 1
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
		p = s.Prog(thumb.AMOVW)
		p.Scond = scond
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbLoweredGetClosurePtr:
		// Closure pointer is R11 (thumb.REGCTXT).
		ssagen.CheckLoweredGetClosurePtr(v)
	case ssa.OpThumbLoweredGetCallerSP:
		// caller's SP is FixedFrameSize below the address of the first arg
		p := s.Prog(thumb.AMOVW)
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = -base.Ctxt.FixedFrameSize()
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbLoweredGetCallerPC:
		p := s.Prog(obj.AGETCALLERPC)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpThumbFlagConstant:
		v.Fatalf("FlagConstant op should never make it to codegen %v", v.LongString())
	case ssa.OpThumbInvertFlags:
		v.Fatalf("InvertFlags should never make it to codegen %v", v.LongString())
	case ssa.OpClobber, ssa.OpClobberReg:
		// TODO: implement for clobberdead experiment. Nop is ok for now.
	case ssa.OpThumbDSB:
		s.Prog(thumb.ADSB)
	case ssa.OpThumbDMB_ST:
		p := s.Prog(thumb.ADMB)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = thumb.REG_MB_ST
	default:
		v.Fatalf("genValue not implemented: %s", v.LongString())
	}
}

var condBits = map[ssa.Op]uint8{
	ssa.OpThumbEqual:         thumb.C_SCOND_EQ,
	ssa.OpThumbNotEqual:      thumb.C_SCOND_NE,
	ssa.OpThumbLessThan:      thumb.C_SCOND_LT,
	ssa.OpThumbLessThanU:     thumb.C_SCOND_LO,
	ssa.OpThumbLessEqual:     thumb.C_SCOND_LE,
	ssa.OpThumbLessEqualU:    thumb.C_SCOND_LS,
	ssa.OpThumbGreaterThan:   thumb.C_SCOND_GT,
	ssa.OpThumbGreaterThanU:  thumb.C_SCOND_HI,
	ssa.OpThumbGreaterEqual:  thumb.C_SCOND_GE,
	ssa.OpThumbGreaterEqualU: thumb.C_SCOND_HS,
}

var blockJump = map[ssa.BlockKind]struct {
	asm, invasm obj.As
}{
	ssa.BlockThumbEQ:     {thumb.ABEQ, thumb.ABNE},
	ssa.BlockThumbNE:     {thumb.ABNE, thumb.ABEQ},
	ssa.BlockThumbLT:     {thumb.ABLT, thumb.ABGE},
	ssa.BlockThumbGE:     {thumb.ABGE, thumb.ABLT},
	ssa.BlockThumbLE:     {thumb.ABLE, thumb.ABGT},
	ssa.BlockThumbGT:     {thumb.ABGT, thumb.ABLE},
	ssa.BlockThumbULT:    {thumb.ABLO, thumb.ABHS},
	ssa.BlockThumbUGE:    {thumb.ABHS, thumb.ABLO},
	ssa.BlockThumbUGT:    {thumb.ABHI, thumb.ABLS},
	ssa.BlockThumbULE:    {thumb.ABLS, thumb.ABHI},
	ssa.BlockThumbLTnoov: {thumb.ABMI, thumb.ABPL},
	ssa.BlockThumbGEnoov: {thumb.ABPL, thumb.ABMI},
}

// To model a 'LEnoov' ('<=' without overflow checking) branching
var leJumps = [2][2]ssagen.IndexJump{
	{{Jump: thumb.ABEQ, Index: 0}, {Jump: thumb.ABPL, Index: 1}}, // next == b.Succs[0]
	{{Jump: thumb.ABMI, Index: 0}, {Jump: thumb.ABEQ, Index: 0}}, // next == b.Succs[1]
}

// To model a 'GTnoov' ('>' without overflow checking) branching
var gtJumps = [2][2]ssagen.IndexJump{
	{{Jump: thumb.ABMI, Index: 1}, {Jump: thumb.ABEQ, Index: 1}}, // next == b.Succs[0]
	{{Jump: thumb.ABEQ, Index: 1}, {Jump: thumb.ABPL, Index: 0}}, // next == b.Succs[1]
}

func ssaGenBlock(s *ssagen.State, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockPlain:
		if b.Succs[0].Block() != next {
			p := s.Prog(obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, ssagen.Branch{P: p, B: b.Succs[0].Block()})
		}

	case ssa.BlockDefer:
		// defer returns in R0:
		// 0 if we should continue executing
		// 1 if we should jump to deferreturn call
		p := s.Prog(thumb.ACMP)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 0
		p.Reg = thumb.REG_R0
		p = s.Prog(thumb.ABNE)
		p.To.Type = obj.TYPE_BRANCH
		s.Branches = append(s.Branches, ssagen.Branch{P: p, B: b.Succs[1].Block()})
		if b.Succs[0].Block() != next {
			p := s.Prog(obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, ssagen.Branch{P: p, B: b.Succs[0].Block()})
		}

	case ssa.BlockExit, ssa.BlockRetJmp:

	case ssa.BlockRet:
		s.Prog(obj.ARET)

	case ssa.BlockThumbEQ, ssa.BlockThumbNE,
		ssa.BlockThumbLT, ssa.BlockThumbGE,
		ssa.BlockThumbLE, ssa.BlockThumbGT,
		ssa.BlockThumbULT, ssa.BlockThumbUGT,
		ssa.BlockThumbULE, ssa.BlockThumbUGE,
		ssa.BlockThumbLTnoov, ssa.BlockThumbGEnoov:
		jmp := blockJump[b.Kind]
		switch next {
		case b.Succs[0].Block():
			s.Br(jmp.invasm, b.Succs[1].Block())
		case b.Succs[1].Block():
			s.Br(jmp.asm, b.Succs[0].Block())
		default:
			if b.Likely != ssa.BranchUnlikely {
				s.Br(jmp.asm, b.Succs[0].Block())
				s.Br(obj.AJMP, b.Succs[1].Block())
			} else {
				s.Br(jmp.invasm, b.Succs[1].Block())
				s.Br(obj.AJMP, b.Succs[0].Block())
			}
		}

	case ssa.BlockThumbLEnoov:
		s.CombJump(b, next, &leJumps)

	case ssa.BlockThumbGTnoov:
		s.CombJump(b, next, &gtJumps)

	default:
		b.Fatalf("branch not implemented: %s", b.LongString())
	}
}
