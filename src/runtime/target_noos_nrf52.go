// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build noos
// +build nrf52840

package runtime

const (
	_OS                        = 0
	noos                       = true
	noosStackCacheSize         = 4 * 1024
	noosNumStackOrders         = 2
	noosHeapAddrBits           = 18 // enough for 256 KiB of nRF52840
	noosLogHeapArenaBytes      = 14 // 16 KiB
	noosArenaBaseOffset        = 0x20000000
	noosFinBlockSize           = 256
	noosSweepMinHeapDistance   = 1024
	noosDefaultHeapMinimum     = 8 * 1024
	noosHeapGoalInc            = 1024
	noosGCSweepBlockEntries    = 64
	noosGCSweepBufInitSpineCap = 32
	noosGCBitsChunkBytes       = 2
)
