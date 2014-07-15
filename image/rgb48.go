// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"

	color_ext "github.com/chai2010/gopkg/image/color"
)

// RGB48 is an in-memory image whose At method returns color.RGB48 values.
type RGB48 struct {
	// Pix holds the image's pixels, in R, G, B order and big-endian format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*6].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *RGB48) ColorModel() color.Model { return color_ext.RGB48Model }

func (p *RGB48) Bounds() image.Rectangle { return p.Rect }

func (p *RGB48) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color_ext.RGB48{}
	}
	i := p.PixOffset(x, y)
	return color_ext.RGB48{
		uint16(p.Pix[i+0])<<8 | uint16(p.Pix[i+1]),
		uint16(p.Pix[i+2])<<8 | uint16(p.Pix[i+3]),
		uint16(p.Pix[i+4])<<8 | uint16(p.Pix[i+5]),
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB48) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*6
}

func (p *RGB48) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color_ext.RGB48Model.Convert(c).(color_ext.RGB48)
	p.Pix[i+0] = uint8(c1.R >> 8)
	p.Pix[i+1] = uint8(c1.R)
	p.Pix[i+2] = uint8(c1.G >> 8)
	p.Pix[i+3] = uint8(c1.G)
	p.Pix[i+4] = uint8(c1.B >> 8)
	p.Pix[i+5] = uint8(c1.B)
}

func (p *RGB48) SetRGB48(x, y int, c color_ext.RGB48) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+0] = uint8(c.R >> 8)
	p.Pix[i+1] = uint8(c.R)
	p.Pix[i+2] = uint8(c.G >> 8)
	p.Pix[i+3] = uint8(c.G)
	p.Pix[i+4] = uint8(c.B >> 8)
	p.Pix[i+5] = uint8(c.B)
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB48) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB48{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB48{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB48) Opaque() bool {
	return true
}

// NewRGB48 returns a new RGB48 with the given bounds.
func NewRGB48(r image.Rectangle) *RGB48 {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 6*w*h)
	return &RGB48{pix, 6 * w, r}
}
