// Inferno utils/5c/list.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/5c/list.c
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
	"cmd/internal/obj"
	"fmt"
)

func init() {
	obj.RegisterRegister(obj.RBaseThumb, MAXREG, rconv)
	obj.RegisterOpcode(obj.ABaseThumb, Anames)
	obj.RegisterRegisterList(obj.RegListARMLo, obj.RegListARMHi, rlconv)
	obj.RegisterOpSuffix("thumb", cconv)
}

var condCode = []string{
	".EQ",
	".NE",
	".CS",
	".CC",
	".MI",
	".PL",
	".VS",
	".VC",
	".HI",
	".LS",
	".GE",
	".LT",
	".GT",
	".LE",
	"",
	".NV",
}

func cconv(s uint8) string {
	sc := condCode[(s&C_SCOND)^C_SCOND_XOR]
	if s&C_SBIT != 0 {
		sc += ".S"
	}
	if s&C_PBIT != 0 {
		sc += ".P"
	}
	if s&C_WBIT != 0 {
		sc += ".W"
	}
	if s&C_UBIT != 0 { /* ambiguous with FBIT */
		sc += ".U"
	}
	return sc
}

func rconv(r int) string {
	if r == 0 {
		return "NONE"
	}
	if r == REGG {
		// Special case.
		return "g"
	}
	if REG_R0 <= r && r <= REG_R15 {
		return fmt.Sprintf("R%d", r-REG_R0)
	}
	if REG_F0 <= r && r <= REG_F15 {
		return fmt.Sprintf("F%d", r-REG_F0)
	}
	if REG_EQ <= r && r <= REG_AL {
		r = (r - REG_EQ) * 2
		return "EQNEHSLOMIPLVSVCHILSGELTGTLEAL"[r : r+2]
	}

	switch r {
	case REG_APSR:
		return "APSR"
	case REG_IAPSR:
		return "IAPSR"
	case REG_EAPSR:
		return "EAPSR"
	case REG_XPSR:
		return "XPSR"
	case REG_IPSR:
		return "IPSR"
	case REG_EPSR:
		return "EPSR"
	case REG_IEPSR:
		return "IEPSR"

	case REG_FPSCR:
		return "FPSCR"

	case REG_MSP:
		return "MSP"
	case REG_PSP:
		return "PSP"

	case REG_PRIMASK:
		return "PRIMASK"
	case REG_BASEPRI:
		return "BASEPRI"
	case REG_BASEPRI_MAX:
		return "BASEPRI_MAX"
	case REG_FAULTMASK:
		return "FAULTMASK"
	case REG_CONTROL:
		return "CONTROL"

	case REG_MB_SY:
		return "MB_SY"
	case REG_MB_ST:
		return "MB_ST"
	case REG_MB_ISH:
		return "MB_ISH"
	case REG_MB_ISHST:
		return "MB_ISHST"
	case REG_MB_NSH:
		return "MB_NSH"
	case REG_MB_NSHST:
		return "MB_NSHST"
	case REG_MB_OSH:
		return "MB_OSH"
	case REG_MB_OSHST:
		return "MB_OSHST"
	}

	return fmt.Sprintf("Rgok(%d)", r-obj.RBaseThumb)
}

func rlconv(list int64) string {
	str := ""
	for i := 0; i < 16; i++ {
		if list&(1<<uint(i)) != 0 {
			if str == "" {
				str += "["
			} else {
				str += ","
			}
			// This is ARM-specific; R10 is g.
			if i == REGG-REG_R0 {
				str += "g"
			} else {
				str += fmt.Sprintf("R%d", i)
			}
		}
	}

	str += "]"
	return str
}
