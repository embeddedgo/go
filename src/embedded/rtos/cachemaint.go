// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import "unsafe"

// CacheMaint performs a cache maintenance operations specified by the
// architecture dependent parameter op. The operation may affect an area of
// memory larger than that specified by the p and size parameters.

// CacheMaint op for ARMv7-M
const (
	ARMv7M_DCacheInvalPoC      = 0
	ARMv7M_DCacheCleanPoC      = 1
	ARMv7M_DCacheCleanInvalPoC = 2
)

func CacheMaint(op int, p unsafe.Pointer, size int) {
	cacheMaint(op, p, size)
}
