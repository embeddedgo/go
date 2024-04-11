// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"syscall"
	"time"
)

type ProcessState struct{}

func findProcess(pid int) (p *Process, err error) {
	return &Process{Pid: pid}, nil
}

func startProcess(name string, argv []string, attr *ProcAttr) (p *Process, err error) {
	return nil, &PathError{"fork/exec", name, syscall.ENOTSUP}
}

func (p *Process) release() error               { return syscall.ENOTSUP }
func (p *Process) kill() error                  { return syscall.ENOTSUP }
func (p *Process) wait() (*ProcessState, error) { return nil, syscall.ENOTSUP }
func (p *Process) signal(sig Signal) error      { return syscall.ENOTSUP }

func (p *ProcessState) userTime() time.Duration   { return 0 }
func (p *ProcessState) systemTime() time.Duration { return 0 }
func (p *ProcessState) exited() bool              { return false }
func (p *ProcessState) success() bool             { return false }
func (p *ProcessState) sys() interface{}          { return nil }
func (p *ProcessState) sysUsage() interface{}     { return nil }

func executable() (string, error) { return "", ErrNotExist }
