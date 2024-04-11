// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package plic provides access to the registers of the Platform-Level Interrupt Controller.
package plic

import (
	"embedded/mmio"
	"unsafe"
)

type Periph struct {
	PRIO [1024]mmio.U32 // interrupt source priorities
	PEND [32]mmio.U32   // interrupt pending bits
	_    [992]uint32
	EN   [15872][32]mmio.U32 // interrupt enable bits
	_    [14336]uint32
	TC   [15872]struct {
		THR mmio.U32 // priority threshold
		CC  mmio.U32 // claim/complete
		_   [1022]uint32
	}
}

func PLIC() *Periph {
	// BUG: The 0x0C000000 address is valid for Qemu Virt, Kendryte K210, some
	// SiFive cores but generally it's implementation specific.
	return (*Periph)(unsafe.Pointer(uintptr(0x0C000000)))
}
