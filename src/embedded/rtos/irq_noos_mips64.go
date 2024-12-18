// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

const (
	intPrioHighest = 1
	intPrioHigh    = 1
	intPrioMid     = 0
	intPrioLow     = 0
	intPrioLowest  = 0

	intPrioSysTimer = intPrioHighest + 1
	intPrioSysCall  = intPrioHighest + 1

	intPrioCurrent = -1
)
