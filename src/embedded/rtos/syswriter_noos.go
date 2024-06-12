// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import (
	"sync"
	_ "unsafe"
)

var setsyswritermx sync.Mutex

func setSysWriter(w func(fd int, p []byte) int) error {
	setsyswritermx.Lock()
	runtime_setsyswriter(w)
	setsyswritermx.Unlock()
	return nil
}

//go:linkname runtime_setsyswriter runtime.setsyswriter
func runtime_setsyswriter(w func(fd int, p []byte) int)
