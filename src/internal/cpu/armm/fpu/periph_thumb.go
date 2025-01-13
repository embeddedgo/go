// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Instances:
//
//	FPU  0xe000ED88  -  -  Floating Point Unit registers
//
// Registers:
//
//	0x000  32  CPACR   Coprocessor Access Control Register
//	0x1AC  32  FPCCR   Floating-point Context Control Register
//	0x1B0  32  FPCAR   Floating-point Context Address Register
//	0x1B4  32  FPDSCR  Floating-point Default Status Control Register
package fpu

const (
	CP0  CPACR = 3 << 0  //+ Coprocessor 0 access privilege.
	CP1  CPACR = 3 << 2  //+ Coprocessor 1 access privilege.
	CP2  CPACR = 3 << 4  //+ Coprocessor 2 access privilege.
	CP3  CPACR = 3 << 6  //+ Coprocessor 3 access privilege.
	CP4  CPACR = 3 << 8  //+ Coprocessor 4 access privilege.
	CP5  CPACR = 3 << 10 //+ Coprocessor 5 access privilege.
	CP6  CPACR = 3 << 12 //+ Coprocessor 6 access privilege.
	CP7  CPACR = 3 << 14 //+ Coprocessor 7 access privilege.
	CP10 CPACR = 3 << 20 //+ Coprocessor 10 (FPU) access privilege.
	CP11 CPACR = 3 << 22 //+ Coprocessor 11 (FPU) access privilege.

	CPACDENY CPACR = 0
	CPACPRIV CPACR = 1
	CPACFULL CPACR = 3
)

const (
	CP0n  = 0
	CP1n  = 2
	CP2n  = 4
	CP3n  = 6
	CP4n  = 8
	CP5n  = 10
	CP6n  = 12
	CP7n  = 14
	CP10n = 20
	CP11n = 22
)

const (
	LSPACT    FPCCR = 1 << 0  //+ Lazy state preservation active.
	USER      FPCCR = 1 << 1  //+ FP stack frame allocated in unprivileged mode.
	S         FPCCR = 1 << 2  //+ Security status of the FP context.
	THREAD    FPCCR = 1 << 3  //+ FP stack frame allocated in thread mode.
	HFRDY     FPCCR = 1 << 4  //+
	MMRDY     FPCCR = 1 << 5  //+
	BFRDY     FPCCR = 1 << 6  //+
	MONRDY    FPCCR = 1 << 8  //+
	SPLIMVIOL FPCCR = 1 << 9  //+ Stack pointer limit violation.
	UFRDY     FPCCR = 1 << 10 //+
	TS        FPCCR = 1 << 26 //+ Treat FP registers as Secure enable.
	CLRONRETS FPCCR = 1 << 27 //+ CLRONRET Secure only.
	CLRONRET  FPCCR = 1 << 28 //+ Clear FP caller saved registers on exception return.
	LSPENS    FPCCR = 1 << 29 //+ LSPEN Secure only.
	LSPEN     FPCCR = 1 << 30 //+ Lazy context save of FP state.
	ASPEN     FPCCR = 1 << 31 //+ Enable CONTROL.FPCA setting on execution of a FP instruction.
)

const (
	ADDRESS FPCAR = 0x3fffffff << 2 //+
)

const (
	RMode FPDSCR = 3 << 22 //+
	FZ    FPDSCR = 1 << 24 //+
	DN    FPDSCR = 1 << 25 //+
	AHP   FPDSCR = 1 << 26 //+
)
