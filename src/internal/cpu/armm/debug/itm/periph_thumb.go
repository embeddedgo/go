// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Peripheral: ITM_Periph
// Instances:
//  ITM  0xE0000000  -  -  Instrumentation Trace Macrocell registers
// Registers:
//  0x000  32  STIM[256]  Stimulus Port Registers
//  0xE00  32  TER[8]     Trace Enable Register
//  0xE40  32  TPR        Trace Privilege Register
//  0xE80  32  TCR        Trace Control Register
//  0xFD0  32  PID[8]     Peripheral Identification registers
//  0xFF0  32  CID[4]     Component Identification registers
package itm

const (
	ITMENA     TCR = 0x01 << 0  //+ enable ITM
	TSENA      TCR = 0x01 << 1  //+ enable local timestamp generation
	SYNCENA    TCR = 0x01 << 2  //+ enables packet transmission for a sync. TPIU
	TXENA      TCR = 0x01 << 3  //+ enable forwarding packets from DWT to ITM
	SWOENA     TCR = 0x01 << 4  //+ enable async. clocking of the timestamp cntr
	TSPrescale TCR = 0x03 << 8  //+ local timestamp prescaler
	TSP0       TCR = 0x00 << 8  //  no prescaling
	TSP4       TCR = 0x01 << 8  //  divide by 4
	TSP16      TCR = 0x02 << 8  //  divide by 16
	TSP64      TCR = 0x03 << 8  //  divide by 64
	GTSFREQ    TCR = 0x03 << 10 //+ global timestamp frequency
	TraceBusID TCR = 0x7F << 16 //+ identifier for multi-source trace stream
	BUSY       TCR = 0x01 << 23 //+ set if ITM is currently processing events
)
