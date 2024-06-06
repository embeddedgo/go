// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !mips64 && !riscv64 && !thumb

package semihostfs

import (
	"errors"
	"io/fs"
)

var errNotSupported = errors.New("not supported")

func openWithFinalizer(fsys *FS, name string, flag int, _ fs.FileMode, closed func()) (f fs.File, err error) {
	return nil, errNotSupported
}

func mkdir(fsys *FS, name string, mode fs.FileMode) error {
	return errNotSupported
}

func remove(fsys *FS, name string) error {
	return errNotSupported
}

func rename(fsys *FS, oldname, newname string) error {
	return errNotSupported
}
