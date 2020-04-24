#include "textflag.h"

TEXT ·load32(SB),NOSPLIT,$0-8
	MOVW  addr+0(FP), A0
	MOVW  (A0), A1
	MOVW  A1, ret+4(FP)
	RET

TEXT ·load16(SB),NOSPLIT,$0-6
	MOVW   addr+0(FP), A0
	MOVHU  (A0), A1
	MOVH   A1, ret+4(FP)
	RET

TEXT ·load8(SB),NOSPLIT,$0-5
	MOVW   addr+0(FP), A0
	MOVBU  (A0), A1
	MOVB   A1, ret+4(FP)
	RET

TEXT ·store32(SB),NOSPLIT,$0-8
	MOVW  addr+0(FP), A0
	MOVW  v+4(FP), A1
	MOVW  A1, (A0)
	RET

TEXT ·store16(SB),NOSPLIT,$0-6
	MOVW   addr+0(FP), A0
	MOVHU  v+4(FP), A1
	MOVH   A1, (A0)
	RET

TEXT ·store8(SB),NOSPLIT,$0-5
	MOVW   addr+0(FP), A0
	MOVBU  v+4(FP), A1
	MOVB   A1, (A0)
	RET

#define FENCE WORD $0x0ff0000f

TEXT ·MB(SB),NOSPLIT,$0
	FENCE
	RET
