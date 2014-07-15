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

// RGB96f is an in-memory image whose At method returns color.RGB96f values.
type RGB96f struct {
	// Pix holds the image's pixels. The pixel at (x, y) starts at
	// Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*12].
	Pix []byte
	// Stride is the Pix stride between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *RGB96f) ColorModel() color.Model { return color_ext.RGB96fModel }

func (p *RGB96f) Bounds() image.Rectangle { return p.Rect }

func (p *RGB96f) At(x, y int) color.Color {
	return p.RGB96fAt(x, y)
}

func (p *RGB96f) RGB96fAt(x, y int) color_ext.RGB96f {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color_ext.RGB96f{}
	}
	i := p.PixOffset(x, y)
	return color_ext.RGB96f{
		R: builtin.Float32(p.Pix[i+0:]),
		G: builtin.Float32(p.Pix[i+4:]),
		B: builtin.Float32(p.Pix[i+8:]),
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB96f) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*16
}

func (p *RGB96f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color_ext.RGB96fModel.Convert(c).(color_ext.RGB96f)
	builtin.PutFloat32(p.Pix[i+0:], c1.R)
	builtin.PutFloat32(p.Pix[i+4:], c1.G)
	builtin.PutFloat32(p.Pix[i+8:], c1.B)
}

func (p *RGB96f) SetRGB96f(x, y int, c color_ext.RGB96f) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	builtin.PutFloat32(p.Pix[i+0:], c.R)
	builtin.PutFloat32(p.Pix[i+4:], c.G)
	builtin.PutFloat32(p.Pix[i+8:], c.B)
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB96f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB96f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB96f{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB96f) Opaque() bool {
	return true
}

// NewRGB96f returns a new RGB96f with the given bounds.
func NewRGB96f(r image.Rectangle) *RGB96f {
	w, h := r.Dx(), r.Dy()
	pix := make([]byte, w*h*16)
	return &RGB96f{pix, w * 16, r}
}
