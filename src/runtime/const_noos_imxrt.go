// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build noos
// +build imxrt1050 imxrt1060

package runtime

const (
	_OS                        = 0
	noos                       = true
	noosScaleDown              = 16
	noosStackCacheSize         = 8 * 1024
	noosNumStackOrders         = 2
	noosHeapAddrBits           = 20         // enough for 1M SRAM
	noosLogHeapArenaBytes      = 14         // 16 KiB
	noosArenaBaseOffset        = 0x20000000 // the begginning of SRAM
	noosMinPhysPageSize        = 256
	noosSpanSetInitSpineCap    = 8
	noosStackMin               = 1024
	noosStackSystem            = 27 * 4 // register stacking at exception entry
	noosStackGuard             = 464
	noosFinBlockSize           = 256
	noosSweepMinHeapDistance   = 1024
	noosDefaultHeapMinimum     = 8 * 1024
	noosHeapGoalInc            = 1024
	noosGCSweepBlockEntries    = 64
	noosGCSweepBufInitSpineCap = 32
	noosGCBitsChunkBytes       = 2 * 1024
	noosSemTabSize             = 31
)
