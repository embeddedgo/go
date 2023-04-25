// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noos

package rtos

func setPrivLevel(newlevel int) (oldlevel int, err error) {
	return -1, ErrNotSupported
}
