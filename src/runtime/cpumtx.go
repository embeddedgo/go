// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !noos !thumb

package runtime

import "runtime/internal/atomic"

// simple spinlock
type cpumtx struct {
	v uint32
}

//go:nosplit
func (l *cpumtx) lock() {
	for {
		if atomic.Cas(&l.v, 0, 1) {
			return
		}
	}
}

//go:nosplit
func (l *cpumtx) unlock() {
	atomic.Store(&l.v, 0)
}
