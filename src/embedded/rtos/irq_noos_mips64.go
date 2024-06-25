package rtos

const (
	intPrioHighest = 1
	intPrioHigh    = 1
	intPrioMid     = 0
	intPrioLow     = 0
	intPrioLowest  = 0

	intPrioSysTimer = IntPrioHighest + 1
	intPrioSysCall  = IntPrioHighest + 1

	intPrioCurrent = -1
)
