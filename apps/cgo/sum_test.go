// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

import (
	"testing"
)

func TestSum(t *testing.T) {
	if v := Sum([]int32{1, 2, 3, 4, 5}); v != 15 {
		t.Fatalf("expect = %v. got = %v", 15, v)
	}
}
