// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jxr

import (
	"errors"
	"image"
	"io"

	"github.com/chai2010/image/convert"
)

// Encode writes the image m to w in JPEG/XR format.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	if opt != nil && opt.ColorModel != nil {
		m = convert.ColorModel(m, opt.ColorModel)
	}
	return errors.New("jxr: Encode, unsupported")
}
