// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import "syscall"

func lookExtensions(path, dir string) (string, error) { return "", syscall.ENOTSUP }

func LookPath(file string) (string, error) { return "", syscall.ENOTSUP }
