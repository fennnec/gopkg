// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"

	color_ext "github.com/chai2010/gopkg/image/color"
)

// RGBA128i is an in-memory image whose At method returns color.RGBA128i values.
type RGBA128i struct {
	// Pix holds the image's pixels. The pixel at (x, y) starts at
	// Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)].
	Pix []color_ext.RGBA128i
	// Stride is the Pix stride between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *RGBA128i) ColorModel() color.Model { return color_ext.RGBA128iModel }

func (p *RGBA128i) Bounds() image.Rectangle { return p.Rect }

func (p *RGBA128i) At(x, y int) color.Color {
	return p.RGBA128iAt(x, y)
}

func (p *RGBA128i) RGBA128iAt(x, y int) color_ext.RGBA128i {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color_ext.RGBA128i{}
	}
	i := p.PixOffset(x, y)
	return p.Pix[i]
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA128i) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x - p.Rect.Min.X)
}

func (p *RGBA128i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i] = color_ext.RGBA128iModel.Convert(c).(color_ext.RGBA128i)
}

func (p *RGBA128i) SetRGB(x, y int, c color_ext.RGBA128i) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i] = c
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA128i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBA128i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGBA128i{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA128i) Opaque() bool {
	return true
}

// NewRGBA128i returns a new RGBA128i with the given bounds.
func NewRGBA128i(r image.Rectangle) *RGBA128i {
	w, h := r.Dx(), r.Dy()
	pix := make([]color_ext.RGBA128i, 1*w*h)
	return &RGBA128i{pix, 1 * w, r}
}
