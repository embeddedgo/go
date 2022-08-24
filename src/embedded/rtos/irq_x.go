// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !noos

package rtos

import _ "unsafe"

const (
	intPrioHighest  = 0
	intPrioHigh     = 0
	intPrioSysTimer = 0
	intPrioMid      = 0
	intPrioSysCall  = 0
	intPrioLow      = 0
	intPrioLowest   = 0
	intPrioCurrent  = -1
)

type intCtx int

func irqEnable(irq IRQ, prio int, ctx IntCtx) error {
	return ErrNotSupported
}

func irqDisable(irq IRQ, ctx IntCtx) error {
	return ErrNotSupported
}

func irqStatus(irq IRQ, ctx IntCtx) (enabled bool, prio int, err error) {
	return false, 0, ErrNotSupported
}

func handlerMode() bool {
	return false
}
