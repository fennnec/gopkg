// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bmp implements a BMP image decoder and encoder.
//
// The BMP specification is at http://www.digicamsoft.com/bmp/bmp.html.
package bmp

import (
	"image"
	"io"

	"code.google.com/p/go.image/bmp"
	image_ext "github.com/chai2010/gopkg/image"
)

// Decode reads a BMP image from r and returns it as an image.Image.
// Limitation: The file must be 8 or 24 bits per pixel.
func Decode(r io.Reader) (image.Image, error) {
	return bmp.Decode(r)
}

// DecodeConfig returns the color model and dimensions of a BMP image without
// decoding the entire image.
// Limitation: The file must be 8 or 24 bits per pixel.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return bmp.DecodeConfig(r)
}

// Encode writes the image m to w in BMP format.
func Encode(w io.Writer, m image.Image) error {
	return bmp.Encode(w, m)
}

func encode(w io.Writer, m image.Image, opt interface{}) error {
	return Encode(w, m)
}

func init() {
	image_ext.RegisterFormat(
		"bmp", "BM????\x00\x00\x00\x00",
		Decode, DecodeConfig,
		encode,
	)
}
