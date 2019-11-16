// Inferno utils/5l/asm.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/5l/asm.c
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
	"cmd/link/internal/sym"
	"debug/elf"
	"fmt"
	"log"
)

const pcoff = 4 // in Thumb mode PC points 4 bytes forward

func lookupFuncSym(syms *sym.Symbols, name string) *sym.Symbol {
	if s := syms.ROLookup(name, sym.SymVerABI0); s != nil && s.FuncInfo != nil {
		return s
	}
	if s := syms.ROLookup(name, sym.SymVerABIInternal); s != nil && s.FuncInfo != nil {
		return s
	}
	return nil
}

func gentext(ctxt *ld.Link) {
	if ctxt.HeadType != objabi.Hnoos {
		return
	}
	vectors := ctxt.Syms.Lookup("runtime.vectors", 0)
	vectors.Type = sym.STEXT
	vectors.Attr |= sym.AttrReachable
	vectors.Align = 4

	vectorsAdd := func(symname string) {
		if s := lookupFuncSym(ctxt.Syms, symname); s != nil {
			s.Attr |= sym.AttrReachable
			rel := vectors.AddRel()
			rel.Off = int32(len(vectors.P))
			rel.Siz = 4
			rel.Type = objabi.R_ADDR
			rel.Sym = s
		}
		vectors.AddUint32(ctxt.Arch, 0)
	}

	ld.Segdata.Laddr = 2048 // communicate the main stack size to Link.address()
	msp := uint32(ld.RAM.Base + ld.Segdata.Laddr)
	vectors.AddUint32(ctxt.Arch, msp) // Main Stack Pointer after reset
	vectorsAdd(*ld.FlagEntrySymbol)   // Reset

	// system exception handlers
	for irqn := -14; irqn < 0; irqn++ {
		vectorsAdd(ld.CortexmHandler(irqn))
	}

	// search for user defined ISRs: //go:linkname functionName IRQ%d_Handler
	var irqHandlers [480]*sym.Symbol
	irqNum := 0
	for i := range irqHandlers {
		s := lookupFuncSym(ctxt.Syms, ld.CortexmHandler(i))
		if s != nil {
			irqHandlers[i] = s
			irqNum = i + 1
		}
	}
	for _, s := range irqHandlers[:irqNum] {
		if s != nil {
			s.Attr |= sym.AttrReachable
			rel := vectors.AddRel()
			rel.Off = int32(len(vectors.P))
			rel.Siz = 4
			rel.Type = objabi.R_ADDR
			rel.Sym = s
		}
		vectors.AddUint32(ctxt.Arch, 0)
	}
	ctxt.Textp = append(ctxt.Textp, nil)
	copy(ctxt.Textp[1:], ctxt.Textp)
	ctxt.Textp[0] = vectors
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	ctxt.Out.Write32(uint32(sectoff))

	elfsym := r.Xsym.ElfsymForReloc()
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR:
		if r.Siz == 4 {
			ctxt.Out.Write32(uint32(elf.R_ARM_ABS32) | uint32(elfsym)<<8)
		} else {
			return false
		}
	case objabi.R_PCREL:
		if r.Siz == 4 {
			ctxt.Out.Write32(uint32(elf.R_ARM_REL32) | uint32(elfsym)<<8)
		} else {
			return false
		}
	case objabi.R_CALLARM:
		if r.Siz != 4 {
			return false
		}
		// r.Add contains branch opcode and initial addend
		op := uint32(r.Add >> 32)
		switch op & 0xD000F800 {
		case 0x9000F000: // B
			ctxt.Out.Write32(uint32(elf.R_ARM_THM_JUMP24) | uint32(elfsym)<<8)
		case 0xD000F000: // BL
			ctxt.Out.Write32(uint32(elf.R_ARM_THM_PC22) | uint32(elfsym)<<8) // R_ARM_THM_CALL
		case 0x8000F000: // Bcond
			ctxt.Out.Write32(uint32(elf.R_ARM_THM_JUMP19) | uint32(elfsym)<<8)
		default:
			return false
		}
	}
	return true
}

// Convert the direct jump relocation r to refer to a trampoline if the target is too far
func trampoline(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol) {
	switch r.Type {
	case objabi.R_CALLARM:
		var maxoffset int64
		switch uint32(r.Add>>32) & 0x9000F800 {
		case 0x9000F000: // B/BL imm24
			maxoffset = 1 << 24
		case 0x8000F000: // Bcond imm20
			maxoffset = 1 << 20
		default:
			ld.Errorf(s, "bad branch opcode")
		}
		t := (ld.Symaddr(r.Sym) + int64(int32(r.Add)) - (s.Value + int64(r.Off)))
		if -maxoffset <= t && t < maxoffset {
			return
		}
		// Direct call too far, need to insert trampoline.
		// Look up existing trampolines first. If we found one within the range
		// of direct call, we can reuse it. Otherwise create a new one.
		offset := t + pcoff
		var tramp *sym.Symbol
		for i := 0; ; i++ {
			name := r.Sym.Name + fmt.Sprintf("%+d-tramp%d", offset, i)
			tramp = ctxt.Syms.Lookup(name, int(r.Sym.Version))
			if tramp.Type == sym.SDYNIMPORT {
				// don't reuse trampoline defined in other module
				continue
			}
			if tramp.Value == 0 {
				// Either the trampoline does not exist -- we need to create one,
				// or found one the address which is not assigned -- this will be
				// laid down immediately after the current function. Use this one.
				break
			}
			t = (ld.Symaddr(tramp) - pcoff - (s.Value + int64(r.Off)))
			if -maxoffset <= t && t < maxoffset {
				// found an existing trampoline that is not too far
				// we can just use it
				break
			}
		}
		if tramp.Type == 0 {
			// trampoline does not exist, create one
			ctxt.AddTramp(tramp)
			gentramp(ctxt.Arch, ctxt.LinkMode, tramp, r.Sym, int64(offset))
		}
		// modify reloc to point to tramp, which will be resolved later
		r.Sym = tramp
		r.Add = r.Add&^0xFFFFFFFF | int64(-pcoff)&0xFFFFFFFF // clear the offset embedded in the instruction
		r.Done = false
	default:
		ld.Errorf(s, "trampoline called with non-jump reloc: %d (%s)", r.Type, sym.RelocName(ctxt.Arch, r.Type))
	}
}

// generate a trampoline to target+offset
func gentramp(arch *sys.Arch, linkmode ld.LinkMode, tramp, target *sym.Symbol, offset int64) {
	t := ld.Symaddr(target) + offset

	o1 := uint16(0x4F00) // MOVW (R15), R7 // R15 is actual pc+4 (points to o3)
	o2 := uint16(0x4738) // B  (R7)
	o3 := uint32(t) | 1  // WORD $(target|1)

	tramp.Size = 8
	tramp.P = make([]byte, tramp.Size)
	arch.ByteOrder.PutUint16(tramp.P, o1)
	arch.ByteOrder.PutUint16(tramp.P[2:], o2)
	arch.ByteOrder.PutUint32(tramp.P[4:], o3)
}

func archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val int64) (int64, bool) {
	if ctxt.LinkMode == ld.LinkExternal {
		log.Fatalf("BUGL: external linking not supported")
		return val, false
	}

	switch r.Type {
	case objabi.R_CONST:
		return r.Add, true
	case objabi.R_GOTOFF:
		return ld.Symaddr(r.Sym) + r.Add - ld.Symaddr(ctxt.Syms.Lookup(".got", 0)), true
	case objabi.R_PLT0, objabi.R_PLT1, objabi.R_PLT2:
		log.Fatalf("BUGL: PLT not supported")
	case objabi.R_CALLARM: // B, BL
		// r.Add contains branch opcode and initial addend
		op := uint32(r.Add >> 32)
		t := (ld.Symaddr(r.Sym) + int64(int32(r.Add)) - (s.Value + int64(r.Off)))
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
			return int64(op | j1<<29 | j2<<27 | imm11<<16 | s<<10 | imm10), true
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
			return int64(op | j1<<29 | j2<<27 | imm11<<16 | s<<10 | imm6), true
		default:
			ld.Errorf(s, "bad branch opcode")
		}
		ld.Errorf(s, "direct call too far: %s %x", r.Sym.Name, t)
	}
	return val, false
}

func archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	log.Fatalf("unexpected relocation variant")
	return t
}

func asmb(ctxt *ld.Link) {
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f asmb\n", ld.Cputime())
	}

	if ctxt.IsELF {
		ld.Asmbelfsetup()
	}

	sect := ld.Segtext.Sections[0]
	ctxt.Out.SeekSet(int64(sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff))
	ld.Codeblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	for _, sect = range ld.Segtext.Sections[1:] {
		ctxt.Out.SeekSet(int64(sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff))
		ld.Datblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	}

	if ld.Segrodata.Filelen > 0 {
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f rodatblk\n", ld.Cputime())
		}
		ctxt.Out.SeekSet(int64(ld.Segrodata.Fileoff))
		ld.Datblk(ctxt, int64(ld.Segrodata.Vaddr), int64(ld.Segrodata.Filelen))
	}
	if ld.Segrelrodata.Filelen > 0 {
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f relrodatblk\n", ld.Cputime())
		}
		ctxt.Out.SeekSet(int64(ld.Segrelrodata.Fileoff))
		ld.Datblk(ctxt, int64(ld.Segrelrodata.Vaddr), int64(ld.Segrelrodata.Filelen))
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f datblk\n", ld.Cputime())
	}

	ctxt.Out.SeekSet(int64(ld.Segdata.Fileoff))
	ld.Datblk(ctxt, int64(ld.Segdata.Vaddr), int64(ld.Segdata.Filelen))

	ctxt.Out.SeekSet(int64(ld.Segdwarf.Fileoff))
	ld.Dwarfblk(ctxt, int64(ld.Segdwarf.Vaddr), int64(ld.Segdwarf.Filelen))
}

func asmb2(ctxt *ld.Link) {
	machlink := uint32(0)
	if ctxt.HeadType == objabi.Hdarwin {
		machlink = uint32(ld.Domacholink(ctxt))
	}

	/* output symbol table */
	ld.Symsize = 0

	ld.Lcsize = 0
	symo := uint32(0)
	if !*ld.FlagS {
		// TODO: rationalize
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f sym\n", ld.Cputime())
		}
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				symo = uint32(ld.Segdwarf.Fileoff + ld.Segdwarf.Filelen)
				symo = uint32(ld.Rnd(int64(symo), int64(*ld.FlagRound)))
			}

		case objabi.Hplan9:
			symo = uint32(ld.Segdata.Fileoff + ld.Segdata.Filelen)

		case objabi.Hdarwin:
			symo = uint32(ld.Segdwarf.Fileoff + uint64(ld.Rnd(int64(ld.Segdwarf.Filelen), int64(*ld.FlagRound))) + uint64(machlink))
		}

		ctxt.Out.SeekSet(int64(symo))
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				if ctxt.Debugvlog != 0 {
					ctxt.Logf("%5.2f elfsym\n", ld.Cputime())
				}
				ld.Asmelfsym(ctxt)
				ctxt.Out.Flush()
				ctxt.Out.Write(ld.Elfstrdat)

				if ctxt.LinkMode == ld.LinkExternal {
					ld.Elfemitreloc(ctxt)
				}
			}

		case objabi.Hplan9:
			ld.Asmplan9sym(ctxt)
			ctxt.Out.Flush()

			sym := ctxt.Syms.Lookup("pclntab", 0)
			if sym != nil {
				ld.Lcsize = int32(len(sym.P))
				ctxt.Out.Write(sym.P)
				ctxt.Out.Flush()
			}

		case objabi.Hdarwin:
			if ctxt.LinkMode == ld.LinkExternal {
				ld.Machoemitreloc(ctxt)
			}
		}
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f header\n", ld.Cputime())
	}
	ctxt.Out.SeekSet(0)
	switch ctxt.HeadType {
	default:
	case objabi.Hplan9: /* plan 9 */
		ctxt.Out.Write32b(0x647)                      /* magic */
		ctxt.Out.Write32b(uint32(ld.Segtext.Filelen)) /* sizes */
		ctxt.Out.Write32b(uint32(ld.Segdata.Filelen))
		ctxt.Out.Write32b(uint32(ld.Segdata.Length - ld.Segdata.Filelen))
		ctxt.Out.Write32b(uint32(ld.Symsize))          /* nsyms */
		ctxt.Out.Write32b(uint32(ld.Entryvalue(ctxt))) /* va of entry */
		ctxt.Out.Write32b(0)
		ctxt.Out.Write32b(uint32(ld.Lcsize))

	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd,
		objabi.Hnacl,
		objabi.Hnoos:
		ld.Asmbelf(ctxt, int64(symo))

	case objabi.Hdarwin:
		ld.Asmbmacho(ctxt)
	}

	ctxt.Out.Flush()
	if *ld.FlagC {
		fmt.Printf("textsize=%d\n", ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", ld.Segdata.Length-ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", ld.Symsize)
		fmt.Printf("lcsize=%d\n", ld.Lcsize)
		fmt.Printf("total=%d\n", ld.Segtext.Filelen+ld.Segdata.Length+uint64(ld.Symsize)+uint64(ld.Lcsize))
	}
}
