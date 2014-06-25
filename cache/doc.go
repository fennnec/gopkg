// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package cache is an in-memory key/value cache.

Package cache is an in-memory key/value store/cache similar to memcached that is
suitable for applications running on a single machine. Its major advantage is
that, being essentially a thread-safe map[string]interface{} with expiration
times, it doesn't need to serialize or transmit its contents over the network.

Any object can be stored, for a given duration or forever, and the cache can be
safely used by multiple goroutines.

Although go-cache isn't meant to be used as a persistent datastore, the entire
cache may be saved to and loaded from a file (or any io.Reader/Writer) to
recover from downtime quickly.

Usage

	import "code.google.com/p/chai2010.gopkg/cache"

	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 30 seconds
	c := cache.New(5*time.Minute, 30*time.Second)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", 0)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, -1)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	// Since Go is statically typed, and cache values can be anything, type
	// assertion is needed when values are being passed to functions that don't
	// take arbitrary types, (i.e. interface{}). The simplest way to do this for
	// values which will only be used once--e.g. for passing to another
	// function--is:
	foo, found := c.Get("foo")
	if found {
		MyFunction(foo.(string))
	}

	// This gets tedious if the value is used several times in the same function.
	// You might do either of the following instead:
	if x, found := c.Get("foo"); found {
		foo := x.(string)
		...
	}
	// or
	var foo string
	if x, found := c.Get("foo"); found {
		foo = x.(string)
	}
	...
	// foo can then be passed around freely as a string

	// Want performance? Store pointers!
	c.Set("foo", &MyStruct, 0)
	if x, found := c.Get("foo"); found {
		foo := x.(*MyStruct)
		...
	}

	// If you store a reference type like a pointer, slice, map or channel, you
	// do not need to run Set if you modify the underlying data. The cached
	// reference points to the same memory, so if you modify a struct whose
	// pointer you've stored in the cache, retrieving that pointer with Get will
	// point you to the same data:
	foo := &MyStruct{Num: 1}
	c.Set("foo", foo, 0)
	...
	x, _ := c.Get("foo")
	foo := x.(*MyStruct)
	fmt.Println(foo.Num)
	...
	foo.Num++
	...
	x, _ := c.Get("foo")
	foo := x.(*MyStruct)
	foo.Println(foo.Num)

	// will print:
	1
	2

LRUCache

The implementation borrows heavily from SmallLRUCache (originally by Nathan
Schrenk). The object maintains a doubly-linked list of elements in the
When an element is accessed it is promoted to the head of the list, and when
space is needed the element at the tail of the list (the least recently used
element) is evicted.
*/
package cache
