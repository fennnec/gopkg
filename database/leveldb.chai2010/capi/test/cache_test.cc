// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "test.h"
#include "leveldb_c.h"
#include "leveldb/slice.h"
#include "leveldb/filter_policy.h"

#include <string.h>
#include <vector>
#include <map>

static struct {
	const char *key;
	const char *val;
} cachePairs[] = {
	{"", ""},
	{"f", "Zg=="},
	{"fo", "Zm8="},
	{"foo", "Zm9v"},
	{"foob", "Zm9vYg=="},
	{"fooba", "Zm9vYmE="},
	{"foobar", "Zm9vYmFy"},
	{"sure.", "c3VyZS4="},
	{"sure", "c3VyZQ=="},
	{"sur", "c3Vy"},
	{"su", "c3U="},
	{"leasure.", "bGVhc3VyZS4="},
	{"easure.", "ZWFzdXJlLg=="},
	{"asure.", "YXN1cmUu"},
	{"sure.", "c3VyZS4="},
};

struct AutoCacheDeleter {
	leveldb_cache_t* cache_;
	AutoCacheDeleter(leveldb_cache_t* p): cache_(p) {}
	~AutoCacheDeleter() { leveldb_cache_destroy(cache_); }
};

TEST(Cache, Simple) {
	auto cache = leveldb_cache_create_lru(1 << 20);
	AutoCacheDeleter cacheDeleter(cache);

	// empty & lookup fail
	for(int i = 0; i < DIM(cachePairs); ++i) {
		leveldb_slice_t key = { cachePairs[i].key, strlen(cachePairs[i].key)+1 };
		leveldb_slice_t val = { cachePairs[i].val, strlen(cachePairs[i].val)+1 };
		auto p = leveldb_cache_lookup(cache, &key);
		ASSERT_TRUE(p == NULL);
	}

	// insert & lookup success
	for(int i = 0; i < DIM(cachePairs); ++i) {
		leveldb_slice_t key = { cachePairs[i].key, strlen(cachePairs[i].key)+1 };
		leveldb_slice_t val = { cachePairs[i].val, strlen(cachePairs[i].val)+1 };
		leveldb_cache_insert(cache, &key, &val);
	}
	for(int i = 0; i < DIM(cachePairs); ++i) {
		leveldb_slice_t key = { cachePairs[i].key, strlen(cachePairs[i].key)+1 };
		leveldb_slice_t val = { cachePairs[i].val, strlen(cachePairs[i].val)+1 };
		auto p = leveldb_cache_lookup(cache, &key);
		ASSERT_TRUE(p != NULL);

		leveldb_slice_t dat = { leveldb_value_data(p), leveldb_value_size(p) };
		ASSERT_TRUE(val.size == dat.size);
		int cmp = memcmp(val.data, dat.data, val.size);
		ASSERT_TRUE(cmp == 0);

		leveldb_value_destroy(p);
	}

	// erase & lookup fail
	for(int i = 0; i < DIM(cachePairs); ++i) {
		leveldb_slice_t key = { cachePairs[i].key, strlen(cachePairs[i].key)+1 };
		leveldb_slice_t val = { cachePairs[i].val, strlen(cachePairs[i].val)+1 };
		leveldb_cache_erase(cache, &key);
	}
	for(int i = 0; i < DIM(cachePairs); ++i) {
		leveldb_slice_t key = { cachePairs[i].key, strlen(cachePairs[i].key)+1 };
		leveldb_slice_t val = { cachePairs[i].val, strlen(cachePairs[i].val)+1 };
		auto p = leveldb_cache_lookup(cache, &key);
		ASSERT_TRUE(p == NULL);
	}
}

