// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips64

package semihostfs

import (
	"fmt"
)

//go:noescape
func hostCall(cmd int, a0, a1, a2 uintptr, avoidGC *byte) int

type Error struct {
	no int
}

func (err *Error) Error() string {
	return fmt.Sprint("semihosting error: ", err.no) // TODO: decode errno
}
