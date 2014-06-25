// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "test.h"
#include "leveldb_c.h"
#include "leveldb/slice.h"
#include "leveldb/options.h"
#include "leveldb/filter_policy.h"

#include <string.h>
#include <vector>
#include <map>

static leveldb_bool_t fromStdBool(bool v) {
	return (v)? 1: 0;
}

static leveldb_slice_t newSlice(const char* data, int32_t size) {
	leveldb_slice_t a = { data, size };
	return a;
}

static bool strHasPrefix(const char* str, const char* prefix) {
	int a = strlen(str);
	int b = strlen(prefix);
	if(a < b) return false;
	for(int i = 0; i < b; ++i) {
		if(str[i] != prefix[i]) return false;
	}
	return true;
}

TEST(Version, Simple) {
	ASSERT_TRUE(leveldb_major_version() == 1);
	ASSERT_TRUE(leveldb_minor_version() == 14);
}

TEST(CompressionType, Simple) {
	ASSERT_TRUE(leveldb_compression_nil == leveldb::kNoCompression);
	ASSERT_TRUE(leveldb_compression_snappy == leveldb::kSnappyCompression);
}

TEST(Value, Simple) {
	leveldb_value_t* value;
	const char* data = "hello go-leveldb!";

	value = leveldb_value_create(NULL, 0);

	ASSERT_TRUE(leveldb_value_data(value) == NULL);
	ASSERT_TRUE(leveldb_value_size(value) == 0);
	leveldb_value_destroy(value);

	value = leveldb_value_create(NULL, 10);
	ASSERT_TRUE(leveldb_value_data(value) == NULL);
	ASSERT_TRUE(leveldb_value_size(value) == 0);
	leveldb_value_destroy(value);

	value = leveldb_value_create(data, 0);
	ASSERT_TRUE(leveldb_value_data(value) == NULL);
	ASSERT_TRUE(leveldb_value_size(value) == 0);
	leveldb_value_destroy(value);

	value = leveldb_value_create(data, strlen(data));
	ASSERT_TRUE(leveldb_value_data(value) != data);
	ASSERT_TRUE(leveldb_value_data(value) != NULL);
	ASSERT_TRUE(leveldb_value_size(value) == strlen(data));
	ASSERT_STREQ(leveldb_value_data(value), data);

	leveldb_value_destroy(value);
}

TEST(Options, Simple) {
	leveldb_options_t* opt;
	leveldb_comparator_t* cmp = (leveldb_comparator_t*)"leveldb_comparator_t";
	leveldb_filterpolicy_t* filter = (leveldb_filterpolicy_t*)"leveldb_filterpolicy_t";
	leveldb_cache_t* cache = (leveldb_cache_t*)"leveldb_cache_t";

	opt = leveldb_options_create();

	ASSERT_TRUE(leveldb_options_get_comparator(opt) != NULL);
	leveldb_options_set_comparator(opt, cmp);
	ASSERT_TRUE(leveldb_options_get_comparator(opt) == cmp);
	leveldb_options_set_comparator(opt, NULL);
	ASSERT_TRUE(leveldb_options_get_comparator(opt) == NULL);

	ASSERT_TRUE(leveldb_options_get_filter_policy(opt) == NULL);
	leveldb_options_set_filter_policy(opt, filter);
	ASSERT_TRUE(leveldb_options_get_filter_policy(opt) == filter);
	leveldb_options_set_filter_policy(opt, NULL);
	ASSERT_TRUE(leveldb_options_get_filter_policy(opt) == NULL);

	ASSERT_TRUE(leveldb_options_get_create_if_missing(opt) == fromStdBool(false));
	leveldb_options_set_create_if_missing(opt, true);
	ASSERT_TRUE(leveldb_options_get_create_if_missing(opt) == fromStdBool(true));
	leveldb_options_set_create_if_missing(opt, false);
	ASSERT_TRUE(leveldb_options_get_create_if_missing(opt) == fromStdBool(false));

	ASSERT_TRUE(leveldb_options_get_error_if_exists(opt) == fromStdBool(false));
	leveldb_options_set_error_if_exists(opt, true);
	ASSERT_TRUE(leveldb_options_get_error_if_exists(opt) == fromStdBool(true));
	leveldb_options_set_error_if_exists(opt, false);
	ASSERT_TRUE(leveldb_options_get_error_if_exists(opt) == fromStdBool(false));

	ASSERT_TRUE(leveldb_options_get_paranoid_checks(opt) == fromStdBool(false));
	leveldb_options_set_paranoid_checks(opt, true);
	ASSERT_TRUE(leveldb_options_get_paranoid_checks(opt) == fromStdBool(true));
	leveldb_options_set_paranoid_checks(opt, false);
	ASSERT_TRUE(leveldb_options_get_paranoid_checks(opt) == fromStdBool(false));

	ASSERT_TRUE(leveldb_options_get_max_open_files(opt) == 1000);
	leveldb_options_set_max_open_files(opt, 500);
	ASSERT_TRUE(leveldb_options_get_max_open_files(opt) == 500);
	leveldb_options_set_max_open_files(opt, 100);
	ASSERT_TRUE(leveldb_options_get_max_open_files(opt) == 100);

	ASSERT_TRUE(leveldb_options_get_cache(opt) == NULL);
	leveldb_options_set_cache(opt, cache);
	ASSERT_TRUE(leveldb_options_get_cache(opt) == cache);
	leveldb_options_set_cache(opt, NULL);
	ASSERT_TRUE(leveldb_options_get_cache(opt) == NULL);

	ASSERT_TRUE(leveldb_options_get_block_size(opt) == 4096);
	leveldb_options_set_block_size(opt, 1024);
	ASSERT_TRUE(leveldb_options_get_block_size(opt) == 1024);
	leveldb_options_set_block_size(opt, 100);
	ASSERT_TRUE(leveldb_options_get_block_size(opt) == 100);

	ASSERT_TRUE(leveldb_options_get_block_restart_interval(opt) == 16);
	leveldb_options_set_block_restart_interval(opt, 10);
	ASSERT_TRUE(leveldb_options_get_block_restart_interval(opt) == 10);
	leveldb_options_set_block_restart_interval(opt, 100);
	ASSERT_TRUE(leveldb_options_get_block_restart_interval(opt) == 100);

	ASSERT_TRUE(leveldb_options_get_compression(opt) == leveldb_compression_snappy);
	leveldb_options_set_compression(opt, leveldb_compression_nil);
	ASSERT_TRUE(leveldb_options_get_compression(opt) == leveldb_compression_nil);
	leveldb_options_set_compression(opt, leveldb_compression_snappy);
	ASSERT_TRUE(leveldb_options_get_compression(opt) == leveldb_compression_snappy);

	leveldb_options_destroy(opt);
}

TEST(ReadOptions, Simple) {
	leveldb_readoptions_t* opt;
	leveldb_snapshot_t* snap = (leveldb_snapshot_t*)"leveldb_snapshot_t";

	opt = leveldb_readoptions_create();

	ASSERT_TRUE(leveldb_readoptions_get_verify_checksums(opt) == fromStdBool(false));
	leveldb_readoptions_set_verify_checksums(opt, true);
	ASSERT_TRUE(leveldb_readoptions_get_verify_checksums(opt) == fromStdBool(true));
	leveldb_readoptions_set_verify_checksums(opt, false);
	ASSERT_TRUE(leveldb_readoptions_get_verify_checksums(opt) == fromStdBool(false));

	ASSERT_TRUE(leveldb_readoptions_get_fill_cache(opt) == fromStdBool(true));
	leveldb_readoptions_set_fill_cache(opt, false);
	ASSERT_TRUE(leveldb_readoptions_get_fill_cache(opt) == fromStdBool(false));
	leveldb_readoptions_set_fill_cache(opt, true);
	ASSERT_TRUE(leveldb_readoptions_get_fill_cache(opt) == fromStdBool(true));

	ASSERT_TRUE(leveldb_readoptions_get_snapshot(opt) == NULL);
	leveldb_readoptions_set_snapshot(opt, snap);
	ASSERT_TRUE(leveldb_readoptions_get_snapshot(opt) == snap);
	leveldb_readoptions_set_snapshot(opt, NULL);
	ASSERT_TRUE(leveldb_readoptions_get_snapshot(opt) == NULL);

	leveldb_readoptions_destroy(opt);
}

TEST(WriteOptions, Simple) {
	leveldb_writeoptions_t* opt;

	opt = leveldb_writeoptions_create();

	ASSERT_TRUE(leveldb_writeoptions_get_sync(opt) == fromStdBool(false));
	leveldb_writeoptions_set_sync(opt, true);
	ASSERT_TRUE(leveldb_writeoptions_get_sync(opt) == fromStdBool(true));
	leveldb_writeoptions_set_sync(opt, false);
	ASSERT_TRUE(leveldb_writeoptions_get_sync(opt) == fromStdBool(false));

	leveldb_writeoptions_destroy(opt);
}

TEST(DB, Simple) {
	// TODO
}

TEST(Iterator, Simple) {
	// TODO
}

TEST(WriteBatch, Simple) {
	struct H {
		static void put(
			void* state,
			const leveldb_slice_t* key,
			const leveldb_slice_t* val
		) {
			auto dict = (std::map<std::string,std::string>*)state;
			(*dict)[std::string(key->data)] = std::string("put:") + val->data;
		}
		static void deleted(
			void* state,
			const leveldb_slice_t* key
		) {
			auto dict = (std::map<std::string,std::string>*)state;
			(*dict)[std::string(key->data)] = std::string("delete:");
		}
	};

	leveldb_writebatch_t* batch;
	leveldb_slice_t key, val;

	batch = leveldb_writebatch_create();

	key = newSlice("key1", strlen("key1")+1);
	val = newSlice("val1", strlen("val1")+1);
	leveldb_writebatch_put(batch, &key, &val);

	key = newSlice("key2", strlen("key2")+1);
	val = newSlice("val2", strlen("val2")+1);
	leveldb_writebatch_put(batch, &key, &val);

	key = newSlice("key1", strlen("key1")+1);
	leveldb_writebatch_delete(batch, &key);

	key = newSlice("key1", strlen("key1")+1);
	val = newSlice("val1-new", strlen("val1-new")+1);
	leveldb_writebatch_put(batch, &key, &val);

	key = newSlice("key-unkown", strlen("key-unkown")+1);
	leveldb_writebatch_delete(batch, &key);

	std::map<std::string,std::string> dict;
	leveldb_writebatch_iterate(batch, &dict, H::put, H::deleted);

	// key-unkown: delete:
	// key1: put:val1-new
	// key2: put:val2
	ASSERT_TRUE(dict.size() == 3);
	ASSERT_STREQ(dict["key-unkown"].c_str(), "delete:");
	ASSERT_STREQ(dict["key1"].c_str(), "put:val1-new");
	ASSERT_STREQ(dict["key2"].c_str(), "put:val2");

	leveldb_writebatch_destroy(batch);
}

TEST(Comparator, Simple) {
	struct H {
		static void destructor(void* state) {
			// Empty
		}
		static int32_t compare(void* state,
			const leveldb_slice_t* a,
			const leveldb_slice_t* b
		) {
			leveldb::Slice a_(a->data, a->size);
			leveldb::Slice b_(b->data, b->size);
			return int32_t(a_.compare(b_));
		}
		static const char* name(void* state) {
			return "goleveldb.comparator.test";
		}
	};

	leveldb_comparator_t* cmp;
	cmp = leveldb_comparator_create(NULL, H::destructor, H::compare, H::name);
	leveldb_comparator_destroy(cmp);
}

TEST(FilterPolicy, Simple) {
	struct H {
		static void destructor(void* state) {
			leveldb::FilterPolicy* bloomFilter = (leveldb::FilterPolicy*)state;
			delete bloomFilter;
		}
		static leveldb_value_t* create_filter(
			void* state,
			const leveldb_slice_t** keys,
			int32_t num_keys
		) {
			leveldb::FilterPolicy* bloomFilter = (leveldb::FilterPolicy*)state;

			std::vector<leveldb::Slice> keys_(num_keys);
			for(int32_t i = 0; i < num_keys; ++i) {
				keys_[i] = leveldb::Slice(keys[i]->data, keys[i]->size);
			}

			std::string tmp;
			bloomFilter->CreateFilter(&keys_[0], num_keys, &tmp);
			return leveldb_value_create(tmp.data(), tmp.size());
		}
		static leveldb_bool_t key_may_match(
			void* state,
			const leveldb_slice_t* key,
			const leveldb_slice_t* filter
		) {
			leveldb::FilterPolicy* bloomFilter = (leveldb::FilterPolicy*)state;

			leveldb::Slice key_(key->data, key->size);
			leveldb::Slice filter_(filter->data, filter->size);
			bool rv = bloomFilter->KeyMayMatch(key_, filter_);
			return rv? 1: 0;
		}
		static const char* name(void* state) {
			return "goleveldb.filterpolicy.test";
		}
	};

	leveldb_filterpolicy_t* filter;

	filter = leveldb_filterpolicy_create_bloom(13);
	leveldb_filterpolicy_destroy(filter);

	filter = leveldb_filterpolicy_create(
		(void*)leveldb::NewBloomFilterPolicy(15),
		H::destructor,
		H::create_filter,
		H::key_may_match,
		H::name
	);
	leveldb_filterpolicy_destroy(filter);
}
