// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raw

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"reflect"

	image_ext "github.com/chai2010/gopkg/image"
	color_ext "github.com/chai2010/gopkg/image/color"
)

func (p *Decoder) DecodeImage(data image.Image) (m draw.Image, err error) {
	// Gray/Gray16/Gray32f
	if p.Channels == 1 && p.DataType == reflect.Uint8 {
		return p.decodeImageGray(data)
	}
	if p.Channels == 1 && p.DataType == reflect.Uint16 {
		return p.decodeImageGray16(data)
	}
	if p.Channels == 1 && p.DataType == reflect.Float32 {
		return p.decodeImageGray32f(data)
	}

	// RGB/RGB48/RGB96f
	if p.Channels == 3 && p.DataType == reflect.Uint8 {
		return p.decodeImageRGB(data)
	}
	if p.Channels == 3 && p.DataType == reflect.Uint16 {
		return p.decodeImageRGB48(data)
	}
	if p.Channels == 3 && p.DataType == reflect.Float32 {
		return p.decodeImageRGB96f(data)
	}

	// RGBA/RGBA64/RGBA128f
	if p.Channels == 4 && p.DataType == reflect.Uint8 {
		return p.decodeImageRGBA(data)
	}
	if p.Channels == 4 && p.DataType == reflect.Uint16 {
		return p.decodeImageRGBA64(data)
	}
	if p.Channels == 4 && p.DataType == reflect.Float32 {
		return p.decodeImageRGBA128f(data)
	}

	// Unknown
	err = fmt.Errorf(
		"image/raw: DecodeImage, unknown image format, channels = %v, dataType = %v",
		p.Channels, p.DataType,
	)
	return
}

func (p *Decoder) decodeImageGray(data image.Image) (m draw.Image, err error) {
	if b := data.Bounds(); b.Dx() != p.Width || b.Dy() != p.Height {
		err = fmt.Errorf("image/raw: bad bounds: %v", data.Bounds())
		return
	}
	if m, ok := data.(*image.Gray); ok {
		return m, nil
	}
	gray := image.NewGray(image.Rect(0, 0, p.Width, p.Height))
	switch data := data.(type) {
	case *image.Gray16:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.Gray16At(x, y)
				gray.SetGray(x, y, color.Gray{uint8(v.Y >> 8)})
			}
		}
	case *image.RGBA:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				gray.SetGray(x, y, color.GrayModel.Convert(data.RGBAAt(x, y)).(color.Gray))
			}
		}
	case *image.RGBA64:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				gray.SetGray(x, y, color.GrayModel.Convert(data.RGBA64At(x, y)).(color.Gray))
			}
		}
	case *image.YCbCr:
		for y := 0; y < p.Height; y++ {
			copy(
				gray.Pix[y*gray.Stride:][:gray.Stride],
				data.Y[y*data.YStride:][:data.YStride],
			)
		}
	default:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				gray.Set(x, y, data.At(x, y))
			}
		}
	}
	m = gray
	return
}

func (p *Decoder) decodeImageGray16(data image.Image) (m draw.Image, err error) {
	if b := data.Bounds(); b.Dx() != p.Width || b.Dy() != p.Height {
		err = fmt.Errorf("image/raw: bad bounds: %v", data.Bounds())
		return
	}
	if m, ok := data.(*image.Gray16); ok {
		return m, nil
	}
	gray16 := image.NewGray16(image.Rect(0, 0, p.Width, p.Height))
	switch data := data.(type) {
	case *image.Gray:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.GrayAt(x, y)
				gray16.SetGray16(x, y, color.Gray16{uint16(v.Y) << 8})
			}
		}
	case *image.RGBA:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				gray16.SetGray16(x, y,
					color.Gray16Model.Convert(data.RGBAAt(x, y)).(color.Gray16),
				)
			}
		}
	case *image.RGBA64:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				gray16.SetGray16(x, y,
					color.Gray16Model.Convert(data.RGBA64At(x, y)).(color.Gray16),
				)
			}
		}
	case *image.YCbCr:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.Y[data.YOffset(x, y)]
				gray16.SetGray16(x, y, color.Gray16{uint16(v) << 8})
			}
		}
	default:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				gray16.Set(x, y, data.At(x, y))
			}
		}
	}
	m = gray16
	return
}

func (p *Decoder) decodeImageGray32f(data image.Image) (m draw.Image, err error) {
	if b := data.Bounds(); b.Dx() != p.Width || b.Dy() != p.Height {
		err = fmt.Errorf("image/raw: bad bounds: %v", data.Bounds())
		return
	}
	if m, ok := data.(*image_ext.Gray32f); ok {
		return m, nil
	}
	gray32f := image_ext.NewGray32f(image.Rect(0, 0, p.Width, p.Height))
	for y := 0; y < p.Height; y++ {
		for x := 0; x < p.Width; x++ {
			gray32f.Set(x, y, data.At(x, y))
		}
	}
	m = gray32f
	return
}

func (p *Decoder) decodeImageRGB(data image.Image) (m draw.Image, err error) {
	if b := data.Bounds(); b.Dx() != p.Width || b.Dy() != p.Height {
		err = fmt.Errorf("image/raw: bad bounds: %v", data.Bounds())
		return
	}
	if m, ok := data.(*image.RGBA); ok {
		return m, nil
	}
	rgba := image.NewRGBA(image.Rect(0, 0, p.Width, p.Height))
	switch data := data.(type) {
	case *image.Gray:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.GrayAt(x, y)
				rgba.SetRGBA(x, y, color.RGBA{
					R: v.Y,
					G: v.Y,
					B: v.Y,
					A: 0xFF,
				})
			}
		}
	case *image.Gray16:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.Gray16At(x, y)
				rgba.SetRGBA(x, y, color.RGBA{
					R: uint8(v.Y >> 8),
					G: uint8(v.Y >> 8),
					B: uint8(v.Y >> 8),
					A: 0xFF,
				})
			}
		}
	default:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				rgba.Set(x, y, data.At(x, y))
			}
		}
	}
	m = rgba
	return
}

func (p *Decoder) decodeImageRGB48(data image.Image) (m draw.Image, err error) {
	if b := data.Bounds(); b.Dx() != p.Width || b.Dy() != p.Height {
		err = fmt.Errorf("image/raw: bad bounds: %v", data.Bounds())
		return
	}
	if m, ok := data.(*image.RGBA64); ok {
		return m, nil
	}
	rgba64 := image.NewRGBA64(image.Rect(0, 0, p.Width, p.Height))
	switch data := data.(type) {
	case *image.Gray:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.GrayAt(x, y)
				rgba64.SetRGBA64(x, y, color.RGBA64{
					R: uint16(v.Y) >> 8,
					G: uint16(v.Y) >> 8,
					B: uint16(v.Y) >> 8,
					A: 0xFFFF,
				})
			}
		}
	case *image.Gray16:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.Gray16At(x, y)
				rgba64.SetRGBA64(x, y, color.RGBA64{
					R: v.Y,
					G: v.Y,
					B: v.Y,
					A: 0xFFFF,
				})
			}
		}
	default:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				rgba64.Set(x, y, data.At(x, y))
			}
		}
	}
	m = rgba64
	return
}

func (p *Decoder) decodeImageRGB96f(data image.Image) (m draw.Image, err error) {
	if b := data.Bounds(); b.Dx() != p.Width || b.Dy() != p.Height {
		err = fmt.Errorf("image/raw: bad bounds: %v", data.Bounds())
		return
	}
	if m, ok := data.(*image_ext.RGBA128f); ok {
		return m, nil
	}
	rgba128f := image_ext.NewRGBA128f(image.Rect(0, 0, p.Width, p.Height))
	switch data := data.(type) {
	case *image.Gray:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.GrayAt(x, y)
				rgba128f.SetRGBA128f(x, y, color_ext.RGBA128f{
					R: float32(uint16(v.Y) >> 8),
					G: float32(uint16(v.Y) >> 8),
					B: float32(uint16(v.Y) >> 8),
					A: 0xFFFF,
				})
			}
		}
	case *image.Gray16:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.Gray16At(x, y)
				rgba128f.SetRGBA128f(x, y, color_ext.RGBA128f{
					R: float32(v.Y),
					G: float32(v.Y),
					B: float32(v.Y),
					A: 0xFFFF,
				})
			}
		}
	default:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				rgba128f.Set(x, y, data.At(x, y))
			}
		}
	}
	m = rgba128f
	return
}

func (p *Decoder) decodeImageRGBA(data image.Image) (m draw.Image, err error) {
	if b := data.Bounds(); b.Dx() != p.Width || b.Dy() != p.Height {
		err = fmt.Errorf("image/raw: bad bounds: %v", data.Bounds())
		return
	}
	if m, ok := data.(*image.RGBA); ok {
		return m, nil
	}
	rgba := image.NewRGBA(image.Rect(0, 0, p.Width, p.Height))
	switch data := data.(type) {
	case *image.Gray:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.GrayAt(x, y)
				rgba.SetRGBA(x, y, color.RGBA{
					R: v.Y,
					G: v.Y,
					B: v.Y,
					A: 0xFF,
				})
			}
		}
	case *image.Gray16:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.Gray16At(x, y)
				rgba.SetRGBA(x, y, color.RGBA{
					R: uint8(v.Y >> 8),
					G: uint8(v.Y >> 8),
					B: uint8(v.Y >> 8),
					A: 0xFF,
				})
			}
		}
	default:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				rgba.Set(x, y, data.At(x, y))
			}
		}
	}
	m = rgba
	return
}

func (p *Decoder) decodeImageRGBA64(data image.Image) (m draw.Image, err error) {
	if b := data.Bounds(); b.Dx() != p.Width || b.Dy() != p.Height {
		err = fmt.Errorf("image/raw: bad bounds: %v", data.Bounds())
		return
	}
	if m, ok := data.(*image.RGBA64); ok {
		return m, nil
	}
	rgba64 := image.NewRGBA64(image.Rect(0, 0, p.Width, p.Height))
	switch data := data.(type) {
	case *image.Gray:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.GrayAt(x, y)
				rgba64.SetRGBA64(x, y, color.RGBA64{
					R: uint16(v.Y) >> 8,
					G: uint16(v.Y) >> 8,
					B: uint16(v.Y) >> 8,
					A: 0xFFFF,
				})
			}
		}
	case *image.Gray16:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.Gray16At(x, y)
				rgba64.SetRGBA64(x, y, color.RGBA64{
					R: v.Y,
					G: v.Y,
					B: v.Y,
					A: 0xFFFF,
				})
			}
		}
	case *image.RGBA:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.RGBAAt(x, y)
				rgba64.SetRGBA64(x, y, color.RGBA64{
					R: uint16(v.R) >> 8,
					G: uint16(v.G) >> 8,
					B: uint16(v.B) >> 8,
					A: uint16(v.A) >> 8,
				})
			}
		}
	default:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				rgba64.Set(x, y, data.At(x, y))
			}
		}
	}
	m = rgba64
	return
}

func (p *Decoder) decodeImageRGBA128f(data image.Image) (m draw.Image, err error) {
	if b := data.Bounds(); b.Dx() != p.Width || b.Dy() != p.Height {
		err = fmt.Errorf("image/raw: bad bounds: %v", data.Bounds())
		return
	}
	if m, ok := data.(*image_ext.RGBA128f); ok {
		return m, nil
	}
	rgba128f := image_ext.NewRGBA128f(image.Rect(0, 0, p.Width, p.Height))
	switch data := data.(type) {
	case *image.Gray:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.GrayAt(x, y)
				rgba128f.SetRGBA128f(x, y, color_ext.RGBA128f{
					R: float32(uint16(v.Y) >> 8),
					G: float32(uint16(v.Y) >> 8),
					B: float32(uint16(v.Y) >> 8),
					A: 0xFFFF,
				})
			}
		}
	case *image.Gray16:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.Gray16At(x, y)
				rgba128f.SetRGBA128f(x, y, color_ext.RGBA128f{
					R: float32(v.Y),
					G: float32(v.Y),
					B: float32(v.Y),
					A: 0xFFFF,
				})
			}
		}
	case *image.RGBA:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				v := data.RGBAAt(x, y)
				rgba128f.SetRGBA128f(x, y, color_ext.RGBA128f{
					R: float32(uint16(v.R) >> 8),
					G: float32(uint16(v.G) >> 8),
					B: float32(uint16(v.B) >> 8),
					A: float32(uint16(v.A) >> 8),
				})
			}
		}
	default:
		for y := 0; y < p.Height; y++ {
			for x := 0; x < p.Width; x++ {
				rgba128f.Set(x, y, data.At(x, y))
			}
		}
	}
	m = rgba128f
	return
}
