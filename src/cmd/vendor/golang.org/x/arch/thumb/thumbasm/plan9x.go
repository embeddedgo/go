// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package thumbasm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"
)

// GoSyntax returns the Go assembler syntax for the instruction.
// The syntax was originally defined by Plan 9.
// The pc is the program counter of the instruction, used for expanding
// PC-relative addresses into absolute ones.
// The symname function queries the symbol table for the program
// being disassembled. Given a target address it returns the name and base
// address of the symbol containing the target, if any; otherwise it returns "", 0.
// The reader r should read from the text segment using text addresses
// as offsets; it is used to display pc-relative loads as constant loads.
func GoSyntax(inst Inst, pc uint64, symname func(uint64) (string, uint64), text io.ReaderAt) string {
	if symname == nil {
		symname = func(uint64) (string, uint64) { return "", 0 }
	}

	//switch inst.Op {
	//case BX, BLX:
	//	if inst.Op == BX {
	//		inst.Op = B
	//	} else {
	//		inst.Op = BL
	//	}
	//	inst.Args[0] = Mem{Base: inst.Args[0].(Reg), Mode: AddrOffset}
	//}

	var args []string
	for _, a := range inst.Args {
		if a == nil {
			break
		}
		args = append(args, plan9Arg(&inst, pc, symname, a))
	}

	op := inst.Op.String()

	switch inst.Op {
	case LDR, LDRB, LDRH, LDRSB, LDRSH, VLDR:
		// Check for RET
		reg, _ := inst.Args[0].(Reg)
		mem, _ := inst.Args[1].(Mem)
		if inst.Op == LDR && reg == R15 && mem.Base == SP && mem.Sign == 0 && mem.Mode == AddrPostIndex {
			return fmt.Sprintf("RET%s #%d", op[3:], mem.Offset)
		}

		// Check for PC-relative load.
		if mem.Base == PC && mem.Sign == 0 && mem.Mode == AddrOffset && text != nil {
			addr := uint32(pc) + 4 + uint32(mem.Offset)
			buf := make([]byte, 8)
			switch inst.Op {
			case LDRB, LDRSB:
				if _, err := text.ReadAt(buf[:1], int64(addr)); err != nil {
					break
				}
				args[1] = fmt.Sprintf("$%#x", buf[0])

			case LDRH, LDRSH:
				if _, err := text.ReadAt(buf[:2], int64(addr)); err != nil {
					break
				}
				args[1] = fmt.Sprintf("$%#x", binary.LittleEndian.Uint16(buf))

			case LDR:
				if _, err := text.ReadAt(buf[:4], int64(addr)); err != nil {
					break
				}
				x := binary.LittleEndian.Uint32(buf)
				if s, base := symname(uint64(x)); s != "" && uint64(x) == base {
					args[1] = fmt.Sprintf("$%s(SB)", s)
				} else {
					args[1] = fmt.Sprintf("$%#x", x)
				}

			case VLDR:
				switch {
				case strings.HasPrefix(args[0], "D"): // VLDR.F64
					if _, err := text.ReadAt(buf, int64(addr)); err != nil {
						break
					}
					args[1] = fmt.Sprintf("$%f", math.Float64frombits(binary.LittleEndian.Uint64(buf)))
				case strings.HasPrefix(args[0], "S"): // VLDR.F32
					if _, err := text.ReadAt(buf[:4], int64(addr)); err != nil {
						break
					}
					args[1] = fmt.Sprintf("$%f", math.Float32frombits(binary.LittleEndian.Uint32(buf)))
				default:
					panic(fmt.Sprintf("wrong FP register: %v", inst))
				}
			}
		}
	}

	// Move addressing mode into opcode suffix.
	suffix := ""
	switch inst.Op {
	case PLD, PLI:
		if mem, ok := inst.Args[0].(Mem); ok {
			args[0], suffix = memOpTrans(mem)
		} else {
			panic(fmt.Sprintf("illegal instruction: %v", inst))
		}
	case LDR, LDRB, LDRSB, LDRH, LDRSH, STR, STRB, STRH, VLDR, VSTR, LDREX, LDREXH, LDREXB:
		if mem, ok := inst.Args[1].(Mem); ok {
			args[1], suffix = memOpTrans(mem)
		} else {
			panic(fmt.Sprintf("illegal instruction: %v", inst))
		}
	case STREX, STREXB, STREXH:
		if mem, ok := inst.Args[2].(Mem); ok {
			args[2], suffix = memOpTrans(mem)
		} else {
			panic(fmt.Sprintf("illegal instruction: %v", inst))
		}
	}

	switch inst.Op {
	case STR, STRB, STRH, CBZ, CBNZ:
		break
	default:
		// Reverse args, placing dest last.
		for i, j := 0, len(args)-1; i < j; i, j = i+1, j-1 {
			args[i], args[j] = args[j], args[i]
		}
	}

	// For MLA-like instructions, the addend is the third operand.
	switch inst.Op {
	case SMLAWT, SMLAWB, MLA, MLS, SMMLA, SMMLS, SMLABB, SMLATB, SMLABT, SMLATT, SMLAD, SMLADX, SMLSD, SMLSDX:
		args = []string{args[1], args[2], args[0], args[3]}
	}
	// For STREX like instructions, the memory operands comes first.
	switch inst.Op {
	case STREX, STREXB, STREXH:
		args = []string{args[1], args[0], args[2]}
	}

	// special process for FP instructions
	op, args = fpTrans(&inst, op, args)

	// LDR/STR like instructions -> MOV like
	switch inst.Op {
	case MOV:
		op = "MOVW" + op[3:]
	case LDR, MSR, MRS:
		op = "MOVW" + op[3:] + suffix
	case VMRS, VMSR:
		op = "MOVW" + op[4:] + suffix
	case LDRB, UXTB:
		op = "MOVBU" + op[4:] + suffix
	case LDRSB:
		op = "MOVBS" + op[5:] + suffix
	case SXTB:
		op = "MOVBS" + op[4:] + suffix
	case LDRH, UXTH:
		op = "MOVHU" + op[4:] + suffix
	case LDRSH:
		op = "MOVHS" + op[5:] + suffix
	case SXTH:
		op = "MOVHS" + op[4:] + suffix
	case STR:
		op = "MOVW" + op[3:] + suffix
	case STRB:
		op = "MOVB" + op[4:] + suffix
	case STRH:
		op = "MOVH" + op[4:] + suffix
	case VSTR:
		args[0], args[1] = args[1], args[0]
	default:
		op = op + suffix
	}

	if args != nil {
		op += " " + strings.Join(args, ", ")
	}

	return op
}

// assembler syntax for the various shifts.
// @x> is a lie; the assembler uses @> 0
// instead of @x> 1, but i wanted to be clear that it
// was a different operation (rotate right extended, not rotate right).
var plan9Shift = []string{"<<", ">>", "->", "@>", "@x>"}

func plan9Arg(inst *Inst, pc uint64, symname func(uint64) (string, uint64), arg Arg) string {
	switch a := arg.(type) {
	case Endian:

	case Imm:
		return fmt.Sprintf("$%d", uint32(a))

	case Mem:
		if a.Mode == AddrOffset && a.Sign == 0 && a.Offset == 0 {
			return fmt.Sprintf("(R%d)", a.Base)
		}

	case PCRel:
		addr := uint64(int64(pc) + 4 + int64(a))
		if s, base := symname(addr); s != "" && addr == base {
			return fmt.Sprintf("%s(SB)", s)
		}
		return fmt.Sprintf("%#x", addr)

	case Reg:
		if a < 16 {
			return fmt.Sprintf("R%d", int(a))
		}

	case RegList:
		var buf bytes.Buffer
		start := -2
		end := -2
		fmt.Fprintf(&buf, "[")
		flush := func() {
			if start >= 0 {
				if buf.Len() > 1 {
					fmt.Fprintf(&buf, ",")
				}
				if start == end {
					fmt.Fprintf(&buf, "R%d", start)
				} else {
					fmt.Fprintf(&buf, "R%d-R%d", start, end)
				}
				start = -2
				end = -2
			}
		}
		for i := 0; i < 16; i++ {
			if a&(1<<uint(i)) != 0 {
				if i == end+1 {
					end++
					continue
				}
				start = i
				end = i
			} else {
				flush()
			}
		}
		flush()
		fmt.Fprintf(&buf, "]")
		return buf.String()

	case RegShift:
		return fmt.Sprintf("R%d%s$%d", int(a.Reg), plan9Shift[a.Shift], int(a.Count))

	}
	return arg.String()
}

// convert memory operand from GNU syntax to Plan 9 syntax, for example,
// [r5] -> (R5)
// [r6, #4080] -> 0xff0(R6)
// [r2, r0, ror #1] -> (R2)(R0@>1)
// inst [r2, -r0, ror #1] -> INST.U (R2)(R0@>1)
// input:
//   a memory operand
// return values:
//   corresponding memory operand in Plan 9 syntax
//   .W/.P/.U suffix
func memOpTrans(mem Mem) (string, string) {
	suffix := ""
	switch mem.Mode {
	case AddrOffset, AddrLDM:
		// no suffix
	case AddrPreIndex, AddrLDM_WB:
		suffix = ".W"
	case AddrPostIndex:
		suffix = ".P"
	}
	off := ""
	if mem.Offset != 0 {
		off = fmt.Sprintf("%#x", mem.Offset)
	}
	base := fmt.Sprintf("(R%d)", int(mem.Base))
	index := ""
	if mem.Sign != 0 {
		sign := ""
		if mem.Sign < 0 {
			suffix += ".U"
		}
		shift := ""
		if mem.Count != 0 {
			shift = fmt.Sprintf("%s%d", plan9Shift[mem.Shift], mem.Count)
		}
		index = fmt.Sprintf("(%sR%d%s)", sign, int(mem.Index), shift)
	}
	return off + base + index, suffix
}

type goFPInfo struct {
	op        Op
	transArgs []int  // indexes of arguments which need transformation
	gnuName   string // instruction name in GNU syntax
	goName    string // instruction name in Plan 9 syntax
}

var fpInst []goFPInfo = []goFPInfo{
	{VADD_F32, []int{2, 1, 0}, "VADD", "ADDF"},
	{VADD_F64, []int{2, 1, 0}, "VADD", "ADDD"},
	{VSUB_F32, []int{2, 1, 0}, "VSUB", "SUBF"},
	{VSUB_F64, []int{2, 1, 0}, "VSUB", "SUBD"},
	{VMUL_F32, []int{2, 1, 0}, "VMUL", "MULF"},
	{VMUL_F64, []int{2, 1, 0}, "VMUL", "MULD"},
	{VNMUL_F32, []int{2, 1, 0}, "VNMUL", "NMULF"},
	{VNMUL_F64, []int{2, 1, 0}, "VNMUL", "NMULD"},
	{VMLA_F32, []int{2, 1, 0}, "VMLA", "MULAF"},
	{VMLA_F64, []int{2, 1, 0}, "VMLA", "MULAD"},
	{VMLS_F32, []int{2, 1, 0}, "VMLS", "MULSF"},
	{VMLS_F64, []int{2, 1, 0}, "VMLS", "MULSD"},
	{VNMLA_F32, []int{2, 1, 0}, "VNMLA", "NMULAF"},
	{VNMLA_F64, []int{2, 1, 0}, "VNMLA", "NMULAD"},
	{VNMLS_F32, []int{2, 1, 0}, "VNMLS", "NMULSF"},
	{VNMLS_F64, []int{2, 1, 0}, "VNMLS", "NMULSD"},
	{VDIV_F32, []int{2, 1, 0}, "VDIV", "DIVF"},
	{VDIV_F64, []int{2, 1, 0}, "VDIV", "DIVD"},
	{VNEG_F32, []int{1, 0}, "VNEG", "NEGF"},
	{VNEG_F64, []int{1, 0}, "VNEG", "NEGD"},
	{VABS_F32, []int{1, 0}, "VABS", "ABSF"},
	{VABS_F64, []int{1, 0}, "VABS", "ABSD"},
	{VSQRT_F32, []int{1, 0}, "VSQRT", "SQRTF"},
	{VSQRT_F64, []int{1, 0}, "VSQRT", "SQRTD"},
	{VCMP_F32, []int{1, 0}, "VCMP", "CMPF"},
	{VCMP_F64, []int{1, 0}, "VCMP", "CMPD"},
	{VCMPE_F32, []int{1, 0}, "VCMPE", "CMPF"},
	{VCMPE_F64, []int{1, 0}, "VCMPE", "CMPD"},
	{VLDR, []int{1}, "VLDR", "MOV"},
	{VSTR, []int{1}, "VSTR", "MOV"},
	{VMOV_F32, []int{1, 0}, "VMOV", "MOVF"},
	{VMOV_F64, []int{1, 0}, "VMOV", "MOVD"},
	{VMOV_32, []int{1, 0}, "VMOV", "MOVW"},
	{VMOV, []int{1, 0}, "VMOV", "MOVW"},
	{VCVT_F64_F32, []int{1, 0}, "VCVT", "MOVFD"},
	{VCVT_F32_F64, []int{1, 0}, "VCVT", "MOVDF"},
	{VCVT_F32_U32, []int{1, 0}, "VCVT", "MOVWF.U"},
	{VCVT_F32_S32, []int{1, 0}, "VCVT", "MOVWF"},
	{VCVT_S32_F32, []int{1, 0}, "VCVT", "MOVFW"},
	{VCVT_U32_F32, []int{1, 0}, "VCVT", "MOVFW.U"},
	{VCVT_F64_U32, []int{1, 0}, "VCVT", "MOVWD.U"},
	{VCVT_F64_S32, []int{1, 0}, "VCVT", "MOVWD"},
	{VCVT_S32_F64, []int{1, 0}, "VCVT", "MOVDW"},
	{VCVT_U32_F64, []int{1, 0}, "VCVT", "MOVDW.U"},
}

// convert FP instructions from GNU syntax to Plan 9 syntax, for example,
// vadd.f32 s0, s3, s4 -> ADDF F0, S3, F2
// vsub.f64 d0, d2, d4 -> SUBD F0, F2, F4
// vldr s2, [r11] -> MOVF (R11), F1
// inputs: instruction name and arguments in GNU syntax
// return values: corresponding instruction name and arguments in Plan 9 syntax
func fpTrans(inst *Inst, op string, args []string) (string, []string) {
	for _, fp := range fpInst {
		if inst.Op == fp.op {
			// remove gnu syntax suffixes
			op = strings.Replace(op, ".F32", "", -1)
			op = strings.Replace(op, ".F64", "", -1)
			op = strings.Replace(op, ".S32", "", -1)
			op = strings.Replace(op, ".U32", "", -1)
			op = strings.Replace(op, ".32", "", -1)
			// compose op name
			if fp.op == VLDR || fp.op == VSTR {
				switch {
				case strings.HasPrefix(args[fp.transArgs[0]], "D"):
					op = "MOVD" + op[len(fp.gnuName):]
				case strings.HasPrefix(args[fp.transArgs[0]], "S"):
					op = "MOVF" + op[len(fp.gnuName):]
				default:
					panic(fmt.Sprintf("wrong FP register: %v", inst))
				}
			} else {
				op = fp.goName + op[len(fp.gnuName):]
			}
			// transform registers
			for ix, ri := range fp.transArgs {
				switch {
				case strings.HasSuffix(args[ri], "[1]"): // MOVW Rx, Dy[1]
					break
				case strings.HasSuffix(args[ri], "[0]"): // Dx[0] -> Fx
					args[ri] = strings.Replace(args[ri], "[0]", "", -1)
					fallthrough
				case strings.HasPrefix(args[ri], "D"): // Dx -> Fx
					args[ri] = "F" + args[ri][1:]
				case strings.HasPrefix(args[ri], "S"):
					if inst.Args[ix].(Reg)&1 == 0 { // Sx -> Fy, y = x/2, if x is even
						args[ri] = fmt.Sprintf("F%d", (inst.Args[ix].(Reg)-S0)/2)
					}
				case strings.HasPrefix(args[ri], "$"): // CMPF/CMPD $0, Fx
					break
				case strings.HasPrefix(args[ri], "R"): // MOVW Rx, Dy[1]
					break
				default:
					panic(fmt.Sprintf("wrong FP register: %v", inst))
				}
			}
			break
		}
	}
	return op, args
}
