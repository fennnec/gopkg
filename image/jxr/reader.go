// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jxr

import "C"
import (
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"

	"github.com/chai2010/builtin"
	image_ext "github.com/chai2010/image"
	color_ext "github.com/chai2010/image/image_color"
	"github.com/chai2010/image/convert"
)

const (
	leHeader = "II\xBC\x00" // Header for little-endian files.
	beHeader = "MM\xBC\x2A" // Header for big-endian files.

	ifdLen = 12 // Length of an IFD entry in bytes.
)

func decodeConfig(data []byte) (config image.Config, err error) {
	width, height, channels, depth, data_type, err := jxr_decode_config(data)
	if err != nil {
		return
	}

	if data_type != jxr_unsigned {
		err = fmt.Errorf("jxr: unsupported data type: %v", data_type)
		return
	}

	config.Width = int(width)
	config.Height = int(height)

	switch data_type {
	case jxr_unsigned:
		switch {
		case channels == 1 && depth == 8:
			config.ColorModel = color.GrayModel
		case channels == 1 && depth == 16:
			config.ColorModel = color.Gray16Model
		case channels == 3 && depth == 8:
			config.ColorModel = color.RGBAModel
		case channels == 3 && depth == 16:
			config.ColorModel = color.RGBA64Model
		case channels == 4 && depth == 8:
			config.ColorModel = color.RGBAModel
		case channels == 4 && depth == 16:
			config.ColorModel = color.RGBA64Model
		}
	case jxr_float:
		switch {
		case channels == 1 && depth == 32:
			config.ColorModel = color_ext.Gray32fModel
		case channels == 3 && depth == 32:
			config.ColorModel = color_ext.RGBA128fModel
		case channels == 4 && depth == 32:
			config.ColorModel = color_ext.RGBA128fModel
		}
	}
	if config.ColorModel == nil {
		err = fmt.Errorf("jxr: unsupported data type: %v", data_type)
		return
	}
	return
}

// Options are the encoding and decoding parameters.
type Options struct {
	ColorModel color.Model
}

// DecodeConfig returns the color model and dimensions of a JPEG/XR image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return decodeConfig(data)
}

// Decode reads a JPEG/XR image from r and returns it as an image.Image.
func Decode(r io.Reader, opt *Options) (m image.Image, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	var config image.Config
	if config, err = decodeConfig(data); err != nil {
		return
	}

	var channels C.int
	switch config.ColorModel {
	case color.GrayModel:
		gray := image.NewGray(image.Rect(0, 0, config.Width, config.Height))
		if _, _, _, _, _, err = jxr_decode(data, gray.Pix, gray.Stride); err != nil {
			return
		}
		m = gray
	case color.Gray16Model:
		gray16 := image.NewGray16(image.Rect(0, 0, config.Width, config.Height))
		if _, _, _, _, _, err = jxr_decode(data, gray16.Pix, gray16.Stride); err != nil {
			return
		}
		m = gray16
	case color_ext.Gray32fModel:
		gray32f := image_ext.NewGray32f(image.Rect(0, 0, config.Width, config.Height))
		if _, _, _, _, _, err = jxr_decode(data, gray32f.Pix, gray32f.Stride); err != nil {
			return
		}
		m = gray32f
	case color.RGBAModel:
		rgba := image.NewRGBA(image.Rect(0, 0, config.Width, config.Height))
		if _, _, channels, _, _, err = jxr_decode(data, rgba.Pix, rgba.Stride); err != nil {
			return
		}
		if channels == 3 {
			b := rgba.Bounds()
			for y := 0; y < b.Max.Y; y++ {
				d := rgba.Pix[y*rgba.Stride:]
				for x := b.Max.X - 1; x >= 0; x-- {
					copy(d[x*4:][:3], d[x*3:])
					d[x*4+3] = 0xff
				}
			}
		}
		m = rgba
	case color.RGBA64Model:
		rgba64 := image.NewRGBA64(image.Rect(0, 0, config.Width, config.Height))
		if _, _, channels, _, _, err = jxr_decode(data, rgba64.Pix, rgba64.Stride); err != nil {
			return
		}
		if channels == 3 {
			b := rgba64.Bounds()
			for y := 0; y < b.Max.Y; y++ {
				d := rgba64.Pix[y*rgba64.Stride:]
				for x := b.Max.X - 1; x >= 0; x-- {
					copy(d[x*8:][:6], d[x*6:])
					d[x*8+7] = 0xff
					d[x*8+6] = 0xff
				}
			}
		}
		m = rgba64
	case color_ext.RGBA128fModel:
		rgba128f := image_ext.NewRGBA128f(image.Rect(0, 0, config.Width, config.Height))
		if _, _, channels, _, _, err = jxr_decode(data, rgba128f.Pix, rgba128f.Stride); err != nil {
			return
		}
		if channels == 3 {
			b := rgba128f.Bounds()
			for y := 0; y < b.Max.Y; y++ {
				d := rgba128f.Pix[y*rgba128f.Stride:]
				for x := b.Max.X - 1; x >= 0; x-- {
					copy(d[x*8:][:6], d[x*6:])
					builtin.PutFloat32(d[x*8+6:], 0xffff)
				}
			}
		}
		m = rgba128f
	}
	if m == nil {
		err = fmt.Errorf("jxr: Decode, unsupported colot model: %T", config.ColorModel)
		return
	}
	if opt != nil && opt.ColorModel != nil {
		m = convert.ColorModel(m, opt.ColorModel)
	}
	return
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
	image.RegisterFormat("jxr", "II\xBC\x00", imageDecode, DecodeConfig)
	image.RegisterFormat("jxr", "II\xBC\x01", imageDecode, DecodeConfig)

	image_ext.RegisterFormat(image_ext.Format{
		Name:         "jxr",
		Extensions:   []string{".jxr", ".wdp"},
		Magics:       []string{"II\xBC\x00", "II\xBC\x01"},
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
