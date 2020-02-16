// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

type FD struct {
	// Lock sysfd and serialize access to Read and Write methods.
	fdmu fdMutex

	// Whether this is a file rather than a network socket.
	isFile bool
}

func (fd *FD) destroy() error { return nil }
