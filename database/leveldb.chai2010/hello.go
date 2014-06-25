// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"log"

	"chai2010.gopkg/database/leveldb.chai2010"
)

func main() {
	fmt.Printf("hello leveldb-%d.%d!\n",
		leveldb.MajorVersion,
		leveldb.MinorVersion,
	)

	opt := &leveldb.Options{}
	opt.SetFlag(leveldb.OFCreateIfMissing)
	opt.SetBlockCacheCapacity(32 << 20)

	db, err := leveldb.Open("testdb", opt)
	if err != nil {
		log.Fatal(err) // *leveldb.Error
	}
	defer db.Close()

	db.Put([]byte("key1"), []byte("value1"), nil)
	db.Put([]byte("key2"), []byte("value2"), nil)
	db.Put([]byte("key3"), []byte("value3"), nil)

	it := db.NewIterator(nil)
	defer it.Release()

	for it.SeekToFirst(); it.Valid(); it.Next() {
		fmt.Printf("%s: %s\n", string(it.Key()), string(it.Value()))
	}
	if err := it.GetError(); err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
