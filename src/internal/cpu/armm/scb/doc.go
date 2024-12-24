// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scb gives access to the registers of the System Control Block.
//
// Notes:
//
// 1. Cortex-M0 does not implement SHPR1, CFSR, HFSR, MMFR, BFAR, AFSR
// registers.
//
// 2. Cortex-M0 supports only word access to SHPR2, SHPR3 so this package does
// not provide a byte access to the individual fields.
package scb