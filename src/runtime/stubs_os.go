// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noos

package runtime

func isr() bool {
	return false
}

func sysReserveMaxArena() (addr, size uintptr, mapMemory bool) {
	return 0, 0, false
}

func sysPersistentAlloc(size, align uintptr, sysStat *sysMemStat) *notInHeap {
	return nil
}
