#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·rt0_target(SB),NOSPLIT|NOFRAME,$0
	JMP machine·rt0(SB)

TEXT ·unhandledExcepton_target(SB),NOSPLIT|NOFRAME,$0
	JMP machine·unhandledException(SB)
