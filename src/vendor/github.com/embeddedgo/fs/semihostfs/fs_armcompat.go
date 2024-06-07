// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build riscv64 || thumb

package semihostfs

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

func openWithFinalizer(fsys *FS, name string, flag int, _ fs.FileMode, closed func()) (f fs.File, err error) {
	mode := 0

	switch flag & (syscall.O_RDONLY | syscall.O_RDWR | syscall.O_WRONLY) {
	case syscall.O_RDONLY:
		// rb: open binary file for reading from the beggining
		mode = 1
	case syscall.O_RDWR:
		switch flag & (syscall.O_CREAT | syscall.O_TRUNC | syscall.O_APPEND) {
		case 0:
			// r+b: open binary file for read/writing at the beggining
			mode = 3
		case syscall.O_CREAT | syscall.O_TRUNC:
			// w+b: truncate or create binary file for writing and reading
			mode = 7
		case syscall.O_CREAT | syscall.O_APPEND:
			// a+b: open or create binary file for appending and reading
			mode = 11
		}
	case syscall.O_WRONLY:
		switch flag & (syscall.O_CREAT | syscall.O_TRUNC | syscall.O_APPEND) {
		case syscall.O_CREAT | syscall.O_TRUNC:
			// wb: truncate or create binary file for writing
			mode = 5
		case syscall.O_CREAT | syscall.O_APPEND:
			// ab: open or create text file for appending
			mode = 9
		}
	default:
		return nil, syscall.ENOTSUP
	}
	hostPath := ":tt"
	switch name {
	case ":stderr":
		mode = 8
	case ":stdout":
		mode = 4
	case ":stdin":
		mode = 0
	default:
		hostPath = filepath.Join(fsys.root, name)
	}
	type args struct {
		path    *byte
		mode    int
		pathLen int
	}
	ptr := unsafe.Pointer(&args{
		unsafe.StringData(hostPath + "\x00"),
		mode,
		len(hostPath),
	})
	mt.Lock()
	fd := hostCall(0x01, uintptr(ptr), ptr)
	if fd == -1 {
		err = hostError()
	}
	mt.Unlock()
	if fd == -1 {
		err = &fs.PathError{Op: "open", Path: name, Err: err}
	} else {
		f = &file{name, fd, closed}
	}
	return
}

func mkdir(fsys *FS, name string, mode fs.FileMode) error {
	return syscall.ENOTSUP
}

func remove(fsys *FS, name string) error {
	type args struct {
		path    *byte
		pathLen int
	}
	hostPath := filepath.Join(fsys.root, name)
	ptr := unsafe.Pointer(&args{
		unsafe.StringData(hostPath + "\x00"),
		len(hostPath),
	})
	mt.Lock()
	errno := hostCall(0x0e, uintptr(ptr), ptr)
	mt.Unlock()
	if errno != 0 {
		return &fs.PathError{Op: "remove", Path: name, Err: &Error{errno}}
	}
	return nil
}

func rename(fsys *FS, oldname, newname string) error {
	type args struct {
		oldName *byte
		oldLen  int
		newName *byte
		newLen  int
	}
	hostOld := filepath.Join(fsys.root, oldname)
	hostNew := filepath.Join(fsys.root, newname)
	ptr := unsafe.Pointer(&args{
		unsafe.StringData(hostOld + "\x00"),
		len(hostOld),
		unsafe.StringData(hostNew + "\x00"),
		len(hostNew),
	})
	mt.Lock()
	errno := hostCall(0x0f, uintptr(ptr), ptr)
	mt.Unlock()
	if errno != 0 {
		return &fs.PathError{Op: "rename", Path: oldname, Err: &Error{errno}}
	}
	return nil
}

func init() {
	if len(os.Args) != 1 || os.Args[0] != "" {
		return
	}
	var buf [256]byte
	args := buf[:]
	ptr := unsafe.Pointer(&args)
	mt.Lock()
	e := hostCall(0x15, uintptr(ptr), ptr)
	mt.Unlock()
	if e == 0 {
		os.Args = strings.Fields(string(args))
	}
}
