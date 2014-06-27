// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raw

import (
	"reflect"
)

func defaultDepthKind(depth int) reflect.Kind {
	switch depth {
	case 8:
		return reflect.Uint8
	case 16:
		return reflect.Uint16
	case 32:
		return reflect.Float32
	}
	return reflect.Uint16
}
