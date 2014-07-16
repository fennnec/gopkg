// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gif implements a GIF image decoder and encoder.
//
// The GIF specification is at http://www.w3.org/Graphics/GIF/spec-gif89a.txt.
package gif

import (
	"image"
	"image/gif"
	"io"

	image_ext "github.com/chai2010/gopkg/image"
)

// Decode reads a GIF image from r and returns the first embedded
// image as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

// DecodeAll reads a GIF image from r and returns the sequential frames
// and timing information.
func DecodeAll(r io.Reader) (*gif.GIF, error) {
	return gif.DecodeAll(r)
}

// DecodeConfig returns the global color model and dimensions of a GIF image
// without decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return gif.DecodeConfig(r)
}

// EncodeAll writes the images in g to w in GIF format with the
// given loop count and delay between frames.
func EncodeAll(w io.Writer, g *gif.GIF) error {
	return gif.EncodeAll(w, g)
}

// Encode writes the Image m to w in GIF format.
func Encode(w io.Writer, m image.Image, o *gif.Options) error {
	return gif.Encode(w, m, o)
}

func encode(w io.Writer, m image.Image, opt interface{}) error {
	if opt, ok := opt.(*gif.Options); ok {
		return Encode(w, m, opt)
	} else {
		return Encode(w, m, nil)
	}
}

func init() {
	image_ext.RegisterFormat(
		"gif", "GIF8?a",
		Decode, DecodeConfig,
		encode,
	)
}
