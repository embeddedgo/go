#include "textflag.h"

// func publicationBarrier()
TEXT ·publicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	WORD  $0x0ff0000f  // FENCE
	RET
