// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

const (
	intPrioHighest  = 255 - 0<<5 // do not use with nRF52 SoftDevice
	intPrioHigh     = 255 - 1<<5
	intPrioSysTimer = 255 - 2<<5
	intPrioMid      = 255 - 3<<5
	intPrioSysCall  = 255 - 4<<5 // compatible with nRF52 SoftDevice
	intPrioLow      = 255 - 5<<5
	intPrioLowest   = 255 - 6<<5
	intPrioPendSV   = 255 - 255 // unusable because can not wake the scheduler
)
