// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

import "syscall"

func open(name string) (uintptr, error) {
	return 0, syscall.ENOENT
}

func closefd(fd uintptr) {
}

func read(fd uintptr, buf []byte) (int, error) {
	return 0, syscall.EINVAL
}

func preadn(fd uintptr, buf []byte, off int) error {
	return syscall.EINVAL
}
