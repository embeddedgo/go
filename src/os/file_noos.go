// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"syscall"
	"time"
)

type file struct {
	name       string
	dirinfo    *dirInfo // nil unless directory being read
	appendMode bool
}

func NewFile(fd uintptr, name string) *File {
	return nil
}

func (f *File) readdir(n int) (fi []FileInfo, err error) {
	return nil, syscall.EINVAL
}

func (f *File) checkValid(op string) error {
	if f == nil {
		return ErrInvalid
	}
	return nil
}

func (f *File) read(b []byte) (n int, err error) {
	return 0, f.wrapErr("read", syscall.ENOTSUP)
}

func (f *File) pread(b []byte, off int64) (n int, err error) {
	return 0, f.wrapErr("pread", syscall.ENOTSUP)
}

func (f *File) write(b []byte) (n int, err error) {
	return 0, f.wrapErr("write", syscall.ENOTSUP)
}

func (f *File) pwrite(b []byte, off int64) (n int, err error) {
	return 0, f.wrapErr("pwrite", syscall.ENOTSUP)
}

func (f *File) seek(offset int64, whence int) (ret int64, err error) {
	return 0, f.wrapErr("seek", syscall.ENOTSUP)
}

// See docs in file.go:(*File).Chmod.
func (f *File) chmod(mode FileMode) error {
	return f.wrapErr("chmod", syscall.ENOTSUP)
}

func (f *File) setDeadline(t time.Time) error {
	return f.wrapErr("setDeadline", syscall.ENOTSUP)
}

func (f *File) setReadDeadline(t time.Time) error {
	return f.wrapErr("setDeadline", syscall.ENOTSUP)
}

func (f *File) setWriteDeadline(t time.Time) error {
	return f.wrapErr("setDeadline", syscall.ENOTSUP)
}

func (f *File) Close() error {
	return f.wrapErr("close", syscall.ENOTSUP)
}

func (f *File) Stat() (FileInfo, error) {
	return nil, f.wrapErr("setDeadline", syscall.ENOTSUP)

}

func epipecheck(file *File, e error) {}

// fixLongPath is a noop on non-Windows platforms.
func fixLongPath(path string) string {
	return path
}

func syscallMode(i FileMode) uint32 {
	return 0
}

func tempDir() string {
	return "/tmp"
}

func openFileNolog(name string, flag int, perm FileMode) (*File, error) {
	return nil, &PathError{"open", name, syscall.ENOTSUP}
}

func rename(oldname, newname string) error {
	return &PathError{"rename", oldname, syscall.ENOTSUP}
}

// See docs in file.go:Chmod.
func chmod(name string, mode FileMode) error {
	return &PathError{"chmod", name, syscall.ENOTSUP}
}

func Remove(name string) error {
	return &PathError{"remove", name, syscall.ENOTSUP}
}

type rawConn struct{}

func (c *rawConn) Control(f func(uintptr)) error {
	return syscall.ENOTSUP
}

func (c *rawConn) Read(f func(uintptr) bool) error {
	return syscall.ENOTSUP
}

func (c *rawConn) Write(f func(uintptr) bool) error {
	return syscall.ENOTSUP
}

func newRawConn(file *File) (*rawConn, error) {
	return nil, syscall.ENOTSUP
}

func hostname() (name string, err error) {
	return "", syscall.ENOTSUP
}
