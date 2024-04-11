// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import _ "unsafe"

func setPrivLevel(newlevel int) (oldlevel int, err error) {
	oldlevel, errno := runtime_setprivlevel(newlevel)
	return oldlevel, errnoError(errno)
}

//go:linkname runtime_setprivlevel runtime.setprivlevel
func runtime_setprivlevel(newlevel int) (oldlevel, errno int)
