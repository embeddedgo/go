// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import (
	"sync"
	_ "unsafe"
)

var setsystimmx sync.Mutex

func setSystemTimer(nanotime func() int64, setalarm func(ns int64)) error {
	setsystimmx.Lock()
	runtime_setsystim(nanotime, setalarm)
	setsystimmx.Unlock()
	return nil
}

//go:linkname runtime_setsystim runtime.setsystim
func runtime_setsystim(nanotime func() int64, setalarm func(ns int64))

