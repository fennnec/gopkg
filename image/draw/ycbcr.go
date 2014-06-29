// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

import (
	"image"
	"image/color"
)

type yCbCr struct {
	*image.YCbCr
}

func (p *yCbCr) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	yi := p.YOffset(x, y)
	ci := p.COffset(x, y)
	c1 := color.YCbCrModel.Convert(c).(color.YCbCr)
	p.Y[yi] = c1.Y
	p.Cb[ci] = c1.Cb
	p.Cr[ci] = c1.Cr
}

func (p *yCbCr) SetYCbCr(x, y int, c color.YCbCr) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	yi := p.YOffset(x, y)
	ci := p.COffset(x, y)
	p.Y[yi] = c.Y
	p.Cb[ci] = c.Cb
	p.Cr[ci] = c.Cr
}
