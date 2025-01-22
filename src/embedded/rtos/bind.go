// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

// NotBound means the thread is not tied to any execution context.
const NotBound ExeCtx = notBound

// An ExeCtx represents a thread execution context. It means the execution
// resources used to execute a thread and/or the environment in which a thread
// is executed. The actual meaning is target specific and may be a CPU/core,
// a set of CPUs/cores, a NUMA node, etc.
type ExeCtx exeCtx

const notBound ExeCtx = -1

type exeCtx int

// Bind binds the current thread to the execution context ctx. It returns the
// previous execution context the thread was tied end an error.
func Bind(ctx ExeCtx) (oldctx ExeCtx, err error) {
	return bind(ctx)
}
