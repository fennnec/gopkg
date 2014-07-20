// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tiff implements a TIFF image decoder and encoder.
//
// The TIFF specification is at http://partners.adobe.com/public/developer/en/tiff/TIFF6.pdf
package tiff

// BUG(chai2010): support Gray32f/RGB/RGB48/RGB96f/RGBA128f.

import (
	"image"
	"image/color"
	"io"

	"code.google.com/p/go.image/tiff"
	image_ext "github.com/chai2010/gopkg/image"
)

const (
	leHeader = "II\x2A\x00" // Header for little-endian files.
	beHeader = "MM\x00\x2A" // Header for big-endian files.
)

// Options are the encoding and decoding parameters.
type Options struct {
	*tiff.Options
	ColorModel color.Model
}

// Decode reads a TIFF image from r and returns it as an image.Image.
// The type of Image returned depends on the contents of the TIFF.
func Decode(r io.Reader, opt interface{}) (image.Image, error) {
	return tiff.Decode(r)
}

// DecodeConfig returns the color model and dimensions of a TIFF image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return tiff.DecodeConfig(r)
}

// Encode writes the image m to w. opt determines the options used for
// encoding, such as the compression type. If opt is nil, an uncompressed
// image is written.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	if opt != nil && opt.Options != nil {
		return tiff.Encode(w, m, opt.Options)
	} else {
		return tiff.Encode(w, m, nil)
	}
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
		Name:         "tiff",
		Extensions:   []string{".tiff", ".tif"},
		Magics:       []string{leHeader, beHeader},
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
