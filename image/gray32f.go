// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"

	"github.com/chai2010/gopkg/builtin"
	color_ext "github.com/chai2010/gopkg/image/color"
)

// Gray32f is an in-memory image whose At method returns color.Gray32f values.
type Gray32f struct {
	// Pix holds the image's pixels. The pixel at (x, y) starts at
	// Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []byte
	// Stride is the Pix stride between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *Gray32f) ColorModel() color.Model { return color_ext.Gray32fModel }

func (p *Gray32f) Bounds() image.Rectangle { return p.Rect }

func (p *Gray32f) At(x, y int) color.Color {
	return p.Gray32fAt(x, y)
}

func (p *Gray32f) Gray32fAt(x, y int) color_ext.Gray32f {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color_ext.Gray32f{}
	}
	v := builtin.Float32(p.Pix[p.PixOffset(x, y):])
	return color_ext.Gray32f{v}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Gray32f) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*4
}

func (p *Gray32f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color_ext.Gray32fModel.Convert(c).(color_ext.Gray32f)
	builtin.PutFloat32(p.Pix[i:], c1.Y)
}

func (p *Gray32f) SetGray32f(x, y int, c color_ext.Gray32f) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	builtin.PutFloat32(p.Pix[i:], c.Y)
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray32f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Gray32f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &Gray32f{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray32f) Opaque() bool {
	return true
}

// NewGray32f returns a new Gray32f with the given bounds.
func NewGray32f(r image.Rectangle) *Gray32f {
	w, h := r.Dx(), r.Dy()
	pix := make([]byte, w*h*4)
	return &Gray32f{pix, w*4, r}
}
