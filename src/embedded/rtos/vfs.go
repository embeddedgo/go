// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import (
	"internal/bytealg"
	"io/fs"
	"path"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// An FS interface is the minimum implementation of a hierarchical file system
// required by the Virtual File System (VFS) provided for os package.
//
// Any FS implementation can also easily implement the fs.FS interface (and
// vice versa) thanks to the use of the common fs.File interface to represent an
// open file.
//
// The FS may implement the following optional methods recognized by VFS:
//
//	// Sync synchronizes cached writes to persistent storage. It can be called
//	// at any time but it is specifically called by Unmount function after
//	// unmounting the file system from VFS.
//	Sync() error
//
//	Chmod(name string, mode fs.FileMode) error // see os.Chmod
//	Mkdir(name, perm fs.FileMode) error        // see os.Mkdir
//	Remove(name string) error                  // see os.Remove
//	Rename(oldname, newname string) error      // see os.Rename
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

	// Usage returns the filesystem usage statistics. All four values are
	// subject to change while the file system is used. Usage should return
	// -1 for any unknown value.
	Usage() (usedItems, maxItems int, usedBytes, maxBytes int64)
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
	mu     sync.RWMutex
	mounts []*MountPoint //  TODO: linked list?
}

var mtab mountTable

// Mount mounts a filesystem with a provided prefix. Prefix can be any
// slash-separated path and does not have to represent an existing directory
// (in this respect it is similar to URL path). Mounted filesystem becomes
// available to os package.
func Mount(fsys FS, prefix string) error {
	if prefix == "" || prefix[0] != '/' || fsys == nil {
		return &fs.PathError{Op: "mount", Path: prefix, Err: syscall.EINVAL}
	}
	prefix = path.Clean(prefix)
	mtab.mu.Lock()
	mtab.mounts = append(mtab.mounts, &MountPoint{prefix, fsys, 0})
	mtab.mu.Unlock()
	return nil
}

// Unmount unmounts the last mounted filesystem that match fsys and prefix.
// At least one parameter must be specified (not empty or nil).
func Unmount(fsys FS, prefix string) error {
	if prefix == "" && fsys == nil {
		return &fs.PathError{Op: "unmount", Path: prefix, Err: syscall.ENOENT}
	}
	prefix = path.Clean(prefix)
	mtab.mu.Lock()
	remove := -1
	for i := len(mtab.mounts) - 1; i >= 0; i-- {
		mp := mtab.mounts[i]
		if (prefix == "" || mp.Prefix == prefix) && (fsys == nil || mp.FS == fsys) {
			remove = i
			fsys = mp.FS
			break
		}
	}
	var err error
	if remove < 0 {
		err = syscall.ENOENT
		goto skip
	}
	if atomic.LoadInt32(&mtab.mounts[remove].OpenCount) != 0 {
		err = syscall.EBUSY
		goto skip
	}
	copy(mtab.mounts[remove:], mtab.mounts[remove+1:])
	mtab.mounts = mtab.mounts[:len(mtab.mounts)-1]
skip:
	mtab.mu.Unlock()
	if err != nil {
		return &fs.PathError{Op: "unmount", Path: prefix, Err: err}
	}

	// close the fsys if it has no another mount point

	mtab.mu.RLock()
	for _, mp := range mtab.mounts {
		if mp.FS == fsys {
			fsys = nil // fsys is still mounted with another prefix
			break
		}
	}
	mtab.mu.RUnlock()
	if fsys == nil {
		return nil
	}
	if fsys, ok := fsys.(interface{ Sync() error }); ok {
		if err = fsys.Sync(); err != nil {
			return &fs.PathError{Op: "unmount", Path: prefix, Err: err}
		}
	}
	return nil
}

// Mounts returns the current list of mount points.
func Mounts() []*MountPoint {
	mtab.mu.RLock()
	list := append([]*MountPoint{}, mtab.mounts...)
	mtab.mu.RUnlock()
	return list
}

func findMountPoint(name string) (mp *MountPoint, fsys FS, unrooted string) {
	name = path.Clean(name)
	if name[0] == '/' || name == "." {
		name = name[1:] // remove leading "/", replace "." or "/" to ""
	}
	// Search the latest mounted FS whose prefix matches the dir.
	for i := len(mtab.mounts) - 1; i >= 0; i-- {
		mp = mtab.mounts[i]
		prefix := mp.Prefix[1:] // remove leading "/"
		switch {
		case len(name) < len(prefix):
			if prefix[len(name)] == '/' && prefix[:len(name)] == name {
				// Dir matches a part of prefix
				return nil, nil, name // virtual directory
			}
		case len(name) == len(prefix):
			if name == prefix {
				return mp, mp.FS, "."
			}
		default: // len(name) > len(prefix)
			if name[len(prefix)] == '/' && name[:len(prefix)] == prefix {
				name = name[len(prefix)+1:]
				if name == "" {
					name = "."
				}
			} else if prefix != "" {
				break
			}
			return mp, mp.FS, name
		}
	}
	// Matching prefix not found.
	if name == "" {
		name = "/" // special case: "/" always exists, virtual directory
	}
	return nil, nil, name
}

func chmod(name string, mode fs.FileMode) error {
	mtab.mu.RLock()
	_, fsys, unrooted := findMountPoint(name)
	mtab.mu.RUnlock()
	if fsys == nil {
		return syscall.ENOENT
	}
	if fsys, ok := fsys.(interface {
		Chmod(name string, mode fs.FileMode) error
	}); ok {
		return fsys.Chmod(unrooted, mode)
	}
	// try to use File.Chmod
	f, err := fsys.OpenWithFinalizer(unrooted, syscall.O_RDONLY, 0, nil)
	if err != nil {
		return err
	}
	if f, ok := f.(interface {
		Chmod(mode fs.FileMode) error
	}); ok {
		err = f.Chmod(mode)
	} else {
		err = syscall.ENOTSUP
	}
	cerr := f.Close()
	if err != nil {
		return err
	}
	return cerr
}

func rename(oldname, newname string) error {
	mtab.mu.RLock()
	_, oldfs, oldunrooted := findMountPoint(oldname)
	_, newfs, newunrooted := findMountPoint(newname)
	mtab.mu.RUnlock()
	if oldfs == nil || newfs == nil {
		return syscall.ENOENT
	}
	if oldfs == newfs {
		if fsys, ok := oldfs.(interface {
			Rename(oldname, newname string) error
		}); ok {
			return fsys.Rename(oldunrooted, newunrooted)
		}
	}
	return syscall.ENOTSUP
}

func mkdir(name string, perm fs.FileMode) (err error) {
	mtab.mu.RLock()
	_, fsys, unrooted := findMountPoint(name)
	mtab.mu.RUnlock()
	if fsys == nil {
		err = syscall.ENOENT
		goto error
	}
	if fsys, ok := fsys.(interface {
		Mkdir(name string, perm fs.FileMode) error
	}); ok {
		if err = fsys.Mkdir(unrooted, perm); err != nil {
			goto error
		}
		return nil
	}
	err = syscall.ENOTSUP
error:
	return &fs.PathError{Op: "mkdir", Path: name, Err: err}
}

func remove(name string) (err error) {
	mtab.mu.RLock()
	_, fsys, unrooted := findMountPoint(name)
	mtab.mu.RUnlock()
	if fsys == nil {
		err = syscall.ENOENT
		goto error
	}
	if fsys, ok := fsys.(interface {
		Remove(name string) error
	}); ok {
		if err = fsys.Remove(unrooted); err != nil {
			goto error
		}
		return nil
	}
	err = syscall.ENOTSUP
error:
	return &fs.PathError{Op: "remove", Path: name, Err: err}
}

func openFile(name string, flag int, perm fs.FileMode) (f fs.File, err error) {
	mtab.mu.RLock()
	mp, fsys, unrooted := findMountPoint(name)
	if mp != nil {
		if atomic.AddInt32(&mp.OpenCount, 1) < 0 {
			atomic.AddInt32(&mp.OpenCount, -1)
			err = syscall.EMFILE
		}
	}
	mtab.mu.RUnlock()
	if err != nil {
		return nil, err
	}
	if mp != nil {
		return fsys.OpenWithFinalizer(unrooted, flag, perm, mp.closed)
	}
	if unrooted != "" {
		if unrooted == "/" {
			unrooted = ""
		}
		return &vdir{unrooted, &vfi_}, nil
	}
	return nil, syscall.ENOENT
}

// a vdir represents a fake directory at the level of a mount-point prefix
type vdir struct {
	name string
	fs.FileInfo
}

func (d *vdir) Read(p []byte) (int, error) { return 0, syscall.EISDIR }
func (d *vdir) Stat() (fs.FileInfo, error) { return d, nil }
func (d *vdir) Close() error               { return nil }

func (d *vdir) ReadDir(n int) (dl []fs.DirEntry, err error) {
	nlen := len(d.name)
	mtab.mu.RLock()
	for _, mp := range mtab.mounts {
		prefix := mp.Prefix
		plen := len(prefix)
		if plen <= nlen || prefix[nlen] != '/' || d.name != prefix[:nlen] {
			continue
		}
		d := &vdir{name: prefix[nlen+1:]}
		if i := bytealg.IndexByteString(d.name, '/'); i >= 0 {
			d.name = d.name[:i]
			d.FileInfo = &vfi_
		} else {
			f, err := openFile(prefix, syscall.O_RDONLY, 0)
			if err != nil {
				return nil, err
			}
			d.FileInfo, err = f.Stat()
			if err != nil {
				return nil, err
			}
			f.Close()
		}
		dl = append(dl, d)
	}
	mtab.mu.RUnlock()
	return
}

// Additional methods to implement fs.DirEntry interface

// overwrite d.FileInfo Name method
func (d *vdir) Name() string               { return d.name }
func (d *vdir) Type() fs.FileMode          { return d.Mode() }
func (d *vdir) Info() (fs.FileInfo, error) { return d, nil }

var vfi_ vfi

type vfi struct{}

func (_ *vfi) Name() string       { return "" }
func (_ *vfi) Size() int64        { return 0 }
func (_ *vfi) IsDir() bool        { return true }
func (_ *vfi) Sys() interface{}   { return nil }
func (_ *vfi) ModTime() time.Time { return time.Time{} }
func (_ *vfi) Mode() fs.FileMode  { return fs.ModeDir | 0555 }
