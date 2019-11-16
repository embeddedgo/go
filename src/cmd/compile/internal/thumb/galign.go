// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package thumb

import (
	"cmd/compile/internal/gc"
	"cmd/compile/internal/ssa"
	"cmd/internal/obj/thumb"
	"cmd/internal/objabi"
)

func Init(arch *gc.Arch) {
	arch.LinkArch = &thumb.Link
	arch.REGSP = thumb.REGSP
	arch.MAXWIDTH = (1 << 32) - 1
	arch.SoftFloat = objabi.GOARM&0xF != 0xD // TODO: handle GOARM==0x7F (32-bit FPU)
	arch.ZeroRange = zerorange
	arch.ZeroAuto = zeroAuto
	arch.Ginsnop = ginsnop      // used as inline mark
	arch.Ginsnopdefer = ginsnop // for stack trace to show right line number for deffered calls

	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
