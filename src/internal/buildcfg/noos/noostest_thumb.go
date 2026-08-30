// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: Do not use GOTARGET tags in runtime. Come up with something smarter
// based on some kind of memory description similar to the -M linker option.

//go:build noos && thumb && noostest

package noos

// see ../../../../noos_thumb.env

const (
	OS                          = 0
	ScaleDown                   = 2 // must be power of 2
	StackCacheSize              = 16 * 1024
	NumStackOrders              = 3
	HeapAddrBits                = 24         // enough for 16MiB PSRAM
	LogHeapArenaBytes           = 17         // 128 KiB
	ArenaBaseOffset             = 0x60000000 // the begginning of PSRAM
	MinPhysPageSize             = 256
	SpanSetInitSpineCap         = 64
	StackMin                    = 2048
	StackSystem                 = 27 * 4 // register stacking at exception entry
	StackGuard                  = 464
	SweepMinHeapDistance        = 8 * 1024
	DefaultHeapMinimum          = 64 * 1024
	MemoryLimitHeapGoalHeadroom = 1 << 18
	GCSweepBlockEntries         = 1024
	GCSweepBufInitSpineCap      = 64
	GCBitsChunkBytes            = 16 * 1024
	SemTabSize                  = 113
	GOGC                        = 50
	TimeHistMaxBucketBits       = 47
	MPaddedSize                 = 2048
)
