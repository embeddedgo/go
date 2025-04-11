// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import (
	"time"
	_ "unsafe"
)

// Cond allows to communicate the occurrence of an event.
//
// Exactly one goroutine can call Wait but it is allowed for multiple goroutines
// or interrupt handlers to call Signal.
//
// A Cond must always be declared as a global variable and must not be copied
// after first use.
type Cond struct {
	// must be in sync with runtime.pollDesc
	key  uintptr
	seq  uintptr
	lock uintptr // incompatible with goexperiment.staticlockranking
	link uintptr
	self *Cond
}

// Wait waits on the Cond to become true and consumes it by setting it back to
// false. If there previously was a call to Signal Wait returns immediately.
// Sleeps indefinitely for a negative timeout. Returns whether the Cond became
// true during the call.
func (n *Cond) Wait(timeout time.Duration) bool { return condwait(n, int64(timeout)) }

// Signal sets the Cond to true.
func (n *Cond) Signal() { condsignal(n) }

//go:linkname condsignal runtime.rtos_condsignal
func condsignal(n *Cond)

//go:linkname condwait runtime.rtos_condwait
func condwait(n *Cond, timeout int64) bool
