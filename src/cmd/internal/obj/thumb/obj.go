// Derived from Inferno utils/5c/swt.c
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/5c/swt.c
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
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package thumb

import (
	"cmd/internal/obj"
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"log"
)

const (
	LEAF = 1 << iota
	AUTOIT
)

// return address size (on stack frame)
const (
	prasize = 4 // parent function
	rasize  = 4 // current function
)

func preprocess(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	if cursym.Func().Text == nil || cursym.Func().Text.Link == nil {
		return
	}

	c := Ctx{ctxt: ctxt, cursym: cursym, newprog: newprog}

	p := c.cursym.Func().Text
	autoffset := int(p.To.Offset)
	if autoffset == -4 {
		// Historical way to mark NOFRAME.
		p.From.Sym.Set(obj.AttrNoFrame, true)
		autoffset = 0
	}
	if autoffset < 0 || autoffset%4 != 0 {
		c.ctxt.Diag("frame size %d not 0 or a positive multiple of 4", autoffset)
	}
	if p.From.Sym.NoFrame() {
		if autoffset != 0 {
			c.ctxt.Diag("NOFRAME functions must have a frame size of 0, not %d", autoffset)
		}
	}

	cursym.Func().Locals = int32(autoffset)
	cursym.Func().Args = p.To.Val.(int32)
	cursym.Func().Text.Mark = AUTOIT
	if !cursym.Func().Text.From.Sym.ISR() {
		cursym.Func().Text.Mark |= LEAF
	}

	for p := cursym.Func().Text; p != nil; p = p.Link {
		switch {
		case AITTTT <= p.As && p.As <= AITEEE:
			cursym.Func().Text.Mark &^= AUTOIT
		case p.As == ABL:
			cursym.Func().Text.Mark &^= LEAF
		}
	}

	var autosize int
	for p := cursym.Func().Text; p != nil; p = p.Link {
		if cursym.Func().Text.Mark&AUTOIT != 0 && p.To.Type == obj.TYPE_BRANCH &&
			ABEQ <= p.As && p.As <= ABLE {
			// convert `Bcond label` to `B.cond label`
			cond := byte(C_SCOND_LE)
			switch p.As {
			case ABEQ:
				cond = C_SCOND_EQ
			case ABNE:
				cond = C_SCOND_NE
			case ABCS, ABHS:
				cond = C_SCOND_HS
			case ABCC, ABLO:
				cond = C_SCOND_LO
			case ABMI:
				cond = C_SCOND_MI
			case ABPL:
				cond = C_SCOND_PL
			case ABVS:
				cond = C_SCOND_VS
			case ABVC:
				cond = C_SCOND_VC
			case ABHI:
				cond = C_SCOND_HI
			case ABLS:
				cond = C_SCOND_LS
			case ABGE:
				cond = C_SCOND_GE
			case ABLT:
				cond = C_SCOND_LT
			case ABGT:
				cond = C_SCOND_GT
			}
			p.As = AB
			p.Scond = p.Scond&^C_SCOND | cond
		}

		switch p.As {
		case obj.ATEXT:
			autosize = autoffset

			if p.Mark&LEAF != 0 && autosize == 0 {
				// A leaf function with no locals has no frame.
				p.From.Sym.Set(obj.AttrNoFrame, true)
			}

			if !p.From.Sym.NoFrame() {
				// If there is a stack frame at all, it includes
				// space to save the LR.
				autosize += rasize
			}

			if autosize == 0 && cursym.Func().Text.Mark&LEAF == 0 {
				// A very few functions that do not return to their caller
				// are not identified as leaves but still have no frame.
				if ctxt.Debugvlog {
					ctxt.Logf("save suppressed in: %s\n", cursym.Name)
				}
				cursym.Func().Text.Mark |= LEAF
			}

			c.autosize = autosize
			p.To.Offset = int64(autosize - rasize) // FP offsets need an updated p.To.Offset.

			if cursym.Func().Text.From.Sym.ISR() {
				// Emit the interrupt handler prologue:
				//
				//  PUSH  [R4-R11]
				//  TST   $0x10, LR
				//  BNE   nofpuctx
				//	MOVW  CONTROL, R0
				//  CPSID i
				//  VPUSH [D8-D15]    // sets CONTROL.FPCA
				//  MOVW  R0, CONTROL // clears FPCA
				//  CPSIE i
				// nofpuctx:
				//
				// The current Go ABI in the presence of the FPU makes the
				// exception entry and exit very inefficient. If a thread uses
				// FPU the above code saves the whole FPU context at exception
				// entry even if the exception handler does not use the FPU. To
				// prevent the higher priority exception to save the same FPU
				// context again the above code clears CONTROL.FPCA just after
				// VPUSH. An another approach can be disabling the FPU at
				// exception entry and re-enabling it only if the UsageFault
				// occurs as descibed in ARM Application Note 298. The undoubted
				// advantages of the ARMv7-M exception model deffinitely require
				// to follow AAPCS.
				//
				p = obj.Appendp(p, newprog)
				p.As = AMOVM
				p.Scond |= C_DB | C_WBIT
				p.From.Type = obj.TYPE_REGLIST
				p.From.Offset = 0xFF0
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = REGSP
				p = obj.Appendp(p, newprog)
				p.As = ATST
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 0x10
				p.Reg = REGLINK
				bne := obj.Appendp(p, newprog)
				bne.As = ABNE
				bne.To.Type = obj.TYPE_BRANCH
				p = obj.Appendp(bne, newprog)
				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_CONTROL
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_R0
				p = obj.Appendp(p, newprog)
				p.As = ACPSID
				p = obj.Appendp(p, newprog)
				p.As = AHWORD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 0xED2D
				p = obj.Appendp(p, newprog)
				p.As = AHWORD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 0x8B10
				p = obj.Appendp(p, newprog)
				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_R0
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_CONTROL
				p = obj.Appendp(p, newprog)
				p.As = ACPSIE
				p = obj.Appendp(p, newprog)
				p.As = obj.ANOP
				bne.To.SetTarget(p)
			}

			if cursym.Func().Text.Mark&LEAF != 0 {
				cursym.Set(obj.AttrLeaf, true)
				if cursym.Func().Text.From.Sym.NoFrame() {
					break
				}
			}

			if !cursym.Func().Text.From.Sym.NoSplit() {
				p = c.stacksplit(p, int32(autosize)) // emit split check
			}

			if autosize <= 255 {
				// fram size fits into MOVW.W R14,-autosize(SP)
				p = obj.Appendp(p, newprog)
				p.As = AMOVW
				p.Scond |= C_WBIT
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGLINK
				p.To.Type = obj.TYPE_MEM
				p.To.Offset = int64(-autosize)
				p.To.Reg = REGSP
				p.Spadj = int32(autosize)
			} else {
				// Frame size is too large for a MOVW.W instruction.
				// Store link register before decrementing SP, so if a signal comes
				// during the execution of the function prologue, the traceback
				// code will not see a half-updated stack frame.
				p = obj.Appendp(p, c.newprog)
				p.Pos = p.Pos
				p.As = ASUB
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = int64(autosize)
				p.Reg = REGSP
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REGTMP

				p = obj.Appendp(p, c.newprog)
				p.Pos = p.Pos
				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGLINK
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = REGTMP

				p = obj.Appendp(p, c.newprog)
				p.Pos = p.Pos
				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGTMP
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REGSP
				p.Spadj = int32(autosize)
			}

			if cursym.Func().Text.From.Sym.ISR() {
				p = obj.Appendp(p, newprog)
				p.As = ABL
				p.To.Type = obj.TYPE_BRANCH
				p.To.Sym = c.ctxt.Lookup("runtime.identcurcpu")
				p = obj.Appendp(p, newprog)
				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_R0
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REGG
			}

			if cursym.Func().Text.From.Sym.Wrapper() {
				// if(g->panic != nil && g->panic->argp == FP) g->panic->argp = bottom-of-frame
				//
				//	MOVW g_panic(g), R1
				//	CMP  $0, R1
				//	BNE checkargp
				// end:
				//	NOP
				// ... function ...
				// checkargp:
				//	MOVW panic_argp(R1), R2
				//	ADD  $(autosize+4), R13, R3
				//	CMP  R2, R3
				//	BNE end
				//	ADD  $4, R13, R4
				//	MOVW R4, panic_argp(R1)
				//	B    end
				//
				// The NOP is needed to give the jumps somewhere to land.
				// It is a liblink NOP, not an ARM NOP: it encodes to 0 instruction bytes.

				p = obj.Appendp(p, newprog)
				p.As = AMOVW
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REGG
				p.From.Offset = 4 * int64(ctxt.Arch.PtrSize) // G.panic
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_R1

				p = obj.Appendp(p, newprog)
				p.As = ACMP
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 0
				p.Reg = REG_R1

				// B.NE checkargp
				bne := obj.Appendp(p, newprog)
				bne.As = ABNE
				bne.To.Type = obj.TYPE_BRANCH

				// end: NOP
				end := obj.Appendp(bne, newprog)
				end.As = obj.ANOP

				// find end of function
				var last *obj.Prog
				for last = end; last.Link != nil; last = last.Link {
				}

				// MOVW panic_argp(R1), R2
				mov := obj.Appendp(last, newprog)
				mov.As = AMOVW
				mov.From.Type = obj.TYPE_MEM
				mov.From.Reg = REG_R1
				mov.From.Offset = 0 // Panic.argp
				mov.To.Type = obj.TYPE_REG
				mov.To.Reg = REG_R2

				// B.NE branch target is MOVW above
				bne.To.SetTarget(mov)

				// ADD $(autosize+4), R13, R3
				p = obj.Appendp(mov, newprog)
				p.As = AADD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = int64(autosize) + 4
				p.Reg = REG_R13
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_R3

				// CMP R2, R3
				p = obj.Appendp(p, newprog)
				p.As = ACMP
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_R2
				p.Reg = REG_R3

				// B.NE end
				p = obj.Appendp(p, newprog)
				p.As = ABNE
				p.To.Type = obj.TYPE_BRANCH
				p.To.SetTarget(end)

				// ADD $4, R13, R4
				p = obj.Appendp(p, newprog)
				p.As = AADD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 4
				p.Reg = REG_R13
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_R4

				// MOVW R4, panic_argp(R1)
				p = obj.Appendp(p, newprog)
				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_R4
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = REG_R1
				p.To.Offset = 0 // Panic.argp

				// B end
				p = obj.Appendp(p, newprog)
				p.As = AB
				p.To.Type = obj.TYPE_BRANCH
				p.To.SetTarget(end)

				// reset for subsequent passes
				p = end
			}

		case obj.ARET:
			nocache(p)
			if cursym.Func().Text.Mark&LEAF != 0 && autosize == 0 {
				p.As = AB
				p.From = obj.Addr{}
				if p.To.Sym != nil { // retjmp
					p.To.Type = obj.TYPE_BRANCH
				} else {
					p.To.Type = obj.TYPE_MEM
					p.To.Offset = 0
					p.To.Reg = REGLINK
				}
				break
			}

			sym := p.To.Sym
			p.To.Sym = nil

			if autosize <= 255 {
				// MOVW.P autosize(SP),PC
				p.As = AMOVW
				p.Scond |= C_PBIT
				p.From.Type = obj.TYPE_MEM
				p.From.Offset = int64(autosize)
				p.From.Reg = REGSP
				p.To.Type = obj.TYPE_REG
				p.To.Offset = 0
				p.To.Reg = REGPC
			} else {
				// MOVW (SP), REGLINK
				// ADD  autosize, SP
				// B    (REGLINK)
				p.As = AMOVW
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REGSP
				p.To.Type = obj.TYPE_REG
				p.To.Offset = 0
				p.To.Reg = REGLINK
				p = obj.Appendp(p, newprog)
				p.As = AADD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = int64(autosize)
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REGSP
				p = obj.Appendp(p, newprog)
				p.As = AB
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = REGLINK
			}

			if cursym.Func().Text.From.Sym.ISR() {
				if p.As == AMOVW {
					// MOVW.P autosize(SP),PC -> MOVW.P autosize(SP),LR
					p.To.Reg = REGLINK
					p = obj.Appendp(p, newprog)
				}
				// emit interrupt handler epilogue:
				//
				//  TST  $0x10, LR
				//  BNE  nofpuctx
				//  MOVW CONTROL, R0
				//  TST  $4, R0 // CONTROL.FPCA tells if FPU have been used
				//  BEQ  nofpuctx
				//  VPOP [D8-D15]
				// nofpuctx:
				//  POP  [R4-R11]
				//
				p.As = ATST
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 0x10
				p.Reg = REGLINK
				bne := obj.Appendp(p, newprog)
				bne.As = ABNE
				bne.To.Type = obj.TYPE_BRANCH
				p = obj.Appendp(bne, newprog)
				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_CONTROL
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_R0
				p = obj.Appendp(p, newprog)
				p.As = ATST
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 4
				p.Reg = REG_R0
				beq := obj.Appendp(p, newprog)
				beq.As = ABEQ
				beq.To.Type = obj.TYPE_BRANCH
				p = obj.Appendp(beq, newprog)
				p.As = AHWORD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 0xECBD
				p = obj.Appendp(p, newprog)
				p.As = AHWORD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 0x8B10
				p = obj.Appendp(p, newprog)
				p.As = AMOVM
				p.Scond |= C_IA | C_WBIT
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REGSP
				p.To.Type = obj.TYPE_REGLIST
				p.To.Offset = 0xFF0
				bne.To.SetTarget(p)
				beq.To.SetTarget(p)
				p = obj.Appendp(p, newprog)
				p.As = AB
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = REGLINK
			}

			// If there are instructions following this ARET, they come from a
			// branch with the same stackframe, so no spadj.
			if sym != nil { // retjmp
				if p.As == AMOVW {
					// MOVW.P autosize(SP),PC -> MOVW.P autosize(SP),LR
					p.To.Reg = REGLINK
					p = obj.Appendp(p, newprog)
				}
				p.As = AB
				p.To.Type = obj.TYPE_BRANCH
				p.To.Sym = sym
			}

		case AADD:
			if p.From.Type == obj.TYPE_CONST && p.From.Reg == 0 && p.To.Type == obj.TYPE_REG && p.To.Reg == REGSP {
				p.Spadj = int32(-p.From.Offset)
			}

		case ASUB:
			if p.From.Type == obj.TYPE_CONST && p.From.Reg == 0 && p.To.Type == obj.TYPE_REG && p.To.Reg == REGSP {
				p.Spadj = int32(p.From.Offset)
			}

		case AMOVW:
			switch {
			case p.From.Type == obj.TYPE_ADDR && (p.From.Name == obj.NAME_NONE || p.From.Name == obj.NAME_PARAM || p.From.Name == obj.NAME_AUTO):
				offset := c.offset(&p.From)
				switch p.From.Name {
				case obj.NAME_PARAM, obj.NAME_AUTO:
					p.Reg = REGSP
				default:
					p.Reg = p.From.Reg
				}
				p.From.Sym = nil
				p.From.Name = obj.NAME_NONE
				if offset == 0 {
					p.From.Type = obj.TYPE_REG
					p.From.Reg = p.Reg
					p.From.Offset = 0
					p.Reg = 0
					break
				}
				p.As = AADD
				p.From.Type = obj.TYPE_CONST
				p.From.Reg = 0
				p.From.Offset = offset
				if p.Reg == REGSP && p.To.Reg == REGSP {
					p.Spadj = int32(-offset)
				}
				p.Scond = C_PBIT // preserve flags
			case (p.Scond&C_WBIT != 0) && p.To.Type == obj.TYPE_MEM && p.To.Reg == REGSP:
				p.Spadj = int32(-p.To.Offset)
				if p.To.Offset == -4 && (p.From.Reg <= REG_R7 || p.From.Reg == REG_R14) {
					// T32 MOVW.W Rt, -4(R13) -> T16 MOVM.DB.W [Rt], (R13) == PUSH [Rt]
					p.As = AMOVM
					p.To.Offset = 0
					p.From.Type = obj.TYPE_REGLIST
					p.From.Offset = 1 << uint(p.From.Reg&15)
					p.From.Reg = 0
					p.Scond = C_DB | C_WBIT
				}
			case (p.Scond&C_PBIT != 0) && p.From.Type == obj.TYPE_MEM && p.From.Reg == REGSP:
				p.Spadj = int32(-p.From.Offset)
				if p.From.Offset == 4 && (p.To.Reg <= REG_R7 || p.To.Reg == REG_R15) {
					// T32 MOVW.P 4(R13), Rt -> T16 MOVM.IA.W (R13), [Rt] == POP [Rt]
					p.As = AMOVM
					p.From.Offset = 0
					p.To.Type = obj.TYPE_REGLIST
					p.To.Offset = 1 << uint(p.To.Reg&15)
					p.To.Reg = 0
					p.Scond = C_IA | C_WBIT
				}
			}

		case AMOVM:
			// TODO: convert invalid MOVM with one register to MOVW

		case obj.AGETCALLERPC:
			if cursym.Leaf() {
				/* MOVW LR, Rd */
				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGLINK
				p.Scond = C_PBIT // preserve flags
			} else {
				/* MOVW (RSP), Rd */
				p.As = AMOVW
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REGSP
			}
		}

		if p.To.Type == obj.TYPE_REG && p.To.Reg == REGSP && p.Spadj == 0 {
			f := c.cursym.Func()
			if f.FuncFlag&objabi.FuncFlag_SPWRITE == 0 {
				c.cursym.Func().FuncFlag |= objabi.FuncFlag_SPWRITE
				if ctxt.Debugvlog || !ctxt.IsAsm {
					ctxt.Logf("auto-SPWRITE: %s %v\n", c.cursym.Name, p)
					if !ctxt.IsAsm {
						ctxt.Diag("invalid auto-SPWRITE in non-assembly")
						ctxt.DiagFlush()
						log.Fatalf("bad SPWRITE")
					}
				}
			}
		}
	}
}

func (c *Ctx) stacksplit(p *obj.Prog, framesize int32) *obj.Prog {
	if c.ctxt.Flag_maymorestack != "" {
		// Save LR and make room for REGCTXT.
		const frameSize = 8
		// MOVW.W R14,$-8(SP)
		p = obj.Appendp(p, c.newprog)
		p.As = AMOVW
		p.Scond |= C_WBIT
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REGLINK
		p.To.Type = obj.TYPE_MEM
		p.To.Offset = -frameSize
		p.To.Reg = REGSP
		p.Spadj = frameSize

		// MOVW REGCTXT, 4(SP)
		p = obj.Appendp(p, c.newprog)
		p.As = AMOVW
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REGCTXT
		p.To.Type = obj.TYPE_MEM
		p.To.Offset = 4
		p.To.Reg = REGSP

		// CALL maymorestack
		p = obj.Appendp(p, c.newprog)
		p.As = obj.ACALL
		p.To.Type = obj.TYPE_BRANCH
		// See ../x86/obj6.go
		p.To.Sym = c.ctxt.LookupABI(c.ctxt.Flag_maymorestack, c.cursym.ABI())

		// Restore REGCTXT and LR.

		// MOVW 4(SP), REGCTXT
		p = obj.Appendp(p, c.newprog)
		p.As = AMOVW
		p.From.Type = obj.TYPE_MEM
		p.From.Offset = 4
		p.From.Reg = REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REGCTXT

		// MOVW.P 8(SP), R14
		p.As = AMOVW
		p.Scond |= C_PBIT
		p.From.Type = obj.TYPE_MEM
		p.From.Offset = frameSize
		p.From.Reg = REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REGLINK
		p.Spadj = -frameSize
	}

	// Jump back to here after morestack returns.
	startPred := p

	// MOVW g_stackguard(g), R1
	p = obj.Appendp(p, c.newprog)

	p.As = AMOVW
	p.From.Type = obj.TYPE_MEM
	p.From.Reg = REGG
	p.From.Offset = 2 * int64(c.ctxt.Arch.PtrSize) // G.stackguard0
	if c.cursym.CFunc() {
		p.From.Offset = 3 * int64(c.ctxt.Arch.PtrSize) // G.stackguard1
	}
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_R1

	// Mark the stack bound check and morestack call async nonpreemptible.
	// If we get preempted here, when resumed the preemption request is
	// cleared, but we'll still call morestack, which will double the stack
	// unnecessarily. See issue #35470.
	p = c.ctxt.StartUnsafePoint(p, c.newprog)

	if framesize <= objabi.StackSmall {
		// small stack: SP < stackguard
		//	CMP	stackguard, SP
		p = obj.Appendp(p, c.newprog)

		p.As = ACMP
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R1
		p.Reg = REGSP
	} else if framesize <= objabi.StackBig {
		// large stack: SP-framesize < stackguard-StackSmall
		//	ADD $-(framesize-StackSmall), SP, R2
		//	CMP stackguard, R2
		p = obj.Appendp(p, c.newprog)

		p.As = AADD
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = -(int64(framesize) - objabi.StackSmall)
		p.Reg = REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		p.As = ACMP
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R1
		p.Reg = REG_R2
	} else {
		// Such a large stack we need to protect against underflow.
		// The runtime guarantees SP > objabi.StackBig, but
		// framesize is large enough that SP-framesize may
		// underflow, causing a direct comparison with the
		// stack guard to incorrectly succeed. We explicitly
		// guard against underflow.
		//
		//	// Try subtracting from SP and check for underflow.
		//	// If this underflows, it sets C to 0.
		//	SUB.S $(framesize-StackSmall), SP, R2
		//	// If C is 1 (unsigned >=), compare with guard.
		//	CMP.HS stackguard, R2

		p = obj.Appendp(p, c.newprog)
		p.As = ASUB
		p.Scond = C_SBIT
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = int64(framesize) - objabi.StackSmall
		p.Reg = REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		p.As = ACMP
		p.Scond = C_SCOND_HS
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R1
		p.Reg = REG_R2
	}

	// BLS call-to-morestack (C is 0 or Z is 1)
	bls := obj.Appendp(p, c.newprog)
	bls.As = ABLS
	bls.To.Type = obj.TYPE_BRANCH

	end := c.ctxt.EndUnsafePoint(bls, c.newprog, -1)

	var last *obj.Prog
	for last = c.cursym.Func().Text; last.Link != nil; last = last.Link {
	}

	// Now we are at the end of the function, but logically
	// we are still in function prologue. We need to fix the
	// SP data and PCDATA.
	spfix := obj.Appendp(last, c.newprog)
	spfix.As = obj.ANOP
	spfix.Spadj = -framesize

	pcdata := c.ctxt.EmitEntryStackMap(c.cursym, spfix, c.newprog)
	pcdata = c.ctxt.StartUnsafePoint(pcdata, c.newprog)

	// MOVW	LR, R3
	movw := obj.Appendp(pcdata, c.newprog)
	movw.As = AMOVW
	movw.From.Type = obj.TYPE_REG
	movw.From.Reg = REGLINK
	movw.To.Type = obj.TYPE_REG
	movw.To.Reg = REG_R3

	bls.To.SetTarget(movw)

	// BL runtime.morestack
	call := obj.Appendp(movw, c.newprog)
	call.As = obj.ACALL
	call.To.Type = obj.TYPE_BRANCH
	morestack := "runtime.morestack"
	switch {
	case c.cursym.CFunc():
		morestack = "runtime.morestackc"
	case !c.cursym.Func().Text.From.Sym.NeedCtxt():
		morestack = "runtime.morestack_noctxt"
	}
	call.To.Sym = c.ctxt.Lookup(morestack)

	pcdata = c.ctxt.EndUnsafePoint(call, c.newprog, -1)

	// B start
	b := obj.Appendp(pcdata, c.newprog)
	b.As = obj.AJMP
	b.To.Type = obj.TYPE_BRANCH
	b.To.SetTarget(startPred.Link)
	b.Spadj = +framesize

	return end
}

func progedit(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {
	p.From.Class = 0
	p.To.Class = 0

	c := &Ctx{ctxt: ctxt, newprog: newprog}

	// Rewrite RSB Rm, Rn, Rd to SUB Rn, Rm, Rd to allow T16 encoding
	if p.As == ARSB && p.From.Type == obj.TYPE_REG {
		p.As = ASUB
		if p.Reg == 0 {
			p.Reg = p.To.Reg
		}
		p.From.Reg, p.Reg = p.Reg, p.From.Reg
	}
	// Rewrite 3-arg to 2-arg
	if p.Reg != 0 && p.To.Type == obj.TYPE_REG && p.To.Reg == p.Reg {
		p.Reg = 0
	}
	// Rewrite TYPE_SHIFT R<<i(R) to TYPE_MEM
	if p.From.Type == obj.TYPE_SHIFT && p.From.Reg != 0 {
		if !shifttomem(&p.From) {
			c.ctxt.Diag("bad src addr mode: %v", p)
		}
	} else if p.To.Type == obj.TYPE_SHIFT && p.To.Reg != 0 {
		if !shifttomem(&p.To) {
			c.ctxt.Diag("bad dst addr mode: %v", p)
		}
	}
	if p.As == obj.ADUFFZERO || p.As == obj.ADUFFCOPY {
		// as long we do not support dynamic linking they are identical to BL
		p.As = ABL
	}
	// Rewrite B/BL to symbol as TYPE_BRANCH.
	if p.As == AB || p.As == ABL || ABEQ <= p.As && p.As <= ACBNZ {
		if p.To.Type == obj.TYPE_MEM && (p.To.Name == obj.NAME_EXTERN || p.To.Name == obj.NAME_STATIC) && p.To.Sym != nil {
			p.To.Type = obj.TYPE_BRANCH
		}
		return
	}
	// Rewrite SLL/SRL/SRA/SRR/ to MOVW
	if ASLL <= p.As && p.As <= ASRR {
		Rn := int(p.To.Reg)
		if p.Reg != 0 {
			Rn = int(p.Reg)
			p.Reg = 0
		}
		typ := int(p.As - ASLL)
		p.As = AMOVW
		if p.From.Type == obj.TYPE_REG {
			p.From.Offset = int64(int(p.From.Reg)&15<<8 | typ<<5 | 1<<4 | Rn&15)
		} else {
			p.From.Offset = int64(int(p.From.Offset)<<7 | typ<<5 | Rn&15)
		}
		p.From.Type = obj.TYPE_SHIFT
		p.From.Reg = 0
		return
	}
	switch p.As {
	// Rewrite float constants to values stored in memory.
	case AMOVF:
		if p.From.Type == obj.TYPE_FCONST && c.chipfloat(p.From.Val.(float64)) < 0 && (c.chipzero(p.From.Val.(float64)) < 0 || p.Scond&C_SCOND != C_SCOND_NONE) {
			f32 := float32(p.From.Val.(float64))
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = ctxt.Float32Sym(f32)
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}
	case AMOVD:
		if p.From.Type == obj.TYPE_FCONST && c.chipfloat(p.From.Val.(float64)) < 0 && (c.chipzero(p.From.Val.(float64)) < 0 || p.Scond&C_SCOND != C_SCOND_NONE) {
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = ctxt.Float64Sym(p.From.Val.(float64))
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}
	}
}

func nocache(p *obj.Prog) {
	p.Optab = 0
	p.From.Class = 0
	if p.GetFrom3() != nil {
		p.GetFrom3().Class = 0
	}
	p.To.Class = 0
}

func onesCount(u uint) int {
	n := 0
	for u != 0 {
		n += int(u & 1)
		u >>= 1
	}
	return n
}

func shifttomem(a *obj.Addr) bool {
	if a.Offset>>4&3 != 0 {
		return false // Thumb2 supports only LSL
	}
	shift := uint(a.Offset) >> 7 & 31
	if shift > 3 {
		return false // Thumb2 supports only 2-bit shift
	}
	a.Type = obj.TYPE_MEM
	a.Index = REG_R0 + int16(a.Offset)&15
	if shift != 0 {
		a.Scale = 1 << shift
	}
	return true
}

var unaryDst = map[obj.As]bool{
	AWORD:      true,
	AB:         true,
	ABL:        true,
	ASWI:       true,
	ABKPT:      true,
	obj.AUNDEF: true,
}

var Link = obj.LinkArch{
	Arch:           sys.ArchThumb,
	Init:           buildop,
	Preprocess:     preprocess,
	Assemble:       span,
	Progedit:       progedit,
	UnaryDst:       unaryDst,
	DWARFRegisters: ARMDWARFRegisters,
}
