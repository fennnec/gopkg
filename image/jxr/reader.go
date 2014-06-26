// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jxr

import (
	"errors"
	"fmt"
	"image"
	"io"
	"io/ioutil"

	image_ext "github.com/chai2010/gopkg/image"
)

const (
	leHeader = "II\xBC\x00" // Header for little-endian files.
	beHeader = "MM\xBC\x2A" // Header for big-endian files.

	ifdLen = 12 // Length of an IFD entry in bytes.
)

// Decode reads a JPEG/XR image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	return nil, errors.New("jxr: unsupported JPEG/XR image")
}

// DecodeConfig returns the color model and dimensions of a JPEG/XR image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	width, height, channels, depth, data_type, n, err := jxr_decode(data, nil)
	if err != nil {
		return
	}

	if data_type != jxr_unsigned {
		err = fmt.Errorf("jxr: unsupported data type: %v", data_type)
		return
	}

	config.Width = int(width)
	config.Height = int(height)

	_ = channels
	_ = depth
	_ = data_type
	_ = n

	_ = image_ext.Gray32f{}

	// TODO(chai2010):
	err = errors.New("jxr: unsupported JPEG/XR image")
	return
}

func init() {
	image.RegisterFormat("jxr", "II\xBC\x00", Decode, DecodeConfig)
	image.RegisterFormat("jxr", "II\xBC\x01", Decode, DecodeConfig)
}
