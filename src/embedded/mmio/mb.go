// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmio

// MB is a full memory barrier. It ensures that any memory access (normal or
// I/O) after it, in program order, do not execute until all explicit memory
// accesses before it complete.
//
// Many embedded systems ensure the access order within their I/O address space
// but the order between load/store in normal RAM and I/O memory is not
// guaranteed mainly because of different buses for I/O and RAM. Some systems
// only ensure the access order to registers of the same peripheral. Use MB to
// properly order operations on different peripherals and RAM.
func MB()
