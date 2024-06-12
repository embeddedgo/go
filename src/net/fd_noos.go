// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"internal/poll"
	"os"
	"syscall"
	"time"
)

type netFD struct {
	pfd poll.FD

	// immutable until Close
	net   string
	laddr Addr
	raddr Addr
}

func (fd *netFD) setAddr(laddr, raddr Addr) {
}

func (fd *netFD) Close() error {
	return syscall.ENOTSUP
}

func (fd *netFD) shutdown(how int) error {
	return syscall.ENOTSUP
}

func (fd *netFD) closeRead() error {
	return syscall.ENOTSUP
}

func (fd *netFD) closeWrite() error {
	return syscall.ENOTSUP
}

func (fd *netFD) Read(p []byte) (n int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) readFrom(p []byte) (n int, sa syscall.Sockaddr, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) readFromInet4(p []byte, from *syscall.SockaddrInet4) (n int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) readFromInet6(p []byte, from *syscall.SockaddrInet6) (n int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) readMsg(p []byte, oob []byte, flags int) (n, oobn, retflags int, sa syscall.Sockaddr, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) readMsgInet4(p []byte, oob []byte, flags int, sa *syscall.SockaddrInet4) (n, oobn, retflags int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) readMsgInet6(p []byte, oob []byte, flags int, sa *syscall.SockaddrInet6) (n, oobn, retflags int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) Write(p []byte) (nn int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) writeTo(p []byte, sa syscall.Sockaddr) (n int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) writeToInet4(p []byte, sa *syscall.SockaddrInet4) (n int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) writeToInet6(p []byte, sa *syscall.SockaddrInet6) (n int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) writeMsg(p []byte, oob []byte, sa syscall.Sockaddr) (n int, oobn int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) writeMsgInet4(p []byte, oob []byte, sa *syscall.SockaddrInet4) (n int, oobn int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) writeMsgInet6(p []byte, oob []byte, sa *syscall.SockaddrInet6) (n int, oobn int, err error) {
	err = syscall.ENOTSUP
	return
}

func (fd *netFD) SetDeadline(t time.Time) error {
	return syscall.ENOTSUP
}

func (fd *netFD) SetReadDeadline(t time.Time) error {
	return syscall.ENOTSUP
}

func (fd *netFD) SetWriteDeadline(t time.Time) error {
	return syscall.ENOTSUP
}

func (fd *netFD) dup() (f *os.File, err error) {
	err = syscall.ENOTSUP
	return
}

func setNoDelay(fd *netFD, noDelay bool) error {
	return syscall.ENOTSUP
}
