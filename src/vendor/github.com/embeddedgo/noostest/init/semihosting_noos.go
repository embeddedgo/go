// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package init

import (
	"embedded/rtos"
	"fmt"
	"os"

	"github.com/embeddedgo/fs/semihostfs"
)

func fatalErr(what string, err error) {
	if err != nil {
		if os.Stderr == nil {
			panic(fmt.Sprintf("%s: %v\n", what, err))
		}
		fmt.Fprintf(os.Stderr, "%s: %v\n", what, err)
		os.Stderr.Close()
	}
}

func debug(fd int, p []byte) int {
	n, _ := os.Stderr.Write(p)
	return n
}

func setupSemihosting() {
	fsys := semihostfs.New("SH1", "")
	rtos.Mount(fsys, "/")

	var err error
	os.Stderr, err = os.Open("/:stderr")
	fatalErr(":stderr", err)
	rtos.SetSystemWriter(debug)

	os.Stdout, err = os.Open("/:stdout")
	fatalErr(":stdout", err)

	os.Stdin, err = os.Open("/:stdin")
	fatalErr(":stdin", err)
}
