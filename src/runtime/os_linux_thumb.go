// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "internal/cpu"

const (
	_HWCAP_VFPv3 = 1 << 13 // introduced in 2.6.30
	_HWCAP_IDIVT = 1 << 18
)

var randomNumber uint32

func checkgoarm() {
	if cpu.HWCap&_HWCAP_IDIVT == 0 {
		print("runtime: hardware division not supported, cannot run GOARCH=thumb binary")
		exit(1)
	}
	if goarm&0xF >= 0xD && cpu.HWCap&_HWCAP_VFPv3 == 0 {
		print("runtime: VFPv3 not supported, cannot run GOARM=7F/7D binary\n")
		exit(1)
	}
}

func archauxv(tag, val uintptr) {
	switch tag {
	case _AT_RANDOM:
		// sysargs filled in startupRandomData, but that
		// pointer may not be word aligned, so we must treat
		// it as a byte array.
		randomNumber = uint32(startupRandomData[4]) | uint32(startupRandomData[5])<<8 |
			uint32(startupRandomData[6])<<16 | uint32(startupRandomData[7])<<24

	case _AT_HWCAP:
		cpu.HWCap = uint(val)
	case _AT_HWCAP2:
		cpu.HWCap2 = uint(val)
	}
}

//go:nosplit
func cputicks() int64 {
	// Currently cputicks() is used in blocking profiler and to seed fastrand().
	// nanotime() is a poor approximation of CPU ticks that is enough for the profiler.
	// randomNumber provides better seeding of fastrand.
	return nanotime() + int64(randomNumber)
}
