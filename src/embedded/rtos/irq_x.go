// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !noos

package rtos

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

func irqEnable(irq IRQ, prio int) error {
	return ErrNotSuppoted
}

func irqDisable(irq IRQ) error {
	return ErrNotSuppoted
}

func irqStatus(irq IRQ) (enabled bool, prio int, err error) {
	return false, 0, ErrNotSuppoted
}
