// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import "runtime"

// ----------------------------------------------------------------------------
// Value
// ----------------------------------------------------------------------------

// Value represents an c memory data.
//
// The caller should not modify the contents of the returned slice.
type Value struct {
	*valueHandler
}
type valueHandler struct {
	value *leveldb_value_t
}

func newValueHandler(p *leveldb_value_t) *Value {
	if p == nil {
		return nil
	}
	h := &valueHandler{p}
	runtime.SetFinalizer(h, (*valueHandler).Release)
	return &Value{h}
}

func (p *valueHandler) Slice() []byte {
	if p.value != nil {
		return leveldb_value_data(p.value)
	}
	return nil
}

func (p *valueHandler) Release() {
	if p.value != nil {
		leveldb_value_destroy(p.value)
	}
	runtime.SetFinalizer(p, nil)
}

// ----------------------------------------------------------------------------
// Options
// ----------------------------------------------------------------------------

// Avoid GC
type optStateContainer struct {
	cmp             *leveldb_comparator_t
	filter          *leveldb_filterpolicy_t
	cache           *Cache
	stateComparator stateComparator
	stateFilter     stateFilter
}

func (p *optStateContainer) Release() {
	if p.cmp != nil {
		leveldb_comparator_destroy(p.cmp)
	}
	if p.filter != nil {
		leveldb_filterpolicy_destroy(p.filter)
	}
	*p = optStateContainer{}
}

func leveldb_options_create_copy(state *optStateContainer, opt *Options) *leveldb_options_t {
	p := leveldb_options_create()
	if opt != nil {
		state.Release()
		if opt.GetComparator() != nil {
			state.stateComparator.Comparator = opt.GetComparator()
			state.cmp = leveldb_comparator_create(&state.stateComparator)
			leveldb_options_set_comparator(p, state.cmp)
		}

		leveldb_options_set_create_if_missing(p, opt.HasFlag(OFCreateIfMissing))
		leveldb_options_set_error_if_exists(p, opt.HasFlag(OFErrorIfExist))
		leveldb_options_set_paranoid_checks(p, opt.HasFlag(OFStrict))

		if v := opt.GetWriteBufferSize(); v > 0 {
			leveldb_options_set_write_buffer_size(p, v)
		}

		if v := opt.GetMaxOpenFiles(); v > 0 {
			leveldb_options_set_max_open_files(p, v)
		}

		if v := opt.GetBlockCache(); v != nil {
			state.cache = v
			leveldb_options_set_cache(p, state.cache.cacheHandler.cache)
		}

		if v := opt.GetBlockSize(); v > 0 {
			leveldb_options_set_block_size(p, v)
		}

		if v := opt.GetBlockRestartInterval(); v > 0 {
			leveldb_options_set_block_restart_interval(p, v)
		}

		switch opt.GetCompression() {
		case NoCompression:
			leveldb_options_set_compression(p, leveldb_compression_nil)
		case SnappyCompression:
			leveldb_options_set_compression(p, leveldb_compression_snappy)
		}

		if opt.GetFilter() != nil {
			state.stateFilter.Filter = opt.GetFilter()
			state.filter = leveldb_filterpolicy_create(&state.stateFilter)
			leveldb_options_set_filter_policy(p, state.filter)
		}
	}
	return p
}

func leveldb_readoptions_create_copy(opt *ReadOptions) *leveldb_readoptions_t {
	p := leveldb_readoptions_create()
	if opt != nil {
		leveldb_readoptions_set_verify_checksums(p, opt.HasFlag(RFVerifyChecksums))
		leveldb_readoptions_set_fill_cache(p, !opt.HasFlag(RFDontFillCache))
		if opt.Snapshot != nil && opt.Snapshot.snap != nil {
			leveldb_readoptions_set_snapshot(p, opt.Snapshot.snap)
		}
	}
	return p
}

func leveldb_writeoptions_create_copy(opt *WriteOptions) *leveldb_writeoptions_t {
	p := leveldb_writeoptions_create()
	if opt != nil {
		leveldb_writeoptions_set_sync(p, opt.HasFlag(WFSync))
	}
	return p
}

// ----------------------------------------------------------------------------
// Callback
// ----------------------------------------------------------------------------

type stateWriteBatchIterater struct {
	WriteBatchIterater
}

type stateComparator struct {
	Comparator
	cname []byte
}

type stateFilter struct {
	Filter
	cname []byte
}

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------
