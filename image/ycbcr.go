// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"
)

// YCbCr is an in-memory image of Y'CbCr colors.
type YCbCr image.YCbCr

func (p *YCbCr) ColorModel() color.Model {
	return color.YCbCrModel
}

func (p *YCbCr) Bounds() image.Rectangle {
	return p.Rect
}

func (p *YCbCr) At(x, y int) color.Color {
	return p.YCbCrAt(x, y)
}

func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr {
	if !(image.Point{x, y}.In(p.Rect)) {
		return zeroYCbCr
	}
	yi := p.YOffset(x, y)
	ci := p.COffset(x, y)
	return color.YCbCr{
		p.Y[yi],
		p.Cb[ci],
		p.Cr[ci],
	}
}

// YOffset returns the index of the first element of Y that corresponds to
// the pixel at (x, y).
func (p *YCbCr) YOffset(x, y int) int {
	return ((*image.YCbCr)(p)).YOffset(x, y)
}

// COffset returns the index of the first element of Cb or Cr that corresponds
// to the pixel at (x, y).
func (p *YCbCr) COffset(x, y int) int {
	return ((*image.YCbCr)(p)).COffset(x, y)
}

func (p *YCbCr) Set(x, y int, c color.Color) {
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

func (p *YCbCr) SetYCbCr(x, y int, c color.YCbCr) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	yi := p.YOffset(x, y)
	ci := p.COffset(x, y)
	p.Y[yi] = c.Y
	p.Cb[ci] = c.Cb
	p.Cr[ci] = c.Cr
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *YCbCr) SubImage(r image.Rectangle) image.Image {
	return ((*image.YCbCr)(p)).SubImage(r)
}

func (p *YCbCr) Opaque() bool {
	return ((*image.YCbCr)(p)).Opaque()
}

// NewYCbCr returns a new YCbCr with the given bounds and subsample ratio.
func NewYCbCr(r image.Rectangle, subsampleRatio image.YCbCrSubsampleRatio) *YCbCr {
	p := image.NewYCbCr(r, subsampleRatio)
	for i := 0; i < len(p.Cb); i++ {
		p.Cb[i] = zeroYCbCr.Cb
	}
	for i := 0; i < len(p.Cr); i++ {
		p.Cr[i] = zeroYCbCr.Cr
	}
	return (*YCbCr)(p)
}

var zeroYCbCr = func() (c color.YCbCr) {
	c.Y, c.Cb, c.Cr = color.RGBToYCbCr(0, 0, 0)
	return
}()
