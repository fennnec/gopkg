// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"image/color"
	"io"
	"io/ioutil"

	image_ext "github.com/chai2010/gopkg/image"
	color_ext "github.com/chai2010/gopkg/image/color"
)

// Decode reads a WEBP image from r and returns it as an image.Image.
func Decode(r io.Reader) (m image.Image, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	_, _, hasAlpha, err := GetInfo(data)
	if err != nil {
		return
	}
	if hasAlpha {
		return DecodeRGBA(data)
	} else {
		return DecodeRGB(data)
	}
}

// DecodeConfig returns the color model and dimensions of a WEBP image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	width, height, hasAlpha, err := GetInfo(data)
	if err != nil {
		return
	}
	config.Width = width
	config.Height = height
	if hasAlpha {
		config.ColorModel = color.RGBAModel
	} else {
		config.ColorModel = color_ext.RGBModel
	}
	return
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
