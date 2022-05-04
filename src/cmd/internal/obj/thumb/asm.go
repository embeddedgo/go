// Inferno utils/5l/span.c
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/5l/span.c
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
	"fmt"
	"math"

	"cmd/internal/obj"
	"internal/buildcfg"
)

type Ctx struct {
	ctxt     *obj.Link
	newprog  obj.ProgAlloc
	cursym   *obj.LSym
	blitrl   *obj.Prog
	elitrl   *obj.Prog
	it       *obj.Prog
	beforeit *obj.Prog
	autosize int // frame size: auto + ret
	pool     struct {
		start int
		size  int
	}
}

func regclass(r int16) Aclass {
	if REG_R0 <= r && r <= REG_R7 {
		return C_RLO
	}
	if REG_R0 <= r && r <= REG_R15 {
		switch r {
		case REGSP:
			return C_SP
		case REGPC:
			return C_PC
		}
		return C_REG
	}
	if REG_F0 <= r && r <= REG_F15 {
		return C_FREG
	}
	if REG_APSR <= r && r <= REG_CONTROL {
		return C_SPEC
	}
	if r == REG_FPSCR {
		return C_FCR
	}
	if REG_MB_OSHST <= r && r <= REG_MB_SY {
		return C_MB
	}
	if REG_EQ <= r && r <= REG_AL {
		return C_IT
	}
	return C_GOK
}

func rol(x uint32, k int) uint32 {
	m := uint(k) & 31
	return x<<m | x>>(32-m)
}

func mic(u uint32) bool {
	if u>>8 == 0 {
		return true // 00000000_00000000_00000000_abcdefgh
	}
	for n := 31; n >= 8; n-- {
		if rol(u, n)&^0xFF == 0 {
			return true // rotated 00000000_00000000_00000000_1bcdefgh
		}
	}
	if hi, lo := u>>16, u&0xFFFF; hi == lo {
		hi = lo >> 8
		lo &= 0xFF
		switch {
		case hi == 0:
			return true // 00000000_abcdefgh_00000000_abcdefgh
		case lo == 0:
			return true // abcdefgh_00000000_abcdefgh_00000000
		case hi == lo:
			return true // abcdefgh_abcdefgh_abcdefgh_abcdefgh
		}
	}
	return false
}

func conclass(u uint32) Aclass {
	if u&3 == 0 {
		switch {
		case u == 0:
			return C_ZCON
		case u < 1<<3:
			return C_U1CON2
		case u < 1<<8:
			return C_U6CON2
		case u < 1<<9:
			return C_U7CON2
		case u < 1<<10:
			return C_U8CON2
		}
	} else {
		switch {
		case u < 1<<3:
			return C_U3CON
		case u < 1<<8:
			return C_U8CON
		}
	}
	for n := uint(1); n <= 8; n++ {
		if u&(1<<n-1) == 0 && u>>n <= 0xFF {
			if n <= 4 {
				return C_U8CON1_4
			}
			return C_U8CON5_8
		}
	}
	switch {
	case u < 1<<12:
		return C_U12CON
	case u < 1<<16:
		return C_U16CON
	case mic(u):
		return C_E32CON
	}
	return C_LCON
}

func (c *Ctx) offset(a *obj.Addr) int64 {
	switch a.Name {
	case obj.NAME_PARAM:
		return a.Offset + int64(c.autosize+prasize)
	case obj.NAME_AUTO:
		return a.Offset + int64(c.autosize)
	}
	return a.Offset
}

func (c *Ctx) chipzero(e float64) int {
	switch buildcfg.GOARM {
	case 0x7F, 0x7D: // ARMv7-M with floating point extension
		break
	default:
		return -1
	}
	if math.Float64bits(e) == 0 {
		return 0
	}
	return -1
}

func (c *Ctx) chipfloat(e float64) int {
	switch buildcfg.GOARM {
	case 0x7F, 0x7D: // ARMv7-M with floating point extension
		break
	default:
		return -1
	}
	ei := math.Float64bits(e)
	l := uint32(ei)
	h := uint32(ei >> 32)

	if l != 0 || h&0xffff != 0 {
		return -1
	}
	h1 := h & 0x7fc00000
	if h1 != 0x40000000 && h1 != 0x3fc00000 {
		return -1
	}
	n := 0

	// sign bit (a)
	if h&0x80000000 != 0 {
		n |= 1 << 7
	}

	// exp sign bit (b)
	if h1 == 0x3fc00000 {
		n |= 1 << 6
	}

	// rest of exp and mantissa (cd-efgh)
	n |= int((h >> 16) & 0x3f)

	//print("match %.8lux %.8lux %d\n", l, h, n);
	return n
}

func (c *Ctx) aclass(a *obj.Addr, as *obj.As) Aclass {
	switch a.Type {
	case obj.TYPE_NONE:
		return C_NONE

	case obj.TYPE_REG:
		return regclass(a.Reg)

	case obj.TYPE_REGREG:
		return C_REGREG

	case obj.TYPE_REGREG2:
		return C_REGREG2

	case obj.TYPE_REGLIST:
		if a.Offset&^0xC0FF == 0 {
			switch a.Offset >> 12 {
			case 0:
				return C_LISTLO
			case 4:
				return C_LISTLOLR
			case 8:
				return C_LISTLOPC
			}
		}
		return C_LIST

	case obj.TYPE_SHIFT:
		if a.Reg == 0 {
			rlo := a.Offset&15 <= 7
			if a.Offset>>4&1 == 0 {
				// R>>i
				if rlo && a.Offset>>5&3 != 3 {
					// rlo and not @>
					return C_SHIFTILO
				}
				return C_SHIFTI
			} else {
				// R>>R
				if rlo && a.Offset>>8&15 <= 7 {
					return C_SHIFTRLO
				}
				return C_SHIFTR
			}
		}

	case obj.TYPE_MEM:
		switch a.Name {
		case obj.NAME_STATIC, obj.NAME_EXTERN:
			if a.Sym == nil || a.Sym.Name == "" {
				return C_GOK
			}
			return C_LORLO // MOVW litoffset(PC), R7 and use (R7) as address

		case obj.NAME_NONE, obj.NAME_AUTO, obj.NAME_PARAM:
			if a.Index != 0 {
				if a.Reg <= REG_R7 && a.Index <= REG_R7 && uint(a.Scale) <= 1 {
					return C_RORLO
				}
				return C_ROREG
			}
			if a.Name == obj.NAME_AUTO || a.Name == obj.NAME_PARAM {
				a.Reg = REGSP
			}
			offset := c.offset(a)
			switch {
			case offset > 0:
				switch {
				case a.Reg <= REG_R7:
					switch {
					case offset&3 != 0:
						break
					case offset < 1<<5:
						return C_U3ORLO2
					case offset < 1<<6:
						return C_U4ORLO2
					case offset < 1<<7:
						return C_U5ORLO2
					}
					switch {
					case offset&1 != 0:
						break
					case offset < 1<<5:
						return C_U4ORLO1
					case offset < 1<<6:
						return C_U5ORLO1
					}
					if offset < 1<<5 {
						return C_U5ORLO
					}
				case a.Reg == REGSP:
					switch {
					case offset&3 != 0:
						break
					case offset < 1<<8:
						return C_U6OSP2
					case offset < 1<<10:
						return C_U8OSP2
					}
				case a.Reg == REGPC:
					switch {
					case offset&3 != 0:
						break
					case offset < 1<<8:
						return C_U6OPC2
					case offset < 1<<10:
						return C_U8OPC2
					}
					if offset < 1<<12 {
						return C_U12OPC
					}
					return C_LOREG // there is no wider class than U12
				}
				switch {
				case offset&3 != 0:
					break
				case offset < 1<<8:
					return C_U6OREG2
				case offset < 1<<10:
					return C_U8OREG2
				}
				switch {
				case offset < 1<<8:
					return C_U8OREG
				case offset < 1<<12:
					return C_U12OREG
				}
			case offset == 0:
				switch {
				case a.Reg <= REG_R7:
					return C_U0ORLO
				case a.Reg == REG_R13:
					return C_U0OSP
				case a.Reg == REG_R15:
					return C_U0OPC
				}
				return C_U0OREG
			default: // offset < 0
				offset = -offset
				if a.Reg == REGPC {
					switch {
					case offset&3 != 0:
						break
					case offset < 1<<10:
						return C_S8OPC2
					}
					switch {
					case offset < 1<<8:
						return C_S8OPC
					case offset < 1<<12:
						return C_S12OPC
					}
					return C_LOREG // there is no wider class than S12
				}
				switch {
				case offset&3 != 0:
					break
				case offset < 1<<8:
					return C_S6OREG2
				case offset < 1<<10:
					return C_S8OREG2
				}
				if offset < 1<<8 {
					return C_S8OREG
				}
			}
			if a.Reg <= REG_R7 {
				return C_LORLO
			}
			return C_LOREG
		}

	case obj.TYPE_FCONST:
		if c.chipzero(a.Val.(float64)) >= 0 {
			return C_ZFCON
		}
		if c.chipfloat(a.Val.(float64)) >= 0 {
			return C_SFCON
		}
		return C_LFCON

	case obj.TYPE_TEXTSIZE:
		return C_TEXTSIZE

	case obj.TYPE_CONST:
		if a.Offset > 4294967295 || a.Offset < -2147483648 {
			return C_GOK // doesn't fit in uint32 or int32
		}
		con := uint32(a.Offset)
		cc := conclass(con)
		if as != nil && cc >= C_E32CON {
			switch *as {
			case AADD, ASUB, ACMP, ACMN:
				con = -con
				nc := conclass(con)
				if nc >= cc {
					break
				}
				cc = nc
				a.Offset = int64(con)
				switch *as {
				case AADD:
					*as = ASUB
				case ASUB:
					*as = AADD
				case ACMP:
					*as = ACMN
				default: // ACMN
					*as = ACMP
				}
			case AAND, ABIC, AORR, AORN, AMOVW, AMVN:
				con = ^con
				nc := conclass(con)
				if nc >= cc {
					break
				}
				cc = nc
				a.Offset = int64(con)
				switch *as {
				case AAND:
					*as = ABIC
				case ABIC:
					*as = AAND
				case AORR:
					*as = AORN
				case AORN:
					*as = AORR
				case AMOVW:
					*as = AMVN
				default: // AMVN
					*as = AMOVW
				}
			}
		}
		return cc

	case obj.TYPE_ADDR:
		switch a.Name {
		case obj.NAME_STATIC, obj.NAME_EXTERN:
			if a.Sym == nil || a.Sym.Name == "" {
				return C_GOK
			}
			return C_LCON
		}
		// MOV with NAME_NONE, NAME_PARAM, NAME_AUTO is converted to ADD with TYPE_CONST

	case obj.TYPE_BRANCH:
		// can't determine branch length here so assume some tentative value
		switch *as {
		case AB, ABL:
			return C_S24BRA
		case ACBZ, ACBNZ:
			return C_U6BRA
		default: // Bcond
			return C_S20BRA
		}
	}

	return C_GOK
}

func itstate(p *obj.Prog) uint16 {
	mask := int(p.As-AITTTT) + 1
	// IT      mask = 0b1000
	// IT T    mask = 0b0100
	// IT E    mask = 0b1100
	// IT TT   mask = 0b0010
	// IT ET   mask = 0b1010
	// IT TE   mask = 0b0110
	// IT EE   mask = 0b1110
	// IT TTT  mask = 0b0001
	// IT ETT  mask = 0b1001
	// IT TET  mask = 0b0101
	// IT EET  mask = 0b1101
	// IT TTE  mask = 0b0011
	// IT ETE  mask = 0b1011
	// IT TEE  mask = 0b0111
	// IT EEE  mask = 0b1111
	if fc0 := int(p.Scond) & 1; fc0 != 0 {
		n := uint(0)
		for mask>>n&1 == 0 {
			n++ // count trailing zeros
		}
		for n++; n < 4; n++ {
			mask ^= fc0 << n
		}
	}
	return uint16(mask)
}

func (c *Ctx) itclose() {
	c.it.As = AITTTT + obj.As(c.it.Mark) - 1
	c.it.Mark = uint16(c.it.Scond<<4) | itstate(c.it)
	c.it = nil
}

func (c *Ctx) oplook(p *obj.Prog) (ret *Optab) {
	if k := int(p.Optab) - 1; k >= 0 {
		return &oprange[p.As&obj.AMask][k]
	}
	a1 := Aclass(p.From.Class)
	if a1 == 0 {
		a1 = c.aclass(&p.From, &p.As) + 1
		p.From.Class = int8(a1)
	}
	a1--
	a3 := Aclass(p.To.Class)
	if a3 == 0 {
		a3 = c.aclass(&p.To, &p.As)
		if a1 == C_SHIFTRLO && p.As == AMOVW && int(p.From.Offset)&7 == int(p.To.Reg)&7 {
			switch a3 {
			case C_RLO:
				a3 = C_NONE // special case to generate 16bit `MOVW Rdn<v>Rm, Rdn`
			case C_NONE:
				a3 = C_GOK
			}
		}
		a3++
		p.To.Class = int8(a3)
	}
	a3--
	a2 := C_NONE
	if p.Reg != 0 {
		a2 = regclass(p.Reg)
	}

	//fmt.Printf("\t%-12v %-12v %-12v %04b\n", a1, a2, a3, p.Scond>>4)

	autoit := c.cursym.Func().Text.Mark&AUTOIT != 0
	cond := p.Scond & 0xF
again:
	if autoit {
		if c.it != nil {
			if (cond^C_SCOND_XOR)>>1 != c.it.Scond>>1 {
				// no condition or condition does not match the current it block
				c.itclose()
			}
		}
		if c.it == nil {
			if cond != C_SCOND_NONE {
				if p.As == AB && p.To.Type == obj.TYPE_BRANCH && p.To.Name == obj.NAME_NONE {
					// convert `B.cond label` to `Bcond label`
					switch cond {
					case C_SCOND_EQ:
						p.As = ABEQ
					case C_SCOND_NE:
						p.As = ABNE
					case C_SCOND_HS:
						p.As = ABHS
					case C_SCOND_LO:
						p.As = ABLO
					case C_SCOND_MI:
						p.As = ABMI
					case C_SCOND_PL:
						p.As = ABPL
					case C_SCOND_VS:
						p.As = ABVS
					case C_SCOND_VC:
						p.As = ABVC
					case C_SCOND_HI:
						p.As = ABHI
					case C_SCOND_LS:
						p.As = ABLS
					case C_SCOND_GE:
						p.As = ABGE
					case C_SCOND_LT:
						p.As = ABLT
					case C_SCOND_GT:
						p.As = ABGT
					case C_SCOND_LE:
						p.As = ABLE
					default:
						c.ctxt.Diag("invalid flags: %v", p)
					}
					p.Scond &^= 0xF
					a3 = C_S20BRA
					p.To.Class = int8(C_S20BRA + 1)
				} else {
					// insert IT by replacing p to avoid jumping in middle of IT block
					// BUG: this only works for one instruction IT blocks (jump target check is need)
					q := c.newprog()
					*q = *p
					p.Link = q
					p.As = AIT
					p.Spadj = 0
					p.Scond = cond ^ C_SCOND_XOR
					p.Mark = 0x10 // invalid (empty IT block)
					p.Reg = 4     // number of free slots in IT block
					p.Optab = 1
					c.it = p
					return &oprange[AIT&obj.AMask][0]
				}
			}
		}
	} else {
		if c.it != nil {
			switch {
			case AITTTT <= p.As && p.As <= AITEEE:
				c.ctxt.Diag("IT in previous IT block: %v", p)
			case ABEQ <= p.As && p.As <= ACBNZ:
				c.ctxt.Diag("not allowed in IT block: %v", p)
			case c.it.RegTo2<<1&0xF != 0 && (p.As == AB || p.As == ABL):
				// BUG: take into account other cases when PC is modified
				c.ctxt.Diag("only allowed as the last instruction in IT block: %v", p)
			case int16(cond^C_SCOND_XOR) != c.it.RegTo2>>4:
				c.ctxt.Diag("condition code does not match IT code: %v", p)
			}
		} else if cond != 0 {
			c.ctxt.Diag("condition code outisde IT block: %v", p)
		}
	}
	spuw := p.Scond &^ 0xF
	var spxor uint8
	if spuw&(C_SBIT|C_PBIT) == 0 {
		spxor = C_SBIT
	} else if c.it != nil {
		spxor = C_SBIT | C_PBIT
	}
	ops := oprange[p.As&obj.AMask]
	for k := range ops {
		op := &ops[k]
		if op.flag&NOIT != 0 && c.it != nil {
			continue
		}
		rscond := op.rscond
		oscond := op.oscond
		if rscond&C_SBIT != 0 {
			// 16-bit data processing instruction that sets flags outside IT block
			rscond ^= spxor
			oscond ^= spxor
		}
		if spuw&rscond == rscond && spuw&oscond == spuw &&
			(op.a2 == a2 || op.a2 == C_REG && (C_RLO <= a2 && a2 < C_REG)) &&
			match(op.a1, a1) && match(op.a3, a3) {

			p.Optab = uint16(k + 1)
			/*fmt.Printf(
				"\t%-12v %-12v %-12v %04b %04b\n",
				op.a1, op.a2, op.a3, rscond>>4, oscond>>4,
			)*/
			ret = op
			break
		}
	}
	if ret == nil {
		c.ctxt.Diag(
			"illegal combination %v; %v %v %v; from %v %v; to %v %v",
			p, a1, a2, a3, p.From.Type, p.From.Name, p.To.Type, p.To.Name,
		)
		return &Optab{as: obj.AUNDEF}
	}
	if autoit {
		if c.it != nil {
			n := int(ret.size) >> 4
			if n > int(c.it.Reg) {
				// can not fit n instructions in current it block
				c.itclose()
				p.Optab = 0
				ret = nil
				goto again
			}
			cond ^= C_SCOND_XOR
			for ; n > 0; n-- {
				c.it.Reg--
				ito := uint16(1 << uint(c.it.Reg))
				if cond&1 == c.it.Scond&1 {
					c.it.Mark -= ito
				} else {
					c.it.Mark += ito
				}
			}
			if c.it.Reg == 0 || p.As == AB || p.As == ABL {
				c.itclose()
			}
		}
	} else {
		if AITTTT <= p.As && p.As <= AITEEE {
			p.Scond = uint8(p.From.Reg - REG_EQ)
			p.Mark = itstate(p)
			p.RegTo2 = int16(int(p.Scond)<<4 | int(p.Mark))
			c.it = p
		} else if c.it != nil {
			c.it.RegTo2 = c.it.RegTo2&^0x1F | c.it.RegTo2<<(ret.size>>4)&0x1F
			if c.it.RegTo2&0xF == 0 {
				c.it = nil
			}
		}
	}
	return ret
}

func (c *Ctx) brlook(p *obj.Prog) *Optab {
	v := (p.To.Target().Pc - p.Pc - 4) >> 1
	var a3 Aclass
	switch {
	case 0 <= v && v < 1<<6:
		a3 = C_U6BRA
	case -1<<7 <= v && v < 1<<7:
		a3 = C_S8BRA
	case -1<<10 <= v && v < 1<<10:
		a3 = C_S11BRA
	case -1<<19 <= v && v < 1<<19:
		a3 = C_S20BRA
	case -1<<23 <= v && v < 1<<23:
		a3 = C_S24BRA
	default:
		c.ctxt.Diag("branch too long: %d %v", v*2, p)
		return &Optab{as: obj.AUNDEF}
	}
	p.To.Class = int8(a3) + 1
	ops := oprange[p.As&obj.AMask]
	for k := range ops {
		op := &ops[k]
		if a3 <= op.a3 && op.a3 <= C_S24BRA {
			p.Optab = uint16(k + 1)
			return op
		}
	}
	c.ctxt.Diag("illegal branch combination %v; %v; to %v %v", p, a3, p.To.Type, p.To.Name)
	return &Optab{as: obj.AUNDEF}
}

func debug(p *obj.Prog) {
	fmt.Printf("\n%v\n", p)
	if p.Reg != 0 {
		fmt.Printf(
			"\t%02d %-5v %v(%v), R%d, %v(%v)\n\n",
			p.Pc, p.As, p.From.Type, p.From.Name, p.Reg-obj.RBaseThumb, p.To.Type, p.To.Name,
		)
	} else {
		fmt.Printf(
			"\t%02d %-5v %v(%v), %v(%v)\n\n",
			p.Pc, p.As, p.From.Type, p.From.Name, p.To.Type, p.To.Name,
		)
	}
}

func (c *Ctx) litoffset(p *obj.Prog, reglo bool, pcdiff int) (offset int, short bool) {
	v := int(p.Pool.Pc - p.Pc&^3 - 4)
	if v > 0 {
		v += pcdiff // forward literal
	}
	if v <= -1<<12 || 1<<12 <= v {
		c.ctxt.Diag("|literal offset| >= 1<<12: %d %v", v, p)
		return 0, false
	}
	return v, reglo && uint(v) < 1<<10
}

func span(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	if ctxt.Retpoline {
		ctxt.Diag("-spectre=ret not supported on thumb")
		ctxt.Retpoline = false // don't keep printing
	}
	p := cursym.Func().Text
	if p == nil || p.Link == nil {
		return // external functions or ELF section symbol
	}
	if oprange[AAND&obj.AMask] == nil {
		ctxt.Diag("thumb ops not initialized, call thumb.buildop first")
	}
	c := &Ctx{ctxt: ctxt, newprog: newprog, cursym: cursym, autosize: int(p.To.Offset) + rasize}
	p.Pc = 0
	pc := 0
	dbg := false
	if false && p.From.Sym.Name == "\"\".main" {
		dbg = true
		fmt.Println("-->", p)
	}
	for p.Link != nil {
		prev := p
		p = p.Link
		if p.As == AWORD && pc&3 != 0 {
			pc += 2
		}
		p.Pc = int64(pc)
		if dbg {
			debug(p)
		}
		o := c.oplook(p)
		if o.as == AIT {
			c.beforeit = prev
		}
		switch o.flag & (LFROM | LTO) {
		case LFROM:
			c.addpool(p, &p.From, pc)
		case LTO:
			c.addpool(p, &p.To, pc)
		}
		pc += int(o.size) & 0xF
		if c.blitrl != nil {
			if p.Link == nil {
				p, pc = c.flushpool(p, pc, false)
			} else {
				p, pc = c.checkpool(p, pc)
				if dbg && p != prev.Link {
					if c.it == nil {
						fmt.Printf("\tpool flushed\n")
					} else {
						fmt.Printf("\tpool flushed before last IT - restarting\n")
					}
				}
			}
		}
	}

	// determine exect instruction sizes
	// TODO: jump across functions needs reloc and always uses 32-bit instruction

	for {
		p = c.cursym.Func().Text
		pc = 0
		changed := false
		for p.Link != nil {
			p = p.Link
			if p.As == AWORD && pc&3 != 0 {
				pc += 2
			}
			pcdiff := int(int64(pc) - p.Pc)
			if pcdiff != 0 {
				changed = true
				p.Pc = int64(pc)
			}
			var o *Optab
			var m int
			if p.To.Type == obj.TYPE_BRANCH && p.To.Target() != nil && p.To.Sym == nil {
				o = c.brlook(p)
				m = int(o.size) & 0xF
			} else {
				o = c.oplook(p)
				m = int(o.size) & 0xF
				if o.flag&(LFROM|LTO) != 0 {
					// in most cases the literal is loaded into R7 (REGTMP)
					loreg := true
					// but in the following cases the literal is loaded straight into destination register
					if p.As == AMOVW || p.As == AMVN {
						if p.From.Type == obj.TYPE_CONST || p.From.Type == obj.TYPE_ADDR &&
							(p.From.Name == obj.NAME_STATIC || p.From.Name == obj.NAME_EXTERN) {
							loreg = p.To.Reg <= REG_R7
						}
					}
					if _, short := c.litoffset(p, loreg, pcdiff); short {
						// optab assumees 32-bit load literal instruction;
						// here we are correcting this assumption
						m -= 2
						//fmt.Println(-2, loreg)
					}
					if p.As == AMOVF || p.As == AMOVD {
						name := p.From.Name
						if o.flag&LTO != 0 {
							name = p.To.Name
						}
						if name == obj.NAME_STATIC || name == obj.NAME_EXTERN {
							// optab assumes the offset(regBase) requires
							// 16-bit ADD regBase,REGTMP instruction but
							// in this case the offset is absolute addres so
							// there is no ADD instruction
							m -= 2
						}
					}
				}
			}
			//fmt.Println(m, p)
			pc += m
		}
		if !changed {
			break
		}
	}
	c.cursym.Size = int64(pc)

	// lay out the code

	c.cursym.Grow(c.cursym.Size)
	bp := c.cursym.P
	var out [8]uint16
	p = c.cursym.Func().Text
	pc = int(p.Pc)
	for p.Link != nil {
		p = p.Link
		var o *Optab
		if p.To.Type == obj.TYPE_BRANCH && p.To.Target() != nil {
			o = c.brlook(p)
		} else {
			o = c.oplook(p)
		}
		if o.asmout == nil {
			continue
		}
		m := o.asmout(c, p, out[:])
		if m == 0 {
			return
		}
		if int64(pc) > p.Pc {
			ctxt.Diag("PC padding invalid: want %#d, has %#d: %v", p.Pc, pc, p)
			return
		}
		for int64(pc) != p.Pc {
			// emit NOP
			bp[1] = 0xBF
			bp[0] = 0x00
			bp = bp[2:]
			pc += 2
		}
		for _, v := range out[:m/2] {
			bp[0] = byte(v)
			bp[1] = byte(v >> 8)
			bp = bp[2:]
		}
		pc += m
		// unset base register in case of SP/FP for better printing
		var a *obj.Addr
		if p.From.Type == obj.TYPE_MEM || p.From.Type == obj.TYPE_ADDR {
			a = &p.From
		} else if p.To.Type == obj.TYPE_MEM || p.To.Type == obj.TYPE_ADDR {
			a = &p.To
		} else {
			continue
		}
		if a.Reg == REGSP && (a.Name == obj.NAME_AUTO || a.Name == obj.NAME_PARAM) {
			a.Reg = obj.REG_NONE
		}
	}
}

func (c *Ctx) addpool(p *obj.Prog, a *obj.Addr, pc int) {
	t := c.newprog()
	t.As = AWORD
	switch a.Name {
	case obj.NAME_EXTERN, obj.NAME_STATIC:
		t.To.Type = a.Type
		t.To.Name = a.Name
		t.To.Sym = a.Sym
		t.To.Offset = a.Offset
	default:
		t.To.Type = obj.TYPE_CONST
		lit := c.offset(a)
		if p.As == AMVN {
			lit = ^lit // MVN lit, Rd is implemented as MOVW litoffset(PC), Rd
		}
		t.To.Offset = lit
	}
	for q := c.blitrl; q != nil; q = q.Link {
		if q.Rel == nil && q.To == t.To {
			p.Pool = q
			return
		}
	}
	if c.blitrl == nil {
		c.blitrl = t
		c.pool.start = pc
	} else {
		c.elitrl.Link = t
	}
	c.elitrl = t
	c.pool.size += 4
	p.Pool = t // store the link to the pool entry in Pool
}

func (c *Ctx) flushpool(p *obj.Prog, pc int, skip bool) (*obj.Prog, int) {
	if skip {
		q := c.newprog()
		q.As = AB
		q.To.Type = obj.TYPE_BRANCH
		q.To.SetTarget(p.Link)
		q.Pos = p.Pos
		q.Link = c.blitrl
		c.blitrl = q
	}
	if pc&3 != 0 {
		pc += 2
	}
	for q := c.blitrl; q != nil; q = q.Link {
		q.Pos = p.Pos // the line number of the preceding instruction (no deltas in the pcln)
		q.Pc = int64(pc)
		pc += 4
	}
	c.pool.size = 0
	c.pool.start = 0
	c.elitrl.Link = p.Link
	p.Link = c.blitrl
	c.blitrl = nil // BUG: should refer back to values until out-of-range
	return c.elitrl, pc
}

func (c *Ctx) checkpool(p *obj.Prog, pc int) (*obj.Prog, int) {
	poolLast := pc
	skip := !(p.As == AB && p.Scond&0xF == 0)
	if skip {
		poolLast += 4 // the AB instruction to jump around the pool
	}
	poolLast += c.pool.size - 4 // the offset of the last pool entry
	refpc := c.pool.start       // PC of the first pool reference
	v := poolLast - refpc - 4   // PC-relative offset
	if v >= 0xFF0 {
		goto flush
	}
	if !skip && v >= 0xFD0 {
		if q := p.Link.Link; q != nil && q.Link != nil {
			goto flush
		}
	}
	return p, pc
flush:
	if c.it == nil {
		return c.flushpool(p, pc, skip)
	}
	// cannot flush literal pool inside IT block
	q := c.beforeit
	skip = !(q.As == AB && q.Scond&0xF == 0)
	return c.flushpool(q, int(c.it.Pc), skip)
}

var xcmporeg = [...][14]byte{
	//B H  W  W  W  U  S  S  F  Z  Z  Z  L  L
	//O O  O  O  O  O  O  O  O  O  O  O  O  O
	//R R  R  S  P  R  R  P  R  R  S  R  R  R
	//L L  L  P  C  E  E  C  E  L  P  E  E  L
	//O O  O        G  G     G  O     G  G  O

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1}, // C_LORLO
	{1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1}, // C_U0ORLO
	{1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1}, // C_U3ORLO2
	{0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1}, // C_U4ORLO2
	{0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1}, // C_U5ORLO2
	{1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1}, // C_U4ORLO1
	{0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1}, // C_U5ORLO1
	{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1}, // C_U5ORLO
	{0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0}, // C_U0OPC
	{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0}, // C_U6OPC2
	{0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0}, // C_U8OPC2
	{0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0}, // C_U12OPC
	{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0}, // C_S8OPC
	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0}, // C_S8OPC2
	{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0}, // C_S12OPC
	{0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0}, // C_U0OSP
	{0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0}, // C_U6OSP2
	{0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0}, // C_U8OSP2
	{0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0}, // C_U0OREG
	{0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0}, // C_U6OREG2
	{0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0}, // C_U8OREG2
	{0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0}, // C_U8OREG
	{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0}, // C_U12OREG
	{0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0}, // C_S6OREG2
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0}, // C_S8OREG2
	{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0}, // C_S8OREG
}

var xcmpcon = [...][8]byte{
	//U U  U  U  U  U  E  L
	//7 8  3  8  1  1  3
	//<<<<       2  6  2
	//2 2

	{1, 1, 1, 1, 1, 1, 1, 1}, // C_ZCON
	{1, 1, 1, 1, 1, 1, 1, 1}, // C_U1CON2
	{1, 1, 0, 1, 1, 1, 1, 1}, // C_U6CON2
	{0, 0, 0, 0, 1, 1, 1, 1}, // C_U8CON1_4 // not contain C_U8CON2
	{0, 0, 0, 0, 0, 1, 1, 1}, // C_U8CON5_8
	{1, 1, 0, 0, 1, 1, 1, 1}, // C_U7CON2
	{0, 1, 0, 0, 1, 1, 1, 1}, // C_U8CON2
	{0, 0, 1, 1, 1, 1, 1, 1}, // C_U3CON
	{0, 0, 0, 1, 1, 1, 1, 1}, // C_U8CON
	{0, 0, 0, 0, 1, 1, 0, 1}, // C_U12CON
	{0, 0, 0, 0, 0, 1, 0, 1}, // C_U16CON
	{0, 0, 0, 0, 0, 0, 1, 1}, // C_E32CON
}

var xcmpreg = [...][6]byte{
	//R S  P  R  S  S
	//L P  C  E  I  I
	//O       G  L

	{1, 0, 0, 1, 1, 1}, // C_RLO
	{0, 1, 0, 1, 0, 1}, // C_SP
	{0, 0, 1, 1, 0, 1}, // C_PC
	{0, 0, 0, 1, 0, 1}, // C_REG
	{0, 0, 0, 0, 1, 1}, // C_SIFTILO
}

func match(op, code Aclass) bool {
	if op == code {
		return true
	}
	switch {
	case C_RLO <= code && code <= C_SHIFTILO:
		if op < C_RLO || C_SHIFTI < op {
			return false
		}
		return xcmpreg[code-C_RLO][op-C_RLO] != 0
	case C_LORLO <= code && code <= C_S8OREG:
		if op < C_BORLO || C_LORLO < op {
			return false
		}
		return xcmporeg[code-C_LORLO][op-C_BORLO] != 0
	case C_ZCON <= code && code <= C_E32CON:
		if op < C_U7CON2 || C_LCON < op {
			return false
		}
		return xcmpcon[code-C_ZCON][op-C_U7CON2] != 0
	case op == C_ROREG && code == C_RORLO:
		return true
	case op == C_SHIFTR && code == C_SHIFTRLO:
		return true
	case op == C_LIST && (C_LISTLO <= code && code <= C_LISTLOPC):
		return true
	case (op == C_LISTLOLR || op == C_LISTLOPC) && code == C_LISTLO:
		return true
	}
	return false
}

var oprange [ALAST & obj.AMask][]Optab

var (
	symdiv      *obj.LSym
	symdivu     *obj.LSym
	symmod      *obj.LSym
	symmodu     *obj.LSym
)

func buildop(ctxt *obj.Link) {
	if oprange[AAND&obj.AMask] != nil {
		// Already initialized; stop now. This happens in the
		// cmd/asm tests, each of which re-initializes the arch.
		return
	}

	symdiv = ctxt.Lookup("runtime._div")
	symdivu = ctxt.Lookup("runtime._divu")
	symmod = ctxt.Lookup("runtime._mod")
	symmodu = ctxt.Lookup("runtime._modu")

	for i := range optab {
		oi := &optab[i]
		if !(oi.size != 0 || oi.a1 != 0 || oi.a2 != 0 || oi.a3 != 0) {
			continue
		}
		oi.oscond |= oi.rscond
	}

	for i := 0; i < len(optab); {
		r := optab[i].as
		if r == obj.AXXX {
			continue
		}
		r0 := r & obj.AMask
		start := i
		for i++; i < len(optab) && optab[i].as == r; i++ {
		}
		opt := optab[start:i]
		oprange[r0] = append(oprange[r0], opt...)
		for i < len(optab) {
			oi := &optab[i]
			if oi.size != 0 || oi.a1 != 0 || oi.a2 != 0 || oi.a3 != 0 || oi.flag != 0 {
				break
			}
			op := &oprange[oi.as&obj.AMask]
			*op = append(*op, opt...)
			i++
		}
	}
}
