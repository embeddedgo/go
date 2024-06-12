// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import "unsafe"

// copied from /home/michal/P/go/goroot/src/runtime/stubs.go
func bool2int(x bool) int { return int(uint8(*(*uint8)(unsafe.Pointer(&x)))) }
