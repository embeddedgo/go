// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

const (
	blockprofilerate = 0
	mutexprofilerate = 0
)

var MemProfileRate int = 0

//go:notinheap
type bucket struct{}

func blockevent(cycles int64, skip int)                     {}
func tracealloc(p unsafe.Pointer, size uintptr, typ *_type) {}
func mProf_Malloc(p unsafe.Pointer, size uintptr)           {}
func mProf_PostSweep()                                      {}
func mProf_NextCycle()                                      {}
func mProf_Flush()                                          {}
func tracegc()                                              {}
func tracefree(p unsafe.Pointer, size uintptr)              {}
func mProf_Free(b *bucket, size uintptr)                    {}
func mutexevent(cycles int64, skip int)                     {}
