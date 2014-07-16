// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package jpeg implements a JPEG image decoder and encoder.
//
// JPEG is defined in ITU-T T.81: http://www.w3.org/Graphics/JPEG/itu-t81.pdf.
package jpeg

import (
	"image"
	"image/jpeg"
	"io"

	image_ext "github.com/chai2010/gopkg/image"
)

// Decode reads a JPEG image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

// DecodeConfig returns the color model and dimensions of a JPEG image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return jpeg.DecodeConfig(r)
}

// Encode writes the Image m to w in JPEG 4:2:0 baseline format with the given
// options. Default parameters are used if a nil *Options is passed.
func Encode(w io.Writer, m image.Image, o *jpeg.Options) error {
	return jpeg.Encode(w, m, o)
}

func encode(w io.Writer, m image.Image, opt interface{}) error {
	if opt, ok := opt.(*jpeg.Options); ok {
		return Encode(w, m, opt)
	} else {
		return Encode(w, m, nil)
	}
}

func init() {
	image_ext.RegisterFormat(
		"jpeg", "\xff\xd8",
		Decode, DecodeConfig,
		encode,
	)
	image_ext.RegisterFormat(
		"jpg", "\xff\xd8",
		Decode, DecodeConfig,
		encode,
	)
}
