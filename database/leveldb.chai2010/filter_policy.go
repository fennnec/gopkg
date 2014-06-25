// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import (
	"runtime"
)

// Filter is the filter.
type Filter interface {
	// Return the name of this policy.  Note that if the filter encoding
	// changes in an incompatible way, the name returned by this method
	// must be changed.  Otherwise, old incompatible filters may be
	// passed to methods of this type.
	Name() string

	// keys[0,n-1] contains a list of keys (potentially with duplicates)
	// that are ordered according to the user supplied comparator.
	// Append a filter that summarizes keys[0,n-1] to *dst.
	//
	// Warning: do not change the initial contents of *dst.  Instead,
	// append the newly constructed filter to *dst.
	CreateFilter(keys [][]byte) []byte

	// "filter" contains the data appended by a preceding call to
	// CreateFilter() on this class.  This method must return true if
	// the key was in the list of keys passed to CreateFilter().
	// This method may return true or false if the key was not on the
	// list, but it should aim to return false with a high probability.
	KeyMayMatch(key, filter []byte) bool
}

type BloomFilter struct {
	*bloomFilterHandler
}
type bloomFilterHandler struct {
	filter *leveldb_filterpolicy_t
}

func newBloomFilter(bitsPerKey int) *BloomFilter {
	p := &bloomFilterHandler{
		filter: leveldb_filterpolicy_create_bloom(bitsPerKey),
	}
	runtime.SetFinalizer(p, (*bloomFilterHandler).Release)
	return &BloomFilter{p}
}

func (p *bloomFilterHandler) Release() {
	runtime.SetFinalizer(p, nil)
	leveldb_filterpolicy_destroy(p.filter)
	*p = bloomFilterHandler{}
}

func (p *bloomFilterHandler) Name() string {
	return leveldb_filterpolicy_name(p.filter)
}

func (p *bloomFilterHandler) CreateFilter(keys [][]byte) []byte {
	d := leveldb_filterpolicy_create_filter(p.filter, keys)
	defer leveldb_value_destroy(d)
	return append([]byte{}, leveldb_value_data(d)...)
}

func (p *bloomFilterHandler) KeyMayMatch(key, filter []byte) bool {
	return leveldb_filterpolicy_key_may_match(p.filter, key, filter)
}

// NewBloomFilter creates a new initialized bloom filter for given
// bitsPerKey.
//
// Since bitsPerKey is persisted individually for each bloom filter
// serialization, bloom filters are backwards compatible with respect to
// changing bitsPerKey. This means that no big performance penalty will
// be experienced when changing the parameter. See documentation for
// opt.Options.Filter for more information.
func NewBloomFilterPolicy(bitsPerKey int) *BloomFilter {
	return newBloomFilter(bitsPerKey)
}
