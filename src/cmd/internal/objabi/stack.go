// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

import "internal/buildcfg"

// For the linkers. Must match Go definitions.

const (
	STACKSYSTEM = 0
	StackBig    = 4096
	StackSmall  = 128
)

var StackGuard, StackLimit, StackSystem int

// Initialize StackGuard and StackLimit according to target system.
func init() {
	if buildcfg.GOOS == "noos" && buildcfg.GOARCH == "thumb" {
		StackSystem = 27 * 4
		StackGuard = 464 + StackSystem
	} else {
		StackSystem = STACKSYSTEM
		StackGuard = 928*stackGuardMultiplier() + StackSystem
	}
	StackLimit = StackGuard - StackSystem - StackSmall
}

// stackGuardMultiplier returns a multiplier to apply to the default
// stack guard size. Larger multipliers are used for non-optimized
// builds that have larger stack frames or for specific targets.
func stackGuardMultiplier() int {
	// On AIX, a larger stack is needed for syscalls.
	if buildcfg.GOOS == "aix" {
		return 2
	}
	return stackGuardMultiplierDefault
}
