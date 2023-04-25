// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import "unsafe"

//go:linkname reset runtime.reset
func reset(level int, addr unsafe.Pointer) bool
