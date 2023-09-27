#include "textflag.h"

TEXT ·mtc0(SB),NOSPLIT,$0-12
	MOVV   creg+0(FP), R4
	MOVWU  val+8(FP), R5
	MOVW   $0, R6
	BEQ    R6, R4, creg0
	MOVW   $1, R6
	BEQ    R6, R4, creg1
	MOVW   $2, R6
	BEQ    R6, R4, creg2
	MOVW   $3, R6
	BEQ    R6, R4, creg3
	MOVW   $4, R6
	BEQ    R6, R4, creg4
	MOVW   $5, R6
	BEQ    R6, R4, creg5
	MOVW   $6, R6
	BEQ    R6, R4, creg6
	MOVW   $7, R6
	BEQ    R6, R4, creg7
	MOVW   $8, R6
	BEQ    R6, R4, creg8
	MOVW   $9, R6
	BEQ    R6, R4, creg9
	MOVW   $10, R6
	BEQ    R6, R4, creg10
	MOVW   $11, R6
	BEQ    R6, R4, creg11
	MOVW   $12, R6
	BEQ    R6, R4, creg12
	MOVW   $13, R6
	BEQ    R6, R4, creg13
	MOVW   $14, R6
	BEQ    R6, R4, creg14
	MOVW   $15, R6
	BEQ    R6, R4, creg15
	MOVW   $16, R6
	BEQ    R6, R4, creg16
	MOVW   $17, R6
	BEQ    R6, R4, creg17
	MOVW   $18, R6
	BEQ    R6, R4, creg18
	MOVW   $19, R6
	BEQ    R6, R4, creg19
	MOVW   $20, R6
	BEQ    R6, R4, creg20
	MOVW   $21, R6
	BEQ    R6, R4, creg21
	MOVW   $22, R6
	BEQ    R6, R4, creg22
	MOVW   $23, R6
	BEQ    R6, R4, creg23
	MOVW   $24, R6
	BEQ    R6, R4, creg24
	MOVW   $25, R6
	BEQ    R6, R4, creg25
	MOVW   $26, R6
	BEQ    R6, R4, creg26
	MOVW   $27, R6
	BEQ    R6, R4, creg27
	MOVW   $28, R6
	BEQ    R6, R4, creg28
	MOVW   $29, R6
	BEQ    R6, R4, creg29
	MOVW   $30, R6
	BEQ    R6, R4, creg30
	MOVW   $31, R6
	BEQ    R6, R4, creg31
creg0:
	MOVW   R5, M(0)
	RET
creg1:
	MOVW   R5, M(1)
	RET
creg2:
	MOVW   R5, M(2)
	RET
creg3:
	MOVW   R5, M(3)
	RET
creg4:
	MOVW   R5, M(4)
	RET
creg5:
	MOVW   R5, M(5)
	RET
creg6:
	MOVW   R5, M(6)
	RET
creg7:
	MOVW   R5, M(7)
	RET
creg8:
	MOVW   R5, M(8)
	RET
creg9:
	MOVW   R5, M(9)
	RET
creg10:
	MOVW   R5, M(10)
	RET
creg11:
	MOVW   R5, M(11)
	RET
creg12:
	MOVW   R5, M(12)
	RET
creg13:
	MOVW   R5, M(13)
	RET
creg14:
	MOVW   R5, M(14)
	RET
creg15:
	MOVW   R5, M(15)
	RET
creg16:
	MOVW   R5, M(16)
	RET
creg17:
	MOVW   R5, M(17)
	RET
creg18:
	MOVW   R5, M(18)
	RET
creg19:
	MOVW   R5, M(19)
	RET
creg20:
	MOVW   R5, M(20)
	RET
creg21:
	MOVW   R5, M(21)
	RET
creg22:
	MOVW   R5, M(22)
	RET
creg23:
	MOVW   R5, M(23)
	RET
creg24:
	MOVW   R5, M(24)
	RET
creg25:
	MOVW   R5, M(25)
	RET
creg26:
	MOVW   R5, M(26)
	RET
creg27:
	MOVW   R5, M(27)
	RET
creg28:
	MOVW   R5, M(28)
	RET
creg29:
	MOVW   R5, M(29)
	RET
creg30:
	MOVW   R5, M(30)
	RET
creg31:
	MOVW   R5, M(31)
	RET

TEXT ·mfc0(SB),NOSPLIT,$0-12
	MOVV   creg+0(FP), R4
	MOVW   $0, R6
	BEQ    R6, R4, creg0
	MOVW   $1, R6
	BEQ    R6, R4, creg1
	MOVW   $2, R6
	BEQ    R6, R4, creg2
	MOVW   $3, R6
	BEQ    R6, R4, creg3
	MOVW   $4, R6
	BEQ    R6, R4, creg4
	MOVW   $5, R6
	BEQ    R6, R4, creg5
	MOVW   $6, R6
	BEQ    R6, R4, creg6
	MOVW   $7, R6
	BEQ    R6, R4, creg7
	MOVW   $8, R6
	BEQ    R6, R4, creg8
	MOVW   $9, R6
	BEQ    R6, R4, creg9
	MOVW   $10, R6
	BEQ    R6, R4, creg10
	MOVW   $11, R6
	BEQ    R6, R4, creg11
	MOVW   $12, R6
	BEQ    R6, R4, creg12
	MOVW   $13, R6
	BEQ    R6, R4, creg13
	MOVW   $14, R6
	BEQ    R6, R4, creg14
	MOVW   $15, R6
	BEQ    R6, R4, creg15
	MOVW   $16, R6
	BEQ    R6, R4, creg16
	MOVW   $17, R6
	BEQ    R6, R4, creg17
	MOVW   $18, R6
	BEQ    R6, R4, creg18
	MOVW   $19, R6
	BEQ    R6, R4, creg19
	MOVW   $20, R6
	BEQ    R6, R4, creg20
	MOVW   $21, R6
	BEQ    R6, R4, creg21
	MOVW   $22, R6
	BEQ    R6, R4, creg22
	MOVW   $23, R6
	BEQ    R6, R4, creg23
	MOVW   $24, R6
	BEQ    R6, R4, creg24
	MOVW   $25, R6
	BEQ    R6, R4, creg25
	MOVW   $26, R6
	BEQ    R6, R4, creg26
	MOVW   $27, R6
	BEQ    R6, R4, creg27
	MOVW   $28, R6
	BEQ    R6, R4, creg28
	MOVW   $29, R6
	BEQ    R6, R4, creg29
	MOVW   $30, R6
	BEQ    R6, R4, creg30
	MOVW   $31, R6
	BEQ    R6, R4, creg31
creg0:
	MOVW   M(0), R5
	MOVW   R5, ret+8(FP)
	RET
creg1:
	MOVW   M(1), R5
	MOVW   R5, ret+8(FP)
	RET
creg2:
	MOVW   M(2), R5
	MOVW   R5, ret+8(FP)
	RET
creg3:
	MOVW   M(3), R5
	MOVW   R5, ret+8(FP)
	RET
creg4:
	MOVW   M(4), R5
	MOVW   R5, ret+8(FP)
	RET
creg5:
	MOVW   M(5), R5
	MOVW   R5, ret+8(FP)
	RET
creg6:
	MOVW   M(6), R5
	MOVW   R5, ret+8(FP)
	RET
creg7:
	MOVW   M(7), R5
	MOVW   R5, ret+8(FP)
	RET
creg8:
	MOVW   M(8), R5
	MOVW   R5, ret+8(FP)
	RET
creg9:
	MOVW   M(9), R5
	MOVW   R5, ret+8(FP)
	RET
creg10:
	MOVW   M(10), R5
	MOVW   R5, ret+8(FP)
	RET
creg11:
	MOVW   M(11), R5
	MOVW   R5, ret+8(FP)
	RET
creg12:
	MOVW   M(12), R5
	MOVW   R5, ret+8(FP)
	RET
creg13:
	MOVW   M(13), R5
	MOVW   R5, ret+8(FP)
	RET
creg14:
	MOVW   M(14), R5
	MOVW   R5, ret+8(FP)
	RET
creg15:
	MOVW   M(15), R5
	MOVW   R5, ret+8(FP)
	RET
creg16:
	MOVW   M(16), R5
	MOVW   R5, ret+8(FP)
	RET
creg17:
	MOVW   M(17), R5
	MOVW   R5, ret+8(FP)
	RET
creg18:
	MOVW   M(18), R5
	MOVW   R5, ret+8(FP)
	RET
creg19:
	MOVW   M(19), R5
	MOVW   R5, ret+8(FP)
	RET
creg20:
	MOVW   M(20), R5
	MOVW   R5, ret+8(FP)
	RET
creg21:
	MOVW   M(21), R5
	MOVW   R5, ret+8(FP)
	RET
creg22:
	MOVW   M(22), R5
	MOVW   R5, ret+8(FP)
	RET
creg23:
	MOVW   M(23), R5
	MOVW   R5, ret+8(FP)
	RET
creg24:
	MOVW   M(24), R5
	MOVW   R5, ret+8(FP)
	RET
creg25:
	MOVW   M(25), R5
	MOVW   R5, ret+8(FP)
	RET
creg26:
	MOVW   M(26), R5
	MOVW   R5, ret+8(FP)
	RET
creg27:
	MOVW   M(27), R5
	MOVW   R5, ret+8(FP)
	RET
creg28:
	MOVW   M(28), R5
	MOVW   R5, ret+8(FP)
	RET
creg29:
	MOVW   M(29), R5
	MOVW   R5, ret+8(FP)
	RET
creg30:
	MOVW   M(30), R5
	MOVW   R5, ret+8(FP)
	RET
creg31:
	MOVW   M(31), R5
	MOVW   R5, ret+8(FP)
	RET
