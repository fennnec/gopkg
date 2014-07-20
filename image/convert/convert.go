// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package convert implements some image convert functions.
package convert

import (
	"image"
	"image/color"

	image_ext "github.com/chai2010/gopkg/image"
	color_ext "github.com/chai2010/gopkg/image/color"
)

func Gray(m image.Image) *image.Gray {
	if gray, ok := m.(*image.Gray); ok {
		return gray
	}
	b := m.Bounds()
	gray := image.NewGray(b)
	switch m := m.(type) {
	case *image.Gray16:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				gray.SetGray(x, y, color.Gray{uint8(v.Y >> 8)})
			}
		}
	case *image.RGBA:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray.SetGray(x, y, color.GrayModel.Convert(m.RGBAAt(x, y)).(color.Gray))
			}
		}
	case *image.RGBA64:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray.SetGray(x, y, color.GrayModel.Convert(m.RGBA64At(x, y)).(color.Gray))
			}
		}
	case *image.YCbCr:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			copy(
				gray.Pix[y*gray.Stride:][:gray.Stride],
				m.Y[y*m.YStride:][:m.YStride],
			)
		}
	default:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray.Set(x, y, m.At(x, y))
			}
		}
	}
	return gray
}

func Gray16(m image.Image) *image.Gray16 {
	if gray16, ok := m.(*image.Gray16); ok {
		return gray16
	}
	b := m.Bounds()
	gray16 := image.NewGray16(b)
	switch m := m.(type) {
	case *image.Gray:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				gray16.SetGray16(x, y, color.Gray16{uint16(v.Y) << 8})
			}
		}
	case *image.RGBA:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray16.SetGray16(x, y,
					color.Gray16Model.Convert(m.RGBAAt(x, y)).(color.Gray16),
				)
			}
		}
	case *image.RGBA64:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray16.SetGray16(x, y,
					color.Gray16Model.Convert(m.RGBA64At(x, y)).(color.Gray16),
				)
			}
		}
	case *image.YCbCr:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Y[m.YOffset(x, y)]
				gray16.SetGray16(x, y, color.Gray16{uint16(v) << 8})
			}
		}
	default:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				gray16.Set(x, y, m.At(x, y))
			}
		}
	}
	return gray16
}

func Gray32f(m image.Image) *image_ext.Gray32f {
	if gray32f, ok := m.(*image_ext.Gray32f); ok {
		return gray32f
	}
	b := m.Bounds()
	gray32f := image_ext.NewGray32f(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			gray32f.Set(x, y, m.At(x, y))
		}
	}
	return gray32f
}

func RGB(m image.Image) *image_ext.RGB {
	if rgb, ok := m.(*image_ext.RGB); ok {
		return rgb
	}
	b := m.Bounds()
	rgb := image_ext.NewRGB(b)
	switch m := m.(type) {
	case *image.Gray:
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
	case *image.Gray16:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				rgb.SetRGB(x, y, color_ext.RGB{
					R: uint8(v.Y >> 8),
					G: uint8(v.Y >> 8),
					B: uint8(v.Y >> 8),
				})
			}
		}
	default:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				rgb.Set(x, y, m.At(x, y))
			}
		}
	}
	return rgb
}

func RGB48(m image.Image) *image_ext.RGB48 {
	if rgb48, ok := m.(*image_ext.RGB48); ok {
		return rgb48
	}
	b := m.Bounds()
	rgb48 := image_ext.NewRGB48(b)
	switch m := m.(type) {
	case *image.Gray:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				rgb48.SetRGB48(x, y, color_ext.RGB48{
					R: uint16(v.Y) >> 8,
					G: uint16(v.Y) >> 8,
					B: uint16(v.Y) >> 8,
				})
			}
		}
	case *image.Gray16:
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
	default:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				rgb48.Set(x, y, m.At(x, y))
			}
		}
	}
	return rgb48
}

func RGB96f(m image.Image) *image_ext.RGB96f {
	if rgb96f, ok := m.(*image_ext.RGB96f); ok {
		return rgb96f
	}
	b := m.Bounds()
	rgb96f := image_ext.NewRGB96f(b)
	switch m := m.(type) {
	case *image.Gray:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				rgb96f.SetRGB96f(x, y, color_ext.RGB96f{
					R: float32(uint16(v.Y) >> 8),
					G: float32(uint16(v.Y) >> 8),
					B: float32(uint16(v.Y) >> 8),
				})
			}
		}
	case *image.Gray16:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				rgb96f.SetRGB96f(x, y, color_ext.RGB96f{
					R: float32(v.Y),
					G: float32(v.Y),
					B: float32(v.Y),
				})
			}
		}
	default:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				rgb96f.Set(x, y, m.At(x, y))
			}
		}
	}
	return rgb96f
}

func RGBA(m image.Image) *image.RGBA {
	if rgba, ok := m.(*image.RGBA); ok {
		return rgba
	}
	b := m.Bounds()
	rgba := image.NewRGBA(b)
	switch m := m.(type) {
	case *image.Gray:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				rgba.SetRGBA(x, y, color.RGBA{
					R: v.Y,
					G: v.Y,
					B: v.Y,
					A: 0xFF,
				})
			}
		}
	case *image.Gray16:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				rgba.SetRGBA(x, y, color.RGBA{
					R: uint8(v.Y >> 8),
					G: uint8(v.Y >> 8),
					B: uint8(v.Y >> 8),
					A: 0xFF,
				})
			}
		}
	default:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				rgba.Set(x, y, m.At(x, y))
			}
		}
	}
	return rgba
}

func RGBA64(m image.Image) *image.RGBA64 {
	if rgba64, ok := m.(*image.RGBA64); ok {
		return rgba64
	}
	b := m.Bounds()
	rgba64 := image.NewRGBA64(b)
	switch m := m.(type) {
	case *image.Gray:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				rgba64.SetRGBA64(x, y, color.RGBA64{
					R: uint16(v.Y) >> 8,
					G: uint16(v.Y) >> 8,
					B: uint16(v.Y) >> 8,
					A: 0xFFFF,
				})
			}
		}
	case *image.Gray16:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				rgba64.SetRGBA64(x, y, color.RGBA64{
					R: v.Y,
					G: v.Y,
					B: v.Y,
					A: 0xFFFF,
				})
			}
		}
	case *image.RGBA:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAAt(x, y)
				rgba64.SetRGBA64(x, y, color.RGBA64{
					R: uint16(v.R) >> 8,
					G: uint16(v.G) >> 8,
					B: uint16(v.B) >> 8,
					A: uint16(v.A) >> 8,
				})
			}
		}
	default:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				rgba64.Set(x, y, m.At(x, y))
			}
		}
	}
	return rgba64
}

func RGBA128f(m image.Image) *image_ext.RGBA128f {
	if rgba128f, ok := m.(*image_ext.RGBA128f); ok {
		return rgba128f
	}
	b := m.Bounds()
	rgba128f := image_ext.NewRGBA128f(b)
	switch m := m.(type) {
	case *image.Gray:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				rgba128f.SetRGBA128f(x, y, color_ext.RGBA128f{
					R: float32(uint16(v.Y) >> 8),
					G: float32(uint16(v.Y) >> 8),
					B: float32(uint16(v.Y) >> 8),
					A: 0xFFFF,
				})
			}
		}
	case *image.Gray16:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				rgba128f.SetRGBA128f(x, y, color_ext.RGBA128f{
					R: float32(v.Y),
					G: float32(v.Y),
					B: float32(v.Y),
					A: 0xFFFF,
				})
			}
		}
	case *image.RGBA:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAAt(x, y)
				rgba128f.SetRGBA128f(x, y, color_ext.RGBA128f{
					R: float32(uint16(v.R) >> 8),
					G: float32(uint16(v.G) >> 8),
					B: float32(uint16(v.B) >> 8),
					A: float32(uint16(v.A) >> 8),
				})
			}
		}
	default:
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				rgba128f.Set(x, y, m.At(x, y))
			}
		}
	}
	return rgba128f
}

func Paletted(m image.Image, p color.Palette) *image.Paletted {
	if m, ok := m.(*image.Paletted); ok {
		if len(m.Palette) == len(p) {
			if &m.Palette[0] == &p[0] {
				return m
			}
			isEqual := true
			for i, c := range m.Palette {
				r0, g0, b0, a0 := c.RGBA()
				r1, g1, b1, a1 := p[i].RGBA()
				if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
					isEqual = false
					break
				}
			}
			if isEqual {
				return m
			}
		}
	}
	b := m.Bounds()
	paletted := image.NewPaletted(b, p)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			paletted.Set(x, y, m.At(x, y))
		}
	}
	return paletted
}

func YCbCr(m image.Image, subsampleRatio image.YCbCrSubsampleRatio) *image.YCbCr {
	if m, ok := m.(*image.YCbCr); ok {
		if m.SubsampleRatio == subsampleRatio {
			return m
		}
	}
	b := m.Bounds()
	yCbCr := image_ext.NewYCbCr(b, subsampleRatio)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			yCbCr.Set(x, y, m.At(x, y))
		}
	}
	return (*image.YCbCr)(yCbCr)
}

func Color(m image.Image, isColor bool) image.Image {
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
		rgb := image_ext.NewRGB(b)
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
		rgb48 := image_ext.NewRGB48(b)
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
	case *image_ext.Gray32f:
		b := m.Bounds()
		rgb96f := image_ext.NewRGB96f(b)
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
		rgb := image_ext.NewRGB(b)
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
			rgb := image_ext.NewRGB(b)
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
			rgb48 := image_ext.NewRGB48(b)
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
			rgb96f := image_ext.NewRGB96f(b)
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
			rgb := image_ext.NewRGB(b)
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
	case *image_ext.RGB:
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
	case *image_ext.RGB48:
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
	case *image_ext.RGB96f:
		b := m.Bounds()
		gray32f := image_ext.NewGray32f(b)
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
	case *image_ext.RGBA128f:
		b := m.Bounds()
		gray32f := image_ext.NewGray32f(b)
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
			gray32f := image_ext.NewGray32f(b)
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
