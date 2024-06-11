// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

import "unsafe"

type T64 interface{ ~int64 | ~uint64 }

// An R64 represents 64-bit memory mapped register of type T.
// BUG: go:notinheap broken in go 1.18
type R64[T T64] struct {
	r uint64
}

// Addr returns the address of r as uintptr.
//
//go:nosplit
func (r *R64[_]) Addr() uintptr {
	return uintptr(unsafe.Pointer(r))
}

// LoadBits returns the value od r logicaly anded with mask. It is a convenient
// replacement for r.Load()&mask.
//
//go:nosplit
func (r *R64[T]) LoadBits(mask T) T {
	return T(load64(&r.r)) & mask
}

// StoreBits stores bits in r selected by mask. It is convenient replacement for
// r.Store(r.Load()&^mask | bits&mask). This is not an atomic operation.
//
//go:nosplit
func (r *R64[T]) StoreBits(mask, bits T) {
	store64(&r.r, load64(&r.r)&^uint64(mask)|uint64(bits&mask))
}

// SetBits sets bits in r selected by mask. This is not an atomic operation.
//
//go:nosplit
func (r *R64[T]) SetBits(mask T) {
	store64(&r.r, load64(&r.r)|uint64(mask))
}

// ClearBits clears bits in r selected by mask. This is not an atomic operation.
//
//go:nosplit
func (r *R64[T]) ClearBits(mask T) {
	store64(&r.r, load64(&r.r)&^uint64(mask))
}

// Load returns the value of r.
//
//go:nosplit
func (r *R64[T]) Load() T {
	return T(load64(&r.r))
}

// Store stores v in r.
//
//go:nosplit
func (r *R64[T]) Store(v T) {
	store64(&r.r, uint64(v))
}

// An RM64 represents a set of bits in R selected by Mask.
type RM64[T T64] struct {
	R    *R64[T]
	Mask T
}

// Set sets all bits in rm. This is not an atomic operation.
//
//go:nosplit
func (rm RM64[T]) Set() { rm.R.SetBits(rm.Mask) }

// Clear clears all bits in b. This is not an atomic operation.
//
//go:nosplit
func (rm RM64[T]) Clear() { rm.R.ClearBits(rm.Mask) }

// Load returns the value of b.
//
//go:nosplit
func (rm RM64[T]) Load() T { return rm.R.LoadBits(rm.Mask) }

// Store stores bits in b. This is not an atomic operation.
//
//go:nosplit
func (rm RM64[T]) Store(bits T) { rm.R.StoreBits(rm.Mask, bits) }

// U64 is an alias for the R64[uint64] type kept for backward compatibility.
//type U64 = R64[uint64]

// UM64 is an alias for the RM64[uint64] type kept for backward compatibility.
//type UM64 = RM64[uint64]
