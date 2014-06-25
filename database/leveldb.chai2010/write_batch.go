// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import (
	"runtime"
)

// WriteBatch is a batching of Puts, and Deletes to be written atomically to a
// database. A WriteBatch is written when passed to DB.Write.
type WriteBatch struct {
	*writeBatchHandler
}
type writeBatchHandler struct {
	batch    *leveldb_writebatch_t
	iterater stateWriteBatchIterater
}

// WriteBatchIterater iterates over a WriteBatch.
type WriteBatchIterater interface {
	Put(key, val []byte)
	Delete(key []byte)
}

// NewWriteBatch creates a fully allocated WriteBatch.
func NewWriteBatch() *WriteBatch {
	p := &writeBatchHandler{
		batch: leveldb_writebatch_create(),
	}
	runtime.SetFinalizer(p, (*writeBatchHandler).Release)
	return &WriteBatch{writeBatchHandler: p}
}

// Release releases the underlying memory of a WriteBatch.
func (p *writeBatchHandler) Release() {
	runtime.SetFinalizer(p, nil)
	leveldb_writebatch_destroy(p.batch)
	*p = writeBatchHandler{}
}

// Put places a key-value pair into the WriteBatch for writing later.
//
// Both the key and value byte slices may be reused as WriteBatch takes a copy
// of them before returning.
func (p *writeBatchHandler) Put(key, val []byte) {
	leveldb_writebatch_put(p.batch, key, val)
}

// Delete queues a deletion of the data at key to be deleted later.
//
// The key byte slice may be reused safely. Delete takes a copy of
// them before returning.
func (p *writeBatchHandler) Delete(key []byte) {
	leveldb_writebatch_delete(p.batch, key)
}

// Clear removes all the enqueued Put and Deletes in the WriteBatch.
func (p *writeBatchHandler) Clear() {
	leveldb_writebatch_clear(p.batch)
}

// Iterate iterating over the contents of a batch.
func (p *writeBatchHandler) Iterate(handler WriteBatchIterater) {
	p.iterater.WriteBatchIterater = handler
	leveldb_writebatch_iterate(p.batch, &p.iterater)
}
