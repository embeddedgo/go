// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import "unsafe"

// CacheMaint performs a cache maintenance operations specified by the
// architecture dependent parameter op. The operation may affect an area of
// memory larger than that specified by the p and size parameters.
//
// The meaning of "op" for different architectures.
//
// thumb:
//  0. data cache invalidate to the point of coherency
//  1. data cache clean to the point of coherency
//  2. data cache clean and invalidate to the point of coherency
//
func CacheMaint(op int, p unsafe.Pointer, size int) {
	cacheMaint(op, p, size)
}
