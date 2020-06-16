// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

// SetPrivLevel sets privilege level for current thread to newlevel. Level 0 is
// the most privileged and allows access to all system resources. Any resource
// available to the level n is also available to the level 0..n. If n < 0 the
// privilege level is not changed. SetPrivLevel returns previous level number
// and error. SetPrivLevel should be usually preceded by runtime.LockOSThread()
// to ensure effect for the current goroutine.
func SetPrivLevel(newlevel int) (oldlevel int, err error) {
	return setPrivLevel(newlevel)
}
