// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jxr

import (
	"errors"
	"io"

	"github.com/chai2010/gopkg.image"
)

const (
	leHeader = "II\xBC\x00" // Header for little-endian files.
	beHeader = "MM\xBC\x2A" // Header for big-endian files.

	ifdLen = 12 // Length of an IFD entry in bytes.
)

// ErrUnsupported means that the input JPEG/XR image uses a valid but unsupported
// feature.
var ErrUnsupported = errors.New("jxr: unsupported JPEG/XR image")

// Decode reads a JPEG/XR image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	return nil, ErrUnsupported
}

// DecodeConfig returns the color model and dimensions of a JPEG/XR image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	err = ErrUnsupported
	return
}

func init() {
	image.RegisterFormat("jxr", "II\xBC\x00", Decode, DecodeConfig)
	image.RegisterFormat("jxr", "II\xBC\x01", Decode, DecodeConfig)
}
