// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import (
	"time"
	_ "unsafe"
)

// Note allows to communicate the occurrence of an event.
//
// Before any calls to Sleep or Wakeup, must call Clear to initialize the Note.
//
// Exactly one gorutine can call Sleep but there is allowed to multiple
// gorutines or interrupt handlers to call Wakeup. Future Sleep will return
// immediately.
//
// Subsequent Clear must be called only after previous Sleep has returned, e.g.
// it is disallowed to call Clear straight after Wakeup.
type Note struct {
	// must be in sync with runtime.notel
	key  uintptr
	link uintptr
}

// Sleep sleeps on the cleared note until other goroutine or interrupt handler
// call Wakeup or until the timeout. Reports whether it was awakened before
// timeout.
func (n *Note) Sleep(timeout time.Duration) bool {
	return runtime_notetsleepg(n, int64(timeout))
}

// Wakeup wakeups the goroutine that sleeps on the note. The Wakeup remains in
// effect until subequent Clear so future Sleep will return immediately.
func (n *Note) Wakeup() { notewakeup(n) }

// Clear clears the note. Any Wakeup that happened before Clear is forgotten.
// Clear works as a publication barrier, that is, the Clear itself and any
// memory writes preceding it in the program order happens before any memory
// writes that follows it.
func (n *Note) Clear() {
	n.key = 0
	publicationBarrier()
}

//go:linkname runtime_notetsleepg runtime.notetsleepg
func runtime_notetsleepg(n *Note, ns int64) bool
