// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package semihostfs provieds access to files located on a debuging host.
// Debuger or emulator must support it.
//
// Example QEMU semihosting options:
//
// 	-serial none -semihosting --semihosting-config enable=on,target=native,userspace=on
//
package semihostfs
