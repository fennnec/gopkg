// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package webp implements a PNG image decoder and encoder.
//
// The PNG specification is at http://www.w3.org/TR/PNG/.
package webp

import (
	"errors"
	"image"
	"io"

	"code.google.com/p/go.image/webp"
	image_ext "github.com/chai2010/gopkg/image"
)

// Decode reads a WEBP image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	return webp.Decode(r)
}

// DecodeConfig returns the color model and dimensions of a WEBP image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return webp.DecodeConfig(r)
}

// Encode writes the image m to w in WEBP format.
func Encode(w io.Writer, m image.Image) error {
	return errors.New("webp: Encode, unsupported")
}

func encode(w io.Writer, m image.Image, opt interface{}) error {
	return Encode(w, m)
}

func init() {
	image_ext.RegisterFormat(
		"webp", "RIFF????WEBPVP8 ",
		Decode, DecodeConfig,
		encode,
	)
}
