// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

import (
	"reflect"
	"testing"
)

func TestMapSlice(t *testing.T) {
	a := MapSlice([]int{1, 2, 3, 4}, func(val interface{}) interface{} {
		return val.(int) * 2
	})
	if !reflect.DeepEqual(a, []int{2, 4, 6, 8}) {
		t.Fatal("not equal")
	}
}
