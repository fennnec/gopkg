// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

// Comparator defines a total ordering over the space
// of []byte keys: a 'less than' relationship.
type Comparator interface {
	// Three-way comparison.  Returns value:
	//   <  0 if "a" <  "b",
	//   == 0 if "a" == "b",
	//   >  0 if "a" >  "b"
	Compare(a, b []byte) int

	// The name of the comparator.  Used to check for comparator
	// mismatches (i.e., a DB created with one comparator is
	// accessed using a different comparator.
	//
	// The client of this package should switch to a new name whenever
	// the comparator implementation changes in a way that will cause
	// the relative ordering of any two keys to change.
	//
	// Names starting with "leveldb." are reserved and should not be used
	// by any clients of this package.
	Name() string

	// Separator appends a sequence of bytes x to dst such that a <= x && x < b,
	// where 'less than' is consistent with Compare. An implementation should
	// return nil if x equal to a.
	//
	// Either contents of a or b should not by any means modified. Doing so
	// may cause corruption on the internal state.
	Separator(dst, a, b []byte) []byte

	// Successor appends a sequence of bytes x to dst such that x >= b, where
	// 'less than' is consistent with Compare. An implementation should return
	// nil if x equal to b.
	//
	// Contents of b should not by any means modified. Doing so may cause
	// corruption on the internal state.
	Successor(dst, b []byte) []byte
}
