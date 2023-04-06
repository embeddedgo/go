// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "syscall"

func setDefaultSockopts(s, family, sotype int, ipv6only bool) error {
	return nil
}

func setDefaultListenerSockopts(s int) error {
	return nil
}

func setDefaultMulticastSockopts(s int) error {
	return nil
}

func setReadBuffer(fd *netFD, bytes int) error {
	return syscall.ENOTSUP
}

func setWriteBuffer(fd *netFD, bytes int) error {
	return syscall.ENOTSUP
}

func setKeepAlive(fd *netFD, keepalive bool) error {
	return syscall.ENOTSUP
}

func setLinger(fd *netFD, sec int) error {
	return syscall.ENOTSUP
}
