// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import _ "unsafe"

//go:linkname notewakeup runtime.notelwakeup
func notewakeup(n *Note)
