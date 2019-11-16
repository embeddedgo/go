// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Instances:
//  CMT  0xE000EF50  -  -  Cache maintenance registers
// Registers:
//  0x00 32  ICIALLU  Instr. cache invalidate all to the Point of Unification
//  0x08 32  ICIMVAU  Instr. cache invalidate by address to the PoU
//  0x0C 32  DCIMVAC  Data cache invalidate by address to the Point of Coherency
//  0x10 32  DCISW    Data cache invalidate by set/way
//  0x14 32  DCCMVAU  Data cache clean by address to the PoU
//  0x18 32  DCCMVAC  Data cache clean by address to the PoC
//  0x1c 32  DCCSW    Data cache clean by set/way
//  0x20 32  DCCIMVAC Data cache clean and invalidate by address to the PoC
//  0x24 32  DCCISW   Data cache clean and invalidate by set/way
package cmt
