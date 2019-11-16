// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	"runtime"
	"testing"
)

// Unity tests that can be run on the host machine (not on the target MCU)

func TestMQ(t *testing.T) {
	if s := runtime.MQTest(); s != "" {
		t.Error(s)
	}
}
