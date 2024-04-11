// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noos

package rtos

import "unsafe"

func reset(level int, addr unsafe.Pointer) bool {
	return false
}
