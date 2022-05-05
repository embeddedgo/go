// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

var scavenge struct {
	sysmonWake uint32
}

func bgscavenge(c chan int)                                    {}
func gcPaceScavenger(heapGoal, lastHeapGoal uint64)            {}
func wakeScavenger()                                           {}
func readyForScavenger()                                       {}
func heapRetained() uint64                                     { return 0 }
func printScavTrace(gen uint32, released uintptr, forced bool) {}

func (p *pageAlloc) scavengeStartGen()               {}
func (p *pageAlloc) scavenge(nbytes uintptr) uintptr { return 0 }
