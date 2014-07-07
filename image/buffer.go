// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/draw"
)

type ImageBuffer interface {
	draw.Image
	SubImage(r image.Rectangle) image.Image
}
