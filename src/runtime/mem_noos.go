// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

// Simple memory allocator that emulates OS allocator
//
// There are sysReserveMaxArena and sysPersistentAlloc functions specific to
// noos target.
//
// sysReserveMaxArena allocates maximum arena space heapArenaBytes aligned for
// mheap. It is intended to be run only once by mallocinit.
//
// sysReserve allocates down from sysMem.end always returning _PageSize aligned
// memory.
//
// sysPersistentAlloc is fast and memory efficient implementation of
// persistentalloc1. It uses sysReserve for _PageSize alignned allocations.
// Otherwise it allocates from sysMem.start.

var sysMem struct {
	free, nodma           pamem
	arenaStart, arenaSize uintptr
	mx                    mutex
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

// sysReserveMaxArena is for initial reservation, must not be called
// concurrently with any other reservation. TODO: Consider replace this by
// direct initialization of mheap_.arena.
//go:nosplit
func sysReserveMaxArena() (addr, size uintptr, mapMemory bool) {
	return sysMem.arenaStart, sysMem.arenaSize, false
}

//go:nosplit
func sysReserve1(size uintptr) unsafe.Pointer {
	lock(&sysMem.mx)
	p := sysMem.free.allocPages(size)
	if p == nil {
		p = sysMem.nodma.allocPages(size)
	}
	unlock(&sysMem.mx)
	return p
}

//go:nosplit
func sysReserve(v unsafe.Pointer, size uintptr) unsafe.Pointer {
	size += (_PageSize - 1)
	size &^= (_PageSize - 1)
	return sysReserve1(size)
}

//go:nosplit
func sysAlloc(size uintptr, sysStat *sysMemStat) unsafe.Pointer {
	size += (_PageSize - 1)
	size &^= (_PageSize - 1)
	p := sysReserve1(size)
	if p == nil {
		throw("runtime: cannot allocate memory")
	}
	sysStat.add(int64(size))
	return p
}

// align must be power of two
//go:nosplit
func sysPersistentAlloc(size, align uintptr, sysStat *sysMemStat) (p *notInHeap) {
	if size&(_PageSize-1) == 0 {
		p = (*notInHeap)(sysReserve1(size))
	} else {
		if align == 0 {
			align = 8
		}
		lock(&sysMem.mx)
		p = (*notInHeap)(sysMem.free.alloc(size, align))
		if p == nil {
			p = (*notInHeap)(sysMem.nodma.alloc(size, align))
		}
		unlock(&sysMem.mx)
	}
	if p == nil {
		throw("runtime: cannot allocate memory")
	}
	sysStat.add(int64(size))
	return
}

//go:nosplit
func sysMap(v unsafe.Pointer, n uintptr, sysStat *sysMemStat) {
	sysStat.add(int64(n))
}

//go:nosplit
func sysFree(v unsafe.Pointer, n uintptr, sysStat *sysMemStat) {
	sysStat.add(-int64(n))
}

func sysUnused(v unsafe.Pointer, n uintptr)   {}
func sysUsed(v unsafe.Pointer, n uintptr)     {}
func sysFault(v unsafe.Pointer, n uintptr)    {}
func sysHugePage(v unsafe.Pointer, n uintptr) {}
