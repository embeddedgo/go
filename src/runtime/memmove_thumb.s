// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// The ARMv7-M architecture supports unaligned MOVW, MOVH. An implementation can
// be configured to force alignment faults but no real one does this.

// Writes are often buffered so this code aligns reads and allows unaligned
// writes. Ensure CCR.UNALIGN_TRP=0.

// TODO: better performance (eg. full aligned copy using MOVM)

// func memmove(to, from unsafe.Pointer, n uintptr)
TEXT runtime·memmove(SB), NOSPLIT|NOFRAME, $0-12
	MOVW  to+0(FP), R0
	MOVW  from+4(FP), R1
	MOVW  n+8(FP), R2

	CMP  $4, R2
	BHI  more

// Fast path for 0, 1, 2, 3, 4 bytes.

	TBB    R2, R15
	HWORD  $0x0305
	HWORD  $0x0906
	HWORD  $0x0E

	MOVBU  (R1), R3  // n == 1 (TBB offset 0x03)
	MOVBU  R3, (R0)
	RET        // n == 0 (TBB offset 0x05)

	MOVHU  (R1), R3  // n == 2 (TBB offset 0x06)
	MOVHU  R3, (R0)
	RET    

	MOVHU  (R1), R3  // n == 3 (TBB offset 0x09)
	MOVBU  2(R1),R2
	MOVHU  R3, (R0)
	MOVBU  R2, 2(R0)
	RET    

	MOVW  (R1), R3  // n == 4 (TBB offset 0x0E)
	MOVW  R3, (R0)
end:
	RET  

more: // now we are sure that n > 4
	CMP  R0, R1
	BLO  backward

// Forward copy

fhead: // head copy (up to 3 bytes) to make src (R1) word aligned
	TST      $3, R1
	BEQ      fwords
	MOVBU.P  1(R1), R3
	MOVBU.P  R3, 1(R0)
	SUB      $1, R2
	B        fhead

fwords: // copy words
	SUB.S   $4, R2
	BLT     fwend
	MOVW.P  4(R1), R3
	MOVW.P  R3, 4(R0)
	B       fwords
fwend:
	ADD.S  $4, R2

ftail: // tail copy (up to 3 bytes)
	BEQ      end
	MOVBU.P  1(R1), R3
	MOVBU.P  R3, 1(R0)
	SUB.S    $1, R2
	B        ftail

// Backward copy

backward:
	ADD  R2, R0
	ADD  R2, R1

btail: // tail copy (up to 3 bytes) to make src (R1) word aligned
	TST      $3, R1
	BEQ      bwords
	MOVBU.W  -1(R1), R3
	MOVBU.W  R3, -1(R0)
	SUB      $1, R2
	B        btail

bwords: // copy words
	SUB.S   $4, R2
	BLT     bwend
	MOVW.W  -4(R1), R3
	MOVW.W  R3, -4(R0)
	B       bwords
bwend:
	ADD.S  $4, R2

bhead: // head copy (up to 3 bytes)
	BEQ      end
	MOVBU.W  -1(R1), R3
	MOVBU.W  R3, -1(R0)
	SUB.S    $1, R2
	B        bhead
