// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import (
	"runtime"
)

// Cache is a cache used to store data read from data in memory.
type Cache struct {
	*cacheHandler
}
type cacheHandler struct {
	cache *leveldb_cache_t
}

// NewCache creates a new cache with a fixed size capacity.
func NewCache(capacity int64) *Cache {
	p := &cacheHandler{
		cache: leveldb_cache_create_lru(capacity),
	}
	runtime.SetFinalizer(p, (*cacheHandler).Release)
	return &Cache{p}
}

// Release releases deallocates the underlying memory of the Cache object.
func (p *cacheHandler) Release() {
	runtime.SetFinalizer(p, nil)
	leveldb_cache_destroy(p.cache)
	*p = cacheHandler{}
}

// Insert inserts a mapping from key->value into the cache and assign it
// the specified charge against the total cache capacity.
func (p *cacheHandler) Insert(key, val []byte) {
	leveldb_cache_insert(p.cache, key, val)
}

// Lookup get the cached value by key.
// If the cache has no mapping for "key", returns nil.
func (p *cacheHandler) Lookup(key []byte) *Value {
	return newValueHandler(leveldb_cache_lookup(p.cache, key))
}

// Erase erase the cached value by key.
func (p *cacheHandler) Erase(key []byte) {
	leveldb_cache_erase(p.cache, key)
}
