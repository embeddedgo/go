// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

// For the linkers. Must match Go definitions.

const (
	STACKSYSTEM = 0
	StackBig    = 4096
	StackSmall  = 128
)

const (
	StackPreempt = -1314 // 0xfff...fade
)

var StackGuard, StackLimit, StackSystem int

// Initialize StackGuard and StackLimit according to target system.
func init() {
	if GOOS == "noos" && GOARCH == "thumb" {
		StackSystem = 27 * 4
		StackGuard = 440 + StackSystem
	} else {
		StackSystem = STACKSYSTEM
		StackGuard = 880*stackGuardMultiplier() + StackSystem
	}
	StackLimit = StackGuard - StackSystem - StackSmall
}

// stackGuardMultiplier returns a multiplier to apply to the default
// stack guard size. Larger multipliers are used for non-optimized
// builds that have larger stack frames or for specific targets.
func stackGuardMultiplier() int {
	// On AIX, a larger stack is needed for syscalls.
	if GOOS == "aix" {
		return 2
	}
	return stackGuardMultiplierDefault
}
