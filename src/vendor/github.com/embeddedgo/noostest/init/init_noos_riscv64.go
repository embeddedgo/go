// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package init

import (
	"embedded/arch/riscv/systim"
	"runtime"
)

func init() {
	setupSemihosting()
	systim.Setup(10e6)
	runtime.GOMAXPROCS(2)
}
