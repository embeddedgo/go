#include "textflag.h"

// func publicationBarrier()
TEXT ·publicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	SYNC
	RET
