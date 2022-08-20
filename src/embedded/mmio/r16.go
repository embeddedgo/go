// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

import "unsafe"

type T16 interface{ ~int16 | ~uint16 }

// An R16 represents 16-bit memory mapped register of type T.
// BUG: go:notinheap broken in go 1.18
type R16[T T16] struct {
	r uint16
}

// Addr returns the address of r as uintptr.
func (r *R16[_]) Addr() uintptr {
	return uintptr(unsafe.Pointer(r))
}

// LoadBits returns the value od r logicaly anded with mask. It is a convenient
// replacement for r.Load()&mask.
func (r *R16[T]) LoadBits(mask T) T {
	return T(load16(&r.r)) & mask
}

// StoreBits stores bits in r selected by mask. It is convenient replacement for
// r.Store(r.Load()&^mask | bits&mask). This is not an atomic operation.
func (r *R16[T]) StoreBits(mask, bits T) {
	store16(&r.r, load16(&r.r)&^uint16(mask)|uint16(bits&mask))
}

// SetBits sets bits in r selected by mask. This is not an atomic operation.
func (r *R16[T]) SetBits(mask T) {
	store16(&r.r, load16(&r.r)|uint16(mask))
}

// ClearBits clears bits in r selected by mask. This is not an atomic operation.
func (r *R16[T]) ClearBits(mask T) {
	store16(&r.r, load16(&r.r)&^uint16(mask))
}

// Load returns the value of r.
func (r *R16[T]) Load() T {
	return T(load16(&r.r))
}

// Store stores v in r.
func (r *R16[T]) Store(v T) {
	store16(&r.r, uint16(v))
}

// An RM16 represents a set of bits in R selected by Mask.
type RM16[T T16] struct {
	R    *R16[T]
	Mask T
}

// Set sets all bits in rm. This is not an atomic operation.
func (rm RM16[T]) Set() { rm.R.SetBits(rm.Mask) }

// Clear clears all bits in b. This is not an atomic operation.
func (rm RM16[T]) Clear() { rm.R.ClearBits(rm.Mask) }

// Load returns the value of b.
func (rm RM16[T]) Load() T { return rm.R.LoadBits(rm.Mask) }

// Store stores bits in b. This is not an atomic operation.
func (rm RM16[T]) Store(bits T) { rm.R.StoreBits(rm.Mask, bits) }

// U16 is an alias for the R16[uint16] type kept for backward compatibility.
//type U16 = R16[uint16]

// UM16 is an alias for the RM16[uint16] type kept for backward compatibility.
//type UM16 = RM16[uint16]
