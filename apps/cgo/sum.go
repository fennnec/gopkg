// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Linking golang statically demo.
package cgo

// BUG(chai2010): on windows/amd64, sizeof(C.int) is 4

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -static

#include "sum.h"
*/
import "C"
import "unsafe"

func Sum(a []int32) int {
	v := C.sum((*C.int)(unsafe.Pointer(&a[0])), C.int(len(a)))
	return int(v)
}
