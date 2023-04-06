// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

import (
	"syscall"
	"time"
)

type FD struct {
	// Lock sysfd and serialize access to Read and Write methods.
	fdmu fdMutex

	// Whether this is a file rather than a network socket.
	isFile bool
}

func (fd *FD) destroy() error { return nil }

func (fd *FD) RawControl(f func(uintptr)) error {
	return syscall.ENOTSUP
}

func (fd *FD) RawRead(f func(uintptr) bool) error {
	return syscall.ENOTSUP
}

func (fd *FD) RawWrite(f func(uintptr) bool) error {
	return syscall.ENOTSUP
}

func (fd *FD) SetDeadline(t time.Time) error {
	return syscall.ENOTSUP
}
