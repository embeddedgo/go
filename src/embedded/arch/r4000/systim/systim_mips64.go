package systim

import (
	"embedded/rtos"
	"internal/cpu/r4000/creg"
	_ "unsafe" // for linkname
)

var hz2ns int64

// Setup setups COUNT register to work as system timer.
func Setup(clkhz int64) {
	hz2ns = 2 * 1e9 / clkhz
	creg.COMPARE.Store(0xffffffff)

	creg.STATUS.SetBits(creg.IP_TIMER)
	creg.CAUSE.ClearBits(creg.IP_TIMER)
	rtos.SetSystemTimer(nanotime, setAlarm)
}

//go:nosplit
func nanotime() int64 {
	return int64(creg.COUNT.Load()) * hz2ns
}

//go:nosplit
func setAlarm(ns int64) {
	var compare uint32
	if ns >= 0 {
		compare = uint32(ns / hz2ns)
	}
	creg.COMPARE.Store(compare)
}
