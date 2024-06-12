// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import "syscall"

type dirInfo struct{}

func (f *File) readdirnames(n int) (names []string, err error) {
	return nil, syscall.ENOTSUP
}
