// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Current implementation has limitations:
//   - Overflows at 2^56 nanoseconds (834 days) because 8 bits are used to
//     improve precision of hz2ns variable.
//   - Calls between nanotime must not be more than 2^32 nanoseconds apart,
//     otherwise overflow detection fails.  Since the GC calls nanoseconds
//     frequently, it's currently not a issue.
package systim

import (
	"embedded/rtos"
	"internal/cpu/r4000/creg"
	_ "unsafe" // for linkname
)

var hz2ns int64 // 56.8 fixed-point integer
var lastTicks int64

// Setups COUNT register to work as system timer.
func Setup(clkhz int64) {
	hz2ns = 2 * 1e9
	hz2ns = hz2ns << 8
	hz2ns /= clkhz

	setAlarm(nanotime()) // init lastTicks and alarm

	rtos.SetSystemTimer(nanotime, setAlarm)
}

//go:nosplit
func nanotime() int64 {
	// TODO handle overflow in interrupt
	ticks32 := int64(creg.COUNT.Load())
	if ticks32 < lastTicks&0xffff_ffff {
		lastTicks += (1 << 32)
	}
	lastTicks = (lastTicks & ^0xffff_ffff) | ticks32

	return (lastTicks * hz2ns) >> 8
}

//go:nosplit
func setAlarm(ns int64) {
	// Only needed for a tickless implementation. As long as curcpuSleep is
	// a noop, there is no need to implement this.
}
