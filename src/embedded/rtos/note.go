// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import _ "unsafe"

// Note allows sleep and wakeup on one-time events.
//
// Before any calls to Sleep or Wakeup, must call Clear to initialize the Note.
//
// Exactly one gorutine can call Sleep and exactly one gorutine or interrupt
// handler can call Wakeup (once). Future Sleep will return immediately.
//
// Subsequent Clear must be called only after previous Sleep has returned, e.g.
// it is disallowed to call Clear straight after Wakeup.
type Note struct {
	// must be in sync with runtime.notel
	key  uintptr
	link uintptr
}

// Sleep sleeps on cleared note until other goroutine or interrupt handler
// call Wakeup or until the timeout.
func (n *Note) Sleep(ns int64) bool { return runtime_notetsleepg(n, ns) }

// Wakeup wakeups the goroutine that sleeps or will try to sleep on the note.
func (n *Note) Wakeup() { notewakeup(n) }

// Clear clears the note. It also behaves as full memory barrier in a way that
// the Clear itself and any memory access preceding it in the program order
// happens before any memory access that follows it.
func (n *Note) Clear() { atomic_Storeuintptr(&n.key, 0) }

//go:linkname runtime_notetsleepg runtime.notetsleepg
func runtime_notetsleepg(n *Note, ns int64) bool

//go:linkname atomic_Storeuintptr runtime/internal/atomic.Storeuintptr
func atomic_Storeuintptr(p *uintptr, v uintptr)
