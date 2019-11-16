// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"runtime/internal/atomic"
	"unsafe"
)

var _cgo_setenv unsafe.Pointer   // pointer to C function
var _cgo_unsetenv unsafe.Pointer // pointer to C function

//go:nosplit
func gogetenv(key string) string { return "" }

// The following functions should be defined in the runtime/debug package but
// for now they are here because the runtime/debug imports time package, not
// ported yet.

// SetGCTrace is the alternate way to set GODEBUG=gctrace=n at the runtime.
func SetGCTrace(in int32) {
	atomic.Store((*uint32)(unsafe.Pointer(&debug.gctrace)), uint32(in))
}

// SetGCPercent is the alternate way to set GOGC=percent at the runtime.
func SetGCPercent(percent int) int {
	return int(setGCPercent(int32(percent)))
}
