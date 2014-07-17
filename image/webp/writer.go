// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"image/color"
	"io"

	image_ext "github.com/chai2010/gopkg/image"
)

const DefaulQuality = 90

// Options are the encoding parameters.
type Options struct {
	Lossless bool
	Quality  float32 // 0 ~ 100
}

// Encode writes the image m to w in WEBP format.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	if opt != nil && opt.Lossless {
		switch m := adjustImage(m).(type) {
		case *image.Gray:
			return EncodeLosslessGray(w, m)
		case *image_ext.RGB:
			return EncodeLosslessRGB(w, m)
		case *image.RGBA:
			return EncodeLosslessRGBA(w, m)
		}
	} else {
		quality := float32(DefaulQuality)
		if opt != nil {
			quality = opt.Quality
		}
		switch m := adjustImage(m).(type) {
		case *image.Gray:
			return EncodeGray(w, m, quality)
		case *image_ext.RGB:
			return EncodeRGB(w, m, quality)
		case *image.RGBA:
			return EncodeRGBA(w, m, quality)
		}
	}
	panic("image/webp: Encode, unreachable!")
}

func adjustImage(m image.Image) image.Image {
	switch m := m.(type) {
	case *image.Gray, *image_ext.RGB, *image.RGBA:
		return m
	default:
		b := m.Bounds()
		rgba := image.NewRGBA(b)
		dstColorRGBA64 := &color.RGBA64{}
		dstColor := color.Color(dstColorRGBA64)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				pr, pg, pb, pa := m.At(x, y).RGBA()
				dstColorRGBA64.R = uint16(pr)
				dstColorRGBA64.G = uint16(pg)
				dstColorRGBA64.B = uint16(pb)
				dstColorRGBA64.A = uint16(pa)
				rgba.Set(x, y, dstColor)
			}
		}
		return rgba
	}
}
