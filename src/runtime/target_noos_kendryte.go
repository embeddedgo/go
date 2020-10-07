// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build noos
// +build k210

package runtime

const (
	_OS                        = 0
	noos                       = true
	noosStackCacheSize         = 16 * 1024
	noosNumStackOrders         = 3
	noosHeapAddrBits           = 23 // enough for 8 MiB K210 SRAM
	noosLogHeapArenaBytes      = 17 // 128 KiB
	noosArenaBaseOffset        = 0x80000000
	noosMinPhysPageSize        = 256
	noosFinBlockSize           = 2 * 1024
	noosSweepMinHeapDistance   = 8 * 1024
	noosDefaultHeapMinimum     = 64 * 1024
	noosHeapGoalInc            = 8 * 1024
	noosGCSweepBlockEntries    = 1024
	noosGCSweepBufInitSpineCap = 64
	noosGCBitsChunkBytes       = 16 * 1024
)
