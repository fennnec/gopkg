// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
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

func encode(w io.Writer, m image.Image, opt interface{}) error {
	if opt, ok := opt.(*Options); ok {
		return Encode(w, m, opt)
	} else {
		return Encode(w, m, nil)
	}
}

func init() {
	image_ext.RegisterFormat(
		"webp", "RIFF????WEBPVP8 ",
		Decode, DecodeConfig,
		encode,
	)
}
