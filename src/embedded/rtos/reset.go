// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import "unsafe"

/// Reset levels
const (
	// SystemReset resets all major system components. If possible it uses the
	// default reset mechanism provided by the architecture.
	SystemReset = 0

	// UnsafeReboot disables all interrupts and restarts the whole software
	// without resetting the hardware. Use with caution because other masters
	// in the system like DMA controllers may keep going, used peripherals may
	// be active and IO lines may be configured as outputs so they can affect
	// the external world in an unexpected way.
	UnsafeReboot = 15
)

// Reset performs a reset operation on a given level. Generally it does not
// return in case of successful reset. It can return true only if the
// architecture provides a such exotic reset level that allows the software keep
// going. The addr parameter is architecture specific and generally may allow to
// boot an another image (set to nil to boot the current image).
func Reset(level int, addr unsafe.Pointer) (success bool) {
	return reset(level, addr)
}
