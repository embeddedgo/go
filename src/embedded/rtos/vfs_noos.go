// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import _ "unsafe"

//go:linkname openFile os.openFile

//go:linkname chmod os.chmod
//go:linkname mkdir syscall.Mkdir
//go:linkname rename os.rename
//go:linkname remove os.Remove
