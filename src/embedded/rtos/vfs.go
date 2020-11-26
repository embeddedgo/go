// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import (
	"io/fs"
	"path"
	"sync"
	"sync/atomic"
	"syscall"
)

// An FS interface is the minimum implementation of a hierarchical file system
// required by the Virtual File System (VFS) provided for os package.
//
// Any FS implementation can also easily implement the fs.FS interface (and
// vice versa) thanks to the use of the common fs.File interface to represent an
// open file.
type FS interface {
	// OpenWithFinalizer works like the os.OpenFile function but has an
	// additional parameter. The returned fs.File implementation is obliged to
	// call the closed function if the underlying file has been closed.
	OpenWithFinalizer(name string, flag int, perm fs.FileMode, closed func()) (fs.File, error)

	// Type returns the file system type, e.g. fat32, ext4, nfs, ramfs. The type
	// name cannot be an empty string. It is recommended to use only lowercase
	// letters and numbers in the type name.
	Type() string

	// Name returns the name (label) of a specific file system instance.
	Name() string
}

// An UsageFS is a file system with an Usage method.
type UsageFS interface {
	FS

	// Usage returns the filesystem usage statistics. All four values are
	// subject to change while the file system is used. Usage should return
	// -1 for any unknown value.
	Usage() (usedItems, maxItems int, usedBytes, maxBytes int64)
}

// A MkdirFS is a file system with a Mkdir method.
type MkdirFS interface {
	FS

	Mkdir(name, perm fs.FileMode) error
}

// A SyncFS is a file system with a Sync method.
type SyncFS interface {
	FS

	// Sync synchronizes cached writes to persistent storage. It can be called
	// at any time but it is specifically called by Unmount function after
	// unmounting the file system from VFS.
	Sync() error
}

// A MountPoint represents a mounted file system.
type MountPoint struct {
	Prefix    string // path to FS
	FS        FS     // mounted file system
	OpenCount int32  // number of open files
}

func (mp *MountPoint) closed() {
	if atomic.AddInt32(&mp.OpenCount, -1) < 0 {
		panic("open count < 0")
	}
}

type mountTable struct {
	lock   sync.RWMutex
	mounts []*MountPoint //  TODO: linked list?
}

var mtab mountTable

// Mount mounts a filesystem with a provided prefix. Prefix can be any
// slash-separated path and does not have to represent an existing directory
// (in this respect it is similar to URL path). Mounted filesystem becomes
// available to os package.
func Mount(prefix string, fsys FS) error {
	if prefix == "" || prefix[0] != '/' || fsys == nil {
		return &fs.PathError{Op: "mount", Path: prefix, Err: syscall.EINVAL}
	}
	prefix = path.Clean(prefix)
	mtab.lock.Lock()
	mtab.mounts = append(mtab.mounts, &MountPoint{prefix, fsys, 0})
	mtab.lock.Unlock()
	return nil
}

// Unmount unmounts the last mounted filesystem that match prefix and/or fsys.
// At least one parameter must be specified (not empty or nil).
func Unmount(prefix string, fsys FS) error {
	if prefix == "" && fsys == nil {
		return &fs.PathError{Op: "unmount", Path: prefix, Err: syscall.ENOENT}
	}
	prefix = path.Clean(prefix)
	mtab.lock.Lock()
	remove := -1
	for i := len(mtab.mounts) - 1; i >= 0; i-- {
		mp := mtab.mounts[i]
		if (prefix == "" || mp.Prefix == prefix) && (fsys == nil || mp.FS == fsys) {
			remove = i
			fsys = mp.FS
			break
		}
	}
	if remove < 0 {
		mtab.lock.Unlock()
		return &fs.PathError{Op: "unmount", Path: prefix, Err: syscall.ENOENT}
	}
	if atomic.LoadInt32(&mtab.mounts[remove].OpenCount) != 0 {
		mtab.lock.Unlock()
		return &fs.PathError{Op: "unmount", Path: prefix, Err: syscall.EBUSY}
	}
	copy(mtab.mounts[remove:], mtab.mounts[remove+1:])
	mtab.mounts = mtab.mounts[:len(mtab.mounts)-1]
	mtab.lock.Unlock()

	// close the fsys if it has no another mount point

	mtab.lock.RLock()
	defer mtab.lock.RUnlock()
	for _, mp := range mtab.mounts {
		if mp.FS == fsys {
			fsys = nil // fsys is still mounted with another prefix
			break
		}
	}
	if fsys == nil {
		return nil
	}
	if syncfs, ok := fsys.(SyncFS); ok {
		if err := syncfs.Sync(); err != nil {
			return &fs.PathError{Op: "unmount", Path: prefix, Err: err}
		}
	}
	return nil
}

// Mounts returns the current list of mount points.
func Mounts() []*MountPoint {
	mtab.lock.RLock()
	defer mtab.lock.RUnlock()
	return append([]*MountPoint{}, mtab.mounts...)
}

func findMountPoint(name string) (mp *MountPoint, unrooted string, err error) {
	name = path.Clean(name)
	for i := len(mtab.mounts) - 1; i >= 0; i-- {
		mp = mtab.mounts[i]
		plen := len(mp.Prefix)
		if len(name) < plen || name[:plen] != mp.Prefix {
			continue
		}
		if len(name) == plen {
			unrooted = "."
			break
		}
		if name[plen] == '/' {
			unrooted = name[plen+1:]
			break
		}
	}
	if unrooted == "" {
		return nil, "", syscall.ENOENT
	}
	return
}

func openFile(name string, flag int, perm fs.FileMode) (fs.File, error) {
	mtab.lock.RLock()
	defer mtab.lock.RUnlock()
	mp, unrooted, err := findMountPoint(name)
	if err != nil {
		return nil, err
	}
	if atomic.AddInt32(&mp.OpenCount, 1) < 0 {
		atomic.AddInt32(&mp.OpenCount, -1)
		return nil, err
	}
	return mp.FS.OpenWithFinalizer(unrooted, flag, perm, mp.closed)
}

func mkdir(name string, perm fs.FileMode) error {
	return syscall.ENOTSUP
}

func chmod(name string, mode fs.FileMode) error {
	mtab.lock.RLock()
	defer mtab.lock.RUnlock()
	mp, unrooted, err := findMountPoint(name)
	if err != nil {
		return err
	}
	if fsys, ok := mp.FS.(interface {
		Chmod(name string, mode fs.FileMode) error
	}); ok {
		if err := fsys.Chmod(unrooted, mode); err != nil {
			return err
		}
	}
	// try to use File.Chmod
	f, err := mp.FS.OpenWithFinalizer(unrooted, syscall.O_RDONLY, 0, nil)
	if err != nil {
		return err
	}
	defer f.Close()
	if ff, ok := f.(interface {
		Chmod(mode fs.FileMode) error
	}); ok {
		if err := ff.Chmod(mode); err != nil {
			return err
		}
	}
	return syscall.ENOTSUP
}

func rename(oldname, newname string) error {
	mtab.lock.RLock()
	defer mtab.lock.RUnlock()
	mp, unrooted, err := findMountPoint(oldname)
	if err != nil {
		return err
	}
	newmp, newunrooted, err := findMountPoint(newname)
	if err != nil {
		return err
	}
	if mp == newmp {
		if fsys, ok := mp.FS.(interface {
			Rename(oldname, newname string) error
		}); ok {
			return fsys.Rename(unrooted, newunrooted)
		}
	}
	return syscall.ENOTSUP
}

func remove(name string) error {
	mtab.lock.RLock()
	defer mtab.lock.RUnlock()
	mp, unrooted, err := findMountPoint(name)
	if err != nil {
		return err
	}
	if fsys, ok := mp.FS.(interface {
		Remove(name string) error
	}); ok {
		return fsys.Remove(unrooted)
	}
	return syscall.ENOTSUP
}
