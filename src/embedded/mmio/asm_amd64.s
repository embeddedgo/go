#include "textflag.h"

TEXT ·load32(SB),NOSPLIT,$0-12
	MOVQ  addr+0(FP), AX
	MOVL  (AX), BX
	MOVL  BX, ret+8(FP)
	RET

TEXT ·load16(SB),NOSPLIT,$0-10
	MOVQ  addr+0(FP), AX
	MOVW  (AX), BX
	MOVW  BX, ret+8(FP)
	RET

TEXT ·load8(SB),NOSPLIT,$0-9
	MOVQ  addr+0(FP), AX
	MOVB  (AX), BX
	MOVB  BX, ret+8(FP)
	RET

TEXT ·store32(SB),NOSPLIT,$0-12
	MOVQ  addr+0(FP), AX
	MOVL  v+8(FP), BX
	MOVL  BX, (AX)
	RET

TEXT ·store16(SB),NOSPLIT,$0-10
	MOVQ  addr+0(FP), AX
	MOVW  v+8(FP), BX
	MOVW  BX, (AX)
	RET

TEXT ·store8(SB),NOSPLIT,$0-9
	MOVQ  addr+0(FP), AX
	MOVB  v+8(FP), BX
	MOVB  BX, (AX)
	RET

TEXT ·MB(SB),NOSPLIT,$0
	// BUG: memory barrier instruction
	RET
