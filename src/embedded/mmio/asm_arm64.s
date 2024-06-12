#include "textflag.h"

TEXT ·load32(SB),NOSPLIT,$0-8
	MOVW  addr+0(FP), R0
	MOVW  (R0), R1
	MOVW  R1, ret+4(FP)
	RET

TEXT ·load16(SB),NOSPLIT,$0-6
	MOVW   addr+0(FP), R0
	MOVHU  (R0), R1
	MOVH   R1, ret+4(FP)
	RET

TEXT ·load8(SB),NOSPLIT,$0-5
	MOVW   addr+0(FP), R0
	MOVBU  (R0), R1
	MOVB   R1, ret+4(FP)
	RET

TEXT ·store32(SB),NOSPLIT,$0-8
	MOVW  addr+0(FP), R0
	MOVW  v+4(FP), R1
	MOVW  R1, (R0)
	RET

TEXT ·store16(SB),NOSPLIT,$0-6
	MOVW   addr+0(FP), R0
	MOVHU  v+4(FP), R1
	MOVH   R1, (R0)
	RET

TEXT ·store8(SB),NOSPLIT,$0-5
	MOVW   addr+0(FP), R0
	MOVBU  v+4(FP), R1
	MOVB   R1, (R0)
	RET

TEXT ·MB(SB),NOSPLIT,$0
	// use DSB instead of DMB because an IO access can affect CPU directly (eg:
	// generate interrupt, change CPU behavior via memory-mapped control reg.)
	DSB $0xF
	RET
