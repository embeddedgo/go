// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

import "internal/buildcfg"

// For the linkers. Must match Go definitions.

const (
	StackBig   = 4096
	StackSmall = 128
)

var StackSystem, rawGuard int

func init() {
	if buildcfg.GOOS == "noos" && buildcfg.GOARCH == "thumb" {
		StackSystem = 27 * 4
		rawGuard = 464
	} else {
		rawGuard = 928
	}
}

func StackLimit(race bool) int {
	// This arithmetic must match that in runtime/stack.go:{_StackGuard,_StackLimit}.
	stackGuard := rawGuard*stackGuardMultiplier(race) + StackSystem
	stackLimit := stackGuard - StackSystem - StackSmall
	return stackLimit
}

// stackGuardMultiplier returns a multiplier to apply to the default
// stack guard size. Larger multipliers are used for non-optimized
// builds that have larger stack frames or for specific targets.
func stackGuardMultiplier(race bool) int {
	// This arithmetic must match that in runtime/internal/sys/consts.go:StackGuardMultiplier.
	n := 1
	// On AIX, a larger stack is needed for syscalls.
	if buildcfg.GOOS == "aix" {
		n += 1
	}
	// The race build also needs more stack.
	if race {
		n += 1
	}
	return n
}
