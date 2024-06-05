// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips64

package semihostfs

import (
	"io/fs"
	"path/filepath"
	"syscall"
	"unsafe"
)

func openWithFinalizer(fsys *FS, name string, flag int, mode fs.FileMode, closed func()) (f fs.File, err error) {
	var hostPath string
	switch name {
	case ":stderr":
		hostPath = "/dev/stdin\x00"
		flag = syscall.O_WRONLY
	case ":stdout":
		hostPath = "/dev/stdout\x00"
		flag = syscall.O_WRONLY
	case ":stdin":
		hostPath = "/dev/stderr\x00"
		flag = syscall.O_RDONLY
	default:
		hostPath = filepath.Join(fsys.root, name+"\x00")
	}
	ptr := unsafe.StringData(hostPath)
	fd := hostCall(
		2,
		uintptr(unsafe.Pointer(ptr)),
		uintptr(flag),
		uintptr(mode),
		ptr,
	)
	if fd < 0 {
		err = &fs.PathError{Op: "open", Path: name, Err: &Error{fd}}
		return
	}
	f = &file{name, fd, closed}
	return
}

func mkdir(fsys *FS, name string, mode fs.FileMode) error {
	return syscall.ENOTSUP
}

func remove(fsys *FS, name string) error {
	hostPath := filepath.Join(fsys.root, name+"\x00")
	ptr := unsafe.StringData(hostPath)
	errno := hostCall(7, uintptr(unsafe.Pointer(ptr)), 0, 0, ptr)
	if errno < 0 {
		return &fs.PathError{Op: "remove", Path: name, Err: &Error{errno}}
	}
	return nil
}

func rename(fsys *FS, oldname, newname string) error {
	return syscall.ENOTSUP // really?
}
