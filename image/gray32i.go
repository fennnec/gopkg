// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"

	color_ext "github.com/chai2010/gopkg/image/color"
)

// Gray32i is an in-memory image whose At method returns color.Gray32i values.
type Gray32i struct {
	// Pix holds the image's pixels. The pixel at (x, y) starts at
	// Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)].
	Pix []color_ext.Gray32i
	// Stride is the Pix stride between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *Gray32i) ColorModel() color.Model { return color_ext.Gray32iModel }

func (p *Gray32i) Bounds() image.Rectangle { return p.Rect }

func (p *Gray32i) At(x, y int) color.Color {
	return p.Gray32iAt(x, y)
}

func (p *Gray32i) Gray32iAt(x, y int) color_ext.Gray32i {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color_ext.Gray32i{}
	}
	i := p.PixOffset(x, y)
	return p.Pix[i]
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Gray32i) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x - p.Rect.Min.X)
}

func (p *Gray32i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color_ext.Gray32iModel.Convert(c).(color_ext.Gray32i)
	p.Pix[i] = c1
}

func (p *Gray32i) SetGray16(x, y int, c color_ext.Gray32i) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i] = c
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray32i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Gray32i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &Gray32i{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray32i) Opaque() bool {
	return true
}

// NewGray32i returns a new Gray32i with the given bounds.
func NewGray32i(r image.Rectangle) *Gray32i {
	w, h := r.Dx(), r.Dy()
	pix := make([]color_ext.Gray32i, w*h)
	return &Gray32i{pix, w, r}
}
