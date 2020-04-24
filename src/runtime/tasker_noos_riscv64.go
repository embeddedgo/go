// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"embedded/mmio"
	"internal/cpu/cortexm"
	"internal/cpu/cortexm/debug/itm"
	"internal/cpu/cortexm/mpu"
	"internal/cpu/cortexm/nvic"
	"internal/cpu/cortexm/scb"
	"internal/cpu/cortexm/scid"
	"unsafe"
)


type mOS [1]uint64
