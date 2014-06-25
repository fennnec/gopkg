// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

// #include <leveldb/c.h>
import "C"

var (
	MajorVersion = int(C.leveldb_major_version()) // leveldb-1.14
	MinorVersion = int(C.leveldb_minor_version())
)
