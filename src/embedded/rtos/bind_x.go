
// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noos

package rtos

const notBound ExeCtx = -1

type exeCtx int

func bind(ctx ExeCtx) (oldctx ExeCtx, err error) {
	return -1, ErrNotSupported
}
