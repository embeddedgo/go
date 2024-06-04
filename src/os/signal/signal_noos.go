// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package signal

import (
	"os"
)

//var sigtab = make(map[os.Signal]int)

// Defined by the runtime package.
func signal_disable(uint32)
func signal_enable(uint32)
func signal_ignore(uint32)
func signal_ignored(uint32) bool
func signal_recv() string

const numSig = 0

func signum(sig os.Signal) int   { return -1 }
func enableSignal(sig int)       {}
func disableSignal(sig int)      {}
func ignoreSignal(sig int)       {}
func signalIgnored(sig int) bool { return true }
