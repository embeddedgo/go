// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

import "unsafe"

type T8 interface{ ~int8 | ~uint8 }

// An R8 represents 8-bit memory mapped register of type T.
// BUG: go:notinheap broken in go 1.18
type R8[T T8] struct {
	r uint8
}

// Addr returns the address of r as uintptr.
func (r *R8[_]) Addr() uintptr {
	return uintptr(unsafe.Pointer(r))
}

// LoadBits returns the value od r logicaly anded with mask. It is a convenient
// replacement for r.Load()&mask.
func (r *R8[T]) LoadBits(mask T) T {
	return T(load8(&r.r)) & mask
}

// StoreBits stores bits in r selected by mask. It is convenient replacement for
// r.Store(r.Load()&^mask | bits&mask). This is not an atomic operation.
func (r *R8[T]) StoreBits(mask, bits T) {
	store8(&r.r, load8(&r.r)&^uint8(mask)|uint8(bits&mask))
}

// SetBits sets bits in r selected by mask. This is not an atomic operation.
func (r *R8[T]) SetBits(mask T) {
	store8(&r.r, load8(&r.r)|uint8(mask))
}

// ClearBits clears bits in r selected by mask. This is not an atomic operation.
func (r *R8[T]) ClearBits(mask T) {
	store8(&r.r, load8(&r.r)&^uint8(mask))
}

// Load returns the value of r.
func (r *R8[T]) Load() T {
	return T(load8(&r.r))
}

// Store stores v in r.
func (r *R8[T]) Store(v T) {
	store8(&r.r, uint8(v))
}

// An RM8 represents a set of bits in R selected by Mask.
type RM8[T T8] struct {
	R    *R8[T]
	Mask T
}

// Set sets all bits in rm. This is not an atomic operation.
func (rm RM8[T]) Set() { rm.R.SetBits(rm.Mask) }

// Clear clears all bits in b. This is not an atomic operation.
func (rm RM8[T]) Clear() { rm.R.ClearBits(rm.Mask) }

// Load returns the value of b.
func (rm RM8[T]) Load() T { return rm.R.LoadBits(rm.Mask) }

// Store stores bits in b. This is not an atomic operation.
func (rm RM8[T]) Store(bits T) { rm.R.StoreBits(rm.Mask, bits) }

// U8 is an alias for the R8[uint8] type kept for backward compatibility.
//type U8 = R8[uint8]

// UM8 is an alias for the RM8[uint8] type kept for backward compatibility.
//type UM8 = RM8[uint8]
