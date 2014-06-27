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

// RGBA128f is an in-memory image whose At method returns color.RGBA128f values.
type RGBA128f struct {
	// Pix holds the image's pixels. The pixel at (x, y) starts at
	// Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*16].
	Pix []byte
	// Stride is the Pix stride between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *RGBA128f) ColorModel() color.Model { return color_ext.RGBA128fModel }

func (p *RGBA128f) Bounds() image.Rectangle { return p.Rect }

func (p *RGBA128f) At(x, y int) color.Color {
	return p.RGBA128fAt(x, y)
}

func (p *RGBA128f) RGBA128fAt(x, y int) color_ext.RGBA128f {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color_ext.RGBA128f{}
	}
	i := p.PixOffset(x, y)
	return color_ext.RGBA128f{
		R: builtin.Float32(p.Pix[i+0:]),
		G: builtin.Float32(p.Pix[i+4:]),
		B: builtin.Float32(p.Pix[i+8:]),
		A: builtin.Float32(p.Pix[i+16:]),
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA128f) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*1
}

func (p *RGBA128f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color_ext.RGBA128fModel.Convert(c).(color_ext.RGBA128f)
	builtin.PutFloat32(p.Pix[i+0:], c1.R)
	builtin.PutFloat32(p.Pix[i+4:], c1.G)
	builtin.PutFloat32(p.Pix[i+8:], c1.B)
	builtin.PutFloat32(p.Pix[i+16:], c1.A)
}

func (p *RGBA128f) SetRGBA128f(x, y int, c color_ext.RGBA128f) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	builtin.PutFloat32(p.Pix[i+0:], c.R)
	builtin.PutFloat32(p.Pix[i+4:], c.G)
	builtin.PutFloat32(p.Pix[i+8:], c.B)
	builtin.PutFloat32(p.Pix[i+16:], c.A)
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA128f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBA128f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGBA128f{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA128f) Opaque() bool {
	return true
}

// NewRGBA128f returns a new RGBA128f with the given bounds.
func NewRGBA128f(r image.Rectangle) *RGBA128f {
	w, h := r.Dx(), r.Dy()
	pix := make([]byte, w*h*16)
	return &RGBA128f{pix, w*16, r}
}
