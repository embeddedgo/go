// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

// SetSysTimer registers two functions that the thread scheduler uses to
// communicate with the system timer.
//
// Nanotime should return the monotonic time in nanoseconds.
//
// Setalarm is called by the scheduler with the ns >= 0 to ask the system timer
// to wake the current CPU at the specified time. Sheduler can also pass ns < 0
// if it do not want to be woken up by the system timer. A spurious or
// inaccurate wakeups are acceptable.
//
// There are two main types of system timer implementations.
//
// The ticking implementation can simply wake up the CPU with a constant period.
// In this case setalarm can be nil.
//
// The tickless implementation schould try to wake up the CPU only once, at ns
// or just after ns.
//
// The System Timer interrupt handler, if used, must run with IntPrioSysTimer
// priority.
//
// SetSystemTimer is intended to be called only once, at the very beginning of
// the system initialization. If you want to change the system timer
// implementation at a later stage (which is discouraged) you must ensure a
// continuity of time and that all gorutines that sleeeps using old setalarm
// function will be wakeup.
func SetSystemTimer(nanotime func() int64, setalarm func(ns int64)) error {
	return setSystemTimer(nanotime, setalarm)
}
