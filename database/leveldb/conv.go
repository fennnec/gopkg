// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import "C"

func boolToUchar(b bool) C.uchar {
	if b {
		return C.uchar(1)
	}
	return C.uchar(0)
}

func ucharToBool(uc C.uchar) bool {
	if uc == C.uchar(0) {
		return false
	}
	return true
}
