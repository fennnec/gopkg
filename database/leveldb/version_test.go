// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if MajorVersion != 1 || MinorVersion != 14 {
		t.Fatal("invalid version")
	}
}
