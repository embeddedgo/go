// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import _ "unsafe"

type intCtx int

func irqEnable(irq IRQ, prio int, ctx IntCtx) error {
	if uint(prio+1) > intPrioHighest+1 {
		return ErrBadIntPrio
	}
	_, _, errno := runtime_irqctl(int(irq), prio, int(ctx))
	return errnoError(errno)
}

func irqDisable(irq IRQ, ctx IntCtx) error {
	_, _, errno := runtime_irqctl(int(irq), -2, int(ctx))
	return errnoError(errno)
}

func irqStatus(irq IRQ, ctx IntCtx) (enabled bool, prio int, err error) {
	en, prio, errno := runtime_irqctl(int(irq), -3, int(ctx))
	return en != 0, prio, errnoError(errno)
}

//go:linkname runtime_irqctl runtime.irqctl
func runtime_irqctl(irq, ctl, ctxid int) (enabled, prio, errno int)

//go:linkname handlerMode runtime.isr
func handlerMode() bool
