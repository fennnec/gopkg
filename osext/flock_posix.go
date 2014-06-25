// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package osext

import (
	"os"
	"syscall"
)

// A FileLock is a across multiple processes file lock.
type FileLock struct {
	fp *os.File
}

// NewFileLock try to lock the file, if failed return nil.
func NewFileLock(path string) (*FileLock, error) {
	fp, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	if err = setFileLock(f, true); err != nil {
		f.Close()
		return nil, err
	}
	return &FileLock{fp: fp}, nil
}

// Release releases the FileLock.
func (p *FileLock) Release() error {
	if err := setFileLock(p.fp, false); err != nil {
		return err
	}
	return p.fp.Close()
}

func setFileLock(fp *os.File, lock bool) error {
	if lock {
		return syscall.Flock(int(fp.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	} else {
		return syscall.Flock(int(fp.Fd()), syscall.LOCK_UN|syscall.LOCK_NB)
	}
}
