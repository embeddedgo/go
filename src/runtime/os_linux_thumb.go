// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/cpu"
	"unsafe"
)

const (
	_HWCAP_VFPv3 = 1 << 13 // introduced in 2.6.30
	_HWCAP_IDIVT = 1 << 18
)

func vdsoCall()

func checkgoarm() {
	if cpu.HWCap&_HWCAP_IDIVT == 0 {
		print("runtime: hardware division not supported, cannot run GOARCH=thumb binary")
		exit(1)
	}
	if cpu.HWCap&_HWCAP_VFPv3 == 0 && goarmsoftfp == 0 {
		print("runtime: this CPU has no VFPv3 floating point hardware, so it cannot run\n")
		print("a binary compiled for VFPv3 hard floating point. Recompile adding ,softfloat\n"to GOARM\n")
		exit(1)
	}
}

func archauxv(tag, val uintptr) {
	switch tag {
	case _AT_HWCAP:
		cpu.HWCap = uint(val)
	case _AT_HWCAP2:
		cpu.HWCap2 = uint(val)
	case _AT_PLATFORM:
		cpu.Platform = gostringnocopy((*byte)(unsafe.Pointer(val)))
	}
}

func osArchInit() {}

//go:nosplit
func cputicks() int64 {
	// nanotime() is a poor approximation of CPU ticks that is enough for the profiler.
	return nanotime()
}
