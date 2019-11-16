// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import _ "unsafe"

const (
	intPrioCurrent = -1
	intPrioLowest  = 0
)

func irqEnable(irq IRQ, prio int) error {
	if uint(prio+1) > intPrioHighest+1 {
		return ErrBadIntPrio
	}
	_, _, errno := runtime_irqctl(int(irq), prio)
	return errnoError(errno)
}

func irqDisable(irq IRQ) error {
	_, _, errno := runtime_irqctl(int(irq), -2)
	return errnoError(errno)
}

func irqStatus(irq IRQ) (enabled bool, prio int, err error) {
	en, prio, errno := runtime_irqctl(int(irq), -3)
	return en != 0, prio, errnoError(errno)
}

//go:linkname runtime_irqctl runtime.irqctl
func runtime_irqctl(irq, ctl int) (enabled, prio, errno int)
