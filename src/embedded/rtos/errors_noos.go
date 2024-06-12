// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

var errorsByNumber = [...]*Error{
	0: ErrUknown,
	1: ErrNotSupported,
	2: ErrBadPrivLevel,
	3: ErrInsufPrivLevel,
	4: ErrBadIntNumber,
	5: ErrBadIntPrio,
	6: ErrBadIntCtx,
}

func errnoError(errno int) error {
	if errno == 0 {
		return nil
	}
	if uint(errno) > uint(len(errorsByNumber)) {
		errno = 0
	}
	return errorsByNumber[errno]
}
