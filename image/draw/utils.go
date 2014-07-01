// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

func mergeRgbaFast(rgba0, rgba1 uint32) uint32 {
	return (((rgba0 & 0xFEFEFEFE >> 1) + (rgba1 & 0xFEFEFEFE >> 1)) & 0xFEFEFEFE) >> 1
}
