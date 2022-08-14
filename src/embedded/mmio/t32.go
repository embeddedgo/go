// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || arm || thumb

package mmio

type T8 interface {
	~int8 | ~uint8
}

type T16 interface {
	~int16 | ~uint16
}

type T32 interface {
	~int32 | ~uint32 | ~int | ~uint | ~uintptr
}

type T64 interface {
	~int64 | ~uint64
}
