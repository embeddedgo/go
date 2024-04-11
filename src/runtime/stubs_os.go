// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noos

package runtime

func isr() bool {
	return false
}

func noosMemory() (heapBase, heapSize, limit uintptr) {
	return 0, 0, 0
}

func noosPersistentAlloc(size, align uintptr, sysStat *sysMemStat) *notInHeap {
	return nil
}
