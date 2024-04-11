// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

import "unsafe"

// A P64 represents 64-bit memory mapped register that stores a pointer.
//BUG: go:notinheap broken in go 1.18
type P64 struct {
	r uint64
}

// Addr returns the address of r as uintptr.
func (r *P64) Addr() uintptr {
	return uintptr(unsafe.Pointer(r))
}

// Load returns the value of r.
func (r *P64) Load() uintptr {
	return uintptr(load64(&r.r))
}

// Store stores v in r.
func (r *P64) Store(p unsafe.Pointer) {
	store64(&r.r, uint64(uintptr(p)))
}
