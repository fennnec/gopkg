// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package leveldb provides Go bindings for LevelDB.

leveldb.Open opens and creates databases.

	opt := &leveldb.Options{}
	opt.SetBlockCacheCapacity(64 << 20)
	opt.SetFlag(leveldb.OFCreateIfMissing)
	db, err := leveldb.Open("/path/to/db", opt)
	if err != nil {
		log.Fatal(err) // *leveldb.Error
	}
	defer db.Close()

The DB struct returned by Open provides DB.Get, DB.Put and DB.Delete to modify
and query the database.

	data, err := db.Get([]byte("key"), nil)
	...
	err = db.Put([]byte("anotherkey"), data, nil)
	...
	err = db.Delete([]byte("key"), nil)

For bulk reads, use an Iterator. If you want to avoid disturbing your live
traffic while doing the bulk read, be sure to call SetFillCache(false) on the
ReadOptions you use when creating the Iterator.

	ro := &leveldb.ReadOptions{}
	ro.SetFlag(leveldb.RFDontFillCache)
	it := db.NewIterator(ro)
	defer it.Release()

	for it.Seek(mykey); it.Valid(); it.Next() {
		doSomething(it.Key(), it.Value())
	}
	if err := it.GetError(); err != nil {
		// ...
	}

Batched, atomic writes can be performed with a WriteBatch and
DB.Write.

	b := leveldb.NewWriteBatch()
	defer b.Release()

	b.Delete([]byte("removed"))
	b.Put([]byte("added"), []byte("data"))
	b.Put([]byte("anotheradded"), []byte("more"))
	err := db.Write(b, nil)

If your working dataset does not fit in memory, you'll want to add a bloom
filter to your database. NewBloomFilter and Options.SetFilterPolicy is what
you want. NewBloomFilter is amount of bits in the filter to use per key in
your database.

	filter := leveldb.NewBloomFilter(10)
	opt := &leveldb.Options{}
	opt.SetFilter(leveldb.NewBloomFilter(10))
	db, err := leveldb.Open("/path/to/db", opt)

If you're using a custom comparator in your code, be aware you may have to
make your own filter policy object.

This documentation is not a complete discussion of LevelDB. Please read the
LevelDB documentation <http://code.google.com/p/leveldb> for information on
its operation. You'll find lots of goodies there.
*/
package leveldb
