// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

// Package rtos defines a seven well known interrupt priority levels ordered by
// increasing urgency as follows: IntPrioLowest, IntPrioLow, IntPrioSysCall,
// IntPrioMid, IntPrioSysTimer, IntPrioHigh, IntPrioHighest. There can be
// additional priority levels between the defined ones and some can have the
// same level. You can use ordinary equality operators to compare priority
// levels.
//
// The architecture specific code should provide (if possible) at least four
// different real priority levels so the defined levels can satisfy the
// following inequality:
//
// Lowest <= Low < SysCall < IntPrioMid <= SysTimer < High <= Highest
//
// Architecture that supports interrupt nesting must ensure that the incoming
// interrupt request can not preempt the interrupt handler that runs with the
// same or higher priority.
const (
	IntPrioHighest  int = intPrioHighest
	IntPrioHigh     int = intPrioHigh
	IntPrioSysTimer int = intPrioSysTimer
	IntPrioMid      int = intPrioMid
	IntPrioSysCall  int = intPrioSysCall
	IntPrioLow      int = intPrioLow
	IntPrioLowest   int = intPrioLowest

	IntPrioCurrent int = intPrioCurrent
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
