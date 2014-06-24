// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

import (
	"fmt"
	"reflect"
)

func MapSlice(slice interface{}, fn func(a interface{}) interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("MapSlice called with non-slice value of type %T", slice))
	}
	val := reflect.ValueOf(slice)
	out := reflect.MakeSlice(reflect.TypeOf(slice), val.Len(), val.Cap())
	for i := 0; i < val.Len(); i++ {
		out.Index(i).Set(
			reflect.ValueOf(fn(val.Index(i).Interface())),
		)
	}
	return out.Interface()
}
