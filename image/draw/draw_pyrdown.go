// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

import (
	"image"

	image_ext "github.com/chai2010/gopkg/image"
)

type Filter int

const (
	Filter_Average Filter = iota
	Filter_Interlace
)

// DrawPyrDown aligns r.Min in dst with sp in src and then replaces
// the rectangle r in dst with downsamples src.
func DrawPyrDown(
	dst Image, r image.Rectangle, src image.Image, sp image.Point,
	filter Filter,
) {
	switch filter {
	case Filter_Average:
		switch dst := dst.(type) {
		case *image.Gray:
			drawPyrDownGray_Average(dst, r, src, sp)
		case *image.Gray16:
			drawPyrDownGray16_Average(dst, r, src, sp)
		case *image_ext.Gray32f:
			drawPyrDownGray32f_Average(dst, r, src, sp)
		case *image.RGBA:
			drawPyrDownRGBA_Average(dst, r, src, sp)
		case *image.RGBA64:
			drawPyrDownRGBA64_Average(dst, r, src, sp)
		case *image_ext.RGBA128f:
			drawPyrDownRGBA128f_Average(dst, r, src, sp)
		//case *image.YCbCr:
		//	drawPyrDownYCbCr_Average(&yCbCr{dst}, r, src, sp)
		default:
			drawPyrDown_Average(dst, r, src, sp)
		}
	case Filter_Interlace:
		switch dst := dst.(type) {
		case *image.Gray:
			drawPyrDownGray_Interlace(dst, r, src, sp)
		case *image.Gray16:
			drawPyrDownGray16_Interlace(dst, r, src, sp)
		case *image_ext.Gray32f:
			drawPyrDownGray32f_Interlace(dst, r, src, sp)
		case *image.RGBA:
			drawPyrDownRGBA_Interlace(dst, r, src, sp)
		case *image.RGBA64:
			drawPyrDownRGBA64_Interlace(dst, r, src, sp)
		case *image_ext.RGBA128f:
			drawPyrDownRGBA128f_Interlace(dst, r, src, sp)
		//case *image.YCbCr:
		//	drawPyrDownYCbCr_Interlace(&yCbCr{dst}, r, src, sp)
		default:
			drawPyrDown_Interlace(dst, r, src, sp)
		}
	}
	panic("unreachable")
}
