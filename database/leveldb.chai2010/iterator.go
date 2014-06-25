// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import (
	"runtime"
)

// Iterator iterates over a DB's key/value pairs in key order.
type Iterator struct {
	*iteratorHandler
}
type iteratorHandler struct {
	it *leveldb_iterator_t
	ro *leveldb_readoptions_t
}

// NewIterator returns an Iterator over the the database that uses the
// ReadOptions given.
//
// Often, this is used for large, offline bulk reads while serving live
// traffic. In that case, it may be wise to disable caching so that the data
// processed by the returned Iterator does not displace the already cached
// data. This can be done by calling SetFillCache(false) on the ReadOptions
// before passing it here.
//
// Similiarly, ReadOptions.SetSnapshot is also useful.
func (p *dbHandler) NewIterator(opt *ReadOptions) *Iterator {
	ro := leveldb_readoptions_create_copy(opt)
	it := leveldb_create_iterator(p.db, ro)
	h := &iteratorHandler{
		it: it,
		ro: ro,
	}
	runtime.SetFinalizer(h, (*iteratorHandler).Release)
	return &Iterator{h}
}

// Release releases the Iterator.
func (p *iteratorHandler) Release() {
	runtime.SetFinalizer(p, nil)

	leveldb_readoptions_destroy(p.ro)
	leveldb_iter_destroy(p.it)
	*p = iteratorHandler{}
}

// An iterator is either positioned at a key/value pair, or
// not valid.  This method returns true iff the iterator is valid.
func (p *iteratorHandler) Valid() bool {
	return leveldb_iter_valid(p.it)
}

// Position at the first key in the source.  The iterator is Valid()
// after this call iff the source is not empty.
func (p *iteratorHandler) SeekToFirst() {
	leveldb_iter_seek_to_first(p.it)
}

// Position at the last key in the source.  The iterator is
// Valid() after this call iff the source is not empty.
func (p *iteratorHandler) SeekToLast() {
	leveldb_iter_seek_to_last(p.it)
}

// Position at the first key in the source that at or past target
// The iterator is Valid() after this call iff the source contains
// an entry that comes at or past target.
func (p *iteratorHandler) Seek(target []byte) {
	leveldb_iter_seek(p.it, target)
}

// Moves to the next entry in the source.  After this call, Valid() is
// true iff the iterator was not positioned at the last entry in the source.
// REQUIRES: Valid()
func (p *iteratorHandler) Next() {
	leveldb_iter_next(p.it)
}

// Moves to the previous entry in the source.  After this call, Valid() is
// true iff the iterator was not positioned at the first entry in source.
// REQUIRES: Valid()
func (p *iteratorHandler) Prev() {
	leveldb_iter_prev(p.it)
}

// Return the key for the current entry.
//
// The underlying storage for the returned slice is valid only until
// the next modification of the iterator.
func (p *iteratorHandler) Key() []byte {
	k := leveldb_iter_key(p.it)
	return leveldb_slice_data(k)
}

// Return the value for the current entry.
//
// The underlying storage for the returned slice is valid only until
// the next modification of the iterator.
func (p *iteratorHandler) Value() []byte {
	k := leveldb_iter_value(p.it)
	return leveldb_slice_data(k)
}

// GetError returns an *Error from LevelDB if it had one during iteration.
//
// This method is safe to call when Valid returns false.
func (p *iteratorHandler) GetError() error {
	return leveldb_iter_get_error(p.it)
}
