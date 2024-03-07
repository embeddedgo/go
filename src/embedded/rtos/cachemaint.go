// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import "unsafe"

const (
	DCacheInval      = 0
	DCacheFlush      = 1
	DCacheFlushInval = 2

	DCacheClean      = DCacheFlush      // deprecated
	DCacheCleanInval = DCacheFlushInval // deprecated
)

// CacheMaint performs a cache maintenance operations specified by op. This
// operation applies to all cache lines that contain data from the memory
// described by p and size. For this reason, the adjacent memory that shares
// the same cache lines will be affected too. You must take this into account
// if there are other bus masters (e.g. DMA) in the system that may use this
// adjacent memory.
func CacheMaint(op int, p unsafe.Pointer, size int) {
	cacheMaint(op, p, size)
}
