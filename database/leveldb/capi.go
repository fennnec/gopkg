// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* Copyright (c) 2011 The LevelDB Authors. All rights reserved.
Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file. See the AUTHORS file for names of contributors.

C bindings for leveldb.  May be useful as a stable ABI that can be
used by programs that keep leveldb in a shared library, or for
a JNI api.

Does not support:
. getters for the option types
. custom comparators that implement key shortening
. capturing post-write-snapshot
. custom iter, db, env, cache implementations using just the C bindings

Some conventions:

(1) We expose just opaque struct pointers and functions to clients.
This allows us to change internal representations without having to
recompile clients.

(2) For simplicity, there is no equivalent to the Slice type.  Instead,
the caller has to pass the pointer and length as separate
arguments.

(3) Errors are represented by a null-terminated c string.  NULL
means no error.  All operations that can raise an error are passed
a "char** errptr" as the last argument.  One of the following must
be true on entry:
   *errptr == NULL
   *errptr points to a malloc()ed null-terminated error message
     (On Windows, *errptr must have been malloc()-ed by this library.)
On success, a leveldb routine leaves *errptr unchanged.
On failure, leveldb frees the old value of *errptr and
set *errptr to a malloc()ed error message.

(4) Bools have the type unsigned char (0 == false; rest == true)

(5) All of the pointer arguments must be non-NULL.
*/

package leveldb

/*
#cgo windows,amd64 LDFLAGS: -L. -l"leveldb-cgo-win64"
#cgo windows,386 LDFLAGS: -L. -l"leveldb-cgo-win32"
#cgo linux,amd64 LDFLAGS: -L. -l"leveldb-cgo-posix64"
#cgo linux,386 LDFLAGS: -L. -l"leveldb-cgo-posix32"

#cgo windows CFLAGS: -I./include -fno-stack-check -fno-stack-protector -mno-stack-arg-probe
#cgo linux   CFLAGS: -I./include

#include "capi_helper.h"
*/
import "C"
import "unsafe"

/* Exported types */
type (
	leveldb_t              C.leveldb_t
	leveldb_cache_t        C.leveldb_cache_t
	leveldb_comparator_t   C.leveldb_comparator_t
	leveldb_env_t          C.leveldb_env_t
	leveldb_filelock_t     C.leveldb_filelock_t
	leveldb_filterpolicy_t C.leveldb_filterpolicy_t
	leveldb_iterator_t     C.leveldb_iterator_t
	leveldb_logger_t       C.leveldb_logger_t
	leveldb_options_t      C.leveldb_options_t
	leveldb_randomfile_t   C.leveldb_randomfile_t
	leveldb_readoptions_t  C.leveldb_readoptions_t
	leveldb_seqfile_t      C.leveldb_seqfile_t
	leveldb_snapshot_t     C.leveldb_snapshot_t
	leveldb_writablefile_t C.leveldb_writablefile_t
	leveldb_writebatch_t   C.leveldb_writebatch_t
	leveldb_writeoptions_t C.leveldb_writeoptions_t
)

/* DB operations */

func leveldb_open(name string, options *leveldb_options_t) (*leveldb_t, error) {
	if options == nil {
		options = leveldb_options_create()
		defer leveldb_options_destroy(options)
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var cerr *C.char
	db := C.leveldb_open((*C.leveldb_options_t)(options), cname, &cerr)
	if cerr != nil {
		err := C.GoString(cerr)
		C.leveldb_free(unsafe.Pointer(cerr))
		return nil, leveldb_error(err)
	}
	return (*leveldb_t)(db), nil
}

func leveldb_close(db *leveldb_t) {
	C.leveldb_close((*C.leveldb_t)(db))
}

func leveldb_put(db *leveldb_t, key, val []byte, options *leveldb_writeoptions_t) error {
	if options == nil {
		options = leveldb_writeoptions_create()
		defer leveldb_writeoptions_destroy(options)
	}

	var cerr *C.char
	var ckey, cval *C.char
	if len(key) != 0 {
		ckey = (*C.char)(unsafe.Pointer(&key[0]))
	}
	if len(val) != 0 {
		cval = (*C.char)(unsafe.Pointer(&val[0]))
	}
	C.leveldb_put(
		(*C.leveldb_t)(db),
		(*C.leveldb_writeoptions_t)(options),
		ckey, C.size_t(len(key)),
		cval, C.size_t(len(val)),
		&cerr,
	)
	if cerr != nil {
		err := C.GoString(cerr)
		C.leveldb_free(unsafe.Pointer(cerr))
		return leveldb_error(err)
	}
	return nil
}

func leveldb_delete(db *leveldb_t, key []byte, options *leveldb_writeoptions_t) error {
	if options == nil {
		options = leveldb_writeoptions_create()
		defer leveldb_writeoptions_destroy(options)
	}

	var cerr *C.char
	var ckey *C.char
	if len(key) != 0 {
		ckey = (*C.char)(unsafe.Pointer(&key[0]))
	}
	C.leveldb_delete(
		(*C.leveldb_t)(db),
		(*C.leveldb_writeoptions_t)(options),
		ckey, C.size_t(len(key)),
		&cerr,
	)
	if cerr != nil {
		err := C.GoString(cerr)
		C.leveldb_free(unsafe.Pointer(cerr))
		return leveldb_error(err)
	}
	return nil
}

func leveldb_write(db *leveldb_t, batch *leveldb_writebatch_t, options *leveldb_writeoptions_t) error {
	if options == nil {
		options = leveldb_writeoptions_create()
		defer leveldb_writeoptions_destroy(options)
	}

	var cerr *C.char
	C.leveldb_write(
		(*C.leveldb_t)(db),
		(*C.leveldb_writeoptions_t)(options),
		(*C.leveldb_writebatch_t)(batch),
		&cerr,
	)
	if cerr != nil {
		err := C.GoString(cerr)
		C.leveldb_free(unsafe.Pointer(cerr))
		return leveldb_error(err)
	}
	return nil
}

func leveldb_get(db *leveldb_t, key []byte, options *leveldb_readoptions_t) ([]byte, error) {
	if options == nil {
		options = leveldb_readoptions_create()
		defer leveldb_readoptions_destroy(options)
	}

	var cerr *C.char
	var ckey *C.char
	if len(key) != 0 {
		ckey = (*C.char)(unsafe.Pointer(&key[0]))
	}
	var valLen C.size_t
	cval := C.leveldb_get(
		(*C.leveldb_t)(db),
		(*C.leveldb_readoptions_t)(options),
		ckey, C.size_t(len(key)),
		&valLen,
		&cerr,
	)
	if cerr != nil {
		err := C.GoString(cerr)
		C.leveldb_free(unsafe.Pointer(cerr))
		return nil, leveldb_error(err)
	}
	// https://code.google.com/p/leveldb/issues/detail?id=207
	if cval == nil {
		err := "NotFound:"
		return nil, leveldb_error(err)
	}

	val := C.GoBytes(unsafe.Pointer(cval), C.int(valLen))
	C.leveldb_free(unsafe.Pointer(cval))
	return val, nil
}

func leveldb_create_iterator(db *leveldb_t, options *leveldb_readoptions_t) *leveldb_iterator_t {
	if options == nil {
		options = leveldb_readoptions_create()
		defer leveldb_readoptions_destroy(options)
	}

	p := C.leveldb_create_iterator((*C.leveldb_t)(db), (*C.leveldb_readoptions_t)(options))
	return (*leveldb_iterator_t)(p)
}

func leveldb_create_snapshot(db *leveldb_t) *leveldb_snapshot_t {
	p := C.leveldb_create_snapshot((*C.leveldb_t)(db))
	return (*leveldb_snapshot_t)(p)
}

func leveldb_release_snapshot(db *leveldb_t, snapshot *leveldb_snapshot_t) {
	C.leveldb_release_snapshot((*C.leveldb_t)(db), (*C.leveldb_snapshot_t)(snapshot))
}

func leveldb_property_value(db *leveldb_t, name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cval := C.leveldb_property_value((*C.leveldb_t)(db), cname)
	if cval == nil {
		return ""
	}
	defer C.leveldb_free(unsafe.Pointer(cval))
	return C.GoString(cval)
}

func leveldb_approximate_sizes(db *leveldb_t, start, limit [][]byte) []uint64 {
	minLen := minInt(len(start), len(limit))
	if minLen <= 0 {
		return nil
	}
	cstart := make([]*C.char, minLen)
	climit := make([]*C.char, minLen)
	startLens := make([]C.size_t, minLen)
	limitLens := make([]C.size_t, minLen)
	sizes := make([]uint64, minLen)
	for i := 0; i < len(start); i++ {
		cstart[i] = C.CString(string(start[i]))
		defer C.free(unsafe.Pointer(cstart[i]))
		climit[i] = C.CString(string(limit[i]))
		defer C.free(unsafe.Pointer(climit[i]))
		startLens[i] = C.size_t(len(start[i]))
		limitLens[i] = C.size_t(len(limit[i]))
	}
	C.leveldb_approximate_sizes(
		(*C.leveldb_t)(db),
		C.int(minLen),
		&cstart[0], &startLens[0],
		&climit[0], &limitLens[0],
		(*C.uint64_t)(&sizes[0]),
	)
	return sizes
}

func leveldb_compact_range(db *leveldb_t, start, limit []byte) {
	var cstart, climit *C.char
	if len(start) != 0 {
		cstart = (*C.char)(unsafe.Pointer(&start[0]))
	}
	if len(limit) != 0 {
		climit = (*C.char)(unsafe.Pointer(&limit[0]))
	}
	C.leveldb_compact_range(
		(*C.leveldb_t)(db),
		cstart, C.size_t(len(start)),
		climit, C.size_t(len(limit)),
	)
}

/* Management operations */

func leveldb_destroy_db(name string, options *leveldb_options_t) error {
	if options == nil {
		options = leveldb_options_create()
		defer leveldb_options_destroy(options)
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var cerr *C.char
	C.leveldb_destroy_db((*C.leveldb_options_t)(options), cname, &cerr)
	if cerr != nil {
		err := C.GoString(cerr)
		C.leveldb_free(unsafe.Pointer(cerr))
		return leveldb_error(err)
	}
	return nil
}

func leveldb_repair_db(name string, options *leveldb_options_t) error {
	if options == nil {
		options = leveldb_options_create()
		defer leveldb_options_destroy(options)
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var cerr *C.char
	C.leveldb_repair_db((*C.leveldb_options_t)(options), cname, &cerr)
	if cerr != nil {
		err := C.GoString(cerr)
		C.leveldb_free(unsafe.Pointer(cerr))
		return leveldb_error(err)
	}
	return nil
}

/* Iterator */

func leveldb_iter_destroy(it *leveldb_iterator_t) {
	C.leveldb_iter_destroy((*C.leveldb_iterator_t)(it))
}

func leveldb_iter_valid(it *leveldb_iterator_t) bool {
	v := C.leveldb_iter_valid((*C.leveldb_iterator_t)(it))
	return v != C.uchar(0)
}

func leveldb_iter_seek_to_first(it *leveldb_iterator_t) {
	C.leveldb_iter_seek_to_first((*C.leveldb_iterator_t)(it))
}

func leveldb_iter_seek_to_last(it *leveldb_iterator_t) {
	C.leveldb_iter_seek_to_last((*C.leveldb_iterator_t)(it))
}

func leveldb_iter_seek(it *leveldb_iterator_t, key []byte) {
	C.leveldb_iter_seek(
		(*C.leveldb_iterator_t)(it),
		(*C.char)(unsafe.Pointer(&key[0])),
		C.size_t(len(key)),
	)
}

func leveldb_iter_next(it *leveldb_iterator_t) {
	C.leveldb_iter_next((*C.leveldb_iterator_t)(it))
}

func leveldb_iter_prev(it *leveldb_iterator_t) {
	C.leveldb_iter_prev((*C.leveldb_iterator_t)(it))
}

func leveldb_iter_key(it *leveldb_iterator_t) []byte {
	var klen C.size_t
	ckey := C.leveldb_iter_key((*C.leveldb_iterator_t)(it), &klen)
	key := C.GoBytes(unsafe.Pointer(ckey), C.int(klen))
	C.leveldb_free(unsafe.Pointer(ckey))
	return key
}

func leveldb_iter_value(it *leveldb_iterator_t) []byte {
	var vlen C.size_t
	cval := C.leveldb_iter_value((*C.leveldb_iterator_t)(it), &vlen)
	val := C.GoBytes(unsafe.Pointer(cval), C.int(vlen))
	C.leveldb_free(unsafe.Pointer(cval))
	return val
}

func leveldb_iter_get_error(it *leveldb_iterator_t) error {
	var cerr *C.char
	C.leveldb_iter_get_error((*C.leveldb_iterator_t)(it), &cerr)
	if cerr != nil {
		err := C.GoString(cerr)
		C.leveldb_free(unsafe.Pointer(cerr))
		return leveldb_error(err)
	}
	return nil
}

/* Write batch */

func leveldb_writebatch_create() *leveldb_writebatch_t {
	p := C.leveldb_writebatch_create()
	return (*leveldb_writebatch_t)(p)
}

func leveldb_writebatch_destroy(b *leveldb_writebatch_t) {
	C.leveldb_writebatch_destroy((*C.leveldb_writebatch_t)(b))
}

func leveldb_writebatch_clear(b *leveldb_writebatch_t) {
	C.leveldb_writebatch_clear((*C.leveldb_writebatch_t)(b))
}

func leveldb_writebatch_put(b *leveldb_writebatch_t, key, val []byte) {
	C.leveldb_writebatch_put(
		(*C.leveldb_writebatch_t)(b),
		(*C.char)(unsafe.Pointer(&key[0])), C.size_t(len(key)),
		(*C.char)(unsafe.Pointer(&val[0])), C.size_t(len(val)),
	)
}

func leveldb_writebatch_delete(b *leveldb_writebatch_t, key []byte) {
	C.leveldb_writebatch_delete(
		(*C.leveldb_writebatch_t)(b),
		(*C.char)(unsafe.Pointer(&key[0])), C.size_t(len(key)),
	)
}

func leveldb_writebatch_iterate(
	b *leveldb_writebatch_t,
	put func(key, val []byte),
	del func(key []byte),
) {
	state := leveldb_writebatch_iterate_state_create(put, del)
	C.go_leveldb_writebatch_iterate_helper((*C.leveldb_writebatch_t)(b), state)
}

//export go_leveldb_writebatch_iterate_put
func go_leveldb_writebatch_iterate_put(p unsafe.Pointer, k *C.char, klen C.size_t, v *C.char, vlen C.size_t) {
	state := (*leveldb_writebatch_iterate_state_t)(p)
	key := C.GoBytes(unsafe.Pointer(k), C.int(klen))
	val := C.GoBytes(unsafe.Pointer(v), C.int(vlen))
	state.put(key, val)
}

//export go_leveldb_writebatch_iterate_del
func go_leveldb_writebatch_iterate_del(p unsafe.Pointer, k *C.char, klen C.size_t) {
	state := (*leveldb_writebatch_iterate_state_t)(p)
	key := C.GoBytes(unsafe.Pointer(k), C.int(klen))
	state.del(key)
}

/* Options */

type leveldb_compression_t int

const (
	leveldb_no_compression     = leveldb_compression_t(C.leveldb_no_compression)
	leveldb_snappy_compression = leveldb_compression_t(C.leveldb_snappy_compression)
)

func leveldb_options_create() *leveldb_options_t {
	p := C.leveldb_options_create()
	return (*leveldb_options_t)(p)
}

func leveldb_options_destroy(opt *leveldb_options_t) {
	C.leveldb_options_destroy((*C.leveldb_options_t)(opt))
}

func leveldb_options_set_comparator(opt *leveldb_options_t, cmp *leveldb_comparator_t) {
	C.leveldb_options_set_comparator(
		(*C.leveldb_options_t)(opt),
		(*C.leveldb_comparator_t)(cmp),
	)
}

func leveldb_options_set_filter_policy(opt *leveldb_options_t, policy *leveldb_filterpolicy_t) {
	C.leveldb_options_set_filter_policy(
		(*C.leveldb_options_t)(opt),
		(*C.leveldb_filterpolicy_t)(policy),
	)
}

func leveldb_options_set_create_if_missing(opt *leveldb_options_t, create_if_missing bool) {
	var cval C.uchar
	if create_if_missing {
		cval = C.uchar(1)
	}
	C.leveldb_options_set_create_if_missing(
		(*C.leveldb_options_t)(opt),
		cval,
	)
}

func leveldb_options_set_error_if_exists(opt *leveldb_options_t, error_if_exists bool) {
	var cval C.uchar
	if error_if_exists {
		cval = C.uchar(1)
	}
	C.leveldb_options_set_error_if_exists(
		(*C.leveldb_options_t)(opt),
		cval,
	)
}

func leveldb_options_set_paranoid_checks(opt *leveldb_options_t, paranoid_checks bool) {
	var cval C.uchar
	if paranoid_checks {
		cval = C.uchar(1)
	}
	C.leveldb_options_set_paranoid_checks(
		(*C.leveldb_options_t)(opt),
		cval,
	)
}

func leveldb_options_set_env(opt *leveldb_options_t, env *leveldb_env_t) {
	C.leveldb_options_set_env(
		(*C.leveldb_options_t)(opt),
		(*C.leveldb_env_t)(env),
	)
}

func leveldb_options_set_info_log(opt *leveldb_options_t, logger *leveldb_logger_t) {
	C.leveldb_options_set_info_log(
		(*C.leveldb_options_t)(opt),
		(*C.leveldb_logger_t)(logger),
	)
}

func leveldb_options_set_write_buffer_size(opt *leveldb_options_t, size int) {
	C.leveldb_options_set_write_buffer_size(
		(*C.leveldb_options_t)(opt),
		C.size_t(size),
	)
}

func leveldb_options_set_max_open_files(opt *leveldb_options_t, size int) {
	C.leveldb_options_set_max_open_files(
		(*C.leveldb_options_t)(opt),
		C.int(size),
	)
}

func leveldb_options_set_cache(opt *leveldb_options_t, cache *leveldb_cache_t) {
	C.leveldb_options_set_cache(
		(*C.leveldb_options_t)(opt),
		(*C.leveldb_cache_t)(cache),
	)
}

func leveldb_options_set_block_size(opt *leveldb_options_t, size int) {
	C.leveldb_options_set_block_size(
		(*C.leveldb_options_t)(opt),
		C.size_t(size),
	)
}

func leveldb_options_set_block_restart_interval(opt *leveldb_options_t, size int) {
	C.leveldb_options_set_block_restart_interval(
		(*C.leveldb_options_t)(opt),
		C.int(size),
	)
}

func leveldb_options_set_compression(opt *leveldb_options_t, v leveldb_compression_t) {
	C.leveldb_options_set_compression(
		(*C.leveldb_options_t)(opt),
		C.int(v),
	)
}

/* Comparator */

func leveldb_comparator_create(
	destructor func(),
	compare func(a, b []byte) int,
	name func() string,
) *leveldb_comparator_t {
	state := leveldb_comparator_create_state_create(destructor, compare, name)
	p := C.leveldb_comparator_create_helper(state)
	return (*leveldb_comparator_t)(p)
}

func leveldb_comparator_destroy(cmp *leveldb_comparator_t) {
	C.leveldb_comparator_destroy((*C.leveldb_comparator_t)(cmp))
}

//export go_leveldb_comparator_create_state_destructor
func go_leveldb_comparator_create_state_destructor(p unsafe.Pointer) {
	state := (*leveldb_comparator_create_state_t)(p)
	if state.cname != nil {
		C.free(unsafe.Pointer(state.cname))
	}
	state.destructor()
}

//export go_leveldb_comparator_create_state_compare
func go_leveldb_comparator_create_state_compare(
	p unsafe.Pointer, a *C.char, alen C.size_t, b *C.char, blen C.size_t,
) int {
	state := (*leveldb_comparator_create_state_t)(p)
	ga := C.GoBytes(unsafe.Pointer(a), C.int(alen))
	gb := C.GoBytes(unsafe.Pointer(b), C.int(blen))
	return state.compare(ga, gb)
}

//export go_leveldb_comparator_create_state_name
func go_leveldb_comparator_create_state_name(p unsafe.Pointer) *C.char {
	state := (*leveldb_comparator_create_state_t)(p)
	return state.cname
}

/* Filter policy */

// panic: Memory Leak (Issue185)
func leveldb_filterpolicy_create(
	destructor func(),
	create_filter func(keys [][]byte) []byte,
	key_may_match func(key, filter []byte) bool,
	name func() string,
) *leveldb_filterpolicy_t {
	panic("TODO")
	/*
		leveldb_filterpolicy_t* leveldb_filterpolicy_create(
			void* state,
			void (*destructor)(void*),
			char* (*create_filter)(
				void*,
				const char* const* key_array, const size_t* key_length_array,
				int num_keys,
				size_t* filter_length),
			void (*free_filter)(void*),
			unsigned char (*key_may_match)(
				void*,
				const char* key, size_t length,
				const char* filter, size_t filter_length),
			const char* (*name)(void*),
		);
	*/
}

func leveldb_filterpolicy_destroy(policy *leveldb_filterpolicy_t) {
	C.leveldb_filterpolicy_destroy((*C.leveldb_filterpolicy_t)(policy))
}

func leveldb_filterpolicy_create_bloom(bits_per_key int) *leveldb_filterpolicy_t {
	p := C.leveldb_filterpolicy_create_bloom(C.int(bits_per_key))
	return (*leveldb_filterpolicy_t)(p)
}

/* Read options */

func leveldb_readoptions_create() *leveldb_readoptions_t {
	p := C.leveldb_readoptions_create()
	return (*leveldb_readoptions_t)(p)
}

func leveldb_readoptions_destroy(opt *leveldb_readoptions_t) {
	C.leveldb_readoptions_destroy((*C.leveldb_readoptions_t)(opt))
}

func leveldb_readoptions_set_verify_checksums(opt *leveldb_readoptions_t, verify_checksums bool) {
	var cval C.uchar
	if verify_checksums {
		cval = C.uchar(1)
	}
	C.leveldb_readoptions_set_verify_checksums((*C.leveldb_readoptions_t)(opt), cval)
}

func leveldb_readoptions_set_fill_cache(opt *leveldb_readoptions_t, fill_cache bool) {
	var cval C.uchar
	if fill_cache {
		cval = C.uchar(1)
	}
	C.leveldb_readoptions_set_fill_cache((*C.leveldb_readoptions_t)(opt), cval)
}

func leveldb_readoptions_set_snapshot(opt *leveldb_readoptions_t, snapshot *leveldb_snapshot_t) {
	C.leveldb_readoptions_set_snapshot(
		(*C.leveldb_readoptions_t)(opt),
		(*C.leveldb_snapshot_t)(snapshot),
	)
}

/* Write options */

func leveldb_writeoptions_create() *leveldb_writeoptions_t {
	p := C.leveldb_writeoptions_create()
	return (*leveldb_writeoptions_t)(p)
}

func leveldb_writeoptions_destroy(opt *leveldb_writeoptions_t) {
	C.leveldb_writeoptions_destroy((*C.leveldb_writeoptions_t)(opt))
}

func leveldb_writeoptions_set_sync(opt *leveldb_writeoptions_t, sync bool) {
	var cval C.uchar
	if sync {
		cval = C.uchar(1)
	}
	C.leveldb_writeoptions_set_sync((*C.leveldb_writeoptions_t)(opt), cval)
}

/* Cache */

func leveldb_cache_create_lru(capacity int) *leveldb_cache_t {
	p := C.leveldb_cache_create_lru(C.size_t(capacity))
	return (*leveldb_cache_t)(p)
}

func leveldb_cache_destroy(cache *leveldb_cache_t) {
	C.leveldb_cache_destroy((*C.leveldb_cache_t)(cache))
}

/* Env */

func leveldb_create_default_env() *leveldb_env_t {
	p := C.leveldb_create_default_env()
	return (*leveldb_env_t)(p)
}

func leveldb_env_destroy(env *leveldb_env_t) {
	C.leveldb_env_destroy((*C.leveldb_env_t)(env))
}

/* Utility */

func leveldb_free(p unsafe.Pointer) {
	C.leveldb_free(p)
}

func leveldb_major_version() int {
	v := C.leveldb_major_version()
	return int(v)
}

func leveldb_minor_version() int {
	v := C.leveldb_minor_version()
	return int(v)
}
