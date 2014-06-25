// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import (
	"testing"
)

func TestOptions(t *testing.T) {
	//
}

func TestOptionsDefault(t *testing.T) {
	//
}

func TestReadOptions_leveldb_options_t(t *testing.T) {
	//
}

func TestReadOptions_leveldb_options_t_convert(t *testing.T) {
	//
}

func TestReadOptions(t *testing.T) {
	opt := &ReadOptions{}
	if v := opt.HasFlag(RFVerifyChecksums); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
	if v := opt.HasFlag(RFDontFillCache); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}

	opt.SetFlag(RFVerifyChecksums)
	if v := opt.HasFlag(RFVerifyChecksums); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}
	if v := opt.HasFlag(RFDontFillCache); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}

	opt.SetFlag(RFDontFillCache)
	if v := opt.HasFlag(RFVerifyChecksums); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}
	if v := opt.HasFlag(RFDontFillCache); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}

	opt.ClearFlag(RFVerifyChecksums)
	if v := opt.HasFlag(RFVerifyChecksums); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
	if v := opt.HasFlag(RFDontFillCache); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}

	opt.ClearFlag(RFDontFillCache)
	if v := opt.HasFlag(RFVerifyChecksums); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
	if v := opt.HasFlag(RFDontFillCache); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
}

func TestReadOptions_leveldb_readoptions_t(t *testing.T) {
	opt := leveldb_readoptions_create()
	defer leveldb_readoptions_destroy(opt)

	if v := leveldb_readoptions_get_verify_checksums(opt); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
	if v := leveldb_readoptions_get_fill_cache(opt); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}

	leveldb_readoptions_set_verify_checksums(opt, true)
	if v := leveldb_readoptions_get_verify_checksums(opt); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}

	leveldb_readoptions_set_fill_cache(opt, false)
	if v := leveldb_readoptions_get_fill_cache(opt); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
}

func TestReadOptions_leveldb_readoptions_t_convert(t *testing.T) {
	var ro = &ReadOptions{}
	var opt *leveldb_readoptions_t

	opt = leveldb_readoptions_create_copy(ro)
	if v := leveldb_readoptions_get_verify_checksums(opt); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
	if v := leveldb_readoptions_get_fill_cache(opt); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}
	leveldb_readoptions_destroy(opt)

	ro.SetFlag(RFVerifyChecksums)
	opt = leveldb_readoptions_create_copy(ro)
	if v := leveldb_readoptions_get_verify_checksums(opt); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}
	leveldb_readoptions_destroy(opt)

	ro.SetFlag(RFDontFillCache)
	opt = leveldb_readoptions_create_copy(ro)
	if v := leveldb_readoptions_get_fill_cache(opt); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
	leveldb_readoptions_destroy(opt)
}

func TestWriteOptions(t *testing.T) {
	opt := &WriteOptions{}
	if v := opt.HasFlag(WFSync); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
	opt.SetFlag(WFSync)
	if v := opt.HasFlag(WFSync); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}
	opt.ClearFlag(WFSync)
	if v := opt.HasFlag(WFSync); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
}

func TestWriteOptions_leveldb_writeoptions_t(t *testing.T) {
	opt := leveldb_writeoptions_create()
	defer leveldb_writeoptions_destroy(opt)

	if v := leveldb_writeoptions_get_sync(opt); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
	leveldb_writeoptions_set_sync(opt, true)
	if v := leveldb_writeoptions_get_sync(opt); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}
	leveldb_writeoptions_set_sync(opt, false)
	if v := leveldb_writeoptions_get_sync(opt); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
}

func TestWriteOptions_leveldb_writeoptions_t_convert(t *testing.T) {
	var wo = &WriteOptions{}
	var opt *leveldb_writeoptions_t

	opt = leveldb_writeoptions_create_copy(wo)
	if v := leveldb_writeoptions_get_sync(opt); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
	leveldb_writeoptions_destroy(opt)

	opt = leveldb_writeoptions_create_copy(wo)
	if v := leveldb_writeoptions_get_sync(opt); v != false {
		t.Fatalf("expect = false, got = %v", v)
	}
	leveldb_writeoptions_destroy(opt)

	wo.SetFlag(WFSync)
	opt = leveldb_writeoptions_create_copy(wo)
	if v := leveldb_writeoptions_get_sync(opt); v != true {
		t.Fatalf("expect = true, got = %v", v)
	}
	leveldb_writeoptions_destroy(opt)
}
