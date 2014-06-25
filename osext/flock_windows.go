// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package osext

import (
	"syscall"
)

// A FileLock is a across multiple processes file lock.
type FileLock struct {
	fd syscall.Handle
}

// NewFileLock try to lock the file, if failed return nil.
func NewFileLock(path string) (*FileLock, error) {
	pathp, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}
	fd, err := syscall.CreateFile(pathp,
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		0, nil, syscall.CREATE_ALWAYS,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if err != nil {
		return nil, err
	}
	return &FileLock{fd: fd}, nil
}

// Release releases the FileLock.
func (p *FileLock) Release() error {
	return syscall.Close(p.fd)
}
