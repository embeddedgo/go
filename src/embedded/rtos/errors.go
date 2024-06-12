// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

type Error struct{ s string }

func (e *Error) Error() string { return e.s }

var (
	ErrUknown         = &Error{"rtos: unknown error"}
	ErrNotSupported   = &Error{"rtos: operation not supported"}
	ErrBadPrivLevel   = &Error{"rtos: bad privilege level"}
	ErrInsufPrivLevel = &Error{"rtos: insufficient privilege level"}
	ErrBadIntNumber   = &Error{"rtos: bad interrupt number"}
	ErrBadIntPrio     = &Error{"rtos: bad interrupt priority"}
	ErrBadIntCtx      = &Error{"rtos: bad interrupt context"}
)
