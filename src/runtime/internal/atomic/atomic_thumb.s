// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·Load(SB),NOSPLIT,$0-8
	MOVW  addr+0(FP), R0
	MOVW  (R0), R1
	DMB   MB_ISH
	MOVW  R1, ret+4(FP)
	RET

TEXT ·Load8(SB),NOSPLIT,$0-5
	MOVW   addr+0(FP), R0
	MOVBU  (R0), R1
	DMB    MB_ISH
	MOVB   R1, ret+4(FP)
	RET

TEXT ·Store(SB),NOSPLIT,$0-8
	MOVW  addr+0(FP), R1
	MOVW  v+4(FP), R2
	DMB   MB_ISH
	MOVW  R2, (R1)
	DMB   MB_ISH
	RET

TEXT ·Store8(SB),NOSPLIT,$0-5
	MOVW   addr+0(FP), R1
	MOVBU  v+4(FP), R2
	DMB    MB_ISH
	MOVB   R2, (R1)
	DMB    MB_ISH
	RET

TEXT ·Cas(SB),NOSPLIT|NOFRAME,$0
	MOVW  ptr+0(FP), R1
	MOVW  old+4(FP), R2
	MOVW  new+8(FP), R3
casl:
	LDREX  (R1), R0
	CMP    R0, R2
	BNE    end
	DMB    MB_ISHST
	STREX  R3, (R1), R0
	CMP    $0, R0
	BNE    casl
end:
	MOVW.NE    $0, R0
	MOVW.P.EQ  $1, R0
	DMB.EQ     MB_ISH
	MOVB       R0, ret+12(FP)
	RET

TEXT ·Xadd(SB),NOSPLIT|NOFRAME,$0-12
	MOVW  val+0(FP), R1
	MOVW  delta+4(FP), R2
loop:
	LDREX  (R1), R0
	DMB    MB_ISHST
	ADD    R2, R0
	STREX  R0, (R1), R3
	CMP    $0, R3
	BNE    loop
	DMB    MB_ISH
	MOVW   R0, ret+8(FP)
	RET

TEXT ·Or(SB),NOSPLIT|NOFRAME,$0-8
	MOVW  addr+0(FP), R1
	MOVW  v+4(FP), R2
loop:
	LDREX  (R1), R0
	DMB    MB_ISHST
	ORR    R2, R0
	STREX  R0, (R1), R3
	CMP    $0, R3
	BNE    loop
	DMB    MB_ISH
	RET

TEXT ·And(SB),NOSPLIT|NOFRAME,$0-8
	MOVW  addr+0(FP), R1
	MOVW  v+4(FP), R2
loop:
	LDREX  (R1), R0
	DMB    MB_ISHST
	AND    R2, R0
	STREX  R0, (R1), R3
	CMP    $0, R3
	BNE    loop
	DMB    MB_ISH
	RET

// stubs

TEXT ·Loadp(SB),NOSPLIT|NOFRAME,$0-8
	B   ·Load(SB)

TEXT ·LoadAcq(SB),NOSPLIT|NOFRAME,$0-8
	B   ·Load(SB)

TEXT ·LoadAcquintptr(SB),NOSPLIT|NOFRAME,$0-8
	B 	·Load(SB)

TEXT ·Casint32(SB),NOSPLIT,$0-13
	B	·Cas(SB)

TEXT ·Casint64(SB),NOSPLIT,$-4-21
	B	·Cas64(SB)

TEXT ·Casuintptr(SB),NOSPLIT,$0-13
	B	·Cas(SB)

TEXT ·Casp1(SB),NOSPLIT,$0-13
	B   ·Cas(SB)

TEXT ·CasRel(SB),NOSPLIT,$0-13
	B   ·Cas(SB)

TEXT ·Loadint32(SB),NOSPLIT,$0-8
	B	·Load(SB)

TEXT ·Loadint64(SB),NOSPLIT,$-4-12
	B	·Load64(SB)

TEXT ·Loaduintptr(SB),NOSPLIT,$0-8
	B	·Load(SB)

TEXT ·Loaduint(SB),NOSPLIT,$0-8
	B	·Load(SB)

TEXT ·Storeint32(SB),NOSPLIT,$0-8
	B	·Store(SB)

TEXT ·Storeint64(SB),NOSPLIT,$0-12
	B	·Store64(SB)

TEXT ·Storeuintptr(SB),NOSPLIT,$0-8
	B	·Store(SB)

TEXT ·StorepNoWB(SB),NOSPLIT,$0-8
	B   ·Store(SB)

TEXT ·StoreRel(SB),NOSPLIT,$0-8
	B   ·Store(SB)

TEXT ·StoreReluintptr(SB),NOSPLIT,$0-8
	B	·Store(SB)

TEXT ·Xaddint32(SB),NOSPLIT,$0-12
	B	·Xadd(SB)

TEXT ·Xaddint64(SB),NOSPLIT,$-4-20
	B	·Xadd64(SB)

TEXT ·Xadduintptr(SB),NOSPLIT,$0-12
	B	·Xadd(SB)

TEXT ·Xchgint32(SB),NOSPLIT,$0-12
	B	·Xchg(SB)

TEXT ·Xchgint64(SB),NOSPLIT,$-4-20
	B	·Xchg64(SB)

TEXT ·Cas64(SB),NOSPLIT,$0-21
	B   ·goCas64(SB)

TEXT ·Xadd64(SB),NOSPLIT,$0-20
	B   ·goXadd64(SB)

TEXT ·Xchg64(SB),NOSPLIT,$0-20
	B   ·goXchg64(SB)

TEXT ·Load64(SB),NOSPLIT,$0-12
	B   ·goLoad64(SB)

TEXT ·Store64(SB),NOSPLIT,$0-12
	B   ·goStore64(SB)
