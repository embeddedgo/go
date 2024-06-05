// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build riscv64 || thumb

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
	ptr := unsafe.Pointer(&f.fd)
	mt.Lock()
	if hostCall(0x02, uintptr(ptr), ptr) == -1 {
		err = &fs.PathError{Op: "close", Path: f.name, Err: hostError()}
	}
	mt.Unlock()
	if f.closed != nil {
		f.closed()
		f.closed = nil
	}
	if f.name == ":stderr" {
		const (
			sysExitApplicationExit = 0x20026 // graceful exit
			ptrSize                = 4 << (^uintptr(0) >> 63)
		)
		if ptrSize == 32 {
			hostCall(0x18, sysExitApplicationExit, nil)
		} else {
			ptr := unsafe.Pointer(&[2]int{sysExitApplicationExit, 0})
			hostCall(0x18, uintptr(ptr), ptr)
		}
	}
	f.name = ""
	return
}

type rwargs struct {
	fd int
	p  *byte
	n  int
}

func (f *file) Read(p []byte) (n int, err error) {
	if f.name == "" {
		return 0, &fs.PathError{Op: "read", Path: f.name, Err: syscall.EBADF}
	}
	if len(p) == 0 {
		return
	}
	ptr := unsafe.Pointer(&rwargs{
		f.fd,
		unsafe.SliceData(p),
		len(p),
	})

	mt.Lock()
	notRead := hostCall(0x06, uintptr(ptr), ptr)
	mt.Unlock()
	n = len(p) - notRead
	if n == 0 {
		err = io.EOF
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
	ptr := unsafe.Pointer(&rwargs{
		f.fd,
		unsafe.StringData(s),
		len(s),
	})
	mt.Lock()
	notWritten := hostCall(0x05, uintptr(ptr), ptr)
	if notWritten != 0 {
		err = hostError()
	}
	mt.Unlock()
	if notWritten != 0 {
		err = &fs.PathError{Op: "write", Path: f.name, Err: err}
	}
	n = len(s) - notWritten
	return

}

func (f *file) Write(p []byte) (int, error) {
	return f.WriteString(*(*string)(unsafe.Pointer(&p)))
}

type fileInfo struct {
	name string
	size int
}

func (f *file) Stat() (fi fs.FileInfo, err error) {
	if f.name == "" {
		return nil, &fs.PathError{Op: "stat", Path: f.name, Err: syscall.EBADF}
	}
	ptr := unsafe.Pointer(&f.fd)
	mt.Lock()
	size := hostCall(0x0c, uintptr(ptr), ptr)
	if size == -1 {
		err = hostError()
	}
	mt.Unlock()
	if size == -1 {
		err = &fs.PathError{Op: "stat", Path: f.name, Err: err}
	} else {
		fi = &fileInfo{
			filepath.Base(f.name),
			size,
		}
	}
	return
}

func (fi *fileInfo) Name() string       { return fi.name }
func (fi *fileInfo) Size() int64        { return int64(fi.size) }
func (fi *fileInfo) Mode() fs.FileMode  { return 0666 }
func (fi *fileInfo) ModTime() time.Time { return time.Time{} }
func (fi *fileInfo) IsDir() bool        { return false }
func (fi *fileInfo) Sys() any           { return nil }
