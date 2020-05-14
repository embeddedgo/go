// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build riscv64

// Instances:
//  CLINT  0x2000000  -  -  Core Local Interrupter
// Registers:
//  0x000  32  MSIP[2]      Hart Software Interrupt Register
//  0x4000 64  MTIMECMP[2]  Hart Time Comparator Register
//  0xBFF8 64  MTIME        Timer Register
package clint
