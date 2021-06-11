// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	_ "embedded/rtos"
	"io"
	"io/fs"
	"syscall"
	"time"
)

func tempDir() string {
	return "/tmp"
}

type file struct {
	f          fs.File
	name       string
	dirinfo    *dirInfo // nil unless directory being read
	appendMode bool
}

func openFileNolog(name string, flag int, perm FileMode) (*File, error) {
	f, err := openFile(name, flag, perm)
	if err != nil {
		return nil, &PathError{Op: "open", Path: name, Err: err}
	}
	return &File{&file{f: f, name: name}}, nil
}

// NewFile is not supported by GOOS=noos.
func NewFile(fd uintptr, name string) *File {
	return nil
}

func (f *File) readdir(n int, mode readdirMode) (names []string, dirents []DirEntry, fi []FileInfo, err error) {
	{
		ff, ok := f.f.(interface {
			ReadDir(n int) ([]fs.DirEntry, error)
		})
		if !ok {
			err = syscall.ENOTSUP
			goto error
		}
		dirents, err := ff.ReadDir(n)
		if err != nil {
			goto error
		}
		switch mode {
		case readdirName:
			names = make([]string, len(dirents))
			for i, de := range dirents {
				names[i] = de.Name()
			}
			return names, nil, nil, err
		case readdirDirEntry:
			return nil, dirents, nil, err
		default:
			fi = make([]FileInfo, len(dirents))
			for i, de := range dirents {
				fi[i], err = de.Info()
				if err != nil {
					goto error
				}
			}
			return nil, nil, fi, err
		}
	}
error:
	return nil, nil, nil, f.wrapErr("readdir", err)
}

func (f *File) checkValid(op string) error {
	if f == nil {
		return ErrInvalid
	}
	return nil
}

func (f *File) read(p []byte) (n int, err error) {
	return f.f.Read(p)
}

func (f *File) pread(b []byte, off int64) (n int, err error) {
	if ff, ok := f.f.(interface {
		ReadAt(b []byte, off int64) (n int, err error)
	}); ok {
		return ff.ReadAt(b, off)
	}
	return 0, syscall.ENOTSUP
}

func (f *File) write(p []byte) (n int, err error) {
	if ff, ok := f.f.(io.Writer); ok {
		return ff.Write(p)
	}
	return 0, syscall.ENOTSUP
}

func (f *File) pwrite(b []byte, off int64) (n int, err error) {
	if ff, ok := f.f.(interface {
		WriteAt(b []byte, off int64) (n int, err error)
	}); ok {
		return ff.WriteAt(b, off)
	}
	return 0, syscall.ENOTSUP
}

func (f *File) seek(offset int64, whence int) (ret int64, err error) {
	if ff, ok := f.f.(interface {
		Seek(offset int64, whence int) (ret int64, err error)
	}); ok {
		return ff.Seek(offset, whence)
	}
	return 0, syscall.ENOTSUP
}

// See docs in file.go:(*File).Chmod.
func (f *File) chmod(mode FileMode) (err error) {
	err = syscall.ENOTSUP
	if ff, ok := f.f.(interface {
		Chmod(mode fs.FileMode) error
	}); ok {
		err = ff.Chmod(mode)
		if err == nil {
			return nil
		}
	}
	return f.wrapErr("chmod", err)
}

func (f *File) setDeadline(t time.Time) (err error) {
	err = syscall.ENOTSUP
	if ff, ok := f.f.(interface {
		SetDeadline(t time.Time) error
	}); ok {
		err = ff.SetDeadline(t)
		if err == nil {
			return nil
		}
	}
	return f.wrapErr("setDeadline", err)
}

func (f *File) setReadDeadline(t time.Time) (err error) {
	err = syscall.ENOTSUP
	if ff, ok := f.f.(interface {
		SetReadDeadline(t time.Time) error
	}); ok {
		err = ff.SetReadDeadline(t)
		if err == nil {
			return nil
		}
	}
	return f.wrapErr("setReadDeadline", err)
}

func (f *File) setWriteDeadline(t time.Time) (err error) {
	err = syscall.ENOTSUP
	if ff, ok := f.f.(interface {
		SetWriteDeadline(t time.Time) error
	}); ok {
		err = ff.SetWriteDeadline(t)
		if err == nil {
			return nil
		}
	}
	return f.wrapErr("setWriteDeadline", err)
}

// Close closes the File, rendering it unusable for I/O.
// On files that support SetDeadline, any pending I/O operations will
// be canceled and return immediately with an error.
// Close will return an error if it has already been called.
func (f *File) Close() error {
	if f == nil {
		return ErrInvalid
	}
	if err := f.f.Close(); err != nil {
		return f.wrapErr("close", err)
	}
	return nil
}

// Stat returns the FileInfo structure describing file.
// If there is an error, it will be of type *PathError.
func (f *File) Stat() (fi FileInfo, err error) {
	err = syscall.ENOTSUP
	if ff, ok := f.f.(interface {
		Stat() (FileInfo, error)
	}); ok {
		fi, err = ff.Stat()
		if err == nil {
			return fi, nil
		}
	}
	return nil, f.wrapErr("stat", err)
}

func epipecheck(file *File, e error) {}

// fixLongPath is a noop on non-Windows platforms.
func fixLongPath(path string) string {
	return path
}

func syscallMode(i FileMode) uint32 {
	return uint32(i.Perm())
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

func ignoringEINTR(fn func() error) error {
	return fn()
}

// provided by package embedded/rtos

func openFile(name string, flag int, perm fs.FileMode) (f fs.File, err error)
func chmod(name string, mode FileMode) error
func rename(oldname, newname string) error

// Remove removes the named file or (empty) directory.
// If there is an error, it will be of type *PathError.
func Remove(name string) error
