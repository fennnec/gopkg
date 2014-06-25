// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

const (
	errNil = iota
	errInvalidArgument
	errNotFound
	errCorruption
	errIOError
	errUnknown
)

// Error represents an LevelDB error.
type Error struct {
	code int
	err  string
}

func newError(code int, err string) *Error {
	return &Error{code: code, err: err}
}

func (e *Error) InvalidArgument() bool {
	return e.code == errInvalidArgument
}

func (e *Error) IsNotFound() bool {
	return e.code == errNotFound
}

func (e *Error) IsCorruption() bool {
	return e.code == errCorruption
}

func (e *Error) IsIOError() bool {
	return e.code == errIOError
}

func (e *Error) IsUnknown() bool {
	return e.code == errUnknown
}

func (e *Error) Error() string {
	if e == nil || e.code == errNil {
		return "leveldb: <nil>"
	}
	switch e.code {
	case errInvalidArgument:
		return "leveldb: Invalid Argument"
	case errNotFound:
		return "leveldb: Not Found"
	case errCorruption:
		return "leveldb: Corruption"
	case errIOError:
		return "leveldb: IO Error"
	}
	return "leveldb: " + e.err
}
