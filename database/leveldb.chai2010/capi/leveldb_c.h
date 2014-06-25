// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef LEVELDB_C_H_
#define LEVELDB_C_H_

#ifdef __cplusplus
extern "C" {
#endif

#include <stdarg.h>
#include <stddef.h>
#include <stdint.h>

// ----------------------------------------------------------------------------
// Exported types
// ----------------------------------------------------------------------------

struct leveldb_slice_t {
	const char* data;
	int32_t     size;
};

enum leveldb_status_t {
	leveldb_status_ok               = 0,
	leveldb_status_invalid_argument = 1,
	leveldb_status_not_found        = 2,
	leveldb_status_corruption       = 3,
	leveldb_status_io_error         = 4,
	leveldb_status_unknown          = 5,
	leveldb_status_max              = 6,
};

enum leveldb_compression_t {
	leveldb_compression_nil    = 0,
	leveldb_compression_snappy = 1,
	leveldb_compression_max    = 2,
};

// --------------------------------------------------------

typedef unsigned char                  leveldb_bool_t;
typedef struct leveldb_slice_t         leveldb_slice_t;
typedef struct leveldb_value_t         leveldb_value_t;
typedef enum   leveldb_compression_t   leveldb_compression_t;
typedef enum   leveldb_status_t        leveldb_status_t;

typedef struct leveldb_t               leveldb_t;
typedef struct leveldb_options_t       leveldb_options_t;
typedef struct leveldb_readoptions_t   leveldb_readoptions_t;
typedef struct leveldb_writeoptions_t  leveldb_writeoptions_t;

typedef struct leveldb_cache_t         leveldb_cache_t;
typedef struct leveldb_writebatch_t    leveldb_writebatch_t;
typedef struct leveldb_iterator_t      leveldb_iterator_t;
typedef struct leveldb_filterpolicy_t  leveldb_filterpolicy_t;
typedef struct leveldb_comparator_t    leveldb_comparator_t;
typedef struct leveldb_snapshot_t      leveldb_snapshot_t;

// ----------------------------------------------------------------------------
// Version
// ----------------------------------------------------------------------------

int32_t leveldb_major_version();
int32_t leveldb_minor_version();

// ----------------------------------------------------------------------------
// Slice
// ----------------------------------------------------------------------------

leveldb_slice_t leveldb_slice(
	const char* data, int32_t size
);

// ----------------------------------------------------------------------------
// Value
// ----------------------------------------------------------------------------

leveldb_value_t* leveldb_value_create(
	const char* data, int32_t size
);
leveldb_value_t* leveldb_value_create_copy(
	leveldb_value_t* value
);
void leveldb_value_destroy(
	leveldb_value_t* value
);

int32_t leveldb_value_size(
	const leveldb_value_t* value
);
const char* leveldb_value_data(
	const leveldb_value_t* value
);
const char* leveldb_value_cstr(
	const leveldb_value_t* value
);

// ----------------------------------------------------------------------------
// Options
// ----------------------------------------------------------------------------

leveldb_options_t* leveldb_options_create();
void leveldb_options_destroy(leveldb_options_t* opt);

void leveldb_options_set_comparator(
	leveldb_options_t* opt,
	leveldb_comparator_t* cmp
);
leveldb_comparator_t* leveldb_options_get_comparator(
	leveldb_options_t* opt
);

void leveldb_options_set_filter_policy(
	leveldb_options_t* opt,
	leveldb_filterpolicy_t* filter
);
leveldb_filterpolicy_t* leveldb_options_get_filter_policy(
	leveldb_options_t* opt
);

void leveldb_options_set_create_if_missing(
	leveldb_options_t* opt, leveldb_bool_t value
);
leveldb_bool_t leveldb_options_get_create_if_missing(
	leveldb_options_t* opt
);

void leveldb_options_set_error_if_exists(
	leveldb_options_t* opt, leveldb_bool_t value
);
leveldb_bool_t leveldb_options_get_error_if_exists(
	leveldb_options_t* opt
);

void leveldb_options_set_paranoid_checks(
	leveldb_options_t* opt, leveldb_bool_t value
);
leveldb_bool_t leveldb_options_get_paranoid_checks(
	leveldb_options_t* opt 
);

void leveldb_options_set_write_buffer_size(
	leveldb_options_t* opt, int32_t size
);
int32_t leveldb_options_get_write_buffer_size(
	leveldb_options_t* opt
);

void leveldb_options_set_max_open_files(
	leveldb_options_t* opt, int32_t num
);
int32_t leveldb_options_get_max_open_files(
	leveldb_options_t* opt
);

void leveldb_options_set_cache(
	leveldb_options_t* opt, leveldb_cache_t* cache
);
leveldb_cache_t* leveldb_options_get_cache(
	leveldb_options_t* opt
);

void leveldb_options_set_block_size(
	leveldb_options_t* opt, int32_t size
);
int32_t leveldb_options_get_block_size(
	leveldb_options_t* opt
);

void leveldb_options_set_block_restart_interval(
	leveldb_options_t* opt, int32_t size
);
int32_t leveldb_options_get_block_restart_interval(
	leveldb_options_t* opt
);

void leveldb_options_set_compression(
	leveldb_options_t* opt, leveldb_compression_t type
);
leveldb_compression_t leveldb_options_get_compression(
	leveldb_options_t* opt
);

// ----------------------------------------------------------------------------
// ReadOptions
// ----------------------------------------------------------------------------

leveldb_readoptions_t* leveldb_readoptions_create();
void leveldb_readoptions_destroy(
	leveldb_readoptions_t* opt
);

void leveldb_readoptions_set_verify_checksums(
	leveldb_readoptions_t* opt,
	leveldb_bool_t value
);
leveldb_bool_t leveldb_readoptions_get_verify_checksums(
	leveldb_readoptions_t* opt
);

void leveldb_readoptions_set_fill_cache(
	leveldb_readoptions_t* opt,
	leveldb_bool_t value
);
leveldb_bool_t leveldb_readoptions_get_fill_cache(
	leveldb_readoptions_t* opt
);

void leveldb_readoptions_set_snapshot(
	leveldb_readoptions_t* opt,
	const leveldb_snapshot_t* snap
);
leveldb_snapshot_t* leveldb_readoptions_get_snapshot(
	leveldb_readoptions_t* opt
);

// ----------------------------------------------------------------------------
// WriteOptions
// ----------------------------------------------------------------------------

leveldb_writeoptions_t* leveldb_writeoptions_create();
void leveldb_writeoptions_destroy(leveldb_writeoptions_t*);

void leveldb_writeoptions_set_sync(
	leveldb_writeoptions_t* opt,
	leveldb_bool_t value
);
leveldb_bool_t leveldb_writeoptions_get_sync(
	leveldb_writeoptions_t* opt
);

// ----------------------------------------------------------------------------
// DB
// ----------------------------------------------------------------------------

leveldb_status_t leveldb_repair_db(
	const leveldb_options_t* options,
	const char* name,
	leveldb_value_t** errptr
);

leveldb_status_t leveldb_destroy_db(
	const leveldb_options_t* options,
	const char* name,
	leveldb_value_t** errptr
);

leveldb_status_t leveldb_open(
	const leveldb_options_t* options,
	const char* name,
	leveldb_t** db,
	leveldb_value_t** errptr
);

void leveldb_close(leveldb_t* db);

leveldb_status_t leveldb_get(
	leveldb_t* db,
	const leveldb_readoptions_t* options,
	const leveldb_slice_t* key,
	leveldb_value_t** val,
	leveldb_value_t** errptr
);

leveldb_status_t leveldb_put(
	leveldb_t* db,
	const leveldb_writeoptions_t* options,
	const leveldb_slice_t* key,
	const leveldb_slice_t* val,
	leveldb_value_t** errptr
);

leveldb_status_t leveldb_delete(
	leveldb_t* db,
	const leveldb_writeoptions_t* options,
	const leveldb_slice_t* key,
	leveldb_value_t** errptr
);

leveldb_status_t leveldb_write(
	leveldb_t* db,
	const leveldb_writeoptions_t* options,
	leveldb_writebatch_t* batch,
	leveldb_value_t** errptr
);

leveldb_bool_t leveldb_property_value(
	leveldb_t* db,
	const char* propname,
	leveldb_value_t** value
);

void leveldb_approximate_sizes(
	leveldb_t* db,
	int32_t num_ranges,
	const leveldb_slice_t* range_start_key[],
	const leveldb_slice_t* range_limit_key[],
	uint64_t sizes[]
);

void leveldb_compact_range(
	leveldb_t* db,
	const leveldb_slice_t* range_start_key,
	const leveldb_slice_t* range_limit_key
);

const leveldb_snapshot_t* leveldb_create_snapshot(
	leveldb_t* db
);

void leveldb_release_snapshot(
	leveldb_t* db,
	const leveldb_snapshot_t* snapshot
);

// ----------------------------------------------------------------------------
// Iterator
// ----------------------------------------------------------------------------

leveldb_iterator_t* leveldb_create_iterator(
	leveldb_t* db,
	const leveldb_readoptions_t* options
);
void leveldb_iter_destroy(
	leveldb_iterator_t* it
);

leveldb_bool_t leveldb_iter_valid(
	const leveldb_iterator_t* it
);

void leveldb_iter_seek_to_first(
	leveldb_iterator_t* it
);
void leveldb_iter_seek_to_last(
	leveldb_iterator_t* it
);
void leveldb_iter_seek(
	leveldb_iterator_t* it,
	const leveldb_slice_t* key
);

void leveldb_iter_next(leveldb_iterator_t* it);
void leveldb_iter_prev(leveldb_iterator_t* it);

leveldb_slice_t leveldb_iter_key(
	const leveldb_iterator_t* it
);
leveldb_slice_t leveldb_iter_value(
	const leveldb_iterator_t* it
);

leveldb_status_t leveldb_iter_get_error(const leveldb_iterator_t* it,
	leveldb_value_t** errptr
);

// ----------------------------------------------------------------------------
// WriteBatch
// ----------------------------------------------------------------------------

leveldb_writebatch_t* leveldb_writebatch_create();
void leveldb_writebatch_destroy(leveldb_writebatch_t*);

void leveldb_writebatch_put(
	leveldb_writebatch_t* batch,
	const leveldb_slice_t* key,
	const leveldb_slice_t* val
);
void leveldb_writebatch_delete(
	leveldb_writebatch_t* batch,
	const leveldb_slice_t* key
);
void leveldb_writebatch_clear(
	leveldb_writebatch_t* batch
);

void leveldb_writebatch_iterate(
	leveldb_writebatch_t* batch,
	void* state,
	void (*put)(
		void* state,
		const leveldb_slice_t* key,
		const leveldb_slice_t* value
	),
	void (*deleted)(
		void* state,
		const leveldb_slice_t* key
	)
);

// ----------------------------------------------------------------------------
// Comparator
// ----------------------------------------------------------------------------

leveldb_comparator_t* leveldb_comparator_create(
	void* state,
	void (*destructor)(void*),
	int32_t (*compare)(
		void* state,
		const leveldb_slice_t* a,
		const leveldb_slice_t* b
	),
	const char* (*name)(void*)
);
void leveldb_comparator_destroy(leveldb_comparator_t* cmp);

int32_t leveldb_comparator_compare(
	leveldb_comparator_t* cmp,
	const leveldb_slice_t* a,
	const leveldb_slice_t* b
);
const char* leveldb_comparator_name(
	leveldb_comparator_t* cmp
);

// ----------------------------------------------------------------------------
// FilterPolicy
// ----------------------------------------------------------------------------

leveldb_filterpolicy_t* leveldb_filterpolicy_create(
	void* state,
	void (*destructor)(void* state),
	leveldb_value_t* (*create_filter)(
		void* state,
		const leveldb_slice_t** keys,
		int32_t num_keys
	),
	leveldb_bool_t (*key_may_match)(
		void* state,
		const leveldb_slice_t* key,
		const leveldb_slice_t* filter
	),
	const char* (*name)(void* state)
);
leveldb_filterpolicy_t* leveldb_filterpolicy_create_bloom(
	int32_t bits_per_key
);
void leveldb_filterpolicy_destroy(
	leveldb_filterpolicy_t* filterpolicy
);

leveldb_value_t* leveldb_filterpolicy_create_filter(
	leveldb_filterpolicy_t* filterpolicy,
	const leveldb_slice_t** keys,
	int32_t num_keys
);
leveldb_bool_t leveldb_filterpolicy_key_may_match(
	leveldb_filterpolicy_t* filterpolicy,
	const leveldb_slice_t* key,
	const leveldb_slice_t* filter
);
const char* leveldb_filterpolicy_name(
	leveldb_filterpolicy_t* filterpolicy
);

// ----------------------------------------------------------------------------
// Cache
// ----------------------------------------------------------------------------

leveldb_cache_t* leveldb_cache_create_lru(int64_t capacity);
void leveldb_cache_destroy(leveldb_cache_t* cache);

void leveldb_cache_insert(
	leveldb_cache_t* cache,
	const leveldb_slice_t* key,
	const leveldb_slice_t* val
);
leveldb_value_t* leveldb_cache_lookup(
	leveldb_cache_t* cache,
	const leveldb_slice_t* key
);
void leveldb_cache_erase(
	leveldb_cache_t* cache,
	const leveldb_slice_t* key
);

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------

#ifdef __cplusplus
} // extern "C"
#endif
#endif // LEVELDB_C_H_
