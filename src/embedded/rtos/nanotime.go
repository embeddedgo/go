// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import (
	"time"
	_ "unsafe"
)

// Nanotime returns a time duration from some event in the past. Typically it's
// roughly corresponds to the system runtime.
func Nanotime() time.Duration {
	return time.Duration(runtime_nanotime())
}

//go:linkname runtime_nanotime runtime.nanotime
func runtime_nanotime() int64
