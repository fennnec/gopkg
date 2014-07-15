// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package color implements a basic color library.
package color

import (
	"image/color"
)

// Gray32f represents a float32 grayscale color.
type Gray32f struct {
	Y float32
}

func (c Gray32f) RGBA() (r, g, b, a uint32) {
	y := uint32(f32ToU16(c.Y))
	return y, y, y, 0xffff
}

// RGB represents a traditional 24-bit fully opaque color,
// having 8 bits for each of red, green and blue.
type RGB struct {
	R, G, B uint8
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 0xFFFF
	return
}

// RGB48 represents a 48-bit fully opaque color,
// having 16 bits for each of red, green and blue.
type RGB48 struct {
	R, G, B uint16
}

func (c RGB48) RGBA() (r, g, b, a uint32) {
	return uint32(c.R), uint32(c.G), uint32(c.B), 0xFFFF
}

// RGB96f represents a 48-bit fully opaque color,
// having float32 for each of red, green and blue.
type RGB96f struct {
	R, G, B float32
}

func (c RGB96f) RGBA() (r, g, b, a uint32) {
	r = uint32(f32ToU16(c.R))
	g = uint32(f32ToU16(c.G))
	b = uint32(f32ToU16(c.B))
	a = 0xFFFF
	return
}

// RGBA128f represents a 64-bit alpha-premultiplied color,
// having float32 for each of red, green, blue and alpha.
type RGBA128f struct {
	R, G, B, A float32
}

func (c RGBA128f) RGBA() (r, g, b, a uint32) {
	r = uint32(f32ToU16(c.R))
	g = uint32(f32ToU16(c.G))
	b = uint32(f32ToU16(c.B))
	a = uint32(f32ToU16(c.A))
	return
}

// Models for the standard color types.
var (
	Gray32fModel  color.Model = color.ModelFunc(gray32fModel)
	RGBModel      color.Model = color.ModelFunc(rgbModel)
	RGB48Model    color.Model = color.ModelFunc(rgb48Model)
	RGB96fModel   color.Model = color.ModelFunc(rgb96fModel)
	RGBA128fModel color.Model = color.ModelFunc(rgba128fModel)
)

func gray32fModel(c color.Color) color.Color {
	if _, ok := c.(Gray32f); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := (299*r + 587*g + 114*b + 500) / 1000
	return Gray32f{float32(y)}
}

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGB); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
}

func rgb48Model(c color.Color) color.Color {
	if _, ok := c.(RGB48); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB48{uint16(r), uint16(g), uint16(b)}
}

func rgb96fModel(c color.Color) color.Color {
	if _, ok := c.(RGB96f); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB96f{float32(r), float32(g), float32(b)}
}

func rgba128fModel(c color.Color) color.Color {
	if _, ok := c.(RGBA128f); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA128f{float32(r), float32(g), float32(b), float32(a)}
}
