// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import _ "unsafe"

func irqEnable(irq IRQ, prio, ctxid int) error {
	if uint(prio+1) > intPrioHighest+1 {
		return ErrBadIntPrio
	}
	_, _, errno := runtime_irqctl(int(irq), prio, ctxid)
	return errnoError(errno)
}

func irqDisable(irq IRQ, ctxid int) error {
	_, _, errno := runtime_irqctl(int(irq), -2, ctxid)
	return errnoError(errno)
}

func irqStatus(irq IRQ, ctxid int) (enabled bool, prio int, err error) {
	en, prio, errno := runtime_irqctl(int(irq), -3, ctxid)
	return en != 0, prio, errnoError(errno)
}

//go:linkname runtime_irqctl runtime.irqctl
func runtime_irqctl(irq, ctl, ctxid int) (enabled, prio, errno int)
