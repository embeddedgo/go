// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

const (
	intPrioLow      = intPrioHighest / 4
	intPrioSysTimer = intPrioHighest / 2
	intPrioHigh     = intPrioHighest * 3 / 4
	intPrioHighest  = 255
)
