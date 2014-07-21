// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rawp

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"

	"code.google.com/p/snappy-go/snappy"
	image_ext "github.com/chai2010/gopkg/image"
	"github.com/chai2010/gopkg/image/convert"
)

type Options struct {
	ColorModel color.Model
	UseSnappy  bool // 0=disabled, 1=enabled (RawPHeader.Data)
}

func DecodeConfig(r io.Reader) (config image.Config, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	hdr, err := rawpDecodeHeader(data)
	if err != nil {
		return
	}

	model, err := rawpColorModel(hdr)
	if err != nil {
		return
	}

	config = image.Config{model, int(hdr.Width), int(hdr.Height)}
	return
}

func Decode(r io.Reader, opt *Options) (m image.Image, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	hdr, err := rawpDecodeHeader(data)
	if err != nil {
		return
	}

	// new decoder
	decoder, err := rawpPixDecoder(hdr)
	if err != nil {
		return
	}

	// decode snappy
	pix := hdr.Data
	if hdr.UseSnappy != 0 {
		if pix, err = snappy.Decode(nil, hdr.Data); err != nil {
			err = fmt.Errorf("image/rawp: Decode, snappy err: %v", err)
			return
		}
	}

	// decode raw pix
	m, err = decoder.Decode(pix, nil)
	if err != nil {
		return
	}

	// convert color model
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
	image.RegisterFormat("rawp", "RAWP\x1B\xF2\x38\x0A", imageDecode, DecodeConfig)

	image_ext.RegisterFormat(image_ext.Format{
		Name:         "rawp",
		Extensions:   []string{".rawp"},
		Magics:       []string{"RAWP\x1B\xF2\x38\x0A"}, // rawSig + rawpMagic
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
