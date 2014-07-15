// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

import (
	"image"
	"image/draw"

	image_ext "github.com/chai2010/gopkg/image"
)

func drawPyrDownGray_Interlace(dst *image.Gray, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.Gray:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				dst.SetGray(x, y, src.GrayAt(x0, y0))
			}
		}
	default:
		drawPyrDown_Interlace(dst, r, src, sp)
	}
}

func drawPyrDownGray16_Interlace(dst *image.Gray16, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.Gray16:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				dst.SetGray16(x, y, src.Gray16At(x0, y0))
			}
		}
	default:
		drawPyrDown_Interlace(dst, r, src, sp)
	}
}

func drawPyrDownGray32f_Interlace(dst *image_ext.Gray32f, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image_ext.Gray32f:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				dst.SetGray32f(x, y, src.Gray32fAt(x0, y0))
			}
		}
	default:
		drawPyrDown_Interlace(dst, r, src, sp)
	}
}

func drawPyrDownRGB_Interlace(dst *image_ext.RGB, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image_ext.RGB:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				dst.SetRGB(x, y, src.RGBAt(x0, y0))
			}
		}
	default:
		drawPyrDown_Interlace(dst, r, src, sp)
	}
}

func drawPyrDownRGB48_Interlace(dst *image_ext.RGB48, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image_ext.RGB48:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				dst.SetRGB48(x, y, src.RGB48At(x0, y0))
			}
		}
	default:
		drawPyrDown_Interlace(dst, r, src, sp)
	}
}

func drawPyrDownRGB96f_Interlace(dst *image_ext.RGB96f, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image_ext.RGB96f:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				dst.SetRGB96f(x, y, src.RGB96fAt(x0, y0))
			}
		}
	default:
		drawPyrDown_Interlace(dst, r, src, sp)
	}
}

func drawPyrDownRGBA_Interlace(dst *image.RGBA, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.RGBA:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				dst.SetRGBA(x, y, src.RGBAAt(x0, y0))
			}
		}
	default:
		drawPyrDown_Interlace(dst, r, src, sp)
	}
}

func drawPyrDownRGBA64_Interlace(dst *image.RGBA64, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.RGBA64:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				dst.SetRGBA64(x, y, src.RGBA64At(x0, y0))
			}
		}
	default:
		drawPyrDown_Interlace(dst, r, src, sp)
	}
}

func drawPyrDownRGBA128f_Interlace(dst *image_ext.RGBA128f, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image_ext.RGBA128f:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				dst.SetRGBA128f(x, y, src.RGBA128fAt(x0, y0))
			}
		}
	default:
		drawPyrDown_Interlace(dst, r, src, sp)
	}
}

func drawPyrDownYCbCr_Interlace(dst *yCbCr, r image.Rectangle, src image.Image, sp image.Point) {
	drawPyrDown_Interlace(dst, r, src, sp)
}

func drawPyrDown_Interlace(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			x0 := (x-r.Min.X)*2 + sp.X
			y0 := (y-r.Min.Y)*2 + sp.Y
			dst.Set(x, y, src.At(x0+0, y0+0))
		}
	}
}
