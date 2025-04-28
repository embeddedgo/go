// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "internal/runtime/atomic"

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
	v := atomic.Load(&l.v)
	for v == 0 {
		breakpoint() // catch the case of unlocking a not locked mutex
	}
	atomic.Store(&l.v, 0)
}
