// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import "unsafe"

// CacheMaint performs a cache maintenance operations specified by the
// architecture dependent parameter op. The operation may affect the memory
// area bigger than specified by the p and size parameters.
func CacheMaint(op int, p unsafe.Pointer, size int) {
	cacheMaint(op, p, size)
}
