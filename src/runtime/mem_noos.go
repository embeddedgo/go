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
// sysReserve allocates down from noosMem.*.end.
//
// noosPersistentAlloc provides fast and memory efficient implementation of
// the persistentalloc1 function.

var noosMem struct {
	free, arena, nodma    pamem
	arenaStart, arenaSize uintptr
	size                  uintptr // initial sum of free, nodma and arena bytes
	mx                    mutex
}

//go:nosplit
func meminit(freeStart, freeEnd, nodmaStart, nodmaEnd, stackTop uintptr) (nodmaStack bool) {
	if nodmaStart < stackTop && stackTop < nodmaEnd {
		nodmaStart = stackTop // ISR stack(s) in the NoDMA memory
		nodmaStack = true
	}
	freeSize := freeEnd - freeStart
	nodmaSize := nodmaEnd - nodmaStart
	size := freeSize + nodmaSize

	arenaStart := alignUp(freeStart, heapArenaBytes)
	arenaSize := freeEnd - arenaStart

	noosMem.free = pamem{freeStart, arenaStart} // free memory before the arena
	noosMem.arena = pamem{arenaStart, freeEnd}  // memory also reserved by the arena
	noosMem.nodma = pamem{nodmaStart, nodmaEnd} // extra nadma memory
	noosMem.arenaStart = arenaStart
	noosMem.arenaSize = arenaSize
	noosMem.size = size

	physPageSize = _PageSize
	return
}

type pamem struct {
	start, end uintptr
}

//go:nosplit
func (m *pamem) alloc(size, align uintptr) unsafe.Pointer {
	p := alignDown(m.end-size, align)
	if p < m.start || p > m.end {
		return nil
	}
	m.end = p
	return unsafe.Pointer(p)
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
	return noosRawAlloc(size, 8)

}

//go:nosplit
func sysAllocOS(size uintptr) unsafe.Pointer {
	p := noosRawAlloc(size, 8)
	if p == nil {
		throw("runtime: cannot allocate memory")
	}
	return p
}

func sysUsedOS(v unsafe.Pointer, n uintptr) {
	lock(&noosMem.mx)
	noosMem.arena.start = max(noosMem.arena.start, uintptr(v)+n)
	if noosMem.arena.start > noosMem.arena.end {
		throw("runtime: cannot allocate memory")
	}
	unlock(&noosMem.mx)
}

func sysFreeOS(v unsafe.Pointer, n uintptr)             {}
func sysMapOS(v unsafe.Pointer, n uintptr)              {}
func sysUnusedOS(v unsafe.Pointer, n uintptr)           {}
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

func noosRawAlloc(size, align uintptr) unsafe.Pointer {
	lock(&noosMem.mx)
	p := noosMem.free.alloc(size, align)
	if p == nil {
		p = noosMem.nodma.alloc(size, align)
	}
	if p == nil {
		// If both free and nodma don't have enough space left, start
		// allocating in the arena. While the arena allocates from
		// bottom, pamem allocates from top.
		p = noosMem.arena.alloc(size, align)
	}
	unlock(&noosMem.mx)
	return p
}

// align must be power of two
//
//go:nosplit
func noosPersistentAlloc(size, align uintptr, sysStat *sysMemStat) (p *notInHeap) {
	align = max(align, 8)
	p = (*notInHeap)(noosRawAlloc(size, align))
	if p == nil {
		throw("runtime: cannot allocate memory")
	}
	sysStat.add(int64(size))
	gcController.mappedReady.Add(int64(size))
	return
}
