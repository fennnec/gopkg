// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jxr

import (
	"errors"
	"image"
	"io"
)

// Encode writes the image m to w in JPEG/XR format.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	return errors.New("jxr: Encode, unsupported")
}
