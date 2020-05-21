#include "textflag.h"

#define CSRSI(U5,CSR) WORD $(0x6073 + U5<<15 + CSR<<20)
#define mie 0x304
#define MTIE (1<<7)

TEXT ·enableTimerInterrupt(SB),NOSPLIT|NOFRAME,$0
	CSRSI  (MTIE, mie)
	RET
