// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import (
	"bytes"
	"os"
	"testing"
)

var testdb = "testdb"

// https://code.google.com/p/leveldb/issues/detail?id=207
func TestIssue207(t *testing.T) {
	opt := leveldb_options_create()
	defer leveldb_options_destroy(opt)

	leveldb_options_set_create_if_missing(opt, true)
	db, err := leveldb_open(testdb, opt)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(testdb)
	defer leveldb_close(db)

	key := []byte("key")
	val := []byte("value")

	data, err := leveldb_get(db, key, nil)
	if err != ErrNotFound {
		t.Fatalf("expect ErrNotFound, got %v", err)
	}

	err = leveldb_put(db, key, val, nil)
	if err != nil {
		t.Fatal(err)
	}
	data, err = leveldb_get(db, key, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, val) {
		t.Fatalf("data not equal, got %v", data)
	}

	err = leveldb_delete(db, key, nil)
	if err != nil {
		t.Fatal(err)
	}
	data, err = leveldb_get(db, key, nil)
	if err != ErrNotFound {
		t.Fatalf("expect ErrNotFound, got %v", err)
	}
}
