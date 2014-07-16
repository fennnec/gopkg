// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import (
	"bytes"
	"testing"
)

type testpair struct {
	key, value string
}

var pairs = []testpair{
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
}

func TestCache(t *testing.T) {
	cache := NewCache(1 << 20)
	defer cache.Release()

	// empty & lookup fail
	for _, p := range pairs {
		if v := cache.Lookup([]byte(p.key)); v != nil {
			t.Fatalf("%s: expect = nil, got = not nil\n", p.key)
		}
	}

	// insert & lookup success
	for _, p := range pairs {
		cache.Insert([]byte(p.key), []byte(p.value))
	}
	for _, p := range pairs {
		v := cache.Lookup([]byte(p.key))
		if v == nil {
			t.Fatalf("%s: expect = not nil, got = nil\n", p.key)
		}
		if !bytes.Equal([]byte(p.value), v.Slice()) {
			t.Fatalf("%s: expect = %v, got = %v\n", p.key, p.value, string(v.Slice()))
		}
		v.Release()
	}

	// erase & lookup fail
	for _, p := range pairs {
		cache.Erase([]byte(p.key))
	}
	for _, p := range pairs {
		if v := cache.Lookup([]byte(p.key)); v != nil {
			t.Fatalf("%s: expect = nil, got = not nil\n", p.key)
		}
	}
}
