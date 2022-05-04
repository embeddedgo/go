// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

import "unsafe"

//go:noescape
func load16(addr *uint16) uint16

//go:noescape
func store16(addr *uint16, v uint16)

// An U16 represents 16-bit memory mapped register.
//BUG: go:notinheap broken in go 1.18
type U16 struct {
	r uint16
}

// Addr returns the address of r as uintptr.
func (r *U16) Addr() uintptr {
	return uintptr(unsafe.Pointer(r))
}

// SetBit sets n-th bit in r. This is not an atomic operation.
func (r *U16) SetBit(n int) {
	store16(&r.r, load16(&r.r)|uint16(1)<<uint(n))
}

// ClearBit clears n-th bit in r. This is not an atomic operation.
func (r *U16) ClearBit(n int) {
	store16(&r.r, load16(&r.r)&^uint16(1)<<uint(n))
}

// Bit returns the value of n-th bit in r (0 or 1).
func (r *U16) LoadBit(n int) int {
	return int(load16(&r.r)>>uint(n)) & 1
}

// StoreBit sets the value of n-th bit in r to least significant bit of v. This
// is not an atomic operation.
func (r *U16) StoreBit(n, v int) {
	mask := uint16(1) << uint(n)
	store16(&r.r, load16(&r.r)&^mask|uint16(v<<uint(n))&mask)
}

// Bits returns the value od r logicaly anded with mask. It is a convenient
// replacement for r.Load()&mask.
func (r *U16) LoadBits(mask uint16) uint16 {
	return load16(&r.r) & mask
}

// StoreBits stores bits in r selected by mask. It is convenient replacement for
// r.Store(r.Load()&^mask | bits&mask). This is not an atomic operation.
func (r *U16) StoreBits(mask, bits uint16) {
	store16(&r.r, load16(&r.r)&^mask|bits&mask)
}

// SetBits sets bits in r selected by mask. This is not an atomic operation.
func (r *U16) SetBits(mask uint16) {
	store16(&r.r, load16(&r.r)|mask)
}

// ClearBits clears bits in r selected by mask. This is not an atomic operation.
func (r *U16) ClearBits(mask uint16) {
	store16(&r.r, load16(&r.r)&^mask)
}

// Load returns the value of r.
func (r *U16) Load() uint16 {
	return load16(&r.r)
}

// Store stores v in r.
func (r *U16) Store(v uint16) {
	store16(&r.r, v)
}

//func (r *U16) Field(mask uint16) int {
//	return bits.Field16(r.r, mask)
//}
//func (r *U16) SetField(mask uint16, v int) {
//	r.StoreBits(mask, bits.MakeField16(v, mask))
//}

// An UM16 represents a set of bits in R selected by Mask.
type UM16 struct {
	R    *U16
	Mask uint16
}

// Set sets all bits in b. This is not an atomic operation.
func (b UM16) Set() { b.R.SetBits(b.Mask) }

// Clear clears all bits in b. This is not an atomic operation.
func (b UM16) Clear() { b.R.ClearBits(b.Mask) }

// Load returns the value of b.
func (b UM16) Load() uint16 { return b.R.LoadBits(b.Mask) }

// Store stores bits in b. This is not an atomic operation.
func (b UM16) Store(bits uint16) { b.R.StoreBits(b.Mask, bits) }

//func (b UM16) LoadVal() int   { return b.R.Field(uint16(b.Mask)) }
//func (b UM16) StoreVal(v int) { b.R.SetField(b.Mask, v) }
