// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package thumb

import (
	"cmd/compile/internal/ssa"
	"cmd/compile/internal/ssagen"
	"cmd/internal/obj/thumb"
	"internal/buildcfg"
)

func Init(arch *ssagen.ArchInfo) {
	arch.LinkArch = &thumb.Link
	arch.REGSP = thumb.REGSP
	arch.MAXWIDTH = (1 << 32) - 1
	arch.SoftFloat = buildcfg.GOARM&0xF != 0xD // TODO: handle GOARM==0x7F (32-bit FPU)
	arch.ZeroRange = zerorange
	arch.Ginsnop = ginsnop      // used as inline mark

	arch.SSAMarkMoves = func(s *ssagen.State, b *ssa.Block) {}
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
