// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package thumb

import "encoding/binary"

const (
	header = 0xffffded3
	footer = 0xab123579
)

// Item types
const (
	itemImageDef    = 0x42
	itemVectorTable = 0x03
	itemLast        = 0xff
)

// IMAGE_DEF items
const (
	imdImageTypeInvalid = 0 << 0
	imdImageTypeExe     = 1 << 0
	imdImageTypeData    = 2 << 0

	imdExeSecUnspec = 0 << 4
	imdExeSecNS     = 1 << 4
	imdExeSecS      = 2 << 4

	imdExeARM   = 0 << 8
	imdExeRISCV = 1 << 8

	imdExeRP2040 = 0 << 12
	imdExeRP2350 = 1 << 12

	imdExeTBYB = 1 << 15
)

func picoMeta() []byte {
	var imd []byte // IMAGE_DEF
	le := binary.LittleEndian
	imd = le.AppendUint32(imd, header)

	imd = append(imd, itemImageDef, 1) // IMAGE_DEF, size = 1w
	imd = le.AppendUint16(imd, imdImageTypeExe|imdExeSecS|imdExeARM|imdExeRP2350)

	imd = append(imd, itemVectorTable, 2, 0, 0) // VECTOR_TABLE, size = 2w, pad
	imd = le.AppendUint32(imd, 0x1000_0000)     // Vector table (runtime) address

	imd = append(imd, itemLast)
	imd = le.AppendUint16(imd, uint16((len(imd)-4-1)/4)) // other items' size
	imd = append(imd, 0)                                 // pad

	imd = le.AppendUint32(imd, 0) // link to the next block relative to header

	imd = le.AppendUint32(imd, footer)

	// For now we only produce IMAGE_DEF.
	return imd
}
