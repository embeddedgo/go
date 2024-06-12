// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"context"
	"syscall"
)

func (c *IPConn) readFrom(b []byte) (int, *IPAddr, error) {
	return 0, nil, syscall.ENOTSUP
}

func (c *IPConn) readMsg(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error) {
	return 0, 0, 0, nil, syscall.ENOTSUP
}

func (c *IPConn) writeTo(b []byte, addr *IPAddr) (int, error) {
	return 0, syscall.ENOTSUP
}

func (c *IPConn) writeMsg(b, oob []byte, addr *IPAddr) (n, oobn int, err error) {
	return 0, 0, syscall.ENOTSUP
}

func (sd *sysDialer) dialIP(ctx context.Context, laddr, raddr *IPAddr) (*IPConn, error) {
	return nil, syscall.ENOTSUP
}

func (sl *sysListener) listenIP(ctx context.Context, laddr *IPAddr) (*IPConn, error) {
	return nil, syscall.ENOTSUP
}
