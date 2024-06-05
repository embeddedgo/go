// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build riscv64 || thumb

package semihostfs

import (
	"fmt"
	"sync"
	"unsafe"
)

var mt sync.Mutex // for hostCall, hostError pair

//go:noescape
func hostCall(cmd int, arg uintptr, avoidGC unsafe.Pointer) int

type Error struct {
	no int
}

func (err *Error) Error() string {
	return fmt.Sprint("semihosting error: ", err.no)
}

func hostError() *Error {
	return &Error{hostCall(0x13, 0, nil)}
}
