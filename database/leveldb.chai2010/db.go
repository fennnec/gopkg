// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import (
	"runtime"
)

const (
	MajorVersion = 1
	MinorVersion = 14
)

// Snapshot provides a consistent view of read operations in a DB. It is set
// on to a ReadOptions and passed in. It is only created by DB.NewSnapshot.
type Snapshot struct {
	snap *leveldb_snapshot_t
}

// Range is a key range.
type Range struct {
	Start []byte // Start of the key range, include in the range.
	Limit []byte // Limit of the key range, not include in the range.
}

// DB is a reusable handle to a LevelDB database on disk, created by Open.
//
// To avoid memory and file descriptor leaks, call Close when the process no
// longer needs the handle. Calls to any DB method made after Close will
// panic.
//
// The DB instance may be shared between goroutines. The usual data race
// conditions will occur if the same key is written to from more than one, of
// course.
type DB struct {
	*dbHandler
}
type dbHandler struct {
	db              *leveldb_t
	defaultOpt      *leveldb_options_t
	defaultReadOpt  *leveldb_readoptions_t
	defaultWriteOpt *leveldb_writeoptions_t
	optState        optStateContainer
	opt             Options
}

// Open opens a database.
//
// Creating a new database is done by calling SetCreateIfMissing(true) on the
// Options passed to Open.
//
// It is usually wise to set a Cache object on the Options with SetCache to
// keep recently used data from that database in memory.
func Open(name string, opt *Options) (*DB, error) {
	var err error

	h := &dbHandler{}
	if opt != nil {
		h.opt = *opt
	}
	h.defaultOpt = leveldb_options_create_copy(&h.optState, &h.opt)
	h.defaultReadOpt = leveldb_readoptions_create()
	h.defaultWriteOpt = leveldb_writeoptions_create()
	defer func() {
		if err != nil {
			leveldb_options_destroy(h.defaultOpt)
			leveldb_readoptions_destroy(h.defaultReadOpt)
			leveldb_writeoptions_destroy(h.defaultWriteOpt)
		}
	}()

	h.db, err = leveldb_open(name, h.defaultOpt)
	if err != nil {
		return nil, err
	}

	runtime.SetFinalizer(h, (*dbHandler).Close)
	return &DB{h}, nil
}

// Close closes the database, rendering it unusable for I/O, by deallocating
// the underlying handle.
//
// Any attempts to use the DB after Close is called will panic.
func (p *dbHandler) Close() {
	runtime.SetFinalizer(p, nil)

	leveldb_close(p.db)
	leveldb_options_destroy(p.defaultOpt)
	leveldb_readoptions_destroy(p.defaultReadOpt)
	leveldb_writeoptions_destroy(p.defaultWriteOpt)
	*p = dbHandler{}
}

// Get returns the data associated with the key from the database.
//
// If the key does not exist in the database, a nil Value is returned. If the
// key does exist, but the data is zero-length in the database, a zero-length
// Value will be returned.
//
// When the process no longer needs the Value, call Value.Release.
func (p *dbHandler) Get(key []byte, opt *ReadOptions) (*Value, error) {
	ro := p.defaultReadOpt
	if opt != nil {
		ro = leveldb_readoptions_create_copy(opt)
		defer leveldb_readoptions_destroy(ro)
	}
	d, err := leveldb_get(p.db, key, ro)
	if err != nil {
		return nil, err
	}
	return newValueHandler(d), nil
}

// Put writes data associated with a key to the database.
//
// If a nil []byte is passed in as value, it will be returned by Get as an
// zero-length slice.
//
// The key and value byte slices may be reused safely. Put takes a copy of
// them before returning.
func (p *dbHandler) Put(key, value []byte, opt *WriteOptions) error {
	wo := p.defaultWriteOpt
	if opt != nil {
		wo = leveldb_writeoptions_create_copy(opt)
		defer leveldb_writeoptions_destroy(wo)
	}
	return leveldb_put(p.db, key, value, wo)
}

// Delete removes the data associated with the key from the database.
//
// The key byte slice may be reused safely. Delete takes a copy of
// them before returning.
func (p *dbHandler) Delete(key []byte, opt *WriteOptions) error {
	wo := p.defaultWriteOpt
	if opt != nil {
		wo = leveldb_writeoptions_create_copy(opt)
		defer leveldb_writeoptions_destroy(wo)
	}
	return leveldb_delete(p.db, key, wo)
}

// Write atomically writes a WriteBatch to disk.
func (p *dbHandler) Write(batch *WriteBatch, opt *WriteOptions) error {
	wo := p.defaultWriteOpt
	if opt != nil {
		wo = leveldb_writeoptions_create_copy(opt)
		defer leveldb_writeoptions_destroy(wo)
	}
	return leveldb_write(p.db, batch.batch, wo)
}

// GetProperty returns the value of a database property.
//
// Examples of properties include "leveldb.stats", "leveldb.sstables",
// and "leveldb.num-files-at-level0".
//
//
// When the process no longer needs the Value, call Value.Release.
func (p *dbHandler) GetProperty(property string) *Value {
	return newValueHandler(leveldb_property_value(p.db, property))
}

// GetApproximateSizes returns the approximate number of bytes of file system
// space used by one or more key ranges.
//
// The keys counted will begin at Range.Start and end on the key before
// Range.Limit.
func (p *dbHandler) GetApproximateSizes(keyRanges []Range) []uint64 {
	rangeStartKey := make([][]byte, len(keyRanges))
	rangeLimitKey := make([][]byte, len(keyRanges))
	for i := 0; i < len(keyRanges); i++ {
		rangeStartKey[i] = keyRanges[i].Start
		rangeLimitKey[i] = keyRanges[i].Limit
	}
	return leveldb_approximate_sizes(p.db, rangeStartKey, rangeLimitKey)
}

// CompactRange runs a manual compaction on the Range of keys given. This is
// not likely to be needed for typical usage.
func (p *dbHandler) CompactRange(begin, end []byte) {
	leveldb_compact_range(p.db, begin, end)
}

// GetSnapshot creates a new snapshot of the database.
//
// The snapshot, when used in a ReadOptions, provides a consistent view of
// state of the database at the the snapshot was created.
//
// To prevent memory leaks and resource strain in the database, the snapshot
// returned must be released with DB.ReleaseSnapshot method on the DB that
// created it.
//
// See the LevelDB documentation for details.
func (p *dbHandler) GetSnapshot() *Snapshot {
	return &Snapshot{
		snap: leveldb_create_snapshot(p.db),
	}
}

// ReleaseSnapshot removes the snapshot from the database's list of snapshots,
// and deallocates it.
func (p *dbHandler) ReleaseSnapshot(snap *Snapshot) {
	leveldb_release_snapshot(p.db, snap.snap)
	*snap = Snapshot{}
}

// RepairDB attempts to repair a database.
// If the database is unrepairable, an error is returned.
func RepairDB(name string, opt *Options) error {
	var opt_ *leveldb_options_t
	var state optStateContainer
	if opt != nil {
		opt_ = leveldb_options_create_copy(&state, opt)
		defer leveldb_options_destroy(opt_)
	}
	return leveldb_repair_db(name, opt_)
}

// DestroyDB removes a database entirely,
// removing everything from the filesystem.
func DestroyDB(name string, opt *Options) error {
	var opt_ *leveldb_options_t
	var state optStateContainer
	if opt != nil {
		opt_ = leveldb_options_create_copy(&state, opt)
		defer leveldb_options_destroy(opt_)
	}
	return leveldb_destroy_db(name, opt_)
}
