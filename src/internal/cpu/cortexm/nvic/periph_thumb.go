// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// BUG: Cortex-M0 supports only 32-bit access on the Private Peripheral Bus but
// the IPRs are defined as 8-bit registers. We prefer 8-bit access because it
// allows changing the priority of different interrupts concurrently. If we
// ever will support ARMv6-M all code that uses IPRs need to be fixed.

// Instances:
//  NVIC  0xE000E100  -  -  Nested Vectored Interrupt Controller
// Registers:
//  0x000  32  ISER[16]  Interrupt Set-Enable Registers
//  0x080  32  ICER[16]  Interrupt Clear-Enable Registers
//  0x100  32  ISPR[16]  Interrupt Set-Pending Registers
//  0x180  32  ICPR[16]  Interrupt Clear-Pending Registers
//  0x200  32  IABR[16]  Interrupt Active Bit Registers
//  0x280  32  ITNS[16]  Interrupt Target Non-secure Registers
//  0x300   8  IPR[496]  Interrupt Priority Registers
//  0xE00  32  STIR      Software Trigger Interrupt Register
package nvic
