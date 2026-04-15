// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var (
	// Constants that we use and will transfer to the runtime.
	minHeapAlign = 8
	maxSmallSize = 32 << 10
	smallSizeDiv = 8
	smallSizeMax = 1024
	largeSizeDiv = 128
	pageShift    = 13
	tinySize     = 16

	// Derived constants.
	pageSize = 1 << pageShift
)

var (
	maxPtrSize = uintptr(max(4, 8))
	maxPtrBits = 8 * maxPtrSize

	// Maximum size smallScanNoHeader would be called for, which is the
	// maximum value gc.MinSizeForMallocHeader can have on any platform.
	// gc.MinSizeForMallocHeader is defined as goarch.PtrSize * goarch.PtrBits.
	smallScanNoHeaderMax = maxPtrSize * maxPtrBits
)

var build = "!noos"

func constsNoos() {
	switch *noos {
	case 32:
		maxSmallSize = 1 << 9
		smallSizeMax = 256
		pageShift    = 9
		build = "noos && thumb"
	case 64:
		maxSmallSize = 1 << 12
		smallSizeMax = 512
		pageShift    = 11
		build = "noss && (riscv64 || mips64)"
	}
	pageSize = 1 << pageShift
}