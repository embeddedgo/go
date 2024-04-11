// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"context"
	"net/netip"
	"syscall"
)

func (sd *sysDialer) dialUDP(ctx context.Context, laddr, raddr *UDPAddr) (*UDPConn, error) {
	return nil, syscall.ENOTSUP
}

func (sl *sysListener) listenUDP(ctx context.Context, laddr *UDPAddr) (*UDPConn, error) {
	return nil, syscall.ENOTSUP
}

func (c *UDPConn) readFrom(b []byte, addr *UDPAddr) (int, *UDPAddr, error) {
	return 0, nil, syscall.ENOTSUP
}

func (c *UDPConn) readFromAddrPort(b []byte) (n int, addr netip.AddrPort, err error) {
	err = syscall.ENOTSUP
	return
}

func (c *UDPConn) readMsg(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error) {
	err = syscall.ENOTSUP
	return
}

func (c *UDPConn) writeTo(b []byte, addr *UDPAddr) (int, error) {
	return 0, syscall.ENOTSUP
}

func (c *UDPConn) writeToAddrPort(b []byte, addr netip.AddrPort) (int, error) {
	return 0, syscall.ENOTSUP
}

func (c *UDPConn) writeMsg(b, oob []byte, addr *UDPAddr) (n, oobn int, err error) {
	err = syscall.ENOTSUP
	return
}

func (c *UDPConn) writeMsgAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error) {
	err = syscall.ENOTSUP
	return
}

func (sl *sysListener) listenMulticastUDP(ctx context.Context, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error) {
	return nil, syscall.ENOTSUP
}
