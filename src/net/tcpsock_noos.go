// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"context"
	"io"
	"os"
	"syscall"
)

func (c *TCPConn) readFrom(r io.Reader) (int64, error) {
	return 0, syscall.ENOTSUP
}

func (sd *sysDialer) dialTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
	return nil, syscall.ENOTSUP
}

func (ln *TCPListener) ok() bool {
	return false
}

func (ln *TCPListener) accept() (*TCPConn, error) {
	return nil, syscall.ENOTSUP
}

func (ln *TCPListener) close() error {
	return syscall.ENOTSUP
}

func (ln *TCPListener) file() (*os.File, error) {
	return nil, syscall.ENOTSUP
}

func (sl *sysListener) listenTCP(ctx context.Context, laddr *TCPAddr) (*TCPListener, error) {
	return nil, syscall.ENOTSUP
}
