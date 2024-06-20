package runtime

import "unsafe"

const pallocMin = 20 * 1024
const signalStackSize = 4096

//go:nosplit
func meminit(freeStart, freeEnd, nodmaStart, nodmaEnd uintptr) {
	freeSize := freeEnd - freeStart
	nodmaSize := nodmaEnd - nodmaStart
	size := freeSize + nodmaSize

	// Estimate the space needed for non-heap allocations
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
