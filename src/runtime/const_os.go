// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noos
// +build !noos

package runtime

const (
	_OS                             = 1
	noos                            = false
	noosScaleDown                   = 1
	noosStackCacheSize              = 0
	noosNumStackOrders              = 0
	noosHeapAddrBits                = 0
	noosLogHeapArenaBytes           = 0
	noosArenaBaseOffset             = 0
	noosMinPhysPageSize             = 0
	noosSpanSetInitSpineCap         = 0
	noosStackMin                    = 0
	noosStackSystem                 = 0
	noosStackGuard                  = 0
	noosFinBlockSize                = 0
	noosSweepMinHeapDistance        = 0
	noosDefaultHeapMinimum          = 0
	noosMemoryLimitHeapGoalHeadroom = 0
	noosGCSweepBlockEntries         = 0
	noosGCSweepBufInitSpineCap      = 0
	noosGCBitsChunkBytes            = 0
	noosSemTabSize                  = 0
	noosGOGC                        = 0
)
