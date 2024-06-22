package rtos

import "runtime"

const (
	intPrioHighest = runtime.IntPrioHigh
	intPrioHigh    = runtime.IntPrioHigh
	intPrioMid     = runtime.IntPrioNormal
	intPrioLow     = runtime.IntPrioNormal
	intPrioLowest  = runtime.IntPrioNormal

	intPrioSysTimer = runtime.IntPrioHigh + 1
	intPrioSysCall  = runtime.IntPrioHigh + 1

	intPrioCurrent = runtime.IntPrioCurrent
)
