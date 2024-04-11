package thumbasm

const (
	_ Op = iota
	ADD
	ADD_S
	ADC
	ADC_S
	AND
	AND_S
	ASR
	ASR_S
	B
	BCC
	BCS
	BEQ
	BFC
	BFI
	BGE
	BGT
	BHI
	BIC
	BIC_S
	BKPT
	BL
	BLE
	BLS
	BLT
	BLX
	BMI
	BNE
	BPL
	BVC
	BVS
	BX
	CBNZ
	CBZ
	CLREX
	CLZ
	CMN
	CMP
	CPSID
	CPSIE
	DMB
	DSB
	EOR
	EOR_S
	ISB
	IT
	ITE
	ITEE
	ITEEE
	ITEET
	ITET
	ITETE
	ITETT
	ITT
	ITTE
	ITTEE
	ITTET
	ITTT
	ITTTE
	ITTTT
	LDMIA
	LDMDB
	LDR
	LDRB
	LDRD
	LDREX
	LDREXB
	LDREXH
	LDRH
	LDRSB
	LDRSH
	LSL
	LSL_S
	LSR
	LSR_S
	MLA
	MLS
	MOV
	MOV_S
	MOVT
	MOVW
	MRS
	MSR
	MVN
	MVN_S
	MUL
	NOP
	ORN
	ORN_S
	ORR
	ORR_S
	PLD
	PLI
	POP
	PUSH
	RBIT
	REV
	REV16
	REVSH
	ROR
	ROR_S
	RSB
	RSB_S
	SBC
	SBC_S
	SBFX
	SDIV
	SEL
	SEV
	SMLABB
	SMLABT
	SMLAD
	SMLADX
	SMLAL
	SMLATB
	SMLATT
	SMLAWB
	SMLAWT
	SMLSD
	SMLSDX
	SMMLA
	SMMLS
	SMULL
	STMIA
	STMDB
	STR
	STRB
	STRD
	STREX
	STREXB
	STREXH
	STRH
	SUB
	SUB_S
	SVC
	SXTB
	SXTH
	TBB
	TBH
	TST
	TEQ
	UBFX
	UDF
	UDIV
	UMLAL
	UMULL
	UXTB
	UXTH

	VABS_F32
	VABS_F64
	VADD_F32
	VADD_F64
	VCMP_F32
	VCMP_F64
	VCMPE_F32
	VCMPE_F64
	VCVT_F32_FXS16
	VCVT_F32_FXS32
	VCVT_F32_FXU16
	VCVT_F32_FXU32
	VCVT_F64_FXS16
	VCVT_F64_FXS32
	VCVT_F64_FXU16
	VCVT_F64_FXU32
	VCVT_F32_U32
	VCVT_F32_S32
	VCVT_F64_U32
	VCVT_F64_S32
	VCVT_F64_F32
	VCVT_F32_F64
	VCVT_FXS16_F32
	VCVT_FXS16_F64
	VCVT_FXS32_F32
	VCVT_FXS32_F64
	VCVT_FXU16_F32
	VCVT_FXU16_F64
	VCVT_FXU32_F32
	VCVT_FXU32_F64
	VCVTB_F32_F16
	VCVTB_F16_F32
	VCVTT_F32_F16
	VCVTT_F16_F32
	VCVTR_U32_F32
	VCVTR_U32_F64
	VCVTR_S32_F32
	VCVTR_S32_F64
	VCVT_U32_F32
	VCVT_U32_F64
	VCVT_S32_F32
	VCVT_S32_F64
	VDIV_F32
	VDIV_F64
	VLDR
	VMLA_F32
	VMLA_F64
	VMLS_F32
	VMLS_F64
	VMOV
	VMOV_32
	VMOV_F32
	VMOV_F64
	VMRS
	VMSR
	VMUL_F32
	VMUL_F64
	VNEG_F32
	VNEG_F64
	VNMLS_F32
	VNMLS_F64
	VNMLA_F32
	VNMLA_F64
	VNMUL_F32
	VNMUL_F64
	VSQRT_F32
	VSQRT_F64
	VSTR
	VSUB_F32
	VSUB_F64

	WFE
	WFI
	YIELD
)

var opstr = [...]string{
	ADD:            "ADD",
	ADD_S:          "ADD.S",
	ADC:            "ADC",
	ADC_S:          "ADC.S",
	AND:            "AND",
	AND_S:          "AND.S",
	ASR:            "ASR",
	ASR_S:          "ASR.S",
	B:              "B",
	BCC:            "BCC",
	BCS:            "BCS",
	BEQ:            "BEQ",
	BFC:            "BFC",
	BFI:            "BFI",
	BGE:            "BGE",
	BGT:            "BGT",
	BHI:            "BHI",
	BIC:            "BIC",
	BIC_S:          "BIC.S",
	BKPT:           "BKPT",
	BL:             "BL",
	BLE:            "BLE",
	BLS:            "BLS",
	BLT:            "BLT",
	BLX:            "BLX",
	BMI:            "BMI",
	BNE:            "BNE",
	BPL:            "BPL",
	BVC:            "BVC",
	BVS:            "BVS",
	BX:             "BX",
	CBNZ:           "CBNZ",
	CBZ:            "CBZ",
	CLREX:          "CLREX",
	CLZ:            "CLZ",
	CMN:            "CMN",
	CMP:            "CMP",
	CPSID:          "CPSID",
	CPSIE:          "CPSIE",
	DMB:            "DMB",
	DSB:            "DSB",
	EOR:            "EOR",
	EOR_S:          "EOR.S",
	ISB:            "ISB",
	IT:             "IT",
	ITE:            "ITE",
	ITEE:           "ITEE",
	ITEEE:          "ITEEE",
	ITEET:          "ITEET",
	ITET:           "ITET",
	ITETE:          "ITETE",
	ITETT:          "ITETT",
	ITT:            "ITT",
	ITTE:           "ITTE",
	ITTEE:          "ITTEE",
	ITTET:          "ITTET",
	ITTT:           "ITTT",
	ITTTE:          "ITTTE",
	ITTTT:          "ITTTT",
	LDMIA:          "LDMIA",
	LDMDB:          "LDMDB",
	LDR:            "LDR",
	LDRB:           "LDRB",
	LDRD:           "LDRD",
	LDREX:          "LDREX",
	LDREXB:         "LDREXB",
	LDREXH:         "LDREXH",
	LDRH:           "LDRH",
	LDRSB:          "LDRSB",
	LDRSH:          "LDRSH",
	LSL:            "LSL",
	LSL_S:          "LSL.S",
	LSR:            "LSR",
	LSR_S:          "LSR.S",
	MLA:            "MLA",
	MLS:            "MLS",
	MOV:            "MOV",
	MOV_S:          "MOV.S",
	MOVT:           "MOVT",
	MOVW:           "MOVW",
	MRS:            "MRS",
	MSR:            "MSR",
	MVN:            "MVN",
	MVN_S:          "MVN.S",
	MUL:            "MUL",
	NOP:            "NOP",
	ORN:            "ORN",
	ORN_S:          "ORN.S",
	ORR:            "ORR",
	ORR_S:          "ORR.S",
	PLD:            "PLD",
	PLI:            "PLI",
	POP:            "POP",
	PUSH:           "PUSH",
	RBIT:           "RBIT",
	REV:            "REV",
	REV16:          "REV16",
	REVSH:          "REVSH",
	ROR:            "ROR",
	ROR_S:          "ROR.S",
	RSB:            "RSB",
	RSB_S:          "RSB.S",
	SBC:            "SBC",
	SBC_S:          "SBC.S",
	SBFX:           "SBFX",
	SDIV:           "SDIV",
	SEL:            "SEL",
	SEV:            "SEV",
	SMLABB:         "SMLABB",
	SMLABT:         "SMLABT",
	SMLAD:          "SMLAD",
	SMLADX:         "SMLADX",
	SMLAL:          "SMLAL",
	SMLATB:         "SMLATB",
	SMLATT:         "SMLATT",
	SMLAWB:         "SMLAWB",
	SMLAWT:         "SMLAWT",
	SMLSD:          "SMLSD",
	SMLSDX:         "SMLSDX",
	SMMLA:          "SMMLA",
	SMMLS:          "SMMLS",
	SMULL:          "SMULL",
	STMIA:          "STMIA",
	STMDB:          "STMDB",
	STR:            "STR",
	STRB:           "STRB",
	STRD:           "STRD",
	STREX:          "STREX",
	STREXB:         "STREXB",
	STREXH:         "STREXH",
	STRH:           "STRH",
	SUB:            "SUB",
	SUB_S:          "SUB.S",
	SVC:            "SVC",
	SXTB:           "SXTB",
	SXTH:           "SXTH",
	TBB:            "TBB",
	TBH:            "TBH",
	TST:            "TST",
	TEQ:            "TEQ",
	UBFX:           "UBFX",
	UDF:            "UDF",
	UDIV:           "UDIV",
	UMLAL:          "UMLAL",
	UMULL:          "UMULL",
	UXTB:           "UXTB",
	UXTH:           "UXTH",
	VABS_F32:       "VABS.F32",
	VABS_F64:       "VABS.F64",
	VADD_F32:       "VADD.F32",
	VADD_F64:       "VADD.F64",
	VCMP_F32:       "VCMP.F32",
	VCMP_F64:       "VCMP.F64",
	VCMPE_F32:      "VCMPE.F32",
	VCMPE_F64:      "VCMPE.F64",
	VCVT_F32_FXS16: "VCVT.F32.FXS16",
	VCVT_F32_FXS32: "VCVT.F32.FXS32",
	VCVT_F32_FXU16: "VCVT.F32.FXU16",
	VCVT_F32_FXU32: "VCVT.F32.FXU32",
	VCVT_F64_FXS16: "VCVT.F64.FXS16",
	VCVT_F64_FXS32: "VCVT.F64.FXS32",
	VCVT_F64_FXU16: "VCVT.F64.FXU16",
	VCVT_F64_FXU32: "VCVT.F64.FXU32",
	VCVT_F32_U32:   "VCVT.F32.U32",
	VCVT_F32_S32:   "VCVT.F32.S32",
	VCVT_F64_U32:   "VCVT.F64.U32",
	VCVT_F64_S32:   "VCVT.F64.S32",
	VCVT_F64_F32:   "VCVT.F64.F32",
	VCVT_F32_F64:   "VCVT.F32.F64",
	VCVT_FXS16_F32: "VCVT.FXS16.F32",
	VCVT_FXS16_F64: "VCVT.FXS16.F64",
	VCVT_FXS32_F32: "VCVT.FXS32.F32",
	VCVT_FXS32_F64: "VCVT.FXS32.F64",
	VCVT_FXU16_F32: "VCVT.FXU16.F32",
	VCVT_FXU16_F64: "VCVT.FXU16.F64",
	VCVT_FXU32_F32: "VCVT.FXU32.F32",
	VCVT_FXU32_F64: "VCVT.FXU32.F64",
	VCVTB_F32_F16:  "VCVTB.F32.F16",
	VCVTB_F16_F32:  "VCVTB.F16.F32",
	VCVTT_F32_F16:  "VCVTT.F32.F16",
	VCVTT_F16_F32:  "VCVTT.F16.F32",
	VCVTR_U32_F32:  "VCVTR.U32.F32",
	VCVTR_U32_F64:  "VCVTR.U32.F64",
	VCVTR_S32_F32:  "VCVTR.S32.F32",
	VCVTR_S32_F64:  "VCVTR.S32.F64",
	VCVT_U32_F32:   "VCVT.U32.F32",
	VCVT_U32_F64:   "VCVT.U32.F64",
	VCVT_S32_F32:   "VCVT.S32.F32",
	VCVT_S32_F64:   "VCVT.S32.F64",
	VDIV_F32:       "VDIV.F32",
	VDIV_F64:       "VDIV.F64",
	VLDR:           "VLDR",
	VMLA_F32:       "VMLA.F32",
	VMLA_F64:       "VMLA.F64",
	VMLS_F32:       "VMLS.F32",
	VMLS_F64:       "VMLS.F64",
	VMOV:           "VMOV",
	VMOV_32:        "VMOV.32",
	VMOV_F32:       "VMOV.F32",
	VMOV_F64:       "VMOV.F64",
	VMRS:           "VMRS",
	VMSR:           "VMSR",
	VMUL_F32:       "VMUL.F32",
	VMUL_F64:       "VMUL.F64",
	VNEG_F32:       "VNEG.F32",
	VNEG_F64:       "VNEG.F64",
	VNMLS_F32:      "VNMLS.F32",
	VNMLS_F64:      "VNMLS.F64",
	VNMLA_F32:      "VNMLA.F32",
	VNMLA_F64:      "VNMLA.F64",
	VNMUL_F32:      "VNMUL.F32",
	VNMUL_F64:      "VNMUL.F64",
	VSQRT_F32:      "VSQRT.F32",
	VSQRT_F64:      "VSQRT.F64",
	VSTR:           "VSTR",
	VSUB_F32:       "VSUB.F32",
	VSUB_F64:       "VSUB.F64",
	WFE:            "WFE",
	WFI:            "WFI",
	YIELD:          "YIELD",
}

var inst16formats = [...]inst16format{
	{0xFFFF, 0xBF00, NOP, nil},
	{0xFFFF, 0xBF10, YIELD, nil},
	{0xFFFF, 0xBF20, WFE, nil},
	{0xFFFF, 0xBF30, WFI, nil},
	{0xFFFF, 0xBF40, SEV, nil},

	{0xFFFC, 0xB660, CPSIE, _CPSIE},
	{0xFFFC, 0xB670, CPSID, _CPSIE},

	{0xFF0F, 0xBF00 | 8, IT, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 0 | 4, ITT, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 8 | 4, ITE, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 0 | 0 | 2, ITTT, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 8 | 0 | 2, ITET, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 0 | 4 | 2, ITTE, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 8 | 4 | 2, ITEE, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 0 | 0 | 0 | 1, ITTTT, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 8 | 0 | 0 | 1, ITETT, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 0 | 4 | 0 | 1, ITTET, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 8 | 4 | 0 | 1, ITEET, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 0 | 0 | 2 | 1, ITTTE, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 8 | 0 | 2 | 1, ITETE, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 0 | 4 | 2 | 1, ITTEE, _ITmask__firstcond},
	{0xFF1F, 0xBF00 | 8 | 4 | 2 | 1, ITEEE, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 0 | 4, ITE, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 8 | 4, ITT, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 0 | 0 | 2, ITEE, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 8 | 0 | 2, ITTE, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 0 | 4 | 2, ITET, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 8 | 4 | 2, ITTT, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 0 | 0 | 0 | 1, ITEEE, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 8 | 0 | 0 | 1, ITTEE, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 0 | 4 | 0 | 1, ITETE, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 8 | 4 | 0 | 1, ITTTE, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 0 | 0 | 2 | 1, ITEET, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 8 | 0 | 2 | 1, ITTET, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 0 | 4 | 2 | 1, ITETT, _ITmask__firstcond},
	{0xFF1F, 0xBF10 | 8 | 4 | 2 | 1, ITTTT, _ITmask__firstcond},

	{0xFF87, 0x4700, BX, _B__Rm},
	{0xFF87, 0x4780, BLX, _B__Rm},

	{0xFFC0, 0x0000, MOV_S, _AND__Rm__Rdn},
	{0xFFC0, 0x4000, AND, _AND__Rm__Rdn},
	{0xFFC0, 0x4040, EOR, _AND__Rm__Rdn},
	{0xFFC0, 0x4140, ADC, _AND__Rm__Rdn},
	{0xFFC0, 0x4180, SBC, _AND__Rm__Rdn},
	{0xFFC0, 0x4200, TST, _AND__Rm__Rdn},
	{0xFFC0, 0x4240, RSB, _AND__Rm__Rdn},
	{0xFFC0, 0x4280, CMP, _AND__Rm__Rdn},
	{0xFFC0, 0x42C0, CMN, _AND__Rm__Rdn},
	{0xFFC0, 0x4300, ORR, _AND__Rm__Rdn},
	{0xFFC0, 0x4340, MUL, _AND__Rm__Rdn},
	{0xFFC0, 0x4380, BIC, _AND__Rm__Rdn},
	{0xFFC0, 0x43C0, MVN, _AND__Rm__Rdn},
	{0xFFC0, 0x4080, LSL, _AND__Rm__Rdn},
	{0xFFC0, 0x40C0, LSR, _AND__Rm__Rdn},
	{0xFFC0, 0x4100, ASR, _AND__Rm__Rdn},
	{0xFFC0, 0x41C0, ROR, _AND__Rm__Rdn},
	{0xFFC0, 0xBA00, REV, _AND__Rm__Rdn},
	{0xFFC0, 0xBA40, REV16, _AND__Rm__Rdn},
	{0xFFC0, 0xBAC0, REVSH, _AND__Rm__Rdn},
	{0xFFC0, 0xB200, SXTH, _AND__Rm__Rdn},
	{0xFFC0, 0xB240, SXTB, _AND__Rm__Rdn},
	{0xFFC0, 0xB280, UXTH, _AND__Rm__Rdn},
	{0xFFC0, 0xB2C0, UXTB, _AND__Rm__Rdn},

	{0xFF80, 0xB000, ADD, _ADD__u7_2__R13},
	{0xFF80, 0xB080, SUB, _ADD__u7_2__R13},

	{0xFF00, 0x4400, ADD, _ADD__Rm__Rdn},
	{0xFF00, 0x4600, MOV, _ADD__Rm__Rdn},
	{0xFF00, 0x4500, CMP, _ADD__Rm__Rdn},

	{0xFF00, 0xD000, BEQ, _Bcond__i8_1},
	{0xFF00, 0xD100, BNE, _Bcond__i8_1},
	{0xFF00, 0xD200, BCS, _Bcond__i8_1},
	{0xFF00, 0xD300, BCC, _Bcond__i8_1},
	{0xFF00, 0xD400, BMI, _Bcond__i8_1},
	{0xFF00, 0xD500, BPL, _Bcond__i8_1},
	{0xFF00, 0xD600, BVS, _Bcond__i8_1},
	{0xFF00, 0xD700, BVC, _Bcond__i8_1},
	{0xFF00, 0xD800, BHI, _Bcond__i8_1},
	{0xFF00, 0xD900, BLS, _Bcond__i8_1},
	{0xFF00, 0xDA00, BGE, _Bcond__i8_1},
	{0xFF00, 0xDB00, BLT, _Bcond__i8_1},
	{0xFF00, 0xDC00, BGT, _Bcond__i8_1},
	{0xFF00, 0xDD00, BLE, _Bcond__i8_1},

	{0xFF00, 0xDE00, UDF, _UDF__u8},
	{0xFF00, 0xDF00, SVC, _UDF__u8},
	{0xFF00, 0xBE00, BKPT, _UDF__u8},

	{0xFE00, 0x1800, ADD, _ADD__Rm__Rn__Rd},
	{0xFE00, 0x1A00, SUB, _ADD__Rm__Rn__Rd},

	{0xFE00, 0xB400, PUSH, _PUSH__reglist},
	{0xFE00, 0xBC00, POP, _PUSH__reglist},

	{0xFE00, 0x1C00, ADD, _ADD__u3__Rn__Rd},
	{0xFE00, 0x1E00, SUB, _ADD__u3__Rn__Rd},

	{0xFD00, 0xB100, CBZ, _CBZ__Rn__u6_1},
	{0xFD00, 0xB900, CBNZ, _CBZ__Rn__u6_1},

	{0xF800, 0x0000, LSL, _MOVW__Rm_v_u5__Rd},
	{0xF800, 0x0800, LSR, _MOVW__Rm_v_u5__Rd},
	{0xF800, 0x1000, ASR, _MOVW__Rm_v_u5__Rd},

	{0xF800, 0x2000, MOV, _MOVW__u8__Rd},
	{0xF800, 0x2800, CMP, _MOVW__u8__Rd},
	{0xF800, 0x3000, ADD, _MOVW__u8__Rd},
	{0xF800, 0x3800, SUB, _MOVW__u8__Rd},

	{0xF800, 0x4800, LDR, _MOVW__u8_2_R13__Rt},
	{0xF800, 0x9000, STR, _MOVW__u8_2_R13__Rt},
	{0xF800, 0x9800, LDR, _MOVW__u8_2_R13__Rt},

	{0xF800, 0x6000, STR, _MOVW__u5_2_Rn__Rt},
	{0xF800, 0x7000, STRB, _MOVW__u5_2_Rn__Rt},
	{0xF800, 0x8000, STRH, _MOVW__u5_2_Rn__Rt},
	{0xF800, 0x6800, LDR, _MOVW__u5_2_Rn__Rt},
	{0xF800, 0x7800, LDRB, _MOVW__u5_2_Rn__Rt},
	{0xF800, 0x8800, LDRH, _MOVW__u5_2_Rn__Rt},

	{0xF800, 0xE000, B, _B__i11_1},

	{0xF800, 0xC000, STMIA, _MOVM_IAW},
	{0xF800, 0xC800, LDMIA, _MOVM_IAW},

	{0xF000, 0xA000, ADD, _ADD__u8_2__R13__Rd},

	{0xFE00, 0x5800, LDR, _MOVW__Rn_Rm__Rt},
	{0xFE00, 0x5E00, LDRSH, _MOVW__Rn_Rm__Rt},
	{0xFE00, 0x5A00, LDRH, _MOVW__Rn_Rm__Rt},
	{0xFE00, 0x5600, LDRSB, _MOVW__Rn_Rm__Rt},
	{0xFE00, 0x5C00, LDRB, _MOVW__Rn_Rm__Rt},
	{0xFE00, 0x5000, STR, _MOVW__Rn_Rm__Rt},
	{0xFE00, 0x5200, STRH, _MOVW__Rn_Rm__Rt},
	{0xFE00, 0x5400, STRB, _MOVW__Rn_Rm__Rt},
}

var inst32formats = [...]inst32format{
	{0xFFFFFFFF, 0xF3AF8000, NOP, nil},
	{0xFFFFFFFF, 0xF3BF8F2F, CLREX, nil},

	{0xFFFF0FFF, 0xF84D0D04, PUSH, _PUSH__Rt},
	{0xFFFF0FFF, 0xF85D0B04, POP, _PUSH__Rt},

	{0xFFFFFFF0, 0xF3BF8F40, DSB, _DSB__opt},
	{0xFFFFFFF0, 0xF3BF8F50, DMB, _DSB__opt},
	{0xFFFFFFF0, 0xF3BF8F60, ISB, _DSB__opt},

	{0xFFFF0FFF, 0xEEE10A10, VMSR, _MOVW__FPSCR__Rt},
	{0xFFFF0FFF, 0xEEF10A10, VMRS, _MOVW__FPSCR__Rt},

	{0xFFF0FFF0, 0xE8D0F000, TBB, _TBB__Rm__Rn},
	{0xFFF0FFF0, 0xE8D0F010, TBH, _TBB__Rm__Rn},

	{0xFFF0F0F0, 0xFB00F000, MUL, _MUL__Rm__Rn__Rd},
	{0xFFF0F0F0, 0xFB90F0F0, SDIV, _MUL__Rm__Rn__Rd},
	{0xFFF0F0F0, 0xFBB0F0F0, UDIV, _MUL__Rm__Rn__Rd},
	{0xFFF0F0F0, 0xFAA0F080, SEL, _MUL__Rm__Rn__Rd},

	{0xFFF0F0F0, 0xFAB0F080, CLZ, _CLZ__Rm__Rd},
	{0xFFF0F0F0, 0xFA90F080, REV, _CLZ__Rm__Rd},
	{0xFFF0F0F0, 0xFA90F090, REV16, _CLZ__Rm__Rd},
	{0xFFF0F0F0, 0xFA90F0A0, RBIT, _CLZ__Rm__Rd},
	{0xFFF0F0F0, 0xFA90F0B0, REVSH, _CLZ__Rm__Rd},

	{0xFFF0F0F0, 0xFA00F000, LSL, _MOVWs__Rn_v_Rm__Rd},
	{0xFFF0F0F0, 0xFA10F000, LSL_S, _MOVWs__Rn_v_Rm__Rd},
	{0xFFF0F0F0, 0xFA20F000, LSR, _MOVWs__Rn_v_Rm__Rd},
	{0xFFF0F0F0, 0xFA30F000, LSR_S, _MOVWs__Rn_v_Rm__Rd},
	{0xFFF0F0F0, 0xFA40F000, ASR, _MOVWs__Rn_v_Rm__Rd},
	{0xFFF0F0F0, 0xFA50F000, ASR_S, _MOVWs__Rn_v_Rm__Rd},
	{0xFFF0F0F0, 0xFA60F000, ROR, _MOVWs__Rn_v_Rm__Rd},
	{0xFFF0F0F0, 0xFA70F000, ROR_S, _MOVWs__Rn_v_Rm__Rd},

	{0xFFF00FFF, 0xE8D00F4F, LDREXB, _LDREXB__Rn__Rt},
	{0xFFF00FFF, 0xE8D00F5F, LDREXB, _LDREXB__Rn__Rt},

	{0xFFF00FF0, 0xE8C00F40, STREXB, _STREXB__Rn__Rt},
	{0xFFF00FF0, 0xE8C00F50, STREXH, _STREXB__Rn__Rt},

	{0xFFFFF0C0, 0xFA0FF080, SXTH, _MOVH__Rm_rot__Rd},
	{0xFFFFF0C0, 0xFA4FF080, SXTB, _MOVH__Rm_rot__Rd},
	{0xFFFFF0C0, 0xFA1FF080, UXTH, _MOVH__Rm_rot__Rd},
	{0xFFFFF0C0, 0xFA5FF080, UXTB, _MOVH__Rm_rot__Rd},

	{0xFFFFF000, 0xF3EF8000, MRS, _MOVW__SYSm__Rd},
	{0xFFF0FF00, 0xF3808800, MSR, _MOVW__SYSm__Rd},

	{0xFFF08F00, 0xEA100F00, TST, _TST__Rm_v_u5__Rn},
	{0xFFF08F00, 0xEA900F00, TEQ, _TST__Rm_v_u5__Rn},
	{0xFFF08F00, 0xEB100F00, CMN, _TST__Rm_v_u5__Rn},
	{0xFFF08F00, 0xEBB00F00, CMP, _TST__Rm_v_u5__Rn},

	{0xFFFF8000, 0xEA4F0000, MOV, _MOVWs__Rm_v_u5__Rn},
	{0xFFFF8000, 0xEA5F0000, MOV_S, _MOVWs__Rm_v_u5__Rn},
	{0xFFFF8000, 0xEA6F0000, MVN, _MOVWs__Rm_v_u5__Rn},
	{0xFFFF8000, 0xEA7F0000, MVN_S, _MOVWs__Rm_v_u5__Rn},

	{0xFFF000F0, 0xFB800000, SMULL, _MULL__Rm__Rn__Rdh_Rdl},
	{0xFFF000F0, 0xFBA00000, UMULL, _MULL__Rm__Rn__Rdh_Rdl},
	{0xFFF000F0, 0xFBC00000, SMLAL, _MULL__Rm__Rn__Rdh_Rdl},
	{0xFFF000F0, 0xFBE00000, UMLAL, _MULL__Rm__Rn__Rdh_Rdl},
	{0xFFF000F0, 0xFB000000, MLA, _MULL__Rm__Rn__Rdh_Rdl},
	{0xFFF000F0, 0xFB000010, MLS, _MULL__Rm__Rn__Rdh_Rdl},
	{0xFFF000F0, 0xFB300000, SMLAWB, _MULL__Rm__Rn__Rdh_Rdl},
	{0xFFF000F0, 0xFB300010, SMLAWT, _MULL__Rm__Rn__Rdh_Rdl},

	{0xFF7F0000, 0xF81F0000, LDRB, _MOVW__s12_Rn__Rt},
	{0xFF7F0000, 0xF83F0000, LDRH, _MOVW__s12_Rn__Rt},
	{0xFF7F0000, 0xF85F0000, LDR, _MOVW__s12_Rn__Rt},
	{0xFF7F0000, 0xF91F0000, LDRSB, _MOVW__s12_Rn__Rt},
	{0xFF7F0000, 0xF93F0000, LDRSH, _MOVW__s12_Rn__Rt},

	{0xFFF00FC0, 0xF8500000, LDR, _MOVW__Rn_Rm_1_u2__Rt},
	{0xFFF00FC0, 0xF9300000, LDRSH, _MOVW__Rn_Rm_1_u2__Rt},
	{0xFFF00FC0, 0xF8300000, LDRH, _MOVW__Rn_Rm_1_u2__Rt},
	{0xFFF00FC0, 0xF9100000, LDRSB, _MOVW__Rn_Rm_1_u2__Rt},
	{0xFFF00FC0, 0xF8100000, LDRB, _MOVW__Rn_Rm_1_u2__Rt},
	{0xFFF00FC0, 0xF8400000, STR, _MOVW__Rn_Rm_1_u2__Rt},
	{0xFFF00FC0, 0xF8200000, STRH, _MOVW__Rn_Rm_1_u2__Rt},
	{0xFFF00FC0, 0xF8000000, STRB, _MOVW__Rn_Rm_1_u2__Rt},

	{0xFBFF8000, 0xF04F0000, MOV, _MOVWs__e32__Rd},
	{0xFBFF8000, 0xF05F0000, MOV_S, _MOVWs__e32__Rd},
	{0xFBFF8000, 0xF06F0000, MVN, _MOVWs__e32__Rd},
	{0xFBFF8000, 0xF07F0000, MVN_S, _MOVWs__e32__Rd},

	{0xFFF00800, 0xF8000800, STRB, _MOVWpw__s8_Rn__Rt},
	{0xFFF00800, 0xF8100800, LDRB, _MOVWpw__s8_Rn__Rt},
	{0xFFF00800, 0xF8200800, STRH, _MOVWpw__s8_Rn__Rt},
	{0xFFF00800, 0xF8300800, LDRH, _MOVWpw__s8_Rn__Rt},
	{0xFFF00800, 0xF8400800, STR, _MOVWpw__s8_Rn__Rt},
	{0xFFF00800, 0xF8500800, LDR, _MOVWpw__s8_Rn__Rt},
	{0xFFF00800, 0xF9100800, LDRSB, _MOVWpw__s8_Rn__Rt},
	{0xFFF00800, 0xF9300800, LDRSH, _MOVWpw__s8_Rn__Rt},

	{0xFFF08000, 0xEA000000, AND, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEA100000, AND_S, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEA200000, BIC, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEA300000, BIC_S, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEA400000, ORR, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEA500000, ORR_S, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEA600000, ORN, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEA700000, ORN_S, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEA800000, EOR, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEA900000, EOR_S, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEB000000, ADD, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEB100000, ADD_S, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEB400000, ADC, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEB500000, ADC_S, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEB600000, SBC, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEB700000, SBC_S, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEBA00000, SUB, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEBB00000, SUB_S, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEBC00000, RSB, _ANDs__Rm_v_u5__Rn__Rd},
	{0xFFF08000, 0xEBD00000, RSB_S, _ANDs__Rm_v_u5__Rn__Rd},

	{0xFFF00000, 0xF8900000, LDRB, _MOVW__s12_Rn__Rt},
	{0xFFF00000, 0xF8800000, STRB, _MOVW__s12_Rn__Rt},
	{0xFFF00000, 0xF8B00000, LDRH, _MOVW__s12_Rn__Rt},
	{0xFFF00000, 0xF8A00000, STRH, _MOVW__s12_Rn__Rt},
	{0xFFF00000, 0xF8D00000, LDR, _MOVW__s12_Rn__Rt},
	{0xFFF00000, 0xF8C00000, STR, _MOVW__s12_Rn__Rt},
	{0xFFF00000, 0xF9900000, LDRSB, _MOVW__s12_Rn__Rt},
	{0xFFF00000, 0xF9B00000, LDRSH, _MOVW__s12_Rn__Rt},

	{0xFFD02000, 0xE8900000, LDMIA, _MOVM_IAw},
	{0xFFD02000, 0xE9100000, LDMDB, _MOVM_IAw},
	{0xFFD02000, 0xE8800000, STMIA, _MOVM_IAw},
	{0xFFD02000, 0xE9000000, STMDB, _MOVM_IAw},

	{0xFBF08F00, 0xF0100F00, TST, _TST__e32__Rn},
	{0xFBF08F00, 0xF0900F00, TEQ, _TST__e32__Rn},
	{0xFBF08F00, 0xF1100F00, CMN, _TST__e32__Rn},
	{0xFBF08F00, 0xF1B00F00, CMP, _TST__e32__Rn},

	{0xFBF08000, 0xF0000000, AND, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF0100000, AND_S, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF0200000, BIC, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF0300000, BIC_S, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF0400000, ORR, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF0500000, ORR_S, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF0600000, ORN, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF0700000, ORN_S, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF0800000, EOR, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF0900000, EOR_S, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF1000000, ADD, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF1100000, ADD_S, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF1400000, ADC, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF1500000, ADC_S, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF1600000, SBC, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF1700000, SBC_S, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF1A00000, SUB, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF1B00000, SUB_S, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF1C00000, RSB, _ANDs__e32__Rn__Rd},
	{0xFBF08000, 0xF1D00000, RSB_S, _ANDs__e32__Rn__Rd},

	{0xFBF08000, 0xF2400000, MOVW, _MOVW__uyz16__Rd},
	{0xFBF08000, 0xF2C00000, MOVT, _MOVW__uyz16__Rd},

	{0xFBF08000, 0xF2000000, ADD, _ADD__u12__Rn__Rd},
	{0xFBF08000, 0xF2A00000, SUB, _ADD__u12__Rn__Rd},

	{0xF800D000, 0xF0009000, B, _B__ji24_1},
	{0xF800D000, 0xF000D000, BL, _B__ji24_1},

	{0xFFF00F00, 0xE8500F00, LDREX, _LDREX__u8_2_Rn__Rt},
	{0xFFF00000, 0xE8400000, STREX, _STREX__Rt__u8_2_Rn__Rd},

	{0xFFF08020, 0xF3400000, SBFX, _BFX__width__ulsb__Rn__Rd},
	{0xFFF08020, 0xF3C00000, UBFX, _BFX__width__ulsb__Rn__Rd},

	{0xFFFF8020, 0xF36F0000, BFC, _BFC__width__ulsb__Rd},
	{0xFFF08020, 0xF3600000, BFI, _BFC__width__ulsb__Rd},

	{0xFF300E00, 0xED100A00, VLDR, _MOVF__s8_2_Rn__Fd},
	{0xFF300E00, 0xED000A00, VSTR, _MOVF__s8_2_Rn__Fd},

	{0xFFBF0FFF, 0xEEB50A40, VCMP_F32, _CMPF__0__Fd},
	{0xFFBF0FFF, 0xEEB50B40, VCMP_F64, _CMPF__0__Fd},
	{0xFFBF0FFF, 0xEEB50AC0, VCMPE_F32, _CMPF__0__Fd},
	{0xFFBF0FFF, 0xEEB50BC0, VCMPE_F64, _CMPF__0__Fd},

	{0xFFBF0FD0, 0xEEB40A40, VCMP_F32, _CMPF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB40B40, VCMP_F64, _CMPF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB40AC0, VCMPE_F32, _CMPF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB40BC0, VCMPE_F64, _CMPF__Fm__Fd},

	{0xFFBF0FD0, 0xEEB10AC0, VSQRT_F32, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB10BC0, VSQRT_F64, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB10A40, VNEG_F32, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB10B40, VNEG_F64, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB00A40, VMOV_F32, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB00B40, VMOV_F64, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB00AC0, VABS_F32, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB00BC0, VABS_F64, _SQRTF__Fm__Fd},

	{0xFFBF0FD0, 0xEEB70AC0, VCVT_F64_F32, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB70BC0, VCVT_F32_F64, _SQRTF__Fm__Fd},

	// ARMv7-A
	{0xFFBF0FD0, 0xEEBD0AC0, VCVT_S32_F32, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEBC0AC0, VCVT_U32_F32, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEBD0BC0, VCVT_S32_F64, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEBC0BC0, VCVT_U32_F64, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB80AC0, VCVT_F32_S32, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB80A40, VCVT_F32_U32, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB80BC0, VCVT_F64_S32, _SQRTF__Fm__Fd},
	{0xFFBF0FD0, 0xEEB80B40, VCVT_F64_U32, _SQRTF__Fm__Fd},

	// ARMv7-M version of above (according to RM)
	//{0xFFBF0FD0, 0xFEBD0AC0, VCVT_S32_F32, _SQRTF__Fm__Fd},
	//{0xFFBF0FD0, 0xFEBC0AC0, VCVT_U32_F32, _SQRTF__Fm__Fd},
	//{0xFFBF0FD0, 0xFEBD0BC0, VCVT_S32_F64, _SQRTF__Fm__Fd},
	//{0xFFBF0FD0, 0xFEBC0BC0, VCVT_U32_F64, _SQRTF__Fm__Fd},
	//{0xFFBF0FD0, 0xFEB80AC0, VCVT_F32_S32, _SQRTF__Fm__Fd},
	//{0xFFBF0FD0, 0xFEB80A40, VCVT_F32_U32, _SQRTF__Fm__Fd},
	//{0xFFBF0FD0, 0xFEB80BC0, VCVT_F64_S32, _SQRTF__Fm__Fd},
	//{0xFFBF0FD0, 0xFEB80B40, VCVT_F64_U32, _SQRTF__Fm__Fd},

	{0xFFF00F7F, 0xEE000B10, VMOV, _MOVW__Fm__Rd},
	{0xFFF00F7F, 0xEE100B10, VMOV, _MOVW__Fm__Rd},

	{0xFFB00FF0, 0xEEB00A00, VMOV_F32, _MOVF__f8__Fd},
	{0xFFB00FF0, 0xEEB00B00, VMOV_F64, _MOVF__f8__Fd},

	{0xFFB00F50, 0xEE300A00, VADD_F32, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE300B00, VADD_F64, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE300A40, VSUB_F32, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE300B40, VSUB_F64, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE200A00, VMUL_F32, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE200B00, VMUL_F64, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE800A00, VDIV_F32, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE800B00, VDIV_F64, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE000A00, VMLA_F32, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE000B00, VMLA_F64, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE000A40, VMLS_F32, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE000B40, VMLS_F64, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE200A40, VNMUL_F32, _ADDF__Fm__Fn__Fd},
	{0xFFB00F50, 0xEE200B40, VNMUL_F64, _ADDF__Fm__Fn__Fd},
}
