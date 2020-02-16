// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import "syscall"

func statNolog(name string) (FileInfo, error) {
	return nil, &PathError{"stat", name, syscall.ENOTSUP}
}

func lstatNolog(name string) (FileInfo, error) {
	return nil, &PathError{"lstat", name, syscall.ENOTSUP}
}
