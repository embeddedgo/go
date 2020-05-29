// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

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
// use them to set or calclulate an interrupt priority if they are outside the
// [IntPrioLowest, IntPrioHighest] range.
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

// IRQ represents an user accessible interrupt. It provides interface to basic
// operations such as enabling/disabling handling of the interrupt requests and
// setting interrupt priorities. There can be other (system) interrupts not
// exposed by this interface.
type IRQ int

// Enable sets the priority of the interrupt and enables handling interrupt
// requests on the CPU (core/hart) specified by cpuid. The special cpuid -1 lets
// the Enable to select one or more CPUs automaticaly.
func (irq IRQ) Enable(prio, cpuid int) error {
	return irqEnable(irq, prio)
}

// Disable disables handling of interrupt requests by the CPU (core/hart)
// specified by cpuid. The special cpuid -1 disables irq on all CPUs.
func (irq IRQ) Disable(cpuid int) error {
	return irqDisable(irq)
}

// Status reports whether the irq is enabled on the CPU (core/hart) specifie by
// cpuid and returns its priority.
func (irq IRQ) Status(cpuid int) (enabled bool, prio int, err error) {
	return irqStatus(irq)
}
