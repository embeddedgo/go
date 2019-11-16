#include "../../../../../runtime/textflag.h"

TEXT foo(SB), DUPOK|NOSPLIT, $0

// Load/store
	MOVW  0x28(R15), R7  // 4f0a
	MOVW  R1, (R2)(R7)   // 51d1
	MOVW  R8, (R9)(R7)   // f849 8007
	MOVW  R1, (R7)       // 6039
	MOVW  g, (R7)        // f8c7 a000
	MOVW  (R2)(R7), R1   // 59d1
	MOVW  (R9)(R7), R8   // f859 8007
	MOVW  (R7), R1       // 6839
	MOVW  (R7), R10      // f8d7 a000

	MOVB   (R7), R3  // f997 3000
	MOVH   (R7), R3  // f9b7 3000
	MOVBU  (R7), R3  // 783b
	MOVHU  (R7), R3  // 883b
	MOVW   (R7), R3  // 683b
	MOVB   R3, (R7)  // 703b
	MOVH   R3, (R7)  // 803b
	MOVW   R3, (R7)  // 603b

// FP load/store
	MOVF  0x10(R4), F0   // ed94 0a10
	MOVF  0x20(R4), F1   // ed94 1a20
	MOVF  0x30(R4), F2   // ed94 2a30
	MOVD  -0x10(R4), F0  // ed14 0b10
	MOVD  -0x20(R4), F1  // ed14 1b20
	MOVD  -0x30(R4), F2  // ed14 2b30
	MOVF  F0, -0x10(R4)  // ed04 0a10
	MOVF  F1, -0x20(R4)  // ed04 1a20
	MOVF  F2, -0x30(R4)  // ed04 2a30
	MOVD  F0, 0x10(R4)   // ed84 0b10
	MOVD  F1, 0x20(R4)   // ed84 1b20
	MOVD  F2, 0x30(R4)   // ed84 2b30

// FP data processing
	SQRTF  F3, F4  // eeb1 4ac3
	SQRTD  F5, F6  // eeb1 6bc5

// System registers
	MOVW  R1, IAPSR  // f381 8801
	MOVW  APSR, R3   // f3ef 8300

// Branch
	B   (R14)  // 4770

