// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package thumb

import (
	"cmd/compile/internal/ir"
	"cmd/compile/internal/objw"
	"cmd/compile/internal/types"
	"cmd/internal/obj"
	"cmd/internal/obj/thumb"
)

func zerorange(pp *objw.Progs, p *obj.Prog, off, cnt int64, r0 *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	if *r0 == 0 {
		p = pp.Append(p, thumb.AMOVW, obj.TYPE_CONST, 0, 0, obj.TYPE_REG, thumb.REG_R0, 0)
		*r0 = 1
	}

	if cnt < int64(4*types.PtrSize) {
		for i := int64(0); i < cnt; i += int64(types.PtrSize) {
			p = pp.Append(p, thumb.AMOVW, obj.TYPE_REG, thumb.REG_R0, 0, obj.TYPE_MEM, thumb.REGSP, 4+off+i)
		}
	} else if cnt <= int64(128*types.PtrSize) {
		p = pp.Append(p, thumb.AADD, obj.TYPE_CONST, 0, 4+off, obj.TYPE_REG, thumb.REG_R1, 0)
		p.Reg = thumb.REGSP
		p = pp.Append(p, obj.ADUFFZERO, obj.TYPE_NONE, 0, 0, obj.TYPE_MEM, 0, 0)
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = ir.Syms.Duffzero
		p.To.Offset = 4 * (128 - cnt/int64(types.PtrSize))
	} else {
		p = pp.Append(p, thumb.AADD, obj.TYPE_CONST, 0, 4+off, obj.TYPE_REG, thumb.REG_R1, 0)
		p.Reg = thumb.REGSP
		p = pp.Append(p, thumb.AADD, obj.TYPE_CONST, 0, cnt, obj.TYPE_REG, thumb.REG_R2, 0)
		p.Reg = thumb.REG_R1
		p = pp.Append(p, thumb.AMOVW, obj.TYPE_REG, thumb.REG_R0, 0, obj.TYPE_MEM, thumb.REG_R1, 4)
		p1 := p
		p.Scond |= thumb.C_PBIT
		p = pp.Append(p, thumb.ACMP, obj.TYPE_REG, thumb.REG_R1, 0, obj.TYPE_NONE, 0, 0)
		p.Reg = thumb.REG_R2
		p = pp.Append(p, thumb.ABNE, obj.TYPE_NONE, 0, 0, obj.TYPE_BRANCH, 0, 0)
		p.To.SetTarget(p1)
	}

	return p
}

func ginsnop(pp *objw.Progs) *obj.Prog {
	return pp.Prog(thumb.ANOP2)
}
