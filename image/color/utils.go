// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"math"
)

func i32ToU16(v int32) uint16 {
	switch {
	case v < 0:
		return 0
	case v > math.MaxUint16:
		return math.MaxUint16
	default:
		return uint16(v)
	}
}

func f32ToU16(v float32) uint16 {
	switch {
	case v < 0:
		return 0
	case v > math.MaxUint16:
		return math.MaxUint16
	default:
		return uint16(v)
	}
}
