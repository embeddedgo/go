#include "textflag.h"


TEXT ·load64(SB),NOSPLIT,$0-16
	MOV  addr+0(FP), A0
	MOV  (A0), A1
	MOV  A1, ret+8(FP)
	RET

TEXT ·load32(SB),NOSPLIT,$0-12
	MOV    addr+0(FP), A0
	MOVWU  (A0), A1
	MOVW   A1, ret+8(FP)
	RET

TEXT ·load16(SB),NOSPLIT,$0-10
	MOVW   addr+0(FP), A0
	MOVHU  (A0), A1
	MOVH   A1, ret+8(FP)
	RET

TEXT ·load8(SB),NOSPLIT,$0-9
	MOV    addr+0(FP), A0
	MOVBU  (A0), A1
	MOVB   A1, ret+8(FP)
	RET

TEXT ·store64(SB),NOSPLIT,$0-16
	MOV  addr+0(FP), A0
	MOV  v+8(FP), A1
	MOV  A1, (A0)
	RET

TEXT ·store32(SB),NOSPLIT,$0-12
	MOV    addr+0(FP), A0
	MOVWU  v+8(FP), A1
	MOVW   A1, (A0)
	RET

TEXT ·store16(SB),NOSPLIT,$0-10
	MOV    addr+0(FP), A0
	MOVHU  v+8(FP), A1
	MOVH   A1, (A0)
	RET

TEXT ·store8(SB),NOSPLIT,$0-9
	MOV    addr+0(FP), A0
	MOVBU  v+8(FP), A1
	MOVB   A1, (A0)
	RET

#define FENCE WORD $0x0ff0000f

TEXT ·MB(SB),NOSPLIT,$0
	FENCE
	RET
