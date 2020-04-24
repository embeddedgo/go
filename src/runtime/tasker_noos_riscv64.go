// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

type mOS [1]uint64

func curcpuSleep()          { breakpoint() }
func curcpuSavectxSched()   { breakpoint() }
func curcpuWakeup()         { breakpoint() }
func archnewm(m *m)         { breakpoint() }
func curcpuSchedule()       { breakpoint() }
func curcpuSavectxCall()    { breakpoint() }
func (cpu *cpuctx) wakeup() { breakpoint() }
