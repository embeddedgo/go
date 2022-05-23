// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: Do not use GOTARGET tags in runtime. Come up with something smarter
// based on some kind of memory description similar to the -M linker option.

// +build noos
// +build k210

package runtime

const (
	_OS                        = 0
	noos                       = true
	noosScaleDown              = 2
	noosStackCacheSize         = 16 * 1024
	noosNumStackOrders         = 3
	noosHeapAddrBits           = 23 // enough for 8 MiB K210 SRAM
	noosLogHeapArenaBytes      = 17 // 128 KiB
	noosArenaBaseOffset        = 0x80000000
	noosMinPhysPageSize        = 256
	noosSpanSetInitSpineCap    = 64
	noosStackMin               = 2048
	noosStackSystem            = 0
	noosStackGuard             = 928
	noosFinBlockSize           = 2 * 1024
	noosSweepMinHeapDistance   = 8 * 1024
	noosDefaultHeapMinimum     = 64 * 1024
	noosHeapGoalInc            = 8 * 1024
	noosGCSweepBlockEntries    = 1024
	noosGCSweepBufInitSpineCap = 64
	noosGCBitsChunkBytes       = 16 * 1024
	noosSemTabSize             = 113
)
