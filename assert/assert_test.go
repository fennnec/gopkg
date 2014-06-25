// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assert

import (
	"testing"
)

func TestLineNumbers(t *testing.T) {
	Equal(t, "foo", "foo", "msg!")
	//Equal(t, "foo", "bar", "this should blow up")
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, "foo", "bar", "msg!")
	//NotEqual(t, "foo", "foo", "this should blow up")
}
