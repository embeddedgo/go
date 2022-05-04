// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

const (
	// The number of levels in the radix tree.
	summaryLevels = 2

	// Constants for testing.
	pageAlloc32Bit = 0
	pageAlloc64Bit = 0

	// Number of bits needed to represent all indices into the L1 of the
	// chunks map.
	//
	// See (*pageAlloc).chunks for more details. Update the documentation
	// there should this number change.
	pallocChunksL1Bits = 0
)

// See comment in mpagealloc_64bit.go.
var levelBits = [summaryLevels]uint{
	summaryL0Bits,
	summaryLevelBits,
}

// See comment in mpagealloc_64bit.go.
var levelShift = [summaryLevels]uint{
	heapAddrBits - summaryL0Bits,
	heapAddrBits - summaryL0Bits - 1*summaryLevelBits,
}

// See comment in mpagealloc_64bit.go.
var levelLogPages = [summaryLevels]uint{
	logPallocChunkPages + 1*summaryLevelBits,
	logPallocChunkPages,
}

// See mpagealloc_64bit.go for details.
func (p *pageAlloc) sysInit() {
	// Calculate how much memory all our entries will take up.
	//
	// This should be around 12 KiB or less.
	totalSize := uintptr(0)
	for l := 0; l < summaryLevels; l++ {
		totalSize += (uintptr(1) << (heapAddrBits - levelShift[l])) * pallocSumBytes
	}
	totalSize = alignUp(totalSize, physPageSize)

	// Reserve memory for all levels in one go.
	reservation := sysAlloc(totalSize, p.sysStat)
	if reservation == nil {
		throw("failed to reserve page summary memory")
	}

	// Iterate over the reservation and cut it up into slices.
	//
	// Maintain i as the byte offset from reservation where
	// the new slice should start.
	for l, shift := range levelShift {
		entries := 1 << (heapAddrBits - shift)

		// Put this reservation into a slice.
		sl := notInHeapSlice{(*notInHeap)(reservation), 0, entries}
		p.summary[l] = *(*[]pallocSum)(unsafe.Pointer(&sl))

		reservation = add(reservation, uintptr(entries)*pallocSumBytes)
	}
}

// See mpagealloc_64bit.go for details.
func (p *pageAlloc) sysGrow(base, limit uintptr) {
	if base%pallocChunkBytes != 0 || limit%pallocChunkBytes != 0 {
		print("runtime: base = ", hex(base), ", limit = ", hex(limit), "\n")
		throw("sysGrow bounds not aligned to pallocChunkBytes")
	}

	// Walk up the tree and update the summary slices.
	for l := len(p.summary) - 1; l >= 0; l-- {
		// Figure out what part of the summary array this new address space needs.
		// Note that we need to align the ranges to the block width (1<<levelBits[l])
		// at this level because the full block is needed to compute the summary for
		// the next level.
		lo, hi := addrsToSummaryRange(l, base, limit)
		_, hi = blockAlignSummaryRange(l, lo, hi)
		if hi > len(p.summary[l]) {
			p.summary[l] = p.summary[l][:hi]
		}
	}
}
