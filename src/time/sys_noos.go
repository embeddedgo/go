// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

import "syscall"

func open(name string) (uintptr, error) {
	return 0, syscall.ENOENT
}

func closefd(fd uintptr) {
}

func read(fd uintptr, buf []byte) (int, error) {
	return 0, syscall.EINVAL
}

func preadn(fd uintptr, buf []byte, off int) error {
	return syscall.EINVAL
}

func move(sec int64, nsec int32) (sec0 int64, nsec0 int32)

// Set sets the wall clock in a way that the old time will become the new time.
// In fact it moves the wall clock by new - old amount of time. The monotonic
// clock is unaffected. Set returns the new wall time that corresponds to the
// the zero monotonic time.
func Set(old, new Time) Time {
	sec := new.sec() - old.sec()
	nsec := new.nsec() - old.nsec()
	if nsec < 0 {
		nsec += 1e9
		sec--
	} else if nsec >= 1e9 {
		nsec -= 1e9
		sec++
	}
	sec0, nsec0 := move(sec, nsec)
	return Unix(sec0, int64(nsec0))
}
