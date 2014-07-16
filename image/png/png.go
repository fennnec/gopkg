// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package png implements a PNG image decoder and encoder.
//
// The PNG specification is at http://www.w3.org/TR/PNG/.
package png

import (
	"image"
	"image/png"
	"io"

	image_ext "github.com/chai2010/gopkg/image"
)

const pngHeader = "\x89PNG\r\n\x1a\n"

// Decode reads a PNG image from r and returns it as an image.Image.
// The type of Image returned depends on the PNG contents.
func Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

// DecodeConfig returns the color model and dimensions of a PNG image
// without decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return png.DecodeConfig(r)
}

// Encode writes the Image m to w in PNG format.
// Any Image may be encoded, but images that are not image.NRGBA
// might be encoded lossily.
func Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

func encode(w io.Writer, m image.Image, opt interface{}) error {
	return Encode(w, m)
}

func init() {
	image_ext.RegisterFormat(
		"png", "BM????\x00\x00\x00\x00",
		Decode, DecodeConfig,
		encode,
	)
}
