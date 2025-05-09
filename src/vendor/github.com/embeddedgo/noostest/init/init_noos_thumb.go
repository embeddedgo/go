// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package init

import (
	"embedded/arch/cortexm/systim"
	"embedded/rtos"
	"runtime"
)

func init() {
	setupSemihosting()

	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)
	systim.Setup(10e6, 1e6, true)
	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()

	rtos.SetSystemTimer(systim.Nanotime, nil)
}
