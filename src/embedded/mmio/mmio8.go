// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

import "unsafe"

//go:noescape
func load8(addr *uint8) uint8

//go:noescape
func store8(addr *uint8, v uint8)

// An U8 represents 8-bit memory mapped register.
//BUG: go:notinheap broken in go 1.18
type U8 struct {
	r uint8
}

// Addr returns the address of r as uintptr.
func (r *U8) Addr() uintptr {
	return uintptr(unsafe.Pointer(r))
}

// SetBit sets n-th bit in r. This is not an atomic operation.
func (r *U8) SetBit(n int) {
	store8(&r.r, load8(&r.r)|uint8(1)<<uint(n))
}

// ClearBit clears n-th bit in r. This is not an atomic operation.
func (r *U8) ClearBit(n int) {
	store8(&r.r, load8(&r.r)&^uint8(1)<<uint(n))
}

// Bit returns the value of n-th bit in r (0 or 1).
func (r *U8) LoadBit(n int) int {
	return int(load8(&r.r)>>uint(n)) & 1
}

// StoreBit sets the value of n-th bit in r to least significant bit of v. This
// is not an atomic operation.
func (r *U8) StoreBit(n, v int) {
	mask := uint8(1) << uint(n)
	store8(&r.r, load8(&r.r)&^mask|uint8(v<<uint(n))&mask)
}

// Bits returns the value od r logicaly anded with mask. It is a convenient
// replacement for r.Load()&mask.
func (r *U8) LoadBits(mask uint8) uint8 {
	return load8(&r.r) & mask
}

// StoreBits stores bits in r selected by mask. It is convenient replacement for
// r.Store(r.Load()&^mask | bits&mask). This is not an atomic operation.
func (r *U8) StoreBits(mask, bits uint8) {
	store8(&r.r, load8(&r.r)&^mask|bits&mask)
}

// SetBits sets bits in r selected by mask. This is not an atomic operation.
func (r *U8) SetBits(mask uint8) {
	store8(&r.r, load8(&r.r)|mask)
}

// ClearBits clears bits in r selected by mask. This is not an atomic operation.
func (r *U8) ClearBits(mask uint8) {
	store8(&r.r, load8(&r.r)&^mask)
}

// Load returns the value of r.
func (r *U8) Load() uint8 {
	return load8(&r.r)
}

// Store stores v in r.
func (r *U8) Store(v uint8) {
	store8(&r.r, v)
}

//func (r *U8) Field(mask uint8) int {
//	return bits.Field8(r.r, mask)
//}
//func (r *U8) SetField(mask uint8, v int) {
//	r.StoreBits(mask, bits.MakeField8(v, mask))
//}

// An UM8 represents a set of bits in R selected by Mask.
type UM8 struct {
	R    *U8
	Mask uint8
}

// Set sets all bits in b. This is not an atomic operation.
func (b UM8) Set() { b.R.SetBits(b.Mask) }

// Clear clears all bits in b. This is not an atomic operation.
func (b UM8) Clear() { b.R.ClearBits(b.Mask) }

// Load returns the value of b.
func (b UM8) Load() uint8 { return b.R.LoadBits(b.Mask) }

// Store stores bits in b. This is not an atomic operation.
func (b UM8) Store(bits uint8) { b.R.StoreBits(b.Mask, bits) }

//func (b UM8) LoadVal() int   { return b.R.Field(uint8(b.Mask)) }
//func (b UM8) StoreVal(v int) { b.R.SetField(b.Mask, v) }
