// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

import "internal/goarch"

// Map constants common to several packages
// runtime/runtime-gdb.py:MapTypePrinter contains its own copy
const (
	MapBucketCountBits = 3 // log2 of number of elements in a bucket.
	MapBucketCount     = 1 << MapBucketCountBits
	MapMaxKeyBytes     = 128 // Must fit in a uint8.
	MapMaxElemBytes    = 128 // Must fit in a uint8.
)

// ZeroValSize is the size in bytes of runtime.zeroVal.
const ZeroValSize = 512 << (goarch.PtrSize / 64) // Embedded Go: 512 on 32-bit, 1024 on 64-bit architectures
