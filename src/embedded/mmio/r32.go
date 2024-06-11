// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

import "unsafe"

type T32 interface{ ~int32 | ~uint32 }

// An R32 represents 32-bit memory mapped register of type T.
// BUG: go:notinheap broken in go 1.18
type R32[T T32] struct {
	r uint32
}

// Addr returns the address of r as uintptr.
//
//go:nosplit
func (r *R32[_]) Addr() uintptr {
	return uintptr(unsafe.Pointer(r))
}

// LoadBits returns the value od r logicaly anded with mask. It is a convenient
// replacement for r.Load()&mask.
//
//go:nosplit
func (r *R32[T]) LoadBits(mask T) T {
	return T(load32(&r.r)) & mask
}

// StoreBits stores bits in r selected by mask. It is convenient replacement for
// r.Store(r.Load()&^mask | bits&mask). This is not an atomic operation.
//
//go:nosplit
func (r *R32[T]) StoreBits(mask, bits T) {
	store32(&r.r, load32(&r.r)&^uint32(mask)|uint32(bits&mask))
}

// SetBits sets bits in r selected by mask. This is not an atomic operation.
//
//go:nosplit
func (r *R32[T]) SetBits(mask T) {
	store32(&r.r, load32(&r.r)|uint32(mask))
}

// ClearBits clears bits in r selected by mask. This is not an atomic operation.
//
//go:nosplit
func (r *R32[T]) ClearBits(mask T) {
	store32(&r.r, load32(&r.r)&^uint32(mask))
}

// Load returns the value of r.
//
//go:nosplit
func (r *R32[T]) Load() T {
	return T(load32(&r.r))
}

// Store stores v in r.
//
//go:nosplit
func (r *R32[T]) Store(v T) {
	store32(&r.r, uint32(v))
}

// An RM32 represents a set of bits in R selected by Mask.
type RM32[T T32] struct {
	R    *R32[T]
	Mask T
}

// Set sets all bits in rm. This is not an atomic operation.
//
//go:nosplit
func (rm RM32[T]) Set() { rm.R.SetBits(rm.Mask) }

// Clear clears all bits in b. This is not an atomic operation.
//
//go:nosplit
func (rm RM32[T]) Clear() { rm.R.ClearBits(rm.Mask) }

// Load returns the value of b.
//
//go:nosplit
func (rm RM32[T]) Load() T { return rm.R.LoadBits(rm.Mask) }

// Store stores bits in b. This is not an atomic operation.
//
//go:nosplit
func (rm RM32[T]) Store(bits T) { rm.R.StoreBits(rm.Mask, bits) }

// U32 is an alias for the R32[uint32] type kept for backward compatibility.
//type U32 = R32[uint32]

// UM32 is an alias for the RM32[uint32] type kept for backward compatibility.
//type UM32 = RM32[uint32]
