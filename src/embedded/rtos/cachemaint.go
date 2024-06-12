// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import "unsafe"


const (
	DCacheInval      = 0
	DCacheClean      = 1
	DCacheCleanInval = 2
)

// CacheMaint performs a cache maintenance operations specified by op.
func CacheMaint(op int, p unsafe.Pointer, size int) {
	cacheMaint(op, p, size)
}
