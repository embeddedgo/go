// Code generated from gen/Thumb.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "cmd/compile/internal/types"

func rewriteValueThumb(v *Value) bool {
	switch v.Op {
	case OpAbs:
		v.Op = OpThumbABSD
		return true
	case OpAdd16:
		v.Op = OpThumbADD
		return true
	case OpAdd32:
		v.Op = OpThumbADD
		return true
	case OpAdd32F:
		v.Op = OpThumbADDF
		return true
	case OpAdd32carry:
		v.Op = OpThumbADDS
		return true
	case OpAdd32withcarry:
		v.Op = OpThumbADC
		return true
	case OpAdd64F:
		v.Op = OpThumbADDD
		return true
	case OpAdd8:
		v.Op = OpThumbADD
		return true
	case OpAddPtr:
		v.Op = OpThumbADD
		return true
	case OpAddr:
		return rewriteValueThumb_OpAddr(v)
	case OpAnd16:
		v.Op = OpThumbAND
		return true
	case OpAnd32:
		v.Op = OpThumbAND
		return true
	case OpAnd8:
		v.Op = OpThumbAND
		return true
	case OpAndB:
		v.Op = OpThumbAND
		return true
	case OpAvg32u:
		return rewriteValueThumb_OpAvg32u(v)
	case OpBitLen32:
		return rewriteValueThumb_OpBitLen32(v)
	case OpBswap32:
		v.Op = OpThumbREV
		return true
	case OpClosureCall:
		v.Op = OpThumbCALLclosure
		return true
	case OpCom16:
		v.Op = OpThumbMVN
		return true
	case OpCom32:
		v.Op = OpThumbMVN
		return true
	case OpCom8:
		v.Op = OpThumbMVN
		return true
	case OpConst16:
		return rewriteValueThumb_OpConst16(v)
	case OpConst32:
		return rewriteValueThumb_OpConst32(v)
	case OpConst32F:
		return rewriteValueThumb_OpConst32F(v)
	case OpConst64F:
		return rewriteValueThumb_OpConst64F(v)
	case OpConst8:
		return rewriteValueThumb_OpConst8(v)
	case OpConstBool:
		return rewriteValueThumb_OpConstBool(v)
	case OpConstNil:
		return rewriteValueThumb_OpConstNil(v)
	case OpCtz16:
		return rewriteValueThumb_OpCtz16(v)
	case OpCtz16NonZero:
		v.Op = OpCtz32
		return true
	case OpCtz32:
		return rewriteValueThumb_OpCtz32(v)
	case OpCtz32NonZero:
		v.Op = OpCtz32
		return true
	case OpCtz8:
		return rewriteValueThumb_OpCtz8(v)
	case OpCtz8NonZero:
		v.Op = OpCtz32
		return true
	case OpCvt32Fto32:
		v.Op = OpThumbMOVFW
		return true
	case OpCvt32Fto32U:
		v.Op = OpThumbMOVFWU
		return true
	case OpCvt32Fto64F:
		v.Op = OpThumbMOVFD
		return true
	case OpCvt32Uto32F:
		v.Op = OpThumbMOVWUF
		return true
	case OpCvt32Uto64F:
		v.Op = OpThumbMOVWUD
		return true
	case OpCvt32to32F:
		v.Op = OpThumbMOVWF
		return true
	case OpCvt32to64F:
		v.Op = OpThumbMOVWD
		return true
	case OpCvt64Fto32:
		v.Op = OpThumbMOVDW
		return true
	case OpCvt64Fto32F:
		v.Op = OpThumbMOVDF
		return true
	case OpCvt64Fto32U:
		v.Op = OpThumbMOVDWU
		return true
	case OpCvtBoolToUint8:
		v.Op = OpCopy
		return true
	case OpDiv16:
		return rewriteValueThumb_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueThumb_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueThumb_OpDiv32(v)
	case OpDiv32F:
		v.Op = OpThumbDIVF
		return true
	case OpDiv32u:
		v.Op = OpThumbDIVU
		return true
	case OpDiv64F:
		v.Op = OpThumbDIVD
		return true
	case OpDiv8:
		return rewriteValueThumb_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueThumb_OpDiv8u(v)
	case OpEq16:
		return rewriteValueThumb_OpEq16(v)
	case OpEq32:
		return rewriteValueThumb_OpEq32(v)
	case OpEq32F:
		return rewriteValueThumb_OpEq32F(v)
	case OpEq64F:
		return rewriteValueThumb_OpEq64F(v)
	case OpEq8:
		return rewriteValueThumb_OpEq8(v)
	case OpEqB:
		return rewriteValueThumb_OpEqB(v)
	case OpEqPtr:
		return rewriteValueThumb_OpEqPtr(v)
	case OpFMA:
		return rewriteValueThumb_OpFMA(v)
	case OpGetCallerPC:
		v.Op = OpThumbLoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpThumbLoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpThumbLoweredGetClosurePtr
		return true
	case OpHmul32:
		v.Op = OpThumbHMUL
		return true
	case OpHmul32u:
		v.Op = OpThumbHMULU
		return true
	case OpInterCall:
		v.Op = OpThumbCALLinter
		return true
	case OpIsInBounds:
		return rewriteValueThumb_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValueThumb_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValueThumb_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValueThumb_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueThumb_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueThumb_OpLeq32(v)
	case OpLeq32F:
		return rewriteValueThumb_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValueThumb_OpLeq32U(v)
	case OpLeq64F:
		return rewriteValueThumb_OpLeq64F(v)
	case OpLeq8:
		return rewriteValueThumb_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueThumb_OpLeq8U(v)
	case OpLess16:
		return rewriteValueThumb_OpLess16(v)
	case OpLess16U:
		return rewriteValueThumb_OpLess16U(v)
	case OpLess32:
		return rewriteValueThumb_OpLess32(v)
	case OpLess32F:
		return rewriteValueThumb_OpLess32F(v)
	case OpLess32U:
		return rewriteValueThumb_OpLess32U(v)
	case OpLess64F:
		return rewriteValueThumb_OpLess64F(v)
	case OpLess8:
		return rewriteValueThumb_OpLess8(v)
	case OpLess8U:
		return rewriteValueThumb_OpLess8U(v)
	case OpLoad:
		return rewriteValueThumb_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueThumb_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueThumb_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueThumb_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueThumb_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueThumb_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueThumb_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueThumb_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueThumb_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueThumb_OpLsh32x8(v)
	case OpLsh8x16:
		return rewriteValueThumb_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueThumb_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueThumb_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueThumb_OpLsh8x8(v)
	case OpMMIOLoad16:
		return rewriteValueThumb_OpMMIOLoad16(v)
	case OpMMIOLoad32:
		return rewriteValueThumb_OpMMIOLoad32(v)
	case OpMMIOLoad8:
		return rewriteValueThumb_OpMMIOLoad8(v)
	case OpMMIOMB:
		v.Op = OpThumbDSB
		return true
	case OpMMIOStore16:
		return rewriteValueThumb_OpMMIOStore16(v)
	case OpMMIOStore32:
		return rewriteValueThumb_OpMMIOStore32(v)
	case OpMMIOStore8:
		return rewriteValueThumb_OpMMIOStore8(v)
	case OpMod16:
		return rewriteValueThumb_OpMod16(v)
	case OpMod16u:
		return rewriteValueThumb_OpMod16u(v)
	case OpMod32:
		return rewriteValueThumb_OpMod32(v)
	case OpMod32u:
		return rewriteValueThumb_OpMod32u(v)
	case OpMod8:
		return rewriteValueThumb_OpMod8(v)
	case OpMod8u:
		return rewriteValueThumb_OpMod8u(v)
	case OpMove:
		return rewriteValueThumb_OpMove(v)
	case OpMul16:
		v.Op = OpThumbMUL
		return true
	case OpMul32:
		v.Op = OpThumbMUL
		return true
	case OpMul32F:
		v.Op = OpThumbMULF
		return true
	case OpMul32uhilo:
		v.Op = OpThumbMULLU
		return true
	case OpMul64F:
		v.Op = OpThumbMULD
		return true
	case OpMul8:
		v.Op = OpThumbMUL
		return true
	case OpNeg16:
		return rewriteValueThumb_OpNeg16(v)
	case OpNeg32:
		return rewriteValueThumb_OpNeg32(v)
	case OpNeg32F:
		v.Op = OpThumbNEGF
		return true
	case OpNeg64F:
		v.Op = OpThumbNEGD
		return true
	case OpNeg8:
		return rewriteValueThumb_OpNeg8(v)
	case OpNeq16:
		return rewriteValueThumb_OpNeq16(v)
	case OpNeq32:
		return rewriteValueThumb_OpNeq32(v)
	case OpNeq32F:
		return rewriteValueThumb_OpNeq32F(v)
	case OpNeq64F:
		return rewriteValueThumb_OpNeq64F(v)
	case OpNeq8:
		return rewriteValueThumb_OpNeq8(v)
	case OpNeqB:
		v.Op = OpThumbXOR
		return true
	case OpNeqPtr:
		return rewriteValueThumb_OpNeqPtr(v)
	case OpNilCheck:
		v.Op = OpThumbLoweredNilCheck
		return true
	case OpNot:
		return rewriteValueThumb_OpNot(v)
	case OpOffPtr:
		return rewriteValueThumb_OpOffPtr(v)
	case OpOr16:
		v.Op = OpThumbOR
		return true
	case OpOr32:
		v.Op = OpThumbOR
		return true
	case OpOr8:
		v.Op = OpThumbOR
		return true
	case OpOrB:
		v.Op = OpThumbOR
		return true
	case OpPanicBounds:
		return rewriteValueThumb_OpPanicBounds(v)
	case OpPanicExtend:
		return rewriteValueThumb_OpPanicExtend(v)
	case OpPublicationBarrier:
		v.Op = OpThumbDMB_ST
		return true
	case OpRotateLeft16:
		return rewriteValueThumb_OpRotateLeft16(v)
	case OpRotateLeft32:
		return rewriteValueThumb_OpRotateLeft32(v)
	case OpRotateLeft8:
		return rewriteValueThumb_OpRotateLeft8(v)
	case OpRound32F:
		v.Op = OpCopy
		return true
	case OpRound64F:
		v.Op = OpCopy
		return true
	case OpRsh16Ux16:
		return rewriteValueThumb_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueThumb_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueThumb_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueThumb_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueThumb_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueThumb_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueThumb_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueThumb_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueThumb_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueThumb_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueThumb_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueThumb_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueThumb_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueThumb_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueThumb_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueThumb_OpRsh32x8(v)
	case OpRsh8Ux16:
		return rewriteValueThumb_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueThumb_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueThumb_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueThumb_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueThumb_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueThumb_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueThumb_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueThumb_OpRsh8x8(v)
	case OpSignExt16to32:
		v.Op = OpThumbMOVHreg
		return true
	case OpSignExt8to16:
		v.Op = OpThumbMOVBreg
		return true
	case OpSignExt8to32:
		v.Op = OpThumbMOVBreg
		return true
	case OpSignmask:
		return rewriteValueThumb_OpSignmask(v)
	case OpSlicemask:
		return rewriteValueThumb_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpThumbSQRTD
		return true
	case OpStaticCall:
		v.Op = OpThumbCALLstatic
		return true
	case OpStore:
		return rewriteValueThumb_OpStore(v)
	case OpSub16:
		v.Op = OpThumbSUB
		return true
	case OpSub32:
		v.Op = OpThumbSUB
		return true
	case OpSub32F:
		v.Op = OpThumbSUBF
		return true
	case OpSub32carry:
		v.Op = OpThumbSUBS
		return true
	case OpSub32withcarry:
		v.Op = OpThumbSBC
		return true
	case OpSub64F:
		v.Op = OpThumbSUBD
		return true
	case OpSub8:
		v.Op = OpThumbSUB
		return true
	case OpSubPtr:
		v.Op = OpThumbSUB
		return true
	case OpThumbADC:
		return rewriteValueThumb_OpThumbADC(v)
	case OpThumbADCconst:
		return rewriteValueThumb_OpThumbADCconst(v)
	case OpThumbADCshiftLL:
		return rewriteValueThumb_OpThumbADCshiftLL(v)
	case OpThumbADCshiftRA:
		return rewriteValueThumb_OpThumbADCshiftRA(v)
	case OpThumbADCshiftRL:
		return rewriteValueThumb_OpThumbADCshiftRL(v)
	case OpThumbADD:
		return rewriteValueThumb_OpThumbADD(v)
	case OpThumbADDD:
		return rewriteValueThumb_OpThumbADDD(v)
	case OpThumbADDF:
		return rewriteValueThumb_OpThumbADDF(v)
	case OpThumbADDS:
		return rewriteValueThumb_OpThumbADDS(v)
	case OpThumbADDSshiftLL:
		return rewriteValueThumb_OpThumbADDSshiftLL(v)
	case OpThumbADDSshiftRA:
		return rewriteValueThumb_OpThumbADDSshiftRA(v)
	case OpThumbADDSshiftRL:
		return rewriteValueThumb_OpThumbADDSshiftRL(v)
	case OpThumbADDconst:
		return rewriteValueThumb_OpThumbADDconst(v)
	case OpThumbADDshiftLL:
		return rewriteValueThumb_OpThumbADDshiftLL(v)
	case OpThumbADDshiftRA:
		return rewriteValueThumb_OpThumbADDshiftRA(v)
	case OpThumbADDshiftRL:
		return rewriteValueThumb_OpThumbADDshiftRL(v)
	case OpThumbAND:
		return rewriteValueThumb_OpThumbAND(v)
	case OpThumbANDconst:
		return rewriteValueThumb_OpThumbANDconst(v)
	case OpThumbANDshiftLL:
		return rewriteValueThumb_OpThumbANDshiftLL(v)
	case OpThumbANDshiftRA:
		return rewriteValueThumb_OpThumbANDshiftRA(v)
	case OpThumbANDshiftRL:
		return rewriteValueThumb_OpThumbANDshiftRL(v)
	case OpThumbBFX:
		return rewriteValueThumb_OpThumbBFX(v)
	case OpThumbBFXU:
		return rewriteValueThumb_OpThumbBFXU(v)
	case OpThumbBIC:
		return rewriteValueThumb_OpThumbBIC(v)
	case OpThumbBICconst:
		return rewriteValueThumb_OpThumbBICconst(v)
	case OpThumbBICshiftLL:
		return rewriteValueThumb_OpThumbBICshiftLL(v)
	case OpThumbBICshiftRA:
		return rewriteValueThumb_OpThumbBICshiftRA(v)
	case OpThumbBICshiftRL:
		return rewriteValueThumb_OpThumbBICshiftRL(v)
	case OpThumbCMN:
		return rewriteValueThumb_OpThumbCMN(v)
	case OpThumbCMNconst:
		return rewriteValueThumb_OpThumbCMNconst(v)
	case OpThumbCMNshiftLL:
		return rewriteValueThumb_OpThumbCMNshiftLL(v)
	case OpThumbCMNshiftRA:
		return rewriteValueThumb_OpThumbCMNshiftRA(v)
	case OpThumbCMNshiftRL:
		return rewriteValueThumb_OpThumbCMNshiftRL(v)
	case OpThumbCMOVWHSconst:
		return rewriteValueThumb_OpThumbCMOVWHSconst(v)
	case OpThumbCMOVWLSconst:
		return rewriteValueThumb_OpThumbCMOVWLSconst(v)
	case OpThumbCMP:
		return rewriteValueThumb_OpThumbCMP(v)
	case OpThumbCMPD:
		return rewriteValueThumb_OpThumbCMPD(v)
	case OpThumbCMPF:
		return rewriteValueThumb_OpThumbCMPF(v)
	case OpThumbCMPconst:
		return rewriteValueThumb_OpThumbCMPconst(v)
	case OpThumbCMPshiftLL:
		return rewriteValueThumb_OpThumbCMPshiftLL(v)
	case OpThumbCMPshiftRA:
		return rewriteValueThumb_OpThumbCMPshiftRA(v)
	case OpThumbCMPshiftRL:
		return rewriteValueThumb_OpThumbCMPshiftRL(v)
	case OpThumbDIV:
		return rewriteValueThumb_OpThumbDIV(v)
	case OpThumbDIVU:
		return rewriteValueThumb_OpThumbDIVU(v)
	case OpThumbEqual:
		return rewriteValueThumb_OpThumbEqual(v)
	case OpThumbGreaterEqual:
		return rewriteValueThumb_OpThumbGreaterEqual(v)
	case OpThumbGreaterEqualU:
		return rewriteValueThumb_OpThumbGreaterEqualU(v)
	case OpThumbGreaterThan:
		return rewriteValueThumb_OpThumbGreaterThan(v)
	case OpThumbGreaterThanU:
		return rewriteValueThumb_OpThumbGreaterThanU(v)
	case OpThumbLessEqual:
		return rewriteValueThumb_OpThumbLessEqual(v)
	case OpThumbLessEqualU:
		return rewriteValueThumb_OpThumbLessEqualU(v)
	case OpThumbLessThan:
		return rewriteValueThumb_OpThumbLessThan(v)
	case OpThumbLessThanU:
		return rewriteValueThumb_OpThumbLessThanU(v)
	case OpThumbLoadOnce16:
		return rewriteValueThumb_OpThumbLoadOnce16(v)
	case OpThumbLoadOnce16idx:
		return rewriteValueThumb_OpThumbLoadOnce16idx(v)
	case OpThumbLoadOnce32:
		return rewriteValueThumb_OpThumbLoadOnce32(v)
	case OpThumbLoadOnce32idx:
		return rewriteValueThumb_OpThumbLoadOnce32idx(v)
	case OpThumbLoadOnce8:
		return rewriteValueThumb_OpThumbLoadOnce8(v)
	case OpThumbLoadOnce8idx:
		return rewriteValueThumb_OpThumbLoadOnce8idx(v)
	case OpThumbMOVBUload:
		return rewriteValueThumb_OpThumbMOVBUload(v)
	case OpThumbMOVBUloadidx:
		return rewriteValueThumb_OpThumbMOVBUloadidx(v)
	case OpThumbMOVBUloadshiftLL:
		return rewriteValueThumb_OpThumbMOVBUloadshiftLL(v)
	case OpThumbMOVBUreg:
		return rewriteValueThumb_OpThumbMOVBUreg(v)
	case OpThumbMOVBload:
		return rewriteValueThumb_OpThumbMOVBload(v)
	case OpThumbMOVBloadidx:
		return rewriteValueThumb_OpThumbMOVBloadidx(v)
	case OpThumbMOVBloadshiftLL:
		return rewriteValueThumb_OpThumbMOVBloadshiftLL(v)
	case OpThumbMOVBreg:
		return rewriteValueThumb_OpThumbMOVBreg(v)
	case OpThumbMOVBstore:
		return rewriteValueThumb_OpThumbMOVBstore(v)
	case OpThumbMOVBstoreidx:
		return rewriteValueThumb_OpThumbMOVBstoreidx(v)
	case OpThumbMOVBstoreshiftLL:
		return rewriteValueThumb_OpThumbMOVBstoreshiftLL(v)
	case OpThumbMOVDload:
		return rewriteValueThumb_OpThumbMOVDload(v)
	case OpThumbMOVDstore:
		return rewriteValueThumb_OpThumbMOVDstore(v)
	case OpThumbMOVFload:
		return rewriteValueThumb_OpThumbMOVFload(v)
	case OpThumbMOVFstore:
		return rewriteValueThumb_OpThumbMOVFstore(v)
	case OpThumbMOVHUload:
		return rewriteValueThumb_OpThumbMOVHUload(v)
	case OpThumbMOVHUloadidx:
		return rewriteValueThumb_OpThumbMOVHUloadidx(v)
	case OpThumbMOVHUloadshiftLL:
		return rewriteValueThumb_OpThumbMOVHUloadshiftLL(v)
	case OpThumbMOVHUreg:
		return rewriteValueThumb_OpThumbMOVHUreg(v)
	case OpThumbMOVHload:
		return rewriteValueThumb_OpThumbMOVHload(v)
	case OpThumbMOVHloadidx:
		return rewriteValueThumb_OpThumbMOVHloadidx(v)
	case OpThumbMOVHloadshiftLL:
		return rewriteValueThumb_OpThumbMOVHloadshiftLL(v)
	case OpThumbMOVHreg:
		return rewriteValueThumb_OpThumbMOVHreg(v)
	case OpThumbMOVHstore:
		return rewriteValueThumb_OpThumbMOVHstore(v)
	case OpThumbMOVHstoreidx:
		return rewriteValueThumb_OpThumbMOVHstoreidx(v)
	case OpThumbMOVHstoreshiftLL:
		return rewriteValueThumb_OpThumbMOVHstoreshiftLL(v)
	case OpThumbMOVWload:
		return rewriteValueThumb_OpThumbMOVWload(v)
	case OpThumbMOVWloadidx:
		return rewriteValueThumb_OpThumbMOVWloadidx(v)
	case OpThumbMOVWloadshiftLL:
		return rewriteValueThumb_OpThumbMOVWloadshiftLL(v)
	case OpThumbMOVWreg:
		return rewriteValueThumb_OpThumbMOVWreg(v)
	case OpThumbMOVWstore:
		return rewriteValueThumb_OpThumbMOVWstore(v)
	case OpThumbMOVWstoreidx:
		return rewriteValueThumb_OpThumbMOVWstoreidx(v)
	case OpThumbMOVWstoreshiftLL:
		return rewriteValueThumb_OpThumbMOVWstoreshiftLL(v)
	case OpThumbMUL:
		return rewriteValueThumb_OpThumbMUL(v)
	case OpThumbMULA:
		return rewriteValueThumb_OpThumbMULA(v)
	case OpThumbMULD:
		return rewriteValueThumb_OpThumbMULD(v)
	case OpThumbMULF:
		return rewriteValueThumb_OpThumbMULF(v)
	case OpThumbMULS:
		return rewriteValueThumb_OpThumbMULS(v)
	case OpThumbMVN:
		return rewriteValueThumb_OpThumbMVN(v)
	case OpThumbMVNshiftLL:
		return rewriteValueThumb_OpThumbMVNshiftLL(v)
	case OpThumbMVNshiftRA:
		return rewriteValueThumb_OpThumbMVNshiftRA(v)
	case OpThumbMVNshiftRL:
		return rewriteValueThumb_OpThumbMVNshiftRL(v)
	case OpThumbNEGD:
		return rewriteValueThumb_OpThumbNEGD(v)
	case OpThumbNEGF:
		return rewriteValueThumb_OpThumbNEGF(v)
	case OpThumbNMULD:
		return rewriteValueThumb_OpThumbNMULD(v)
	case OpThumbNMULF:
		return rewriteValueThumb_OpThumbNMULF(v)
	case OpThumbNotEqual:
		return rewriteValueThumb_OpThumbNotEqual(v)
	case OpThumbOR:
		return rewriteValueThumb_OpThumbOR(v)
	case OpThumbORN:
		return rewriteValueThumb_OpThumbORN(v)
	case OpThumbORNconst:
		return rewriteValueThumb_OpThumbORNconst(v)
	case OpThumbORNshiftLL:
		return rewriteValueThumb_OpThumbORNshiftLL(v)
	case OpThumbORNshiftRA:
		return rewriteValueThumb_OpThumbORNshiftRA(v)
	case OpThumbORNshiftRL:
		return rewriteValueThumb_OpThumbORNshiftRL(v)
	case OpThumbORconst:
		return rewriteValueThumb_OpThumbORconst(v)
	case OpThumbORshiftLL:
		return rewriteValueThumb_OpThumbORshiftLL(v)
	case OpThumbORshiftRA:
		return rewriteValueThumb_OpThumbORshiftRA(v)
	case OpThumbORshiftRL:
		return rewriteValueThumb_OpThumbORshiftRL(v)
	case OpThumbRSB:
		return rewriteValueThumb_OpThumbRSB(v)
	case OpThumbRSBSshiftLL:
		return rewriteValueThumb_OpThumbRSBSshiftLL(v)
	case OpThumbRSBSshiftRA:
		return rewriteValueThumb_OpThumbRSBSshiftRA(v)
	case OpThumbRSBSshiftRL:
		return rewriteValueThumb_OpThumbRSBSshiftRL(v)
	case OpThumbRSBconst:
		return rewriteValueThumb_OpThumbRSBconst(v)
	case OpThumbRSBshiftLL:
		return rewriteValueThumb_OpThumbRSBshiftLL(v)
	case OpThumbRSBshiftRA:
		return rewriteValueThumb_OpThumbRSBshiftRA(v)
	case OpThumbRSBshiftRL:
		return rewriteValueThumb_OpThumbRSBshiftRL(v)
	case OpThumbSBC:
		return rewriteValueThumb_OpThumbSBC(v)
	case OpThumbSBCconst:
		return rewriteValueThumb_OpThumbSBCconst(v)
	case OpThumbSBCshiftLL:
		return rewriteValueThumb_OpThumbSBCshiftLL(v)
	case OpThumbSBCshiftRA:
		return rewriteValueThumb_OpThumbSBCshiftRA(v)
	case OpThumbSBCshiftRL:
		return rewriteValueThumb_OpThumbSBCshiftRL(v)
	case OpThumbSLL:
		return rewriteValueThumb_OpThumbSLL(v)
	case OpThumbSLLconst:
		return rewriteValueThumb_OpThumbSLLconst(v)
	case OpThumbSRA:
		return rewriteValueThumb_OpThumbSRA(v)
	case OpThumbSRAcond:
		return rewriteValueThumb_OpThumbSRAcond(v)
	case OpThumbSRAconst:
		return rewriteValueThumb_OpThumbSRAconst(v)
	case OpThumbSRL:
		return rewriteValueThumb_OpThumbSRL(v)
	case OpThumbSRLconst:
		return rewriteValueThumb_OpThumbSRLconst(v)
	case OpThumbSUB:
		return rewriteValueThumb_OpThumbSUB(v)
	case OpThumbSUBD:
		return rewriteValueThumb_OpThumbSUBD(v)
	case OpThumbSUBF:
		return rewriteValueThumb_OpThumbSUBF(v)
	case OpThumbSUBS:
		return rewriteValueThumb_OpThumbSUBS(v)
	case OpThumbSUBSshiftLL:
		return rewriteValueThumb_OpThumbSUBSshiftLL(v)
	case OpThumbSUBSshiftRA:
		return rewriteValueThumb_OpThumbSUBSshiftRA(v)
	case OpThumbSUBSshiftRL:
		return rewriteValueThumb_OpThumbSUBSshiftRL(v)
	case OpThumbSUBconst:
		return rewriteValueThumb_OpThumbSUBconst(v)
	case OpThumbSUBshiftLL:
		return rewriteValueThumb_OpThumbSUBshiftLL(v)
	case OpThumbSUBshiftRA:
		return rewriteValueThumb_OpThumbSUBshiftRA(v)
	case OpThumbSUBshiftRL:
		return rewriteValueThumb_OpThumbSUBshiftRL(v)
	case OpThumbStoreOnce16:
		return rewriteValueThumb_OpThumbStoreOnce16(v)
	case OpThumbStoreOnce16idx:
		return rewriteValueThumb_OpThumbStoreOnce16idx(v)
	case OpThumbStoreOnce32:
		return rewriteValueThumb_OpThumbStoreOnce32(v)
	case OpThumbStoreOnce32idx:
		return rewriteValueThumb_OpThumbStoreOnce32idx(v)
	case OpThumbStoreOnce8:
		return rewriteValueThumb_OpThumbStoreOnce8(v)
	case OpThumbStoreOnce8idx:
		return rewriteValueThumb_OpThumbStoreOnce8idx(v)
	case OpThumbTEQ:
		return rewriteValueThumb_OpThumbTEQ(v)
	case OpThumbTEQconst:
		return rewriteValueThumb_OpThumbTEQconst(v)
	case OpThumbTEQshiftLL:
		return rewriteValueThumb_OpThumbTEQshiftLL(v)
	case OpThumbTEQshiftRA:
		return rewriteValueThumb_OpThumbTEQshiftRA(v)
	case OpThumbTEQshiftRL:
		return rewriteValueThumb_OpThumbTEQshiftRL(v)
	case OpThumbTST:
		return rewriteValueThumb_OpThumbTST(v)
	case OpThumbTSTconst:
		return rewriteValueThumb_OpThumbTSTconst(v)
	case OpThumbTSTshiftLL:
		return rewriteValueThumb_OpThumbTSTshiftLL(v)
	case OpThumbTSTshiftRA:
		return rewriteValueThumb_OpThumbTSTshiftRA(v)
	case OpThumbTSTshiftRL:
		return rewriteValueThumb_OpThumbTSTshiftRL(v)
	case OpThumbXOR:
		return rewriteValueThumb_OpThumbXOR(v)
	case OpThumbXORconst:
		return rewriteValueThumb_OpThumbXORconst(v)
	case OpThumbXORshiftLL:
		return rewriteValueThumb_OpThumbXORshiftLL(v)
	case OpThumbXORshiftRA:
		return rewriteValueThumb_OpThumbXORshiftRA(v)
	case OpThumbXORshiftRL:
		return rewriteValueThumb_OpThumbXORshiftRL(v)
	case OpThumbXORshiftRR:
		return rewriteValueThumb_OpThumbXORshiftRR(v)
	case OpTrunc16to8:
		v.Op = OpCopy
		return true
	case OpTrunc32to16:
		v.Op = OpCopy
		return true
	case OpTrunc32to8:
		v.Op = OpCopy
		return true
	case OpWB:
		v.Op = OpThumbLoweredWB
		return true
	case OpXor16:
		v.Op = OpThumbXOR
		return true
	case OpXor32:
		v.Op = OpThumbXOR
		return true
	case OpXor8:
		v.Op = OpThumbXOR
		return true
	case OpZero:
		return rewriteValueThumb_OpZero(v)
	case OpZeroExt16to32:
		v.Op = OpThumbMOVHUreg
		return true
	case OpZeroExt8to16:
		v.Op = OpThumbMOVBUreg
		return true
	case OpZeroExt8to32:
		v.Op = OpThumbMOVBUreg
		return true
	case OpZeromask:
		return rewriteValueThumb_OpZeromask(v)
	}
	return false
}
func rewriteValueThumb_OpAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Addr {sym} base)
	// result: (MOVWaddr {sym} base)
	for {
		sym := auxToSym(v.Aux)
		base := v_0
		v.reset(OpThumbMOVWaddr)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
}
func rewriteValueThumb_OpAvg32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg32u <t> x y)
	// result: (ADD (SRLconst <t> (SUB <t> x y) [1]) y)
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, t)
		v0.AuxInt = int32ToAuxInt(1)
		v1 := b.NewValue0(v.Pos, OpThumbSUB, t)
		v1.AddArg2(x, y)
		v0.AddArg(v1)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueThumb_OpBitLen32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (BitLen32 <t> x)
	// result: (RSBconst [32] (CLZ <t> x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(32)
		v0 := b.NewValue0(v.Pos, OpThumbCLZ, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpConst16(v *Value) bool {
	// match: (Const16 [val])
	// result: (MOVWconst [int32(val)])
	for {
		val := auxIntToInt16(v.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(val))
		return true
	}
}
func rewriteValueThumb_OpConst32(v *Value) bool {
	// match: (Const32 [val])
	// result: (MOVWconst [int32(val)])
	for {
		val := auxIntToInt32(v.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(val))
		return true
	}
}
func rewriteValueThumb_OpConst32F(v *Value) bool {
	// match: (Const32F [val])
	// result: (MOVFconst [float64(val)])
	for {
		val := auxIntToFloat32(v.AuxInt)
		v.reset(OpThumbMOVFconst)
		v.AuxInt = float64ToAuxInt(float64(val))
		return true
	}
}
func rewriteValueThumb_OpConst64F(v *Value) bool {
	// match: (Const64F [val])
	// result: (MOVDconst [float64(val)])
	for {
		val := auxIntToFloat64(v.AuxInt)
		v.reset(OpThumbMOVDconst)
		v.AuxInt = float64ToAuxInt(float64(val))
		return true
	}
}
func rewriteValueThumb_OpConst8(v *Value) bool {
	// match: (Const8 [val])
	// result: (MOVWconst [int32(val)])
	for {
		val := auxIntToInt8(v.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(val))
		return true
	}
}
func rewriteValueThumb_OpConstBool(v *Value) bool {
	// match: (ConstBool [b])
	// result: (MOVWconst [b2i32(b)])
	for {
		b := auxIntToBool(v.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(b))
		return true
	}
}
func rewriteValueThumb_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVWconst [0])
	for {
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
}
func rewriteValueThumb_OpCtz16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz16 <t> x)
	// result: (CLZ <t> (RBIT <typ.UInt32> (ORconst <typ.UInt32> [0x10000] x)))
	for {
		t := v.Type
		x := v_0
		v.reset(OpThumbCLZ)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpThumbRBIT, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpThumbORconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(0x10000)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpCtz32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Ctz32 <t> x)
	// result: (CLZ <t> (RBIT <t> x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpThumbCLZ)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpThumbRBIT, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpCtz8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz8 <t> x)
	// result: (CLZ <t> (RBIT <typ.UInt32> (ORconst <typ.UInt32> [0x100] x)))
	for {
		t := v.Type
		x := v_0
		v.reset(OpThumbCLZ)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpThumbRBIT, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpThumbORconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(0x100)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 [c] x y)
	// result: (DIV (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbDIV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (DIVU (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbDIVU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div32 [c] x y)
	// result: (DIV x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbDIV)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueThumb_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (DIV (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbDIV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (DIVU (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbDIVU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (Equal (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32 x y)
	// result: (Equal (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32F x y)
	// result: (Equal (CMPF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPF, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64F x y)
	// result: (Equal (CMPD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (Equal (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (XORconst [1] (XOR <typ.Bool> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpThumbXOR, typ.Bool)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (EqPtr x y)
	// result: (Equal (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpFMA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMA x y z)
	// result: (FMULAD z x y)
	for {
		x := v_0
		y := v_1
		z := v_2
		v.reset(OpThumbFMULAD)
		v.AddArg3(z, x, y)
		return true
	}
}
func rewriteValueThumb_OpIsInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsInBounds idx len)
	// result: (LessThanU (CMP idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpThumbLessThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(idx, len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsNonNil ptr)
	// result: (NotEqual (CMPconst [0] ptr))
	for {
		ptr := v_0
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = int32ToAuxInt(0)
		v0.AddArg(ptr)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpIsSliceInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsSliceInBounds idx len)
	// result: (LessEqualU (CMP idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpThumbLessEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(idx, len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (LessEqual (CMP (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (LessEqualU (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32 x y)
	// result: (LessEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32F x y)
	// result: (GreaterEqual (CMPF y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPF, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32U x y)
	// result: (LessEqualU (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64F x y)
	// result: (GreaterEqual (CMPD y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPD, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (LessEqual (CMP (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (LessEqualU (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (LessThan (CMP (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (LessThanU (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32 x y)
	// result: (LessThan (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32F x y)
	// result: (GreaterThan (CMPF y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbGreaterThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMPF, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32U x y)
	// result: (LessThanU (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64F x y)
	// result: (GreaterThan (CMPD y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbGreaterThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMPD, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (LessThan (CMP (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (LessThanU (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbLessThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpThumbMOVBUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && isSigned(t))
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpThumbMOVBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && !isSigned(t))
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpThumbMOVBUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && isSigned(t))
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpThumbMOVHload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && !isSigned(t))
	// result: (MOVHUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpThumbMOVHUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) || isPtr(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpThumbMOVWload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (MOVFload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpThumbMOVFload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpThumbMOVDload)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpLocalAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LocalAddr {sym} base _)
	// result: (MOVWaddr {sym} base)
	for {
		sym := auxToSym(v.Aux)
		base := v_0
		v.reset(OpThumbMOVWaddr)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
}
func rewriteValueThumb_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = int32ToAuxInt(256)
		v2.AddArg(v1)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueThumb_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh16x32 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = int32ToAuxInt(256)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SLLconst x [int32(c)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = int32ToAuxInt(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (Const16 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = int16ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 x y)
	// result: (SLL x (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSLL)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueThumb_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = int32ToAuxInt(256)
		v2.AddArg(v1)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueThumb_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh32x32 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = int32ToAuxInt(256)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SLLconst x [int32(c)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = int32ToAuxInt(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (Const32 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 x y)
	// result: (SLL x (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSLL)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueThumb_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = int32ToAuxInt(256)
		v2.AddArg(v1)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueThumb_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh8x32 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = int32ToAuxInt(256)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SLLconst x [int32(c)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = int32ToAuxInt(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (Const8 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = int8ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 x y)
	// result: (SLL x (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSLL)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueThumb_OpMMIOLoad16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MMIOLoad16 ptr mem)
	// result: (LoadOnce16 [0] ptr mem)
	for {
		ptr := v_0
		mem := v_1
		v.reset(OpThumbLoadOnce16)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg2(ptr, mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOLoad32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MMIOLoad32 ptr mem)
	// result: (LoadOnce32 [0] ptr mem)
	for {
		ptr := v_0
		mem := v_1
		v.reset(OpThumbLoadOnce32)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg2(ptr, mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOLoad8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MMIOLoad8 ptr mem)
	// result: (LoadOnce8 [0] ptr mem)
	for {
		ptr := v_0
		mem := v_1
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg2(ptr, mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOStore16(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MMIOStore16 ptr val mem)
	// result: (StoreOnce16 [0] ptr val mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce16)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg3(ptr, val, mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOStore32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MMIOStore32 ptr val mem)
	// result: (StoreOnce32 [0] ptr val mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce32)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg3(ptr, val, mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOStore8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MMIOStore8 ptr val mem)
	// result: (StoreOnce8 [0] ptr val mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce8)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg3(ptr, val, mem)
		return true
	}
}
func rewriteValueThumb_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (Mod32 (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (Mod32u (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Mod32 x y)
	// result: (SUB x (MUL <y.Type> y (DIV <x.Type> x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSUB)
		v0 := b.NewValue0(v.Pos, OpThumbMUL, y.Type)
		v1 := b.NewValue0(v.Pos, OpThumbDIV, x.Type)
		v1.AddArg2(x, y)
		v0.AddArg2(y, v1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueThumb_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Mod32u x y)
	// result: (SUB x (MUL <y.Type> y (DIVU <x.Type> x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSUB)
		v0 := b.NewValue0(v.Pos, OpThumbMUL, y.Type)
		v1 := b.NewValue0(v.Pos, OpThumbDIVU, x.Type)
		v1.AddArg2(x, y)
		v0.AddArg2(y, v1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueThumb_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (Mod32 (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (Mod32u (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.copyOf(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBUload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpThumbMOVBstore)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHUload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v0 := b.NewValue0(v.Pos, OpThumbMOVHUload, typ.UInt16)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v0.AuxInt = int32ToAuxInt(1)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpThumbMOVWstore)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWload, typ.UInt32)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore [2] dst (MOVHUload [2] src mem) (MOVHstore dst (MOVHUload src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpThumbMOVHUload, typ.UInt16)
		v0.AuxInt = int32ToAuxInt(2)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpThumbMOVHstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpThumbMOVHUload, typ.UInt16)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (MOVBstore [3] dst (MOVBUload [3] src mem) (MOVBstore [2] dst (MOVBUload [2] src mem) (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(3)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v0.AuxInt = int32ToAuxInt(3)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(2)
		v2 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v2.AuxInt = int32ToAuxInt(2)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(1)
		v4 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v4.AuxInt = int32ToAuxInt(1)
		v4.AddArg2(src, mem)
		v5 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v6 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v6.AddArg2(src, mem)
		v5.AddArg3(dst, v6, mem)
		v3.AddArg3(dst, v4, v5)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [3] dst src mem)
	// result: (MOVBstore [2] dst (MOVBUload [2] src mem) (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 3 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v0.AuxInt = int32ToAuxInt(2)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(1)
		v2 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v2.AuxInt = int32ToAuxInt(1)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: s%4 == 0 && s > 4 && s <= 512 && t.Alignment()%4 == 0 && !config.noDuffDevice && logLargeCopy(v, s)
	// result: (DUFFCOPY [8 * (128 - s/4)] dst src mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(s%4 == 0 && s > 4 && s <= 512 && t.Alignment()%4 == 0 && !config.noDuffDevice && logLargeCopy(v, s)) {
			break
		}
		v.reset(OpThumbDUFFCOPY)
		v.AuxInt = int64ToAuxInt(8 * (128 - s/4))
		v.AddArg3(dst, src, mem)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: ((s > 512 || config.noDuffDevice) || t.Alignment()%4 != 0) && logLargeCopy(v, s)
	// result: (LoweredMove [t.Alignment()] dst src (ADDconst <src.Type> src [int32(s-moveSize(t.Alignment(), config))]) mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(((s > 512 || config.noDuffDevice) || t.Alignment()%4 != 0) && logLargeCopy(v, s)) {
			break
		}
		v.reset(OpThumbLoweredMove)
		v.AuxInt = int64ToAuxInt(t.Alignment())
		v0 := b.NewValue0(v.Pos, OpThumbADDconst, src.Type)
		v0.AuxInt = int32ToAuxInt(int32(s - moveSize(t.Alignment(), config)))
		v0.AddArg(src)
		v.AddArg4(dst, src, v0, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpNeg16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg16 x)
	// result: (RSBconst [0] x)
	for {
		x := v_0
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpNeg32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg32 x)
	// result: (RSBconst [0] x)
	for {
		x := v_0
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpNeg8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg8 x)
	// result: (RSBconst [0] x)
	for {
		x := v_0
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (NotEqual (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32 x y)
	// result: (NotEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32F x y)
	// result: (NotEqual (CMPF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPF, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64F x y)
	// result: (NotEqual (CMPD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (NotEqual (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (NeqPtr x y)
	// result: (NotEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Not x)
	// result: (XORconst [1] x)
	for {
		x := v_0
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (OffPtr [off] ptr:(SP))
	// result: (MOVWaddr [int32(off)] ptr)
	for {
		off := auxIntToInt64(v.AuxInt)
		ptr := v_0
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpThumbMOVWaddr)
		v.AuxInt = int32ToAuxInt(int32(off))
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// result: (ADDconst [int32(off)] ptr)
	for {
		off := auxIntToInt64(v.AuxInt)
		ptr := v_0
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(int32(off))
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueThumb_OpPanicBounds(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicBoundsA [kind] x y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpThumbLoweredPanicBoundsA)
		v.AuxInt = int64ToAuxInt(kind)
		v.AddArg3(x, y, mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicBoundsB [kind] x y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpThumbLoweredPanicBoundsB)
		v.AuxInt = int64ToAuxInt(kind)
		v.AddArg3(x, y, mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicBoundsC [kind] x y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpThumbLoweredPanicBoundsC)
		v.AuxInt = int64ToAuxInt(kind)
		v.AddArg3(x, y, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpPanicExtend(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicExtendA [kind] hi lo y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		hi := v_0
		lo := v_1
		y := v_2
		mem := v_3
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpThumbLoweredPanicExtendA)
		v.AuxInt = int64ToAuxInt(kind)
		v.AddArg4(hi, lo, y, mem)
		return true
	}
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicExtendB [kind] hi lo y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		hi := v_0
		lo := v_1
		y := v_2
		mem := v_3
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpThumbLoweredPanicExtendB)
		v.AuxInt = int64ToAuxInt(kind)
		v.AddArg4(hi, lo, y, mem)
		return true
	}
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicExtendC [kind] hi lo y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		hi := v_0
		lo := v_1
		y := v_2
		mem := v_3
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpThumbLoweredPanicExtendC)
		v.AuxInt = int64ToAuxInt(kind)
		v.AddArg4(hi, lo, y, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVWconst [c]))
	// result: (Or16 (Lsh16x32 <t> x (MOVWconst [c&15])) (Rsh16Ux32 <t> x (MOVWconst [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x32, t)
		v1 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(c & 15)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux32, t)
		v3 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(-c & 15)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueThumb_OpRotateLeft32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RotateLeft32 x (MOVWconst [c]))
	// result: (SRRconst [-c&31] x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSRRconst)
		v.AuxInt = int32ToAuxInt(-c & 31)
		v.AddArg(x)
		return true
	}
	// match: (RotateLeft32 x y)
	// result: (SRR x (RSBconst [0] <y.Type> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRR)
		v0 := b.NewValue0(v.Pos, OpThumbRSBconst, y.Type)
		v0.AuxInt = int32ToAuxInt(0)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueThumb_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVWconst [c]))
	// result: (Or8 (Lsh8x32 <t> x (MOVWconst [c&7])) (Rsh8Ux32 <t> x (MOVWconst [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x32, t)
		v1 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(c & 7)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux32, t)
		v3 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(-c & 7)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 x y)
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt16to32 x) (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v3.AuxInt = int32ToAuxInt(256)
		v3.AddArg(v2)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueThumb_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 x y)
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt16to32 x) y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = int32ToAuxInt(256)
		v2.AddArg(y)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueThumb_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRLconst (SLLconst <typ.UInt32> x [16]) [int32(c+16)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(int32(c + 16))
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (Const16 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = int16ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 x y)
	// result: (SRL (ZeroExt16to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRL)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 x y)
	// result: (SRAcond (SignExt16to32 x) (ZeroExt16to32 y) (CMPconst [256] (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = int32ToAuxInt(256)
		v2.AddArg(v1)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueThumb_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 x y)
	// result: (SRAcond (SignExt16to32 x) y (CMPconst [256] y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = int32ToAuxInt(256)
		v1.AddArg(y)
		v.AddArg3(v0, y, v1)
		return true
	}
}
func rewriteValueThumb_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRAconst (SLLconst <typ.UInt32> x [16]) [int32(c+16)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(int32(c + 16))
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (SRAconst (SLLconst <typ.UInt32> x [16]) [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 x y)
	// result: (SRA (SignExt16to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 x y)
	// result: (CMOVWHSconst (SRL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = int32ToAuxInt(256)
		v2.AddArg(v1)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueThumb_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32Ux32 x y)
	// result: (CMOVWHSconst (SRL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = int32ToAuxInt(256)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Rsh32Ux64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRLconst x [int32(c)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (Const32 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 x y)
	// result: (SRL x (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRL)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueThumb_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 x y)
	// result: (SRAcond x (ZeroExt16to32 y) (CMPconst [256] (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRAcond)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = int32ToAuxInt(256)
		v1.AddArg(v0)
		v.AddArg3(x, v0, v1)
		return true
	}
}
func rewriteValueThumb_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32x32 x y)
	// result: (SRAcond x y (CMPconst [256] y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRAcond)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = int32ToAuxInt(256)
		v0.AddArg(y)
		v.AddArg3(x, y, v0)
		return true
	}
}
func rewriteValueThumb_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRAconst x [int32(c)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (SRAconst x [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 x y)
	// result: (SRA x (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRA)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueThumb_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 x y)
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt8to32 x) (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v3.AuxInt = int32ToAuxInt(256)
		v3.AddArg(v2)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueThumb_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 x y)
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt8to32 x) y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = int32ToAuxInt(256)
		v2.AddArg(y)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueThumb_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRLconst (SLLconst <typ.UInt32> x [24]) [int32(c+24)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(int32(c + 24))
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(24)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (Const8 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = int8ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 x y)
	// result: (SRL (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRL)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 x y)
	// result: (SRAcond (SignExt8to32 x) (ZeroExt16to32 y) (CMPconst [256] (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = int32ToAuxInt(256)
		v2.AddArg(v1)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueThumb_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 x y)
	// result: (SRAcond (SignExt8to32 x) y (CMPconst [256] y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = int32ToAuxInt(256)
		v1.AddArg(y)
		v.AddArg3(v0, y, v1)
		return true
	}
}
func rewriteValueThumb_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRAconst (SLLconst <typ.UInt32> x [24]) [int32(c+24)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(int32(c + 24))
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(24)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (SRAconst (SLLconst <typ.UInt32> x [24]) [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(24)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 x y)
	// result: (SRA (SignExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpThumbSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueThumb_OpSignmask(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Signmask x)
	// result: (SRAconst x [31])
	for {
		x := v_0
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SRAconst (RSBconst <t> [0] x) [31])
	for {
		t := v.Type
		x := v_0
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v0 := b.NewValue0(v.Pos, OpThumbRSBconst, t)
		v0.AuxInt = int32ToAuxInt(0)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 1) {
			break
		}
		v.reset(OpThumbMOVBstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 2) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 4 && !is32BitFloat(val.Type)
	// result: (MOVWstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 4 && !is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpThumbMOVWstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 4 && is32BitFloat(val.Type)
	// result: (MOVFstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpThumbMOVFstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 8 && is64BitFloat(val.Type)
	// result: (MOVDstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpThumbMOVDstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADC(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADC (MOVWconst [c]) x flags)
	// result: (ADCconst [c] x flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_0.AuxInt)
			x := v_1
			flags := v_2
			v.reset(OpThumbADCconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, flags)
			return true
		}
		break
	}
	// match: (ADC x (SLLconst [c] y) flags)
	// result: (ADCshiftLL x y [c] flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSLLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			flags := v_2
			v.reset(OpThumbADCshiftLL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg3(x, y, flags)
			return true
		}
		break
	}
	// match: (ADC x (SRLconst [c] y) flags)
	// result: (ADCshiftRL x y [c] flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			flags := v_2
			v.reset(OpThumbADCshiftRL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg3(x, y, flags)
			return true
		}
		break
	}
	// match: (ADC x (SRAconst [c] y) flags)
	// result: (ADCshiftRA x y [c] flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRAconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			flags := v_2
			v.reset(OpThumbADCshiftRA)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg3(x, y, flags)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbADCconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADCconst [c] (ADDconst [d] x) flags)
	// result: (ADCconst [c+d] x flags)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		flags := v_1
		v.reset(OpThumbADCconst)
		v.AuxInt = int32ToAuxInt(c + d)
		v.AddArg2(x, flags)
		return true
	}
	// match: (ADCconst [c] (SUBconst [d] x) flags)
	// result: (ADCconst [c-d] x flags)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		flags := v_1
		v.reset(OpThumbADCconst)
		v.AuxInt = int32ToAuxInt(c - d)
		v.AddArg2(x, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADCshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADCshiftLL (MOVWconst [c]) x [d] flags)
	// result: (ADCconst [c] (SLLconst <x.Type> x [d]) flags)
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		flags := v_2
		v.reset(OpThumbADCconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg2(v0, flags)
		return true
	}
	// match: (ADCshiftLL x (MOVWconst [c]) [d] flags)
	// result: (ADCconst x [c<<uint64(d)] flags)
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		flags := v_2
		v.reset(OpThumbADCconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg2(x, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADCshiftRA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADCshiftRA (MOVWconst [c]) x [d] flags)
	// result: (ADCconst [c] (SRAconst <x.Type> x [d]) flags)
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		flags := v_2
		v.reset(OpThumbADCconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg2(v0, flags)
		return true
	}
	// match: (ADCshiftRA x (MOVWconst [c]) [d] flags)
	// result: (ADCconst x [c>>uint64(d)] flags)
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		flags := v_2
		v.reset(OpThumbADCconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg2(x, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADCshiftRL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADCshiftRL (MOVWconst [c]) x [d] flags)
	// result: (ADCconst [c] (SRLconst <x.Type> x [d]) flags)
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		flags := v_2
		v.reset(OpThumbADCconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg2(v0, flags)
		return true
	}
	// match: (ADCshiftRL x (MOVWconst [c]) [d] flags)
	// result: (ADCconst x [int32(uint32(c)>>uint64(d))] flags)
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		flags := v_2
		v.reset(OpThumbADCconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg2(x, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADD x (MOVWconst [c]))
	// result: (ADDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpThumbADDconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADD x (SLLconst [c] y))
	// result: (ADDshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSLLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbADDshiftLL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (ADD x (SRLconst [c] y))
	// result: (ADDshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbADDshiftRL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (ADD x (SRAconst [c] y))
	// result: (ADDshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRAconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbADDshiftRA)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (ADD x (RSBconst [0] y))
	// result: (SUB x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbRSBconst || auxIntToInt32(v_1.AuxInt) != 0 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpThumbSUB)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (ADD <t> (RSBconst [c] x) (RSBconst [d] y))
	// result: (RSBconst [c+d] (ADD <t> x y))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpThumbRSBconst {
				continue
			}
			c := auxIntToInt32(v_0.AuxInt)
			x := v_0.Args[0]
			if v_1.Op != OpThumbRSBconst {
				continue
			}
			d := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbRSBconst)
			v.AuxInt = int32ToAuxInt(c + d)
			v0 := b.NewValue0(v.Pos, OpThumbADD, t)
			v0.AddArg2(x, y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (ADD (MUL x y) a)
	// result: (MULA x y a)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpThumbMUL {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			a := v_1
			v.reset(OpThumbMULA)
			v.AddArg3(x, y, a)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbADDD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDD a (MULD x y))
	// cond: a.Uses == 1
	// result: (MULAD a x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			a := v_0
			if v_1.Op != OpThumbMULD {
				continue
			}
			y := v_1.Args[1]
			x := v_1.Args[0]
			if !(a.Uses == 1) {
				continue
			}
			v.reset(OpThumbMULAD)
			v.AddArg3(a, x, y)
			return true
		}
		break
	}
	// match: (ADDD a (NMULD x y))
	// cond: a.Uses == 1
	// result: (MULSD a x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			a := v_0
			if v_1.Op != OpThumbNMULD {
				continue
			}
			y := v_1.Args[1]
			x := v_1.Args[0]
			if !(a.Uses == 1) {
				continue
			}
			v.reset(OpThumbMULSD)
			v.AddArg3(a, x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbADDF(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDF a (MULF x y))
	// cond: a.Uses == 1
	// result: (MULAF a x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			a := v_0
			if v_1.Op != OpThumbMULF {
				continue
			}
			y := v_1.Args[1]
			x := v_1.Args[0]
			if !(a.Uses == 1) {
				continue
			}
			v.reset(OpThumbMULAF)
			v.AddArg3(a, x, y)
			return true
		}
		break
	}
	// match: (ADDF a (NMULF x y))
	// cond: a.Uses == 1
	// result: (MULSF a x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			a := v_0
			if v_1.Op != OpThumbNMULF {
				continue
			}
			y := v_1.Args[1]
			x := v_1.Args[0]
			if !(a.Uses == 1) {
				continue
			}
			v.reset(OpThumbMULSF)
			v.AddArg3(a, x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbADDS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDS x (MOVWconst [c]))
	// result: (ADDSconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpThumbADDSconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDS x (SLLconst [c] y))
	// result: (ADDSshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSLLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbADDSshiftLL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (ADDS x (SRLconst [c] y))
	// result: (ADDSshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbADDSshiftRL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (ADDS x (SRAconst [c] y))
	// result: (ADDSshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRAconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbADDSshiftRA)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbADDSshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDSshiftLL (MOVWconst [c]) x [d])
	// result: (ADDSconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbADDSconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftLL x (MOVWconst [c]) [d])
	// result: (ADDSconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbADDSconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDSshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDSshiftRA (MOVWconst [c]) x [d])
	// result: (ADDSconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbADDSconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftRA x (MOVWconst [c]) [d])
	// result: (ADDSconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbADDSconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDSshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDSshiftRL (MOVWconst [c]) x [d])
	// result: (ADDSconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbADDSconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftRL x (MOVWconst [c]) [d])
	// result: (ADDSconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbADDSconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDconst [off1] (MOVWaddr [off2] {sym} ptr))
	// result: (MOVWaddr [off1+off2] {sym} ptr)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVWaddr)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg(ptr)
		return true
	}
	// match: (ADDconst [0] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (ADDconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c+d])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(c + d)
		return true
	}
	// match: (ADDconst [c] (ADDconst [d] x))
	// result: (ADDconst [c+d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(c + d)
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (SUBconst [d] x))
	// result: (ADDconst [c-d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(c - d)
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (RSBconst [d] x))
	// result: (RSBconst [c+d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbRSBconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(c + d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ADDshiftLL (MOVWconst [c]) x [d])
	// result: (ADDconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftLL x (MOVWconst [c]) [d])
	// result: (ADDconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftLL [c] (SRLconst x [32-c]) x)
	// result: (SRRconst [32-c] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSRLconst || auxIntToInt32(v_0.AuxInt) != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = int32ToAuxInt(32 - c)
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftLL <typ.UInt16> [8] (BFXU <typ.UInt16> [int32(armBFAuxInt(8, 8))] x) x)
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || auxIntToInt32(v.AuxInt) != 8 || v_0.Op != OpThumbBFXU || v_0.Type != typ.UInt16 || auxIntToInt32(v_0.AuxInt) != int32(armBFAuxInt(8, 8)) {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbREV16)
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftLL <typ.UInt16> [8] (SRLconst <typ.UInt16> [24] (SLLconst [16] x)) x)
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || auxIntToInt32(v.AuxInt) != 8 || v_0.Op != OpThumbSRLconst || v_0.Type != typ.UInt16 || auxIntToInt32(v_0.AuxInt) != 24 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbSLLconst || auxIntToInt32(v_0_0.AuxInt) != 16 {
			break
		}
		x := v_0_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbREV16)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDshiftRA (MOVWconst [c]) x [d])
	// result: (ADDconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftRA x (MOVWconst [c]) [d])
	// result: (ADDconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDshiftRL (MOVWconst [c]) x [d])
	// result: (ADDconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftRL x (MOVWconst [c]) [d])
	// result: (ADDconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftRL [c] (SLLconst x [32-c]) x)
	// result: (SRRconst [ c] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSLLconst || auxIntToInt32(v_0.AuxInt) != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbAND(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AND x (MOVWconst [c]))
	// result: (ANDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpThumbANDconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND x (SLLconst [c] y))
	// result: (ANDshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSLLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbANDshiftLL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (AND x (SRLconst [c] y))
	// result: (ANDshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbANDshiftRL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (AND x (SRAconst [c] y))
	// result: (ANDshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRAconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbANDshiftRA)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (AND x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (AND x (MVN y))
	// result: (BIC x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMVN {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpThumbBIC)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (AND x (MVNshiftLL y [c]))
	// result: (BICshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMVNshiftLL {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbBICshiftLL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (AND x (MVNshiftRL y [c]))
	// result: (BICshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMVNshiftRL {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbBICshiftRL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (AND x (MVNshiftRA y [c]))
	// result: (BICshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMVNshiftRA {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbBICshiftRA)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbANDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ANDconst [c] y:(MOVBUreg _))
	// cond: c&0xFF == 0xFF
	// result: y
	for {
		c := auxIntToInt32(v.AuxInt)
		y := v_0
		if y.Op != OpThumbMOVBUreg || !(c&0xFF == 0xFF) {
			break
		}
		v.copyOf(y)
		return true
	}
	// match: (ANDconst [c] y:(MOVHUreg _))
	// cond: c&0xFFFF == 0xFFFF
	// result: y
	for {
		c := auxIntToInt32(v.AuxInt)
		y := v_0
		if y.Op != OpThumbMOVHUreg || !(c&0xFFFF == 0xFFFF) {
			break
		}
		v.copyOf(y)
		return true
	}
	// match: (ANDconst [c] (MOVBUreg x))
	// result: (ANDconst [c&0xFF] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVBUreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c & 0xFF)
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVHUreg x))
	// result: (ANDconst [c&0xFFFF] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVHUreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c & 0xFFFF)
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [0] _)
	// result: (MOVWconst [0])
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (ANDconst [c] x)
	// cond: int32(c)==-1
	// result: x
	for {
		c := auxIntToInt32(v.AuxInt)
		x := v_0
		if !(int32(c) == -1) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ANDconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c&d])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(c & d)
		return true
	}
	// match: (ANDconst [c] (ANDconst [d] x))
	// result: (ANDconst [c&d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbANDconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c & d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbANDshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ANDshiftLL (MOVWconst [c]) x [d])
	// result: (ANDconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftLL x (MOVWconst [c]) [d])
	// result: (ANDconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (ANDshiftLL x y:(SLLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		y := v_1
		if y.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(y.AuxInt)
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.copyOf(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbANDshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ANDshiftRA (MOVWconst [c]) x [d])
	// result: (ANDconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftRA x (MOVWconst [c]) [d])
	// result: (ANDconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (ANDshiftRA x y:(SRAconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		y := v_1
		if y.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(y.AuxInt)
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.copyOf(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbANDshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ANDshiftRL (MOVWconst [c]) x [d])
	// result: (ANDconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftRL x (MOVWconst [c]) [d])
	// result: (ANDconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ANDshiftRL x y:(SRLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		y := v_1
		if y.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(y.AuxInt)
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.copyOf(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBFX(v *Value) bool {
	v_0 := v.Args[0]
	// match: (BFX [c] (MOVWconst [d]))
	// result: (MOVWconst [d<<(32-uint32(c&0xff)-uint32(c>>8))>>(32-uint32(c>>8))])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(d << (32 - uint32(c&0xff) - uint32(c>>8)) >> (32 - uint32(c>>8)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBFXU(v *Value) bool {
	v_0 := v.Args[0]
	// match: (BFXU [c] (MOVWconst [d]))
	// result: (MOVWconst [int32(uint32(d)<<(32-uint32(c&0xff)-uint32(c>>8))>>(32-uint32(c>>8)))])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(d) << (32 - uint32(c&0xff) - uint32(c>>8)) >> (32 - uint32(c>>8))))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBIC(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BIC x (MOVWconst [c]))
	// result: (BICconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbBICconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (BIC x (SLLconst [c] y))
	// result: (BICshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbBICshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (BIC x (SRLconst [c] y))
	// result: (BICshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbBICshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (BIC x (SRAconst [c] y))
	// result: (BICshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbBICshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (BIC x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBICconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (BICconst [0] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (BICconst [c] _)
	// cond: int32(c)==-1
	// result: (MOVWconst [0])
	for {
		c := auxIntToInt32(v.AuxInt)
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (BICconst [c] (MOVWconst [d]))
	// result: (MOVWconst [d&^c])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(d &^ c)
		return true
	}
	// match: (BICconst [c] (BICconst [d] x))
	// result: (BICconst [c|d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbBICconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbBICconst)
		v.AuxInt = int32ToAuxInt(c | d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBICshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BICshiftLL x (MOVWconst [c]) [d])
	// result: (BICconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbBICconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (BICshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBICshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BICshiftRA x (MOVWconst [c]) [d])
	// result: (BICconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbBICconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (BICshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBICshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BICshiftRL x (MOVWconst [c]) [d])
	// result: (BICconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbBICconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (BICshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMN(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMN x (MOVWconst [c]))
	// result: (CMNconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpThumbCMNconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (CMN x (SLLconst [c] y))
	// result: (CMNshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSLLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbCMNshiftLL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (CMN x (SRLconst [c] y))
	// result: (CMNshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbCMNshiftRL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (CMN x (SRAconst [c] y))
	// result: (CMNshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRAconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbCMNshiftRA)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbCMNconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMNconst (MOVWconst [x]) [y])
	// result: (FlagConstant [addFlags32(x,y)])
	for {
		y := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbFlagConstant)
		v.AuxInt = flagConstantToAuxInt(addFlags32(x, y))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMNshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMNshiftLL (MOVWconst [c]) x [d])
	// result: (CMNconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbCMNconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftLL x (MOVWconst [c]) [d])
	// result: (CMNconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbCMNconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMNshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMNshiftRA (MOVWconst [c]) x [d])
	// result: (CMNconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbCMNconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftRA x (MOVWconst [c]) [d])
	// result: (CMNconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbCMNconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMNshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMNshiftRL (MOVWconst [c]) x [d])
	// result: (CMNconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbCMNconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftRL x (MOVWconst [c]) [d])
	// result: (CMNconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbCMNconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMOVWHSconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMOVWHSconst _ (FlagConstant [fc]) [c])
	// cond: fc.uge()
	// result: (MOVWconst [c])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_1.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_1.AuxInt)
		if !(fc.uge()) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(c)
		return true
	}
	// match: (CMOVWHSconst x (FlagConstant [fc]) [c])
	// cond: fc.ult()
	// result: x
	for {
		x := v_0
		if v_1.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_1.AuxInt)
		if !(fc.ult()) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (CMOVWHSconst x (InvertFlags flags) [c])
	// result: (CMOVWLSconst x flags [c])
	for {
		c := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbInvertFlags {
			break
		}
		flags := v_1.Args[0]
		v.reset(OpThumbCMOVWLSconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMOVWLSconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMOVWLSconst _ (FlagConstant [fc]) [c])
	// cond: fc.ule()
	// result: (MOVWconst [c])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_1.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_1.AuxInt)
		if !(fc.ule()) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(c)
		return true
	}
	// match: (CMOVWLSconst x (FlagConstant [fc]) [c])
	// cond: fc.ugt()
	// result: x
	for {
		x := v_0
		if v_1.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_1.AuxInt)
		if !(fc.ugt()) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (CMOVWLSconst x (InvertFlags flags) [c])
	// result: (CMOVWHSconst x flags [c])
	for {
		c := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbInvertFlags {
			break
		}
		flags := v_1.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMP(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMP x (MOVWconst [c]))
	// result: (CMPconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbCMPconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (CMP (MOVWconst [c]) x)
	// result: (InvertFlags (CMPconst [c] x))
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = int32ToAuxInt(c)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x y)
	// cond: x.ID > y.ID
	// result: (InvertFlags (CMP y x))
	for {
		x := v_0
		y := v_1
		if !(x.ID > y.ID) {
			break
		}
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SLLconst [c] y))
	// result: (CMPshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbCMPshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (CMP (SLLconst [c] y) x)
	// result: (InvertFlags (CMPshiftLL x y [c]))
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPshiftLL, types.TypeFlags)
		v0.AuxInt = int32ToAuxInt(c)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SRLconst [c] y))
	// result: (CMPshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbCMPshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (CMP (SRLconst [c] y) x)
	// result: (InvertFlags (CMPshiftRL x y [c]))
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRL, types.TypeFlags)
		v0.AuxInt = int32ToAuxInt(c)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SRAconst [c] y))
	// result: (CMPshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbCMPshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (CMP (SRAconst [c] y) x)
	// result: (InvertFlags (CMPshiftRA x y [c]))
	for {
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRA, types.TypeFlags)
		v0.AuxInt = int32ToAuxInt(c)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMPD x (MOVDconst [0]))
	// result: (CMPD0 x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVDconst || auxIntToFloat64(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpThumbCMPD0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPF(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMPF x (MOVFconst [0]))
	// result: (CMPF0 x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVFconst || auxIntToFloat64(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpThumbCMPF0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMPconst (MOVWconst [x]) [y])
	// result: (FlagConstant [subFlags32(x,y)])
	for {
		y := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbFlagConstant)
		v.AuxInt = flagConstantToAuxInt(subFlags32(x, y))
		return true
	}
	// match: (CMPconst (MOVBUreg _) [c])
	// cond: 0xff < c
	// result: (FlagConstant [subFlags32(0, 1)])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVBUreg || !(0xff < c) {
			break
		}
		v.reset(OpThumbFlagConstant)
		v.AuxInt = flagConstantToAuxInt(subFlags32(0, 1))
		return true
	}
	// match: (CMPconst (MOVHUreg _) [c])
	// cond: 0xffff < c
	// result: (FlagConstant [subFlags32(0, 1)])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVHUreg || !(0xffff < c) {
			break
		}
		v.reset(OpThumbFlagConstant)
		v.AuxInt = flagConstantToAuxInt(subFlags32(0, 1))
		return true
	}
	// match: (CMPconst (ANDconst _ [m]) [n])
	// cond: 0 <= m && m < n
	// result: (FlagConstant [subFlags32(0, 1)])
	for {
		n := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbANDconst {
			break
		}
		m := auxIntToInt32(v_0.AuxInt)
		if !(0 <= m && m < n) {
			break
		}
		v.reset(OpThumbFlagConstant)
		v.AuxInt = flagConstantToAuxInt(subFlags32(0, 1))
		return true
	}
	// match: (CMPconst (SRLconst _ [c]) [n])
	// cond: 0 <= n && 0 < c && c <= 32 && (1<<uint32(32-c)) <= uint32(n)
	// result: (FlagConstant [subFlags32(0, 1)])
	for {
		n := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		if !(0 <= n && 0 < c && c <= 32 && (1<<uint32(32-c)) <= uint32(n)) {
			break
		}
		v.reset(OpThumbFlagConstant)
		v.AuxInt = flagConstantToAuxInt(subFlags32(0, 1))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPshiftLL (MOVWconst [c]) x [d])
	// result: (InvertFlags (CMPconst [c] (SLLconst <x.Type> x [d])))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = int32ToAuxInt(c)
		v1 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v1.AuxInt = int32ToAuxInt(d)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftLL x (MOVWconst [c]) [d])
	// result: (CMPconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbCMPconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPshiftRA (MOVWconst [c]) x [d])
	// result: (InvertFlags (CMPconst [c] (SRAconst <x.Type> x [d])))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = int32ToAuxInt(c)
		v1 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v1.AuxInt = int32ToAuxInt(d)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftRA x (MOVWconst [c]) [d])
	// result: (CMPconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbCMPconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPshiftRL (MOVWconst [c]) x [d])
	// result: (InvertFlags (CMPconst [c] (SRLconst <x.Type> x [d])))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = int32ToAuxInt(c)
		v1 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v1.AuxInt = int32ToAuxInt(d)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftRL x (MOVWconst [c]) [d])
	// result: (CMPconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbCMPconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbDIV(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (DIV x (MOVWconst [1]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst || auxIntToInt32(v_1.AuxInt) != 1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (DIV x (MOVWconst [c]))
	// cond: isPowerOfTwo32(c)
	// result: (SRAconst [int32(log32(c))] x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if !(isPowerOfTwo32(c)) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(int32(log32(c)))
		v.AddArg(x)
		return true
	}
	// match: (DIV (MOVWconst [c]) (MOVWconst [d]))
	// cond: d != 0
	// result: (MOVWconst [c/d])
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_1.AuxInt)
		if !(d != 0) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(c / d)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbDIVU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (DIVU x (MOVWconst [1]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst || auxIntToInt32(v_1.AuxInt) != 1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (DIVU x (MOVWconst [c]))
	// cond: isPowerOfTwo32(c)
	// result: (SRLconst [int32(log32(c))] x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if !(isPowerOfTwo32(c)) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(int32(log32(c)))
		v.AddArg(x)
		return true
	}
	// match: (DIVU (MOVWconst [c]) (MOVWconst [d]))
	// cond: d != 0
	// result: (MOVWconst [int32(uint32(c)/uint32(d))])
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_1.AuxInt)
		if !(d != 0) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) / uint32(d)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbEqual(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Equal (FlagConstant [fc]))
	// result: (MOVWconst [b2i32(fc.eq())])
	for {
		if v_0.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(fc.eq()))
		return true
	}
	// match: (Equal (InvertFlags x))
	// result: (Equal x)
	for {
		if v_0.Op != OpThumbInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbGreaterEqual(v *Value) bool {
	v_0 := v.Args[0]
	// match: (GreaterEqual (FlagConstant [fc]))
	// result: (MOVWconst [b2i32(fc.ge())])
	for {
		if v_0.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(fc.ge()))
		return true
	}
	// match: (GreaterEqual (InvertFlags x))
	// result: (LessEqual x)
	for {
		if v_0.Op != OpThumbInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbLessEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbGreaterEqualU(v *Value) bool {
	v_0 := v.Args[0]
	// match: (GreaterEqualU (FlagConstant [fc]))
	// result: (MOVWconst [b2i32(fc.uge())])
	for {
		if v_0.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(fc.uge()))
		return true
	}
	// match: (GreaterEqualU (InvertFlags x))
	// result: (LessEqualU x)
	for {
		if v_0.Op != OpThumbInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbLessEqualU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbGreaterThan(v *Value) bool {
	v_0 := v.Args[0]
	// match: (GreaterThan (FlagConstant [fc]))
	// result: (MOVWconst [b2i32(fc.gt())])
	for {
		if v_0.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(fc.gt()))
		return true
	}
	// match: (GreaterThan (InvertFlags x))
	// result: (LessThan x)
	for {
		if v_0.Op != OpThumbInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbLessThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbGreaterThanU(v *Value) bool {
	v_0 := v.Args[0]
	// match: (GreaterThanU (FlagConstant [fc]))
	// result: (MOVWconst [b2i32(fc.ugt())])
	for {
		if v_0.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(fc.ugt()))
		return true
	}
	// match: (GreaterThanU (InvertFlags x))
	// result: (LessThanU x)
	for {
		if v_0.Op != OpThumbInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbLessThanU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLessEqual(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LessEqual (FlagConstant [fc]))
	// result: (MOVWconst [b2i32(fc.le())])
	for {
		if v_0.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(fc.le()))
		return true
	}
	// match: (LessEqual (InvertFlags x))
	// result: (GreaterEqual x)
	for {
		if v_0.Op != OpThumbInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbGreaterEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLessEqualU(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LessEqualU (FlagConstant [fc]))
	// result: (MOVWconst [b2i32(fc.ule())])
	for {
		if v_0.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(fc.ule()))
		return true
	}
	// match: (LessEqualU (InvertFlags x))
	// result: (GreaterEqualU x)
	for {
		if v_0.Op != OpThumbInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbGreaterEqualU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLessThan(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LessThan (FlagConstant [fc]))
	// result: (MOVWconst [b2i32(fc.lt())])
	for {
		if v_0.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(fc.lt()))
		return true
	}
	// match: (LessThan (InvertFlags x))
	// result: (GreaterThan x)
	for {
		if v_0.Op != OpThumbInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbGreaterThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLessThanU(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LessThanU (FlagConstant [fc]))
	// result: (MOVWconst [b2i32(fc.ult())])
	for {
		if v_0.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(fc.ult()))
		return true
	}
	// match: (LessThanU (InvertFlags x))
	// result: (GreaterThanU x)
	for {
		if v_0.Op != OpThumbInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbGreaterThanU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoadOnce16 [off1] (ADDconst [off2] ptr) mem)
	// result: (LoadOnce16 [off1+off2] ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbLoadOnce16)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce16 [off1] (SUBconst [off2] ptr) mem)
	// result: (LoadOnce16 [off1-off2] ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbLoadOnce16)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce16 [0] (ADD ptr idx) mem)
	// result: (LoadOnce16idx ptr idx mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbLoadOnce16idx)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (LoadOnce16 [0] (ADDshiftLL ptr idx [c]) mem)
	// cond: c <= 3
	// result: (LoadOnce16shiftLL ptr idx [c] mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce16shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce16idx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoadOnce16idx ptr (MOVWconst [c]) mem)
	// result: (LoadOnce16 [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbLoadOnce16)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce16idx (MOVWconst [c]) ptr mem)
	// result: (LoadOnce16 [c] ptr mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		mem := v_2
		v.reset(OpThumbLoadOnce16)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce16idx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (LoadOnce16shiftLL ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce16shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoadOnce32 [off1] (ADDconst [off2] ptr) mem)
	// result: (LoadOnce32 [off1+off2] ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbLoadOnce32)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce32 [off1] (SUBconst [off2] ptr) mem)
	// result: (LoadOnce32 [off1-off2] ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbLoadOnce32)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce32 [0] (ADD ptr idx) mem)
	// result: (LoadOnce32idx ptr idx mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbLoadOnce32idx)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (LoadOnce32 [0] (ADDshiftLL ptr idx [c]) mem)
	// cond: c <= 3
	// result: (LoadOnce32shiftLL ptr idx [c] mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce32shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce32idx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoadOnce32idx ptr (MOVWconst [c]) mem)
	// result: (LoadOnce32 [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbLoadOnce32)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce32idx (MOVWconst [c]) ptr mem)
	// result: (LoadOnce32 [c] ptr mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		mem := v_2
		v.reset(OpThumbLoadOnce32)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce32idx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (LoadOnce32shiftLL ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce32shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoadOnce8 [off1] (ADDconst [off2] ptr) mem)
	// result: (LoadOnce8 [off1+off2] ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce8 [off1] (SUBconst [off2] ptr) mem)
	// result: (LoadOnce8 [off1-off2] ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce8 [0] (ADD ptr idx) mem)
	// result: (LoadOnce8idx ptr idx mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbLoadOnce8idx)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (LoadOnce8 [0] (ADDshiftLL ptr idx [c]) mem)
	// cond: c <= 3
	// result: (LoadOnce8shiftLL ptr idx [c] mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce8shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce8idx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoadOnce8idx ptr (MOVWconst [c]) mem)
	// result: (LoadOnce8 [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce8idx (MOVWconst [c]) ptr mem)
	// result: (LoadOnce8 [c] ptr mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		mem := v_2
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (LoadOnce8idx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (LoadOnce8shiftLL ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce8shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBUload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVBUload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVBUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVBUload [off1-off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVBUload)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVBUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBUreg x)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVBstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVBUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVBUloadidx ptr idx mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVBUloadidx)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVBUload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (MOVBUloadshiftLL ptr idx [c] mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVBUloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVBUload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVWconst [int32(read8(sym, int64(off)))])
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSB || !(symIsRO(sym)) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(read8(sym, int64(off))))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBUloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBUloadidx ptr idx (MOVBstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVBUreg x)
	for {
		ptr := v_0
		idx := v_1
		if v_2.Op != OpThumbMOVBstoreidx {
			break
		}
		x := v_2.Args[2]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] || !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVBUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUloadidx ptr (MOVWconst [c]) mem)
	// result: (MOVBUload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbMOVBUload)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUloadidx (MOVWconst [c]) ptr mem)
	// result: (MOVBUload [c] ptr mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		mem := v_2
		v.reset(OpThumbMOVBUload)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUloadidx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (MOVBUloadshiftLL ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBUloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVBUloadidx (SLLconst idx [c]) ptr mem)
	// cond: c <= 3
	// result: (MOVBUloadshiftLL ptr idx [c] mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBUloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBUloadshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBUloadshiftLL ptr idx [c] (MOVBstoreshiftLL ptr2 idx [d] x _))
	// cond: c==d && isSamePtr(ptr, ptr2)
	// result: (MOVBUreg x)
	for {
		c := auxIntToInt32(v.AuxInt)
		ptr := v_0
		idx := v_1
		if v_2.Op != OpThumbMOVBstoreshiftLL {
			break
		}
		d := auxIntToInt32(v_2.AuxInt)
		x := v_2.Args[2]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] || !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVBUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUloadshiftLL ptr (MOVWconst [c]) [d] mem)
	// result: (MOVBUload [int32(uint32(c)<<uint64(d))] ptr mem)
	for {
		d := auxIntToInt32(v.AuxInt)
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbMOVBUload)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) << uint64(d)))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBUreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVBUreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0xFF
	// result: y
	for {
		y := v_0
		if y.Op != OpThumbANDconst {
			break
		}
		c := auxIntToInt32(y.AuxInt)
		if !(uint64(c) <= 0xFF) {
			break
		}
		v.copyOf(y)
		return true
	}
	// match: (MOVBUreg (MOVBreg x))
	// result: (MOVBUreg x)
	for {
		if v_0.Op != OpThumbMOVBreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMOVBUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (MOVBUreg x))
	// result: (MOVBUreg x)
	for {
		if v_0.Op != OpThumbMOVBUreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMOVBUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg <t> (SRLconst [c] (MOVBUreg x)))
	// result: (SRLconst [c] (MOVBUreg <t> x))
	for {
		t := v.Type
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbMOVBUreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBUreg, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVBUreg (BFXU [c] x))
	// cond: (c>>8 <= 8)
	// result: (BFXU [c] x)
	for {
		if v_0.Op != OpThumbBFXU {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c>>8 <= 8) {
			break
		}
		v.reset(OpThumbBFXU)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (SRLconst [c] x))
	// cond: c>=24
	// result: (SRLconst [c] x)
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c >= 24) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBUload {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg x:(MOVBUloadidx _ _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBUloadidx {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (ANDconst [c] x))
	// result: (ANDconst [c&0xff] x)
	for {
		if v_0.Op != OpThumbANDconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c & 0xff)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBUreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (MOVWconst [c]))
	// result: (MOVWconst [int32(uint8(c))])
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(uint8(c)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVBload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVBload [off1-off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVBload)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVBload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBreg x)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVBstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVBloadidx ptr idx mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVBloadidx)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVBload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (MOVBloadshiftLL ptr idx [c] mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVBloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBloadidx ptr idx (MOVBstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVBreg x)
	for {
		ptr := v_0
		idx := v_1
		if v_2.Op != OpThumbMOVBstoreidx {
			break
		}
		x := v_2.Args[2]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] || !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBloadidx ptr (MOVWconst [c]) mem)
	// result: (MOVBload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbMOVBload)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBloadidx (MOVWconst [c]) ptr mem)
	// result: (MOVBload [c] ptr mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		mem := v_2
		v.reset(OpThumbMOVBload)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBloadidx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (MOVBloadshiftLL ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVBloadidx (SLLconst idx [c]) ptr mem)
	// cond: c <= 3
	// result: (MOVBloadshiftLL ptr idx [c] mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBloadshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBloadshiftLL ptr idx [c] (MOVBstoreshiftLL ptr2 idx [d] x _))
	// cond: c==d && isSamePtr(ptr, ptr2)
	// result: (MOVBreg x)
	for {
		c := auxIntToInt32(v.AuxInt)
		ptr := v_0
		idx := v_1
		if v_2.Op != OpThumbMOVBstoreshiftLL {
			break
		}
		d := auxIntToInt32(v_2.AuxInt)
		x := v_2.Args[2]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] || !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBloadshiftLL ptr (MOVWconst [c]) [d] mem)
	// result: (MOVBload [int32(uint32(c)<<uint64(d))] ptr mem)
	for {
		d := auxIntToInt32(v.AuxInt)
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbMOVBload)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) << uint64(d)))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVBreg (MOVBUreg x))
	// result: (MOVBreg x)
	for {
		if v_0.Op != OpThumbMOVBUreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVBreg x))
	// result: (MOVBreg x)
	for {
		if v_0.Op != OpThumbMOVBreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg <t> (SRAconst [c] (MOVBreg x)))
	// result: (SRAconst [c] (MOVBreg <t> x))
	for {
		t := v.Type
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbMOVBreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBreg, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVBreg (BFX [c] x))
	// cond: (c>>8 <= 8)
	// result: (BFX [c] x)
	for {
		if v_0.Op != OpThumbBFX {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c>>8 <= 8) {
			break
		}
		v.reset(OpThumbBFX)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (SRLconst [c] x))
	// cond: c>24
	// result: (SRLconst [c] x)
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c > 24) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (SRLconst [c] x))
	// cond: c==24
	// result: (SRAconst [c] x)
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c == 24) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (SRAconst [c] x))
	// cond: c>=24
	// result: (SRAconst [c] x)
	for {
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c >= 24) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBload {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBloadidx _ _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBloadidx {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (ANDconst [c] x))
	// cond: c & 0x80 == 0
	// result: (ANDconst [c&0x7f] x)
	for {
		if v_0.Op != OpThumbANDconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c&0x80 == 0) {
			break
		}
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c & 0x7f)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVWconst [c]))
	// result: (MOVWconst [int32(int8(c))])
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(int8(c)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// result: (MOVBstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBUreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHUreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (MOVBstoreidx ptr idx val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVBstoreidx)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (MOVBstore [0] {sym} (ADDshiftLL ptr idx [c]) val mem)
	// cond: sym == nil && c <= 3
	// result: (MOVBstoreshiftLL ptr idx [c] val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVBstoreshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstoreidx ptr (MOVWconst [c]) val mem)
	// result: (MOVBstore [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		val := v_2
		mem := v_3
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstoreidx (MOVWconst [c]) ptr val mem)
	// result: (MOVBstore [c] ptr val mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstoreidx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (MOVBstoreshiftLL ptr idx [c] val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBstoreshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (MOVBstoreidx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (MOVBstoreshiftLL ptr idx [c] val mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBstoreshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBstoreshiftLL(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstoreshiftLL ptr (MOVWconst [c]) [d] val mem)
	// result: (MOVBstore [int32(uint32(c)<<uint64(d))] ptr val mem)
	for {
		d := auxIntToInt32(v.AuxInt)
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		val := v_2
		mem := v_3
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) << uint64(d)))
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVDload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVDload [off1-off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVDload)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVDload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDload [off] {sym} ptr (MOVDstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVDstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbMOVDstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVDstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// result: (MOVDstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbMOVDstore)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVDstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVFload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVFload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVFload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVFload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVFload [off1-off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVFload)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVFload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVFload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVFload [off] {sym} ptr (MOVFstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVFstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVFstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// result: (MOVFstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbMOVFstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVFstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// result: (MOVFstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbMOVFstore)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVFstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVFstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (MOVHUload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVHUload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVHUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVHUload [off1-off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVHUload)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVHUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVHUreg x)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVHstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVHUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVHUloadidx ptr idx mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVHUloadidx)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVHUload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (MOVHUloadshiftLL ptr idx [c] mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVHUloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVHUload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVWconst [int32(read16(sym, int64(off), config.ctxt.Arch.ByteOrder))])
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSB || !(symIsRO(sym)) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(read16(sym, int64(off), config.ctxt.Arch.ByteOrder)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHUloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHUloadidx ptr idx (MOVHstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVHUreg x)
	for {
		ptr := v_0
		idx := v_1
		if v_2.Op != OpThumbMOVHstoreidx {
			break
		}
		x := v_2.Args[2]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] || !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVHUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUloadidx ptr (MOVWconst [c]) mem)
	// result: (MOVHUload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbMOVHUload)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUloadidx (MOVWconst [c]) ptr mem)
	// result: (MOVHUload [c] ptr mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		mem := v_2
		v.reset(OpThumbMOVHUload)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUloadidx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (MOVHUloadshiftLL ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHUloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVHUloadidx (SLLconst idx [c]) ptr mem)
	// cond: c <= 3
	// result: (MOVHUloadshiftLL ptr idx [c] mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHUloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHUloadshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHUloadshiftLL ptr idx [c] (MOVHstoreshiftLL ptr2 idx [d] x _))
	// cond: c==d && isSamePtr(ptr, ptr2)
	// result: (MOVHUreg x)
	for {
		c := auxIntToInt32(v.AuxInt)
		ptr := v_0
		idx := v_1
		if v_2.Op != OpThumbMOVHstoreshiftLL {
			break
		}
		d := auxIntToInt32(v_2.AuxInt)
		x := v_2.Args[2]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] || !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVHUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUloadshiftLL ptr (MOVWconst [c]) [d] mem)
	// result: (MOVHUload [int32(uint32(c)<<uint64(d))] ptr mem)
	for {
		d := auxIntToInt32(v.AuxInt)
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbMOVHUload)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) << uint64(d)))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHUreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVHUreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0xFFFF
	// result: y
	for {
		y := v_0
		if y.Op != OpThumbANDconst {
			break
		}
		c := auxIntToInt32(y.AuxInt)
		if !(uint64(c) <= 0xFFFF) {
			break
		}
		v.copyOf(y)
		return true
	}
	// match: (MOVHUreg (MOVBUreg x))
	// result: (MOVBUreg x)
	for {
		if v_0.Op != OpThumbMOVBUreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMOVBUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (MOVHUreg x))
	// result: (MOVHUreg x)
	for {
		if v_0.Op != OpThumbMOVHUreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMOVHUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (MOVHreg x))
	// result: (MOVHUreg x)
	for {
		if v_0.Op != OpThumbMOVHreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMOVHUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg <t> (SRLconst [c] (MOVBUreg x)))
	// result: (SRLconst [c] (MOVBUreg <t> x))
	for {
		t := v.Type
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbMOVBUreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBUreg, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVHUreg <t> (SRLconst [c] (MOVHUreg x)))
	// result: (SRLconst [c] (MOVHUreg <t> x))
	for {
		t := v.Type
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbMOVHUreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbMOVHUreg, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVHUreg (BFXU [c] x))
	// cond: (c>>8 <= 8)
	// result: (BFXU [c] x)
	for {
		if v_0.Op != OpThumbBFXU {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c>>8 <= 8) {
			break
		}
		v.reset(OpThumbBFXU)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (BFXU [c] x))
	// cond: (c>>8 <= 16)
	// result: (BFXU [c] x)
	for {
		if v_0.Op != OpThumbBFXU {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c>>8 <= 16) {
			break
		}
		v.reset(OpThumbBFXU)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (SRLconst [c] x))
	// cond: c>=16
	// result: (SRLconst [c] x)
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c >= 16) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBUload {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVHUload {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVBUloadidx _ _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBUloadidx {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUloadidx _ _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVHUloadidx {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (ANDconst [c] x))
	// result: (ANDconst [c&0xffff] x)
	for {
		if v_0.Op != OpThumbANDconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c & 0xffff)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBUreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVHUreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (MOVWconst [c]))
	// result: (MOVWconst [int32(uint16(c))])
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(uint16(c)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVHload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVHload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVHload [off1-off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVHload)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVHload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVHreg x)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVHstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVHloadidx ptr idx mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVHloadidx)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVHload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (MOVHloadshiftLL ptr idx [c] mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVHloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHloadidx ptr idx (MOVHstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVHreg x)
	for {
		ptr := v_0
		idx := v_1
		if v_2.Op != OpThumbMOVHstoreidx {
			break
		}
		x := v_2.Args[2]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] || !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHloadidx ptr (MOVWconst [c]) mem)
	// result: (MOVHload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbMOVHload)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHloadidx (MOVWconst [c]) ptr mem)
	// result: (MOVHload [c] ptr mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		mem := v_2
		v.reset(OpThumbMOVHload)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHloadidx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (MOVHloadshiftLL ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVHloadidx (SLLconst idx [c]) ptr mem)
	// cond: c <= 3
	// result: (MOVHloadshiftLL ptr idx [c] mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHloadshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHloadshiftLL ptr idx [c] (MOVHstoreshiftLL ptr2 idx [d] x _))
	// cond: c==d && isSamePtr(ptr, ptr2)
	// result: (MOVHreg x)
	for {
		c := auxIntToInt32(v.AuxInt)
		ptr := v_0
		idx := v_1
		if v_2.Op != OpThumbMOVHstoreshiftLL {
			break
		}
		d := auxIntToInt32(v_2.AuxInt)
		x := v_2.Args[2]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] || !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHloadshiftLL ptr (MOVWconst [c]) [d] mem)
	// result: (MOVHload [int32(uint32(c)<<uint64(d))] ptr mem)
	for {
		d := auxIntToInt32(v.AuxInt)
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbMOVHload)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) << uint64(d)))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVHreg (MOVBreg x))
	// result: (MOVBreg x)
	for {
		if v_0.Op != OpThumbMOVBreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVHreg x))
	// result: (MOVHreg x)
	for {
		if v_0.Op != OpThumbMOVHreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVHUreg x))
	// result: (MOVHreg x)
	for {
		if v_0.Op != OpThumbMOVHUreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg <t> (SRAconst [c] (MOVBreg x)))
	// result: (SRAconst [c] (MOVBreg <t> x))
	for {
		t := v.Type
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbMOVBreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBreg, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVHreg <t> (SRAconst [c] (MOVHreg x)))
	// result: (SRAconst [c] (MOVHreg <t> x))
	for {
		t := v.Type
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbMOVHreg {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbMOVHreg, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MOVHreg (BFX [c] x))
	// cond: (c>>8 <= 8)
	// result: (BFX [c] x)
	for {
		if v_0.Op != OpThumbBFX {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c>>8 <= 8) {
			break
		}
		v.reset(OpThumbBFX)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (BFX [c] x))
	// cond: (c>>8 <= 16)
	// result: (BFX [c] x)
	for {
		if v_0.Op != OpThumbBFX {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c>>8 <= 16) {
			break
		}
		v.reset(OpThumbBFX)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (SRLconst [c] x))
	// cond: c>16
	// result: (SRLconst [c] x)
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c > 16) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (SRAconst [c] x))
	// cond: c>=16
	// result: (SRAconst [c] x)
	for {
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c >= 16) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (SRLconst [c] x))
	// cond: c==16
	// result: (SRAconst [c] x)
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c == 16) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBload {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBUload {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVHload {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBloadidx _ _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBloadidx {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUloadidx _ _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBUloadidx {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHloadidx _ _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVHloadidx {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (ANDconst [c] x))
	// cond: c & 0x8000 == 0
	// result: (ANDconst [c&0x7fff] x)
	for {
		if v_0.Op != OpThumbANDconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c&0x8000 == 0) {
			break
		}
		v.reset(OpThumbANDconst)
		v.AuxInt = int32ToAuxInt(c & 0x7fff)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVBUreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpThumbMOVHreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVWconst [c]))
	// result: (MOVWconst [int32(int16(c))])
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(int16(c)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbMOVHstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// result: (MOVHstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbMOVHstore)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpThumbMOVHstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHUreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpThumbMOVHstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVHstore [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (MOVHstoreidx ptr idx val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVHstoreidx)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (MOVHstore [0] {sym} (ADDshiftLL ptr idx [c]) val mem)
	// cond: sym == nil && c <= 3
	// result: (MOVHstoreshiftLL ptr idx [c] val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVHstoreshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstoreidx ptr (MOVWconst [c]) val mem)
	// result: (MOVHstore [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		val := v_2
		mem := v_3
		v.reset(OpThumbMOVHstore)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstoreidx (MOVWconst [c]) ptr val mem)
	// result: (MOVHstore [c] ptr val mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpThumbMOVHstore)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstoreidx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (MOVHstoreshiftLL ptr idx [c] val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHstoreshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (MOVHstoreidx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (MOVHstoreshiftLL ptr idx [c] val mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHstoreshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHstoreshiftLL(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstoreshiftLL ptr (MOVWconst [c]) [d] val mem)
	// result: (MOVHstore [int32(uint32(c)<<uint64(d))] ptr val mem)
	for {
		d := auxIntToInt32(v.AuxInt)
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		val := v_2
		mem := v_3
		v.reset(OpThumbMOVHstore)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) << uint64(d)))
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (MOVWload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVWload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVWload [off1-off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpThumbMOVWload)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVWload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWload [off] {sym} ptr (MOVWstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpThumbMOVWstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (MOVWload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVWloadidx ptr idx mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVWloadidx)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVWload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (MOVWloadshiftLL ptr idx [c] mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVWloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVWload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVWconst [int32(read32(sym, int64(off), config.ctxt.Arch.ByteOrder))])
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSB || !(symIsRO(sym)) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(read32(sym, int64(off), config.ctxt.Arch.ByteOrder)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWloadidx ptr idx (MOVWstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: x
	for {
		ptr := v_0
		idx := v_1
		if v_2.Op != OpThumbMOVWstoreidx {
			break
		}
		x := v_2.Args[2]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] || !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (MOVWloadidx ptr (MOVWconst [c]) mem)
	// result: (MOVWload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbMOVWload)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWloadidx (MOVWconst [c]) ptr mem)
	// result: (MOVWload [c] ptr mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		mem := v_2
		v.reset(OpThumbMOVWload)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWloadidx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (MOVWloadshiftLL ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVWloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	// match: (MOVWloadidx (SLLconst idx [c]) ptr mem)
	// cond: c <= 3
	// result: (MOVWloadshiftLL ptr idx [c] mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVWloadshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, idx, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWloadshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWloadshiftLL ptr idx [c] (MOVWstoreshiftLL ptr2 idx [d] x _))
	// cond: c==d && isSamePtr(ptr, ptr2)
	// result: x
	for {
		c := auxIntToInt32(v.AuxInt)
		ptr := v_0
		idx := v_1
		if v_2.Op != OpThumbMOVWstoreshiftLL {
			break
		}
		d := auxIntToInt32(v_2.AuxInt)
		x := v_2.Args[2]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] || !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (MOVWloadshiftLL ptr (MOVWconst [c]) [d] mem)
	// result: (MOVWload [int32(uint32(c)<<uint64(d))] ptr mem)
	for {
		d := auxIntToInt32(v.AuxInt)
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpThumbMOVWload)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) << uint64(d)))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVWreg x)
	// cond: x.Uses == 1
	// result: (MOVWnop x)
	for {
		x := v_0
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpThumbMOVWnop)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (MOVWconst [c]))
	// result: (MOVWconst [c])
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(c)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbMOVWstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// result: (MOVWstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbMOVWstore)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVWstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (MOVWstoreidx ptr idx val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVWstoreidx)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (MOVWstore [0] {sym} (ADDshiftLL ptr idx [c]) val mem)
	// cond: sym == nil && c <= 3
	// result: (MOVWstoreshiftLL ptr idx [c] val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		sym := auxToSym(v.Aux)
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVWstoreshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstoreidx ptr (MOVWconst [c]) val mem)
	// result: (MOVWstore [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		val := v_2
		mem := v_3
		v.reset(OpThumbMOVWstore)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstoreidx (MOVWconst [c]) ptr val mem)
	// result: (MOVWstore [c] ptr val mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpThumbMOVWstore)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstoreidx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (MOVWstoreshiftLL ptr idx [c] val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVWstoreshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (MOVWstoreidx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (MOVWstoreshiftLL ptr idx [c] val mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVWstoreshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWstoreshiftLL(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstoreshiftLL ptr (MOVWconst [c]) [d] val mem)
	// result: (MOVWstore [int32(uint32(c)<<uint64(d))] ptr val mem)
	for {
		d := auxIntToInt32(v.AuxInt)
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		val := v_2
		mem := v_3
		v.reset(OpThumbMOVWstore)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) << uint64(d)))
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMUL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (MUL x (MOVWconst [c]))
	// cond: int32(c) == -1
	// result: (RSBconst [0] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			if !(int32(c) == -1) {
				continue
			}
			v.reset(OpThumbRSBconst)
			v.AuxInt = int32ToAuxInt(0)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL _ (MOVWconst [0]))
	// result: (MOVWconst [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_1.Op != OpThumbMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
				continue
			}
			v.reset(OpThumbMOVWconst)
			v.AuxInt = int32ToAuxInt(0)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [1]))
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst || auxIntToInt32(v_1.AuxInt) != 1 {
				continue
			}
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: isPowerOfTwo32(c)
	// result: (SLLconst [int32(log32(c))] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			if !(isPowerOfTwo32(c)) {
				continue
			}
			v.reset(OpThumbSLLconst)
			v.AuxInt = int32ToAuxInt(int32(log32(c)))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: isPowerOfTwo32(c-1) && c >= 3
	// result: (ADDshiftLL x x [int32(log32(c-1))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			if !(isPowerOfTwo32(c-1) && c >= 3) {
				continue
			}
			v.reset(OpThumbADDshiftLL)
			v.AuxInt = int32ToAuxInt(int32(log32(c - 1)))
			v.AddArg2(x, x)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: isPowerOfTwo32(c+1) && c >= 7
	// result: (RSBshiftLL x x [int32(log32(c+1))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			if !(isPowerOfTwo32(c+1) && c >= 7) {
				continue
			}
			v.reset(OpThumbRSBshiftLL)
			v.AuxInt = int32ToAuxInt(int32(log32(c + 1)))
			v.AddArg2(x, x)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%3 == 0 && isPowerOfTwo32(c/3)
	// result: (SLLconst [int32(log32(c/3))] (ADDshiftLL <x.Type> x x [1]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			if !(c%3 == 0 && isPowerOfTwo32(c/3)) {
				continue
			}
			v.reset(OpThumbSLLconst)
			v.AuxInt = int32ToAuxInt(int32(log32(c / 3)))
			v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
			v0.AuxInt = int32ToAuxInt(1)
			v0.AddArg2(x, x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%5 == 0 && isPowerOfTwo32(c/5)
	// result: (SLLconst [int32(log32(c/5))] (ADDshiftLL <x.Type> x x [2]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			if !(c%5 == 0 && isPowerOfTwo32(c/5)) {
				continue
			}
			v.reset(OpThumbSLLconst)
			v.AuxInt = int32ToAuxInt(int32(log32(c / 5)))
			v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
			v0.AuxInt = int32ToAuxInt(2)
			v0.AddArg2(x, x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%7 == 0 && isPowerOfTwo32(c/7)
	// result: (SLLconst [int32(log32(c/7))] (RSBshiftLL <x.Type> x x [3]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			if !(c%7 == 0 && isPowerOfTwo32(c/7)) {
				continue
			}
			v.reset(OpThumbSLLconst)
			v.AuxInt = int32ToAuxInt(int32(log32(c / 7)))
			v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
			v0.AuxInt = int32ToAuxInt(3)
			v0.AddArg2(x, x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%9 == 0 && isPowerOfTwo32(c/9)
	// result: (SLLconst [int32(log32(c/9))] (ADDshiftLL <x.Type> x x [3]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			if !(c%9 == 0 && isPowerOfTwo32(c/9)) {
				continue
			}
			v.reset(OpThumbSLLconst)
			v.AuxInt = int32ToAuxInt(int32(log32(c / 9)))
			v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
			v0.AuxInt = int32ToAuxInt(3)
			v0.AddArg2(x, x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (MUL (MOVWconst [c]) (MOVWconst [d]))
	// result: (MOVWconst [c*d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_0.AuxInt)
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			d := auxIntToInt32(v_1.AuxInt)
			v.reset(OpThumbMOVWconst)
			v.AuxInt = int32ToAuxInt(c * d)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbMULA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c == -1
	// result: (SUB a x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(c == -1) {
			break
		}
		v.reset(OpThumbSUB)
		v.AddArg2(a, x)
		return true
	}
	// match: (MULA _ (MOVWconst [0]) a)
	// result: a
	for {
		if v_1.Op != OpThumbMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		a := v_2
		v.copyOf(a)
		return true
	}
	// match: (MULA x (MOVWconst [1]) a)
	// result: (ADD x a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst || auxIntToInt32(v_1.AuxInt) != 1 {
			break
		}
		a := v_2
		v.reset(OpThumbADD)
		v.AddArg2(x, a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: isPowerOfTwo32(c)
	// result: (ADD (SLLconst <x.Type> [int32(log32(c))] x) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(isPowerOfTwo32(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c)))
		v0.AddArg(x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: isPowerOfTwo32(c-1) && c >= 3
	// result: (ADD (ADDshiftLL <x.Type> x x [int32(log32(c-1))]) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(isPowerOfTwo32(c-1) && c >= 3) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c - 1)))
		v0.AddArg2(x, x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: isPowerOfTwo32(c+1) && c >= 7
	// result: (ADD (RSBshiftLL <x.Type> x x [int32(log32(c+1))]) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(isPowerOfTwo32(c+1) && c >= 7) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c + 1)))
		v0.AddArg2(x, x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%3 == 0 && isPowerOfTwo32(c/3)
	// result: (ADD (SLLconst <x.Type> [int32(log32(c/3))] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(c%3 == 0 && isPowerOfTwo32(c/3)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 3)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(1)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%5 == 0 && isPowerOfTwo32(c/5)
	// result: (ADD (SLLconst <x.Type> [int32(log32(c/5))] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(c%5 == 0 && isPowerOfTwo32(c/5)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 5)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(2)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%7 == 0 && isPowerOfTwo32(c/7)
	// result: (ADD (SLLconst <x.Type> [int32(log32(c/7))] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(c%7 == 0 && isPowerOfTwo32(c/7)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 7)))
		v1 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(3)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%9 == 0 && isPowerOfTwo32(c/9)
	// result: (ADD (SLLconst <x.Type> [int32(log32(c/9))] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(c%9 == 0 && isPowerOfTwo32(c/9)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 9)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(3)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c == -1
	// result: (SUB a x)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(c == -1) {
			break
		}
		v.reset(OpThumbSUB)
		v.AddArg2(a, x)
		return true
	}
	// match: (MULA (MOVWconst [0]) _ a)
	// result: a
	for {
		if v_0.Op != OpThumbMOVWconst || auxIntToInt32(v_0.AuxInt) != 0 {
			break
		}
		a := v_2
		v.copyOf(a)
		return true
	}
	// match: (MULA (MOVWconst [1]) x a)
	// result: (ADD x a)
	for {
		if v_0.Op != OpThumbMOVWconst || auxIntToInt32(v_0.AuxInt) != 1 {
			break
		}
		x := v_1
		a := v_2
		v.reset(OpThumbADD)
		v.AddArg2(x, a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: isPowerOfTwo32(c)
	// result: (ADD (SLLconst <x.Type> [int32(log32(c))] x) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(isPowerOfTwo32(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c)))
		v0.AddArg(x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: isPowerOfTwo32(c-1) && c >= 3
	// result: (ADD (ADDshiftLL <x.Type> x x [int32(log32(c-1))]) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(isPowerOfTwo32(c-1) && c >= 3) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c - 1)))
		v0.AddArg2(x, x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: isPowerOfTwo32(c+1) && c >= 7
	// result: (ADD (RSBshiftLL <x.Type> x x [int32(log32(c+1))]) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(isPowerOfTwo32(c+1) && c >= 7) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c + 1)))
		v0.AddArg2(x, x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%3 == 0 && isPowerOfTwo32(c/3)
	// result: (ADD (SLLconst <x.Type> [int32(log32(c/3))] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(c%3 == 0 && isPowerOfTwo32(c/3)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 3)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(1)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%5 == 0 && isPowerOfTwo32(c/5)
	// result: (ADD (SLLconst <x.Type> [int32(log32(c/5))] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(c%5 == 0 && isPowerOfTwo32(c/5)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 5)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(2)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%7 == 0 && isPowerOfTwo32(c/7)
	// result: (ADD (SLLconst <x.Type> [int32(log32(c/7))] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(c%7 == 0 && isPowerOfTwo32(c/7)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 7)))
		v1 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(3)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%9 == 0 && isPowerOfTwo32(c/9)
	// result: (ADD (SLLconst <x.Type> [int32(log32(c/9))] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(c%9 == 0 && isPowerOfTwo32(c/9)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 9)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(3)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULA (MOVWconst [c]) (MOVWconst [d]) a)
	// result: (ADDconst [c*d] a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_1.AuxInt)
		a := v_2
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(c * d)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMULD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MULD (NEGD x) y)
	// result: (NMULD x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpThumbNEGD {
				continue
			}
			x := v_0.Args[0]
			y := v_1
			v.reset(OpThumbNMULD)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbMULF(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MULF (NEGF x) y)
	// result: (NMULF x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpThumbNEGF {
				continue
			}
			x := v_0.Args[0]
			y := v_1
			v.reset(OpThumbNMULF)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbMULS(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c == -1
	// result: (ADD a x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(c == -1) {
			break
		}
		v.reset(OpThumbADD)
		v.AddArg2(a, x)
		return true
	}
	// match: (MULS _ (MOVWconst [0]) a)
	// result: a
	for {
		if v_1.Op != OpThumbMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		a := v_2
		v.copyOf(a)
		return true
	}
	// match: (MULS x (MOVWconst [1]) a)
	// result: (RSB x a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst || auxIntToInt32(v_1.AuxInt) != 1 {
			break
		}
		a := v_2
		v.reset(OpThumbRSB)
		v.AddArg2(x, a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: isPowerOfTwo32(c)
	// result: (RSB (SLLconst <x.Type> [int32(log32(c))] x) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(isPowerOfTwo32(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c)))
		v0.AddArg(x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: isPowerOfTwo32(c-1) && c >= 3
	// result: (RSB (ADDshiftLL <x.Type> x x [int32(log32(c-1))]) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(isPowerOfTwo32(c-1) && c >= 3) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c - 1)))
		v0.AddArg2(x, x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: isPowerOfTwo32(c+1) && c >= 7
	// result: (RSB (RSBshiftLL <x.Type> x x [int32(log32(c+1))]) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(isPowerOfTwo32(c+1) && c >= 7) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c + 1)))
		v0.AddArg2(x, x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%3 == 0 && isPowerOfTwo32(c/3)
	// result: (RSB (SLLconst <x.Type> [int32(log32(c/3))] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(c%3 == 0 && isPowerOfTwo32(c/3)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 3)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(1)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%5 == 0 && isPowerOfTwo32(c/5)
	// result: (RSB (SLLconst <x.Type> [int32(log32(c/5))] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(c%5 == 0 && isPowerOfTwo32(c/5)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 5)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(2)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%7 == 0 && isPowerOfTwo32(c/7)
	// result: (RSB (SLLconst <x.Type> [int32(log32(c/7))] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(c%7 == 0 && isPowerOfTwo32(c/7)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 7)))
		v1 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(3)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%9 == 0 && isPowerOfTwo32(c/9)
	// result: (RSB (SLLconst <x.Type> [int32(log32(c/9))] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		a := v_2
		if !(c%9 == 0 && isPowerOfTwo32(c/9)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 9)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(3)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c == -1
	// result: (ADD a x)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(c == -1) {
			break
		}
		v.reset(OpThumbADD)
		v.AddArg2(a, x)
		return true
	}
	// match: (MULS (MOVWconst [0]) _ a)
	// result: a
	for {
		if v_0.Op != OpThumbMOVWconst || auxIntToInt32(v_0.AuxInt) != 0 {
			break
		}
		a := v_2
		v.copyOf(a)
		return true
	}
	// match: (MULS (MOVWconst [1]) x a)
	// result: (RSB x a)
	for {
		if v_0.Op != OpThumbMOVWconst || auxIntToInt32(v_0.AuxInt) != 1 {
			break
		}
		x := v_1
		a := v_2
		v.reset(OpThumbRSB)
		v.AddArg2(x, a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: isPowerOfTwo32(c)
	// result: (RSB (SLLconst <x.Type> [int32(log32(c))] x) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(isPowerOfTwo32(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c)))
		v0.AddArg(x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: isPowerOfTwo32(c-1) && c >= 3
	// result: (RSB (ADDshiftLL <x.Type> x x [int32(log32(c-1))]) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(isPowerOfTwo32(c-1) && c >= 3) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c - 1)))
		v0.AddArg2(x, x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: isPowerOfTwo32(c+1) && c >= 7
	// result: (RSB (RSBshiftLL <x.Type> x x [int32(log32(c+1))]) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(isPowerOfTwo32(c+1) && c >= 7) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c + 1)))
		v0.AddArg2(x, x)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%3 == 0 && isPowerOfTwo32(c/3)
	// result: (RSB (SLLconst <x.Type> [int32(log32(c/3))] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(c%3 == 0 && isPowerOfTwo32(c/3)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 3)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(1)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%5 == 0 && isPowerOfTwo32(c/5)
	// result: (RSB (SLLconst <x.Type> [int32(log32(c/5))] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(c%5 == 0 && isPowerOfTwo32(c/5)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 5)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(2)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%7 == 0 && isPowerOfTwo32(c/7)
	// result: (RSB (SLLconst <x.Type> [int32(log32(c/7))] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(c%7 == 0 && isPowerOfTwo32(c/7)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 7)))
		v1 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(3)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%9 == 0 && isPowerOfTwo32(c/9)
	// result: (RSB (SLLconst <x.Type> [int32(log32(c/9))] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		a := v_2
		if !(c%9 == 0 && isPowerOfTwo32(c/9)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(int32(log32(c / 9)))
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = int32ToAuxInt(3)
		v1.AddArg2(x, x)
		v0.AddArg(v1)
		v.AddArg2(v0, a)
		return true
	}
	// match: (MULS (MOVWconst [c]) (MOVWconst [d]) a)
	// result: (SUBconst [c*d] a)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_1.AuxInt)
		a := v_2
		v.reset(OpThumbSUBconst)
		v.AuxInt = int32ToAuxInt(c * d)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMVN(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MVN (MOVWconst [c]))
	// result: (MOVWconst [^c])
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(^c)
		return true
	}
	// match: (MVN (SLLconst [c] x))
	// result: (MVNshiftLL x [c])
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbMVNshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MVN (SRLconst [c] x))
	// result: (MVNshiftRL x [c])
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbMVNshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (MVN (SRAconst [c] x))
	// result: (MVNshiftRA x [c])
	for {
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbMVNshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMVNshiftLL(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MVNshiftLL (MOVWconst [c]) [d])
	// result: (MOVWconst [^(c<<uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(^(c << uint64(d)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMVNshiftRA(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MVNshiftRA (MOVWconst [c]) [d])
	// result: (MOVWconst [^(int32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(^(int32(c) >> uint64(d)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMVNshiftRL(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MVNshiftRL (MOVWconst [c]) [d])
	// result: (MOVWconst [^int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(^int32(uint32(c) >> uint64(d)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbNEGD(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NEGD (MULD x y))
	// result: (NMULD x y)
	for {
		if v_0.Op != OpThumbMULD {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpThumbNMULD)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbNEGF(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NEGF (MULF x y))
	// result: (NMULF x y)
	for {
		if v_0.Op != OpThumbMULF {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpThumbNMULF)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbNMULD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NMULD (NEGD x) y)
	// result: (MULD x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpThumbNEGD {
				continue
			}
			x := v_0.Args[0]
			y := v_1
			v.reset(OpThumbMULD)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbNMULF(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NMULF (NEGF x) y)
	// result: (MULF x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpThumbNEGF {
				continue
			}
			x := v_0.Args[0]
			y := v_1
			v.reset(OpThumbMULF)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbNotEqual(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NotEqual (FlagConstant [fc]))
	// result: (MOVWconst [b2i32(fc.ne())])
	for {
		if v_0.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(fc.ne()))
		return true
	}
	// match: (NotEqual (InvertFlags x))
	// result: (NotEqual x)
	for {
		if v_0.Op != OpThumbInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbNotEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (OR x (MOVWconst [c]))
	// result: (ORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpThumbORconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (OR x (SLLconst [c] y))
	// result: (ORshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSLLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbORshiftLL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (OR x (SRLconst [c] y))
	// result: (ORshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbORshiftRL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (OR x (SRAconst [c] y))
	// result: (ORshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRAconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbORshiftRA)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (OR x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (OR x (MVN y))
	// result: (ORN x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMVN {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpThumbORN)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbORN(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ORN x (MOVWconst [c]))
	// result: (ORNconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbORNconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (ORN x (SLLconst [c] y))
	// result: (ORNshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbORNshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (ORN x (SRLconst [c] y))
	// result: (ORNshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbORNshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (ORN x (SRAconst [c] y))
	// result: (ORNshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbORNshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (ORN x x)
	// result: (MOVWconst [-1])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(-1)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORNconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ORNconst [0] _)
	// result: (MOVWconst [-1])
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(-1)
		return true
	}
	// match: (ORNconst [c] x)
	// cond: int32(c)==-1
	// result: x
	for {
		c := auxIntToInt32(v.AuxInt)
		x := v_0
		if !(int32(c) == -1) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ORNconst [c] (MOVWconst [d]))
	// result: (MOVWconst [^c|d])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(^c | d)
		return true
	}
	// match: (ORNconst [c] (ORNconst [d] x))
	// result: (ORconst [^c|^d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbORNconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbORconst)
		v.AuxInt = int32ToAuxInt(^c | ^d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORNshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ORNshiftLL x (MOVWconst [c]) [d])
	// result: (ORNconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbORNconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORNshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ORNshiftRA x (MOVWconst [c]) [d])
	// result: (ORNconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbORNconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORNshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ORNshiftRL x (MOVWconst [c]) [d])
	// result: (ORNconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbORNconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ORconst [0] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (ORconst [c] _)
	// cond: int32(c)==-1
	// result: (MOVWconst [-1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(-1)
		return true
	}
	// match: (ORconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c|d])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(c | d)
		return true
	}
	// match: (ORconst [c] (ORconst [d] x))
	// result: (ORconst [c|d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbORconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbORconst)
		v.AuxInt = int32ToAuxInt(c | d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORshiftLL (MOVWconst [c]) x [d])
	// result: (ORconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbORconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftLL x (MOVWconst [c]) [d])
	// result: (ORconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbORconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	// match: ( ORshiftLL [c] (SRLconst x [32-c]) x)
	// result: (SRRconst [32-c] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSRLconst || auxIntToInt32(v_0.AuxInt) != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = int32ToAuxInt(32 - c)
		v.AddArg(x)
		return true
	}
	// match: (ORshiftLL <typ.UInt16> [8] (BFXU <typ.UInt16> [int32(armBFAuxInt(8, 8))] x) x)
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || auxIntToInt32(v.AuxInt) != 8 || v_0.Op != OpThumbBFXU || v_0.Type != typ.UInt16 || auxIntToInt32(v_0.AuxInt) != int32(armBFAuxInt(8, 8)) {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbREV16)
		v.AddArg(x)
		return true
	}
	// match: (ORshiftLL <typ.UInt16> [8] (SRLconst <typ.UInt16> [24] (SLLconst [16] x)) x)
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || auxIntToInt32(v.AuxInt) != 8 || v_0.Op != OpThumbSRLconst || v_0.Type != typ.UInt16 || auxIntToInt32(v_0.AuxInt) != 24 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbSLLconst || auxIntToInt32(v_0_0.AuxInt) != 16 {
			break
		}
		x := v_0_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbREV16)
		v.AddArg(x)
		return true
	}
	// match: (ORshiftLL x y:(SLLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		y := v_1
		if y.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(y.AuxInt)
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.copyOf(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ORshiftRA (MOVWconst [c]) x [d])
	// result: (ORconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbORconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftRA x (MOVWconst [c]) [d])
	// result: (ORconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbORconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (ORshiftRA x y:(SRAconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		y := v_1
		if y.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(y.AuxInt)
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.copyOf(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ORshiftRL (MOVWconst [c]) x [d])
	// result: (ORconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbORconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftRL x (MOVWconst [c]) [d])
	// result: (ORconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbORconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: ( ORshiftRL [c] (SLLconst x [32-c]) x)
	// result: (SRRconst [ c] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSLLconst || auxIntToInt32(v_0.AuxInt) != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (ORshiftRL x y:(SRLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		y := v_1
		if y.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(y.AuxInt)
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.copyOf(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RSB (MOVWconst [c]) x)
	// result: (SUBconst [c] x)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbSUBconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (RSB x (MOVWconst [c]))
	// result: (RSBconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (RSB x (SLLconst [c] y))
	// result: (RSBshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbRSBshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (RSB (SLLconst [c] y) x)
	// result: (SUBshiftLL x y [c])
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbSUBshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (RSB x (SRLconst [c] y))
	// result: (RSBshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbRSBshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (RSB (SRLconst [c] y) x)
	// result: (SUBshiftRL x y [c])
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbSUBshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (RSB x (SRAconst [c] y))
	// result: (RSBshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbRSBshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (RSB (SRAconst [c] y) x)
	// result: (SUBshiftRA x y [c])
	for {
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbSUBshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (RSB x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (RSB (MUL x y) a)
	// result: (MULS x y a)
	for {
		if v_0.Op != OpThumbMUL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		a := v_1
		v.reset(OpThumbMULS)
		v.AddArg3(x, y, a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBSshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBSshiftLL (MOVWconst [c]) x [d])
	// result: (SUBSconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbSUBSconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftLL x (MOVWconst [c]) [d])
	// result: (RSBSconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbRSBSconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBSshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBSshiftRA (MOVWconst [c]) x [d])
	// result: (SUBSconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbSUBSconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftRA x (MOVWconst [c]) [d])
	// result: (RSBSconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbRSBSconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBSshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBSshiftRL (MOVWconst [c]) x [d])
	// result: (SUBSconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbSUBSconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftRL x (MOVWconst [c]) [d])
	// result: (RSBSconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbRSBSconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (RSBconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c-d])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(c - d)
		return true
	}
	// match: (RSBconst [c] (RSBconst [d] x))
	// result: (ADDconst [c-d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbRSBconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(c - d)
		v.AddArg(x)
		return true
	}
	// match: (RSBconst [c] (ADDconst [d] x))
	// result: (RSBconst [c-d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(c - d)
		v.AddArg(x)
		return true
	}
	// match: (RSBconst [c] (SUBconst [d] x))
	// result: (RSBconst [c+d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(c + d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBshiftLL (MOVWconst [c]) x [d])
	// result: (SUBconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbSUBconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftLL x (MOVWconst [c]) [d])
	// result: (RSBconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (RSBshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBshiftRA (MOVWconst [c]) x [d])
	// result: (SUBconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbSUBconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftRA x (MOVWconst [c]) [d])
	// result: (RSBconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (RSBshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBshiftRL (MOVWconst [c]) x [d])
	// result: (SUBconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbSUBconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftRL x (MOVWconst [c]) [d])
	// result: (RSBconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (RSBshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSBC(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SBC x (MOVWconst [c]) flags)
	// result: (SBCconst [c] x flags)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		flags := v_2
		v.reset(OpThumbSBCconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, flags)
		return true
	}
	// match: (SBC x (SLLconst [c] y) flags)
	// result: (SBCshiftLL x y [c] flags)
	for {
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		flags := v_2
		v.reset(OpThumbSBCshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(x, y, flags)
		return true
	}
	// match: (SBC x (SRLconst [c] y) flags)
	// result: (SBCshiftRL x y [c] flags)
	for {
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		flags := v_2
		v.reset(OpThumbSBCshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(x, y, flags)
		return true
	}
	// match: (SBC x (SRAconst [c] y) flags)
	// result: (SBCshiftRA x y [c] flags)
	for {
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		flags := v_2
		v.reset(OpThumbSBCshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(x, y, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSBCconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SBCconst [c] (ADDconst [d] x) flags)
	// result: (SBCconst [c-d] x flags)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		flags := v_1
		v.reset(OpThumbSBCconst)
		v.AuxInt = int32ToAuxInt(c - d)
		v.AddArg2(x, flags)
		return true
	}
	// match: (SBCconst [c] (SUBconst [d] x) flags)
	// result: (SBCconst [c+d] x flags)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		flags := v_1
		v.reset(OpThumbSBCconst)
		v.AuxInt = int32ToAuxInt(c + d)
		v.AddArg2(x, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSBCshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SBCshiftLL x (MOVWconst [c]) [d] flags)
	// result: (SBCconst x [c<<uint64(d)] flags)
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		flags := v_2
		v.reset(OpThumbSBCconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg2(x, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSBCshiftRA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SBCshiftRA x (MOVWconst [c]) [d] flags)
	// result: (SBCconst x [c>>uint64(d)] flags)
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		flags := v_2
		v.reset(OpThumbSBCconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg2(x, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSBCshiftRL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SBCshiftRL x (MOVWconst [c]) [d] flags)
	// result: (SBCconst x [int32(uint32(c)>>uint64(d))] flags)
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		flags := v_2
		v.reset(OpThumbSBCconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg2(x, flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SLL x (MOVWconst [c]))
	// result: (SLLconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSLLconst)
		v.AuxInt = int32ToAuxInt(c & 31)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSLLconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SLLconst [c] (MOVWconst [d]))
	// result: (MOVWconst [d<<uint64(c)])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(d << uint64(c))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRA x (MOVWconst [c]))
	// result: (SRAconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(c & 31)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSRAcond(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRAcond x _ (FlagConstant [fc]))
	// cond: fc.uge()
	// result: (SRAconst x [31])
	for {
		x := v_0
		if v_2.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_2.AuxInt)
		if !(fc.uge()) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v.AddArg(x)
		return true
	}
	// match: (SRAcond x y (FlagConstant [fc]))
	// cond: fc.ult()
	// result: (SRA x y)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpThumbFlagConstant {
			break
		}
		fc := auxIntToFlagConstant(v_2.AuxInt)
		if !(fc.ult()) {
			break
		}
		v.reset(OpThumbSRA)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSRAconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRAconst [c] (MOVWconst [d]))
	// result: (MOVWconst [d>>uint64(c)])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(d >> uint64(c))
		return true
	}
	// match: (SRAconst (SLLconst x [c]) [d])
	// cond: uint64(d)>=uint64(c) && uint64(d)<=31
	// result: (BFX [(d-c)|(32-d)<<8] x)
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(uint64(d) >= uint64(c) && uint64(d) <= 31) {
			break
		}
		v.reset(OpThumbBFX)
		v.AuxInt = int32ToAuxInt((d - c) | (32-d)<<8)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRL x (MOVWconst [c]))
	// result: (SRLconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSRLconst)
		v.AuxInt = int32ToAuxInt(c & 31)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSRLconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRLconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int32(uint32(d)>>uint64(c))])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(d) >> uint64(c)))
		return true
	}
	// match: (SRLconst (SLLconst x [c]) [d])
	// cond: uint64(d)>=uint64(c) && uint64(d)<=31
	// result: (BFXU [(d-c)|(32-d)<<8] x)
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(uint64(d) >= uint64(c) && uint64(d) <= 31) {
			break
		}
		v.reset(OpThumbBFXU)
		v.AuxInt = int32ToAuxInt((d - c) | (32-d)<<8)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUB (MOVWconst [c]) x)
	// result: (RSBconst [c] x)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (SUB x (MOVWconst [c]))
	// result: (SUBconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSUBconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (SUB x (SLLconst [c] y))
	// result: (SUBshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbSUBshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUB (SLLconst [c] y) x)
	// result: (RSBshiftLL x y [c])
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbRSBshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUB x (SRLconst [c] y))
	// result: (SUBshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbSUBshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUB (SRLconst [c] y) x)
	// result: (RSBshiftRL x y [c])
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbRSBshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUB x (SRAconst [c] y))
	// result: (SUBshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbSUBshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUB (SRAconst [c] y) x)
	// result: (RSBshiftRA x y [c])
	for {
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbRSBshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUB x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (SUB a (MUL x y))
	// result: (MULS x y a)
	for {
		a := v_0
		if v_1.Op != OpThumbMUL {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		v.reset(OpThumbMULS)
		v.AddArg3(x, y, a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBD a (MULD x y))
	// cond: a.Uses == 1
	// result: (MULSD a x y)
	for {
		a := v_0
		if v_1.Op != OpThumbMULD {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULSD)
		v.AddArg3(a, x, y)
		return true
	}
	// match: (SUBD a (NMULD x y))
	// cond: a.Uses == 1
	// result: (MULAD a x y)
	for {
		a := v_0
		if v_1.Op != OpThumbNMULD {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULAD)
		v.AddArg3(a, x, y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBF(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBF a (MULF x y))
	// cond: a.Uses == 1
	// result: (MULSF a x y)
	for {
		a := v_0
		if v_1.Op != OpThumbMULF {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULSF)
		v.AddArg3(a, x, y)
		return true
	}
	// match: (SUBF a (NMULF x y))
	// cond: a.Uses == 1
	// result: (MULAF a x y)
	for {
		a := v_0
		if v_1.Op != OpThumbNMULF {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULAF)
		v.AddArg3(a, x, y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBS x (MOVWconst [c]))
	// result: (SUBSconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSUBSconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (SUBS x (SLLconst [c] y))
	// result: (SUBSshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbSUBSshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUBS (SLLconst [c] y) x)
	// result: (RSBSshiftLL x y [c])
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbRSBSshiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUBS x (SRLconst [c] y))
	// result: (SUBSshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbSUBSshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUBS (SRLconst [c] y) x)
	// result: (RSBSshiftRL x y [c])
	for {
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbRSBSshiftRL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUBS x (SRAconst [c] y))
	// result: (SUBSshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		y := v_1.Args[0]
		v.reset(OpThumbSUBSshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUBS (SRAconst [c] y) x)
	// result: (RSBSshiftRA x y [c])
	for {
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_0.Args[0]
		x := v_1
		v.reset(OpThumbRSBSshiftRA)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBSshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBSshiftLL (MOVWconst [c]) x [d])
	// result: (RSBSconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbRSBSconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftLL x (MOVWconst [c]) [d])
	// result: (SUBSconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSUBSconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBSshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBSshiftRA (MOVWconst [c]) x [d])
	// result: (RSBSconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbRSBSconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftRA x (MOVWconst [c]) [d])
	// result: (SUBSconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSUBSconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBSshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBSshiftRL (MOVWconst [c]) x [d])
	// result: (RSBSconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbRSBSconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftRL x (MOVWconst [c]) [d])
	// result: (SUBSconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSUBSconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SUBconst [off1] (MOVWaddr [off2] {sym} ptr))
	// result: (MOVWaddr [off2-off1] {sym} ptr)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVWaddr)
		v.AuxInt = int32ToAuxInt(off2 - off1)
		v.Aux = symToAux(sym)
		v.AddArg(ptr)
		return true
	}
	// match: (SUBconst [0] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (SUBconst [c] (MOVWconst [d]))
	// result: (MOVWconst [d-c])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(d - c)
		return true
	}
	// match: (SUBconst [c] (SUBconst [d] x))
	// result: (ADDconst [-c-d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(-c - d)
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (ADDconst [d] x))
	// result: (ADDconst [-c+d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = int32ToAuxInt(-c + d)
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (RSBconst [d] x))
	// result: (RSBconst [-c+d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbRSBconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(-c + d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBshiftLL (MOVWconst [c]) x [d])
	// result: (RSBconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftLL x (MOVWconst [c]) [d])
	// result: (SUBconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSUBconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (SUBshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBshiftRA (MOVWconst [c]) x [d])
	// result: (RSBconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftRA x (MOVWconst [c]) [d])
	// result: (SUBconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSUBconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (SUBshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBshiftRL (MOVWconst [c]) x [d])
	// result: (RSBconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbRSBconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftRL x (MOVWconst [c]) [d])
	// result: (SUBconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbSUBconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (SUBshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce16(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (StoreOnce16 [off1] (ADDconst [off2] ptr) val mem)
	// result: (StoreOnce16 [off1+off2] ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce16)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce16 [off1] (SUBconst [off2] ptr) val mem)
	// result: (StoreOnce16 [off1-off2] ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce16)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce16 [0] (ADD ptr idx) val mem)
	// result: (StoreOnce16idx ptr idx val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce16idx)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (StoreOnce16 [0] (ADDshiftLL ptr idx [c]) val mem)
	// cond: c <= 3
	// result: (StoreOnce16shiftLL ptr idx [c] val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce16shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce16idx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (StoreOnce16idx ptr (MOVWconst [c]) val mem)
	// result: (StoreOnce16 [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		val := v_2
		mem := v_3
		v.reset(OpThumbStoreOnce16)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce16idx (MOVWconst [c]) ptr val mem)
	// result: (StoreOnce16 [c] ptr val mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpThumbStoreOnce16)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce16idx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (StoreOnce16shiftLL ptr idx [c] val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce16shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (StoreOnce16idx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (StoreOnce16shiftLL ptr idx [c] val mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce16shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (StoreOnce32 [off1] (ADDconst [off2] ptr) val mem)
	// result: (StoreOnce32 [off1+off2] ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce32)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce32 [off1] (SUBconst [off2] ptr) val mem)
	// result: (StoreOnce32 [off1-off2] ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce32)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce32 [0] (ADD ptr idx) val mem)
	// result: (StoreOnce32idx ptr idx val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce32idx)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (StoreOnce32 [0] (ADDshiftLL ptr idx [c]) val mem)
	// cond: c <= 3
	// result: (StoreOnce32shiftLL ptr idx [c] val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce32shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce32idx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (StoreOnce32idx ptr (MOVWconst [c]) val mem)
	// result: (StoreOnce32 [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		val := v_2
		mem := v_3
		v.reset(OpThumbStoreOnce32)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce32idx (MOVWconst [c]) ptr val mem)
	// result: (StoreOnce32 [c] ptr val mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpThumbStoreOnce32)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce32idx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (StoreOnce32shiftLL ptr idx [c] val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce32shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (StoreOnce32idx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (StoreOnce32shiftLL ptr idx [c] val mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce32shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (StoreOnce8 [off1] (ADDconst [off2] ptr) val mem)
	// result: (StoreOnce8 [off1+off2] ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce8)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce8 [off1] (SUBconst [off2] ptr) val mem)
	// result: (StoreOnce8 [off1-off2] ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce8)
		v.AuxInt = int32ToAuxInt(off1 - off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce8 [0] (ADD ptr idx) val mem)
	// result: (StoreOnce8idx ptr idx val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpThumbStoreOnce8idx)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (StoreOnce8 [0] (ADDshiftLL ptr idx [c]) val mem)
	// cond: c <= 3
	// result: (StoreOnce8shiftLL ptr idx [c] val mem)
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce8shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce8idx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (StoreOnce8idx ptr (MOVWconst [c]) val mem)
	// result: (StoreOnce8 [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		val := v_2
		mem := v_3
		v.reset(OpThumbStoreOnce8)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce8idx (MOVWconst [c]) ptr val mem)
	// result: (StoreOnce8 [c] ptr val mem)
	for {
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpThumbStoreOnce8)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (StoreOnce8idx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (StoreOnce8shiftLL ptr idx [c] val mem)
	for {
		ptr := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce8shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	// match: (StoreOnce8idx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (StoreOnce8shiftLL ptr idx [c] val mem)
	for {
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		idx := v_0.Args[0]
		ptr := v_1
		val := v_2
		mem := v_3
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce8shiftLL)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg4(ptr, idx, val, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTEQ(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (TEQ x (MOVWconst [c]))
	// result: (TEQconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpThumbTEQconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (TEQ x (SLLconst [c] y))
	// result: (TEQshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSLLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbTEQshiftLL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (TEQ x (SRLconst [c] y))
	// result: (TEQshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbTEQshiftRL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (TEQ x (SRAconst [c] y))
	// result: (TEQshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRAconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbTEQshiftRA)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbTEQconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (TEQconst (MOVWconst [x]) [y])
	// result: (FlagConstant [logicFlags32(x^y)])
	for {
		y := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbFlagConstant)
		v.AuxInt = flagConstantToAuxInt(logicFlags32(x ^ y))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTEQshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TEQshiftLL (MOVWconst [c]) x [d])
	// result: (TEQconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbTEQconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftLL x (MOVWconst [c]) [d])
	// result: (TEQconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbTEQconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTEQshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TEQshiftRA (MOVWconst [c]) x [d])
	// result: (TEQconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbTEQconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftRA x (MOVWconst [c]) [d])
	// result: (TEQconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbTEQconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTEQshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TEQshiftRL (MOVWconst [c]) x [d])
	// result: (TEQconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbTEQconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftRL x (MOVWconst [c]) [d])
	// result: (TEQconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbTEQconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTST(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (TST x (MOVWconst [c]))
	// result: (TSTconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpThumbTSTconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (TST x (SLLconst [c] y))
	// result: (TSTshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSLLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbTSTshiftLL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (TST x (SRLconst [c] y))
	// result: (TSTshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbTSTshiftRL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (TST x (SRAconst [c] y))
	// result: (TSTshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRAconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbTSTshiftRA)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueThumb_OpThumbTSTconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (TSTconst (MOVWconst [x]) [y])
	// result: (FlagConstant [logicFlags32(x&y)])
	for {
		y := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbFlagConstant)
		v.AuxInt = flagConstantToAuxInt(logicFlags32(x & y))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTSTshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TSTshiftLL (MOVWconst [c]) x [d])
	// result: (TSTconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbTSTconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftLL x (MOVWconst [c]) [d])
	// result: (TSTconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbTSTconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTSTshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TSTshiftRA (MOVWconst [c]) x [d])
	// result: (TSTconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbTSTconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftRA x (MOVWconst [c]) [d])
	// result: (TSTconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbTSTconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTSTshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TSTshiftRL (MOVWconst [c]) x [d])
	// result: (TSTconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbTSTconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftRL x (MOVWconst [c]) [d])
	// result: (TSTconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbTSTconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XOR x (MOVWconst [c]))
	// result: (XORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpThumbXORconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XOR x (SLLconst [c] y))
	// result: (XORshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSLLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbXORshiftLL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (XOR x (SRLconst [c] y))
	// result: (XORshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRLconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbXORshiftRL)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (XOR x (SRAconst [c] y))
	// result: (XORshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRAconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbXORshiftRA)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (XOR x (SRRconst [c] y))
	// result: (XORshiftRR x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpThumbSRRconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			y := v_1.Args[0]
			v.reset(OpThumbXORshiftRR)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (XOR x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (XORconst [0] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (XORconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c^d])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(c ^ d)
		return true
	}
	// match: (XORconst [c] (XORconst [d] x))
	// result: (XORconst [c^d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbXORconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(c ^ d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXORshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (XORshiftLL (MOVWconst [c]) x [d])
	// result: (XORconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftLL x (MOVWconst [c]) [d])
	// result: (XORconst x [c<<uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(c << uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL [c] (SRLconst x [32-c]) x)
	// result: (SRRconst [32-c] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSRLconst || auxIntToInt32(v_0.AuxInt) != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = int32ToAuxInt(32 - c)
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL <typ.UInt16> [8] (BFXU <typ.UInt16> [int32(armBFAuxInt(8, 8))] x) x)
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || auxIntToInt32(v.AuxInt) != 8 || v_0.Op != OpThumbBFXU || v_0.Type != typ.UInt16 || auxIntToInt32(v_0.AuxInt) != int32(armBFAuxInt(8, 8)) {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbREV16)
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL <typ.UInt16> [8] (SRLconst <typ.UInt16> [24] (SLLconst [16] x)) x)
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || auxIntToInt32(v.AuxInt) != 8 || v_0.Op != OpThumbSRLconst || v_0.Type != typ.UInt16 || auxIntToInt32(v_0.AuxInt) != 24 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbSLLconst || auxIntToInt32(v_0_0.AuxInt) != 16 {
			break
		}
		x := v_0_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbREV16)
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXORshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (XORshiftRA (MOVWconst [c]) x [d])
	// result: (XORconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRA x (MOVWconst [c]) [d])
	// result: (XORconst x [c>>uint64(d)])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(c >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (XORshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXORshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (XORshiftRL (MOVWconst [c]) x [d])
	// result: (XORconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRL x (MOVWconst [c]) [d])
	// result: (XORconst x [int32(uint32(c)>>uint64(d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (XORshiftRL [c] (SLLconst x [32-c]) x)
	// result: (SRRconst [ c] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbSLLconst || auxIntToInt32(v_0.AuxInt) != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (XORshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXORshiftRR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (XORshiftRR (MOVWconst [c]) x [d])
	// result: (XORconst [c] (SRRconst <x.Type> x [d]))
	for {
		d := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpThumbSRRconst, x.Type)
		v0.AuxInt = int32ToAuxInt(d)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRR x (MOVWconst [c]) [d])
	// result: (XORconst x [int32(uint32(c)>>uint64(d)|uint32(c)<<uint64(32-d))])
	for {
		d := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpThumbXORconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c)>>uint64(d) | uint32(c)<<uint64(32-d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_1
		v.copyOf(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// result: (MOVBstore ptr (MOVWconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpThumbMOVBstore)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVWconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// result: (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(0)
		v1.AddArg3(ptr, v0, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVWconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpThumbMOVWstore)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore [2] ptr (MOVWconst [0]) (MOVHstore [0] ptr (MOVWconst [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVHstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(0)
		v1.AddArg3(ptr, v0, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [4] ptr mem)
	// result: (MOVBstore [3] ptr (MOVWconst [0]) (MOVBstore [2] ptr (MOVWconst [0]) (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(3)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(2)
		v2 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(1)
		v3 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(0)
		v3.AddArg3(ptr, v0, mem)
		v2.AddArg3(ptr, v0, v3)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [3] ptr mem)
	// result: (MOVBstore [2] ptr (MOVWconst [0]) (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 3 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpThumbMOVBstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(1)
		v2 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(0)
		v2.AddArg3(ptr, v0, mem)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: s%4 == 0 && s > 4 && s <= 512 && t.Alignment()%4 == 0 && !config.noDuffDevice
	// result: (DUFFZERO [4 * (128 - s/4)] ptr (MOVWconst [0]) mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(s%4 == 0 && s > 4 && s <= 512 && t.Alignment()%4 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpThumbDUFFZERO)
		v.AuxInt = int64ToAuxInt(4 * (128 - s/4))
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: (s > 512 || config.noDuffDevice) || t.Alignment()%4 != 0
	// result: (LoweredZero [t.Alignment()] ptr (ADDconst <ptr.Type> ptr [int32(s-moveSize(t.Alignment(), config))]) (MOVWconst [0]) mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !((s > 512 || config.noDuffDevice) || t.Alignment()%4 != 0) {
			break
		}
		v.reset(OpThumbLoweredZero)
		v.AuxInt = int64ToAuxInt(t.Alignment())
		v0 := b.NewValue0(v.Pos, OpThumbADDconst, ptr.Type)
		v0.AuxInt = int32ToAuxInt(int32(s - moveSize(t.Alignment(), config)))
		v0.AddArg(ptr)
		v1 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(0)
		v.AddArg4(ptr, v0, v1, mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpZeromask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Zeromask x)
	// result: (SRAconst (RSBshiftRL <typ.Int32> x x [1]) [31])
	for {
		x := v_0
		v.reset(OpThumbSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftRL, typ.Int32)
		v0.AuxInt = int32ToAuxInt(1)
		v0.AddArg2(x, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteBlockThumb(b *Block) bool {
	switch b.Kind {
	case BlockThumbEQ:
		// match: (EQ (FlagConstant [fc]) yes no)
		// cond: fc.eq()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.eq()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (EQ (FlagConstant [fc]) yes no)
		// cond: !fc.eq()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.eq()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (EQ (InvertFlags cmp) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbEQ, cmp)
			return true
		}
		// match: (EQ (CMP x (RSBconst [0] y)))
		// result: (EQ (CMN x y))
		for b.Controls[0].Op == OpThumbCMP {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
				break
			}
			y := v_0_1.Args[0]
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMN x (RSBconst [0] y)))
		// result: (EQ (CMP x y))
		for b.Controls[0].Op == OpThumbCMN {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
					continue
				}
				y := v_0_1.Args[0]
				v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbEQ, v0)
				return true
			}
			break
		}
		// match: (EQ (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMP x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMN x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbEQ, v0)
				return true
			}
			break
		}
		// match: (EQ (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TST x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTST, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbEQ, v0)
				return true
			}
			break
		}
		// match: (EQ (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQ x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTEQ, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbEQ, v0)
				return true
			}
			break
		}
		// match: (EQ (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbEQ, v0)
			return true
		}
	case BlockThumbGE:
		// match: (GE (FlagConstant [fc]) yes no)
		// cond: fc.ge()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.ge()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (GE (FlagConstant [fc]) yes no)
		// cond: !fc.ge()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.ge()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GE (InvertFlags cmp) yes no)
		// result: (LE cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbLE, cmp)
			return true
		}
		// match: (GE (CMP x (RSBconst [0] y)))
		// result: (GE (CMN x y))
		for b.Controls[0].Op == OpThumbCMP {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
				break
			}
			y := v_0_1.Args[0]
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGE, v0)
			return true
		}
		// match: (GE (CMN x (RSBconst [0] y)))
		// result: (GE (CMP x y))
		for b.Controls[0].Op == OpThumbCMN {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
					continue
				}
				y := v_0_1.Args[0]
				v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbGE, v0)
				return true
			}
			break
		}
		// match: (GE (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMP x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMN x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbGEnoov, v0)
				return true
			}
			break
		}
		// match: (GE (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (TST x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTST, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbGEnoov, v0)
				return true
			}
			break
		}
		// match: (GE (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (TEQ x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTEQ, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbGEnoov, v0)
				return true
			}
			break
		}
		// match: (GE (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GEnoov (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGEnoov, v0)
			return true
		}
	case BlockThumbGEnoov:
		// match: (GEnoov (FlagConstant [fc]) yes no)
		// cond: fc.geNoov()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.geNoov()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (GEnoov (FlagConstant [fc]) yes no)
		// cond: !fc.geNoov()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.geNoov()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GEnoov (InvertFlags cmp) yes no)
		// result: (LEnoov cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbLEnoov, cmp)
			return true
		}
	case BlockThumbGT:
		// match: (GT (FlagConstant [fc]) yes no)
		// cond: fc.gt()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.gt()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (GT (FlagConstant [fc]) yes no)
		// cond: !fc.gt()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.gt()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GT (InvertFlags cmp) yes no)
		// result: (LT cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbLT, cmp)
			return true
		}
		// match: (GT (CMP x (RSBconst [0] y)))
		// result: (GT (CMN x y))
		for b.Controls[0].Op == OpThumbCMP {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
				break
			}
			y := v_0_1.Args[0]
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGT, v0)
			return true
		}
		// match: (GT (CMN x (RSBconst [0] y)))
		// result: (GT (CMP x y))
		for b.Controls[0].Op == OpThumbCMN {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
					continue
				}
				y := v_0_1.Args[0]
				v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbGT, v0)
				return true
			}
			break
		}
		// match: (GT (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMP x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMN x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbGTnoov, v0)
				return true
			}
			break
		}
		// match: (GT (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (TST x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTST, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbGTnoov, v0)
				return true
			}
			break
		}
		// match: (GT (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (TEQ x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTEQ, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbGTnoov, v0)
				return true
			}
			break
		}
		// match: (GT (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GTnoov (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbGTnoov, v0)
			return true
		}
	case BlockThumbGTnoov:
		// match: (GTnoov (FlagConstant [fc]) yes no)
		// cond: fc.gtNoov()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.gtNoov()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (GTnoov (FlagConstant [fc]) yes no)
		// cond: !fc.gtNoov()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.gtNoov()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GTnoov (InvertFlags cmp) yes no)
		// result: (LTnoov cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbLTnoov, cmp)
			return true
		}
	case BlockIf:
		// match: (If (Equal cc) yes no)
		// result: (EQ cc yes no)
		for b.Controls[0].Op == OpThumbEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(BlockThumbEQ, cc)
			return true
		}
		// match: (If (NotEqual cc) yes no)
		// result: (NE cc yes no)
		for b.Controls[0].Op == OpThumbNotEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(BlockThumbNE, cc)
			return true
		}
		// match: (If (LessThan cc) yes no)
		// result: (LT cc yes no)
		for b.Controls[0].Op == OpThumbLessThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(BlockThumbLT, cc)
			return true
		}
		// match: (If (LessThanU cc) yes no)
		// result: (ULT cc yes no)
		for b.Controls[0].Op == OpThumbLessThanU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(BlockThumbULT, cc)
			return true
		}
		// match: (If (LessEqual cc) yes no)
		// result: (LE cc yes no)
		for b.Controls[0].Op == OpThumbLessEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(BlockThumbLE, cc)
			return true
		}
		// match: (If (LessEqualU cc) yes no)
		// result: (ULE cc yes no)
		for b.Controls[0].Op == OpThumbLessEqualU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(BlockThumbULE, cc)
			return true
		}
		// match: (If (GreaterThan cc) yes no)
		// result: (GT cc yes no)
		for b.Controls[0].Op == OpThumbGreaterThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(BlockThumbGT, cc)
			return true
		}
		// match: (If (GreaterThanU cc) yes no)
		// result: (UGT cc yes no)
		for b.Controls[0].Op == OpThumbGreaterThanU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(BlockThumbUGT, cc)
			return true
		}
		// match: (If (GreaterEqual cc) yes no)
		// result: (GE cc yes no)
		for b.Controls[0].Op == OpThumbGreaterEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(BlockThumbGE, cc)
			return true
		}
		// match: (If (GreaterEqualU cc) yes no)
		// result: (UGE cc yes no)
		for b.Controls[0].Op == OpThumbGreaterEqualU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(BlockThumbUGE, cc)
			return true
		}
		// match: (If cond yes no)
		// result: (NE (CMPconst [0] cond) yes no)
		for {
			cond := b.Controls[0]
			v0 := b.NewValue0(cond.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(0)
			v0.AddArg(cond)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
	case BlockThumbLE:
		// match: (LE (FlagConstant [fc]) yes no)
		// cond: fc.le()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.le()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (LE (FlagConstant [fc]) yes no)
		// cond: !fc.le()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.le()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LE (InvertFlags cmp) yes no)
		// result: (GE cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbGE, cmp)
			return true
		}
		// match: (LE (CMP x (RSBconst [0] y)))
		// result: (LE (CMN x y))
		for b.Controls[0].Op == OpThumbCMP {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
				break
			}
			y := v_0_1.Args[0]
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLE, v0)
			return true
		}
		// match: (LE (CMN x (RSBconst [0] y)))
		// result: (LE (CMP x y))
		for b.Controls[0].Op == OpThumbCMN {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
					continue
				}
				y := v_0_1.Args[0]
				v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbLE, v0)
				return true
			}
			break
		}
		// match: (LE (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMP x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMN x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbLEnoov, v0)
				return true
			}
			break
		}
		// match: (LE (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (TST x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTST, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbLEnoov, v0)
				return true
			}
			break
		}
		// match: (LE (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (TEQ x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTEQ, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbLEnoov, v0)
				return true
			}
			break
		}
		// match: (LE (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LEnoov (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLEnoov, v0)
			return true
		}
	case BlockThumbLEnoov:
		// match: (LEnoov (FlagConstant [fc]) yes no)
		// cond: fc.leNoov()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.leNoov()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (LEnoov (FlagConstant [fc]) yes no)
		// cond: !fc.leNoov()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.leNoov()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LEnoov (InvertFlags cmp) yes no)
		// result: (GEnoov cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbGEnoov, cmp)
			return true
		}
	case BlockThumbLT:
		// match: (LT (FlagConstant [fc]) yes no)
		// cond: fc.lt()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.lt()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (LT (FlagConstant [fc]) yes no)
		// cond: !fc.lt()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.lt()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LT (InvertFlags cmp) yes no)
		// result: (GT cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbGT, cmp)
			return true
		}
		// match: (LT (CMP x (RSBconst [0] y)))
		// result: (LT (CMN x y))
		for b.Controls[0].Op == OpThumbCMP {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
				break
			}
			y := v_0_1.Args[0]
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLT, v0)
			return true
		}
		// match: (LT (CMN x (RSBconst [0] y)))
		// result: (LT (CMP x y))
		for b.Controls[0].Op == OpThumbCMN {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
					continue
				}
				y := v_0_1.Args[0]
				v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbLT, v0)
				return true
			}
			break
		}
		// match: (LT (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMP x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMN x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbLTnoov, v0)
				return true
			}
			break
		}
		// match: (LT (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (TST x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTST, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbLTnoov, v0)
				return true
			}
			break
		}
		// match: (LT (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (TEQ x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTEQ, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbLTnoov, v0)
				return true
			}
			break
		}
		// match: (LT (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LTnoov (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbLTnoov, v0)
			return true
		}
	case BlockThumbLTnoov:
		// match: (LTnoov (FlagConstant [fc]) yes no)
		// cond: fc.ltNoov()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.ltNoov()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (LTnoov (FlagConstant [fc]) yes no)
		// cond: !fc.ltNoov()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.ltNoov()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LTnoov (InvertFlags cmp) yes no)
		// result: (GTnoov cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbGTnoov, cmp)
			return true
		}
	case BlockThumbNE:
		// match: (NE (CMPconst [0] (Equal cc)) yes no)
		// result: (EQ cc yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpThumbEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.resetWithControl(BlockThumbEQ, cc)
			return true
		}
		// match: (NE (CMPconst [0] (NotEqual cc)) yes no)
		// result: (NE cc yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpThumbNotEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.resetWithControl(BlockThumbNE, cc)
			return true
		}
		// match: (NE (CMPconst [0] (LessThan cc)) yes no)
		// result: (LT cc yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpThumbLessThan {
				break
			}
			cc := v_0_0.Args[0]
			b.resetWithControl(BlockThumbLT, cc)
			return true
		}
		// match: (NE (CMPconst [0] (LessThanU cc)) yes no)
		// result: (ULT cc yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpThumbLessThanU {
				break
			}
			cc := v_0_0.Args[0]
			b.resetWithControl(BlockThumbULT, cc)
			return true
		}
		// match: (NE (CMPconst [0] (LessEqual cc)) yes no)
		// result: (LE cc yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpThumbLessEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.resetWithControl(BlockThumbLE, cc)
			return true
		}
		// match: (NE (CMPconst [0] (LessEqualU cc)) yes no)
		// result: (ULE cc yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpThumbLessEqualU {
				break
			}
			cc := v_0_0.Args[0]
			b.resetWithControl(BlockThumbULE, cc)
			return true
		}
		// match: (NE (CMPconst [0] (GreaterThan cc)) yes no)
		// result: (GT cc yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpThumbGreaterThan {
				break
			}
			cc := v_0_0.Args[0]
			b.resetWithControl(BlockThumbGT, cc)
			return true
		}
		// match: (NE (CMPconst [0] (GreaterThanU cc)) yes no)
		// result: (UGT cc yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpThumbGreaterThanU {
				break
			}
			cc := v_0_0.Args[0]
			b.resetWithControl(BlockThumbUGT, cc)
			return true
		}
		// match: (NE (CMPconst [0] (GreaterEqual cc)) yes no)
		// result: (GE cc yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpThumbGreaterEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.resetWithControl(BlockThumbGE, cc)
			return true
		}
		// match: (NE (CMPconst [0] (GreaterEqualU cc)) yes no)
		// result: (UGE cc yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpThumbGreaterEqualU {
				break
			}
			cc := v_0_0.Args[0]
			b.resetWithControl(BlockThumbUGE, cc)
			return true
		}
		// match: (NE (FlagConstant [fc]) yes no)
		// cond: fc.ne()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.ne()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (NE (FlagConstant [fc]) yes no)
		// cond: !fc.ne()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.ne()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (NE (InvertFlags cmp) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbNE, cmp)
			return true
		}
		// match: (NE (CMP x (RSBconst [0] y)))
		// result: (NE (CMN x y))
		for b.Controls[0].Op == OpThumbCMP {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
				break
			}
			y := v_0_1.Args[0]
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMN x (RSBconst [0] y)))
		// result: (NE (CMP x y))
		for b.Controls[0].Op == OpThumbCMN {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpThumbRSBconst || auxIntToInt32(v_0_1.AuxInt) != 0 {
					continue
				}
				y := v_0_1.Args[0]
				v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbNE, v0)
				return true
			}
			break
		}
		// match: (NE (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMP x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMP, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMN x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbNE, v0)
				return true
			}
			break
		}
		// match: (NE (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMN, types.TypeFlags)
			v1 := b.NewValue0(v_0.Pos, OpThumbMUL, x.Type)
			v1.AddArg2(x, y)
			v0.AddArg2(a, v1)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (TST x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTST, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbNE, v0)
				return true
			}
			break
		}
		// match: (NE (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQ x y) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				v0 := b.NewValue0(v_0.Pos, OpThumbTEQ, types.TypeFlags)
				v0.AddArg2(x, y)
				b.resetWithControl(BlockThumbNE, v0)
				return true
			}
			break
		}
		// match: (NE (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg(x)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpThumbCMPconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := auxIntToInt32(l.AuxInt)
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			v0 := b.NewValue0(v_0.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = int32ToAuxInt(c)
			v0.AddArg2(x, y)
			b.resetWithControl(BlockThumbNE, v0)
			return true
		}
	case BlockThumbUGE:
		// match: (UGE (FlagConstant [fc]) yes no)
		// cond: fc.uge()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.uge()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (UGE (FlagConstant [fc]) yes no)
		// cond: !fc.uge()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.uge()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGE (InvertFlags cmp) yes no)
		// result: (ULE cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbULE, cmp)
			return true
		}
	case BlockThumbUGT:
		// match: (UGT (FlagConstant [fc]) yes no)
		// cond: fc.ugt()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.ugt()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (UGT (FlagConstant [fc]) yes no)
		// cond: !fc.ugt()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.ugt()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGT (InvertFlags cmp) yes no)
		// result: (ULT cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbULT, cmp)
			return true
		}
	case BlockThumbULE:
		// match: (ULE (FlagConstant [fc]) yes no)
		// cond: fc.ule()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.ule()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULE (FlagConstant [fc]) yes no)
		// cond: !fc.ule()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.ule()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (ULE (InvertFlags cmp) yes no)
		// result: (UGE cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbUGE, cmp)
			return true
		}
	case BlockThumbULT:
		// match: (ULT (FlagConstant [fc]) yes no)
		// cond: fc.ult()
		// result: (First yes no)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(fc.ult()) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULT (FlagConstant [fc]) yes no)
		// cond: !fc.ult()
		// result: (First no yes)
		for b.Controls[0].Op == OpThumbFlagConstant {
			v_0 := b.Controls[0]
			fc := auxIntToFlagConstant(v_0.AuxInt)
			if !(!fc.ult()) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (ULT (InvertFlags cmp) yes no)
		// result: (UGT cmp yes no)
		for b.Controls[0].Op == OpThumbInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockThumbUGT, cmp)
			return true
		}
	}
	return false
}
