// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package color implements a basic color library.
package color

import (
	"image/color"
)

// Gray32i represents a int32 grayscale color.
type Gray32i struct {
	Y int32
}

func (c Gray32i) RGBA() (r, g, b, a uint32) {
	y := uint32(i32ToU16(c.Y))
	return y, y, y, 0xffff
}

// Gray32f represents a float32 grayscale color.
type Gray32f struct {
	Y float32
}

func (c Gray32f) RGBA() (r, g, b, a uint32) {
	y := uint32(f32ToU16(c.Y))
	return y, y, y, 0xffff
}

// RGBA128i represents a 64-bit alpha-premultiplied color,
// having int32 for each of red, green, blue and alpha.
type RGBA128i struct {
	R, G, B, A int32
}

func (c RGBA128i) RGBA() (r, g, b, a uint32) {
	r = uint32(i32ToU16(c.R))
	g = uint32(i32ToU16(c.G))
	b = uint32(i32ToU16(c.B))
	a = uint32(i32ToU16(c.A))
	return
}

// RGBA128f represents a 64-bit alpha-premultiplied color,
// having int32 for each of red, green, blue and alpha.
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
	Gray32iModel  color.Model = color.ModelFunc(gray32iModel)
	Gray32fModel  color.Model = color.ModelFunc(gray32fModel)
	RGBA128iModel color.Model = color.ModelFunc(rgba128iModel)
	RGBA128fModel color.Model = color.ModelFunc(rgba128fModel)
)

func gray32iModel(c color.Color) color.Color {
	if _, ok := c.(Gray32i); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := (299*r + 587*g + 114*b + 500) / 1000
	return Gray32i{int32(y)}
}

func gray32fModel(c color.Color) color.Color {
	if _, ok := c.(Gray32f); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := (299*r + 587*g + 114*b + 500) / 1000
	return Gray32f{float32(y)}
}

func rgba128iModel(c color.Color) color.Color {
	if _, ok := c.(RGBA128i); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA128i{int32(r), int32(g), int32(b), int32(a)}
}

func rgba128fModel(c color.Color) color.Color {
	if _, ok := c.(RGBA128f); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA128f{float32(r), float32(g), float32(b), float32(a)}
}
