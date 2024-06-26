// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: Do not use GOTARGET tags in runtime. Come up with something smarter
// based on some kind of memory description similar to the -M linker option.

//go:build noos && thumb && noostest

package runtime

// qemu-system-arm -machine mps2-an500

const (
	_OS                             = 0
	noos                            = true
	noosScaleDown                   = 2 // must be power of 2
	noosStackCacheSize              = 16 * 1024
	noosNumStackOrders              = 3
	noosHeapAddrBits                = 24         // enough for 16MiB PSRAM
	noosLogHeapArenaBytes           = 17         // 128 KiB
	noosArenaBaseOffset             = 0x60000000 // the begginning of PSRAM
	noosMinPhysPageSize             = 256
	noosSpanSetInitSpineCap         = 64
	noosStackMin                    = 2048
	noosStackSystem                 = 27 * 4 // register stacking at exception entry
	noosStackGuard                  = 464
	noosFinBlockSize                = 2 * 1024
	noosSweepMinHeapDistance        = 8 * 1024
	noosDefaultHeapMinimum          = 64 * 1024
	noosMemoryLimitHeapGoalHeadroom = 1 << 18
	noosGCSweepBlockEntries         = 1024
	noosGCSweepBufInitSpineCap      = 64
	noosGCBitsChunkBytes            = 16 * 1024
	noosSemTabSize                  = 113
	noosGOGC                        = 50
	noosTimeHistMaxBucketBits       = 47
)
