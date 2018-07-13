// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bmp implements a BMP image decoder and encoder.
//
// The BMP specification is at http://www.digicamsoft.com/bmp/bmp.html.
package bmp

import (
	"image"
	"image/color"
	"io"

	"code.google.com/p/go.image/bmp"
	image_ext "github.com/chai2010/image"
	"github.com/chai2010/image/convert"
)

// Options are the encoding and decoding parameters.
type Options struct {
	ColorModel color.Model
}

// DecodeConfig returns the color model and dimensions of a BMP image without
// decoding the entire image.
// Limitation: The file must be 8 or 24 bits per pixel.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return bmp.DecodeConfig(r)
}

// Decode reads a BMP image from r and returns it as an image.Image.
// Limitation: The file must be 8 or 24 bits per pixel.
func Decode(r io.Reader, opt *Options) (m image.Image, err error) {
	if m, err = bmp.Decode(r); err != nil {
		return
	}
	if opt != nil && opt.ColorModel != nil {
		m = convert.ColorModel(m, opt.ColorModel)
	}
	return
}

// Encode writes the image m to w in BMP format.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	if opt != nil && opt.ColorModel != nil {
		m = convert.ColorModel(m, opt.ColorModel)
	}
	return bmp.Encode(w, m)
}

func imageExtDecode(r io.Reader, opt interface{}) (image.Image, error) {
	if opt, ok := opt.(*Options); ok {
		return Decode(r, opt)
	} else {
		return Decode(r, nil)
	}
}

func imageExtEncode(w io.Writer, m image.Image, opt interface{}) error {
	if opt, ok := opt.(*Options); ok {
		return Encode(w, m, opt)
	} else {
		return Encode(w, m, nil)
	}
}

func init() {
	image_ext.RegisterFormat(image_ext.Format{
		Name:         "bmp",
		Extensions:   []string{".bmp"},
		Magics:       []string{"BM????\x00\x00\x00\x00"},
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
