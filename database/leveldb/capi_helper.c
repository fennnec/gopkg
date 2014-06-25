// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "capi_helper.h"
#include "_cgo_export.h"

/* Write batch */

static void leveldb_writebatch_iterate_put(void* p,
	const char* k, size_t klen,
	const char* v, size_t vlen
) {
	go_leveldb_writebatch_iterate_put(p, (char*)k, klen, (char*)v, vlen);
}

static void leveldb_writebatch_iterate_del(void* p,
	const char* k, size_t klen
) {
	go_leveldb_writebatch_iterate_del(p, (char*)k, klen);
}

void go_leveldb_writebatch_iterate_helper(leveldb_writebatch_t* batch, void* state) {
	leveldb_writebatch_iterate(batch, state,
		leveldb_writebatch_iterate_put,
		leveldb_writebatch_iterate_del
	);
}

/* Comparator */

static void leveldb_comparator_create_destructor(void* state) {
	go_leveldb_comparator_create_state_destructor(state);
}

static int leveldb_comparator_create_compare(void* state,
	const char* a, size_t alen,
	const char* b, size_t blen
) {
	return leveldb_comparator_create_compare(state, (char*)a, alen, (char*)b, blen);
}

static const char* leveldb_comparator_create_name(void* state) {
	return go_leveldb_comparator_create_state_name(state);
}

leveldb_comparator_t* leveldb_comparator_create_helper(void* state) {
	return leveldb_comparator_create(
		state,
		leveldb_comparator_create_destructor,
		leveldb_comparator_create_compare,
		leveldb_comparator_create_name
	);
}
