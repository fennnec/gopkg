// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

import (
	"image"
	"image/color"

	image_ext "github.com/chai2010/gopkg/image"
	color_ext "github.com/chai2010/gopkg/image/color"
)

func drawPyrDownGray_Average(dst *image.Gray, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.Gray:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				y00 := uint16(src.GrayAt(x0+0, y0+0).Y)
				y01 := uint16(src.GrayAt(x0+0, y0+1).Y)
				y11 := uint16(src.GrayAt(x0+1, y0+1).Y)
				y10 := uint16(src.GrayAt(x0+1, y0+0).Y)

				dst.SetGray(x, y, color.Gray{
					Y: uint8((y00 + y01 + y11 + y10) / 4),
				})
			}
		}
	default:
		drawPyrDown_Average(dst, r, src, sp)
	}
}
func drawPyrDownGray16_Average(dst *image.Gray16, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.Gray16:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				y00 := uint32(src.Gray16At(x0+0, y0+0).Y)
				y01 := uint32(src.Gray16At(x0+0, y0+1).Y)
				y11 := uint32(src.Gray16At(x0+1, y0+1).Y)
				y10 := uint32(src.Gray16At(x0+1, y0+0).Y)

				dst.SetGray16(x, y, color.Gray16{
					Y: uint16((y00 + y01 + y11 + y10) / 4),
				})
			}
		}
	default:
		drawPyrDown_Average(dst, r, src, sp)
	}
}
func drawPyrDownGray32f_Average(dst *image_ext.Gray32f, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image_ext.Gray32f:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				y00 := src.Gray32fAt(x0+0, y0+0).Y
				y01 := src.Gray32fAt(x0+0, y0+1).Y
				y11 := src.Gray32fAt(x0+1, y0+1).Y
				y10 := src.Gray32fAt(x0+1, y0+0).Y

				dst.SetGray32f(x, y, color_ext.Gray32f{
					Y: (y00 + y01 + y11 + y10) / 4,
				})
			}
		}
	default:
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func drawPyrDownRGBA_Average(dst *image.RGBA, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.RGBA:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				rgba00 := src.RGBAAt(x0+0, y0+0)
				rgba01 := src.RGBAAt(x0+0, y0+1)
				rgba11 := src.RGBAAt(x0+1, y0+1)
				rgba10 := src.RGBAAt(x0+1, y0+0)

				dst.SetRGBA(x, y, color.RGBA{
					R: uint8((uint16(rgba00.R) + uint16(rgba01.R) + uint16(rgba11.R) + uint16(rgba10.R)) / 4),
					G: uint8((uint16(rgba00.G) + uint16(rgba01.G) + uint16(rgba11.G) + uint16(rgba10.G)) / 4),
					B: uint8((uint16(rgba00.B) + uint16(rgba01.B) + uint16(rgba11.B) + uint16(rgba10.B)) / 4),
					A: uint8((uint16(rgba00.A) + uint16(rgba01.A) + uint16(rgba11.A) + uint16(rgba10.A)) / 4),
				})
			}
		}
	default:
		drawPyrDown_Average(dst, r, src, sp)
	}
}
func drawPyrDownRGBA64_Average(dst *image.RGBA64, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.RGBA64:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				rgba00 := src.RGBA64At(x0+0, y0+0)
				rgba01 := src.RGBA64At(x0+0, y0+1)
				rgba11 := src.RGBA64At(x0+1, y0+1)
				rgba10 := src.RGBA64At(x0+1, y0+0)

				dst.SetRGBA64(x, y, color.RGBA64{
					R: uint16((uint32(rgba00.R) + uint32(rgba01.R) + uint32(rgba11.R) + uint32(rgba10.R)) / 4),
					G: uint16((uint32(rgba00.G) + uint32(rgba01.G) + uint32(rgba11.G) + uint32(rgba10.G)) / 4),
					B: uint16((uint32(rgba00.B) + uint32(rgba01.B) + uint32(rgba11.B) + uint32(rgba10.B)) / 4),
					A: uint16((uint32(rgba00.A) + uint32(rgba01.A) + uint32(rgba11.A) + uint32(rgba10.A)) / 4),
				})
			}
		}
	default:
		drawPyrDown_Average(dst, r, src, sp)
	}
}
func drawPyrDownRGBA128f_Average(dst *image_ext.RGBA128f, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image_ext.RGBA128f:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				x0 := (x-r.Min.X)*2 + sp.X
				y0 := (y-r.Min.Y)*2 + sp.Y

				rgba00 := src.RGBA128fAt(x0+0, y0+0)
				rgba01 := src.RGBA128fAt(x0+0, y0+1)
				rgba11 := src.RGBA128fAt(x0+1, y0+1)
				rgba10 := src.RGBA128fAt(x0+1, y0+0)

				dst.SetRGBA128f(x, y, color_ext.RGBA128f{
					R: (rgba00.R + rgba01.R + rgba11.R + rgba10.R) / 4,
					G: (rgba00.G + rgba01.G + rgba11.G + rgba10.G) / 4,
					B: (rgba00.B + rgba01.B + rgba11.B + rgba10.B) / 4,
					A: (rgba00.A + rgba01.A + rgba11.A + rgba10.A) / 4,
				})
			}
		}
	default:
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func drawPyrDown_Average(dst Image, r image.Rectangle, src image.Image, sp image.Point) {
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			x0 := (x-r.Min.X)*2 + sp.X
			y0 := (y-r.Min.Y)*2 + sp.Y

			r00, g00, b00, a00 := src.At(x0+0, y0+0).RGBA()
			r01, g01, b01, a01 := src.At(x0+0, y0+1).RGBA()
			r11, g11, b11, a11 := src.At(x0+1, y0+1).RGBA()
			r10, g10, b10, a10 := src.At(x0+1, y0+0).RGBA()

			dst.Set(x, y, color.RGBA64{
				R: uint16((r00 + r01 + r11 + r10) / 4),
				G: uint16((g00 + g01 + g11 + g10) / 4),
				B: uint16((b00 + b01 + b11 + b10) / 4),
				A: uint16((a00 + a01 + a11 + a10) / 4),
			})
		}
	}
}