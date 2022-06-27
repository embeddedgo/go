// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

var _cgo_setenv unsafe.Pointer   // pointer to C function
var _cgo_unsetenv unsafe.Pointer // pointer to C function

//go:nosplit
func gogetenv(key string) string { return "" }
