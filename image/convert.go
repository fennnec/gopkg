// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"

	color_ext "github.com/chai2010/gopkg/image/color"
)

func Convert(m image.Image, isColor bool) image.Image {
	if isColor {
		return convertToColor(m)
	} else {
		return convertToGray(m)
	}
}

func convertToColor(m image.Image) image.Image {
	switch m := m.(type) {
	case *image.Gray:
		b := m.Bounds()
		rgb := NewRGB(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				rgb.SetRGB(x, y, color_ext.RGB{
					R: v.Y,
					G: v.Y,
					B: v.Y,
				})
			}
		}
		return rgb
	case *image.Gray16:
		b := m.Bounds()
		rgb48 := NewRGB48(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				rgb48.SetRGB48(x, y, color_ext.RGB48{
					R: v.Y,
					G: v.Y,
					B: v.Y,
				})
			}
		}
		return rgb48
	case *Gray32f:
		b := m.Bounds()
		rgb96f := NewRGB96f(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray32fAt(x, y)
				rgb96f.SetRGB96f(x, y, color_ext.RGB96f{
					R: v.Y,
					G: v.Y,
					B: v.Y,
				})
			}
		}
		return rgb96f
	case *image.YCbCr:
		b := m.Bounds()
		rgb := NewRGB(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.YCbCrAt(x, y)
				rr, gg, bb := color.YCbCrToRGB(v.Y, v.Cb, v.Cr)
				rgb.SetRGB(x, y, color_ext.RGB{
					R: rr,
					G: gg,
					B: bb,
				})
			}
		}
		return rgb
	case *image.Paletted:
		switch m.Palette[0].(type) {
		case color.Gray:
			b := m.Bounds()
			rgb := NewRGB(b)
			for y := b.Min.Y; y < b.Max.Y; y++ {
				for x := b.Min.X; x < b.Max.X; x++ {
					v := m.At(x, y).(color.Gray)
					rgb.SetRGB(x, y, color_ext.RGB{
						R: v.Y,
						G: v.Y,
						B: v.Y,
					})
				}
			}
			return rgb
		case color.Gray16:
			b := m.Bounds()
			rgb48 := NewRGB48(b)
			for y := b.Min.Y; y < b.Max.Y; y++ {
				for x := b.Min.X; x < b.Max.X; x++ {
					v := m.At(x, y).(color.Gray16)
					rgb48.SetRGB48(x, y, color_ext.RGB48{
						R: v.Y,
						G: v.Y,
						B: v.Y,
					})
				}
			}
			return rgb48
		case color_ext.Gray32f:
			b := m.Bounds()
			rgb96f := NewRGB96f(b)
			for y := b.Min.Y; y < b.Max.Y; y++ {
				for x := b.Min.X; x < b.Max.X; x++ {
					v := m.At(x, y).(color_ext.Gray32f)
					rgb96f.SetRGB96f(x, y, color_ext.RGB96f{
						R: v.Y,
						G: v.Y,
						B: v.Y,
					})
				}
			}
			return rgb96f
		case color.YCbCr:
			b := m.Bounds()
			rgb := NewRGB(b)
			for y := b.Min.Y; y < b.Max.Y; y++ {
				for x := b.Min.X; x < b.Max.X; x++ {
					v := m.At(x, y).(color.YCbCr)
					rr, gg, bb := color.YCbCrToRGB(v.Y, v.Cb, v.Cr)
					rgb.SetRGB(x, y, color_ext.RGB{
						R: rr,
						G: gg,
						B: bb,
					})
				}
			}
			return rgb
		}
	}
	return m
}

func convertToGray(m image.Image) image.Image {
	switch m := m.(type) {
	case *RGB:
		b := m.Bounds()
		gray := image.NewGray(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray.SetGray(x, y,
					color.GrayModel.Convert(m.RGBAt(x, y)).(color.Gray),
				)
			}
		}
		return gray
	case *RGB48:
		b := m.Bounds()
		gray16 := image.NewGray16(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray16.SetGray16(x, y,
					color.Gray16Model.Convert(m.RGB48At(x, y)).(color.Gray16),
				)
			}
		}
		return gray16
	case *RGB96f:
		b := m.Bounds()
		gray32f := NewGray32f(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray32f.SetGray32f(x, y,
					color_ext.Gray32fModel.Convert(m.RGB96fAt(x, y)).(color_ext.Gray32f),
				)
			}
		}
		return gray32f
	case *image.RGBA:
		b := m.Bounds()
		gray := image.NewGray(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray.SetGray(x, y,
					color.RGBAModel.Convert(m.RGBAAt(x, y)).(color.Gray),
				)
			}
		}
		return gray
	case *image.RGBA64:
		b := m.Bounds()
		gray16 := image.NewGray16(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray16.SetGray16(x, y,
					color.Gray16Model.Convert(m.RGBA64At(x, y)).(color.Gray16),
				)
			}
		}
		return gray16
	case *RGBA128f:
		b := m.Bounds()
		gray32f := NewGray32f(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray32f.SetGray32f(x, y,
					color_ext.Gray32fModel.Convert(m.RGBA128fAt(x, y)).(color_ext.Gray32f),
				)
			}
		}
		return gray32f
	case *image.YCbCr:
		b := m.Bounds()
		gray := image.NewGray(b)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			copy(gray.Pix[y*gray.Stride:][:b.Dx()], m.Y[y*m.YStride:])
		}
		return gray
	case *image.Paletted:
		switch m.Palette[0].(type) {
		case color_ext.RGB, color.RGBA, color.YCbCr:
			b := m.Bounds()
			gray := image.NewGray(b)
			for y := b.Min.Y; y < b.Max.Y; y++ {
				for x := b.Min.X; x < b.Max.X; x++ {
					gray.SetGray(x, y,
						color.GrayModel.Convert(m.At(x, y)).(color.Gray),
					)
				}
			}
			return gray
		case color_ext.RGB48, color.RGBA64:
			b := m.Bounds()
			gray16 := image.NewGray16(b)
			for y := b.Min.Y; y < b.Max.Y; y++ {
				for x := b.Min.X; x < b.Max.X; x++ {
					gray16.SetGray16(x, y,
						color.Gray16Model.Convert(m.At(x, y)).(color.Gray16),
					)
				}
			}
			return gray16
		case color_ext.RGB96f, color_ext.RGBA128f:
			b := m.Bounds()
			gray32f := NewGray32f(b)
			for y := b.Min.Y; y < b.Max.Y; y++ {
				for x := b.Min.X; x < b.Max.X; x++ {
					gray32f.SetGray32f(x, y,
						color_ext.Gray32fModel.Convert(m.At(x, y)).(color_ext.Gray32f),
					)
				}
			}
			return gray32f
		}
	}
	return m
}
