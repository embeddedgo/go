// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64

import (
	"cmd/internal/obj/riscv"
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
	"cmd/link/internal/loader"
	"cmd/link/internal/sym"
	"fmt"
	"log"
	"sync"
)

func lookupFuncSym(ldr *loader.Loader, name string) loader.Sym {
	if s := ldr.Lookup(name, sym.SymVerABI0); s != 0 && ldr.SymType(s) == sym.STEXT {
		return s
	}
	if s := ldr.Lookup(name, sym.SymVerABIInternal); s != 0 && ldr.SymType(s) == sym.STEXT {
		return s
	}
	return 0
}

func gentext2(ctxt *ld.Link, ldr *loader.Loader) {
	if ctxt.HeadType != objabi.Hnoos {
		return
	}
	vectors := ldr.CreateSymForUpdate("runtime.vectors", sym.SymVerABI0)
	vectors.SetType(sym.STEXT)
	vectors.SetReachable(true)
	vectors.SetAlign(8)

	unhandledInterrupt := ldr.Lookup("runtime.unhandledExternalInterrupt", sym.SymVerABI0)
	if unhandledInterrupt == 0 {
		ld.Errorf(nil, "runtime.unhandledExternalInterrupt not defined")
	}

	// search for user defined ISRs: //go:linkname functionName IRQ%d_Handler
	var irqHandlers [1024]loader.Sym // BUG: 1024 is PLIC specific
	irqNum := 1
	for i := 1; i < len(irqHandlers); i++ {
		s := lookupFuncSym(ldr, ld.InterruptHandler(i))
		if s == 0 {
			irqHandlers[i] = unhandledInterrupt
		} else {
			irqHandlers[i] = s
			irqNum = i + 1
		}
	}

	vectors.AddUint64(ctxt.Arch, uint64(irqNum))
	relocs := vectors.AddRelocs(irqNum - 1)
	for i, s := range irqHandlers[1:irqNum] {
		ldr.MakeSymbolUpdater(s).SetReachable(true)
		rel := relocs.At2(i)
		rel.SetSym(s)
		rel.SetType(objabi.R_ADDR)
		rel.SetSiz(8)
		rel.SetOff(int32(vectors.AddUint64(ctxt.Arch, 0)))
	}

	ctxt.Textp2 = append(ctxt.Textp2, vectors.Sym())

	// move the entry symbol at the beggining of the text segment
	entry := lookupFuncSym(ldr, *ld.FlagEntrySymbol)
	if entry == 0 {
		ld.Errorf(nil, "cannot find entry function: %s", *ld.FlagEntrySymbol)
	}
	for i, s := range ctxt.Textp2 {
		if s == entry {
			copy(ctxt.Textp2[1:], ctxt.Textp2[:i])
			ctxt.Textp2[0] = s
			return
		}
	}
	ldr.Errorf(entry, "cannot find symbol in ctxt.Textp")
}

func adddynrela(target *ld.Target, syms *ld.ArchSyms, rel *sym.Symbol, s *sym.Symbol, r *sym.Reloc) {
	log.Fatalf("adddynrela not implemented")
}

func adddynrel(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, s *sym.Symbol, r *sym.Reloc) bool {
	log.Fatalf("adddynrel not implemented")
	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	log.Fatalf("elfreloc1")
	return false
}

func elfsetupplt(ctxt *ld.Link, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym) {
	log.Fatalf("elfsetuplt")
}

func machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	log.Fatalf("machoreloc1 not implemented")
	return false
}

func archreloc(target *ld.Target, syms *ld.ArchSyms, r *sym.Reloc, s *sym.Symbol, val int64) (int64, bool) {
	switch r.Type {
	case objabi.R_CALLRISCV:
		// Nothing to do.
		return val, true

	case objabi.R_RISCV_PCREL_ITYPE, objabi.R_RISCV_PCREL_STYPE:
		pc := s.Value + int64(r.Off)
		off := ld.Symaddr(r.Sym) + r.Add - pc

		// Generate AUIPC and second instruction immediates.
		low, high, err := riscv.Split32BitImmediate(off)
		if err != nil {
			ld.Errorf(s, "R_RISCV_PCREL_ relocation does not fit in 32-bits: %d", off)
		}

		auipcImm, err := riscv.EncodeUImmediate(high)
		if err != nil {
			ld.Errorf(s, "cannot encode R_RISCV_PCREL_ AUIPC relocation offset for %s: %v", r.Sym.Name, err)
		}

		var secondImm, secondImmMask int64
		switch r.Type {
		case objabi.R_RISCV_PCREL_ITYPE:
			secondImmMask = riscv.ITypeImmMask
			secondImm, err = riscv.EncodeIImmediate(low)
			if err != nil {
				ld.Errorf(s, "cannot encode R_RISCV_PCREL_ITYPE I-type instruction relocation offset for %s: %v", r.Sym.Name, err)
			}
		case objabi.R_RISCV_PCREL_STYPE:
			secondImmMask = riscv.STypeImmMask
			secondImm, err = riscv.EncodeSImmediate(low)
			if err != nil {
				ld.Errorf(s, "cannot encode R_RISCV_PCREL_STYPE S-type instruction relocation offset for %s: %v", r.Sym.Name, err)
			}
		default:
			panic(fmt.Sprintf("Unknown relocation type: %v", r.Type))
		}

		auipc := int64(uint32(val))
		second := int64(uint32(val >> 32))

		auipc = (auipc &^ riscv.UTypeImmMask) | int64(uint32(auipcImm))
		second = (second &^ secondImmMask) | int64(uint32(secondImm))

		return second<<32 | auipc, true
	}

	return val, false
}

func archrelocvariant(target *ld.Target, syms *ld.ArchSyms, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	log.Fatalf("archrelocvariant")
	return -1
}

func asmb(ctxt *ld.Link, _ *loader.Loader) {
	if ctxt.IsELF {
		ld.Asmbelfsetup()
	}

	var wg sync.WaitGroup
	sect := ld.Segtext.Sections[0]
	offset := sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff
	ld.WriteParallel(&wg, ld.Codeblk, ctxt, offset, sect.Vaddr, sect.Length)

	for _, sect := range ld.Segtext.Sections[1:] {
		offset := sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff
		ld.WriteParallel(&wg, ld.Datblk, ctxt, offset, sect.Vaddr, sect.Length)
	}

	if ld.Segrodata.Filelen > 0 {
		ld.WriteParallel(&wg, ld.Datblk, ctxt, ld.Segrodata.Fileoff, ld.Segrodata.Vaddr, ld.Segrodata.Filelen)
	}

	if ld.Segrelrodata.Filelen > 0 {
		ld.WriteParallel(&wg, ld.Datblk, ctxt, ld.Segrelrodata.Fileoff, ld.Segrelrodata.Vaddr, ld.Segrelrodata.Filelen)
	}

	ld.WriteParallel(&wg, ld.Datblk, ctxt, ld.Segdata.Fileoff, ld.Segdata.Vaddr, ld.Segdata.Filelen)

	ld.WriteParallel(&wg, ld.Dwarfblk, ctxt, ld.Segdwarf.Fileoff, ld.Segdwarf.Vaddr, ld.Segdwarf.Filelen)
	wg.Wait()
}

func asmb2(ctxt *ld.Link) {
	ld.Symsize = 0
	ld.Lcsize = 0
	symo := uint32(0)

	if !*ld.FlagS {
		if !ctxt.IsELF {
			ld.Errorf(nil, "unsupported executable format")
		}

		symo = uint32(ld.Segdwarf.Fileoff + ld.Segdwarf.Filelen)
		symo = uint32(ld.Rnd(int64(symo), int64(*ld.FlagRound)))
		ctxt.Out.SeekSet(int64(symo))

		ld.Asmelfsym(ctxt)
		ctxt.Out.Write(ld.Elfstrdat)

		if ctxt.LinkMode == ld.LinkExternal {
			ld.Elfemitreloc(ctxt)
		}
	}

	ctxt.Out.SeekSet(0)
	switch ctxt.HeadType {
	case objabi.Hlinux, objabi.Hnoos:
		ld.Asmbelf(ctxt, int64(symo))
	default:
		ld.Errorf(nil, "unsupported operating system")
	}

	if *ld.FlagC {
		fmt.Printf("textsize=%d\n", ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", ld.Segdata.Length-ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", ld.Symsize)
		fmt.Printf("lcsize=%d\n", ld.Lcsize)
		fmt.Printf("total=%d\n", ld.Segtext.Filelen+ld.Segdata.Length+uint64(ld.Symsize)+uint64(ld.Lcsize))
	}
}
