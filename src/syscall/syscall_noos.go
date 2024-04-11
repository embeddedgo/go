// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import (
	"internal/oserror"
	"sync"
)

var (
	Stdin  = 0
	Stdout = 1
	Stderr = 2
)

type Error struct{ s string }

var (
	ENOTSUP      = &Error{"operation not supported"}
	EINVAL       = &Error{"operation not supported"}
	ENOENT       = &Error{"invalid argument"}
	EACCES       = &Error{"no such file or directory"}
	EPERM        = &Error{"operation not permitted"}
	EEXIST       = &Error{"file exists"}
	ENOTEMPTY    = &Error{"directory not empty"}
	EISDIR       = &Error{"is a directory"}
	ENOTDIR      = &Error{"not a directory"}
	ENAMETOOLONG = &Error{"file name too long"}
	EBUSY        = &Error{"device or resource busy"}
	EMFILE       = &Error{"too many open files"}
	ENOSPC       = &Error{"no space left on device"}
	EBADF        = &Error{"bad file descriptor"}
	ECANCELED    = &Error{"operation canceled"}
	EINTR        = &Error{"interrupt"}              // for os package only
	//ENOPROTOOPT  = &Error{"protocol not available"} // for net package only
)

func (e *Error) Is(target error) bool {
	switch target {
	case oserror.ErrPermission:
		return e == EACCES || e == EPERM
	case oserror.ErrExist:
		return e == EEXIST || e == ENOTEMPTY
	case oserror.ErrNotExist:
		return e == ENOENT
	}
	return false
}

func (e *Error) Error() string {
	return e.s
}

type Timespec struct {
	Sec  int64
	Nsec int64
}

type Timeval struct {
	Sec  int32
	Usec int32
}

var env = struct {
	mx   sync.RWMutex
	dict map[string]string
}{dict: map[string]string{
	"HOME": "/tmp",
}}

func Setenv(key, value string) error {
	env.mx.Lock()
	env.dict[key] = value
	env.mx.Unlock()
	return nil
}

func Getenv(key string) (value string, found bool) {
	env.mx.RLock()
	val, ok := env.dict[key]
	env.mx.RUnlock()
	return val, ok
}

func Unsetenv(key string) error {
	env.mx.Lock()
	delete(env.dict, key)
	env.mx.Unlock()
	return nil
}

func Clearenv() {
	env.mx.Lock()
	for key := range env.dict {
		delete(env.dict, key)
	}
	env.mx.Unlock()
}

func Environ() []string {
	env.mx.RLock()
	a := make([]string, 0, len(env.dict))
	for key, val := range env.dict {
		a = append(a, key+"="+val)
	}
	env.mx.RUnlock()
	return a
}

type SysProcAttr struct{}

const (
	O_RDONLY = 0x0
	O_WRONLY = 0x1
	O_RDWR   = 0x2
	O_CREAT  = 0x40
	O_EXCL   = 0x80
	O_TRUNC  = 0x200
	O_APPEND = 0x400
	O_SYNC   = 0x101000
)

type Qid struct {
	Path uint64 // the file server's unique identification for the file
}

type Dir struct {
	// system-modified data
	Type uint16 // server type
	Dev  uint32 // server subtype

	// file data
	Qid Qid // unique id from server
}

func Getpid() int {
	return 1
}

func Getppid() int {
	return 1
}

func Mkdir(path string, mode uint32) (err error)

func Chdir(path string) error {
	return ENOTSUP
}

const ImplementsGetwd = true

func Getwd() (wd string, err error) {
	return "", ENOTSUP
}

func Getuid() int {
	return -1
}

func Geteuid() int {
	return -1
}

func Getgid() int {
	return -1
}

func Getegid() int {
	return -1
}

func Getgroups() (gids []int, err error) {
	return nil, ENOTSUP
}

type Sockaddr struct{}
type SockaddrInet4 struct{}
type SockaddrInet6 struct{}
