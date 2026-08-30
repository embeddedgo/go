// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: Do not use GOTARGET tags in runtime. Come up with something smarter
// based on some kind of memory description similar to the -M linker option.

//go:build noos && riscv64 && noostest

package noos

// see ../../../../noos_riscv64.env

const (
	OS                          = 0
	ScaleDown                   = 2 // must be power of 2
	StackCacheSize              = 16 * 1024
	NumStackOrders              = 3
	HeapAddrBits                = 25         // enough for 32 MiB RAM
	LogHeapArenaBytes           = 18         // 256 KiB
	ArenaBaseOffset             = 0x80000000 // the beginning of RAM
	MinPhysPageSize             = 256
	SpanSetInitSpineCap         = 64
	StackMin                    = 2048
	StackSystem                 = 0
	StackGuard                  = 928
	SweepMinHeapDistance        = 8 * 1024
	DefaultHeapMinimum          = 64 * 1024
	MemoryLimitHeapGoalHeadroom = 1 << 18
	GCSweepBlockEntries         = 1024
	GCSweepBufInitSpineCap      = 64
	GCBitsChunkBytes            = 16 * 1024
	SemTabSize                  = 113
	GOGC                        = 50
	TimeHistMaxBucketBits       = 47
	MPaddedSize                 = 4096
)
