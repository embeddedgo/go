// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

// Simple memory allocator that emulates OS allocator
//
// There are noosMemory and noosPersistentAlloc functions specific to
// noos target.
//
// noosMemory returns the address and size of the memory allocated to the arena
// (heapArenaBytes aligned) and the memory limit for GC.
//
// sysReserve allocates down from noosMem.end always returning _PageSize aligned
// memory.
//
// noosPersistentAlloc provides fast and memory efficient implementation of
// the persistentalloc1 function.

var noosMem struct {
	free, nodma           pamem
	arenaStart, arenaSize uintptr
	size                  uintptr // initial sum of free, nodma and arena bytes
	mx                    mutex
}

//go:nosplit
func meminit(freeStart, freeEnd, nodmaStart, nodmaEnd uintptr) {
	freeSize := freeEnd - freeStart
	nodmaSize := nodmaEnd - nodmaStart
	size := freeSize + nodmaSize

	// Estimate the space needed for non-heap allocations
	const pallocMin = 2 * noosDefaultHeapMinimum
	palloc := int((freeSize>>(pageShift))*unsafe.Sizeof(emptymspan) + pallocMin)

	// We prefer the non-DMA memory for non-heap objects to preserve as much
	// as possible of the DMA capable memory for heap allocations
	pallocInFree := max(palloc-int(nodmaSize), 0)

	// Reduce the arena by the remainder of the non-heap space that did not fit in
	// the non-DMA memory, properly align the arena
	arenaSize := freeSize - uintptr(pallocInFree)
	arenaSize &= ^(uintptr(heapArenaBytes) - 1)
	arenaStart := freeEnd - arenaSize

	noosMem.free.start = freeStart
	noosMem.free.end = freeEnd
	noosMem.nodma.start = nodmaStart
	noosMem.nodma.end = nodmaEnd
	noosMem.arenaStart = arenaStart
	noosMem.arenaSize = arenaSize
	noosMem.size = size
}

type pamem struct {
	start, end uintptr
}

//go:nosplit
func (m *pamem) allocPages(size uintptr) unsafe.Pointer {
	var p uintptr
	if m.end-m.start >= size {
		p = m.end - size
		m.end = p
	}
	return unsafe.Pointer(p)
}

//go:nosplit
func (m *pamem) alloc(size, align uintptr) unsafe.Pointer {
	var p uintptr
	align-- // align must be power of two
	astart := (m.start + align) &^ align
	if m.end-astart >= size {
		p = astart
		m.start = astart + size
	}
	return unsafe.Pointer(p)
}

//go:nosplit
func sysReserve1(size uintptr) unsafe.Pointer {
	lock(&noosMem.mx)
	p := noosMem.free.allocPages(size)
	if p == nil {
		p = noosMem.nodma.allocPages(size)
	}
	unlock(&noosMem.mx)
	return p
}

//go:nosplit
func sysReserveOS(v unsafe.Pointer, size uintptr) unsafe.Pointer {
	if v != nil {
		// The address space of NOOS memory is contiguous,
		// so requesting specific addresses is not supported. We could use
		// a different address, but then mheap.sysAlloc discards the result
		// right away and we don't reuse chunks passed to sysFree.
		return nil
	}
	size += (_PageSize - 1)
	size &^= (_PageSize - 1)
	return sysReserve1(size)
}

//go:nosplit
func sysAllocOS(size uintptr) unsafe.Pointer {
	size += (_PageSize - 1)
	size &^= (_PageSize - 1)
	p := sysReserve1(size)
	if p == nil {
		throw("runtime: cannot allocate memory")
	}
	return p
}

func sysFreeOS(v unsafe.Pointer, n uintptr)             {}
func sysMapOS(v unsafe.Pointer, n uintptr)              {}
func sysUnusedOS(v unsafe.Pointer, n uintptr)           {}
func sysUsedOS(v unsafe.Pointer, n uintptr)             {}
func sysFaultOS(v unsafe.Pointer, n uintptr)            {}
func sysHugePageOS(v unsafe.Pointer, n uintptr)         {}
func sysNoHugePageOS(v unsafe.Pointer, n uintptr)       {}
func sysHugePageCollapseOS(v unsafe.Pointer, n uintptr) {}

// TODO: consider replace this function by direct initialization of mheap_.arena.
//
//go:nosplit
func noosMemory() (heapBase, heapSize, limit uintptr) {
	return noosMem.arenaStart, noosMem.arenaSize, noosMem.size
}

// align must be power of two
//
//go:nosplit
func noosPersistentAlloc(size, align uintptr, sysStat *sysMemStat) (p *notInHeap) {
	if size&(_PageSize-1) == 0 {
		p = (*notInHeap)(sysReserve1(size))
	} else {
		if align == 0 {
			align = 8
		}
		lock(&noosMem.mx)
		p = (*notInHeap)(noosMem.free.alloc(size, align))
		if p == nil {
			p = (*notInHeap)(noosMem.nodma.alloc(size, align))
		}
		unlock(&noosMem.mx)
	}
	if p == nil {
		throw("runtime: cannot allocate memory")
	}
	sysStat.add(int64(size))
	gcController.mappedReady.Add(int64(size))
	return
}
