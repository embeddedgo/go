// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

const iscgo = false

var (
	_cgo_thread_start             unsafe.Pointer
	_cgo_notify_runtime_init_done unsafe.Pointer
	_cgo_set_context_function     unsafe.Pointer
	_cgo_yield                    unsafe.Pointer
)

var (
	ncgocall       uint64
	cgo_yield      = &_cgo_yield
	cgoAlwaysFalse bool
	cgoHasExtraM   bool
)

type cgoCallers [1]uintptr

func cgoCheckSliceCopy(typ *_type, dst, src unsafe.Pointer, n int) {}
func cgoCheckWriteBarrier(dst *uintptr, src uintptr)               {}
func cgocall(fn unsafe.Pointer, arg unsafe.Pointer) int32          { return 0 }

func cgoCheckMemmove(typ *_type, dst unsafe.Pointer, src unsafe.Pointer, off uintptr, size uintptr) {}

