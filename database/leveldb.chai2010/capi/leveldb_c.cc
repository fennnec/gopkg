// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "leveldb_c.h"

#include "leveldb/cache.h"
#include "leveldb/comparator.h"
#include "leveldb/db.h"
#include "leveldb/filter_policy.h"
#include "leveldb/iterator.h"
#include "leveldb/options.h"
#include "leveldb/status.h"
#include "leveldb/write_batch.h"

#include <stdio.h>
#include <stdlib.h>
#include <vector>

// ----------------------------------------------------------------------------
// Utils
// ----------------------------------------------------------------------------

#define DISABLE_NEW_AND_DELETE(TypeName) private: \
	TypeName();                                   \
	TypeName(const TypeName&);                    \
	TypeName& operator=(const TypeName&);         \
	~TypeName()

inline leveldb_bool_t fromStdBool(bool v) {
	return (v)? 1: 0;
}

inline bool toStdBool(leveldb_bool_t v) {
	return (v != 0)? true: false;
}

inline leveldb_slice_t fromLdbSlice(const leveldb::Slice& a) {
	return leveldb_slice(a.data(), int32_t(a.size()));
}

inline leveldb::Slice toLdbSlice(const leveldb_slice_t* a) {
	if(a != NULL && a->data != NULL && a->size > 0) {
		return leveldb::Slice(a->data, size_t(a->size));
	} else {
		return leveldb::Slice(NULL, 0);
	}
}

inline leveldb_status_t fromLdbStatus(const leveldb::Status& s, leveldb_value_t** err) {
	if(s.ok()) {
		return leveldb_status_ok;
	}
	if(err != NULL && *err != NULL) {
		if(*err != NULL) { leveldb_value_destroy(*err); }
		std::string tmp = s.ToString();
		*err = leveldb_value_create(tmp.data(), int32_t(tmp.size()));
	}
	if(s.IsNotFound()) {
		return leveldb_status_not_found;
	}
	if(s.IsCorruption()) {
		return leveldb_status_corruption;
	}
	if(s.IsIOError()) {
		return leveldb_status_io_error;
	}
	return leveldb_status_unknown;
}

// ----------------------------------------------------------------------------
// Version
// ----------------------------------------------------------------------------

int32_t leveldb_major_version() {
	return int32_t(leveldb::kMajorVersion);
}

int32_t leveldb_minor_version() {
	return int32_t(leveldb::kMinorVersion);
}

// ----------------------------------------------------------------------------
// Slice
// ----------------------------------------------------------------------------

leveldb_slice_t leveldb_slice(const char* data, int32_t size) {
	if(data != NULL && size > 0) {
		leveldb_slice_t a = { data, size };
		return a;
	} else {
		leveldb_slice_t a = { NULL, 0 };
		return a;
	}
}

// ----------------------------------------------------------------------------
// Value
// ----------------------------------------------------------------------------

struct leveldb_value_t {
	std::string value_;
	leveldb_value_t(const char* data, int32_t size): value_(data, size) {}
	leveldb_value_t(const std::string& data): value_(data) {}
};

leveldb_value_t* leveldb_value_create(const char*data, int32_t size) {
	if(data != NULL && size > 0) {
		leveldb_value_t* slice = new leveldb_value_t(data, size);
		return slice;
	} else {
		leveldb_value_t* slice = new leveldb_value_t(NULL, 0);
		return slice;
	}
}
leveldb_value_t* leveldb_value_create_copy(
	leveldb_value_t* value
) {
	if(value != NULL) {
		return leveldb_value_create(
			value->value_.data(),
			int32_t(value->value_.size())
		);
	} else {
		return leveldb_value_create(NULL, 0);
	}
}
void leveldb_value_destroy(leveldb_value_t* value) {
	if(value != NULL) {
		delete value;
	}
}

int32_t leveldb_value_size(const leveldb_value_t* value) {
	if(value != NULL && !value->value_.empty()) {
		return int32_t(value->value_.size());
	} else {
		return 0;
	}
}

const char* leveldb_value_data(const leveldb_value_t* value) {
	if(value != NULL && !value->value_.empty()) {
		return value->value_.data();
	} else {
		return NULL;
	}
}

const char* leveldb_value_cstr(const leveldb_value_t* value) {
	if(value != NULL && !value->value_.empty()) {
		return value->value_.data();
	} else {
		return NULL;
	}
}

// ----------------------------------------------------------------------------
// Options
// ----------------------------------------------------------------------------

struct leveldb_options_t: leveldb::Options {
	DISABLE_NEW_AND_DELETE(leveldb_options_t);
};

leveldb_options_t* leveldb_options_create() {
	return (leveldb_options_t*)(new leveldb::Options());
}

void leveldb_options_destroy(leveldb_options_t* opt) {
	if(opt != NULL) {
		delete (leveldb::Options*)opt;
	}
}

void leveldb_options_set_comparator(
	leveldb_options_t* opt,
	leveldb_comparator_t* cmp
) {
	if(opt != NULL) {
		opt->comparator = (leveldb::Comparator*)(cmp);
	}
}

leveldb_comparator_t* leveldb_options_get_comparator(
	leveldb_options_t* opt
) {
	if(opt != NULL) {
		return (leveldb_comparator_t*)(opt->comparator);
	} else {
		return NULL;
	}
}

void leveldb_options_set_filter_policy(
	leveldb_options_t* opt,
	leveldb_filterpolicy_t* policy
) {
	if(opt != NULL) {
		opt->filter_policy = (leveldb::FilterPolicy*)(policy);
	}
}
leveldb_filterpolicy_t* leveldb_options_get_filter_policy(
	leveldb_options_t* opt
) {
	if(opt != NULL) {
		return (leveldb_filterpolicy_t*)(opt->filter_policy);
	} else {
		return NULL;
	}
}

void leveldb_options_set_create_if_missing(
	leveldb_options_t* opt, leveldb_bool_t v
) {
	if(opt != NULL) {
		opt->create_if_missing = toStdBool(v);
	}
}
leveldb_bool_t leveldb_options_get_create_if_missing(
	leveldb_options_t* opt
) {
	if(opt != NULL) {
		return fromStdBool(opt->create_if_missing);
	} else {
		return fromStdBool(false);
	}
}

void leveldb_options_set_error_if_exists(
	leveldb_options_t* opt, leveldb_bool_t v
) {
	if(opt != NULL) {
		opt->error_if_exists = toStdBool(v);
	}
}
leveldb_bool_t leveldb_options_get_error_if_exists(
	leveldb_options_t* opt
) {
	if(opt != NULL) {
		return fromStdBool(opt->error_if_exists);
	} else {
		return fromStdBool(false);
	}
}

void leveldb_options_set_paranoid_checks(
	leveldb_options_t* opt, leveldb_bool_t v
) {
	if(opt != NULL) {
		opt->paranoid_checks = toStdBool(v);
	}
}
leveldb_bool_t leveldb_options_get_paranoid_checks(
	leveldb_options_t* opt 
) {
	if(opt != NULL) {
		return fromStdBool(opt->paranoid_checks);
	} else {
		return fromStdBool(false);
	}
}

void leveldb_options_set_write_buffer_size(
	leveldb_options_t* opt, int32_t size
) {
	if(opt != NULL) {
		opt->write_buffer_size = size_t(size);
	}
}
int32_t leveldb_options_get_write_buffer_size(
	leveldb_options_t* opt
) {
	if(opt != NULL) {
		return int32_t(opt->write_buffer_size);
	} else {
		return int32_t(0);
	}
}

void leveldb_options_set_max_open_files(leveldb_options_t* opt, int32_t n) {
	if(opt != NULL) {
		opt->max_open_files = n;
	}
}
int32_t leveldb_options_get_max_open_files(
	leveldb_options_t* opt
) {
	if(opt != NULL) {
		return int32_t(opt->max_open_files);
	} else {
		return int32_t(0);
	}
}

void leveldb_options_set_cache(leveldb_options_t* opt, leveldb_cache_t* c) {
	if(opt != NULL) {
		opt->block_cache = (leveldb::Cache*)(c);
	}
}
leveldb_cache_t* leveldb_options_get_cache(
	leveldb_options_t* opt
) {
	if(opt != NULL) {
		return (leveldb_cache_t*)(opt->block_cache);
	} else {
		return NULL;
	}
}

void leveldb_options_set_block_size(leveldb_options_t* opt, int32_t size) {
	if(opt != NULL) {
		opt->block_size = size_t(size);
	}
}
int32_t leveldb_options_get_block_size(
	leveldb_options_t* opt
) {
	if(opt != NULL) {
		return int32_t(opt->block_size);
	} else {
		return int32_t(0);
	}
}

void leveldb_options_set_block_restart_interval(leveldb_options_t* opt, int32_t n) {
	if(opt != NULL) {
		opt->block_restart_interval = int32_t(n);
	}
}
int32_t leveldb_options_get_block_restart_interval(
	leveldb_options_t* opt
) {
	if(opt != NULL) {
		return int32_t(opt->block_restart_interval);
	} else {
		return int32_t(0);
	}
}

void leveldb_options_set_compression(
	leveldb_options_t* opt,
	leveldb_compression_t t
) {
	if(opt != NULL) {
		opt->compression = static_cast<leveldb::CompressionType>(t);
	}
}
leveldb_compression_t leveldb_options_get_compression(
	leveldb_options_t* opt
) {
	if(opt != NULL) {
		return leveldb_compression_t(opt->compression);
	} else {
		return leveldb_compression_t(leveldb_compression_nil);
	}
}

// ----------------------------------------------------------------------------
// ReadOptions
// ----------------------------------------------------------------------------

struct leveldb_readoptions_t: leveldb::ReadOptions {
	DISABLE_NEW_AND_DELETE(leveldb_readoptions_t);
};

leveldb_readoptions_t* leveldb_readoptions_create() {
	return (leveldb_readoptions_t*)(new leveldb::ReadOptions());
}

void leveldb_readoptions_destroy(leveldb_readoptions_t* opt) {
	if(opt != NULL) {
		delete (leveldb::ReadOptions*)opt;
	}
}

void leveldb_readoptions_set_verify_checksums(
	leveldb_readoptions_t* opt,
	leveldb_bool_t v
) {
	if(opt != NULL) {
		opt->verify_checksums = toStdBool(v);
	}
}

leveldb_bool_t leveldb_readoptions_get_verify_checksums(
	leveldb_readoptions_t* opt
) {
	if(opt != NULL) {
		return fromStdBool(opt->verify_checksums);
	} else {
		return fromStdBool(false);
	}
}

void leveldb_readoptions_set_fill_cache(
	leveldb_readoptions_t* opt, leveldb_bool_t v
) {
	if(opt != NULL) {
		opt->fill_cache = toStdBool(v);
	}
}

leveldb_bool_t leveldb_readoptions_get_fill_cache(
	leveldb_readoptions_t* opt
) {
	if(opt != NULL) {
		return fromStdBool(opt->fill_cache);
	} else {
		return fromStdBool(false);
	}
}

void leveldb_readoptions_set_snapshot(
	leveldb_readoptions_t* opt,
	const leveldb_snapshot_t* snap
) {
	if(opt != NULL) {
		opt->snapshot = (leveldb::Snapshot*)(snap);
	}
}

leveldb_snapshot_t* leveldb_readoptions_get_snapshot(
	leveldb_readoptions_t* opt
) {
	if(opt != NULL) {
		return (leveldb_snapshot_t*)(opt->snapshot);
	} else {
		return NULL;
	}
}

// ----------------------------------------------------------------------------
// WriteOptions
// ----------------------------------------------------------------------------

struct leveldb_writeoptions_t: leveldb::WriteOptions {
	DISABLE_NEW_AND_DELETE(leveldb_writeoptions_t);
};

leveldb_writeoptions_t* leveldb_writeoptions_create() {
	return (leveldb_writeoptions_t*)(new leveldb::WriteOptions());
}

void leveldb_writeoptions_destroy(leveldb_writeoptions_t* opt) {
	if(opt != NULL) {
		delete (leveldb::WriteOptions*)opt;
	}
}

void leveldb_writeoptions_set_sync(
	leveldb_writeoptions_t* opt, leveldb_bool_t v
) {
	if(opt != NULL) {
		opt->sync = toStdBool(v);
	}
}

leveldb_bool_t leveldb_writeoptions_get_sync(leveldb_writeoptions_t* opt) {
	if(opt != NULL) {
		return fromStdBool(opt->sync);
	} else {
		return fromStdBool(false);
	}
}

// ----------------------------------------------------------------------------
// DB
// ----------------------------------------------------------------------------

struct leveldb_t: leveldb::DB {
	DISABLE_NEW_AND_DELETE(leveldb_t);
};
struct leveldb_snapshot_t: leveldb::Snapshot {
	DISABLE_NEW_AND_DELETE(leveldb_snapshot_t);
};

leveldb_status_t leveldb_repair_db(
	const leveldb_options_t* opt,
	const char* name,
	leveldb_value_t** errptr
) {
	if(name == NULL || name[0] == '\0') {
		return leveldb_status_invalid_argument;
	}
	leveldb::Status s = RepairDB(name,
		(opt!=NULL)? *(leveldb::Options*)(opt): leveldb::Options()
	);
	return fromLdbStatus(s, errptr);
}
leveldb_status_t leveldb_destroy_db(
	const leveldb_options_t* opt,
	const char* name,
	leveldb_value_t** errptr
) {
	if(name == NULL || name[0] == '\0') {
		return leveldb_status_invalid_argument;
	}
	leveldb::Status s = DestroyDB(name,
		(opt!=NULL)? *(leveldb::Options*)(opt): leveldb::Options()
	);
	return fromLdbStatus(s, errptr);
}

leveldb_status_t leveldb_open(
	const leveldb_options_t* opt,
	const char* name,
	leveldb_t** db,
	leveldb_value_t** errptr
) {
	if(name == NULL || name[0] == '\0' || db == NULL || *db == NULL) {
		return leveldb_status_invalid_argument;
	}
	leveldb::Status s = leveldb::DB::Open(
		(opt!=NULL)? *(leveldb::Options*)(opt): leveldb::Options(),
		std::string(name),
		(leveldb::DB**)db
	);
	return fromLdbStatus(s, errptr);
}

void leveldb_close(leveldb_t* db) {
	if(db != NULL) {
		delete (leveldb::DB*)db;
	}
}

leveldb_status_t leveldb_get(
	leveldb_t* db,
	const leveldb_readoptions_t* opt,
	const leveldb_slice_t* key,
	leveldb_value_t** value,
	leveldb_value_t** errptr
) {
	if(db == NULL|| key == NULL|| value == NULL|| *value == NULL) {
		return leveldb_status_invalid_argument;
	}
	std::string tmp;
	leveldb::Status s = db->Get(
		(opt!=NULL)? *(leveldb::ReadOptions*)(opt): leveldb::ReadOptions(),
		toLdbSlice(key),
		&tmp
	);
	*value = leveldb_value_create(tmp.data(), int32_t(tmp.size()));
	return fromLdbStatus(s, errptr);
}

leveldb_status_t leveldb_put(
	leveldb_t* db,
	const leveldb_writeoptions_t* opt,
	const leveldb_slice_t* key,
	const leveldb_slice_t* val,
	leveldb_value_t** errptr
) {
	if(db == NULL || key == NULL) {
		return leveldb_status_invalid_argument;
	}
	leveldb::Status s = db->Put(
		(opt!=NULL)? *(leveldb::WriteOptions*)(opt): leveldb::WriteOptions(),
		toLdbSlice(key),
		toLdbSlice(val)
	);
	return fromLdbStatus(s, errptr);
}

leveldb_status_t leveldb_delete(
	leveldb_t* db,
	const leveldb_writeoptions_t* opt,
	const leveldb_slice_t* key,
	leveldb_value_t** errptr
) {
	if(db == NULL || key == NULL) {
		return leveldb_status_invalid_argument;
	}
	leveldb::Status s = db->Delete(
		(opt!=NULL)? *(leveldb::WriteOptions*)(opt): leveldb::WriteOptions(),
		toLdbSlice(key)
	);
	return fromLdbStatus(s, errptr);
}

leveldb_status_t leveldb_write(
	leveldb_t* db,
	const leveldb_writeoptions_t* opt,
	leveldb_writebatch_t* batch,
	leveldb_value_t** errptr
) {
	if(db == NULL || batch == NULL) {
		return leveldb_status_invalid_argument;
	}
	leveldb::Status s = db->Write(
		(opt!=NULL)? *(leveldb::WriteOptions*)(opt): leveldb::WriteOptions(),
		(leveldb::WriteBatch*)batch
	);
	return fromLdbStatus(s, errptr);
}

leveldb_bool_t leveldb_property_value(
	leveldb_t* db,
	const char* propname,
	leveldb_value_t** value
) {
	if(db == NULL) {
		return fromStdBool(false);
	}
	if(propname == NULL || propname[0] == '\0') {
		return fromStdBool(false);
	}
	if(value == NULL || *value == NULL) {
		return fromStdBool(false);
	}
	std::string tmp;
	if (db->GetProperty(leveldb::Slice(propname), &tmp)) {
		*value = leveldb_value_create(tmp.data(), int32_t(tmp.size()));
		return fromStdBool(true);
	} else {
		return fromStdBool(false);
	}
}

void leveldb_approximate_sizes(
	leveldb_t* db,
	int32_t num_ranges,
	const leveldb_slice_t* range_start_key[],
	const leveldb_slice_t* range_limit_key[],
	uint64_t sizes[]
) {
	if(db == NULL || num_ranges <= 0) {
		return;
	}
	if(range_start_key == NULL || range_limit_key == NULL || sizes == NULL) {
		return;
	}
	leveldb::Range* ranges = new leveldb::Range[num_ranges];
	for (int32_t i = 0; i < num_ranges; i++) {
		ranges[i].start = toLdbSlice(range_start_key[i]);
		ranges[i].limit = toLdbSlice(range_limit_key[i]);
	}
	db->GetApproximateSizes(ranges, num_ranges, sizes);
	delete[] ranges;
}

void leveldb_compact_range(
	leveldb_t* db,
	const leveldb_slice_t* range_start_key,
	const leveldb_slice_t* range_limit_key
) {
	if(db == NULL) return;
	leveldb::Slice a = toLdbSlice(range_start_key);
	leveldb::Slice b = toLdbSlice(range_limit_key);
	db->CompactRange(&a, &b);
}

const leveldb_snapshot_t* leveldb_create_snapshot(
	leveldb_t* db
) {
	if(db != NULL) {
		return (leveldb_snapshot_t*)(db->GetSnapshot());
	} else {
		return NULL;
	}
}

void leveldb_release_snapshot(
	leveldb_t* db,
	const leveldb_snapshot_t* snapshot
) {
	if(db != NULL && snapshot != NULL) {
		db->ReleaseSnapshot((leveldb::Snapshot*)(snapshot));
	}
}

// ----------------------------------------------------------------------------
// Iterator
// ----------------------------------------------------------------------------

struct leveldb_iterator_t: leveldb::Iterator {
	DISABLE_NEW_AND_DELETE(leveldb_iterator_t);
};

leveldb_iterator_t* leveldb_create_iterator(
	leveldb_t* db,
	const leveldb_readoptions_t* opt
) {
	if(db != NULL) {
		return (leveldb_iterator_t*)(db->NewIterator(
			(opt!=NULL)? *(leveldb::ReadOptions*)(opt): leveldb::ReadOptions()
		));
	} else {
		return NULL;
	}
}

void leveldb_iter_destroy(
	leveldb_iterator_t* it
) {
	if(it != NULL) {
		delete (leveldb::Iterator*)it;
	}
}

leveldb_bool_t leveldb_iter_valid(const leveldb_iterator_t* it) {
	if(it != NULL) {
		return fromStdBool(it->Valid());
	} else {
		return fromStdBool(false);
	}
}

void leveldb_iter_seek_to_first(
	leveldb_iterator_t* it
) {
	if(it != NULL) {
		it->SeekToFirst();
	}
}

void leveldb_iter_seek_to_last(
	leveldb_iterator_t* it
) {
	it->SeekToLast();
}

void leveldb_iter_seek(
	leveldb_iterator_t* it, 
	const leveldb_slice_t* key
) {
	if(it != NULL) {
		it->Seek(toLdbSlice(key));
	}
}

void leveldb_iter_next(leveldb_iterator_t* it) {
	if(it != NULL) it->Next();
}

void leveldb_iter_prev(leveldb_iterator_t* it) {
	if(it != NULL) it->Prev();
}

leveldb_slice_t leveldb_iter_key(const leveldb_iterator_t* it) {
	if(it != NULL) {
		leveldb::Slice s = it->key();
		return leveldb_slice(s.data(), int32_t(s.size()));
	} else {
		return leveldb_slice(NULL, int32_t(0));
	}
}

leveldb_slice_t leveldb_iter_value(const leveldb_iterator_t* it) {
	if(it != NULL) {
		leveldb::Slice s = it->value();
		return leveldb_slice(s.data(), int32_t(s.size()));
	} else {
		return leveldb_slice(NULL, int32_t(0));
	}
}

leveldb_status_t leveldb_iter_get_error(const leveldb_iterator_t* it,
	leveldb_value_t** errptr
) {
	if(it == NULL) {
		return leveldb_status_invalid_argument;
	}
	leveldb::Status s = it->status();
	return fromLdbStatus(s, errptr);
}

// ----------------------------------------------------------------------------
// WriteBatch
// ----------------------------------------------------------------------------

struct leveldb_writebatch_t: leveldb::WriteBatch {
	DISABLE_NEW_AND_DELETE(leveldb_writebatch_t);
};

leveldb_writebatch_t* leveldb_writebatch_create() {
	return (leveldb_writebatch_t*)(new leveldb::WriteBatch());
}

void leveldb_writebatch_destroy(leveldb_writebatch_t* batch) {
	if(batch != NULL) {
		delete (leveldb::WriteBatch*)batch;
	}
}

void leveldb_writebatch_put(
	leveldb_writebatch_t* batch,
	const leveldb_slice_t* key,
	const leveldb_slice_t* val
) {
	if(batch != NULL && key != NULL) {
		batch->Put(toLdbSlice(key), toLdbSlice(val));
	}
}
void leveldb_writebatch_delete(
	leveldb_writebatch_t* batch,
	const leveldb_slice_t* key
) {
	if(batch != NULL && key != NULL) {
		batch->Delete(toLdbSlice(key));
	}
}
void leveldb_writebatch_clear(leveldb_writebatch_t* batch) {
	if(batch != NULL) {
		batch->Clear();
	}
}

void leveldb_writebatch_iterate(
	leveldb_writebatch_t* batch,
	void* state,
	void (*put)(
		void* state,
		const leveldb_slice_t* k,
		const leveldb_slice_t* v
	),
	void (*deleted)(
		void* state,
		const leveldb_slice_t* k
	)
) {
	if(batch == NULL) return;
	struct H: public leveldb::WriteBatch::Handler {
		void* state_;
		void (*put_)(
			void*, const leveldb_slice_t* k, const leveldb_slice_t* v
		);
		void (*deleted_)(void*, const leveldb_slice_t* k);
    
		virtual void Put(
			const leveldb::Slice& key, const leveldb::Slice& val
		) {
			if(put_ != NULL) {
				leveldb_slice_t k = fromLdbSlice(key);
				leveldb_slice_t v = fromLdbSlice(val);
				(*put_)(state_, &k, &v);
			}
		}
		virtual void Delete(const leveldb::Slice& key) {
			if(deleted_ != NULL) {
				leveldb_slice_t k = fromLdbSlice(key);
				(*deleted_)(state_, &k);
			}
		}
	};
	H handler;
	handler.state_ = state;
	handler.put_ = put;
	handler.deleted_ = deleted;
	batch->Iterate(&handler);
}

// ----------------------------------------------------------------------------
// Comparator
// ----------------------------------------------------------------------------

struct leveldb_comparator_t : public leveldb::Comparator {
	void* state_;
	void (*destructor_)(void*);
	int32_t (*compare_)(
		void* state,
		const leveldb_slice_t* a,
		const leveldb_slice_t* b
	);
	const char* (*name_)(void*);

	virtual ~leveldb_comparator_t() {
		(*destructor_)(state_);
	}

	virtual int32_t Compare(
		const leveldb::Slice& a, const leveldb::Slice& b
	) const {
		leveldb_slice_t a_ = fromLdbSlice(a);
		leveldb_slice_t b_ = fromLdbSlice(b);
		return (*compare_)(state_, &a_, &b_);
	}

	virtual const char* Name() const {
		return (*name_)(state_);
	}

	// No-ops since the C binding does not support key shortening methods.
	virtual void FindShortestSeparator(
		std::string*, const leveldb::Slice&
	) const {
		// Empty
	}
	virtual void FindShortSuccessor(std::string* key) const {
		// Empty
	}
};

leveldb_comparator_t* leveldb_comparator_create(
	void* state,
	void (*destructor)(void* state),
	int32_t (*compare)(
		void* state,
		const leveldb_slice_t* a,
		const leveldb_slice_t* b
	),
	const char* (*name)(void*)
) {
	if(destructor == NULL || compare == NULL || name == NULL) {
		return NULL;
	}
	leveldb_comparator_t* result = new leveldb_comparator_t;
	result->state_ = state;
	result->destructor_ = destructor;
	result->compare_ = compare;
	result->name_ = name;
	return result;
}

void leveldb_comparator_destroy(leveldb_comparator_t* cmp) {
	if(cmp != NULL) {
		delete cmp;
	}
}

int32_t leveldb_comparator_compare(
	leveldb_comparator_t* cmp,
	const leveldb_slice_t* a,
	const leveldb_slice_t* b
) {
	if(cmp != NULL) {
		return (*cmp->compare_)(cmp->state_, a, b);
	} else {
		return int32_t(0);
	}
}

const char* leveldb_comparator_name(
	leveldb_comparator_t* cmp
) {
	if(cmp != NULL) {
		return (*cmp->name_)(cmp->state_);
	} else {
		return NULL;
	}
}

// ----------------------------------------------------------------------------
// FilterPolicy
// ----------------------------------------------------------------------------

struct leveldb_filterpolicy_t : public leveldb::FilterPolicy {
	void* state_;
	void (*destructor_)(void* state);
	const char* (*name_)(void* state);
	leveldb_value_t* (*create_)(
		void* state,
		const leveldb_slice_t** keys,
		int32_t num_keys
	);
	leveldb_bool_t (*key_match_)(
		void* state,
		const leveldb_slice_t* key,
		const leveldb_slice_t* filter
	);

	virtual ~leveldb_filterpolicy_t() {
		(*destructor_)(state_);
	}

	virtual void CreateFilter(
		const leveldb::Slice* keys, int32_t n, std::string* dst
	) const {
		std::vector<leveldb_slice_t*> keys_(n);
		std::vector<leveldb_slice_t> keys_tmp_(n);
		for(int32_t i = 0; i < n; ++i) {
			keys_tmp_[i] = fromLdbSlice(keys[i]);
			keys_[i] = &keys_tmp_[i];
		}
		leveldb_value_t* filter = (*create_)(
			state_, (const leveldb_slice_t**)(&keys_[0]), int32_t(n)
		);
		dst->append(leveldb_value_data(filter), leveldb_value_size(filter));
		leveldb_value_destroy(filter);
	}

	virtual bool KeyMayMatch(
		const leveldb::Slice& key, const leveldb::Slice& filter
	) const {
		leveldb_slice_t a = fromLdbSlice(key);
		leveldb_slice_t b = fromLdbSlice(filter);
		leveldb_bool_t rv = (*key_match_)(state_, &a, &b);
		return toStdBool(rv);
	}

	virtual const char* Name() const {
		return (*name_)(state_);
	}
};

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
) {
	if(destructor == NULL || create_filter == NULL || key_may_match == NULL || name == NULL) {
		return NULL;
	}
	leveldb_filterpolicy_t* result = new leveldb_filterpolicy_t;
	result->state_ = state;
	result->destructor_ = destructor;
	result->create_ = create_filter;
	result->key_match_ = key_may_match;
	result->name_ = name;
	return result;
}

leveldb_filterpolicy_t* leveldb_filterpolicy_create_bloom(int32_t bits_per_key) {
	// Make a leveldb_filterpolicy_t, but override all of its methods so
	// they delegate to a NewBloomFilterPolicy() instead of user
	// supplied C functions.
	struct Wrapper : public leveldb_filterpolicy_t {
		const FilterPolicy* rep_;
		~Wrapper() { delete rep_; }
		const char* Name() const { return rep_->Name(); }
		void CreateFilter(
			const leveldb::Slice* keys, int32_t n, std::string* dst
		) const {
			return rep_->CreateFilter(keys, n, dst);
		}
		bool KeyMayMatch(
			const leveldb::Slice& key, const leveldb::Slice& filter
		) const {
			return rep_->KeyMayMatch(key, filter);
		}
		static void DoNothing(void*) { }
	};
	Wrapper* wrapper = new Wrapper;
	wrapper->state_ = NULL;
	wrapper->destructor_ = &Wrapper::DoNothing;
	wrapper->rep_ = leveldb::NewBloomFilterPolicy(bits_per_key>0? bits_per_key: 10);
	return wrapper;
}

void leveldb_filterpolicy_destroy(leveldb_filterpolicy_t* filterpolicy) {
	if(filterpolicy != NULL) {
		delete filterpolicy;
	}
}

leveldb_value_t* leveldb_filterpolicy_create_filter(
	leveldb_filterpolicy_t* filterpolicy,
	const leveldb_slice_t** keys,
	int32_t num_keys
) {
	if(filterpolicy == NULL || keys == NULL || num_keys <= 0) {
		return NULL;
	}
	std::vector<leveldb::Slice> keys_;
	keys_.resize(size_t(num_keys));
	for(int32_t i = 0; i < num_keys; ++i) {
		const leveldb_slice_t* pKey = *(keys+int32_t(i));
		keys_[i] = leveldb::Slice(pKey->data, pKey->size);
	}
	std::string tmp;
	filterpolicy->CreateFilter(&keys_[0], int32_t(num_keys), &tmp);
	return leveldb_value_create(tmp.data(), int32_t(tmp.size()));
}
leveldb_bool_t leveldb_filterpolicy_key_may_match(
	leveldb_filterpolicy_t* filterpolicy,
	const leveldb_slice_t* key,
	const leveldb_slice_t* filter
) {
	if(filterpolicy == NULL || key == NULL || filter == NULL) {
		return fromStdBool(false);
	}
	bool rv = filterpolicy->KeyMayMatch(
		leveldb::Slice(key->data, key->size),
		leveldb::Slice(filter->data, filter->size)
	);
	return fromStdBool(rv);
}
const char* leveldb_filterpolicy_name(
	leveldb_filterpolicy_t* filterpolicy
) {
	if(filterpolicy != NULL) {
		return filterpolicy->Name();
	} else {
		return NULL;
	}
}

// ----------------------------------------------------------------------------
// Cache
// ----------------------------------------------------------------------------

struct leveldb_cache_t: leveldb::Cache {
	DISABLE_NEW_AND_DELETE(leveldb_cache_t);
};

leveldb_cache_t* leveldb_cache_create_lru(int64_t capacity) {
	return (leveldb_cache_t*)(leveldb::NewLRUCache(
		size_t(capacity>0? capacity: (8<<20))
	));
}

void leveldb_cache_destroy(leveldb_cache_t* cache) {
	if(cache != NULL) {
		delete (leveldb::Cache*)cache;
	}
}

void leveldb_cache_insert(
	leveldb_cache_t* cache,
	const leveldb_slice_t* key,
	const leveldb_slice_t* val
) {
	if(cache == NULL || key == NULL || val == NULL) {
		return;
	}
	struct H {
		static void deleter(const leveldb::Slice& key, void* value) {
			leveldb_value_destroy((leveldb_value_t*)value);
		}
	};
	leveldb::Cache::Handle* h = cache->Insert(
		toLdbSlice(key),
		leveldb_value_create(val->data, val->size),
		size_t(val->size),
		H::deleter
	);
	cache->Release(h);
}

leveldb_value_t* leveldb_cache_lookup(
	leveldb_cache_t* cache,
	const leveldb_slice_t* key
) {
	if(cache == NULL || key == NULL) {
		return NULL;
	}
	leveldb::Cache::Handle* h = cache->Lookup(toLdbSlice(key));
	if(h == NULL) return NULL;
	leveldb_value_t* v = leveldb_value_create_copy((leveldb_value_t*)cache->Value(h));
	cache->Release(h);
	return v;
}

void leveldb_cache_erase(
	leveldb_cache_t* cache,
	const leveldb_slice_t* key
) {
	if(cache == NULL || key == NULL) {
		return;
	}
	cache->Erase(toLdbSlice(key));
}

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------


