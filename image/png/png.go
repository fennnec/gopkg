// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package png implements a PNG image decoder and encoder.
//
// The PNG specification is at http://www.w3.org/TR/PNG/.
package png

import (
	"image"
	"image/color"
	"image/png"
	"io"

	image_ext "github.com/chai2010/gopkg/image"
)

const pngHeader = "\x89PNG\r\n\x1a\n"

// Options are the encoding and decoding parameters.
type Options struct {
	ColorModel color.Model
}

// DecodeConfig returns the color model and dimensions of a PNG image
// without decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return png.DecodeConfig(r)
}

// Decode reads a PNG image from r and returns it as an image.Image.
// The type of Image returned depends on the PNG contents.
func Decode(r io.Reader, opt *Options) (image.Image, error) {
	return png.Decode(r)
}

// Encode writes the Image m to w in PNG format.
// Any Image may be encoded, but images that are not image.NRGBA
// might be encoded lossily.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	return png.Encode(w, m)
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
		Name:         "png",
		Extensions:   []string{".png"},
		Magics:       []string{pngHeader},
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
