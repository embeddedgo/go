// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semihostfs

import "io/fs"

// An FS represents a semihosting file system.
type FS struct {
	name string
	root string
}

func New(name, rootDir string) *FS {
	return &FS{name, rootDir}
}

// OpenWithFinalizer implements the rtos.FS OpenWithFinalizer method.
func (fsys *FS) OpenWithFinalizer(name string, flag int, mode fs.FileMode, closed func()) (fs.File, error) {
	return openWithFinalizer(fsys, name, flag, mode, closed)
}

// Open implements the fs.FS Open method.
func (fsys *FS) Open(name string) (fs.File, error) {
	return fsys.OpenWithFinalizer(name, 0, 0, nil)
}

// Type implements the rtos.FS Type method.
func (fsys *FS) Type() string { return "semihost" }

// Name implements the rtos.FS Name method.
func (fsys *FS) Name() string { return fsys.name }

// Usage implements the rtos.FS Usage method.
func (fsys *FS) Usage() (usedItems, maxItems int, usedBytes, maxBytes int64) {
	return -1, -1, -1, -1
}

// Mkdir implements the optional rtos.FS method.
func (fsys *FS) Mkdir(name string, mode fs.FileMode) error {
	return mkdir(fsys, name, mode)
}

// Remove implements the optional rtos.FS method.
func (fsys *FS) Remove(name string) error {
	return remove(fsys, name)
}

// Rename implements the optional rtos.FS method.
func (fsys *FS) Rename(oldname, newname string) error {
	return rename(fsys, oldname, newname)
}
