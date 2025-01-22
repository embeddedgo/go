// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import _ "unsafe"

// Positive numbers are PLIC priority levels. It seems the lowest number of the
// priority levels supported by real hardware is 7. CLIC based designs were not
// considered.
const (
	intPrioHighest  = 7
	intPrioHigh     = 6
	intPrioMid      = 4
	intPrioLow      = 2
	intPrioLowest   = 1
	intPrioCurrent  = 0
	intPrioSysTimer = intPrioHighest + 1
	intPrioSysCall  = intPrioHighest + 1
)

//go:linkname irqctl runtime.irqctl
func irqctl(irq, ctl, ctxid int) (enabled, prio, errno int)
