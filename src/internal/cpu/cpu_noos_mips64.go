// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

const CacheLinePadSize = 32 // Must be greatest of all supported CPUs
const CacheLineMinSize = 16 // Must be smallest of all supported CPUs

func doinit() {
	// noos currently will always use MIPS-III ISA.  Once more mips64
	// targets are supported, we should probe here for ISA extensions and
	// enable them.

	MIPS64X.HasMSA = false
}
