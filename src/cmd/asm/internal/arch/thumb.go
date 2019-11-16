// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the Thumb
// instruction set, to minimize its interaction with the core of the
// assembler.

package arch

import (
	"strings"

	"cmd/internal/obj"
	"cmd/internal/obj/thumb"
)

var thumbLS = map[string]uint8{
	"U":  thumb.C_UBIT,
	"S":  thumb.C_SBIT,
	"W":  thumb.C_WBIT,
	"P":  thumb.C_PBIT,
	"PW": thumb.C_WBIT | thumb.C_PBIT,
	"WP": thumb.C_WBIT | thumb.C_PBIT,
}

var thumbSCOND = map[string]uint8{
	"EQ":  thumb.C_SCOND_EQ,
	"NE":  thumb.C_SCOND_NE,
	"CS":  thumb.C_SCOND_HS,
	"HS":  thumb.C_SCOND_HS,
	"CC":  thumb.C_SCOND_LO,
	"LO":  thumb.C_SCOND_LO,
	"MI":  thumb.C_SCOND_MI,
	"PL":  thumb.C_SCOND_PL,
	"VS":  thumb.C_SCOND_VS,
	"VC":  thumb.C_SCOND_VC,
	"HI":  thumb.C_SCOND_HI,
	"LS":  thumb.C_SCOND_LS,
	"GE":  thumb.C_SCOND_GE,
	"LT":  thumb.C_SCOND_LT,
	"GT":  thumb.C_SCOND_GT,
	"LE":  thumb.C_SCOND_LE,
	"AL":  thumb.C_SCOND_NONE,
	"U":   thumb.C_UBIT,
	"S":   thumb.C_SBIT,
	"W":   thumb.C_WBIT,
	"P":   thumb.C_PBIT,
	"PW":  thumb.C_WBIT | thumb.C_PBIT,
	"WP":  thumb.C_WBIT | thumb.C_PBIT,
	"F":   thumb.C_FBIT,
	"IAW": thumb.C_WBIT | thumb.C_UBIT,
	"DBW": thumb.C_WBIT | thumb.C_PBIT,
	"IA":  thumb.C_UBIT,
	"DB":  thumb.C_PBIT,
}

var thumbJump = map[string]bool{
	"B":    true,
	"BL":   true,
	"BEQ":  true,
	"BNE":  true,
	"BCS":  true,
	"BHS":  true,
	"BCC":  true,
	"BLO":  true,
	"BMI":  true,
	"BPL":  true,
	"BVS":  true,
	"BVC":  true,
	"BHI":  true,
	"BLS":  true,
	"BGE":  true,
	"BLT":  true,
	"BGT":  true,
	"BLE":  true,
	"CBZ":  true,
	"CBNZ": true,
	"CALL": true,
	"JMP":  true,
}

func jumpThumb(word string) bool {
	return thumbJump[word]
}

// IsThumbCMP reports whether the op (as defined by an thumb.A* constant) is
// one of the comparison instructions that require special handling.
func IsThumbCMP(op obj.As) bool {
	switch op {
	case thumb.ACMN, thumb.ACMP, thumb.ATEQ, thumb.ATST, thumb.ATBB, thumb.ATBH:
		return true
	}
	return false
}

// IsThumbSTREX reports whether the op (as defined by an thumb.A* constant) is
// one of the STREX-like instructions that require special handling.
func IsThumbSTREX(op obj.As) bool {
	switch op {
	case thumb.ASTREX, thumb.ASTREXB, thumb.ASTREXH:
		return true
	}
	return false
}

// IsThumbBFX reports whether the op (as defined by an thumb.A* constant) is one the
// BFX-like instructions which are in the form of "op $width, $LSB, (Reg,) Reg".
func IsThumbBFX(op obj.As) bool {
	switch op {
	case thumb.ABFX, thumb.ABFXU, thumb.ABFC, thumb.ABFI:
		return true
	}
	return false
}

// IsThumbFloatCmp reports whether the op is a floating comparison instruction.
func IsThumbFloatCmp(op obj.As) bool {
	switch op {
	case thumb.ACMPF, thumb.ACMPD:
		return true
	}
	return false
}

// IsThumbMULA reports whether the op (as defined by an thumb.A* constant) is
// MULA, MULS, MMULA, MMULS, MULABB, MULAWB or MULAWT, the 4-operand instructions.
func IsThumbMULA(op obj.As) bool {
	switch op {
	case thumb.AMULA, thumb.AMULS, thumb.AMMULA, thumb.AMMULS, thumb.AMULABB, thumb.AMULAWB, thumb.AMULAWT:
		return true
	}
	return false
}

var thumbBcode = []obj.As{
	thumb.ABEQ,
	thumb.ABNE,
	thumb.ABCS,
	thumb.ABCC,
	thumb.ABMI,
	thumb.ABPL,
	thumb.ABVS,
	thumb.ABVC,
	thumb.ABHI,
	thumb.ABLS,
	thumb.ABGE,
	thumb.ABLT,
	thumb.ABGT,
	thumb.ABLE,
	thumb.AB,
	obj.ANOP,
}

// ThumbConditionCodes handles the special condition code situation for the Thumb.
// It returns a boolean to indicate success; failure means cond was unrecognized.
func ThumbConditionCodes(prog *obj.Prog, cond string) bool {
	if cond == "" {
		return true
	}
	bits, ok := ParseThumbCondition(cond)
	if ok {
		prog.Scond = bits
	}
	return ok
}

// ParseThumbCondition parses the conditions attached to an Thumb instruction.
// The input is a single string consisting of period-separated condition
// codes, such as ".P.W". An initial period is ignored.
func ParseThumbCondition(cond string) (uint8, bool) {
	return parseThumbCondition(cond, thumbLS, thumbSCOND)
}

func parseThumbCondition(cond string, ls, scond map[string]uint8) (uint8, bool) {
	cond = strings.TrimPrefix(cond, ".")
	if cond == "" {
		return thumb.C_SCOND_NONE, true
	}
	names := strings.Split(cond, ".")
	bits := uint8(0)
	for _, name := range names {
		if b, present := ls[name]; present {
			bits |= b
			continue
		}
		if b, present := scond[name]; present {
			bits = (bits &^ thumb.C_SCOND) | b
			continue
		}
		return 0, false
	}
	return bits, true
}

func thumbRegisterNumber(name string, n int16) (int16, bool) {
	if n < 0 || 15 < n {
		return 0, false
	}
	switch name {
	case "R":
		return thumb.REG_R0 + n, true
	case "F":
		return thumb.REG_F0 + n, true
	}
	return 0, false
}
