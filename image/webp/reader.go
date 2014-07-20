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

const DefaulQuality = 90

// Options are the encoding parameters.
type Options struct {
	ColorModel color.Model
	Lossless   bool
	Quality    float32 // 0 ~ 100
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

// Decode reads a WEBP image from r and returns it as an image.Image.
func Decode(r io.Reader, opt *Options) (m image.Image, err error) {
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

func imageDecode(r io.Reader) (image.Image, error) {
	return Decode(r, nil)
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
	image.RegisterFormat("webp", "RIFF????WEBPVP8 ", imageDecode, DecodeConfig)

	image_ext.RegisterFormat(image_ext.Format{
		Name:         "webp",
		Extensions:   []string{".webp"},
		Magics:       []string{"RIFF????WEBPVP8 "},
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
