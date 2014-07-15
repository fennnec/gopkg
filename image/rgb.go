// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"

	color_ext "github.com/chai2010/gopkg/image/color"
)

// RGB is an in-memory image whose At method returns color.RGB values.
type RGB struct {
	// Pix holds the image's pixels, in R, G, B order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *RGB) ColorModel() color.Model { return color_ext.RGBModel }

func (p *RGB) Bounds() image.Rectangle { return p.Rect }

func (p *RGB) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color_ext.RGB{}
	}
	i := p.PixOffset(x, y)
	return color_ext.RGB{p.Pix[i+0], p.Pix[i+1], p.Pix[i+2]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

func (p *RGB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color_ext.RGBModel.Convert(c).(color_ext.RGB)
	p.Pix[i+0] = c1.R
	p.Pix[i+1] = c1.G
	p.Pix[i+2] = c1.B
}

func (p *RGB) SetRGB(x, y int, c color_ext.RGB) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+0] = c.R
	p.Pix[i+1] = c.G
	p.Pix[i+2] = c.B
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB) Opaque() bool {
	return true
}

// NewRGB returns a new RGB with the given bounds.
func NewRGB(r image.Rectangle) *RGB {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, 3*w*h)
	return &RGB{buf, 3 * w, r}
}
