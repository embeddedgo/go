// Inferno utils/5l/asm.c
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/5l/asm.c
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
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
	"cmd/link/internal/loader"
	"cmd/link/internal/sym"
	"debug/elf"
	"fmt"
	"log"
)

const pcoff = 4 // in Thumb mode PC points 4 bytes forward

func lookupFuncSym(ldr *loader.Loader, name string) loader.Sym {
	if s := ldr.Lookup(name, sym.SymVerABI0); s != 0 && ldr.SymType(s) == sym.STEXT {
		return s
	}
	if s := ldr.Lookup(name, sym.SymVerABIInternal); s != 0 && ldr.SymType(s) == sym.STEXT {
		return s
	}
	return 0
}

func gentext(ctxt *ld.Link, ldr *loader.Loader) {
	if ctxt.HeadType != objabi.Hnoos {
		return
	}
	vectors := ldr.CreateSymForUpdate("runtime.vectors", sym.SymVerABI0)
	vectors.SetType(sym.STEXT)
	vectors.SetReachable(true)
	vectors.SetAlign(4)

	unhandledException := ldr.Lookup("runtime.unhandledException", sym.SymVerABI0)
	if unhandledException == 0 {
		ld.Errorf(nil, "runtime.unhandledException not defined")
	}

	// search for user defined ISRs: //go:linkname functionName IRQ%d_Handler
	var irqHandlers [480]loader.Sym
	irqNum := 0
	for i := range irqHandlers {
		s := lookupFuncSym(ldr, ld.InterruptHandler(i))
		if s == 0 {
			irqHandlers[i] = unhandledException
		} else {
			irqHandlers[i] = s
			irqNum = i + 1 // BUG: saves memory but prevents detect all unhandled interrupts
		}
	}

	ld.Segdata.Laddr = 2048 // communicate the main stack size to Link.address()
	msp := uint32(ld.RAM.Base + ld.Segdata.Laddr)
	vectors.AddUint32(ctxt.Arch, msp) // Main Stack Pointer after reset

	relocs := vectors.AddRelocs(irqNum + 15)
	addHandler := func(irqn int, fname string) {
		s := lookupFuncSym(ldr, fname)
		if s == 0 {
			s = unhandledException
		}
		ldr.MakeSymbolUpdater(s).SetReachable(true)
		rel := relocs.At(irqn + 15)
		rel.SetSym(s)
		rel.SetType(objabi.R_ADDR)
		rel.SetSiz(4)
		rel.SetOff(int32(vectors.AddUint32(ctxt.Arch, 0)))
	}

	// system exception handlers
	addHandler(-15, *ld.FlagEntrySymbol) // reset handler
	for irqn := -14; irqn < 0; irqn++ {
		addHandler(irqn, ld.InterruptHandler(irqn)) // exception handler
	}

	// external interrupt handlers
	for i, s := range irqHandlers[:irqNum] {
		ldr.MakeSymbolUpdater(s).SetReachable(true)
		rel := relocs.At(i + 15)
		rel.SetSym(s)
		rel.SetType(objabi.R_ADDR)
		rel.SetSiz(4)
		rel.SetOff(int32(vectors.AddUint32(ctxt.Arch, 0)))
	}

	// add the vectors symbol at the beggining of the text segment
	ctxt.Textp = append(ctxt.Textp, 0)
	copy(ctxt.Textp[1:], ctxt.Textp)
	ctxt.Textp[0] = vectors.Sym()
}

func elfreloc1(ctxt *ld.Link, out *ld.OutBuf, ldr *loader.Loader, s loader.Sym, r loader.ExtReloc, ri int, sectoff int64) bool {
	out.Write32(uint32(sectoff))

	elfsym := ld.ElfSymForReloc(ctxt, r.Xsym)
	siz := r.Size
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR, objabi.R_DWARFSECREF:
		if siz == 4 {
			out.Write32(uint32(elf.R_ARM_ABS32) | uint32(elfsym)<<8)
		} else {
			return false
		}
	case objabi.R_PCREL:
		if siz == 4 {
			out.Write32(uint32(elf.R_ARM_REL32) | uint32(elfsym)<<8)
		} else {
			return false
		}
	case objabi.R_CALLARM:
		if siz != 4 {
			return false
		}
		relocs := ldr.Relocs(s)
		r := relocs.At(ri)
		// r.Add() contains branch opcode and initial addend
		op := uint32(r.Add() >> 32)
		switch op & 0xD000F800 {
		case 0x9000F000: // B
			out.Write32(uint32(elf.R_ARM_THM_JUMP24) | uint32(elfsym)<<8)
		case 0xD000F000: // BL
			out.Write32(uint32(elf.R_ARM_THM_PC22) | uint32(elfsym)<<8) // R_ARM_THM_CALL
		case 0x8000F000: // Bcond
			out.Write32(uint32(elf.R_ARM_THM_JUMP19) | uint32(elfsym)<<8)
		default:
			return false
		}
	}
	return true
}

// Convert the direct jump relocation r to refer to a trampoline if the target is too far
func trampoline(ctxt *ld.Link, ldr *loader.Loader, ri int, rs, s loader.Sym) {
	relocs := ldr.Relocs(s)
	r := relocs.At(ri)
	switch r.Type() {
	case objabi.R_CALLARM:
		var maxoffset int64
		switch uint32(r.Add()>>32) & 0x9000F800 {
		case 0x9000F000: // B/BL imm24
			maxoffset = 1 << 24
		case 0x8000F000: // Bcond imm20
			maxoffset = 1 << 20
		default:
			ldr.Errorf(s, "bad branch opcode")
		}
		t := (ldr.SymValue(rs) + int64(int32(r.Add())) - (ldr.SymValue(s) + int64(r.Off())))
		if -maxoffset <= t && t < maxoffset {
			return
		}
		// Direct call too far, need to insert trampoline.
		// Look up existing trampolines first. If we found one within the range
		// of direct call, we can reuse it. Otherwise create a new one.
		offset := t + pcoff
		var tramp loader.Sym
		for i := 0; ; i++ {
			oName := ldr.SymName(rs)
			name := oName + fmt.Sprintf("%+d-tramp%d", offset, i)
			tramp = ldr.LookupOrCreateSym(name, int(ldr.SymVersion(rs)))
			ldr.SetAttrReachable(tramp, true)
			if ldr.SymType(tramp) == sym.SDYNIMPORT {
				// don't reuse trampoline defined in other module
				continue
			}
			if oName == "runtime.deferreturn" {
				ldr.SetIsDeferReturnTramp(tramp, true)
			}
			if ldr.SymValue(tramp) == 0 {
				// either the trampoline does not exist -- we need to create one,
				// or found one the address which is not assigned -- this will be
				// laid down immediately after the current function. use this one.
				break
			}
			t = (ldr.SymValue(tramp) - pcoff - (ldr.SymValue(s) + int64(r.Off())))
			if -maxoffset <= t && t < maxoffset {
				// found an existing trampoline that is not too far
				// we can just use it
				break
			}
		}
		if ldr.SymType(tramp) == 0 {
			// trampoline does not exist, create one
			trampb := ldr.MakeSymbolUpdater(tramp)
			ctxt.AddTramp(trampb)
			gentramp(ctxt.Arch, ctxt.LinkMode, ldr, trampb, rs, int64(offset))
		}
		// modify reloc to point to tramp, which will be resolved later
		sb := ldr.MakeSymbolUpdater(s)
		relocs := sb.Relocs()
		r := relocs.At(ri)
		r.SetSym(tramp)
		r.SetAdd(r.Add()&^0xFFFFFFFF | int64(-pcoff)&0xFFFFFFFF) // clear the offset embedded in the instruction
	default:
		ctxt.Errorf(s, "trampoline called with non-jump reloc: %d (%s)", r.Type(), sym.RelocName(ctxt.Arch, r.Type()))
	}
}

// generate a trampoline to target+offset
func gentramp(arch *sys.Arch, linkmode ld.LinkMode, ldr *loader.Loader, tramp *loader.SymbolBuilder, target loader.Sym, offset int64) {
	tramp.SetSize(8) // 2+1 instructions
	P := make([]byte, tramp.Size())
	t := ldr.SymValue(target) + offset
	o1 := uint16(0x4F00) // MOVW (R15), R7 // R15 is actual pc+4 (points to o3)
	o2 := uint16(0x4738) // B  (R7)
	o3 := uint32(t) | 1  // WORD $(target|1)
	arch.ByteOrder.PutUint16(P, o1)
	arch.ByteOrder.PutUint16(P[2:], o2)
	arch.ByteOrder.PutUint32(P[4:], o3)
	tramp.SetData(P)
}

func archreloc(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, r loader.Reloc, s loader.Sym, val int64) (o int64, nExtReloc int, ok bool) {
	rs := r.Sym()
	rs = ldr.ResolveABIAlias(rs)
	if target.IsExternal() {
		log.Fatalf("BUGL: external linking not supported")
		return val, 0, false
	}

	const isOk = true
	const noExtReloc = 0
	switch r.Type() {
	case objabi.R_PLT0, objabi.R_PLT1, objabi.R_PLT2:
		log.Fatalf("BUGL: PLT not supported")
	case objabi.R_CALLARM: // B, BL
		// r.Add() returns branch opcode and initial addend
		op := uint32(r.Add() >> 32)
		t := (ldr.SymValue(rs) + int64(int32(r.Add())) - (ldr.SymValue(s) + int64(r.Off())))
		switch op & 0x9000F800 {
		case 0x9000F000: // B/BL imm24
			if t < -1<<24 || 1<<24 <= t {
				break
			}
			v := uint32(t >> 1)
			s := v >> 23 & 1
			j1 := ^(v>>22 ^ s) & 1
			j2 := ^(v>>21 ^ s) & 1
			imm10 := v >> 11 & 0x3FF
			imm11 := v & 0x7FF
			return int64(op | j1<<29 | j2<<27 | imm11<<16 | s<<10 | imm10), noExtReloc, isOk
		case 0x8000F000: // Bcond imm20
			if t < -1<<20 || 1<<20 <= t {
				break
			}
			v := uint32(t >> 1)
			s := v >> 19 & 1
			j2 := v >> 18 & 1
			j1 := v >> 17 & 1
			imm6 := v >> 11 & 0x3F
			imm11 := v & 0x7FF
			return int64(op | j1<<29 | j2<<27 | imm11<<16 | s<<10 | imm6), noExtReloc, isOk
		default:
			ldr.Errorf(s, "bad branch opcode")
		}
		ldr.Errorf(s, "direct call too far: %s %x", ldr.SymName(rs), t)
	}
	return val, 0, false
}

func archrelocvariant(*ld.Target, *loader.Loader, loader.Reloc, sym.RelocVariant, loader.Sym, int64) int64 {
	log.Fatalf("unexpected relocation variant")
	return -1
}
