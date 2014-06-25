// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

/*
#include "capi_helper.h"

#cgo windows CFLAGS: -I. -I./include -I./capi -fno-stack-check -fno-stack-protector -mno-stack-arg-probe
#cgo windows CXXFLAGS: -I. -I./include -I./capi -fno-stack-check -fno-stack-protector -mno-stack-arg-probe
#cgo windows LDFLAGS: -L./capi -lleveldb_c
// -static -static-libgcc -static-libstdc++

#cgo linux CFLAGS: -I. -I./include -I./capi
#cgo linux CXXFLAGS: -I. -I./include -I./capi
#cgo linux LDFLAGS: -static
*/
import "C"
import "unsafe"
import "reflect"

// ----------------------------------------------------------------------------
// Exported types
// ----------------------------------------------------------------------------

type (
	leveldb_t              C.leveldb_t
	leveldb_snapshot_t     C.leveldb_snapshot_t
	leveldb_options_t      C.leveldb_options_t
	leveldb_readoptions_t  C.leveldb_readoptions_t
	leveldb_writeoptions_t C.leveldb_writeoptions_t
	leveldb_writebatch_t   C.leveldb_writebatch_t
	leveldb_iterator_t     C.leveldb_iterator_t
	leveldb_comparator_t   C.leveldb_comparator_t
	leveldb_filterpolicy_t C.leveldb_filterpolicy_t
	leveldb_cache_t        C.leveldb_cache_t
	leveldb_slice_t        C.leveldb_slice_t
	leveldb_value_t        C.leveldb_value_t
	leveldb_compression_t  C.leveldb_compression_t
	leveldb_status_t       C.leveldb_status_t
	leveldb_bool_t         C.leveldb_bool_t
)

// --------------------------------------------------------

// leveldb_status_t
const (
	leveldb_status_ok               = C.leveldb_status_ok
	leveldb_status_invalid_argument = C.leveldb_status_invalid_argument
	leveldb_status_not_found        = C.leveldb_status_not_found
	leveldb_status_corruption       = C.leveldb_status_corruption
	leveldb_status_io_error         = C.leveldb_status_io_error
	leveldb_status_unknown          = C.leveldb_status_unknown
	leveldb_status_max              = C.leveldb_status_max
)

// leveldb_compression_t
const (
	leveldb_compression_nil    = C.leveldb_compression_nil
	leveldb_compression_snappy = C.leveldb_compression_snappy
	leveldb_compression_max    = C.leveldb_compression_max
)

// ----------------------------------------------------------------------------
// Version
// ----------------------------------------------------------------------------

func leveldb_major_version() int {
	v := C.leveldb_major_version()
	return int(v)
}
func leveldb_minor_version() int {
	v := C.leveldb_minor_version()
	return int(v)
}

// ----------------------------------------------------------------------------
// Slice
// ----------------------------------------------------------------------------

func leveldb_slice_new(data []byte) *leveldb_slice_t {
	if len(data) != 0 {
		return &leveldb_slice_t{
			data: (*C.char)(unsafe.Pointer(&data[0])),
			size: C.int32_t(len(data)),
		}
	} else {
		return &leveldb_slice_t{
			data: (*C.char)(unsafe.Pointer(nil)),
			size: C.int32_t(0),
		}
	}
}

func leveldb_slice_data(p *leveldb_slice_t) (data []byte) {
	h := (*reflect.SliceHeader)((unsafe.Pointer(&data)))
	h.Data = uintptr(unsafe.Pointer(p.data))
	h.Cap = int(p.size)
	h.Len = h.Cap
	return
}

// ----------------------------------------------------------------------------
// Value
// ----------------------------------------------------------------------------

func leveldb_value_create(data []byte) *leveldb_value_t {
	p := C.leveldb_value_create(
		(*C.char)(unsafe.Pointer(&data[0])),
		C.int32_t(len(data)),
	)
	return (*leveldb_value_t)(p)
}
func leveldb_value_create_copy(v *leveldb_value_t) *leveldb_value_t {
	p := C.leveldb_value_create_copy((*C.leveldb_value_t)(v))
	return (*leveldb_value_t)(p)
}
func leveldb_value_destroy(p *leveldb_value_t) {
	C.leveldb_value_destroy((*C.leveldb_value_t)(p))
}

func leveldb_value_data(p *leveldb_value_t) (data []byte) {
	h := (*reflect.SliceHeader)((unsafe.Pointer(&data)))
	h.Data = uintptr(unsafe.Pointer(C.leveldb_value_data((*C.leveldb_value_t)(p))))
	h.Cap = int(C.leveldb_value_size((*C.leveldb_value_t)(p)))
	h.Len = h.Cap
	return
}
func leveldb_value_cstr(p *leveldb_value_t) string {
	var data []byte
	h := (*reflect.SliceHeader)((unsafe.Pointer(&data)))
	h.Data = uintptr(unsafe.Pointer(C.leveldb_value_data((*C.leveldb_value_t)(p))))
	h.Cap = int(C.leveldb_value_size((*C.leveldb_value_t)(p)))
	h.Len = h.Cap
	return string(data)
}

// ----------------------------------------------------------------------------
// Error
// ----------------------------------------------------------------------------

func leveldb_error(e C.leveldb_status_t, errStr *C.leveldb_value_t) *Error {
	switch leveldb_status_t(e) {
	case leveldb_status_ok:
		return nil
	case leveldb_status_invalid_argument:
		return newError(errInvalidArgument, "")
	case leveldb_status_not_found:
		return newError(errNotFound, "")
	case leveldb_status_corruption:
		return newError(errCorruption, "")
	case leveldb_status_io_error:
		return newError(errIOError, "")
	}
	return newError(errUnknown, leveldb_value_cstr((*leveldb_value_t)(errStr)))
}

// ----------------------------------------------------------------------------
// Options
// ----------------------------------------------------------------------------

func leveldb_options_create() *leveldb_options_t {
	p := C.leveldb_options_create()
	return (*leveldb_options_t)(p)
}
func leveldb_options_destroy(p *leveldb_options_t) {
	C.leveldb_options_destroy((*C.leveldb_options_t)(p))
}

func leveldb_options_set_comparator(p *leveldb_options_t, v *leveldb_comparator_t) {
	C.leveldb_options_set_comparator((*C.leveldb_options_t)(p), (*C.leveldb_comparator_t)(v))
}
func leveldb_options_get_comparator(p *leveldb_options_t) *leveldb_comparator_t {
	v := C.leveldb_options_get_comparator((*C.leveldb_options_t)(p))
	return (*leveldb_comparator_t)(v)
}

func leveldb_options_set_filter_policy(p *leveldb_options_t, v *leveldb_filterpolicy_t) {
	C.leveldb_options_set_filter_policy((*C.leveldb_options_t)(p), (*C.leveldb_filterpolicy_t)(v))
}
func leveldb_options_get_filter_policy(p *leveldb_options_t) *leveldb_filterpolicy_t {
	v := C.leveldb_options_get_filter_policy((*C.leveldb_options_t)(p))
	return (*leveldb_filterpolicy_t)(v)
}

func leveldb_options_set_create_if_missing(p *leveldb_options_t, v bool) {
	if v {
		C.leveldb_options_set_create_if_missing((*C.leveldb_options_t)(p), 1)
	} else {
		C.leveldb_options_set_create_if_missing((*C.leveldb_options_t)(p), 0)
	}
}
func leveldb_options_get_create_if_missing(p *leveldb_options_t) bool {
	v := C.leveldb_options_get_create_if_missing((*C.leveldb_options_t)(p))
	return (v != 0)
}

func leveldb_options_set_error_if_exists(p *leveldb_options_t, v bool) {
	if v {
		C.leveldb_options_set_error_if_exists((*C.leveldb_options_t)(p), 1)
	} else {
		C.leveldb_options_set_error_if_exists((*C.leveldb_options_t)(p), 0)
	}
}
func leveldb_options_get_error_if_exists(p *leveldb_options_t) bool {
	v := C.leveldb_options_get_error_if_exists((*C.leveldb_options_t)(p))
	return (v != 0)
}

func leveldb_options_set_paranoid_checks(p *leveldb_options_t, v bool) {
	if v {
		C.leveldb_options_set_paranoid_checks((*C.leveldb_options_t)(p), 1)
	} else {
		C.leveldb_options_set_paranoid_checks((*C.leveldb_options_t)(p), 0)
	}
}
func leveldb_options_get_paranoid_checks(p *leveldb_options_t) bool {
	v := C.leveldb_options_get_paranoid_checks((*C.leveldb_options_t)(p))
	return (v != 0)
}

func leveldb_options_set_write_buffer_size(p *leveldb_options_t, v int) {
	C.leveldb_options_set_write_buffer_size((*C.leveldb_options_t)(p), C.int32_t(v))
}
func leveldb_options_get_write_buffer_size(p *leveldb_options_t) int {
	v := C.leveldb_options_get_write_buffer_size((*C.leveldb_options_t)(p))
	return int(v)
}

func leveldb_options_set_max_open_files(p *leveldb_options_t, v int) {
	C.leveldb_options_set_max_open_files((*C.leveldb_options_t)(p), C.int32_t(v))
}
func leveldb_options_get_max_open_files(p *leveldb_options_t) int {
	v := C.leveldb_options_get_max_open_files((*C.leveldb_options_t)(p))
	return int(v)
}

func leveldb_options_set_cache(p *leveldb_options_t, v *leveldb_cache_t) {
	C.leveldb_options_set_cache((*C.leveldb_options_t)(p), (*C.leveldb_cache_t)(v))
}
func leveldb_options_get_cache(p *leveldb_options_t) *leveldb_cache_t {
	v := C.leveldb_options_get_cache((*C.leveldb_options_t)(p))
	return (*leveldb_cache_t)(v)
}

func leveldb_options_set_block_size(p *leveldb_options_t, v int) {
	C.leveldb_options_set_block_size((*C.leveldb_options_t)(p), C.int32_t(v))
}
func leveldb_options_get_block_size(p *leveldb_options_t) int {
	v := C.leveldb_options_get_block_size((*C.leveldb_options_t)(p))
	return int(v)
}

func leveldb_options_set_block_restart_interval(p *leveldb_options_t, v int) {
	C.leveldb_options_set_block_restart_interval((*C.leveldb_options_t)(p), C.int32_t(v))
}
func leveldb_options_get_block_restart_interval(p *leveldb_options_t) int {
	v := C.leveldb_options_get_block_restart_interval((*C.leveldb_options_t)(p))
	return int(v)
}

func leveldb_options_set_compression(p *leveldb_options_t, v int) {
	C.leveldb_options_set_compression((*C.leveldb_options_t)(p), C.leveldb_compression_t(v))
}
func leveldb_options_get_compression(p *leveldb_options_t) leveldb_compression_t {
	v := C.leveldb_options_get_compression((*C.leveldb_options_t)(p))
	return leveldb_compression_t(v)
}

// ----------------------------------------------------------------------------
// ReadOptions
// ----------------------------------------------------------------------------

func leveldb_readoptions_create() *leveldb_readoptions_t {
	p := C.leveldb_readoptions_create()
	return (*leveldb_readoptions_t)(p)
}
func leveldb_readoptions_destroy(p *leveldb_readoptions_t) {
	C.leveldb_readoptions_destroy((*C.leveldb_readoptions_t)(p))
}

func leveldb_readoptions_set_verify_checksums(p *leveldb_readoptions_t, v bool) {
	if v {
		C.leveldb_readoptions_set_verify_checksums((*C.leveldb_readoptions_t)(p), 1)
	} else {
		C.leveldb_readoptions_set_verify_checksums((*C.leveldb_readoptions_t)(p), 0)
	}
}
func leveldb_readoptions_get_verify_checksums(p *leveldb_readoptions_t) bool {
	v := C.leveldb_readoptions_get_verify_checksums((*C.leveldb_readoptions_t)(p))
	return (v != 0)
}

func leveldb_readoptions_set_fill_cache(p *leveldb_readoptions_t, v bool) {
	if v {
		C.leveldb_readoptions_set_fill_cache((*C.leveldb_readoptions_t)(p), 1)
	} else {
		C.leveldb_readoptions_set_fill_cache((*C.leveldb_readoptions_t)(p), 0)
	}
}
func leveldb_readoptions_get_fill_cache(p *leveldb_readoptions_t) bool {
	v := C.leveldb_readoptions_get_fill_cache((*C.leveldb_readoptions_t)(p))
	return (v != 0)
}

func leveldb_readoptions_set_snapshot(p *leveldb_readoptions_t, v *leveldb_snapshot_t) {
	C.leveldb_readoptions_set_snapshot((*C.leveldb_readoptions_t)(p), (*C.leveldb_snapshot_t)(v))
}
func leveldb_readoptions_get_snapshot(p *leveldb_readoptions_t) *leveldb_snapshot_t {
	v := C.leveldb_readoptions_get_snapshot((*C.leveldb_readoptions_t)(p))
	return (*leveldb_snapshot_t)(v)
}

// ----------------------------------------------------------------------------
// WriteOptions
// ----------------------------------------------------------------------------

func leveldb_writeoptions_create() *leveldb_writeoptions_t {
	p := C.leveldb_writeoptions_create()
	return (*leveldb_writeoptions_t)(p)
}
func leveldb_writeoptions_destroy(p *leveldb_writeoptions_t) {
	C.leveldb_writeoptions_destroy((*C.leveldb_writeoptions_t)(p))
}

func leveldb_writeoptions_set_sync(p *leveldb_writeoptions_t, v bool) {
	if v {
		C.leveldb_writeoptions_set_sync((*C.leveldb_writeoptions_t)(p), 1)
	} else {
		C.leveldb_writeoptions_set_sync((*C.leveldb_writeoptions_t)(p), 0)
	}
}
func leveldb_writeoptions_get_sync(p *leveldb_writeoptions_t) bool {
	v := C.leveldb_writeoptions_get_sync((*C.leveldb_writeoptions_t)(p))
	return (v != 0)
}

// ----------------------------------------------------------------------------
// DB
// ----------------------------------------------------------------------------

func leveldb_repair_db(name string, opt *leveldb_options_t) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var errStr *C.leveldb_value_t
	status := C.leveldb_repair_db((*C.leveldb_options_t)(opt), cname, &errStr)
	if errStr != nil {
		defer C.leveldb_value_destroy(errStr)
	}
	return leveldb_error(status, errStr)
}
func leveldb_destroy_db(name string, opt *leveldb_options_t) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var errStr *C.leveldb_value_t
	status := C.leveldb_destroy_db((*C.leveldb_options_t)(opt), cname, &errStr)
	if errStr != nil {
		defer C.leveldb_value_destroy(errStr)
	}
	return leveldb_error(status, errStr)
}

func leveldb_open(name string, opt *leveldb_options_t) (*leveldb_t, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var db *C.leveldb_t
	var errStr *C.leveldb_value_t
	status := C.leveldb_open((*C.leveldb_options_t)(opt), cname, &db, &errStr)
	if errStr != nil {
		defer C.leveldb_value_destroy(errStr)
	}
	return (*leveldb_t)(db), leveldb_error(status, errStr)
}
func leveldb_close(db *leveldb_t) {
	C.leveldb_close((*C.leveldb_t)(db))
}

func leveldb_get(db *leveldb_t, key []byte, opt *leveldb_readoptions_t) (*leveldb_value_t, error) {
	var data *C.leveldb_value_t
	var errStr *C.leveldb_value_t
	var status = C.leveldb_get(
		(*C.leveldb_t)(db),
		(*C.leveldb_readoptions_t)(opt),
		(*C.leveldb_slice_t)(leveldb_slice_new(key)), &data,
		&errStr,
	)
	if errStr != nil {
		defer C.leveldb_value_destroy(errStr)
	}
	return (*leveldb_value_t)(data), leveldb_error(status, errStr)
}
func leveldb_put(db *leveldb_t, key, val []byte, opt *leveldb_writeoptions_t) error {
	var errStr *C.leveldb_value_t
	var status = C.leveldb_put(
		(*C.leveldb_t)(db),
		(*C.leveldb_writeoptions_t)(opt),
		(*C.leveldb_slice_t)(leveldb_slice_new(key)),
		(*C.leveldb_slice_t)(leveldb_slice_new(val)),
		&errStr,
	)
	if errStr != nil {
		defer C.leveldb_value_destroy(errStr)
	}
	return leveldb_error(status, errStr)
}
func leveldb_delete(db *leveldb_t, key []byte, opt *leveldb_writeoptions_t) error {
	var errStr *C.leveldb_value_t
	var status = C.leveldb_delete(
		(*C.leveldb_t)(db),
		(*C.leveldb_writeoptions_t)(opt),
		(*C.leveldb_slice_t)(leveldb_slice_new(key)),
		&errStr,
	)
	if errStr != nil {
		defer C.leveldb_value_destroy(errStr)
	}
	return leveldb_error(status, errStr)
}

func leveldb_write(db *leveldb_t, batch *leveldb_writebatch_t, opt *leveldb_writeoptions_t) error {
	var errStr *C.leveldb_value_t
	var status = C.leveldb_write(
		(*C.leveldb_t)(db),
		(*C.leveldb_writeoptions_t)(opt),
		(*C.leveldb_writebatch_t)(batch),
		&errStr,
	)
	if errStr != nil {
		defer C.leveldb_value_destroy(errStr)
	}
	return leveldb_error(status, errStr)
}

func leveldb_property_value(db *leveldb_t, name string) *leveldb_value_t {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var value *C.leveldb_value_t
	C.leveldb_property_value(
		(*C.leveldb_t)(db),
		cname, &value,
	)
	return (*leveldb_value_t)(value)
}

func leveldb_approximate_sizes(db *leveldb_t, rangeStartKey, rangeLimitKey [][]byte) []uint64 {
	if len(rangeStartKey) != len(rangeLimitKey) {
		panic("leveldb/capi.go: invalid argument")
	}

	range_start_key := make([]*C.leveldb_slice_t, len(rangeStartKey))
	range_limit_key := make([]*C.leveldb_slice_t, len(rangeStartKey))
	for i := 0; i < len(rangeStartKey); i++ {
		range_start_key[i] = (*C.leveldb_slice_t)(leveldb_slice_new(rangeStartKey[i]))
		range_limit_key[i] = (*C.leveldb_slice_t)(leveldb_slice_new(rangeLimitKey[i]))
	}

	sizes := make([]uint64, len(rangeStartKey))
	C.leveldb_approximate_sizes(
		(*C.leveldb_t)(db),
		(C.int32_t)(len(rangeStartKey)),
		(**C.leveldb_slice_t)(&range_start_key[0]),
		(**C.leveldb_slice_t)(&range_limit_key[0]),
		(*C.uint64_t)(&sizes[0]),
	)
	return sizes
}

func leveldb_compact_range(db *leveldb_t, rangeStartKey, rangeLimitKey []byte) {
	C.leveldb_compact_range(
		(*C.leveldb_t)(db),
		(*C.leveldb_slice_t)(leveldb_slice_new(rangeStartKey)),
		(*C.leveldb_slice_t)(leveldb_slice_new(rangeLimitKey)),
	)
}

func leveldb_create_snapshot(db *leveldb_t) *leveldb_snapshot_t {
	p := C.leveldb_create_snapshot((*C.leveldb_t)(db))
	return (*leveldb_snapshot_t)(p)
}
func leveldb_release_snapshot(db *leveldb_t, p *leveldb_snapshot_t) {
	C.leveldb_release_snapshot((*C.leveldb_t)(db), (*C.leveldb_snapshot_t)(p))
}

// ----------------------------------------------------------------------------
// Iterator
// ----------------------------------------------------------------------------

func leveldb_create_iterator(db *leveldb_t, opt *leveldb_readoptions_t) *leveldb_iterator_t {
	p := C.leveldb_create_iterator((*C.leveldb_t)(db), (*C.leveldb_readoptions_t)(opt))
	return (*leveldb_iterator_t)(p)
}
func leveldb_iter_destroy(it *leveldb_iterator_t) {
	C.leveldb_iter_destroy((*C.leveldb_iterator_t)(it))
}

func leveldb_iter_valid(it *leveldb_iterator_t) bool {
	v := C.leveldb_iter_valid((*C.leveldb_iterator_t)(it))
	return (v != 0)
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
		(*C.leveldb_slice_t)(leveldb_slice_new(key)),
	)
}

func leveldb_iter_next(it *leveldb_iterator_t) {
	C.leveldb_iter_next((*C.leveldb_iterator_t)(it))
}
func leveldb_iter_prev(it *leveldb_iterator_t) {
	C.leveldb_iter_prev((*C.leveldb_iterator_t)(it))
}

func leveldb_iter_key(it *leveldb_iterator_t) *leveldb_slice_t {
	v := C.leveldb_iter_key((*C.leveldb_iterator_t)(it))
	return (*leveldb_slice_t)(&v)
}
func leveldb_iter_value(it *leveldb_iterator_t) *leveldb_slice_t {
	v := C.leveldb_iter_value((*C.leveldb_iterator_t)(it))
	return (*leveldb_slice_t)(&v)
}

func leveldb_iter_get_error(it *leveldb_iterator_t) error {
	var errStr *C.leveldb_value_t
	var status = C.leveldb_iter_get_error(
		(*C.leveldb_iterator_t)(it),
		&errStr,
	)
	if errStr != nil {
		defer C.leveldb_value_destroy(errStr)
	}
	return leveldb_error(status, errStr)
}

// ----------------------------------------------------------------------------
// WriteBatch
// ----------------------------------------------------------------------------

func leveldb_writebatch_create() *leveldb_writebatch_t {
	p := C.leveldb_writebatch_create()
	return (*leveldb_writebatch_t)(p)
}
func leveldb_writebatch_destroy(p *leveldb_writebatch_t) {
	C.leveldb_writebatch_destroy((*C.leveldb_writebatch_t)(p))
}

func leveldb_writebatch_put(p *leveldb_writebatch_t, key, val []byte) {
	C.leveldb_writebatch_put(
		(*C.leveldb_writebatch_t)(p),
		(*C.leveldb_slice_t)(leveldb_slice_new(key)),
		(*C.leveldb_slice_t)(leveldb_slice_new(val)),
	)
}
func leveldb_writebatch_delete(p *leveldb_writebatch_t, key []byte) {
	C.leveldb_writebatch_delete(
		(*C.leveldb_writebatch_t)(p),
		(*C.leveldb_slice_t)(leveldb_slice_new(key)),
	)
}
func leveldb_writebatch_clear(p *leveldb_writebatch_t) {
	C.leveldb_writebatch_clear((*C.leveldb_writebatch_t)(p))
}

func leveldb_writebatch_iterate(p *leveldb_writebatch_t, it *stateWriteBatchIterater) {
	C.leveldb_writebatch_iterate_helper(
		(*C.leveldb_writebatch_t)(p),
		unsafe.Pointer(it),
	)
}

//export go_leveldb_writebatch_iterater_put
func go_leveldb_writebatch_iterater_put(it unsafe.Pointer, key, val *C.leveldb_slice_t) {
	(*stateWriteBatchIterater)(it).Put(
		leveldb_slice_data((*leveldb_slice_t)(key)),
		leveldb_slice_data((*leveldb_slice_t)(val)),
	)
}

//export go_leveldb_writebatch_iterater_delete
func go_leveldb_writebatch_iterater_delete(it unsafe.Pointer, key *C.leveldb_slice_t) {
	(*stateWriteBatchIterater)(it).Delete(
		leveldb_slice_data((*leveldb_slice_t)(key)),
	)
}

// ----------------------------------------------------------------------------
// Comparator
// ----------------------------------------------------------------------------

func leveldb_comparator_create(cmp *stateComparator) *leveldb_comparator_t {
	p := C.leveldb_comparator_create_helper(unsafe.Pointer(cmp))
	return (*leveldb_comparator_t)(p)
}
func leveldb_comparator_destroy(cmp *leveldb_comparator_t) {
	C.leveldb_comparator_destroy((*C.leveldb_comparator_t)(cmp))
}

func leveldb_comparator_compare(cmp *leveldb_comparator_t, a, b []byte) int {
	v := C.leveldb_comparator_compare(
		(*C.leveldb_comparator_t)(cmp),
		(*C.leveldb_slice_t)(leveldb_slice_new(a)),
		(*C.leveldb_slice_t)(leveldb_slice_new(b)),
	)
	return int(v)
}
func leveldb_comparator_name(cmp *leveldb_comparator_t) string {
	p := C.leveldb_comparator_name(
		(*C.leveldb_comparator_t)(cmp),
	)
	return C.GoString(p)
}

//export go_leveldb_comparator_compare
func go_leveldb_comparator_compare(state unsafe.Pointer, a, b *C.leveldb_slice_t) C.int32_t {
	cmp := (*stateComparator)(state)
	v := cmp.Compare(
		leveldb_slice_data((*leveldb_slice_t)(a)),
		leveldb_slice_data((*leveldb_slice_t)(b)),
	)
	return C.int32_t(v)
}

//export go_leveldb_comparator_name
func go_leveldb_comparator_name(state unsafe.Pointer) *C.char {
	cmp := (*stateComparator)(state)
	if cmp.cname == nil {
		name := cmp.Name()
		cmp.cname = make([]byte, len(name)+1)
		copy(cmp.cname, name)
	}
	return (*C.char)(unsafe.Pointer(&cmp.cname[0]))
}

// ----------------------------------------------------------------------------
// FilterPolicy
// ----------------------------------------------------------------------------

func leveldb_filterpolicy_create(filterPolicy *stateFilter) *leveldb_filterpolicy_t {
	if p, ok := filterPolicy.Filter.(*BloomFilter); ok {
		return p.filter
	} else {
		p := C.leveldb_filterpolicy_create_helper(unsafe.Pointer(filterPolicy))
		return (*leveldb_filterpolicy_t)(p)
	}
}
func leveldb_filterpolicy_create_bloom(bitsPerKey int) *leveldb_filterpolicy_t {
	p := C.leveldb_filterpolicy_create_bloom(C.int32_t(bitsPerKey))
	return (*leveldb_filterpolicy_t)(p)
}
func leveldb_filterpolicy_destroy(p *leveldb_filterpolicy_t) {
	C.leveldb_filterpolicy_destroy((*C.leveldb_filterpolicy_t)(p))
}

func leveldb_filterpolicy_create_filter(
	p *leveldb_filterpolicy_t,
	keys [][]byte,
) *leveldb_value_t {
	keys_ := make([]*leveldb_slice_t, len(keys))
	for i := 0; i < len(keys); i++ {
		keys_[i] = leveldb_slice_new(keys[i])
	}
	v := C.leveldb_filterpolicy_create_filter(
		(*C.leveldb_filterpolicy_t)(p),
		(**C.leveldb_slice_t)(unsafe.Pointer(&keys_[0])),
		C.int32_t(len(keys)),
	)
	return (*leveldb_value_t)(v)
}
func leveldb_filterpolicy_key_may_match(p *leveldb_filterpolicy_t, a, b []byte) bool {
	v := C.leveldb_filterpolicy_key_may_match(
		(*C.leveldb_filterpolicy_t)(p),
		(*C.leveldb_slice_t)(leveldb_slice_new(a)),
		(*C.leveldb_slice_t)(leveldb_slice_new(b)),
	)
	return (v != 0)
}
func leveldb_filterpolicy_name(p *leveldb_filterpolicy_t) string {
	v := C.leveldb_filterpolicy_name((*C.leveldb_filterpolicy_t)(p))
	return C.GoString(v)
}

//export go_leveldb_filterpolicy_create_filter
func go_leveldb_filterpolicy_create_filter(
	state unsafe.Pointer,
	keys **C.leveldb_slice_t,
	num_keys C.int32_t,
) unsafe.Pointer /* (*C.leveldb_value_t) */ {
	var keysArray []*C.leveldb_slice_t
	hKeysArray := (*reflect.SliceHeader)((unsafe.Pointer(&keysArray)))
	hKeysArray.Data = uintptr(unsafe.Pointer(keys))
	hKeysArray.Cap = int(num_keys)
	hKeysArray.Len = int(num_keys)

	keys_ := make([][]byte, int(num_keys))
	for i := 0; i < int(num_keys); i++ {
		keys_[i] = leveldb_slice_data((*leveldb_slice_t)(keysArray[i]))
	}
	filterPolicy := (*stateFilter)(state)
	data := filterPolicy.CreateFilter(keys_)
	return unsafe.Pointer(leveldb_value_create(data))
}

//export go_leveldb_filterpolicy_key_may_match
func go_leveldb_filterpolicy_key_may_match(
	state unsafe.Pointer,
	key *C.leveldb_slice_t,
	filter *C.leveldb_slice_t,
) C.leveldb_bool_t {
	filterPolicy := (*stateFilter)(state)
	rv := filterPolicy.KeyMayMatch(
		leveldb_slice_data((*leveldb_slice_t)(key)),
		leveldb_slice_data((*leveldb_slice_t)(filter)),
	)
	if rv {
		return 1
	} else {
		return 0
	}
}

//export go_leveldb_filterpolicy_name
func go_leveldb_filterpolicy_name(state unsafe.Pointer) *C.char {
	filterPolicy := (*stateFilter)(state)
	if filterPolicy.cname == nil {
		name := filterPolicy.Name()
		filterPolicy.cname = make([]byte, len(name)+1)
		copy(filterPolicy.cname, name)
	}
	return (*C.char)(unsafe.Pointer(&filterPolicy.cname[0]))
}

// ----------------------------------------------------------------------------
// Cache
// ----------------------------------------------------------------------------

func leveldb_cache_create_lru(capacity int64) *leveldb_cache_t {
	p := C.leveldb_cache_create_lru(C.int64_t(capacity))
	return (*leveldb_cache_t)(p)
}
func leveldb_cache_destroy(p *leveldb_cache_t) {
	C.leveldb_cache_destroy((*C.leveldb_cache_t)(p))
}

func leveldb_cache_insert(cache *leveldb_cache_t, key, val []byte) {
	C.leveldb_cache_insert(
		(*C.leveldb_cache_t)(cache),
		(*C.leveldb_slice_t)(leveldb_slice_new(key)),
		(*C.leveldb_slice_t)(leveldb_slice_new(val)),
	)
}
func leveldb_cache_lookup(cache *leveldb_cache_t, key []byte) *leveldb_value_t {
	p := C.leveldb_cache_lookup(
		(*C.leveldb_cache_t)(cache),
		(*C.leveldb_slice_t)(leveldb_slice_new(key)),
	)
	return (*leveldb_value_t)(p)
}
func leveldb_cache_erase(cache *leveldb_cache_t, key []byte) {
	C.leveldb_cache_erase(
		(*C.leveldb_cache_t)(cache),
		(*C.leveldb_slice_t)(leveldb_slice_new(key)),
	)
}

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------
