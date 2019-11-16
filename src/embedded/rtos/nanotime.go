// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import _ "unsafe"

// Nanotime returns monotonic time in nanosecond.
func Nanotime() int64 { return runtime_nanotime() }

//go:linkname runtime_nanotime runtime.nanotime
func runtime_nanotime() int64
