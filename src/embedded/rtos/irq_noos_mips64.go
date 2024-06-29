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
