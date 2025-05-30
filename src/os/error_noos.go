// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"syscall"
	_ "unsafe"
)

type syscallErrorType = *syscall.Error

//go:linkname errENOSYS syscall.ENOTSUP
//go:linkname errERANGE syscall.EINVAL
//go:linkname errENOMEM syscall.ENOMEM
var (
	errENOSYS *syscall.Error
	errERANGE *syscall.Error
	errENOMEM *syscall.Error
)
