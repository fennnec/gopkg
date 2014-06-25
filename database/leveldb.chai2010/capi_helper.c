// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "capi_helper.h"

#ifdef __cplusplus
extern "C" {
#	include "_cgo_export.h"
}
#else
#	include "_cgo_export.h"
#endif

// ----------------------------------------------------------------------------
// WriteBatch
// ----------------------------------------------------------------------------

static void _leveldb_writebatch_iterate_put(
	void* state,
	const leveldb_slice_t* key,
	const leveldb_slice_t* val
) {
	go_leveldb_writebatch_iterater_put(state,
		(leveldb_slice_t*)(key),
		(leveldb_slice_t*)(val)
	);
}

static void _leveldb_writebatch_iterate_delete(
	void* state,
	const leveldb_slice_t* key
) {
	go_leveldb_writebatch_iterater_delete(state,
		(leveldb_slice_t*)(key)
	);
}

void leveldb_writebatch_iterate_helper(leveldb_writebatch_t* batch, void* state) {
	leveldb_writebatch_iterate(batch, state,
		_leveldb_writebatch_iterate_put,
		_leveldb_writebatch_iterate_delete
	);
}

// ----------------------------------------------------------------------------
// Comparator
// ----------------------------------------------------------------------------

static void _leveldb_comparator_destructor(void* state) {
	// Empty
}

static int32_t _leveldb_comparator_compare(void* state,
	const leveldb_slice_t* a,
	const leveldb_slice_t* b
) {
	return go_leveldb_comparator_compare(state,
		(leveldb_slice_t*)(a),
		(leveldb_slice_t*)(b)
	);
}

static const char* _leveldb_comparator_name(void* state) {
	return go_leveldb_comparator_name(state);
}

leveldb_comparator_t* leveldb_comparator_create_helper(void* state) {
	return leveldb_comparator_create(
		state,
		_leveldb_comparator_destructor,
		_leveldb_comparator_compare,
		_leveldb_comparator_name
	);
}

// ----------------------------------------------------------------------------
// FilterPolicy
// ----------------------------------------------------------------------------

static void _leveldb_filterpolicy_destructor(void* state) {
	// Empty
}

static leveldb_value_t* _leveldb_filterpolicy_create_filter(void* state,
	const leveldb_slice_t** keys,
	int32_t num_keys
) {
	return (leveldb_value_t*)go_leveldb_filterpolicy_create_filter(state,
		(leveldb_slice_t**)keys, num_keys
	);
}

static leveldb_bool_t _leveldb_filterpolicy_key_may_match(void* state,
	const leveldb_slice_t* key,
	const leveldb_slice_t* filter
) {
	return go_leveldb_filterpolicy_key_may_match(state,
		(leveldb_slice_t*)(key),
		(leveldb_slice_t*)(filter)
	);
}

static const char* _leveldb_filterpolicy_name(void* state) {
	return go_leveldb_filterpolicy_name(state);
}

leveldb_filterpolicy_t* leveldb_filterpolicy_create_helper(void* state) {
	return leveldb_filterpolicy_create(
		state,
		_leveldb_filterpolicy_destructor,
		_leveldb_filterpolicy_create_filter,
		_leveldb_filterpolicy_key_may_match,
		_leveldb_filterpolicy_name
	);
}

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------

