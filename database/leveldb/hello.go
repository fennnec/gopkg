// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	"github.com/p/chai2010/gopkg/database/leveldb"
)

func main() {
	fmt.Printf("leveldb-%d.%d\n", leveldb.MajorVersion, leveldb.MinorVersion)
}
