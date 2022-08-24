// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !noos

package rtos

func setSysWriter(w func(fd int, p []byte) int) error {
	return ErrNotSupported
}
