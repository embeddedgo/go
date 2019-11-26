// Code generated from gen/Thumb.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "fmt"
import "math"
import "cmd/internal/obj"
import "cmd/internal/objabi"
import "cmd/compile/internal/types"

var _ = fmt.Println   // in case not otherwise used
var _ = math.MinInt8  // in case not otherwise used
var _ = obj.ANOP      // in case not otherwise used
var _ = objabi.GOROOT // in case not otherwise used
var _ = types.TypeMem // in case not otherwise used

func rewriteValueThumb(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValueThumb_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValueThumb_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValueThumb_OpAdd32F_0(v)
	case OpAdd32carry:
		return rewriteValueThumb_OpAdd32carry_0(v)
	case OpAdd32withcarry:
		return rewriteValueThumb_OpAdd32withcarry_0(v)
	case OpAdd64F:
		return rewriteValueThumb_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValueThumb_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValueThumb_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValueThumb_OpAddr_0(v)
	case OpAnd16:
		return rewriteValueThumb_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValueThumb_OpAnd32_0(v)
	case OpAnd8:
		return rewriteValueThumb_OpAnd8_0(v)
	case OpAndB:
		return rewriteValueThumb_OpAndB_0(v)
	case OpAvg32u:
		return rewriteValueThumb_OpAvg32u_0(v)
	case OpBitLen32:
		return rewriteValueThumb_OpBitLen32_0(v)
	case OpBswap32:
		return rewriteValueThumb_OpBswap32_0(v)
	case OpClosureCall:
		return rewriteValueThumb_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValueThumb_OpCom16_0(v)
	case OpCom32:
		return rewriteValueThumb_OpCom32_0(v)
	case OpCom8:
		return rewriteValueThumb_OpCom8_0(v)
	case OpConst16:
		return rewriteValueThumb_OpConst16_0(v)
	case OpConst32:
		return rewriteValueThumb_OpConst32_0(v)
	case OpConst32F:
		return rewriteValueThumb_OpConst32F_0(v)
	case OpConst64F:
		return rewriteValueThumb_OpConst64F_0(v)
	case OpConst8:
		return rewriteValueThumb_OpConst8_0(v)
	case OpConstBool:
		return rewriteValueThumb_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValueThumb_OpConstNil_0(v)
	case OpCtz16:
		return rewriteValueThumb_OpCtz16_0(v)
	case OpCtz16NonZero:
		return rewriteValueThumb_OpCtz16NonZero_0(v)
	case OpCtz32:
		return rewriteValueThumb_OpCtz32_0(v)
	case OpCtz32NonZero:
		return rewriteValueThumb_OpCtz32NonZero_0(v)
	case OpCtz8:
		return rewriteValueThumb_OpCtz8_0(v)
	case OpCtz8NonZero:
		return rewriteValueThumb_OpCtz8NonZero_0(v)
	case OpCvt32Fto32:
		return rewriteValueThumb_OpCvt32Fto32_0(v)
	case OpCvt32Fto32U:
		return rewriteValueThumb_OpCvt32Fto32U_0(v)
	case OpCvt32Fto64F:
		return rewriteValueThumb_OpCvt32Fto64F_0(v)
	case OpCvt32Uto32F:
		return rewriteValueThumb_OpCvt32Uto32F_0(v)
	case OpCvt32Uto64F:
		return rewriteValueThumb_OpCvt32Uto64F_0(v)
	case OpCvt32to32F:
		return rewriteValueThumb_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValueThumb_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValueThumb_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValueThumb_OpCvt64Fto32F_0(v)
	case OpCvt64Fto32U:
		return rewriteValueThumb_OpCvt64Fto32U_0(v)
	case OpDiv16:
		return rewriteValueThumb_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValueThumb_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValueThumb_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValueThumb_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValueThumb_OpDiv32u_0(v)
	case OpDiv64F:
		return rewriteValueThumb_OpDiv64F_0(v)
	case OpDiv8:
		return rewriteValueThumb_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValueThumb_OpDiv8u_0(v)
	case OpEq16:
		return rewriteValueThumb_OpEq16_0(v)
	case OpEq32:
		return rewriteValueThumb_OpEq32_0(v)
	case OpEq32F:
		return rewriteValueThumb_OpEq32F_0(v)
	case OpEq64F:
		return rewriteValueThumb_OpEq64F_0(v)
	case OpEq8:
		return rewriteValueThumb_OpEq8_0(v)
	case OpEqB:
		return rewriteValueThumb_OpEqB_0(v)
	case OpEqPtr:
		return rewriteValueThumb_OpEqPtr_0(v)
	case OpGeq16:
		return rewriteValueThumb_OpGeq16_0(v)
	case OpGeq16U:
		return rewriteValueThumb_OpGeq16U_0(v)
	case OpGeq32:
		return rewriteValueThumb_OpGeq32_0(v)
	case OpGeq32F:
		return rewriteValueThumb_OpGeq32F_0(v)
	case OpGeq32U:
		return rewriteValueThumb_OpGeq32U_0(v)
	case OpGeq64F:
		return rewriteValueThumb_OpGeq64F_0(v)
	case OpGeq8:
		return rewriteValueThumb_OpGeq8_0(v)
	case OpGeq8U:
		return rewriteValueThumb_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValueThumb_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValueThumb_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValueThumb_OpGetClosurePtr_0(v)
	case OpGreater16:
		return rewriteValueThumb_OpGreater16_0(v)
	case OpGreater16U:
		return rewriteValueThumb_OpGreater16U_0(v)
	case OpGreater32:
		return rewriteValueThumb_OpGreater32_0(v)
	case OpGreater32F:
		return rewriteValueThumb_OpGreater32F_0(v)
	case OpGreater32U:
		return rewriteValueThumb_OpGreater32U_0(v)
	case OpGreater64F:
		return rewriteValueThumb_OpGreater64F_0(v)
	case OpGreater8:
		return rewriteValueThumb_OpGreater8_0(v)
	case OpGreater8U:
		return rewriteValueThumb_OpGreater8U_0(v)
	case OpHmul32:
		return rewriteValueThumb_OpHmul32_0(v)
	case OpHmul32u:
		return rewriteValueThumb_OpHmul32u_0(v)
	case OpInterCall:
		return rewriteValueThumb_OpInterCall_0(v)
	case OpIsInBounds:
		return rewriteValueThumb_OpIsInBounds_0(v)
	case OpIsNonNil:
		return rewriteValueThumb_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return rewriteValueThumb_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return rewriteValueThumb_OpLeq16_0(v)
	case OpLeq16U:
		return rewriteValueThumb_OpLeq16U_0(v)
	case OpLeq32:
		return rewriteValueThumb_OpLeq32_0(v)
	case OpLeq32F:
		return rewriteValueThumb_OpLeq32F_0(v)
	case OpLeq32U:
		return rewriteValueThumb_OpLeq32U_0(v)
	case OpLeq64F:
		return rewriteValueThumb_OpLeq64F_0(v)
	case OpLeq8:
		return rewriteValueThumb_OpLeq8_0(v)
	case OpLeq8U:
		return rewriteValueThumb_OpLeq8U_0(v)
	case OpLess16:
		return rewriteValueThumb_OpLess16_0(v)
	case OpLess16U:
		return rewriteValueThumb_OpLess16U_0(v)
	case OpLess32:
		return rewriteValueThumb_OpLess32_0(v)
	case OpLess32F:
		return rewriteValueThumb_OpLess32F_0(v)
	case OpLess32U:
		return rewriteValueThumb_OpLess32U_0(v)
	case OpLess64F:
		return rewriteValueThumb_OpLess64F_0(v)
	case OpLess8:
		return rewriteValueThumb_OpLess8_0(v)
	case OpLess8U:
		return rewriteValueThumb_OpLess8U_0(v)
	case OpLoad:
		return rewriteValueThumb_OpLoad_0(v)
	case OpLocalAddr:
		return rewriteValueThumb_OpLocalAddr_0(v)
	case OpLsh16x16:
		return rewriteValueThumb_OpLsh16x16_0(v)
	case OpLsh16x32:
		return rewriteValueThumb_OpLsh16x32_0(v)
	case OpLsh16x64:
		return rewriteValueThumb_OpLsh16x64_0(v)
	case OpLsh16x8:
		return rewriteValueThumb_OpLsh16x8_0(v)
	case OpLsh32x16:
		return rewriteValueThumb_OpLsh32x16_0(v)
	case OpLsh32x32:
		return rewriteValueThumb_OpLsh32x32_0(v)
	case OpLsh32x64:
		return rewriteValueThumb_OpLsh32x64_0(v)
	case OpLsh32x8:
		return rewriteValueThumb_OpLsh32x8_0(v)
	case OpLsh8x16:
		return rewriteValueThumb_OpLsh8x16_0(v)
	case OpLsh8x32:
		return rewriteValueThumb_OpLsh8x32_0(v)
	case OpLsh8x64:
		return rewriteValueThumb_OpLsh8x64_0(v)
	case OpLsh8x8:
		return rewriteValueThumb_OpLsh8x8_0(v)
	case OpMMIOLoad16:
		return rewriteValueThumb_OpMMIOLoad16_0(v)
	case OpMMIOLoad32:
		return rewriteValueThumb_OpMMIOLoad32_0(v)
	case OpMMIOLoad8:
		return rewriteValueThumb_OpMMIOLoad8_0(v)
	case OpMMIOMB:
		return rewriteValueThumb_OpMMIOMB_0(v)
	case OpMMIOStore16:
		return rewriteValueThumb_OpMMIOStore16_0(v)
	case OpMMIOStore32:
		return rewriteValueThumb_OpMMIOStore32_0(v)
	case OpMMIOStore8:
		return rewriteValueThumb_OpMMIOStore8_0(v)
	case OpMod16:
		return rewriteValueThumb_OpMod16_0(v)
	case OpMod16u:
		return rewriteValueThumb_OpMod16u_0(v)
	case OpMod32:
		return rewriteValueThumb_OpMod32_0(v)
	case OpMod32u:
		return rewriteValueThumb_OpMod32u_0(v)
	case OpMod8:
		return rewriteValueThumb_OpMod8_0(v)
	case OpMod8u:
		return rewriteValueThumb_OpMod8u_0(v)
	case OpMove:
		return rewriteValueThumb_OpMove_0(v)
	case OpMul16:
		return rewriteValueThumb_OpMul16_0(v)
	case OpMul32:
		return rewriteValueThumb_OpMul32_0(v)
	case OpMul32F:
		return rewriteValueThumb_OpMul32F_0(v)
	case OpMul32uhilo:
		return rewriteValueThumb_OpMul32uhilo_0(v)
	case OpMul64F:
		return rewriteValueThumb_OpMul64F_0(v)
	case OpMul8:
		return rewriteValueThumb_OpMul8_0(v)
	case OpNeg16:
		return rewriteValueThumb_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValueThumb_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValueThumb_OpNeg32F_0(v)
	case OpNeg64F:
		return rewriteValueThumb_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValueThumb_OpNeg8_0(v)
	case OpNeq16:
		return rewriteValueThumb_OpNeq16_0(v)
	case OpNeq32:
		return rewriteValueThumb_OpNeq32_0(v)
	case OpNeq32F:
		return rewriteValueThumb_OpNeq32F_0(v)
	case OpNeq64F:
		return rewriteValueThumb_OpNeq64F_0(v)
	case OpNeq8:
		return rewriteValueThumb_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValueThumb_OpNeqB_0(v)
	case OpNeqPtr:
		return rewriteValueThumb_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValueThumb_OpNilCheck_0(v)
	case OpNot:
		return rewriteValueThumb_OpNot_0(v)
	case OpOffPtr:
		return rewriteValueThumb_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValueThumb_OpOr16_0(v)
	case OpOr32:
		return rewriteValueThumb_OpOr32_0(v)
	case OpOr8:
		return rewriteValueThumb_OpOr8_0(v)
	case OpOrB:
		return rewriteValueThumb_OpOrB_0(v)
	case OpPanicBounds:
		return rewriteValueThumb_OpPanicBounds_0(v)
	case OpPanicExtend:
		return rewriteValueThumb_OpPanicExtend_0(v)
	case OpPublicationBarrier:
		return rewriteValueThumb_OpPublicationBarrier_0(v)
	case OpRotateLeft16:
		return rewriteValueThumb_OpRotateLeft16_0(v)
	case OpRotateLeft32:
		return rewriteValueThumb_OpRotateLeft32_0(v)
	case OpRotateLeft8:
		return rewriteValueThumb_OpRotateLeft8_0(v)
	case OpRound32F:
		return rewriteValueThumb_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValueThumb_OpRound64F_0(v)
	case OpRsh16Ux16:
		return rewriteValueThumb_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return rewriteValueThumb_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return rewriteValueThumb_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return rewriteValueThumb_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return rewriteValueThumb_OpRsh16x16_0(v)
	case OpRsh16x32:
		return rewriteValueThumb_OpRsh16x32_0(v)
	case OpRsh16x64:
		return rewriteValueThumb_OpRsh16x64_0(v)
	case OpRsh16x8:
		return rewriteValueThumb_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return rewriteValueThumb_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return rewriteValueThumb_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return rewriteValueThumb_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return rewriteValueThumb_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return rewriteValueThumb_OpRsh32x16_0(v)
	case OpRsh32x32:
		return rewriteValueThumb_OpRsh32x32_0(v)
	case OpRsh32x64:
		return rewriteValueThumb_OpRsh32x64_0(v)
	case OpRsh32x8:
		return rewriteValueThumb_OpRsh32x8_0(v)
	case OpRsh8Ux16:
		return rewriteValueThumb_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return rewriteValueThumb_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return rewriteValueThumb_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return rewriteValueThumb_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return rewriteValueThumb_OpRsh8x16_0(v)
	case OpRsh8x32:
		return rewriteValueThumb_OpRsh8x32_0(v)
	case OpRsh8x64:
		return rewriteValueThumb_OpRsh8x64_0(v)
	case OpRsh8x8:
		return rewriteValueThumb_OpRsh8x8_0(v)
	case OpSignExt16to32:
		return rewriteValueThumb_OpSignExt16to32_0(v)
	case OpSignExt8to16:
		return rewriteValueThumb_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValueThumb_OpSignExt8to32_0(v)
	case OpSignmask:
		return rewriteValueThumb_OpSignmask_0(v)
	case OpSlicemask:
		return rewriteValueThumb_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValueThumb_OpSqrt_0(v)
	case OpStaticCall:
		return rewriteValueThumb_OpStaticCall_0(v)
	case OpStore:
		return rewriteValueThumb_OpStore_0(v)
	case OpSub16:
		return rewriteValueThumb_OpSub16_0(v)
	case OpSub32:
		return rewriteValueThumb_OpSub32_0(v)
	case OpSub32F:
		return rewriteValueThumb_OpSub32F_0(v)
	case OpSub32carry:
		return rewriteValueThumb_OpSub32carry_0(v)
	case OpSub32withcarry:
		return rewriteValueThumb_OpSub32withcarry_0(v)
	case OpSub64F:
		return rewriteValueThumb_OpSub64F_0(v)
	case OpSub8:
		return rewriteValueThumb_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValueThumb_OpSubPtr_0(v)
	case OpThumbADC:
		return rewriteValueThumb_OpThumbADC_0(v) || rewriteValueThumb_OpThumbADC_10(v)
	case OpThumbADCconst:
		return rewriteValueThumb_OpThumbADCconst_0(v)
	case OpThumbADCshiftLL:
		return rewriteValueThumb_OpThumbADCshiftLL_0(v)
	case OpThumbADCshiftRA:
		return rewriteValueThumb_OpThumbADCshiftRA_0(v)
	case OpThumbADCshiftRL:
		return rewriteValueThumb_OpThumbADCshiftRL_0(v)
	case OpThumbADD:
		return rewriteValueThumb_OpThumbADD_0(v) || rewriteValueThumb_OpThumbADD_10(v)
	case OpThumbADDD:
		return rewriteValueThumb_OpThumbADDD_0(v)
	case OpThumbADDF:
		return rewriteValueThumb_OpThumbADDF_0(v)
	case OpThumbADDS:
		return rewriteValueThumb_OpThumbADDS_0(v)
	case OpThumbADDSshiftLL:
		return rewriteValueThumb_OpThumbADDSshiftLL_0(v)
	case OpThumbADDSshiftRA:
		return rewriteValueThumb_OpThumbADDSshiftRA_0(v)
	case OpThumbADDSshiftRL:
		return rewriteValueThumb_OpThumbADDSshiftRL_0(v)
	case OpThumbADDconst:
		return rewriteValueThumb_OpThumbADDconst_0(v)
	case OpThumbADDshiftLL:
		return rewriteValueThumb_OpThumbADDshiftLL_0(v)
	case OpThumbADDshiftRA:
		return rewriteValueThumb_OpThumbADDshiftRA_0(v)
	case OpThumbADDshiftRL:
		return rewriteValueThumb_OpThumbADDshiftRL_0(v)
	case OpThumbAND:
		return rewriteValueThumb_OpThumbAND_0(v) || rewriteValueThumb_OpThumbAND_10(v)
	case OpThumbANDconst:
		return rewriteValueThumb_OpThumbANDconst_0(v)
	case OpThumbANDshiftLL:
		return rewriteValueThumb_OpThumbANDshiftLL_0(v)
	case OpThumbANDshiftRA:
		return rewriteValueThumb_OpThumbANDshiftRA_0(v)
	case OpThumbANDshiftRL:
		return rewriteValueThumb_OpThumbANDshiftRL_0(v)
	case OpThumbBFX:
		return rewriteValueThumb_OpThumbBFX_0(v)
	case OpThumbBFXU:
		return rewriteValueThumb_OpThumbBFXU_0(v)
	case OpThumbBIC:
		return rewriteValueThumb_OpThumbBIC_0(v)
	case OpThumbBICconst:
		return rewriteValueThumb_OpThumbBICconst_0(v)
	case OpThumbBICshiftLL:
		return rewriteValueThumb_OpThumbBICshiftLL_0(v)
	case OpThumbBICshiftRA:
		return rewriteValueThumb_OpThumbBICshiftRA_0(v)
	case OpThumbBICshiftRL:
		return rewriteValueThumb_OpThumbBICshiftRL_0(v)
	case OpThumbCMN:
		return rewriteValueThumb_OpThumbCMN_0(v)
	case OpThumbCMNconst:
		return rewriteValueThumb_OpThumbCMNconst_0(v)
	case OpThumbCMNshiftLL:
		return rewriteValueThumb_OpThumbCMNshiftLL_0(v)
	case OpThumbCMNshiftRA:
		return rewriteValueThumb_OpThumbCMNshiftRA_0(v)
	case OpThumbCMNshiftRL:
		return rewriteValueThumb_OpThumbCMNshiftRL_0(v)
	case OpThumbCMOVWHSconst:
		return rewriteValueThumb_OpThumbCMOVWHSconst_0(v)
	case OpThumbCMOVWLSconst:
		return rewriteValueThumb_OpThumbCMOVWLSconst_0(v)
	case OpThumbCMP:
		return rewriteValueThumb_OpThumbCMP_0(v)
	case OpThumbCMPD:
		return rewriteValueThumb_OpThumbCMPD_0(v)
	case OpThumbCMPF:
		return rewriteValueThumb_OpThumbCMPF_0(v)
	case OpThumbCMPconst:
		return rewriteValueThumb_OpThumbCMPconst_0(v)
	case OpThumbCMPshiftLL:
		return rewriteValueThumb_OpThumbCMPshiftLL_0(v)
	case OpThumbCMPshiftRA:
		return rewriteValueThumb_OpThumbCMPshiftRA_0(v)
	case OpThumbCMPshiftRL:
		return rewriteValueThumb_OpThumbCMPshiftRL_0(v)
	case OpThumbDIV:
		return rewriteValueThumb_OpThumbDIV_0(v)
	case OpThumbDIVU:
		return rewriteValueThumb_OpThumbDIVU_0(v)
	case OpThumbEqual:
		return rewriteValueThumb_OpThumbEqual_0(v)
	case OpThumbGreaterEqual:
		return rewriteValueThumb_OpThumbGreaterEqual_0(v)
	case OpThumbGreaterEqualU:
		return rewriteValueThumb_OpThumbGreaterEqualU_0(v)
	case OpThumbGreaterThan:
		return rewriteValueThumb_OpThumbGreaterThan_0(v)
	case OpThumbGreaterThanU:
		return rewriteValueThumb_OpThumbGreaterThanU_0(v)
	case OpThumbLessEqual:
		return rewriteValueThumb_OpThumbLessEqual_0(v)
	case OpThumbLessEqualU:
		return rewriteValueThumb_OpThumbLessEqualU_0(v)
	case OpThumbLessThan:
		return rewriteValueThumb_OpThumbLessThan_0(v)
	case OpThumbLessThanU:
		return rewriteValueThumb_OpThumbLessThanU_0(v)
	case OpThumbLoadOnce16:
		return rewriteValueThumb_OpThumbLoadOnce16_0(v)
	case OpThumbLoadOnce16idx:
		return rewriteValueThumb_OpThumbLoadOnce16idx_0(v)
	case OpThumbLoadOnce32:
		return rewriteValueThumb_OpThumbLoadOnce32_0(v)
	case OpThumbLoadOnce32idx:
		return rewriteValueThumb_OpThumbLoadOnce32idx_0(v)
	case OpThumbLoadOnce8:
		return rewriteValueThumb_OpThumbLoadOnce8_0(v)
	case OpThumbLoadOnce8idx:
		return rewriteValueThumb_OpThumbLoadOnce8idx_0(v)
	case OpThumbMOVBUload:
		return rewriteValueThumb_OpThumbMOVBUload_0(v)
	case OpThumbMOVBUloadidx:
		return rewriteValueThumb_OpThumbMOVBUloadidx_0(v)
	case OpThumbMOVBUreg:
		return rewriteValueThumb_OpThumbMOVBUreg_0(v)
	case OpThumbMOVBload:
		return rewriteValueThumb_OpThumbMOVBload_0(v)
	case OpThumbMOVBloadidx:
		return rewriteValueThumb_OpThumbMOVBloadidx_0(v)
	case OpThumbMOVBreg:
		return rewriteValueThumb_OpThumbMOVBreg_0(v)
	case OpThumbMOVBstore:
		return rewriteValueThumb_OpThumbMOVBstore_0(v)
	case OpThumbMOVBstoreidx:
		return rewriteValueThumb_OpThumbMOVBstoreidx_0(v)
	case OpThumbMOVDload:
		return rewriteValueThumb_OpThumbMOVDload_0(v)
	case OpThumbMOVDstore:
		return rewriteValueThumb_OpThumbMOVDstore_0(v)
	case OpThumbMOVFload:
		return rewriteValueThumb_OpThumbMOVFload_0(v)
	case OpThumbMOVFstore:
		return rewriteValueThumb_OpThumbMOVFstore_0(v)
	case OpThumbMOVHUload:
		return rewriteValueThumb_OpThumbMOVHUload_0(v)
	case OpThumbMOVHUloadidx:
		return rewriteValueThumb_OpThumbMOVHUloadidx_0(v)
	case OpThumbMOVHUreg:
		return rewriteValueThumb_OpThumbMOVHUreg_0(v)
	case OpThumbMOVHload:
		return rewriteValueThumb_OpThumbMOVHload_0(v)
	case OpThumbMOVHloadidx:
		return rewriteValueThumb_OpThumbMOVHloadidx_0(v)
	case OpThumbMOVHreg:
		return rewriteValueThumb_OpThumbMOVHreg_0(v)
	case OpThumbMOVHstore:
		return rewriteValueThumb_OpThumbMOVHstore_0(v)
	case OpThumbMOVHstoreidx:
		return rewriteValueThumb_OpThumbMOVHstoreidx_0(v)
	case OpThumbMOVWload:
		return rewriteValueThumb_OpThumbMOVWload_0(v)
	case OpThumbMOVWloadidx:
		return rewriteValueThumb_OpThumbMOVWloadidx_0(v)
	case OpThumbMOVWloadshiftLL:
		return rewriteValueThumb_OpThumbMOVWloadshiftLL_0(v)
	case OpThumbMOVWreg:
		return rewriteValueThumb_OpThumbMOVWreg_0(v)
	case OpThumbMOVWstore:
		return rewriteValueThumb_OpThumbMOVWstore_0(v)
	case OpThumbMOVWstoreidx:
		return rewriteValueThumb_OpThumbMOVWstoreidx_0(v)
	case OpThumbMOVWstoreshiftLL:
		return rewriteValueThumb_OpThumbMOVWstoreshiftLL_0(v)
	case OpThumbMUL:
		return rewriteValueThumb_OpThumbMUL_0(v) || rewriteValueThumb_OpThumbMUL_10(v) || rewriteValueThumb_OpThumbMUL_20(v)
	case OpThumbMULA:
		return rewriteValueThumb_OpThumbMULA_0(v) || rewriteValueThumb_OpThumbMULA_10(v) || rewriteValueThumb_OpThumbMULA_20(v)
	case OpThumbMULD:
		return rewriteValueThumb_OpThumbMULD_0(v)
	case OpThumbMULF:
		return rewriteValueThumb_OpThumbMULF_0(v)
	case OpThumbMULS:
		return rewriteValueThumb_OpThumbMULS_0(v) || rewriteValueThumb_OpThumbMULS_10(v) || rewriteValueThumb_OpThumbMULS_20(v)
	case OpThumbMVN:
		return rewriteValueThumb_OpThumbMVN_0(v)
	case OpThumbMVNshiftLL:
		return rewriteValueThumb_OpThumbMVNshiftLL_0(v)
	case OpThumbMVNshiftRA:
		return rewriteValueThumb_OpThumbMVNshiftRA_0(v)
	case OpThumbMVNshiftRL:
		return rewriteValueThumb_OpThumbMVNshiftRL_0(v)
	case OpThumbNEGD:
		return rewriteValueThumb_OpThumbNEGD_0(v)
	case OpThumbNEGF:
		return rewriteValueThumb_OpThumbNEGF_0(v)
	case OpThumbNMULD:
		return rewriteValueThumb_OpThumbNMULD_0(v)
	case OpThumbNMULF:
		return rewriteValueThumb_OpThumbNMULF_0(v)
	case OpThumbNotEqual:
		return rewriteValueThumb_OpThumbNotEqual_0(v)
	case OpThumbOR:
		return rewriteValueThumb_OpThumbOR_0(v)
	case OpThumbORconst:
		return rewriteValueThumb_OpThumbORconst_0(v)
	case OpThumbORshiftLL:
		return rewriteValueThumb_OpThumbORshiftLL_0(v)
	case OpThumbORshiftRA:
		return rewriteValueThumb_OpThumbORshiftRA_0(v)
	case OpThumbORshiftRL:
		return rewriteValueThumb_OpThumbORshiftRL_0(v)
	case OpThumbRSB:
		return rewriteValueThumb_OpThumbRSB_0(v)
	case OpThumbRSBSshiftLL:
		return rewriteValueThumb_OpThumbRSBSshiftLL_0(v)
	case OpThumbRSBSshiftRA:
		return rewriteValueThumb_OpThumbRSBSshiftRA_0(v)
	case OpThumbRSBSshiftRL:
		return rewriteValueThumb_OpThumbRSBSshiftRL_0(v)
	case OpThumbRSBconst:
		return rewriteValueThumb_OpThumbRSBconst_0(v)
	case OpThumbRSBshiftLL:
		return rewriteValueThumb_OpThumbRSBshiftLL_0(v)
	case OpThumbRSBshiftRA:
		return rewriteValueThumb_OpThumbRSBshiftRA_0(v)
	case OpThumbRSBshiftRL:
		return rewriteValueThumb_OpThumbRSBshiftRL_0(v)
	case OpThumbSBC:
		return rewriteValueThumb_OpThumbSBC_0(v)
	case OpThumbSBCconst:
		return rewriteValueThumb_OpThumbSBCconst_0(v)
	case OpThumbSBCshiftLL:
		return rewriteValueThumb_OpThumbSBCshiftLL_0(v)
	case OpThumbSBCshiftRA:
		return rewriteValueThumb_OpThumbSBCshiftRA_0(v)
	case OpThumbSBCshiftRL:
		return rewriteValueThumb_OpThumbSBCshiftRL_0(v)
	case OpThumbSLL:
		return rewriteValueThumb_OpThumbSLL_0(v)
	case OpThumbSLLconst:
		return rewriteValueThumb_OpThumbSLLconst_0(v)
	case OpThumbSRA:
		return rewriteValueThumb_OpThumbSRA_0(v)
	case OpThumbSRAcond:
		return rewriteValueThumb_OpThumbSRAcond_0(v)
	case OpThumbSRAconst:
		return rewriteValueThumb_OpThumbSRAconst_0(v)
	case OpThumbSRL:
		return rewriteValueThumb_OpThumbSRL_0(v)
	case OpThumbSRLconst:
		return rewriteValueThumb_OpThumbSRLconst_0(v)
	case OpThumbSUB:
		return rewriteValueThumb_OpThumbSUB_0(v)
	case OpThumbSUBD:
		return rewriteValueThumb_OpThumbSUBD_0(v)
	case OpThumbSUBF:
		return rewriteValueThumb_OpThumbSUBF_0(v)
	case OpThumbSUBS:
		return rewriteValueThumb_OpThumbSUBS_0(v)
	case OpThumbSUBSshiftLL:
		return rewriteValueThumb_OpThumbSUBSshiftLL_0(v)
	case OpThumbSUBSshiftRA:
		return rewriteValueThumb_OpThumbSUBSshiftRA_0(v)
	case OpThumbSUBSshiftRL:
		return rewriteValueThumb_OpThumbSUBSshiftRL_0(v)
	case OpThumbSUBconst:
		return rewriteValueThumb_OpThumbSUBconst_0(v)
	case OpThumbSUBshiftLL:
		return rewriteValueThumb_OpThumbSUBshiftLL_0(v)
	case OpThumbSUBshiftRA:
		return rewriteValueThumb_OpThumbSUBshiftRA_0(v)
	case OpThumbSUBshiftRL:
		return rewriteValueThumb_OpThumbSUBshiftRL_0(v)
	case OpThumbStoreOnce16:
		return rewriteValueThumb_OpThumbStoreOnce16_0(v)
	case OpThumbStoreOnce16idx:
		return rewriteValueThumb_OpThumbStoreOnce16idx_0(v)
	case OpThumbStoreOnce32:
		return rewriteValueThumb_OpThumbStoreOnce32_0(v)
	case OpThumbStoreOnce32idx:
		return rewriteValueThumb_OpThumbStoreOnce32idx_0(v)
	case OpThumbStoreOnce8:
		return rewriteValueThumb_OpThumbStoreOnce8_0(v)
	case OpThumbStoreOnce8idx:
		return rewriteValueThumb_OpThumbStoreOnce8idx_0(v)
	case OpThumbTEQ:
		return rewriteValueThumb_OpThumbTEQ_0(v)
	case OpThumbTEQconst:
		return rewriteValueThumb_OpThumbTEQconst_0(v)
	case OpThumbTEQshiftLL:
		return rewriteValueThumb_OpThumbTEQshiftLL_0(v)
	case OpThumbTEQshiftRA:
		return rewriteValueThumb_OpThumbTEQshiftRA_0(v)
	case OpThumbTEQshiftRL:
		return rewriteValueThumb_OpThumbTEQshiftRL_0(v)
	case OpThumbTST:
		return rewriteValueThumb_OpThumbTST_0(v)
	case OpThumbTSTconst:
		return rewriteValueThumb_OpThumbTSTconst_0(v)
	case OpThumbTSTshiftLL:
		return rewriteValueThumb_OpThumbTSTshiftLL_0(v)
	case OpThumbTSTshiftRA:
		return rewriteValueThumb_OpThumbTSTshiftRA_0(v)
	case OpThumbTSTshiftRL:
		return rewriteValueThumb_OpThumbTSTshiftRL_0(v)
	case OpThumbXOR:
		return rewriteValueThumb_OpThumbXOR_0(v) || rewriteValueThumb_OpThumbXOR_10(v)
	case OpThumbXORconst:
		return rewriteValueThumb_OpThumbXORconst_0(v)
	case OpThumbXORshiftLL:
		return rewriteValueThumb_OpThumbXORshiftLL_0(v)
	case OpThumbXORshiftRA:
		return rewriteValueThumb_OpThumbXORshiftRA_0(v)
	case OpThumbXORshiftRL:
		return rewriteValueThumb_OpThumbXORshiftRL_0(v)
	case OpThumbXORshiftRR:
		return rewriteValueThumb_OpThumbXORshiftRR_0(v)
	case OpTrunc16to8:
		return rewriteValueThumb_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValueThumb_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValueThumb_OpTrunc32to8_0(v)
	case OpWB:
		return rewriteValueThumb_OpWB_0(v)
	case OpXor16:
		return rewriteValueThumb_OpXor16_0(v)
	case OpXor32:
		return rewriteValueThumb_OpXor32_0(v)
	case OpXor8:
		return rewriteValueThumb_OpXor8_0(v)
	case OpZero:
		return rewriteValueThumb_OpZero_0(v)
	case OpZeroExt16to32:
		return rewriteValueThumb_OpZeroExt16to32_0(v)
	case OpZeroExt8to16:
		return rewriteValueThumb_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValueThumb_OpZeroExt8to32_0(v)
	case OpZeromask:
		return rewriteValueThumb_OpZeromask_0(v)
	}
	return false
}
func rewriteValueThumb_OpAdd16_0(v *Value) bool {
	// match: (Add16 x y)
	// cond:
	// result: (ADD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAdd32_0(v *Value) bool {
	// match: (Add32 x y)
	// cond:
	// result: (ADD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAdd32F_0(v *Value) bool {
	// match: (Add32F x y)
	// cond:
	// result: (ADDF x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbADDF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAdd32carry_0(v *Value) bool {
	// match: (Add32carry x y)
	// cond:
	// result: (ADDS x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbADDS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAdd32withcarry_0(v *Value) bool {
	// match: (Add32withcarry x y c)
	// cond:
	// result: (ADC x y c)
	for {
		c := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpThumbADC)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(c)
		return true
	}
}
func rewriteValueThumb_OpAdd64F_0(v *Value) bool {
	// match: (Add64F x y)
	// cond:
	// result: (ADDD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbADDD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAdd8_0(v *Value) bool {
	// match: (Add8 x y)
	// cond:
	// result: (ADD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAddPtr_0(v *Value) bool {
	// match: (AddPtr x y)
	// cond:
	// result: (ADD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAddr_0(v *Value) bool {
	// match: (Addr {sym} base)
	// cond:
	// result: (MOVWaddr {sym} base)
	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpThumbMOVWaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueThumb_OpAnd16_0(v *Value) bool {
	// match: (And16 x y)
	// cond:
	// result: (AND x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAnd32_0(v *Value) bool {
	// match: (And32 x y)
	// cond:
	// result: (AND x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAnd8_0(v *Value) bool {
	// match: (And8 x y)
	// cond:
	// result: (AND x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAndB_0(v *Value) bool {
	// match: (AndB x y)
	// cond:
	// result: (AND x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpAvg32u_0(v *Value) bool {
	b := v.Block
	// match: (Avg32u <t> x y)
	// cond:
	// result: (ADD (SRLconst <t> (SUB <t> x y) [1]) y)
	for {
		t := v.Type
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpThumbSUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpBitLen32_0(v *Value) bool {
	b := v.Block
	// match: (BitLen32 <t> x)
	// cond:
	// result: (RSBconst [32] (CLZ <t> x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpThumbCLZ, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpBswap32_0(v *Value) bool {
	// match: (Bswap32 x)
	// cond:
	// result: (REV x)
	for {
		x := v.Args[0]
		v.reset(OpThumbREV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpClosureCall_0(v *Value) bool {
	// match: (ClosureCall [argwid] entry closure mem)
	// cond:
	// result: (CALLclosure [argwid] entry closure mem)
	for {
		argwid := v.AuxInt
		mem := v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		v.reset(OpThumbCALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpCom16_0(v *Value) bool {
	// match: (Com16 x)
	// cond:
	// result: (MVN x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCom32_0(v *Value) bool {
	// match: (Com32 x)
	// cond:
	// result: (MVN x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCom8_0(v *Value) bool {
	// match: (Com8 x)
	// cond:
	// result: (MVN x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpConst16_0(v *Value) bool {
	// match: (Const16 [val])
	// cond:
	// result: (MOVWconst [val])
	for {
		val := v.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueThumb_OpConst32_0(v *Value) bool {
	// match: (Const32 [val])
	// cond:
	// result: (MOVWconst [val])
	for {
		val := v.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueThumb_OpConst32F_0(v *Value) bool {
	// match: (Const32F [val])
	// cond:
	// result: (MOVFconst [val])
	for {
		val := v.AuxInt
		v.reset(OpThumbMOVFconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueThumb_OpConst64F_0(v *Value) bool {
	// match: (Const64F [val])
	// cond:
	// result: (MOVDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpThumbMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueThumb_OpConst8_0(v *Value) bool {
	// match: (Const8 [val])
	// cond:
	// result: (MOVWconst [val])
	for {
		val := v.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueThumb_OpConstBool_0(v *Value) bool {
	// match: (ConstBool [b])
	// cond:
	// result: (MOVWconst [b])
	for {
		b := v.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueThumb_OpConstNil_0(v *Value) bool {
	// match: (ConstNil)
	// cond:
	// result: (MOVWconst [0])
	for {
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueThumb_OpCtz16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz16 <t> x)
	// cond:
	// result: (CLZ <t> (RBIT <typ.UInt32> (ORconst <typ.UInt32> [0x10000] x)))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpThumbCLZ)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpThumbRBIT, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpThumbORconst, typ.UInt32)
		v1.AuxInt = 0x10000
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpCtz16NonZero_0(v *Value) bool {
	// match: (Ctz16NonZero x)
	// cond:
	// result: (Ctz32 x)
	for {
		x := v.Args[0]
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCtz32_0(v *Value) bool {
	b := v.Block
	// match: (Ctz32 <t> x)
	// cond:
	// result: (CLZ <t> (RBIT <t> x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpThumbCLZ)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpThumbRBIT, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpCtz32NonZero_0(v *Value) bool {
	// match: (Ctz32NonZero x)
	// cond:
	// result: (Ctz32 x)
	for {
		x := v.Args[0]
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCtz8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz8 <t> x)
	// cond:
	// result: (CLZ <t> (RBIT <typ.UInt32> (ORconst <typ.UInt32> [0x100] x)))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpThumbCLZ)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpThumbRBIT, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpThumbORconst, typ.UInt32)
		v1.AuxInt = 0x100
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpCtz8NonZero_0(v *Value) bool {
	// match: (Ctz8NonZero x)
	// cond:
	// result: (Ctz32 x)
	for {
		x := v.Args[0]
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCvt32Fto32_0(v *Value) bool {
	// match: (Cvt32Fto32 x)
	// cond:
	// result: (MOVFW x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVFW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCvt32Fto32U_0(v *Value) bool {
	// match: (Cvt32Fto32U x)
	// cond:
	// result: (MOVFWU x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVFWU)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCvt32Fto64F_0(v *Value) bool {
	// match: (Cvt32Fto64F x)
	// cond:
	// result: (MOVFD x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVFD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCvt32Uto32F_0(v *Value) bool {
	// match: (Cvt32Uto32F x)
	// cond:
	// result: (MOVWUF x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVWUF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCvt32Uto64F_0(v *Value) bool {
	// match: (Cvt32Uto64F x)
	// cond:
	// result: (MOVWUD x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVWUD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCvt32to32F_0(v *Value) bool {
	// match: (Cvt32to32F x)
	// cond:
	// result: (MOVWF x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVWF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCvt32to64F_0(v *Value) bool {
	// match: (Cvt32to64F x)
	// cond:
	// result: (MOVWD x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVWD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCvt64Fto32_0(v *Value) bool {
	// match: (Cvt64Fto32 x)
	// cond:
	// result: (MOVDW x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVDW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCvt64Fto32F_0(v *Value) bool {
	// match: (Cvt64Fto32F x)
	// cond:
	// result: (MOVDF x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVDF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpCvt64Fto32U_0(v *Value) bool {
	// match: (Cvt64Fto32U x)
	// cond:
	// result: (MOVDWU x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVDWU)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpDiv16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// cond:
	// result: (DIV (SignExt16to32 x) (SignExt16to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbDIV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpDiv16u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// cond:
	// result: (DIVU (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbDIVU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpDiv32_0(v *Value) bool {
	// match: (Div32 x y)
	// cond:
	// result: (DIV x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbDIV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpDiv32F_0(v *Value) bool {
	// match: (Div32F x y)
	// cond:
	// result: (DIVF x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbDIVF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpDiv32u_0(v *Value) bool {
	// match: (Div32u x y)
	// cond:
	// result: (DIVU x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbDIVU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpDiv64F_0(v *Value) bool {
	// match: (Div64F x y)
	// cond:
	// result: (DIVD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbDIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpDiv8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// cond:
	// result: (DIV (SignExt8to32 x) (SignExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbDIV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpDiv8u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// cond:
	// result: (DIVU (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbDIVU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpEq16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// cond:
	// result: (Equal (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEq32_0(v *Value) bool {
	b := v.Block
	// match: (Eq32 x y)
	// cond:
	// result: (Equal (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEq32F_0(v *Value) bool {
	b := v.Block
	// match: (Eq32F x y)
	// cond:
	// result: (Equal (CMPF x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEq64F_0(v *Value) bool {
	b := v.Block
	// match: (Eq64F x y)
	// cond:
	// result: (Equal (CMPD x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEq8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// cond:
	// result: (Equal (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEqB_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// cond:
	// result: (XORconst [1] (XOR <typ.Bool> x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpThumbXOR, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpEqPtr_0(v *Value) bool {
	b := v.Block
	// match: (EqPtr x y)
	// cond:
	// result: (Equal (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGeq16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Geq16 x y)
	// cond:
	// result: (GreaterEqual (CMP (SignExt16to32 x) (SignExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGeq16U_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Geq16U x y)
	// cond:
	// result: (GreaterEqualU (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGeq32_0(v *Value) bool {
	b := v.Block
	// match: (Geq32 x y)
	// cond:
	// result: (GreaterEqual (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGeq32F_0(v *Value) bool {
	b := v.Block
	// match: (Geq32F x y)
	// cond:
	// result: (GreaterEqual (CMPF x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGeq32U_0(v *Value) bool {
	b := v.Block
	// match: (Geq32U x y)
	// cond:
	// result: (GreaterEqualU (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGeq64F_0(v *Value) bool {
	b := v.Block
	// match: (Geq64F x y)
	// cond:
	// result: (GreaterEqual (CMPD x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGeq8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Geq8 x y)
	// cond:
	// result: (GreaterEqual (CMP (SignExt8to32 x) (SignExt8to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGeq8U_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Geq8U x y)
	// cond:
	// result: (GreaterEqualU (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGetCallerPC_0(v *Value) bool {
	// match: (GetCallerPC)
	// cond:
	// result: (LoweredGetCallerPC)
	for {
		v.reset(OpThumbLoweredGetCallerPC)
		return true
	}
}
func rewriteValueThumb_OpGetCallerSP_0(v *Value) bool {
	// match: (GetCallerSP)
	// cond:
	// result: (LoweredGetCallerSP)
	for {
		v.reset(OpThumbLoweredGetCallerSP)
		return true
	}
}
func rewriteValueThumb_OpGetClosurePtr_0(v *Value) bool {
	// match: (GetClosurePtr)
	// cond:
	// result: (LoweredGetClosurePtr)
	for {
		v.reset(OpThumbLoweredGetClosurePtr)
		return true
	}
}
func rewriteValueThumb_OpGreater16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Greater16 x y)
	// cond:
	// result: (GreaterThan (CMP (SignExt16to32 x) (SignExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGreater16U_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Greater16U x y)
	// cond:
	// result: (GreaterThanU (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGreater32_0(v *Value) bool {
	b := v.Block
	// match: (Greater32 x y)
	// cond:
	// result: (GreaterThan (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGreater32F_0(v *Value) bool {
	b := v.Block
	// match: (Greater32F x y)
	// cond:
	// result: (GreaterThan (CMPF x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMPF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGreater32U_0(v *Value) bool {
	b := v.Block
	// match: (Greater32U x y)
	// cond:
	// result: (GreaterThanU (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGreater64F_0(v *Value) bool {
	b := v.Block
	// match: (Greater64F x y)
	// cond:
	// result: (GreaterThan (CMPD x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMPD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGreater8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Greater8 x y)
	// cond:
	// result: (GreaterThan (CMP (SignExt8to32 x) (SignExt8to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpGreater8U_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Greater8U x y)
	// cond:
	// result: (GreaterThanU (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpHmul32_0(v *Value) bool {
	// match: (Hmul32 x y)
	// cond:
	// result: (HMUL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbHMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpHmul32u_0(v *Value) bool {
	// match: (Hmul32u x y)
	// cond:
	// result: (HMULU x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbHMULU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpInterCall_0(v *Value) bool {
	// match: (InterCall [argwid] entry mem)
	// cond:
	// result: (CALLinter [argwid] entry mem)
	for {
		argwid := v.AuxInt
		mem := v.Args[1]
		entry := v.Args[0]
		v.reset(OpThumbCALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpIsInBounds_0(v *Value) bool {
	b := v.Block
	// match: (IsInBounds idx len)
	// cond:
	// result: (LessThanU (CMP idx len))
	for {
		len := v.Args[1]
		idx := v.Args[0]
		v.reset(OpThumbLessThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	// match: (IsNonNil ptr)
	// cond:
	// result: (NotEqual (CMPconst [0] ptr))
	for {
		ptr := v.Args[0]
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(ptr)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpIsSliceInBounds_0(v *Value) bool {
	b := v.Block
	// match: (IsSliceInBounds idx len)
	// cond:
	// result: (LessEqualU (CMP idx len))
	for {
		len := v.Args[1]
		idx := v.Args[0]
		v.reset(OpThumbLessEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// cond:
	// result: (LessEqual (CMP (SignExt16to32 x) (SignExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq16U_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// cond:
	// result: (LessEqualU (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq32_0(v *Value) bool {
	b := v.Block
	// match: (Leq32 x y)
	// cond:
	// result: (LessEqual (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq32F_0(v *Value) bool {
	b := v.Block
	// match: (Leq32F x y)
	// cond:
	// result: (GreaterEqual (CMPF y x))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPF, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq32U_0(v *Value) bool {
	b := v.Block
	// match: (Leq32U x y)
	// cond:
	// result: (LessEqualU (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq64F_0(v *Value) bool {
	b := v.Block
	// match: (Leq64F x y)
	// cond:
	// result: (GreaterEqual (CMPD y x))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPD, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// cond:
	// result: (LessEqual (CMP (SignExt8to32 x) (SignExt8to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLeq8U_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// cond:
	// result: (LessEqualU (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessEqualU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// cond:
	// result: (LessThan (CMP (SignExt16to32 x) (SignExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess16U_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// cond:
	// result: (LessThanU (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess32_0(v *Value) bool {
	b := v.Block
	// match: (Less32 x y)
	// cond:
	// result: (LessThan (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess32F_0(v *Value) bool {
	b := v.Block
	// match: (Less32F x y)
	// cond:
	// result: (GreaterThan (CMPF y x))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMPF, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess32U_0(v *Value) bool {
	b := v.Block
	// match: (Less32U x y)
	// cond:
	// result: (LessThanU (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess64F_0(v *Value) bool {
	b := v.Block
	// match: (Less64F x y)
	// cond:
	// result: (GreaterThan (CMPD y x))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbGreaterThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMPD, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// cond:
	// result: (LessThan (CMP (SignExt8to32 x) (SignExt8to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessThan)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLess8U_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// cond:
	// result: (LessThanU (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbLessThanU)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLoad_0(v *Value) bool {
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpThumbMOVBUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && isSigned(t))
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpThumbMOVBload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && !isSigned(t))
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpThumbMOVBUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && isSigned(t))
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpThumbMOVHload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && !isSigned(t))
	// result: (MOVHUload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpThumbMOVHUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) || isPtr(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is32BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpThumbMOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (MOVFload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpThumbMOVFload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpThumbMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpLocalAddr_0(v *Value) bool {
	// match: (LocalAddr {sym} base _)
	// cond:
	// result: (MOVWaddr {sym} base)
	for {
		sym := v.Aux
		_ = v.Args[1]
		base := v.Args[0]
		v.reset(OpThumbMOVWaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueThumb_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 x y)
	// cond:
	// result: (CMOVWHSconst (SLL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueThumb_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	// match: (Lsh16x32 x y)
	// cond:
	// result: (CMOVWHSconst (SLL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpLsh16x64_0(v *Value) bool {
	// match: (Lsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SLLconst x [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (Const16 [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 x y)
	// cond:
	// result: (SLL x (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 x y)
	// cond:
	// result: (CMOVWHSconst (SLL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueThumb_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	// match: (Lsh32x32 x y)
	// cond:
	// result: (CMOVWHSconst (SLL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpLsh32x64_0(v *Value) bool {
	// match: (Lsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SLLconst x [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (Const32 [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 x y)
	// cond:
	// result: (SLL x (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 x y)
	// cond:
	// result: (CMOVWHSconst (SLL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueThumb_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	// match: (Lsh8x32 x y)
	// cond:
	// result: (CMOVWHSconst (SLL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpLsh8x64_0(v *Value) bool {
	// match: (Lsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SLLconst x [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (Const8 [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 x y)
	// cond:
	// result: (SLL x (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpMMIOLoad16_0(v *Value) bool {
	// match: (MMIOLoad16 ptr mem)
	// cond:
	// result: (LoadOnce16 ptr mem)
	for {
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpThumbLoadOnce16)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOLoad32_0(v *Value) bool {
	// match: (MMIOLoad32 ptr mem)
	// cond:
	// result: (LoadOnce32 ptr mem)
	for {
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpThumbLoadOnce32)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOLoad8_0(v *Value) bool {
	// match: (MMIOLoad8 ptr mem)
	// cond:
	// result: (LoadOnce8 ptr mem)
	for {
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpThumbLoadOnce8)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOMB_0(v *Value) bool {
	// match: (MMIOMB mem)
	// cond:
	// result: (DSB mem)
	for {
		mem := v.Args[0]
		v.reset(OpThumbDSB)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOStore16_0(v *Value) bool {
	// match: (MMIOStore16 ptr val mem)
	// cond:
	// result: (StoreOnce16 ptr val mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpThumbStoreOnce16)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOStore32_0(v *Value) bool {
	// match: (MMIOStore32 ptr val mem)
	// cond:
	// result: (StoreOnce32 ptr val mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpThumbStoreOnce32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpMMIOStore8_0(v *Value) bool {
	// match: (MMIOStore8 ptr val mem)
	// cond:
	// result: (StoreOnce8 ptr val mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		v.reset(OpThumbStoreOnce8)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpMod16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// cond:
	// result: (Mod32 (SignExt16to32 x) (SignExt16to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpMod16u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// cond:
	// result: (Mod32u (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpMod32_0(v *Value) bool {
	b := v.Block
	// match: (Mod32 x y)
	// cond:
	// result: (SUB x (MUL <y.Type> y (DIV <x.Type> x y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpThumbMUL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpThumbDIV, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpMod32u_0(v *Value) bool {
	b := v.Block
	// match: (Mod32u x y)
	// cond:
	// result: (SUB x (MUL <y.Type> y (DIVU <x.Type> x y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpThumbMUL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpThumbDIVU, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpMod8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// cond:
	// result: (Mod32 (SignExt8to32 x) (SignExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpMod8u_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// cond:
	// result: (Mod32u (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpMove_0(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// cond:
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// cond:
	// result: (MOVBstore dst (MOVBUload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpThumbMOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHUload src mem) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		t := v.Aux
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpThumbMOVHUload, typ.UInt16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// cond:
	// result: (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = 1
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v0.AuxInt = 1
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpThumbMOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWload, typ.UInt32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [2] dst (MOVHUload [2] src mem) (MOVHstore dst (MOVHUload src mem) mem))
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpThumbMOVHUload, typ.UInt16)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVHstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpThumbMOVHUload, typ.UInt16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// cond:
	// result: (MOVBstore [3] dst (MOVBUload [3] src mem) (MOVBstore [2] dst (MOVBUload [2] src mem) (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v4.AuxInt = 1
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v6.AddArg(src)
		v6.AddArg(mem)
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [3] dst src mem)
	// cond:
	// result: (MOVBstore [2] dst (MOVBUload [2] src mem) (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v2.AuxInt = 1
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpThumbMOVBUload, typ.UInt8)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: s%4 == 0 && s > 4 && s <= 512 && t.(*types.Type).Alignment()%4 == 0 && !config.noDuffDevice
	// result: (DUFFCOPY [8 * (128 - s/4)] dst src mem)
	for {
		s := v.AuxInt
		t := v.Aux
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !(s%4 == 0 && s > 4 && s <= 512 && t.(*types.Type).Alignment()%4 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpThumbDUFFCOPY)
		v.AuxInt = 8 * (128 - s/4)
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: (s > 512 || config.noDuffDevice) || t.(*types.Type).Alignment()%4 != 0
	// result: (LoweredMove [t.(*types.Type).Alignment()] dst src (ADDconst <src.Type> src [s-moveSize(t.(*types.Type).Alignment(), config)]) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		mem := v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		if !((s > 512 || config.noDuffDevice) || t.(*types.Type).Alignment()%4 != 0) {
			break
		}
		v.reset(OpThumbLoweredMove)
		v.AuxInt = t.(*types.Type).Alignment()
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpThumbADDconst, src.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpMul16_0(v *Value) bool {
	// match: (Mul16 x y)
	// cond:
	// result: (MUL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpMul32_0(v *Value) bool {
	// match: (Mul32 x y)
	// cond:
	// result: (MUL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpMul32F_0(v *Value) bool {
	// match: (Mul32F x y)
	// cond:
	// result: (MULF x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpMul32uhilo_0(v *Value) bool {
	// match: (Mul32uhilo x y)
	// cond:
	// result: (MULLU x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbMULLU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpMul64F_0(v *Value) bool {
	// match: (Mul64F x y)
	// cond:
	// result: (MULD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpMul8_0(v *Value) bool {
	// match: (Mul8 x y)
	// cond:
	// result: (MUL x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpNeg16_0(v *Value) bool {
	// match: (Neg16 x)
	// cond:
	// result: (RSBconst [0] x)
	for {
		x := v.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpNeg32_0(v *Value) bool {
	// match: (Neg32 x)
	// cond:
	// result: (RSBconst [0] x)
	for {
		x := v.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpNeg32F_0(v *Value) bool {
	// match: (Neg32F x)
	// cond:
	// result: (NEGF x)
	for {
		x := v.Args[0]
		v.reset(OpThumbNEGF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpNeg64F_0(v *Value) bool {
	// match: (Neg64F x)
	// cond:
	// result: (NEGD x)
	for {
		x := v.Args[0]
		v.reset(OpThumbNEGD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpNeg8_0(v *Value) bool {
	// match: (Neg8 x)
	// cond:
	// result: (RSBconst [0] x)
	for {
		x := v.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpNeq16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// cond:
	// result: (NotEqual (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNeq32_0(v *Value) bool {
	b := v.Block
	// match: (Neq32 x y)
	// cond:
	// result: (NotEqual (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNeq32F_0(v *Value) bool {
	b := v.Block
	// match: (Neq32F x y)
	// cond:
	// result: (NotEqual (CMPF x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNeq64F_0(v *Value) bool {
	b := v.Block
	// match: (Neq64F x y)
	// cond:
	// result: (NotEqual (CMPD x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMPD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNeq8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// cond:
	// result: (NotEqual (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNeqB_0(v *Value) bool {
	// match: (NeqB x y)
	// cond:
	// result: (XOR x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpNeqPtr_0(v *Value) bool {
	b := v.Block
	// match: (NeqPtr x y)
	// cond:
	// result: (NotEqual (CMP x y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbNotEqual)
		v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpNilCheck_0(v *Value) bool {
	// match: (NilCheck ptr mem)
	// cond:
	// result: (LoweredNilCheck ptr mem)
	for {
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpThumbLoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpNot_0(v *Value) bool {
	// match: (Not x)
	// cond:
	// result: (XORconst [1] x)
	for {
		x := v.Args[0]
		v.reset(OpThumbXORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpOffPtr_0(v *Value) bool {
	// match: (OffPtr [off] ptr:(SP))
	// cond:
	// result: (MOVWaddr [off] ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpThumbMOVWaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// cond:
	// result: (ADDconst [off] ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueThumb_OpOr16_0(v *Value) bool {
	// match: (Or16 x y)
	// cond:
	// result: (OR x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpOr32_0(v *Value) bool {
	// match: (Or32 x y)
	// cond:
	// result: (OR x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpOr8_0(v *Value) bool {
	// match: (Or8 x y)
	// cond:
	// result: (OR x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpOrB_0(v *Value) bool {
	// match: (OrB x y)
	// cond:
	// result: (OR x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpPanicBounds_0(v *Value) bool {
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicBoundsA [kind] x y mem)
	for {
		kind := v.AuxInt
		mem := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpThumbLoweredPanicBoundsA)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicBoundsB [kind] x y mem)
	for {
		kind := v.AuxInt
		mem := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpThumbLoweredPanicBoundsB)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicBoundsC [kind] x y mem)
	for {
		kind := v.AuxInt
		mem := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpThumbLoweredPanicBoundsC)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpPanicExtend_0(v *Value) bool {
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicExtendA [kind] hi lo y mem)
	for {
		kind := v.AuxInt
		mem := v.Args[3]
		hi := v.Args[0]
		lo := v.Args[1]
		y := v.Args[2]
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpThumbLoweredPanicExtendA)
		v.AuxInt = kind
		v.AddArg(hi)
		v.AddArg(lo)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicExtendB [kind] hi lo y mem)
	for {
		kind := v.AuxInt
		mem := v.Args[3]
		hi := v.Args[0]
		lo := v.Args[1]
		y := v.Args[2]
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpThumbLoweredPanicExtendB)
		v.AuxInt = kind
		v.AddArg(hi)
		v.AddArg(lo)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicExtendC [kind] hi lo y mem)
	for {
		kind := v.AuxInt
		mem := v.Args[3]
		hi := v.Args[0]
		lo := v.Args[1]
		y := v.Args[2]
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpThumbLoweredPanicExtendC)
		v.AuxInt = kind
		v.AddArg(hi)
		v.AddArg(lo)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpPublicationBarrier_0(v *Value) bool {
	// match: (PublicationBarrier mem)
	// cond:
	// result: (DMB_ST mem)
	for {
		mem := v.Args[0]
		v.reset(OpThumbDMB_ST)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpRotateLeft16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVWconst [c]))
	// cond:
	// result: (Or16 (Lsh16x32 <t> x (MOVWconst [c&15])) (Rsh16Ux32 <t> x (MOVWconst [-c&15])))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x32, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v1.AuxInt = c & 15
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux32, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v3.AuxInt = -c & 15
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueThumb_OpRotateLeft32_0(v *Value) bool {
	// match: (RotateLeft32 x (MOVWconst [c]))
	// cond:
	// result: (SRRconst [-c&31] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSRRconst)
		v.AuxInt = -c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpRotateLeft8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVWconst [c]))
	// cond:
	// result: (Or8 (Lsh8x32 <t> x (MOVWconst [c&7])) (Rsh8Ux32 <t> x (MOVWconst [-c&7])))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x32, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v1.AuxInt = c & 7
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux32, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v3.AuxInt = -c & 7
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueThumb_OpRound32F_0(v *Value) bool {
	// match: (Round32F x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpRound64F_0(v *Value) bool {
	// match: (Round64F x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 x y)
	// cond:
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt16to32 x) (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v3.AuxInt = 256
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueThumb_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 x y)
	// cond:
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt16to32 x) y) (CMPconst [256] y) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueThumb_OpRsh16Ux64_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRLconst (SLLconst <typ.UInt32> x [16]) [c+16])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (Const16 [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 x y)
	// cond:
	// result: (SRL (ZeroExt16to32 x) (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRL)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 x y)
	// cond:
	// result: (SRAcond (SignExt16to32 x) (ZeroExt16to32 y) (CMPconst [256] (ZeroExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueThumb_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 x y)
	// cond:
	// result: (SRAcond (SignExt16to32 x) y (CMPconst [256] y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRAconst (SLLconst <typ.UInt32> x [16]) [c+16])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (SRAconst (SLLconst <typ.UInt32> x [16]) [31])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 x y)
	// cond:
	// result: (SRA (SignExt16to32 x) (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 x y)
	// cond:
	// result: (CMOVWHSconst (SRL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueThumb_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	// match: (Rsh32Ux32 x y)
	// cond:
	// result: (CMOVWHSconst (SRL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpRsh32Ux64_0(v *Value) bool {
	// match: (Rsh32Ux64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRLconst x [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (Const32 [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 x y)
	// cond:
	// result: (SRL x (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 x y)
	// cond:
	// result: (SRAcond x (ZeroExt16to32 y) (CMPconst [256] (ZeroExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRAcond)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	// match: (Rsh32x32 x y)
	// cond:
	// result: (SRAcond x y (CMPconst [256] y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRAcond)
		v.AddArg(x)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = 256
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpRsh32x64_0(v *Value) bool {
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRAconst x [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (SRAconst x [31])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 x y)
	// cond:
	// result: (SRA x (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 x y)
	// cond:
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt8to32 x) (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v3.AuxInt = 256
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueThumb_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 x y)
	// cond:
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt8to32 x) y) (CMPconst [256] y) [0])
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpThumbSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueThumb_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRLconst (SLLconst <typ.UInt32> x [24]) [c+24])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (Const8 [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 x y)
	// cond:
	// result: (SRL (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRL)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 x y)
	// cond:
	// result: (SRAcond (SignExt8to32 x) (ZeroExt16to32 y) (CMPconst [256] (ZeroExt16to32 y)))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueThumb_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 x y)
	// cond:
	// result: (SRAcond (SignExt8to32 x) y (CMPconst [256] y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRAconst (SLLconst <typ.UInt32> x [24]) [c+24])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (SRAconst (SLLconst <typ.UInt32> x [24]) [31])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueThumb_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 x y)
	// cond:
	// result: (SRA (SignExt8to32 x) (ZeroExt8to32 y))
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueThumb_OpSignExt16to32_0(v *Value) bool {
	// match: (SignExt16to32 x)
	// cond:
	// result: (MOVHreg x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpSignExt8to16_0(v *Value) bool {
	// match: (SignExt8to16 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpSignExt8to32_0(v *Value) bool {
	// match: (SignExt8to32 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpSignmask_0(v *Value) bool {
	// match: (Signmask x)
	// cond:
	// result: (SRAconst x [31])
	for {
		x := v.Args[0]
		v.reset(OpThumbSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpSlicemask_0(v *Value) bool {
	b := v.Block
	// match: (Slicemask <t> x)
	// cond:
	// result: (SRAconst (RSBconst <t> [0] x) [31])
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpThumbSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpThumbRSBconst, t)
		v0.AuxInt = 0
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueThumb_OpSqrt_0(v *Value) bool {
	// match: (Sqrt x)
	// cond:
	// result: (SQRTD x)
	for {
		x := v.Args[0]
		v.reset(OpThumbSQRTD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpStaticCall_0(v *Value) bool {
	// match: (StaticCall [argwid] {target} mem)
	// cond:
	// result: (CALLstatic [argwid] {target} mem)
	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpThumbCALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpStore_0(v *Value) bool {
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpThumbMOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)
	// result: (MOVWstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpThumbMOVWstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (MOVFstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpThumbMOVFstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (MOVDstore ptr val mem)
	for {
		t := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		if !(t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpThumbMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpSub16_0(v *Value) bool {
	// match: (Sub16 x y)
	// cond:
	// result: (SUB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpSub32_0(v *Value) bool {
	// match: (Sub32 x y)
	// cond:
	// result: (SUB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpSub32F_0(v *Value) bool {
	// match: (Sub32F x y)
	// cond:
	// result: (SUBF x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSUBF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpSub32carry_0(v *Value) bool {
	// match: (Sub32carry x y)
	// cond:
	// result: (SUBS x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSUBS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpSub32withcarry_0(v *Value) bool {
	// match: (Sub32withcarry x y c)
	// cond:
	// result: (SBC x y c)
	for {
		c := v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpThumbSBC)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(c)
		return true
	}
}
func rewriteValueThumb_OpSub64F_0(v *Value) bool {
	// match: (Sub64F x y)
	// cond:
	// result: (SUBD x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSUBD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpSub8_0(v *Value) bool {
	// match: (Sub8 x y)
	// cond:
	// result: (SUB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpSubPtr_0(v *Value) bool {
	// match: (SubPtr x y)
	// cond:
	// result: (SUB x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpThumbADC_0(v *Value) bool {
	// match: (ADC (MOVWconst [c]) x flags)
	// cond:
	// result: (ADCconst [c] x flags)
	for {
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpThumbADCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (ADC x (MOVWconst [c]) flags)
	// cond:
	// result: (ADCconst [c] x flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (ADC x (MOVWconst [c]) flags)
	// cond:
	// result: (ADCconst [c] x flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (ADC (MOVWconst [c]) x flags)
	// cond:
	// result: (ADCconst [c] x flags)
	for {
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpThumbADCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (ADC x (SLLconst [c] y) flags)
	// cond:
	// result: (ADCshiftLL x y [c] flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (ADC (SLLconst [c] y) x flags)
	// cond:
	// result: (ADCshiftLL x y [c] flags)
	for {
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpThumbADCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (ADC (SLLconst [c] y) x flags)
	// cond:
	// result: (ADCshiftLL x y [c] flags)
	for {
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpThumbADCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (ADC x (SLLconst [c] y) flags)
	// cond:
	// result: (ADCshiftLL x y [c] flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (ADC x (SRLconst [c] y) flags)
	// cond:
	// result: (ADCshiftRL x y [c] flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (ADC (SRLconst [c] y) x flags)
	// cond:
	// result: (ADCshiftRL x y [c] flags)
	for {
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpThumbADCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADC_10(v *Value) bool {
	// match: (ADC (SRLconst [c] y) x flags)
	// cond:
	// result: (ADCshiftRL x y [c] flags)
	for {
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpThumbADCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (ADC x (SRLconst [c] y) flags)
	// cond:
	// result: (ADCshiftRL x y [c] flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (ADC x (SRAconst [c] y) flags)
	// cond:
	// result: (ADCshiftRA x y [c] flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (ADC (SRAconst [c] y) x flags)
	// cond:
	// result: (ADCshiftRA x y [c] flags)
	for {
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpThumbADCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (ADC (SRAconst [c] y) x flags)
	// cond:
	// result: (ADCshiftRA x y [c] flags)
	for {
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpThumbADCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (ADC x (SRAconst [c] y) flags)
	// cond:
	// result: (ADCshiftRA x y [c] flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADCconst_0(v *Value) bool {
	// match: (ADCconst [c] (ADDconst [d] x) flags)
	// cond:
	// result: (ADCconst [int64(int32(c+d))] x flags)
	for {
		c := v.AuxInt
		flags := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbADCconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (ADCconst [c] (SUBconst [d] x) flags)
	// cond:
	// result: (ADCconst [int64(int32(c-d))] x flags)
	for {
		c := v.AuxInt
		flags := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbADCconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADCshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (ADCshiftLL (MOVWconst [c]) x [d] flags)
	// cond:
	// result: (ADCconst [c] (SLLconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpThumbADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (ADCshiftLL x (MOVWconst [c]) [d] flags)
	// cond:
	// result: (ADCconst x [int64(int32(uint32(c)<<uint64(d)))] flags)
	for {
		d := v.AuxInt
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADCconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADCshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (ADCshiftRA (MOVWconst [c]) x [d] flags)
	// cond:
	// result: (ADCconst [c] (SRAconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpThumbADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (ADCshiftRA x (MOVWconst [c]) [d] flags)
	// cond:
	// result: (ADCconst x [int64(int32(c)>>uint64(d))] flags)
	for {
		d := v.AuxInt
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADCconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADCshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (ADCshiftRL (MOVWconst [c]) x [d] flags)
	// cond:
	// result: (ADCconst [c] (SRLconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		flags := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpThumbADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (ADCshiftRL x (MOVWconst [c]) [d] flags)
	// cond:
	// result: (ADCconst x [int64(int32(uint32(c)>>uint64(d)))] flags)
	for {
		d := v.AuxInt
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADCconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADD_0(v *Value) bool {
	// match: (ADD x (MOVWconst [c]))
	// cond:
	// result: (ADDconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ADD (MOVWconst [c]) x)
	// cond:
	// result: (ADDconst [c] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ADD x (SLLconst [c] y))
	// cond:
	// result: (ADDshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADD (SLLconst [c] y) x)
	// cond:
	// result: (ADDshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbADDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADD x (SRLconst [c] y))
	// cond:
	// result: (ADDshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADD (SRLconst [c] y) x)
	// cond:
	// result: (ADDshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbADDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADD x (SRAconst [c] y))
	// cond:
	// result: (ADDshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADD (SRAconst [c] y) x)
	// cond:
	// result: (ADDshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbADDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADD x (RSBconst [0] y))
	// cond:
	// result: (SUB x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbRSBconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpThumbSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADD (RSBconst [0] y) x)
	// cond:
	// result: (SUB x y)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbRSBconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		y := v_0.Args[0]
		v.reset(OpThumbSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADD_10(v *Value) bool {
	b := v.Block
	// match: (ADD <t> (RSBconst [c] x) (RSBconst [d] y))
	// cond:
	// result: (RSBconst [c+d] (ADD <t> x y))
	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbRSBconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbRSBconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = c + d
		v0 := b.NewValue0(v.Pos, OpThumbADD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ADD <t> (RSBconst [d] y) (RSBconst [c] x))
	// cond:
	// result: (RSBconst [c+d] (ADD <t> x y))
	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbRSBconst {
			break
		}
		d := v_0.AuxInt
		y := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbRSBconst {
			break
		}
		c := v_1.AuxInt
		x := v_1.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = c + d
		v0 := b.NewValue0(v.Pos, OpThumbADD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ADD (MUL x y) a)
	// cond:
	// result: (MULA x y a)
	for {
		a := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMUL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpThumbMULA)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(a)
		return true
	}
	// match: (ADD a (MUL x y))
	// cond:
	// result: (MULA x y a)
	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMUL {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		v.reset(OpThumbMULA)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDD_0(v *Value) bool {
	// match: (ADDD a (MULD x y))
	// cond: a.Uses == 1
	// result: (MULAD a x y)
	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMULD {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULAD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDD (MULD x y) a)
	// cond: a.Uses == 1
	// result: (MULAD a x y)
	for {
		a := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMULD {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULAD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDD a (NMULD x y))
	// cond: a.Uses == 1
	// result: (MULSD a x y)
	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbNMULD {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULSD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDD (NMULD x y) a)
	// cond: a.Uses == 1
	// result: (MULSD a x y)
	for {
		a := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbNMULD {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULSD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDF_0(v *Value) bool {
	// match: (ADDF a (MULF x y))
	// cond: a.Uses == 1
	// result: (MULAF a x y)
	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMULF {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULAF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDF (MULF x y) a)
	// cond: a.Uses == 1
	// result: (MULAF a x y)
	for {
		a := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMULF {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULAF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDF a (NMULF x y))
	// cond: a.Uses == 1
	// result: (MULSF a x y)
	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbNMULF {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULSF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDF (NMULF x y) a)
	// cond: a.Uses == 1
	// result: (MULSF a x y)
	for {
		a := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbNMULF {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULSF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDS_0(v *Value) bool {
	// match: (ADDS x (MOVWconst [c]))
	// cond:
	// result: (ADDSconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADDSconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ADDS (MOVWconst [c]) x)
	// cond:
	// result: (ADDSconst [c] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbADDSconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ADDS x (SLLconst [c] y))
	// cond:
	// result: (ADDSshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADDSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDS (SLLconst [c] y) x)
	// cond:
	// result: (ADDSshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbADDSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDS x (SRLconst [c] y))
	// cond:
	// result: (ADDSshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADDSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDS (SRLconst [c] y) x)
	// cond:
	// result: (ADDSshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbADDSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDS x (SRAconst [c] y))
	// cond:
	// result: (ADDSshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbADDSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDS (SRAconst [c] y) x)
	// cond:
	// result: (ADDSshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbADDSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDSshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (ADDSshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (ADDSconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (ADDSconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADDSconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDSshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (ADDSshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (ADDSconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (ADDSconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADDSconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDSshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (ADDSshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (ADDSconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (ADDSconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADDSconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDconst_0(v *Value) bool {
	// match: (ADDconst [off1] (MOVWaddr [off2] {sym} ptr))
	// cond:
	// result: (MOVWaddr [off1+off2] {sym} ptr)
	for {
		off1 := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVWaddr)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		return true
	}
	// match: (ADDconst [0] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(c+d))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(c + d))
		return true
	}
	// match: (ADDconst [c] (ADDconst [d] x))
	// cond:
	// result: (ADDconst [int64(int32(c+d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (SUBconst [d] x))
	// cond:
	// result: (ADDconst [int64(int32(c-d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (RSBconst [d] x))
	// cond:
	// result: (RSBconst [int64(int32(c+d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbRSBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDshiftLL_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ADDshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (ADDconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (ADDconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADDconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftLL [c] (SRLconst x [32-c]) x)
	// cond:
	// result: (SRRconst [32-c] x)
	for {
		c := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		if x != v_0.Args[0] {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftLL <typ.UInt16> [8] (BFXU <typ.UInt16> [armBFAuxInt(8, 8)] x) x)
	// cond:
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 {
			break
		}
		if v.AuxInt != 8 {
			break
		}
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbBFXU {
			break
		}
		if v_0.Type != typ.UInt16 {
			break
		}
		if v_0.AuxInt != armBFAuxInt(8, 8) {
			break
		}
		if x != v_0.Args[0] {
			break
		}
		v.reset(OpThumbREV16)
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftLL <typ.UInt16> [8] (SRLconst <typ.UInt16> [24] (SLLconst [16] x)) x)
	// cond:
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 {
			break
		}
		if v.AuxInt != 8 {
			break
		}
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		if v_0.Type != typ.UInt16 {
			break
		}
		if v_0.AuxInt != 24 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbSLLconst {
			break
		}
		if v_0_0.AuxInt != 16 {
			break
		}
		if x != v_0_0.Args[0] {
			break
		}
		v.reset(OpThumbREV16)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (ADDshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (ADDconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (ADDconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADDconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbADDshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (ADDshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (ADDconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (ADDconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbADDconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftRL [c] (SLLconst x [32-c]) x)
	// cond:
	// result: (SRRconst [ c] x)
	for {
		c := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		if x != v_0.Args[0] {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbAND_0(v *Value) bool {
	// match: (AND x (MOVWconst [c]))
	// cond:
	// result: (ANDconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (AND (MOVWconst [c]) x)
	// cond:
	// result: (ANDconst [c] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (AND x (SLLconst [c] y))
	// cond:
	// result: (ANDshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbANDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND (SLLconst [c] y) x)
	// cond:
	// result: (ANDshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbANDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND x (SRLconst [c] y))
	// cond:
	// result: (ANDshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbANDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND (SRLconst [c] y) x)
	// cond:
	// result: (ANDshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbANDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND x (SRAconst [c] y))
	// cond:
	// result: (ANDshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbANDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND (SRAconst [c] y) x)
	// cond:
	// result: (ANDshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbANDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND x x)
	// cond:
	// result: x
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (AND x (MVN y))
	// cond:
	// result: (BIC x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMVN {
			break
		}
		y := v_1.Args[0]
		v.reset(OpThumbBIC)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbAND_10(v *Value) bool {
	// match: (AND (MVN y) x)
	// cond:
	// result: (BIC x y)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMVN {
			break
		}
		y := v_0.Args[0]
		v.reset(OpThumbBIC)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND x (MVNshiftLL y [c]))
	// cond:
	// result: (BICshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMVNshiftLL {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbBICshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND (MVNshiftLL y [c]) x)
	// cond:
	// result: (BICshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMVNshiftLL {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbBICshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND x (MVNshiftRL y [c]))
	// cond:
	// result: (BICshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMVNshiftRL {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbBICshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND (MVNshiftRL y [c]) x)
	// cond:
	// result: (BICshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMVNshiftRL {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbBICshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND x (MVNshiftRA y [c]))
	// cond:
	// result: (BICshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMVNshiftRA {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbBICshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND (MVNshiftRA y [c]) x)
	// cond:
	// result: (BICshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMVNshiftRA {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbBICshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbANDconst_0(v *Value) bool {
	// match: (ANDconst [0] _)
	// cond:
	// result: (MOVWconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDconst [c] x)
	// cond: int32(c)==-1
	// result: x
	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [c&d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = c & d
		return true
	}
	// match: (ANDconst [c] (ANDconst [d] x))
	// cond:
	// result: (ANDconst [c&d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbANDshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (ANDshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (ANDconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (ANDconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbANDconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ANDshiftLL x y:(SLLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpThumbSLLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbANDshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (ANDshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (ANDconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (ANDconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbANDconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (ANDshiftRA x y:(SRAconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpThumbSRAconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbANDshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (ANDshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (ANDconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (ANDconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbANDconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ANDshiftRL x y:(SRLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpThumbSRLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBFX_0(v *Value) bool {
	// match: (BFX [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(d)<<(32-uint32(c&0xff)-uint32(c>>8))>>(32-uint32(c>>8)))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(d) << (32 - uint32(c&0xff) - uint32(c>>8)) >> (32 - uint32(c>>8)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBFXU_0(v *Value) bool {
	// match: (BFXU [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(uint32(d)<<(32-uint32(c&0xff)-uint32(c>>8))>>(32-uint32(c>>8))))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(uint32(d) << (32 - uint32(c&0xff) - uint32(c>>8)) >> (32 - uint32(c>>8))))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBIC_0(v *Value) bool {
	// match: (BIC x (MOVWconst [c]))
	// cond:
	// result: (BICconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbBICconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (BIC x (SLLconst [c] y))
	// cond:
	// result: (BICshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbBICshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (BIC x (SRLconst [c] y))
	// cond:
	// result: (BICshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbBICshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (BIC x (SRAconst [c] y))
	// cond:
	// result: (BICshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbBICshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (BIC x x)
	// cond:
	// result: (MOVWconst [0])
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBICconst_0(v *Value) bool {
	// match: (BICconst [0] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (BICconst [c] _)
	// cond: int32(c)==-1
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (BICconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [d&^c])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = d &^ c
		return true
	}
	// match: (BICconst [c] (BICconst [d] x))
	// cond:
	// result: (BICconst [int64(int32(c|d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbBICconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbBICconst)
		v.AuxInt = int64(int32(c | d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBICshiftLL_0(v *Value) bool {
	// match: (BICshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (BICconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbBICconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (BICshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBICshiftRA_0(v *Value) bool {
	// match: (BICshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (BICconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbBICconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (BICshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbBICshiftRL_0(v *Value) bool {
	// match: (BICshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (BICconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbBICconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (BICshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMN_0(v *Value) bool {
	// match: (CMN x (MOVWconst [c]))
	// cond:
	// result: (CMNconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbCMNconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMN (MOVWconst [c]) x)
	// cond:
	// result: (CMNconst [c] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbCMNconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMN x (SLLconst [c] y))
	// cond:
	// result: (CMNshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbCMNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMN (SLLconst [c] y) x)
	// cond:
	// result: (CMNshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbCMNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMN x (SRLconst [c] y))
	// cond:
	// result: (CMNshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbCMNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMN (SRLconst [c] y) x)
	// cond:
	// result: (CMNshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbCMNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMN x (SRAconst [c] y))
	// cond:
	// result: (CMNshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbCMNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMN (SRAconst [c] y) x)
	// cond:
	// result: (CMNshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbCMNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMN x (RSBconst [0] y))
	// cond:
	// result: (CMP x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbRSBconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpThumbCMP)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMN (RSBconst [0] y) x)
	// cond:
	// result: (CMP x y)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbRSBconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		y := v_0.Args[0]
		v.reset(OpThumbCMP)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMNconst_0(v *Value) bool {
	// match: (CMNconst (MOVWconst [x]) [y])
	// cond: int32(x)==int32(-y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(-y)) {
			break
		}
		v.reset(OpThumbFlagEQ)
		return true
	}
	// match: (CMNconst (MOVWconst [x]) [y])
	// cond: int32(x)<int32(-y) && uint32(x)<uint32(-y)
	// result: (FlagLT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(-y) && uint32(x) < uint32(-y)) {
			break
		}
		v.reset(OpThumbFlagLT_ULT)
		return true
	}
	// match: (CMNconst (MOVWconst [x]) [y])
	// cond: int32(x)<int32(-y) && uint32(x)>uint32(-y)
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(-y) && uint32(x) > uint32(-y)) {
			break
		}
		v.reset(OpThumbFlagLT_UGT)
		return true
	}
	// match: (CMNconst (MOVWconst [x]) [y])
	// cond: int32(x)>int32(-y) && uint32(x)<uint32(-y)
	// result: (FlagGT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(-y) && uint32(x) < uint32(-y)) {
			break
		}
		v.reset(OpThumbFlagGT_ULT)
		return true
	}
	// match: (CMNconst (MOVWconst [x]) [y])
	// cond: int32(x)>int32(-y) && uint32(x)>uint32(-y)
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(-y) && uint32(x) > uint32(-y)) {
			break
		}
		v.reset(OpThumbFlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMNshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (CMNshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (CMNconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (CMNconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbCMNconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMNshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (CMNshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (CMNconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (CMNconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbCMNconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMNshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (CMNshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (CMNconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (CMNconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbCMNconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMOVWHSconst_0(v *Value) bool {
	// match: (CMOVWHSconst _ (FlagEQ) [c])
	// cond:
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWHSconst x (FlagLT_ULT))
	// cond:
	// result: x
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWHSconst _ (FlagLT_UGT) [c])
	// cond:
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWHSconst x (FlagGT_ULT))
	// cond:
	// result: x
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWHSconst _ (FlagGT_UGT) [c])
	// cond:
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWHSconst x (InvertFlags flags) [c])
	// cond:
	// result: (CMOVWLSconst x flags [c])
	for {
		c := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbInvertFlags {
			break
		}
		flags := v_1.Args[0]
		v.reset(OpThumbCMOVWLSconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMOVWLSconst_0(v *Value) bool {
	// match: (CMOVWLSconst _ (FlagEQ) [c])
	// cond:
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWLSconst _ (FlagLT_ULT) [c])
	// cond:
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWLSconst x (FlagLT_UGT))
	// cond:
	// result: x
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWLSconst _ (FlagGT_ULT) [c])
	// cond:
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWLSconst x (FlagGT_UGT))
	// cond:
	// result: x
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWLSconst x (InvertFlags flags) [c])
	// cond:
	// result: (CMOVWHSconst x flags [c])
	for {
		c := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbInvertFlags {
			break
		}
		flags := v_1.Args[0]
		v.reset(OpThumbCMOVWHSconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMP_0(v *Value) bool {
	b := v.Block
	// match: (CMP x (MOVWconst [c]))
	// cond:
	// result: (CMPconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbCMPconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMP (MOVWconst [c]) x)
	// cond:
	// result: (InvertFlags (CMPconst [c] x))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SLLconst [c] y))
	// cond:
	// result: (CMPshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbCMPshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMP (SLLconst [c] y) x)
	// cond:
	// result: (InvertFlags (CMPshiftLL x y [c]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPshiftLL, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SRLconst [c] y))
	// cond:
	// result: (CMPshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbCMPshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMP (SRLconst [c] y) x)
	// cond:
	// result: (InvertFlags (CMPshiftRL x y [c]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRL, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SRAconst [c] y))
	// cond:
	// result: (CMPshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbCMPshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMP (SRAconst [c] y) x)
	// cond:
	// result: (InvertFlags (CMPshiftRA x y [c]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRA, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (RSBconst [0] y))
	// cond:
	// result: (CMN x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbRSBconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpThumbCMN)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPD_0(v *Value) bool {
	// match: (CMPD x (MOVDconst [0]))
	// cond:
	// result: (CMPD0 x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVDconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpThumbCMPD0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPF_0(v *Value) bool {
	// match: (CMPF x (MOVFconst [0]))
	// cond:
	// result: (CMPF0 x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVFconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpThumbCMPF0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPconst_0(v *Value) bool {
	// match: (CMPconst (MOVWconst [x]) [y])
	// cond: int32(x)==int32(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpThumbFlagEQ)
		return true
	}
	// match: (CMPconst (MOVWconst [x]) [y])
	// cond: int32(x)<int32(y) && uint32(x)<uint32(y)
	// result: (FlagLT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpThumbFlagLT_ULT)
		return true
	}
	// match: (CMPconst (MOVWconst [x]) [y])
	// cond: int32(x)<int32(y) && uint32(x)>uint32(y)
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpThumbFlagLT_UGT)
		return true
	}
	// match: (CMPconst (MOVWconst [x]) [y])
	// cond: int32(x)>int32(y) && uint32(x)<uint32(y)
	// result: (FlagGT_ULT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpThumbFlagGT_ULT)
		return true
	}
	// match: (CMPconst (MOVWconst [x]) [y])
	// cond: int32(x)>int32(y) && uint32(x)>uint32(y)
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpThumbFlagGT_UGT)
		return true
	}
	// match: (CMPconst (MOVBUreg _) [c])
	// cond: 0xff < c
	// result: (FlagLT_ULT)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVBUreg {
			break
		}
		if !(0xff < c) {
			break
		}
		v.reset(OpThumbFlagLT_ULT)
		return true
	}
	// match: (CMPconst (MOVHUreg _) [c])
	// cond: 0xffff < c
	// result: (FlagLT_ULT)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVHUreg {
			break
		}
		if !(0xffff < c) {
			break
		}
		v.reset(OpThumbFlagLT_ULT)
		return true
	}
	// match: (CMPconst (ANDconst _ [m]) [n])
	// cond: 0 <= int32(m) && int32(m) < int32(n)
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbANDconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int32(m) && int32(m) < int32(n)) {
			break
		}
		v.reset(OpThumbFlagLT_ULT)
		return true
	}
	// match: (CMPconst (SRLconst _ [c]) [n])
	// cond: 0 <= n && 0 < c && c <= 32 && (1<<uint32(32-c)) <= uint32(n)
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		if !(0 <= n && 0 < c && c <= 32 && (1<<uint32(32-c)) <= uint32(n)) {
			break
		}
		v.reset(OpThumbFlagLT_ULT)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (CMPshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (InvertFlags (CMPconst [c] (SLLconst <x.Type> x [d])))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v1.AuxInt = d
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (CMPconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbCMPconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (CMPshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (InvertFlags (CMPconst [c] (SRAconst <x.Type> x [d])))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v1.AuxInt = d
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (CMPconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbCMPconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbCMPshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (CMPshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (InvertFlags (CMPconst [c] (SRLconst <x.Type> x [d])))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbInvertFlags)
		v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v1.AuxInt = d
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (CMPconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbCMPconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbDIV_0(v *Value) bool {
	// match: (DIV x (MOVWconst [1]))
	// cond:
	// result: x
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (DIV x (MOVWconst [c]))
	// cond: isPowerOfTwo(c)
	// result: (SRAconst [log2(c)] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (DIV (MOVWconst [c]) (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(c)/int32(d))])
	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(c) / int32(d))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbDIVU_0(v *Value) bool {
	// match: (DIVU x (MOVWconst [1]))
	// cond:
	// result: x
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (DIVU x (MOVWconst [c]))
	// cond: isPowerOfTwo(c)
	// result: (SRLconst [log2(c)] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpThumbSRLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (DIVU (MOVWconst [c]) (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(uint32(c)/uint32(d))])
	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(uint32(c) / uint32(d))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbEqual_0(v *Value) bool {
	// match: (Equal (FlagEQ))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (Equal (FlagLT_ULT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (FlagLT_UGT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (FlagGT_ULT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (FlagGT_UGT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (InvertFlags x))
	// cond:
	// result: (Equal x)
	for {
		v_0 := v.Args[0]
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
func rewriteValueThumb_OpThumbGreaterEqual_0(v *Value) bool {
	// match: (GreaterEqual (FlagEQ))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqual (FlagLT_ULT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterEqual (FlagLT_UGT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterEqual (FlagGT_ULT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqual (FlagGT_UGT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqual (InvertFlags x))
	// cond:
	// result: (LessEqual x)
	for {
		v_0 := v.Args[0]
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
func rewriteValueThumb_OpThumbGreaterEqualU_0(v *Value) bool {
	// match: (GreaterEqualU (FlagEQ))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqualU (FlagLT_ULT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterEqualU (FlagLT_UGT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqualU (FlagGT_ULT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterEqualU (FlagGT_UGT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqualU (InvertFlags x))
	// cond:
	// result: (LessEqualU x)
	for {
		v_0 := v.Args[0]
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
func rewriteValueThumb_OpThumbGreaterThan_0(v *Value) bool {
	// match: (GreaterThan (FlagEQ))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThan (FlagLT_ULT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThan (FlagLT_UGT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThan (FlagGT_ULT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterThan (FlagGT_UGT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterThan (InvertFlags x))
	// cond:
	// result: (LessThan x)
	for {
		v_0 := v.Args[0]
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
func rewriteValueThumb_OpThumbGreaterThanU_0(v *Value) bool {
	// match: (GreaterThanU (FlagEQ))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThanU (FlagLT_ULT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThanU (FlagLT_UGT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterThanU (FlagGT_ULT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThanU (FlagGT_UGT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterThanU (InvertFlags x))
	// cond:
	// result: (LessThanU x)
	for {
		v_0 := v.Args[0]
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
func rewriteValueThumb_OpThumbLessEqual_0(v *Value) bool {
	// match: (LessEqual (FlagEQ))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqual (FlagLT_ULT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqual (FlagLT_UGT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqual (FlagGT_ULT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessEqual (FlagGT_UGT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessEqual (InvertFlags x))
	// cond:
	// result: (GreaterEqual x)
	for {
		v_0 := v.Args[0]
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
func rewriteValueThumb_OpThumbLessEqualU_0(v *Value) bool {
	// match: (LessEqualU (FlagEQ))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqualU (FlagLT_ULT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqualU (FlagLT_UGT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessEqualU (FlagGT_ULT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqualU (FlagGT_UGT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessEqualU (InvertFlags x))
	// cond:
	// result: (GreaterEqualU x)
	for {
		v_0 := v.Args[0]
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
func rewriteValueThumb_OpThumbLessThan_0(v *Value) bool {
	// match: (LessThan (FlagEQ))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThan (FlagLT_ULT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessThan (FlagLT_UGT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessThan (FlagGT_ULT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThan (FlagGT_UGT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThan (InvertFlags x))
	// cond:
	// result: (GreaterThan x)
	for {
		v_0 := v.Args[0]
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
func rewriteValueThumb_OpThumbLessThanU_0(v *Value) bool {
	// match: (LessThanU (FlagEQ))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThanU (FlagLT_ULT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessThanU (FlagLT_UGT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThanU (FlagGT_ULT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessThanU (FlagGT_UGT))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThanU (InvertFlags x))
	// cond:
	// result: (GreaterThanU x)
	for {
		v_0 := v.Args[0]
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
func rewriteValueThumb_OpThumbLoadOnce16_0(v *Value) bool {
	// match: (LoadOnce16 [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond:
	// result: (LoadOnce16 [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbLoadOnce16)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce16 [off1] {sym} (SUBconst [off2] ptr) mem)
	// cond:
	// result: (LoadOnce16 [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbLoadOnce16)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce16 [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (LoadOnce16 [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbLoadOnce16)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce16 [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (LoadOnce16idx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbLoadOnce16idx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce16 [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (LoadOnce16shiftLL ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce16shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce16idx_0(v *Value) bool {
	// match: (LoadOnce16idx ptr (MOVWconst [c]) mem)
	// cond:
	// result: (LoadOnce16 [c] ptr mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbLoadOnce16)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce16idx (MOVWconst [c]) ptr mem)
	// cond:
	// result: (LoadOnce16 [c] ptr mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		v.reset(OpThumbLoadOnce16)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce16idx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (LoadOnce16shiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce16shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce32_0(v *Value) bool {
	// match: (LoadOnce32 [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond:
	// result: (LoadOnce32 [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbLoadOnce32)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce32 [off1] {sym} (SUBconst [off2] ptr) mem)
	// cond:
	// result: (LoadOnce32 [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbLoadOnce32)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce32 [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (LoadOnce32idx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbLoadOnce32idx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce32 [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (LoadOnce32shiftLL ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce32shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce32idx_0(v *Value) bool {
	// match: (LoadOnce32idx ptr (MOVWconst [c]) mem)
	// cond:
	// result: (LoadOnce32 [c] ptr mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbLoadOnce32)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce32idx (MOVWconst [c]) ptr mem)
	// cond:
	// result: (LoadOnce32 [c] ptr mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		v.reset(OpThumbLoadOnce32)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce32idx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (LoadOnce32shiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce32shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce8_0(v *Value) bool {
	// match: (LoadOnce8 [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond:
	// result: (LoadOnce8 [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce8 [off1] {sym} (SUBconst [off2] ptr) mem)
	// cond:
	// result: (LoadOnce8 [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce8 [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (LoadOnce8 [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce8 [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (LoadOnce8 [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce8 [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (LoadOnce8idx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbLoadOnce8idx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce8 [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (LoadOnce8shiftLL ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce8shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbLoadOnce8idx_0(v *Value) bool {
	// match: (LoadOnce8idx ptr (MOVWconst [c]) mem)
	// cond:
	// result: (LoadOnce8 [c] ptr mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce8idx (MOVWconst [c]) ptr mem)
	// cond:
	// result: (LoadOnce8 [c] ptr mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		v.reset(OpThumbLoadOnce8)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (LoadOnce8idx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (LoadOnce8shiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbLoadOnce8shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBUload_0(v *Value) bool {
	// match: (MOVBUload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond:
	// result: (MOVBUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off1] {sym} (SUBconst [off2] ptr) mem)
	// cond:
	// result: (MOVBUload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVBUload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBUreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
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
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVBUloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (MOVBUloadshiftLL ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVBUloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVWconst [int64(read8(sym, off))])
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSB {
			break
		}
		if !(symIsRO(sym)) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(read8(sym, off))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBUloadidx_0(v *Value) bool {
	// match: (MOVBUloadidx ptr idx (MOVBstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVBUreg x)
	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbMOVBstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVBUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUloadidx ptr (MOVWconst [c]) mem)
	// cond:
	// result: (MOVBUload [c] ptr mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbMOVBUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUloadidx (MOVWconst [c]) ptr mem)
	// cond:
	// result: (MOVBUload [c] ptr mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		v.reset(OpThumbMOVBUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUloadidx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (MOVBUloadshiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBUloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUloadidx (SLLconst idx [c]) ptr mem)
	// cond: c <= 3
	// result: (MOVBUloadshiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBUloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBUreg_0(v *Value) bool {
	// match: (MOVBUreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (ANDconst [c] x))
	// cond:
	// result: (ANDconst [c&0xff] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbANDconst)
		v.AuxInt = c & 0xff
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg x:(MOVBUreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVBUreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(uint8(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBload_0(v *Value) bool {
	// match: (MOVBload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond:
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym} (SUBconst [off2] ptr) mem)
	// cond:
	// result: (MOVBload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVBload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
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
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVBloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (MOVBloadshiftLL ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVBloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBloadidx_0(v *Value) bool {
	// match: (MOVBloadidx ptr idx (MOVBstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVBreg x)
	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbMOVBstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBloadidx ptr (MOVWconst [c]) mem)
	// cond:
	// result: (MOVBload [c] ptr mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbMOVBload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBloadidx (MOVWconst [c]) ptr mem)
	// cond:
	// result: (MOVBload [c] ptr mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		v.reset(OpThumbMOVBload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBloadidx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (MOVBloadshiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBloadidx (SLLconst idx [c]) ptr mem)
	// cond: c <= 3
	// result: (MOVBloadshiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBreg_0(v *Value) bool {
	// match: (MOVBreg x:(MOVBload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (ANDconst [c] x))
	// cond: c & 0x80 == 0
	// result: (ANDconst [c&0x7f] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x80 == 0) {
			break
		}
		v.reset(OpThumbANDconst)
		v.AuxInt = c & 0x7f
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVBreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(int8(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBstore_0(v *Value) bool {
	// match: (MOVBstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond:
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// cond:
	// result: (MOVBstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVBreg {
			break
		}
		x := v_1.Args[0]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBUreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVBUreg {
			break
		}
		x := v_1.Args[0]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVHreg {
			break
		}
		x := v_1.Args[0]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHUreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVHUreg {
			break
		}
		x := v_1.Args[0]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (MOVBstoreidx ptr idx val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [0] {sym} (ADDshiftLL ptr idx [c]) val mem)
	// cond: sym == nil && c <= 3
	// result: (MOVBstoreshiftLL ptr idx [c] val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVBstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVBstoreidx_0(v *Value) bool {
	// match: (MOVBstoreidx ptr (MOVWconst [c]) val mem)
	// cond:
	// result: (MOVBstore [c] ptr val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx (MOVWconst [c]) ptr val mem)
	// cond:
	// result: (MOVBstore [c] ptr val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		val := v.Args[2]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (MOVBstoreshiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (MOVBstoreshiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVBstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVDload_0(v *Value) bool {
	// match: (MOVDload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond:
	// result: (MOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym} (SUBconst [off2] ptr) mem)
	// cond:
	// result: (MOVDload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVDload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off] {sym} ptr (MOVDstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVDstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVDstore_0(v *Value) bool {
	// match: (MOVDstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond:
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// cond:
	// result: (MOVDstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbMOVDstore)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVFload_0(v *Value) bool {
	// match: (MOVFload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond:
	// result: (MOVFload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVFload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFload [off1] {sym} (SUBconst [off2] ptr) mem)
	// cond:
	// result: (MOVFload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVFload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVFload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFload [off] {sym} ptr (MOVFstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVFstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVFstore_0(v *Value) bool {
	// match: (MOVFstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond:
	// result: (MOVFstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// cond:
	// result: (MOVFstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbMOVFstore)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHUload_0(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	// match: (MOVHUload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond:
	// result: (MOVHUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off1] {sym} (SUBconst [off2] ptr) mem)
	// cond:
	// result: (MOVHUload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVHUload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVHUreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVHstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
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
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVHUloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (MOVHUloadshiftLL ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVHUloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVWconst [int64(read16(sym, off, config.BigEndian))])
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSB {
			break
		}
		if !(symIsRO(sym)) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(read16(sym, off, config.BigEndian))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHUloadidx_0(v *Value) bool {
	// match: (MOVHUloadidx ptr idx (MOVHstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVHUreg x)
	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbMOVHstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVHUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUloadidx ptr (MOVWconst [c]) mem)
	// cond:
	// result: (MOVHUload [c] ptr mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbMOVHUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUloadidx (MOVWconst [c]) ptr mem)
	// cond:
	// result: (MOVHUload [c] ptr mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		v.reset(OpThumbMOVHUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUloadidx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (MOVHUloadshiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHUloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUloadidx (SLLconst idx [c]) ptr mem)
	// cond: c <= 3
	// result: (MOVHUloadshiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHUloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHUreg_0(v *Value) bool {
	// match: (MOVHUreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (ANDconst [c] x))
	// cond:
	// result: (ANDconst [c&0xffff] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbANDconst)
		v.AuxInt = c & 0xffff
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVBUreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVBUreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVHUreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(uint16(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHload_0(v *Value) bool {
	// match: (MOVHload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond:
	// result: (MOVHload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym} (SUBconst [off2] ptr) mem)
	// cond:
	// result: (MOVHload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVHload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVHreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVHstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
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
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVHloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (MOVHloadshiftLL ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVHloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHloadidx_0(v *Value) bool {
	// match: (MOVHloadidx ptr idx (MOVHstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVHreg x)
	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbMOVHstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpThumbMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHloadidx ptr (MOVWconst [c]) mem)
	// cond:
	// result: (MOVHload [c] ptr mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbMOVHload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHloadidx (MOVWconst [c]) ptr mem)
	// cond:
	// result: (MOVHload [c] ptr mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		v.reset(OpThumbMOVHload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHloadidx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (MOVHloadshiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHloadidx (SLLconst idx [c]) ptr mem)
	// cond: c <= 3
	// result: (MOVHloadshiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHreg_0(v *Value) bool {
	// match: (MOVHreg x:(MOVBload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (ANDconst [c] x))
	// cond: c & 0x8000 == 0
	// result: (ANDconst [c&0x7fff] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x8000 == 0) {
			break
		}
		v.reset(OpThumbANDconst)
		v.AuxInt = c & 0x7fff
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVBreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVBUreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpThumbMOVHreg {
			break
		}
		v.reset(OpThumbMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(int16(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHstore_0(v *Value) bool {
	// match: (MOVHstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond:
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// cond:
	// result: (MOVHstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbMOVHstore)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHreg x) mem)
	// cond:
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVHreg {
			break
		}
		x := v_1.Args[0]
		v.reset(OpThumbMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHUreg x) mem)
	// cond:
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVHUreg {
			break
		}
		x := v_1.Args[0]
		v.reset(OpThumbMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (MOVHstoreidx ptr idx val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [0] {sym} (ADDshiftLL ptr idx [c]) val mem)
	// cond: sym == nil && c <= 3
	// result: (MOVHstoreshiftLL ptr idx [c] val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVHstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVHstoreidx_0(v *Value) bool {
	// match: (MOVHstoreidx ptr (MOVWconst [c]) val mem)
	// cond:
	// result: (MOVHstore [c] ptr val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		v.reset(OpThumbMOVHstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreidx (MOVWconst [c]) ptr val mem)
	// cond:
	// result: (MOVHstore [c] ptr val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		val := v.Args[2]
		v.reset(OpThumbMOVHstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreidx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (MOVHstoreshiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreidx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (MOVHstoreshiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVHstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWload_0(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	// match: (MOVWload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond:
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym} (SUBconst [off2] ptr) mem)
	// cond:
	// result: (MOVWload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVWload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off] {sym} ptr (MOVWstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVWloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVWloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil && c <= 3
	// result: (MOVWloadshiftLL ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVWloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVWconst [int64(int32(read32(sym, off, config.BigEndian)))])
	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSB {
			break
		}
		if !(symIsRO(sym)) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(read32(sym, off, config.BigEndian)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWloadidx_0(v *Value) bool {
	// match: (MOVWloadidx ptr idx (MOVWstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: x
	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbMOVWstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWloadidx ptr (MOVWconst [c]) mem)
	// cond:
	// result: (MOVWload [c] ptr mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbMOVWload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx (MOVWconst [c]) ptr mem)
	// cond:
	// result: (MOVWload [c] ptr mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		v.reset(OpThumbMOVWload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx ptr (SLLconst idx [c]) mem)
	// cond: c <= 3
	// result: (MOVWloadshiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVWloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx (SLLconst idx [c]) ptr mem)
	// cond: c <= 3
	// result: (MOVWloadshiftLL ptr idx [c] mem)
	for {
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVWloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWloadshiftLL_0(v *Value) bool {
	// match: (MOVWloadshiftLL ptr idx [c] (MOVWstoreshiftLL ptr2 idx [d] x _))
	// cond: c==d && isSamePtr(ptr, ptr2)
	// result: x
	for {
		c := v.AuxInt
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbMOVWstoreshiftLL {
			break
		}
		d := v_2.AuxInt
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWloadshiftLL ptr (MOVWconst [c]) [d] mem)
	// cond:
	// result: (MOVWload [int64(uint32(c)<<uint64(d))] ptr mem)
	for {
		d := v.AuxInt
		mem := v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbMOVWload)
		v.AuxInt = int64(uint32(c) << uint64(d))
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWreg_0(v *Value) bool {
	// match: (MOVWreg x)
	// cond: x.Uses == 1
	// result: (MOVWnop x)
	for {
		x := v.Args[0]
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpThumbMOVWnop)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWstore_0(v *Value) bool {
	// match: (MOVWstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond:
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// cond:
	// result: (MOVWstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbMOVWstore)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (MOVWstoreidx ptr idx val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbMOVWstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [0] {sym} (ADDshiftLL ptr idx [c]) val mem)
	// cond: sym == nil && c <= 3
	// result: (MOVWstoreshiftLL ptr idx [c] val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbMOVWstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWstoreidx_0(v *Value) bool {
	// match: (MOVWstoreidx ptr (MOVWconst [c]) val mem)
	// cond:
	// result: (MOVWstore [c] ptr val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		v.reset(OpThumbMOVWstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx (MOVWconst [c]) ptr val mem)
	// cond:
	// result: (MOVWstore [c] ptr val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		val := v.Args[2]
		v.reset(OpThumbMOVWstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (MOVWstoreshiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVWstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (MOVWstoreshiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbMOVWstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMOVWstoreshiftLL_0(v *Value) bool {
	// match: (MOVWstoreshiftLL ptr (MOVWconst [c]) [d] val mem)
	// cond:
	// result: (MOVWstore [int64(uint32(c)<<uint64(d))] ptr val mem)
	for {
		d := v.AuxInt
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		v.reset(OpThumbMOVWstore)
		v.AuxInt = int64(uint32(c) << uint64(d))
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMUL_0(v *Value) bool {
	// match: (MUL x (MOVWconst [c]))
	// cond: int32(c) == -1
	// result: (RSBconst [0] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpThumbRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
	// match: (MUL (MOVWconst [c]) x)
	// cond: int32(c) == -1
	// result: (RSBconst [0] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpThumbRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
	// match: (MUL _ (MOVWconst [0]))
	// cond:
	// result: (MOVWconst [0])
	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (MUL (MOVWconst [0]) _)
	// cond:
	// result: (MOVWconst [0])
	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (MUL x (MOVWconst [1]))
	// cond:
	// result: x
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MUL (MOVWconst [1]) x)
	// cond:
	// result: x
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: isPowerOfTwo(c)
	// result: (SLLconst [log2(c)] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (MUL (MOVWconst [c]) x)
	// cond: isPowerOfTwo(c)
	// result: (SLLconst [log2(c)] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (ADDshiftLL x x [log2(c-1)])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpThumbADDshiftLL)
		v.AuxInt = log2(c - 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (MUL (MOVWconst [c]) x)
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (ADDshiftLL x x [log2(c-1)])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpThumbADDshiftLL)
		v.AuxInt = log2(c - 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMUL_10(v *Value) bool {
	b := v.Block
	// match: (MUL x (MOVWconst [c]))
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (RSBshiftLL x x [log2(c+1)])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpThumbRSBshiftLL)
		v.AuxInt = log2(c + 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (MUL (MOVWconst [c]) x)
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (RSBshiftLL x x [log2(c+1)])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpThumbRSBshiftLL)
		v.AuxInt = log2(c + 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (SLLconst [log2(c/3)] (ADDshiftLL <x.Type> x x [1]))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = 1
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MUL (MOVWconst [c]) x)
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (SLLconst [log2(c/3)] (ADDshiftLL <x.Type> x x [1]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = 1
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (SLLconst [log2(c/5)] (ADDshiftLL <x.Type> x x [2]))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MUL (MOVWconst [c]) x)
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (SLLconst [log2(c/5)] (ADDshiftLL <x.Type> x x [2]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (SLLconst [log2(c/7)] (RSBshiftLL <x.Type> x x [3]))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MUL (MOVWconst [c]) x)
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (SLLconst [log2(c/7)] (RSBshiftLL <x.Type> x x [3]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (SLLconst [log2(c/9)] (ADDshiftLL <x.Type> x x [3]))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (MUL (MOVWconst [c]) x)
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (SLLconst [log2(c/9)] (ADDshiftLL <x.Type> x x [3]))
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbSLLconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMUL_20(v *Value) bool {
	// match: (MUL (MOVWconst [c]) (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(c*d))])
	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(c * d))
		return true
	}
	// match: (MUL (MOVWconst [d]) (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(int32(c*d))])
	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(c * d))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMULA_0(v *Value) bool {
	b := v.Block
	// match: (MULA x (MOVWconst [c]) a)
	// cond: int32(c) == -1
	// result: (SUB a x)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpThumbSUB)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}
	// match: (MULA _ (MOVWconst [0]) a)
	// cond:
	// result: a
	for {
		a := v.Args[2]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [1]) a)
	// cond:
	// result: (ADD x a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpThumbADD)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c)
	// result: (ADD (SLLconst <x.Type> [log2(c)] x) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (ADD (ADDshiftLL <x.Type> x x [log2(c-1)]) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (ADD (RSBshiftLL <x.Type> x x [log2(c+1)]) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/3)] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/5)] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/7)] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/9)] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMULA_10(v *Value) bool {
	b := v.Block
	// match: (MULA (MOVWconst [c]) x a)
	// cond: int32(c) == -1
	// result: (SUB a x)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpThumbSUB)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}
	// match: (MULA (MOVWconst [0]) _ a)
	// cond:
	// result: a
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [1]) x a)
	// cond:
	// result: (ADD x a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		v.reset(OpThumbADD)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c)
	// result: (ADD (SLLconst <x.Type> [log2(c)] x) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (ADD (ADDshiftLL <x.Type> x x [log2(c-1)]) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (ADD (RSBshiftLL <x.Type> x x [log2(c+1)]) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/3)] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/5)] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/7)] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/9)] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbADD)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMULA_20(v *Value) bool {
	// match: (MULA (MOVWconst [c]) (MOVWconst [d]) a)
	// cond:
	// result: (ADDconst [int64(int32(c*d))] a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpThumbADDconst)
		v.AuxInt = int64(int32(c * d))
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMULD_0(v *Value) bool {
	// match: (MULD (NEGD x) y)
	// cond:
	// result: (NMULD x y)
	for {
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbNEGD {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbNMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (MULD y (NEGD x))
	// cond:
	// result: (NMULD x y)
	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbNEGD {
			break
		}
		x := v_1.Args[0]
		v.reset(OpThumbNMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMULF_0(v *Value) bool {
	// match: (MULF (NEGF x) y)
	// cond:
	// result: (NMULF x y)
	for {
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbNEGF {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbNMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (MULF y (NEGF x))
	// cond:
	// result: (NMULF x y)
	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbNEGF {
			break
		}
		x := v_1.Args[0]
		v.reset(OpThumbNMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMULS_0(v *Value) bool {
	b := v.Block
	// match: (MULS x (MOVWconst [c]) a)
	// cond: int32(c) == -1
	// result: (ADD a x)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpThumbADD)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}
	// match: (MULS _ (MOVWconst [0]) a)
	// cond:
	// result: a
	for {
		a := v.Args[2]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [1]) a)
	// cond:
	// result: (RSB x a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpThumbRSB)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c)
	// result: (RSB (SLLconst <x.Type> [log2(c)] x) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (RSB (ADDshiftLL <x.Type> x x [log2(c-1)]) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (RSB (RSBshiftLL <x.Type> x x [log2(c+1)]) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/3)] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/5)] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/7)] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/9)] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		a := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMULS_10(v *Value) bool {
	b := v.Block
	// match: (MULS (MOVWconst [c]) x a)
	// cond: int32(c) == -1
	// result: (ADD a x)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpThumbADD)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}
	// match: (MULS (MOVWconst [0]) _ a)
	// cond:
	// result: a
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [1]) x a)
	// cond:
	// result: (RSB x a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		v.reset(OpThumbRSB)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c)
	// result: (RSB (SLLconst <x.Type> [log2(c)] x) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (RSB (ADDshiftLL <x.Type> x x [log2(c-1)]) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (RSB (RSBshiftLL <x.Type> x x [log2(c+1)]) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/3)] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/5)] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/7)] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpThumbRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/9)] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpThumbRSB)
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpThumbADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMULS_20(v *Value) bool {
	// match: (MULS (MOVWconst [c]) (MOVWconst [d]) a)
	// cond:
	// result: (SUBconst [int64(int32(c*d))] a)
	for {
		a := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpThumbSUBconst)
		v.AuxInt = int64(int32(c * d))
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMVN_0(v *Value) bool {
	// match: (MVN (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [^c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = ^c
		return true
	}
	// match: (MVN (SLLconst [c] x))
	// cond:
	// result: (MVNshiftLL x [c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbMVNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MVN (SRLconst [c] x))
	// cond:
	// result: (MVNshiftRL x [c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbMVNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MVN (SRAconst [c] x))
	// cond:
	// result: (MVNshiftRA x [c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbMVNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMVNshiftLL_0(v *Value) bool {
	// match: (MVNshiftLL (MOVWconst [c]) [d])
	// cond:
	// result: (MOVWconst [^int64(uint32(c)<<uint64(d))])
	for {
		d := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = ^int64(uint32(c) << uint64(d))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMVNshiftRA_0(v *Value) bool {
	// match: (MVNshiftRA (MOVWconst [c]) [d])
	// cond:
	// result: (MOVWconst [^int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = ^int64(int32(c) >> uint64(d))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbMVNshiftRL_0(v *Value) bool {
	// match: (MVNshiftRL (MOVWconst [c]) [d])
	// cond:
	// result: (MOVWconst [^int64(uint32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = ^int64(uint32(c) >> uint64(d))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbNEGD_0(v *Value) bool {
	// match: (NEGD (MULD x y))
	// cond:
	// result: (NMULD x y)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMULD {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpThumbNMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbNEGF_0(v *Value) bool {
	// match: (NEGF (MULF x y))
	// cond:
	// result: (NMULF x y)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMULF {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpThumbNMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbNMULD_0(v *Value) bool {
	// match: (NMULD (NEGD x) y)
	// cond:
	// result: (MULD x y)
	for {
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbNEGD {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (NMULD y (NEGD x))
	// cond:
	// result: (MULD x y)
	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbNEGD {
			break
		}
		x := v_1.Args[0]
		v.reset(OpThumbMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbNMULF_0(v *Value) bool {
	// match: (NMULF (NEGF x) y)
	// cond:
	// result: (MULF x y)
	for {
		y := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbNEGF {
			break
		}
		x := v_0.Args[0]
		v.reset(OpThumbMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (NMULF y (NEGF x))
	// cond:
	// result: (MULF x y)
	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbNEGF {
			break
		}
		x := v_1.Args[0]
		v.reset(OpThumbMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbNotEqual_0(v *Value) bool {
	// match: (NotEqual (FlagEQ))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (NotEqual (FlagLT_ULT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (FlagLT_UGT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (FlagGT_ULT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (FlagGT_UGT))
	// cond:
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (InvertFlags x))
	// cond:
	// result: (NotEqual x)
	for {
		v_0 := v.Args[0]
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
func rewriteValueThumb_OpThumbOR_0(v *Value) bool {
	// match: (OR x (MOVWconst [c]))
	// cond:
	// result: (ORconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (OR (MOVWconst [c]) x)
	// cond:
	// result: (ORconst [c] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (OR x (SLLconst [c] y))
	// cond:
	// result: (ORshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (OR (SLLconst [c] y) x)
	// cond:
	// result: (ORshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (OR x (SRLconst [c] y))
	// cond:
	// result: (ORshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (OR (SRLconst [c] y) x)
	// cond:
	// result: (ORshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (OR x (SRAconst [c] y))
	// cond:
	// result: (ORshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (OR (SRAconst [c] y) x)
	// cond:
	// result: (ORshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (OR x x)
	// cond:
	// result: x
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORconst_0(v *Value) bool {
	// match: (ORconst [0] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ORconst [c] _)
	// cond: int32(c)==-1
	// result: (MOVWconst [-1])
	for {
		c := v.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [c|d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = c | d
		return true
	}
	// match: (ORconst [c] (ORconst [d] x))
	// cond:
	// result: (ORconst [c|d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORshiftLL_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (ORconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (ORconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbORconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ORshiftLL [c] (SRLconst x [32-c]) x)
	// cond:
	// result: (SRRconst [32-c] x)
	for {
		c := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		if x != v_0.Args[0] {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}
	// match: (ORshiftLL <typ.UInt16> [8] (BFXU <typ.UInt16> [armBFAuxInt(8, 8)] x) x)
	// cond:
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 {
			break
		}
		if v.AuxInt != 8 {
			break
		}
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbBFXU {
			break
		}
		if v_0.Type != typ.UInt16 {
			break
		}
		if v_0.AuxInt != armBFAuxInt(8, 8) {
			break
		}
		if x != v_0.Args[0] {
			break
		}
		v.reset(OpThumbREV16)
		v.AddArg(x)
		return true
	}
	// match: (ORshiftLL <typ.UInt16> [8] (SRLconst <typ.UInt16> [24] (SLLconst [16] x)) x)
	// cond:
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 {
			break
		}
		if v.AuxInt != 8 {
			break
		}
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		if v_0.Type != typ.UInt16 {
			break
		}
		if v_0.AuxInt != 24 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbSLLconst {
			break
		}
		if v_0_0.AuxInt != 16 {
			break
		}
		if x != v_0_0.Args[0] {
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
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpThumbSLLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (ORshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (ORconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (ORconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbORconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (ORshiftRA x y:(SRAconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpThumbSRAconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbORshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (ORshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (ORconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (ORconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbORconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ORshiftRL [c] (SLLconst x [32-c]) x)
	// cond:
	// result: (SRRconst [ c] x)
	for {
		c := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		if x != v_0.Args[0] {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ORshiftRL x y:(SRLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpThumbSRLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSB_0(v *Value) bool {
	// match: (RSB (MOVWconst [c]) x)
	// cond:
	// result: (SUBconst [c] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbSUBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (RSB x (MOVWconst [c]))
	// cond:
	// result: (RSBconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbRSBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (RSB x (SLLconst [c] y))
	// cond:
	// result: (RSBshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbRSBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB (SLLconst [c] y) x)
	// cond:
	// result: (SUBshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbSUBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB x (SRLconst [c] y))
	// cond:
	// result: (RSBshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbRSBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB (SRLconst [c] y) x)
	// cond:
	// result: (SUBshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbSUBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB x (SRAconst [c] y))
	// cond:
	// result: (RSBshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbRSBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB (SRAconst [c] y) x)
	// cond:
	// result: (SUBshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbSUBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB x x)
	// cond:
	// result: (MOVWconst [0])
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (RSB (MUL x y) a)
	// cond:
	// result: (MULS x y a)
	for {
		a := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMUL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpThumbMULS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBSshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (RSBSshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (SUBSconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (RSBSconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbRSBSconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBSshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (RSBSshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (SUBSconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (RSBSconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbRSBSconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBSshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (RSBSshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (SUBSconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (RSBSconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbRSBSconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBconst_0(v *Value) bool {
	// match: (RSBconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(c-d))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(c - d))
		return true
	}
	// match: (RSBconst [c] (RSBconst [d] x))
	// cond:
	// result: (ADDconst [int64(int32(c-d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbRSBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}
	// match: (RSBconst [c] (ADDconst [d] x))
	// cond:
	// result: (RSBconst [int64(int32(c-d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}
	// match: (RSBconst [c] (SUBconst [d] x))
	// cond:
	// result: (RSBconst [int64(int32(c+d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (RSBshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (SUBconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (RSBconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbRSBconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (RSBshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (RSBshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (SUBconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (RSBconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbRSBconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (RSBshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbRSBshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (RSBshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (SUBconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (RSBconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbRSBconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (RSBshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSBC_0(v *Value) bool {
	// match: (SBC x (MOVWconst [c]) flags)
	// cond:
	// result: (SBCconst [c] x flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSBCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (SBC x (SLLconst [c] y) flags)
	// cond:
	// result: (SBCshiftLL x y [c] flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbSBCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (SBC x (SRLconst [c] y) flags)
	// cond:
	// result: (SBCshiftRL x y [c] flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbSBCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (SBC x (SRAconst [c] y) flags)
	// cond:
	// result: (SBCshiftRA x y [c] flags)
	for {
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbSBCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSBCconst_0(v *Value) bool {
	// match: (SBCconst [c] (ADDconst [d] x) flags)
	// cond:
	// result: (SBCconst [int64(int32(c-d))] x flags)
	for {
		c := v.AuxInt
		flags := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbSBCconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (SBCconst [c] (SUBconst [d] x) flags)
	// cond:
	// result: (SBCconst [int64(int32(c+d))] x flags)
	for {
		c := v.AuxInt
		flags := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbSBCconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSBCshiftLL_0(v *Value) bool {
	// match: (SBCshiftLL x (MOVWconst [c]) [d] flags)
	// cond:
	// result: (SBCconst x [int64(int32(uint32(c)<<uint64(d)))] flags)
	for {
		d := v.AuxInt
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSBCconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSBCshiftRA_0(v *Value) bool {
	// match: (SBCshiftRA x (MOVWconst [c]) [d] flags)
	// cond:
	// result: (SBCconst x [int64(int32(c)>>uint64(d))] flags)
	for {
		d := v.AuxInt
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSBCconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSBCshiftRL_0(v *Value) bool {
	// match: (SBCshiftRL x (MOVWconst [c]) [d] flags)
	// cond:
	// result: (SBCconst x [int64(int32(uint32(c)>>uint64(d)))] flags)
	for {
		d := v.AuxInt
		flags := v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSBCconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSLL_0(v *Value) bool {
	// match: (SLL x (MOVWconst [c]))
	// cond:
	// result: (SLLconst x [c&31])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSLLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSLLconst_0(v *Value) bool {
	// match: (SLLconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(uint32(d)<<uint64(c)))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(uint32(d) << uint64(c)))
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSRA_0(v *Value) bool {
	// match: (SRA x (MOVWconst [c]))
	// cond:
	// result: (SRAconst x [c&31])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSRAconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSRAcond_0(v *Value) bool {
	// match: (SRAcond x _ (FlagEQ))
	// cond:
	// result: (SRAconst x [31])
	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbFlagEQ {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	// match: (SRAcond x y (FlagLT_ULT))
	// cond:
	// result: (SRA x y)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbFlagLT_ULT {
			break
		}
		v.reset(OpThumbSRA)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAcond x _ (FlagLT_UGT))
	// cond:
	// result: (SRAconst x [31])
	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbFlagLT_UGT {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	// match: (SRAcond x y (FlagGT_ULT))
	// cond:
	// result: (SRA x y)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbFlagGT_ULT {
			break
		}
		v.reset(OpThumbSRA)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAcond x _ (FlagGT_UGT))
	// cond:
	// result: (SRAconst x [31])
	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpThumbFlagGT_UGT {
			break
		}
		v.reset(OpThumbSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSRAconst_0(v *Value) bool {
	// match: (SRAconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(d)>>uint64(c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(d) >> uint64(c))
		return true
	}
	// match: (SRAconst (SLLconst x [c]) [d])
	// cond: uint64(d)>=uint64(c) && uint64(d)<=31
	// result: (BFX [(d-c)|(32-d)<<8] x)
	for {
		d := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(uint64(d) >= uint64(c) && uint64(d) <= 31) {
			break
		}
		v.reset(OpThumbBFX)
		v.AuxInt = (d - c) | (32-d)<<8
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSRL_0(v *Value) bool {
	// match: (SRL x (MOVWconst [c]))
	// cond:
	// result: (SRLconst x [c&31])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSRLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSRLconst_0(v *Value) bool {
	// match: (SRLconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(uint32(d)>>uint64(c)))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(uint32(d) >> uint64(c)))
		return true
	}
	// match: (SRLconst (SLLconst x [c]) [d])
	// cond: uint64(d)>=uint64(c) && uint64(d)<=31
	// result: (BFXU [(d-c)|(32-d)<<8] x)
	for {
		d := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(uint64(d) >= uint64(c) && uint64(d) <= 31) {
			break
		}
		v.reset(OpThumbBFXU)
		v.AuxInt = (d - c) | (32-d)<<8
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUB_0(v *Value) bool {
	// match: (SUB (MOVWconst [c]) x)
	// cond:
	// result: (RSBconst [c] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbRSBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUB x (MOVWconst [c]))
	// cond:
	// result: (SUBconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSUBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUB x (SLLconst [c] y))
	// cond:
	// result: (SUBshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbSUBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB (SLLconst [c] y) x)
	// cond:
	// result: (RSBshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbRSBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB x (SRLconst [c] y))
	// cond:
	// result: (SUBshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbSUBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB (SRLconst [c] y) x)
	// cond:
	// result: (RSBshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbRSBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB x (SRAconst [c] y))
	// cond:
	// result: (SUBshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbSUBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB (SRAconst [c] y) x)
	// cond:
	// result: (RSBshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbRSBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB x x)
	// cond:
	// result: (MOVWconst [0])
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SUB a (MUL x y))
	// cond:
	// result: (MULS x y a)
	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMUL {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		v.reset(OpThumbMULS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBD_0(v *Value) bool {
	// match: (SUBD a (MULD x y))
	// cond: a.Uses == 1
	// result: (MULSD a x y)
	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMULD {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULSD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBD a (NMULD x y))
	// cond: a.Uses == 1
	// result: (MULAD a x y)
	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbNMULD {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULAD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBF_0(v *Value) bool {
	// match: (SUBF a (MULF x y))
	// cond: a.Uses == 1
	// result: (MULSF a x y)
	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMULF {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULSF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBF a (NMULF x y))
	// cond: a.Uses == 1
	// result: (MULAF a x y)
	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbNMULF {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1) {
			break
		}
		v.reset(OpThumbMULAF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBS_0(v *Value) bool {
	// match: (SUBS x (MOVWconst [c]))
	// cond:
	// result: (SUBSconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSUBSconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUBS x (SLLconst [c] y))
	// cond:
	// result: (SUBSshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbSUBSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS (SLLconst [c] y) x)
	// cond:
	// result: (RSBSshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbRSBSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS x (SRLconst [c] y))
	// cond:
	// result: (SUBSshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbSUBSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS (SRLconst [c] y) x)
	// cond:
	// result: (RSBSshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbRSBSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS x (SRAconst [c] y))
	// cond:
	// result: (SUBSshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbSUBSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS (SRAconst [c] y) x)
	// cond:
	// result: (RSBSshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbRSBSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBSshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (SUBSshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (RSBSconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (SUBSconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSUBSconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBSshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (SUBSshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (RSBSconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (SUBSconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSUBSconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBSshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (SUBSshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (RSBSconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (SUBSconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSUBSconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBconst_0(v *Value) bool {
	// match: (SUBconst [off1] (MOVWaddr [off2] {sym} ptr))
	// cond:
	// result: (MOVWaddr [off2-off1] {sym} ptr)
	for {
		off1 := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpThumbMOVWaddr)
		v.AuxInt = off2 - off1
		v.Aux = sym
		v.AddArg(ptr)
		return true
	}
	// match: (SUBconst [0] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(d-c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = int64(int32(d - c))
		return true
	}
	// match: (SUBconst [c] (SUBconst [d] x))
	// cond:
	// result: (ADDconst [int64(int32(-c-d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = int64(int32(-c - d))
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (ADDconst [d] x))
	// cond:
	// result: (ADDconst [int64(int32(-c+d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbADDconst)
		v.AuxInt = int64(int32(-c + d))
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (RSBconst [d] x))
	// cond:
	// result: (RSBconst [int64(int32(-c+d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbRSBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbRSBconst)
		v.AuxInt = int64(int32(-c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (SUBshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (RSBconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (SUBconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSUBconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (SUBshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (SUBshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (RSBconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (SUBconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSUBconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (SUBshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbSUBshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (SUBshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (RSBconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (SUBconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbSUBconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (SUBshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce16_0(v *Value) bool {
	// match: (StoreOnce16 [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond:
	// result: (StoreOnce16 [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbStoreOnce16)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce16 [off1] {sym} (SUBconst [off2] ptr) val mem)
	// cond:
	// result: (StoreOnce16 [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbStoreOnce16)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce16 [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (StoreOnce16 [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbStoreOnce16)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce16 [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (StoreOnce16idx ptr idx val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbStoreOnce16idx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce16 [0] {sym} (ADDshiftLL ptr idx [c]) val mem)
	// cond: sym == nil && c <= 3
	// result: (StoreOnce16shiftLL ptr idx [c] val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce16shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce16idx_0(v *Value) bool {
	// match: (StoreOnce16idx ptr (MOVWconst [c]) val mem)
	// cond:
	// result: (StoreOnce16 [c] ptr val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		v.reset(OpThumbStoreOnce16)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce16idx (MOVWconst [c]) ptr val mem)
	// cond:
	// result: (StoreOnce16 [c] ptr val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		val := v.Args[2]
		v.reset(OpThumbStoreOnce16)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce16idx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (StoreOnce16shiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce16shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce16idx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (StoreOnce16shiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce16shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce32_0(v *Value) bool {
	// match: (StoreOnce32 [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond:
	// result: (StoreOnce32 [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbStoreOnce32)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce32 [off1] {sym} (SUBconst [off2] ptr) val mem)
	// cond:
	// result: (StoreOnce32 [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbStoreOnce32)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce32 [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (StoreOnce32 [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbStoreOnce32)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce32 [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (StoreOnce32idx ptr idx val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbStoreOnce32idx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce32 [0] {sym} (ADDshiftLL ptr idx [c]) val mem)
	// cond: sym == nil && c <= 3
	// result: (StoreOnce32shiftLL ptr idx [c] val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce32shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce32idx_0(v *Value) bool {
	// match: (StoreOnce32idx ptr (MOVWconst [c]) val mem)
	// cond:
	// result: (StoreOnce32 [c] ptr val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		v.reset(OpThumbStoreOnce32)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce32idx (MOVWconst [c]) ptr val mem)
	// cond:
	// result: (StoreOnce32 [c] ptr val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		val := v.Args[2]
		v.reset(OpThumbStoreOnce32)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce32idx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (StoreOnce32shiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce32shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce32idx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (StoreOnce32shiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce32shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce8_0(v *Value) bool {
	// match: (StoreOnce8 [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond:
	// result: (StoreOnce8 [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbStoreOnce8)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce8 [off1] {sym} (SUBconst [off2] ptr) val mem)
	// cond:
	// result: (StoreOnce8 [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		v.reset(OpThumbStoreOnce8)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce8 [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (StoreOnce8 [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpThumbStoreOnce8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce8 [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (StoreOnce8idx ptr idx val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil) {
			break
		}
		v.reset(OpThumbStoreOnce8idx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce8 [0] {sym} (ADDshiftLL ptr idx [c]) val mem)
	// cond: sym == nil && c <= 3
	// result: (StoreOnce8shiftLL ptr idx [c] val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		mem := v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v.Args[1]
		if !(sym == nil && c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce8shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbStoreOnce8idx_0(v *Value) bool {
	// match: (StoreOnce8idx ptr (MOVWconst [c]) val mem)
	// cond:
	// result: (StoreOnce8 [c] ptr val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		v.reset(OpThumbStoreOnce8)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce8idx (MOVWconst [c]) ptr val mem)
	// cond:
	// result: (StoreOnce8 [c] ptr val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		val := v.Args[2]
		v.reset(OpThumbStoreOnce8)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce8idx ptr (SLLconst idx [c]) val mem)
	// cond: c <= 3
	// result: (StoreOnce8shiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce8shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (StoreOnce8idx (SLLconst idx [c]) ptr val mem)
	// cond: c <= 3
	// result: (StoreOnce8shiftLL ptr idx [c] val mem)
	for {
		mem := v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		if !(c <= 3) {
			break
		}
		v.reset(OpThumbStoreOnce8shiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTEQ_0(v *Value) bool {
	// match: (TEQ x (MOVWconst [c]))
	// cond:
	// result: (TEQconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbTEQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (TEQ (MOVWconst [c]) x)
	// cond:
	// result: (TEQconst [c] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbTEQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (TEQ x (SLLconst [c] y))
	// cond:
	// result: (TEQshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbTEQshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (TEQ (SLLconst [c] y) x)
	// cond:
	// result: (TEQshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbTEQshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (TEQ x (SRLconst [c] y))
	// cond:
	// result: (TEQshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbTEQshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (TEQ (SRLconst [c] y) x)
	// cond:
	// result: (TEQshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbTEQshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (TEQ x (SRAconst [c] y))
	// cond:
	// result: (TEQshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbTEQshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (TEQ (SRAconst [c] y) x)
	// cond:
	// result: (TEQshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbTEQshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTEQconst_0(v *Value) bool {
	// match: (TEQconst (MOVWconst [x]) [y])
	// cond: int32(x^y)==0
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x^y) == 0) {
			break
		}
		v.reset(OpThumbFlagEQ)
		return true
	}
	// match: (TEQconst (MOVWconst [x]) [y])
	// cond: int32(x^y)<0
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x^y) < 0) {
			break
		}
		v.reset(OpThumbFlagLT_UGT)
		return true
	}
	// match: (TEQconst (MOVWconst [x]) [y])
	// cond: int32(x^y)>0
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x^y) > 0) {
			break
		}
		v.reset(OpThumbFlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTEQshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (TEQshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (TEQconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (TEQconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbTEQconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTEQshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (TEQshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (TEQconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (TEQconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbTEQconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTEQshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (TEQshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (TEQconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (TEQconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbTEQconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTST_0(v *Value) bool {
	// match: (TST x (MOVWconst [c]))
	// cond:
	// result: (TSTconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbTSTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (TST (MOVWconst [c]) x)
	// cond:
	// result: (TSTconst [c] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbTSTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (TST x (SLLconst [c] y))
	// cond:
	// result: (TSTshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbTSTshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (TST (SLLconst [c] y) x)
	// cond:
	// result: (TSTshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbTSTshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (TST x (SRLconst [c] y))
	// cond:
	// result: (TSTshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbTSTshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (TST (SRLconst [c] y) x)
	// cond:
	// result: (TSTshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbTSTshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (TST x (SRAconst [c] y))
	// cond:
	// result: (TSTshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbTSTshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (TST (SRAconst [c] y) x)
	// cond:
	// result: (TSTshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbTSTshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTSTconst_0(v *Value) bool {
	// match: (TSTconst (MOVWconst [x]) [y])
	// cond: int32(x&y)==0
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) == 0) {
			break
		}
		v.reset(OpThumbFlagEQ)
		return true
	}
	// match: (TSTconst (MOVWconst [x]) [y])
	// cond: int32(x&y)<0
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) < 0) {
			break
		}
		v.reset(OpThumbFlagLT_UGT)
		return true
	}
	// match: (TSTconst (MOVWconst [x]) [y])
	// cond: int32(x&y)>0
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) > 0) {
			break
		}
		v.reset(OpThumbFlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTSTshiftLL_0(v *Value) bool {
	b := v.Block
	// match: (TSTshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (TSTconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (TSTconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbTSTconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTSTshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (TSTshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (TSTconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (TSTconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbTSTconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbTSTshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (TSTshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (TSTconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (TSTconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbTSTconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXOR_0(v *Value) bool {
	// match: (XOR x (MOVWconst [c]))
	// cond:
	// result: (XORconst [c] x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbXORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (XOR (MOVWconst [c]) x)
	// cond:
	// result: (XORconst [c] x)
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbXORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (XOR x (SLLconst [c] y))
	// cond:
	// result: (XORshiftLL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbXORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (XOR (SLLconst [c] y) x)
	// cond:
	// result: (XORshiftLL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbXORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (XOR x (SRLconst [c] y))
	// cond:
	// result: (XORshiftRL x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbXORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (XOR (SRLconst [c] y) x)
	// cond:
	// result: (XORshiftRL x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbXORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (XOR x (SRAconst [c] y))
	// cond:
	// result: (XORshiftRA x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbXORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (XOR (SRAconst [c] y) x)
	// cond:
	// result: (XORshiftRA x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbXORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (XOR x (SRRconst [c] y))
	// cond:
	// result: (XORshiftRR x y [c])
	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRRconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpThumbXORshiftRR)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (XOR (SRRconst [c] y) x)
	// cond:
	// result: (XORshiftRR x y [c])
	for {
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRRconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		v.reset(OpThumbXORshiftRR)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXOR_10(v *Value) bool {
	// match: (XOR x x)
	// cond:
	// result: (MOVWconst [0])
	for {
		x := v.Args[1]
		if x != v.Args[0] {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXORconst_0(v *Value) bool {
	// match: (XORconst [0] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (XORconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [c^d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpThumbMOVWconst)
		v.AuxInt = c ^ d
		return true
	}
	// match: (XORconst [c] (XORconst [d] x))
	// cond:
	// result: (XORconst [c^d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpThumbXORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpThumbXORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXORshiftLL_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (XORshiftLL (MOVWconst [c]) x [d])
	// cond:
	// result: (XORconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftLL x (MOVWconst [c]) [d])
	// cond:
	// result: (XORconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbXORconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL [c] (SRLconst x [32-c]) x)
	// cond:
	// result: (SRRconst [32-c] x)
	for {
		c := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		if x != v_0.Args[0] {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL <typ.UInt16> [8] (BFXU <typ.UInt16> [armBFAuxInt(8, 8)] x) x)
	// cond:
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 {
			break
		}
		if v.AuxInt != 8 {
			break
		}
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbBFXU {
			break
		}
		if v_0.Type != typ.UInt16 {
			break
		}
		if v_0.AuxInt != armBFAuxInt(8, 8) {
			break
		}
		if x != v_0.Args[0] {
			break
		}
		v.reset(OpThumbREV16)
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL <typ.UInt16> [8] (SRLconst <typ.UInt16> [24] (SLLconst [16] x)) x)
	// cond:
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 {
			break
		}
		if v.AuxInt != 8 {
			break
		}
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSRLconst {
			break
		}
		if v_0.Type != typ.UInt16 {
			break
		}
		if v_0.AuxInt != 24 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpThumbSLLconst {
			break
		}
		if v_0_0.AuxInt != 16 {
			break
		}
		if x != v_0_0.Args[0] {
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
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXORshiftRA_0(v *Value) bool {
	b := v.Block
	// match: (XORshiftRA (MOVWconst [c]) x [d])
	// cond:
	// result: (XORconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRA x (MOVWconst [c]) [d])
	// cond:
	// result: (XORconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbXORconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (XORshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXORshiftRL_0(v *Value) bool {
	b := v.Block
	// match: (XORshiftRL (MOVWconst [c]) x [d])
	// cond:
	// result: (XORconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRL x (MOVWconst [c]) [d])
	// cond:
	// result: (XORconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbXORconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (XORshiftRL [c] (SLLconst x [32-c]) x)
	// cond:
	// result: (SRRconst [ c] x)
	for {
		c := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbSLLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		if x != v_0.Args[0] {
			break
		}
		v.reset(OpThumbSRRconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (XORshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpThumbMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueThumb_OpThumbXORshiftRR_0(v *Value) bool {
	b := v.Block
	// match: (XORshiftRR (MOVWconst [c]) x [d])
	// cond:
	// result: (XORconst [c] (SRRconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		x := v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpThumbMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpThumbXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpThumbSRRconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRR x (MOVWconst [c]) [d])
	// cond:
	// result: (XORconst x [int64(int32(uint32(c)>>uint64(d)|uint32(c)<<uint64(32-d)))])
	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpThumbMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpThumbXORconst)
		v.AuxInt = int64(int32(uint32(c)>>uint64(d) | uint32(c)<<uint64(32-d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueThumb_OpTrunc16to8_0(v *Value) bool {
	// match: (Trunc16to8 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpTrunc32to16_0(v *Value) bool {
	// match: (Trunc32to16 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpTrunc32to8_0(v *Value) bool {
	// match: (Trunc32to8 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpWB_0(v *Value) bool {
	// match: (WB {fn} destptr srcptr mem)
	// cond:
	// result: (LoweredWB {fn} destptr srcptr mem)
	for {
		fn := v.Aux
		mem := v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		v.reset(OpThumbLoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueThumb_OpXor16_0(v *Value) bool {
	// match: (Xor16 x y)
	// cond:
	// result: (XOR x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpXor32_0(v *Value) bool {
	// match: (Xor32 x y)
	// cond:
	// result: (XOR x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpXor8_0(v *Value) bool {
	// match: (Xor8 x y)
	// cond:
	// result: (XOR x y)
	for {
		y := v.Args[1]
		x := v.Args[0]
		v.reset(OpThumbXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueThumb_OpZero_0(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// cond:
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v.Args[1]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// cond:
	// result: (MOVBstore ptr (MOVWconst [0]) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpThumbMOVBstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVWconst [0]) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		t := v.Aux
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// cond:
	// result: (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = 1
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVWconst [0]) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpThumbMOVWstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [2] ptr (MOVWconst [0]) (MOVHstore [0] ptr (MOVWconst [0]) mem))
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpThumbMOVHstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVHstore, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [4] ptr mem)
	// cond:
	// result: (MOVBstore [3] ptr (MOVWconst [0]) (MOVBstore [2] ptr (MOVWconst [0]) (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = 3
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [3] ptr mem)
	// cond:
	// result: (MOVBstore [2] ptr (MOVWconst [0]) (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpThumbMOVBstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpThumbMOVBstore, types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: s%4 == 0 && s > 4 && s <= 512 && t.(*types.Type).Alignment()%4 == 0 && !config.noDuffDevice
	// result: (DUFFZERO [4 * (128 - s/4)] ptr (MOVWconst [0]) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		mem := v.Args[1]
		ptr := v.Args[0]
		if !(s%4 == 0 && s > 4 && s <= 512 && t.(*types.Type).Alignment()%4 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpThumbDUFFZERO)
		v.AuxInt = 4 * (128 - s/4)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: (s > 512 || config.noDuffDevice) || t.(*types.Type).Alignment()%4 != 0
	// result: (LoweredZero [t.(*types.Type).Alignment()] ptr (ADDconst <ptr.Type> ptr [s-moveSize(t.(*types.Type).Alignment(), config)]) (MOVWconst [0]) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		mem := v.Args[1]
		ptr := v.Args[0]
		if !((s > 512 || config.noDuffDevice) || t.(*types.Type).Alignment()%4 != 0) {
			break
		}
		v.reset(OpThumbLoweredZero)
		v.AuxInt = t.(*types.Type).Alignment()
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpThumbADDconst, ptr.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpThumbMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueThumb_OpZeroExt16to32_0(v *Value) bool {
	// match: (ZeroExt16to32 x)
	// cond:
	// result: (MOVHUreg x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpZeroExt8to16_0(v *Value) bool {
	// match: (ZeroExt8to16 x)
	// cond:
	// result: (MOVBUreg x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpZeroExt8to32_0(v *Value) bool {
	// match: (ZeroExt8to32 x)
	// cond:
	// result: (MOVBUreg x)
	for {
		x := v.Args[0]
		v.reset(OpThumbMOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueThumb_OpZeromask_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Zeromask x)
	// cond:
	// result: (SRAconst (RSBshiftRL <typ.Int32> x x [1]) [31])
	for {
		x := v.Args[0]
		v.reset(OpThumbSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpThumbRSBshiftRL, typ.Int32)
		v0.AuxInt = 1
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteBlockThumb(b *Block) bool {
	config := b.Func.Config
	typ := &config.Types
	_ = typ
	v := b.Control
	_ = v
	switch b.Kind {
	case BlockThumbEQ:
		// match: (EQ (FlagEQ) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagEQ {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (EQ (FlagLT_ULT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagLT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (EQ (FlagLT_UGT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagLT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (EQ (FlagGT_ULT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagGT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (EQ (FlagGT_UGT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagGT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (EQ (InvertFlags cmp) yes no)
		// cond:
		// result: (EQ cmp yes no)
		for v.Op == OpThumbInvertFlags {
			cmp := v.Args[0]
			b.Kind = BlockThumbEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMP x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMP a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMN x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMN a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TST x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbTST, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQ x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbTEQ, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbEQ
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockThumbGE:
		// match: (GE (FlagEQ) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagEQ {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (GE (FlagLT_ULT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagLT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (GE (FlagLT_UGT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagLT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (GE (FlagGT_ULT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagGT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (GE (FlagGT_UGT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagGT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (GE (InvertFlags cmp) yes no)
		// cond:
		// result: (LE cmp yes no)
		for v.Op == OpThumbInvertFlags {
			cmp := v.Args[0]
			b.Kind = BlockThumbLE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMP x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMP a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMN x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMN a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (GE (TST x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbTST, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQ x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbTEQ, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGE
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockThumbGT:
		// match: (GT (FlagEQ) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagEQ {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagLT_ULT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagLT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagLT_UGT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagLT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagGT_ULT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagGT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (GT (FlagGT_UGT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagGT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (GT (InvertFlags cmp) yes no)
		// cond:
		// result: (LT cmp yes no)
		for v.Op == OpThumbInvertFlags {
			cmp := v.Args[0]
			b.Kind = BlockThumbLT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMP x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMP a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMN x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (GT (TST x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbTST, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMN a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQ x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbTEQ, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbGT
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockIf:
		// match: (If (Equal cc) yes no)
		// cond:
		// result: (EQ cc yes no)
		for v.Op == OpThumbEqual {
			cc := v.Args[0]
			b.Kind = BlockThumbEQ
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (If (NotEqual cc) yes no)
		// cond:
		// result: (NE cc yes no)
		for v.Op == OpThumbNotEqual {
			cc := v.Args[0]
			b.Kind = BlockThumbNE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (If (LessThan cc) yes no)
		// cond:
		// result: (LT cc yes no)
		for v.Op == OpThumbLessThan {
			cc := v.Args[0]
			b.Kind = BlockThumbLT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (If (LessThanU cc) yes no)
		// cond:
		// result: (ULT cc yes no)
		for v.Op == OpThumbLessThanU {
			cc := v.Args[0]
			b.Kind = BlockThumbULT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (If (LessEqual cc) yes no)
		// cond:
		// result: (LE cc yes no)
		for v.Op == OpThumbLessEqual {
			cc := v.Args[0]
			b.Kind = BlockThumbLE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (If (LessEqualU cc) yes no)
		// cond:
		// result: (ULE cc yes no)
		for v.Op == OpThumbLessEqualU {
			cc := v.Args[0]
			b.Kind = BlockThumbULE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (If (GreaterThan cc) yes no)
		// cond:
		// result: (GT cc yes no)
		for v.Op == OpThumbGreaterThan {
			cc := v.Args[0]
			b.Kind = BlockThumbGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (If (GreaterThanU cc) yes no)
		// cond:
		// result: (UGT cc yes no)
		for v.Op == OpThumbGreaterThanU {
			cc := v.Args[0]
			b.Kind = BlockThumbUGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (If (GreaterEqual cc) yes no)
		// cond:
		// result: (GE cc yes no)
		for v.Op == OpThumbGreaterEqual {
			cc := v.Args[0]
			b.Kind = BlockThumbGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (If (GreaterEqualU cc) yes no)
		// cond:
		// result: (UGE cc yes no)
		for v.Op == OpThumbGreaterEqualU {
			cc := v.Args[0]
			b.Kind = BlockThumbUGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (If cond yes no)
		// cond:
		// result: (NE (CMPconst [0] cond) yes no)
		for {
			cond := b.Control
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(cond)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockThumbLE:
		// match: (LE (FlagEQ) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagEQ {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (LE (FlagLT_ULT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagLT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (LE (FlagLT_UGT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagLT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (LE (FlagGT_ULT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagGT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (LE (FlagGT_UGT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagGT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (LE (InvertFlags cmp) yes no)
		// cond:
		// result: (GE cmp yes no)
		for v.Op == OpThumbInvertFlags {
			cmp := v.Args[0]
			b.Kind = BlockThumbGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMP x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMP a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMN x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMN a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (LE (TST x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbTST, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQ x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbTEQ, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLE
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockThumbLT:
		// match: (LT (FlagEQ) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagEQ {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (LT (FlagLT_ULT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagLT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (LT (FlagLT_UGT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagLT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (LT (FlagGT_ULT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagGT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (LT (FlagGT_UGT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagGT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (LT (InvertFlags cmp) yes no)
		// cond:
		// result: (GT cmp yes no)
		for v.Op == OpThumbInvertFlags {
			cmp := v.Args[0]
			b.Kind = BlockThumbGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMP x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMP a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMN x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMN a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (LT (TST x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbTST, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQ x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbTEQ, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbLT
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockThumbNE:
		// match: (NE (CMPconst [0] (Equal cc)) yes no)
		// cond:
		// result: (EQ cc yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpThumbEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockThumbEQ
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] (NotEqual cc)) yes no)
		// cond:
		// result: (NE cc yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpThumbNotEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockThumbNE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] (LessThan cc)) yes no)
		// cond:
		// result: (LT cc yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpThumbLessThan {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockThumbLT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] (LessThanU cc)) yes no)
		// cond:
		// result: (ULT cc yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpThumbLessThanU {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockThumbULT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] (LessEqual cc)) yes no)
		// cond:
		// result: (LE cc yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpThumbLessEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockThumbLE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] (LessEqualU cc)) yes no)
		// cond:
		// result: (ULE cc yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpThumbLessEqualU {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockThumbULE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] (GreaterThan cc)) yes no)
		// cond:
		// result: (GT cc yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpThumbGreaterThan {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockThumbGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] (GreaterThanU cc)) yes no)
		// cond:
		// result: (UGT cc yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpThumbGreaterThanU {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockThumbUGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] (GreaterEqual cc)) yes no)
		// cond:
		// result: (GE cc yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpThumbGreaterEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockThumbGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] (GreaterEqualU cc)) yes no)
		// cond:
		// result: (UGE cc yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpThumbGreaterEqualU {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockThumbUGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}
		// match: (NE (FlagEQ) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagEQ {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (NE (FlagLT_ULT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagLT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (NE (FlagLT_UGT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagLT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (NE (FlagGT_ULT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagGT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (NE (FlagGT_UGT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagGT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (NE (InvertFlags cmp) yes no)
		// cond:
		// result: (NE cmp yes no)
		for v.Op == OpThumbInvertFlags {
			cmp := v.Args[0]
			b.Kind = BlockThumbNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMP x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMP a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMN x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADD {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMN a (MUL <x.Type> x y)) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v.Pos, OpThumbMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (TST x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbAND {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbTST, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQ x y) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXOR {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbTEQ, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQconst [c] x) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftLL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftRL x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftRA x y [c]) yes no)
		for v.Op == OpThumbCMPconst {
			if v.AuxInt != 0 {
				break
			}
			l := v.Args[0]
			if l.Op != OpThumbXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Kind = BlockThumbNE
			v0 := b.NewValue0(v.Pos, OpThumbTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockThumbUGE:
		// match: (UGE (FlagEQ) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagEQ {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (UGE (FlagLT_ULT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagLT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (UGE (FlagLT_UGT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagLT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (UGE (FlagGT_ULT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagGT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (UGE (FlagGT_UGT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagGT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (UGE (InvertFlags cmp) yes no)
		// cond:
		// result: (ULE cmp yes no)
		for v.Op == OpThumbInvertFlags {
			cmp := v.Args[0]
			b.Kind = BlockThumbULE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockThumbUGT:
		// match: (UGT (FlagEQ) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagEQ {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (UGT (FlagLT_ULT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagLT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (UGT (FlagLT_UGT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagLT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (UGT (FlagGT_ULT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagGT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (UGT (FlagGT_UGT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagGT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (UGT (InvertFlags cmp) yes no)
		// cond:
		// result: (ULT cmp yes no)
		for v.Op == OpThumbInvertFlags {
			cmp := v.Args[0]
			b.Kind = BlockThumbULT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockThumbULE:
		// match: (ULE (FlagEQ) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagEQ {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (ULE (FlagLT_ULT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagLT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (ULE (FlagLT_UGT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagLT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (ULE (FlagGT_ULT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagGT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (ULE (FlagGT_UGT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagGT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (ULE (InvertFlags cmp) yes no)
		// cond:
		// result: (UGE cmp yes no)
		for v.Op == OpThumbInvertFlags {
			cmp := v.Args[0]
			b.Kind = BlockThumbUGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockThumbULT:
		// match: (ULT (FlagEQ) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagEQ {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (ULT (FlagLT_ULT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagLT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (ULT (FlagLT_UGT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagLT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (ULT (FlagGT_ULT) yes no)
		// cond:
		// result: (First nil yes no)
		for v.Op == OpThumbFlagGT_ULT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
		// match: (ULT (FlagGT_UGT) yes no)
		// cond:
		// result: (First nil no yes)
		for v.Op == OpThumbFlagGT_UGT {
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
		// match: (ULT (InvertFlags cmp) yes no)
		// cond:
		// result: (UGT cmp yes no)
		for v.Op == OpThumbInvertFlags {
			cmp := v.Args[0]
			b.Kind = BlockThumbUGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	}
	return false
}
