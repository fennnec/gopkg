// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin_test

import (
	"fmt"
	"image/color"
	"reflect"
	"sort"

	. "github.com/chai2010/gopkg/builtin"
)

func ExampleIf() {
	a, b := 42, 9527
	fmt.Printf("max: %d\n", If(a > b, a, b).(int))
	fmt.Printf("min: %d\n", If(a < b, a, b).(int))
	// Output:
	// max: 9527
	// min: 42
}

func ExampleIterN() {
	for i := range IterN(5) {
		fmt.Println(i)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleByteSlice() {
	src := []color.Gray{color.Gray{0xAA}, color.Gray{0xBB}, color.Gray{0xCC}, color.Gray{0xDD}}
	dst := make([]byte, len(src))
	copy(ByteSlice(dst), ByteSlice(src))
	fmt.Printf("%X", dst)
	// Output: AABBCCDD
}

func ExampleSlice() {
	src := []byte{0xAA, 0xBB, 0xCC, 0xDD}
	dst := Slice(src, reflect.TypeOf([]color.Gray(nil))).([]color.Gray)
	fmt.Printf("%X", dst)
	// Output:
	// [{AA} {BB} {CC} {DD}]
}

func ExampleSort_int() {
	arr := []int32{88, 56, 100, 2, 25}

	Sort(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
	// Output:
	// [2 25 56 88 100]
}

func ExampleSort_string() {
	arr := []string{"coffee", "flour", "tea"}

	Sort(arr, func(i, j int) bool {
		return arr[i] > arr[j] // descending
	})
	fmt.Println(arr)
	// Output:
	// [tea flour coffee]
}

func ExampleSortInterface() {
	arr := []int32{88, 56, 100, 2, 25}
	sort.Sort(SortInterface(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	}))
	fmt.Println(arr)
	// Output:
	// [2 25 56 88 100]
}
