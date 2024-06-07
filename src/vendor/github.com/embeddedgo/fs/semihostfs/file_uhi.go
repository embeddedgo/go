// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips64

package semihostfs

import (
	"io"
	"io/fs"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"
)

type file struct {
	name   string
	fd     int
	closed func()
}

func (f *file) Close() (err error) {
	if f.name == "" {
		return &fs.PathError{Op: "close", Path: f.name, Err: syscall.EBADF}
	}
	errno := hostCall(3, uintptr(f.fd), 0, 0, nil)
	if errno < 0 {
		err = &fs.PathError{Op: "close", Path: f.name, Err: &Error{errno}}
	}
	if f.closed != nil {
		f.closed()
		f.closed = nil
	}
	f.name = ""
	return
}

func (f *file) Read(p []byte) (n int, err error) {
	if f.name == "" {
		return 0, &fs.PathError{Op: "read", Path: f.name, Err: syscall.EBADF}
	}
	if len(p) == 0 {
		return
	}
	ptr := &p[0]
	ne := hostCall(
		4,
		uintptr(f.fd),
		uintptr(unsafe.Pointer(ptr)),
		uintptr(len(p)),
		ptr,
	)
	switch {
	case ne == 0:
		err = io.EOF
	case n < 0:
		err = &fs.PathError{Op: "read", Path: f.name, Err: &Error{ne}}
	default:
		n = ne
	}
	return
}

func (f *file) WriteString(s string) (n int, err error) {
	if f.name == "" {
		return 0, &fs.PathError{Op: "write", Path: f.name, Err: syscall.EBADF}
	}
	if len(s) == 0 {
		return
	}
	ptr := unsafe.StringData(s)
	ne := hostCall(
		5,
		uintptr(f.fd),
		uintptr(unsafe.Pointer(ptr)),
		uintptr(len(s)),
		ptr,
	)
	switch {
	case n < 0:
		err = &fs.PathError{Op: "write", Path: f.name, Err: &Error{ne}}
	default:
		n = ne
	}
	return
}

func (f *file) Write(p []byte) (int, error) {
	return f.WriteString(*(*string)(unsafe.Pointer(&p)))
}

func (f *file) Stat() (fi fs.FileInfo, err error) {
	if f.name == "" {
		return nil, &fs.PathError{Op: "stat", Path: f.name, Err: syscall.EBADF}
	}
	var info fileInfo
	ptr := unsafe.Pointer(&info)
	errno := hostCall(8, uintptr(f.fd), uintptr(ptr), 0, (*byte)(ptr))
	if errno < 0 {
		err = &fs.PathError{Op: "stat", Path: f.name, Err: &Error{errno}}
		return
	}
	info.name = filepath.Base(f.name)
	fi = &info
	return
}

type fileInfo struct {
	dev     int16
	ino     uint16
	mode    uint32
	nlink   uint16
	uid     uint16
	gid     uint16
	rdev    int16
	size    uint64
	atime   uint64
	spare1  uint64
	mtime   uint64
	spare2  uint64
	ctime   uint64
	spare3  uint64
	blksize uint64
	blocks  uint64
	spare4  [2]uint64

	name string // not from UHI call
}

func (fi *fileInfo) Name() string       { return fi.name }
func (fi *fileInfo) Size() int64        { return int64(fi.size) }
func (fi *fileInfo) Mode() fs.FileMode  { return 0666 }
func (fi *fileInfo) ModTime() time.Time { return time.Time{} } // TODO
func (fi *fileInfo) IsDir() bool        { return false }       // TODO
func (fi *fileInfo) Sys() any           { return nil }         // TODO
