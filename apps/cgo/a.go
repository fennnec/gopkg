// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"github.com/chai2010/gopkg/apps/cgo"
)

func main() {
	a := []int32{1, 2, 3, 4, 5}
	println(cgo.Sum(a))
}
