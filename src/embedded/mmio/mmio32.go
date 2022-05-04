// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

import "unsafe"

//go:noescape
func load32(addr *uint32) uint32

//go:noescape
func store32(addr *uint32, v uint32)

// An U32 represents 32-bit memory mapped register.
//BUG: go:notinheap broken in go 1.18
type U32 struct {
	r uint32
}

// Addr returns the address of r as uintptr.
func (r *U32) Addr() uintptr {
	return uintptr(unsafe.Pointer(r))
}

// SetBit sets n-th bit in r. This is not an atomic operation.
func (r *U32) SetBit(n int) {
	store32(&r.r, load32(&r.r)|uint32(1)<<uint(n))
}

// ClearBit clears n-th bit in r. This is not an atomic operation.
func (r *U32) ClearBit(n int) {
	store32(&r.r, load32(&r.r)&^uint32(1)<<uint(n))
}

// Bit returns the value of n-th bit in r (0 or 1).
func (r *U32) LoadBit(n int) int {
	return int(load32(&r.r)>>uint(n)) & 1
}

// StoreBit sets the value of n-th bit in r to least significant bit of v. This
// is not an atomic operation.
func (r *U32) StoreBit(n, v int) {
	mask := uint32(1) << uint(n)
	store32(&r.r, load32(&r.r)&^mask|uint32(v<<uint(n))&mask)
}

// Bits returns the value od r logicaly anded with mask. It is a convenient
// replacement for r.Load()&mask.
func (r *U32) LoadBits(mask uint32) uint32 {
	return load32(&r.r) & mask
}

// StoreBits stores bits in r selected by mask. It is convenient replacement for
// r.Store(r.Load()&^mask | bits&mask). This is not an atomic operation.
func (r *U32) StoreBits(mask, bits uint32) {
	store32(&r.r, load32(&r.r)&^mask|bits&mask)
}

// SetBits sets bits in r selected by mask. This is not an atomic operation.
func (r *U32) SetBits(mask uint32) {
	store32(&r.r, load32(&r.r)|mask)
}

// ClearBits clears bits in r selected by mask. This is not an atomic operation.
func (r *U32) ClearBits(mask uint32) {
	store32(&r.r, load32(&r.r)&^mask)
}

// Load returns the value of r.
func (r *U32) Load() uint32 {
	return load32(&r.r)
}

// Store stores v in r.
func (r *U32) Store(v uint32) {
	store32(&r.r, v)
}

//func (r *U32) Field(mask uint32) int {
//	return bits.Field32(r.r, mask)
//}
//func (r *U32) SetField(mask uint32, v int) {
//	r.StoreBits(mask, bits.MakeField32(v, mask))
//}

// An UM32 represents a set of bits in R selected by Mask.
type UM32 struct {
	R    *U32
	Mask uint32
}

// Set sets all bits in b. This is not an atomic operation.
func (b UM32) Set() { b.R.SetBits(b.Mask) }

// Clear clears all bits in b. This is not an atomic operation.
func (b UM32) Clear() { b.R.ClearBits(b.Mask) }

// Load returns the value of b.
func (b UM32) Load() uint32 { return b.R.LoadBits(b.Mask) }

// Store stores bits in b. This is not an atomic operation.
func (b UM32) Store(bits uint32) { b.R.StoreBits(b.Mask, bits) }

//func (b UM32) LoadVal() int   { return b.R.Field(uint32(b.Mask)) }
//func (b UM32) StoreVal(v int) { b.R.SetField(b.Mask, v) }
