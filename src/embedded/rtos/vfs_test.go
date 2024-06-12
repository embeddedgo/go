// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import (
	"errors"
	"io"
	"io/fs"
	"path"
	"testing"
)

type testfile struct {
	closed func()
}

func (f *testfile) Stat() (fi fs.FileInfo, err error) { return }
func (f *testfile) Read(p []byte) (int, error)        { return 0, io.EOF }

func (f *testfile) Close() error {
	if f.closed != nil {
		f.closed()
	}
	f.closed = nil
	return nil
}

type testfs struct {
	name string
}

func (fs *testfs) OpenWithFinalizer(name string, flag int, perm fs.FileMode, closed func()) (fs.File, error) {
	return &testfile{closed}, nil
}

func (fs *testfs) Type() string                    { return "testfs" }
func (fs *testfs) Name() string                    { return fs.name }
func (fs *testfs) Usage() (int, int, int64, int64) { return -1, -1, -1, -1 }

type test struct {
	prefix string
	fs     FS
}

var (
	fs1 = &testfs{name: "FS1"}
	fs2 = &testfs{name: "FS2"}

	tests = []test{
		{"/fs1/prefix1/", fs1},
		{"/fs1/prefix2", fs1},
		{"/fs2//prefix1", fs2},
		{"/../fs2/prefix2", fs2},
		{"/common/prefix", fs1},
		{"/common/prefix", fs2},
	}
)

func checkMounts(t *testing.T) {
	mounts := Mounts()
	i := 0
	for _, test := range tests {
		if test.fs == nil {
			continue
		}
		mp := mounts[i]
		if mp.Prefix != path.Clean(test.prefix) {
			t.Fatal(i, mp.Prefix, "!=", tests[i].prefix)
		}
		if mp.FS != test.fs {
			t.Fatal(i, mp.FS.Name(), "!=", tests[i].fs.Name())
		}
		i++
	}
}

func TestMountOpen(t *testing.T) {
	// mount
	for _, test := range tests {
		if err := Mount(test.fs, test.prefix); err != nil {
			t.Fatal(err)
		}
	}
	checkMounts(t)
	if err := Unmount(nil, "/"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal("expected fs.ErrNotExist, got:", err)
	}

	// unmount
	if err := Unmount(nil, "//fs2/prefix2"); err != nil {
		t.Fatal(err)
	}
	tests[3].fs = nil
	checkMounts(t)
	if err := Unmount(nil, "/common/prefix/"); err != nil {
		t.Fatal(err)
	}
	tests[5].fs = nil
	checkMounts(t)

	// open
	f, err := openFile("/fs2/prefix1", 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if oc := mtab.mounts[2].OpenCount; oc != 1 {
		t.Fatalf("OpenCount=%d, expexted 1:", oc)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	if oc := mtab.mounts[2].OpenCount; oc != 0 {
		t.Fatalf("OpenCount=%d, expexted 0:", oc)
	}
}
