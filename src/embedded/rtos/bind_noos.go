// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import _ "unsafe"

func bind(ctx ExeCtx) (oldctx ExeCtx, err error) {
	old, errno := runtime_bind(int(ctx))
	return ExeCtx(old), errnoError(errno)
}

//go:linkname runtime_bind runtime.bind
func runtime_bind(cpuid int) (oldcpuid, errno int)
