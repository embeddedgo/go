// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rtos defines seven well known interrupt priority levels among which
// five are ordered by increasing urgency as follows: IntPrioLowest, IntPrioLow,
// IntPrioMid, IntPrioHigh, IntPrioHighest. There can be additional priority
// levels between the defined ones and some or all of them can have the same
// effecive level.
//
// You can use ordinary equality operators to compare interrupt priority levels.
// The defined priority levels satisfy the following inequality:
//
// Lowest <= Low <= IntPrioMid <= High <= Highest
//
// Architecture that supports interrupt nesting must ensure the incoming
// interrupt request can not preempt the interrupt handler that runs with the
// same or higher priority.
//
// The IntPrioSysCall and IntPrioSysTimer are special priority levels. Do not
// use them to set or calclulate an interrupt priority if their values are
// outside of [IntPrioLowest, IntPrioHighest] range.
package rtos

const (
	IntPrioHighest = intPrioHighest
	IntPrioHigh    = intPrioHigh
	IntPrioMid     = intPrioMid
	IntPrioLow     = intPrioLow
	IntPrioLowest  = intPrioLowest

	IntPrioSysCall  = intPrioSysCall
	IntPrioSysTimer = intPrioSysTimer

	IntPrioCurrent = intPrioCurrent
)

// IRQ represents a source of interrupts. It provides interface to basic
// operations such as enabling/disabling handling interrupts from this source
// and setting its priority. There can be other (system) interrupt/exception
// sources not exposed by this interface.
type IRQ int

// IntCtx represents an interrupt context. The context has implementation
// specific meaning, can correspond to CPU, core, hardware thread, privilege
// level or any combination thereof.
type IntCtx intCtx

// Enable sets the priority of the interrupt and enables interrupt requests in
// the context specified by ctxid. The context has implementation specific
// meaning, can correspond to CPU, core, hardware thread, privilege level or any
// combination thereof.
func (irq IRQ) Enable(prio int, ctx IntCtx) error {
	return irqEnable(irq, prio, ctx)
}

// Disable disables interrupt requests int the context specified by ctxid.
func (irq IRQ) Disable(ctx IntCtx) error {
	return irqDisable(irq, ctx)
}

// Status reports whether the irq is enabled in context ctxid and returns its
// priority.
func (irq IRQ) Status(ctx IntCtx) (enabled bool, prio int, err error) {
	return irqStatus(irq, ctx)
}

// HandlerMode reports whether the function is called in interupt handler mode.
func HandlerMode() bool {
	return handlerMode()
}
