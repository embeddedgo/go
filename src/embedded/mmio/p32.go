// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

import "unsafe"

// A P32 represents 32-bit memory mapped register that stores a pointer.
//BUG: go:notinheap broken in go 1.18
type P32 struct {
	r uint32
}

// Addr returns the address of r as uintptr.
func (r *P32) Addr() uintptr {
	return uintptr(unsafe.Pointer(r))
}

// Load returns the value of r.
func (r *P32) Load() uintptr {
	return uintptr(load32(&r.r))
}

// Store stores p in r.
func (r *P32) Store(p unsafe.Pointer) {
	store32(&r.r, uint32(uintptr(p)))
}
