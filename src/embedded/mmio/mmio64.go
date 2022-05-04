// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

import "unsafe"

//go:noescape
func load64(addr *uint64) uint64

//go:noescape
func store64(addr *uint64, v uint64)

// An U64 represents 64-bit memory mapped register.
//BUG: go:notinheap broken in go 1.18
type U64 struct {
	r uint64
}

// Addr returns the address of r as uintptr.
func (r *U64) Addr() uintptr {
	return uintptr(unsafe.Pointer(r))
}

// SetBit sets n-th bit in r. This is not an atomic operation.
func (r *U64) SetBit(n int) {
	store64(&r.r, load64(&r.r)|uint64(1)<<uint(n))
}

// ClearBit clears n-th bit in r. This is not an atomic operation.
func (r *U64) ClearBit(n int) {
	store64(&r.r, load64(&r.r)&^uint64(1)<<uint(n))
}

// Bit returns the value of n-th bit in r (0 or 1).
func (r *U64) LoadBit(n int) int {
	return int(load64(&r.r)>>uint(n)) & 1
}

// StoreBit sets the value of n-th bit in r to least significant bit of v. This
// is not an atomic operation.
func (r *U64) StoreBit(n, v int) {
	mask := uint64(1) << uint(n)
	store64(&r.r, load64(&r.r)&^mask|uint64(v<<uint(n))&mask)
}

// Bits returns the value od r logicaly anded with mask. It is a convenient
// replacement for r.Load()&mask.
func (r *U64) LoadBits(mask uint64) uint64 {
	return load64(&r.r) & mask
}

// StoreBits stores bits in r selected by mask. It is convenient replacement for
// r.Store(r.Load()&^mask | bits&mask). This is not an atomic operation.
func (r *U64) StoreBits(mask, bits uint64) {
	store64(&r.r, load64(&r.r)&^mask|bits&mask)
}

// SetBits sets bits in r selected by mask. This is not an atomic operation.
func (r *U64) SetBits(mask uint64) {
	store64(&r.r, load64(&r.r)|mask)
}

// ClearBits clears bits in r selected by mask. This is not an atomic operation.
func (r *U64) ClearBits(mask uint64) {
	store64(&r.r, load64(&r.r)&^mask)
}

// Load returns the value of r.
func (r *U64) Load() uint64 {
	return load64(&r.r)
}

// Store stores v in r.
func (r *U64) Store(v uint64) {
	store64(&r.r, v)
}

//func (r *U64) Field(mask uint64) int {
//	return bits.Field64(r.r, mask)
//}
//func (r *U64) SetField(mask uint64, v int) {
//	r.StoreBits(mask, bits.MakeField64(v, mask))
//}

// An UM64 represents a set of bits in R selected by Mask.
type UM64 struct {
	R    *U64
	Mask uint64
}

// Set sets all bits in b. This is not an atomic operation.
func (b UM64) Set() { b.R.SetBits(b.Mask) }

// Clear clears all bits in b. This is not an atomic operation.
func (b UM64) Clear() { b.R.ClearBits(b.Mask) }

// Load returns the value of b.
func (b UM64) Load() uint64 { return b.R.LoadBits(b.Mask) }

// Store stores bits in b. This is not an atomic operation.
func (b UM64) Store(bits uint64) { b.R.StoreBits(b.Mask, bits) }

//func (b UM64) LoadVal() int   { return b.R.Field(uint64(b.Mask)) }
//func (b UM64) StoreVal(v int) { b.R.SetField(b.Mask, v) }
