// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !noos

package rtos

import "unsafe"

func cacheMaint(op int, p unsafe.Pointer, size int) {
	return -1, ErrNotSuppoted
}
