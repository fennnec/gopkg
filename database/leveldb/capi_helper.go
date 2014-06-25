// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import "C"
import (
	"errors"
	"strings"
	"sync"
	"unsafe"
)

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/* Error types */
var (
	ErrNotFound        = errors.New("leveldb: not found")
	ErrCorruption      = errors.New("leveldb: corruption")
	ErrNotImplemented  = errors.New("leveldb: not implemented")
	ErrInvalidArgument = errors.New("leveldb: invalid argument")
	ErrIO              = errors.New("leveldb: IO error")
	ErrUnknown         = errors.New("leveldb: unknown")
)

func leveldb_error(e string) error {
	if e == "" || e == "OK" {
		return nil
	}
	if strings.HasPrefix(e, "NotFound:") {
		return ErrNotFound
	}
	if strings.HasPrefix(e, "Corruption:") {
		return ErrCorruption
	}
	if strings.HasPrefix(e, "Not implemented:") {
		return ErrNotImplemented
	}
	if strings.HasPrefix(e, "Invalid argument:") {
		return ErrInvalidArgument
	}
	if strings.HasPrefix(e, "IO error:") {
		return ErrIO
	}
	return ErrUnknown
}

/* Write batch */

type leveldb_writebatch_iterate_state_t struct {
	put func(key, val []byte)
	del func(key []byte)
}

// Avoid GC
var leveldb_writebatch_iterate_state_list []*leveldb_writebatch_iterate_state_t
var leveldb_writebatch_iterate_state_mutex sync.Mutex

func leveldb_writebatch_iterate_state_create(
	put func(key, val []byte),
	del func(key []byte),
) unsafe.Pointer {
	leveldb_writebatch_iterate_state_mutex.Lock()
	defer leveldb_writebatch_iterate_state_mutex.Unlock()
	p := &leveldb_writebatch_iterate_state_t{put: put, del: del}
	leveldb_writebatch_iterate_state_list = append(leveldb_writebatch_iterate_state_list, p)
	return unsafe.Pointer(p)
}

/* Comparator */

type leveldb_comparator_create_state_t struct {
	destructor func()
	compare    func(a, b []byte) int
	name       func() string
	cname      *C.char
}

// Avoid GC
var leveldb_comparator_create_state_list []*leveldb_comparator_create_state_t
var leveldb_comparator_create_state_mutex sync.Mutex

func leveldb_comparator_create_state_create(
	destructor func(),
	compare func(a, b []byte) int,
	name func() string,
) unsafe.Pointer {
	leveldb_comparator_create_state_mutex.Lock()
	defer leveldb_comparator_create_state_mutex.Unlock()
	cname := C.CString(name())
	p := &leveldb_comparator_create_state_t{
		destructor: destructor,
		compare:    compare,
		name:       name,
		cname:      cname,
	}
	leveldb_comparator_create_state_list = append(leveldb_comparator_create_state_list, p)
	return unsafe.Pointer(p)
}

/* Filter policy */

type leveldb_filterpolicy_create_state_t struct {
	destructor    func()
	free_filter   func()
	create_filter func(keys [][]byte) []byte
	key_may_match func(key, filter []byte) bool
	name          func() string
	cname         *C.char
	cfilter       *C.char
	cfilter_len   C.int
}

// Avoid GC
var leveldb_filterpolicy_create_state_list []*leveldb_filterpolicy_create_state_t
var leveldb_filterpolicy_create_state_mutex sync.Mutex

func leveldb_filterpolicy_create_state_create(
	destructor func(),
	create_filter func(keys [][]byte) []byte,
	key_may_match func(key, filter []byte) bool,
	name func() string,
) unsafe.Pointer {
	leveldb_filterpolicy_create_state_mutex.Lock()
	defer leveldb_filterpolicy_create_state_mutex.Unlock()
	cname := C.CString(name())
	p := &leveldb_filterpolicy_create_state_t{
		destructor:    destructor,
		create_filter: create_filter,
		key_may_match: key_may_match,
		name:          name,
		cname:         cname,
	}
	leveldb_filterpolicy_create_state_list = append(leveldb_filterpolicy_create_state_list, p)
	return unsafe.Pointer(p)
}
