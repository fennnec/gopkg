// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

/*
#include <stdlib.h>
#include <leveldb/c.h>

// This function exists only to clean up lack-of-const warnings when
// leveldb_approximate_sizes is called from Go-land.
static void go_leveldb_approximate_sizes(
	leveldb_t* db,
	int num_ranges,
	char** range_start_key, const size_t* range_start_key_len,
	char** range_limit_key, const size_t* range_limit_key_len,
	uint64_t* sizes
) {
	leveldb_approximate_sizes(db,
		num_ranges,
		(const char* const*)range_start_key,
		range_start_key_len,
		(const char* const*)range_limit_key,
		range_limit_key_len,
		sizes
	);
}
*/
import "C"

import (
	"unsafe"
)

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
	db *C.leveldb_t
}

// Range is a range of keys in the database. GetApproximateSizes calls with it
// begin at the key Start and end right before the key Limit.
type Range struct {
	Start []byte
	Limit []byte
}

// Snapshot provides a consistent view of read operations in a DB. It is set
// on to a ReadOptions and passed in. It is only created by DB.NewSnapshot.
//
// To prevent memory leaks and resource strain in the database, the snapshot
// returned must be released with DB.ReleaseSnapshot method on the DB that
// created it.
type Snapshot struct {
	snap *C.leveldb_snapshot_t
}

// Open opens a database.
//
// Creating a new database is done by calling SetCreateIfMissing(true) on the
// Options passed to Open.
//
// It is usually wise to set a Cache object on the Options with SetCache to
// keep recently used data from that database in memory.
func Open(name string, o *Options) (*DB, error) {
	var errStr *C.char
	dbname := C.CString(name)
	defer C.free(unsafe.Pointer(dbname))

	leveldb := C.leveldb_open(o.opt, dbname, &errStr)
	if errStr != nil {
		gs := C.GoString(errStr)
		C.leveldb_free(unsafe.Pointer(errStr))
		return nil, leveldb_error(gs)
	}
	return &DB{leveldb}, nil
}

// DestroyDatabase removes a database entirely, removing everything from the
// filesystem.
func DestroyDatabase(name string, o *Options) error {
	var errStr *C.char
	dbname := C.CString(name)
	defer C.free(unsafe.Pointer(dbname))

	C.leveldb_destroy_db(o.opt, dbname, &errStr)
	if errStr != nil {
		gs := C.GoString(errStr)
		C.leveldb_free(unsafe.Pointer(errStr))
		return leveldb_error(gs)
	}
	return nil
}

// RepairDatabase attempts to repair a database.
//
// If the database is unrepairable, an error is returned.
func RepairDatabase(name string, o *Options) error {
	var errStr *C.char
	dbname := C.CString(name)
	defer C.free(unsafe.Pointer(dbname))

	C.leveldb_repair_db(o.opt, dbname, &errStr)
	if errStr != nil {
		gs := C.GoString(errStr)
		C.leveldb_free(unsafe.Pointer(errStr))
		return leveldb_error(gs)
	}
	return nil
}

// Put writes data associated with a key to the database.
//
// If a nil []byte is passed in as value, it will be returned by Get as an
// zero-length slice.
//
// The key and value byte slices may be reused safely. Put takes a copy of
// them before returning.
func (p *DB) Put(wo *WriteOptions, key, value []byte) error {
	var errStr *C.char
	// leveldb_put, _get, and _delete call memcpy() (by way of Memtable::Add)
	// when called, so we do not need to worry about these []byte being
	// reclaimed by GC.
	var k, v *C.char
	if len(key) != 0 {
		k = (*C.char)(unsafe.Pointer(&key[0]))
	}
	if len(value) != 0 {
		v = (*C.char)(unsafe.Pointer(&value[0]))
	}
	C.leveldb_put(
		p.db, wo.opt,
		k, C.size_t(len(key)),
		v, C.size_t(len(value)),
		&errStr,
	)
	if errStr != nil {
		gs := C.GoString(errStr)
		C.leveldb_free(unsafe.Pointer(errStr))
		return leveldb_error(gs)
	}
	return nil
}

// Get returns the data associated with the key from the database.
//
// If the key does not exist in the database, a nil []byte is returned. If the
// key does exist, but the data is zero-length in the database, a zero-length
// []byte will be returned.
//
// The key byte slice may be reused safely. Get takes a copy of
// them before returning.
func (p *DB) Get(ro *ReadOptions, key []byte) ([]byte, error) {
	var errStr *C.char
	var vallen C.size_t
	var k *C.char
	if len(key) != 0 {
		k = (*C.char)(unsafe.Pointer(&key[0]))
	}
	value := C.leveldb_get(
		p.db, ro.opt,
		k, C.size_t(len(key)),
		&vallen,
		&errStr,
	)
	if errStr != nil {
		gs := C.GoString(errStr)
		C.leveldb_free(unsafe.Pointer(errStr))
		return nil, leveldb_error(gs)
	}
	// https://code.google.com/p/leveldb/issues/detail?id=207
	if value == nil {
		err := "NotFound:"
		return nil, leveldb_error(err)
	}

	defer C.leveldb_free(unsafe.Pointer(value))
	return C.GoBytes(unsafe.Pointer(value), C.int(vallen)), nil
}

// Delete removes the data associated with the key from the database.
//
// The key byte slice may be reused safely. Delete takes a copy of
// them before returning.
func (p *DB) Delete(wo *WriteOptions, key []byte) error {
	var errStr *C.char
	var k *C.char
	if len(key) != 0 {
		k = (*C.char)(unsafe.Pointer(&key[0]))
	}
	C.leveldb_delete(
		p.db, wo.opt,
		k, C.size_t(len(key)),
		&errStr,
	)
	if errStr != nil {
		gs := C.GoString(errStr)
		C.leveldb_free(unsafe.Pointer(errStr))
		return leveldb_error(gs)
	}
	return nil
}

// Write atomically writes a WriteBatch to disk.
func (p *DB) Write(wo *WriteOptions, w *WriteBatch) error {
	var errStr *C.char
	C.leveldb_write(p.db, wo.opt, w.wbatch, &errStr)
	if errStr != nil {
		gs := C.GoString(errStr)
		C.leveldb_free(unsafe.Pointer(errStr))
		return leveldb_error(gs)
	}
	return nil
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
func (p *DB) NewIterator(ro *ReadOptions) *Iterator {
	it := C.leveldb_create_iterator(p.db, ro.opt)
	return &Iterator{iter: it}
}

// GetApproximateSizes returns the approximate number of bytes of file system
// space used by one or more key ranges.
//
// The keys counted will begin at Range.Start and end on the key before
// Range.Limit.
func (p *DB) GetApproximateSizes(ranges []Range) []uint64 {
	starts := make([]*C.char, len(ranges))
	limits := make([]*C.char, len(ranges))
	startLens := make([]C.size_t, len(ranges))
	limitLens := make([]C.size_t, len(ranges))
	for i, r := range ranges {
		starts[i] = C.CString(string(r.Start))
		defer C.free(unsafe.Pointer(starts[i]))
		limits[i] = C.CString(string(r.Limit))
		defer C.free(unsafe.Pointer(limits[i]))
		startLens[i] = C.size_t(len(r.Start))
		limitLens[i] = C.size_t(len(r.Limit))
	}
	sizes := make([]uint64, len(ranges))
	C.go_leveldb_approximate_sizes(
		p.db,
		C.int(len(ranges)),
		&starts[0], &startLens[0],
		&limits[0], &limitLens[0],
		(*C.uint64_t)(&sizes[0]),
	)
	return sizes
}

// PropertyValue returns the value of a database property.
//
// Examples of properties include "leveldb.stats", "leveldb.sstables",
// and "leveldb.num-files-at-level0".
func (p *DB) PropertyValue(propName string) string {
	cname := C.CString(propName)
	defer C.free(unsafe.Pointer(cname))
	cval := C.leveldb_property_value(p.db, cname)
	if cval == nil {
		return ""
	}
	defer C.leveldb_free(unsafe.Pointer(cval))
	return C.GoString(cval)
}

// NewSnapshot creates a new snapshot of the database.
//
// The snapshot, when used in a ReadOptions, provides a consistent view of
// state of the database at the the snapshot was created.
//
// To prevent memory leaks and resource strain in the database, the snapshot
// returned must be released with DB.ReleaseSnapshot method on the DB that
// created it.
//
// See the LevelDB documentation for details.
func (p *DB) NewSnapshot() *Snapshot {
	return &Snapshot{C.leveldb_create_snapshot(p.db)}
}

// ReleaseSnapshot removes the snapshot from the database's list of snapshots,
// and deallocates it.
func (p *DB) ReleaseSnapshot(snap *Snapshot) {
	C.leveldb_release_snapshot(p.db, snap.snap)
}

// CompactRange runs a manual compaction on the Range of keys given. This is
// not likely to be needed for typical usage.
func (p *DB) CompactRange(r Range) {
	var start, limit *C.char
	if len(r.Start) != 0 {
		start = (*C.char)(unsafe.Pointer(&r.Start[0]))
	}
	if len(r.Limit) != 0 {
		limit = (*C.char)(unsafe.Pointer(&r.Limit[0]))
	}
	C.leveldb_compact_range(p.db,
		start, C.size_t(len(r.Start)),
		limit, C.size_t(len(r.Limit)),
	)
}

// Close closes the database, rendering it unusable for I/O, by deallocating
// the underlying handle.
//
// Any attempts to use the DB after Close is called will panic.
func (p *DB) Close() {
	C.leveldb_close(p.db)
}
