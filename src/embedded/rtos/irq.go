// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

// Package rtos defines a five well known interrupt priority levels ordered by
// increasing urgency as follows: IntPrioLowest, IRQPrioLow, IRQPrioSysTimer,
// IRQPrioHigh, IRQPrioHighest. There can be additional priority levels between
// the defined ones.
//
// You can use ordinary equality operators to compare priority levels. The
// numeric values of the defined priority levels satisfy the following
// inequality:
//
// IntPrioLowest <= IntPrioLow < IRQPrioSysTimer < IRQPrioHigh <= IRQPrioHighest
//
// Architecture that supports interrupt nesting must ensure that the incoming
// interrupt request can not preempt the interrupt handler that runs with the
// same or higher priority.
const (
	IntPrioLowest   int = intPrioLowest
	IntPrioLow      int = intPrioLow
	IntPrioSysTimer int = intPrioSysTimer
	IntPrioHigh     int = intPrioHigh
	IntPrioHighest  int = intPrioHighest

	IntPrioCurrent int = intPrioCurrent // do not change the current priority
)

// IRQ represents an user accessible interrupt. It provides interface to basic
// operations such as enabling/disabling handling of interrupt requests and
// setting interrupt priorities. There can be other (system) interrupts not
// exposed by this interface.
type IRQ int

// Enable sets the priority of interrupt and enables handling of interrupt
// requests.
func (irq IRQ) Enable(prio int) error {
	return irqEnable(irq, prio)
}

// Disable disables handling of interrupt requests.
func (irq IRQ) Disable() error {
	return irqDisable(irq)
}

// Status reports whether the irq is enabled and returns its priority.
func (irq IRQ) Status() (enabled bool, prio int, err error) {
	return irqStatus(irq)
}
