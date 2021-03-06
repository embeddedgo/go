// DO NOT EDIT THIS FILE. GENERATED BY xgen.

package clint

import (
	"embedded/mmio"
	"unsafe"
)

type Periph struct {
	MSIP     [2]RMSIP
	_        [4094]uint32
	MTIMECMP [2]RMTIMECMP
	_        [8186]uint32
	MTIME    RMTIME
}

func CLINT() *Periph { return (*Periph)(unsafe.Pointer(uintptr(0x2000000))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type MSIP uint32

type RMSIP struct{ mmio.U32 }

func (r *RMSIP) LoadBits(mask MSIP) MSIP { return MSIP(r.U32.LoadBits(uint32(mask))) }
func (r *RMSIP) StoreBits(mask, b MSIP)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMSIP) SetBits(mask MSIP)       { r.U32.SetBits(uint32(mask)) }
func (r *RMSIP) ClearBits(mask MSIP)     { r.U32.ClearBits(uint32(mask)) }
func (r *RMSIP) Load() MSIP              { return MSIP(r.U32.Load()) }
func (r *RMSIP) Store(b MSIP)            { r.U32.Store(uint32(b)) }

type RMMSIP struct{ mmio.UM32 }

func (rm RMMSIP) Load() MSIP   { return MSIP(rm.UM32.Load()) }
func (rm RMMSIP) Store(b MSIP) { rm.UM32.Store(uint32(b)) }

type MTIMECMP uint64

type RMTIMECMP struct{ mmio.U64 }

func (r *RMTIMECMP) LoadBits(mask MTIMECMP) MTIMECMP { return MTIMECMP(r.U64.LoadBits(uint64(mask))) }
func (r *RMTIMECMP) StoreBits(mask, b MTIMECMP)      { r.U64.StoreBits(uint64(mask), uint64(b)) }
func (r *RMTIMECMP) SetBits(mask MTIMECMP)           { r.U64.SetBits(uint64(mask)) }
func (r *RMTIMECMP) ClearBits(mask MTIMECMP)         { r.U64.ClearBits(uint64(mask)) }
func (r *RMTIMECMP) Load() MTIMECMP                  { return MTIMECMP(r.U64.Load()) }
func (r *RMTIMECMP) Store(b MTIMECMP)                { r.U64.Store(uint64(b)) }

type RMMTIMECMP struct{ mmio.UM64 }

func (rm RMMTIMECMP) Load() MTIMECMP   { return MTIMECMP(rm.UM64.Load()) }
func (rm RMMTIMECMP) Store(b MTIMECMP) { rm.UM64.Store(uint64(b)) }

type MTIME uint64

type RMTIME struct{ mmio.U64 }

func (r *RMTIME) LoadBits(mask MTIME) MTIME { return MTIME(r.U64.LoadBits(uint64(mask))) }
func (r *RMTIME) StoreBits(mask, b MTIME)   { r.U64.StoreBits(uint64(mask), uint64(b)) }
func (r *RMTIME) SetBits(mask MTIME)        { r.U64.SetBits(uint64(mask)) }
func (r *RMTIME) ClearBits(mask MTIME)      { r.U64.ClearBits(uint64(mask)) }
func (r *RMTIME) Load() MTIME               { return MTIME(r.U64.Load()) }
func (r *RMTIME) Store(b MTIME)             { r.U64.Store(uint64(b)) }

type RMMTIME struct{ mmio.UM64 }

func (rm RMMTIME) Load() MTIME   { return MTIME(rm.UM64.Load()) }
func (rm RMMTIME) Store(b MTIME) { rm.UM64.Store(uint64(b)) }
