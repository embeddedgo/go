// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: Do not use GOTARGET tags in runtime. Come up with something smarter
// based on some kind of memory description similar to the -M linker option.

//go:build noos && stm32h7x3

package noos

const (
	OS                          = 0
	ScaleDown                   = 8 // must be power of 2
	StackCacheSize              = 8 * 1024
	NumStackOrders              = 2
	HeapAddrBits                = 19         // enough for 512 KiB AXI SRAM (AHB SRAM not supported)
	LogHeapArenaBytes           = 16         // 64 KiB
	ArenaBaseOffset             = 0x24000000 // the begginning of AXI SRAM
	MinPhysPageSize             = 256
	SpanSetInitSpineCap         = 8
	StackMin                    = 1024
	StackSystem                 = 27 * 4 // register stacking at exception entry
	StackGuard                  = 464
	SweepMinHeapDistance        = 1024
	DefaultHeapMinimum          = 8 * 1024
	MemoryLimitHeapGoalHeadroom = 1 << 15
	GCSweepBlockEntries         = 64
	GCSweepBufInitSpineCap      = 32
	GCBitsChunkBytes            = 2 * 1024
	SemTabSize                  = 31
	GOGC                        = 30
	TimeHistMaxBucketBits       = 46
	MPaddedSize                 = 2048
)
