// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package thumbasm

import (
	"encoding/binary"
	"fmt"
)

var (
	errShort   = fmt.Errorf("truncated instruction")
	errUnknown = fmt.Errorf("unknown instruction")
	errBad     = fmt.Errorf("instruction")
)

// Decode decodes the leading bytes in src as a single instruction.
func Decode(src []byte) (Inst, error) {
	if len(src) < 2 {
		return Inst{}, errShort
	}
	if src[1]>>3 > 0x1C {
		if len(src) < 4 {
			return Inst{}, errShort
		}
		enc := binary.LittleEndian.Uint32(src)
		x := enc>>16 | enc<<16 // ARM documentation artifact
		for i := range inst32formats {
			f := &inst32formats[i]
			if f.mask&x != f.value {
				continue
			}
			var args Args
			if f.args != nil {
				args = f.args(x)
				if args[0] == nil {
					return Inst{}, errBad
				}
			}
			return Inst{Op: f.op, Enc: enc, Len: 4, Args: args}, nil
		}
		return Inst{}, errUnknown
	}
	x := binary.LittleEndian.Uint16(src)
	for i := range inst16formats {
		f := &inst16formats[i]
		if f.mask&x != f.value {
			continue
		}
		var args Args
		if f.args != nil {
			args = f.args(x)
			if args[0] == nil {
				return Inst{}, errBad
			}
		}
		return Inst{Op: f.op, Enc: uint32(x), Len: 2, Args: args}, nil
	}
	return Inst{}, errUnknown
}

type inst16format struct {
	mask  uint16
	value uint16
	op    Op
	args  func(enc uint16) Args
}

type inst32format struct {
	mask  uint32
	value uint32
	op    Op
	args  func(enc uint32) Args
}

// 16-bit instructions

// 0100 01x0 dmmm mddd
func _ADD__Rm__Rdn(enc uint16) Args {
	return Args{Reg(enc>>4&8 | enc&7), Reg(enc >> 3 & 15)} // Rd, Rm
}

// 0001 10xm mmnn nddd
func _ADD__Rm__Rn__Rd(enc uint16) Args {
	return Args{Reg(enc & 7), Reg(enc >> 3 & 7), Reg(enc >> 6 & 7)}
}

// 0001 11xu uunn nddd
func _ADD__u3__Rn__Rd(enc uint16) Args {
	return Args{Reg(enc & 7), Reg(enc >> 3 & 7), Imm(enc >> 6 & 7)}
}

// 1011 0000 xuuu uuuu
func _ADD__u7_2__R13(enc uint16) Args {
	return Args{R13, Imm(enc & 0x7F << 2)} // R13, u7<<2
}

// 1010 xddd uuuu uuuu
func _ADD__u8_2__R13__Rd(enc uint16) Args {
	Rn := R15
	if enc&0x800 != 0 {
		Rn = R13
	}
	return Args{Reg(enc >> 8 & 7), Rn, Imm(enc & 0xFF << 2)}
}

// 0100 xxxx xxmm mddd
func _AND__Rm__Rdn(enc uint16) Args {
	args := Args{Reg(enc & 7), Reg(enc >> 3 & 7)} // Rdn, Rm
	if enc&0xFFC0 == 0x4240 {
		args[2] = Imm(0)
	}
	return args
}

// 1110 0iii iiii iiii
func _B__i11_1(enc uint16) Args {
	return Args{PCRel(enc) << 21 >> 20}
}

// 0100 0111 xmmm m000
func _B__Rm(enc uint16) Args {
	return Args{Reg(enc >> 3 & 15)} // Rm
}

// 1101 cccc iiii iiii
func _Bcond__i8_1(enc uint16) Args {
	return Args{PCRel(enc) << 24 >> 23}
}

// 1011 x0u1 uuuu unnn
func _CBZ__Rn__u6_1(enc uint16) Args {
	return Args{Reg(enc & 7), PCRel(enc>>4&0x20 | enc>>3&0x1F)}
}

// 1011 0110 011x 00if
func _CPSIE(enc uint16) Args {
	var s string
	switch enc & 3 {
	case 1:
		s = "f"
	case 2:
		s = "i"
	case 3:
		s = "if"
	}
	return Args{Str(s)}
}

// 1011 1111 cccc mmmm
func _ITmask__firstcond(enc uint16) Args {
	return Args{Cond(enc >> 4 & 0xF)}
}

// 1100 xnnn rrrr rrrr
func _MOVM_IAW(enc uint16) Args {
	rlist := RegList(enc & 0xFF)
	Rn := Reg(enc >> 8 & 7)
	mode := AddrLDM_WB
	if enc&0x800 != 0 && 1<<Rn&rlist != 0 {
		mode = AddrLDM
	}
	return Args{Mem{Base: Rn, Mode: mode}, rlist}
}

// 0101 xxxm mmnn nttt
func _MOVW__Rn_Rm__Rt(enc uint16) Args {
	return Args{
		Reg(enc & 7),
		Mem{Base: Reg(enc >> 3 & 7), Mode: AddrOffset, Sign: 1, Index: Reg(enc >> 6 & 7)},
	}
}

// 000v vuuu uumm mddd (vv != 11)
func _MOVW__Rm_v_u5__Rd(enc uint16) Args {
	return Args{Reg(enc & 7), Reg(enc >> 3 & 7), Imm(enc >> 6 & 0x1F)}
}

// xxxx xuuu uunn nttt
func _MOVW__u5_2_Rn__Rt(enc uint16) Args {
	Rt := Reg(enc & 7)
	Rn := Reg(enc >> 3 & 7)
	offset := uint(enc >> 6 & 0x1F)
	switch enc >> 12 {
	case 6:
		offset <<= 2
	case 8:
		offset <<= 1
	}
	return Args{Rt, Mem{Base: Rn, Offset: int16(offset)}}
}

// 001x xddd uuuu uuuu
func _MOVW__u8__Rd(enc uint16) Args {
	return Args{Reg(enc >> 8 & 7), Imm(enc & 0xFF)}
}

// xx0x xttt uuuu uuuu
func _MOVW__u8_2_R13__Rt(enc uint16) Args {
	Rt := Reg(enc >> 8 & 7)
	mem := Mem{Mode: AddrOffset, Offset: int16(enc & 0xFF << 2)}
	switch enc >> 11 {
	case 0x09: // Rt, [R15, u8<<2]
		mem.Base = R15
	case 0x12, 0x13: // Rt, [R13, u8<<2]
		mem.Base = R13
	}
	return Args{Rt, mem}
}

// 1011 x10r rrrr rrrr
func _PUSH__reglist(enc uint16) Args {
	rlist := enc & 0xFF
	lrpc := enc & 0x100
	if enc&0x800 == 0 {
		rlist |= lrpc << 6 // PUSH
	} else {
		rlist |= lrpc << 7 // POP
	}
	return Args{RegList(rlist)}
}

// 1xx1 111x uuuu uuuu
func _UDF__u8(enc uint16) Args {
	return Args{Imm(enc & 0xFF)}
}

// 32-bit instructions

// 1111 0u10 x0x0 nnnn  0uuu dddd uuuu uuuu
func _ADD__u12__Rn__Rd(enc uint32) Args {
	return Args{
		Reg(enc >> 8 & 15),
		Reg(enc >> 16 & 15),
		Imm(enc>>15&0x800 | enc>>4&0x700 | enc&0xFF),
	}
}

// 1111 0e0x xxxs nnnn  0eee dddd eeee eeee   RSB.s     e32, Rn, Rd
func _ANDs__e32__Rn__Rd(enc uint32) Args {
	return Args{Reg(enc >> 8 & 15), Reg(enc >> 16 & 15), decodeMIC(enc)}
}

// 1110 101x xxxs nnnn  0uuu dddd uuvv mmmm
func _ANDs__Rm_v_u5__Rn__Rd(enc uint32) Args {
	a := _TST__Rm_v_u5__Rn(enc)
	return Args{Reg(enc >> 8 & 15), a[0], a[1]}
}

// 1111 0jii iiii iiii  1xj1 jiii iiii iiii
func _B__ji24_1(enc uint32) Args {
	s := enc >> 26 & 1
	imm10 := enc >> 16 & 0x3FF
	i1 := ^(enc>>13 ^ s) & 1
	i2 := ^(enc>>11 ^ s) & 1
	imm11 := enc & 0x7FF
	enc = s<<23 | i1<<22 | i2<<21 | imm10<<11 | imm11
	return Args{PCRel(enc) << 8 >> 7}
}

// 1111 0011 0110 nnnn  0uuu dddd uu0k kkkk
func _BFC__width__ulsb__Rd(enc uint32) Args {
	a := _BFX__width__ulsb__Rn__Rd(enc)
	Rd := a[0]
	Rn := a[1]
	lsb := a[2]
	width := a[3].(Imm) - lsb.(Imm)
	if Rn.(Reg) == 15 {
		return Args{Rd, lsb, width} // BFC
	}
	return Args{Rd, Rn, lsb, width} // BFI
}

// 1111 0011 x100 nnnn  0uuu dddd uu0w wwww
func _BFX__width__ulsb__Rn__Rd(enc uint32) Args {
	width := Imm(enc&0x1F) + 1
	lsb := Imm(enc>>10&0x1C | enc>>6&0x3)
	Rn := Reg(enc >> 16 & 15)
	Rd := Reg(enc >> 8 & 15)
	return Args{Rd, Rn, lsb, width}
}

// 1111 1010 10x1 mmmm  1111 dddd 10xx mmmm
func _CLZ__Rm__Rd(enc uint32) Args {
	Rn := Reg(enc >> 16 & 15)
	Rd := Reg(enc >> 8 & 15)
	Rm := Reg(enc & 15)
	if Rn != Rm {
		return Args{Rd, ArgErr(fmt.Sprintf("{!inconsistent Rm: %v/%v}", Rm, Rn))}
	}
	return Args{Rd, Rm}
}

// 1111 0011 1011 1111  1000 1111 0100 oooo
func _DSB__opt(enc uint32) Args {
	return Args{Imm(enc & 15)}
}

// 1110 1000 0101 nnnn  tttt 1111 uuuu uuuu
func _LDREX__u8_2_Rn__Rt(enc uint32) Args {
	return Args{
		Reg(enc >> 12 & 15),
		Mem{Base: Reg(enc >> 16 & 15), Mode: AddrOffset, Offset: int16(enc & 0xFF << 2)},
	}
}

// 1110 1000 1101 nnnn  tttt 1111 010x 1111
func _LDREXB__Rn__Rt(enc uint32) Args {
	return Args{
		Reg(enc >> 12 & 15),
		Mem{Base: Reg(enc >> 16 & 15), Mode: AddrOffset},
	}
}

// 1111 1010 0x0x 1111  1111 dddd 10rr mmmm
func _MOVH__Rm_rot__Rd(enc uint32) Args {
	return Args{Reg(enc >> 8 & 15), RegShift{Reg(enc & 15), RotateRight, uint8(enc >> 1 & 0x18)}}
}

// 1110 100x x0wx nnnn  rr0r rrrr rrrr rrrr
func _MOVM_IAw(enc uint32) Args {
	mode := AddrLDM
	if enc&0x200000 != 0 {
		mode = AddrLDM_WB
	}
	Rn := Reg(enc >> 16 & 15)
	return Args{Mem{Base: Rn, Mode: mode}, RegList(enc)}
}

// 1111 100x 0xxx nnnn  tttt 0000 00uu mmmm
func _MOVW__Rn_Rm_1_u2__Rt(enc uint32) Args {
	return Args{
		Reg(enc >> 12 & 15),
		Mem{
			Base:  Reg(enc >> 16 & 15),
			Mode:  AddrOffset,
			Sign:  1,
			Index: Reg(enc & 15),
			Shift: ShiftLeft,
			Count: uint8(enc >> 4 & 3),
		},
	}
}

// 1111 100x 1xxx nnnn  tttt uuuu uuuu uuuu
// 1111 100x ±xxx 1111  tttt uuuu uuuu uuuu
func _MOVW__s12_Rn__Rt(enc uint32) Args {
	Rn := Reg(enc >> 16 & 15)
	Rt := Reg(enc >> 12 & 15)
	offset := int16(enc & 0xFFF)
	if enc&0x800000 == 0 {
		offset = -offset
	}
	return Args{Rt, Mem{Base: Rn, Offset: offset}}
}

// 1111 0y10 x100 uuuu  0zzz dddd zzzz zzzz
func _MOVW__uyz16__Rd(enc uint32) Args {
	Rd := Reg(enc >> 8 & 15)
	imm := Imm(enc>>4&0xF000 | enc>>15&0x800 | enc>>4&0x700 | enc&0xFF)
	return Args{Rd, imm}
}

// 1111 100x 0xxx nnnn  tttt 1p±w uuuu uuuu
func _MOVWpw__s8_Rn__Rt(enc uint32) Args {
	Rn := Reg(enc >> 16 & 15)
	Rt := Reg(enc >> 12 & 15)
	offset := int16(enc & 0xFF)
	if enc&0x200 == 0 {
		offset = -offset
	}
	p0w := enc >> 8 & 5
	var mode AddrMode
	switch p0w {
	case 1:
		mode = AddrPostIndex
	case 4:
		mode = AddrOffset
	case 5:
		mode = AddrPreIndex
	default:
		return Args{} // undefined
	}
	return Args{Rt, Mem{Base: Rn, Mode: mode, Offset: offset}}
}

// 1111 0e00 01xs 1111  0eee dddd eeee eeee
func _MOVWs__e32__Rd(enc uint32) Args {
	return Args{Reg(enc >> 8 & 15), decodeMIC(enc)}
}

// 1111 1010 0vvs nnnn  1111 dddd 0000 mmmm
func _MOVWs__Rn_v_Rm__Rd(enc uint32) Args {
	return Args{Reg(enc >> 8 & 15), Reg(enc >> 16 & 15), Reg(enc & 15)}
}

// 1111 0011 1110 1111  1000 dddd mmmm mmmm
// 1111 0011 1000 nnnn  1000 kk00 mmmm mmmm
func _MOVW__SYSm__Rd(enc uint32) Args {
	SYSm := APSR + Reg(enc&0xFF)
	if enc>>20 == 0xF3E {
		// MRS
		return Args{Reg(enc >> 8 & 15), SYSm}
	} else {
		// MSR
		return Args{SYSm, Reg(enc >> 16 & 15)}
	}
}

// 1111 1011 xxxx nnnn  1111 dddd xxxx mmmm
func _MUL__Rm__Rn__Rd(enc uint32) Args {
	return Args{Reg(enc >> 8 & 15), Reg(enc >> 16 & 15), Reg(enc & 15)}
}

// 1111 1011 xxxx nnnn  llll hhhh 000x mmmm
func _MULL__Rm__Rn__Rdh_Rdl(enc uint32) Args {
	return Args{Reg(enc >> 12 & 15), Reg(enc>>8) & 15, Reg(enc >> 16 & 15), Reg(enc & 15)}
}

// 1111 1000 010x 1101 tttt 1xx1 0000 0100
func _PUSH__Rt(enc uint32) Args {
	return Args{RegList(1 << (enc >> 12 & 0xF))}
}

// 1110 1000 1101 nnnn  1111 0000 000h mmmm
func _TBB__Rm__Rn(enc uint32) Args {
	return Args{Reg(enc >> 16 & 15), Reg(enc & 15)}
}

// 1111 0e0x x0x1 nnnn  0eee 1111 eeee eeee
func _TST__e32__Rn(enc uint32) Args {
	return Args{Reg(enc >> 16 & 15), decodeMIC(enc)}
}

// 1110 101x x0x1 nnnn  0uuu 1111 uuvv mmmm
func _TST__Rm_v_u5__Rn(enc uint32) Args {
	return Args{Reg(enc >> 16 & 15), decodeShiftI(enc)}
}

// 1110 1010 01xs 1111  0uuu dddd uuvv mmmm
func _MOVWs__Rm_v_u5__Rn(enc uint32) Args {
	return Args{Reg(enc >> 8 & 15), decodeShiftI(enc)}
}

// 1110 1000 0100 nnnn  tttt dddd uuuu uuuu
func _STREX__Rt__u8_2_Rn__Rd(enc uint32) Args {
	a := _LDREX__u8_2_Rn__Rt(enc)
	return Args{Reg(enc >> 8 & 15), a[0], a[1]}
}

// 1110 1000 1100 nnnn  tttt 1111 010x dddd
func _STREXB__Rn__Rt(enc uint32) Args {
	a := _LDREXB__Rn__Rt(enc)
	return Args{Reg(enc & 15), a[0], a[1]}
}

// 1110 1101 ±d0x nnnn  dddd 101x uuuu uuuu
func _MOVF__s8_2_Rn__Fd(enc uint32) Args {
	var Vd Reg
	if (enc>>8)&1 != 0 {
		Vd = Reg(int(D0) + int((enc>>18)&16|(enc>>12)&15))
	} else {
		Vd = Reg(int(S0) + int((enc>>11)&30|(enc>>22)&1))
	}
	offset := int(enc) & 0xFF
	if (enc>>23)&1 == 0 {
		offset = -offset
	}
	return Args{Vd, Mem{Mode: AddrOffset, Base: Reg(enc >> 16 & 15), Offset: int16(offset)}}
}

// 1110 1110 1d11 0001  dddd 101x 11m0 mmmm
func _SQRTF__Fm__Fd(enc uint32) Args {
	var Vm, Vd Reg
	if (enc>>8)&1 != 0 {
		Vm = Reg(int(D0) + int((enc>>1)&16|enc&15))
		Vd = Reg(int(D0) + int((enc>>18)&16|(enc>>12)&15))
	} else {
		Vm = Reg(int(S0) + int((enc<<1)&30|(enc>>5)&1))
		Vd = Reg(int(S0) + int((enc>>11)&30|(enc>>22)&1))
	}
	return Args{Vd, Vm}
}

func decodeMIC(enc uint32) ImmAlt {
	rot := int(enc>>22&0x10 | enc>>11&0x0E | enc>>7&0x01)
	if rot < 8 {
		enc &= 0xFF
	} else {
		enc = enc&0x7F | 0x80
	}
	return ImmAlt{uint8(enc), uint8(rot)}
}

func decodeShiftI(enc uint32) RegShift {
	Rm := Reg(enc & 15)
	shift := Shift(enc >> 4 & 3)
	count := uint8(enc>>10&0x1C | enc>>6&3)
	if shift == RotateRight && count == 0 {
		shift = RotateRightExt
	}
	return RegShift{Rm, shift, count}
}
