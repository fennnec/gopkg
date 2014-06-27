// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package draw provides image composition functions.
package draw

import (
	"image"
	"image/color"

	image_ext "github.com/chai2010/gopkg/image"
)

// Image is an image.Image with a Set method to change a single pixel.
type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}

// Draw aligns r.Min in dst with sp in src and then replaces the rectangle r in dst with src.
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point) {
	switch dst := dst.(type) {
	case *image.Gray:
		drawGray(dst, r, src, sp)
	case *image.Gray16:
		drawGray16(dst, r, src, sp)
	case *image_ext.Gray32f:
		drawGray32f(dst, r, src, sp)
	case *image.RGBA:
		drawRGBA(dst, r, src, sp)
	case *image.RGBA64:
		drawRGBA64(dst, r, src, sp)
	case *image_ext.RGBA128f:
		drawRGBA128f(dst, r, src, sp)
	default:
		drawImage(dst, r, src, sp)
	}
}

func drawGray(dst *image.Gray, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.Gray:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			off0 := dst.PixOffset(r.Min.X, y)
			off1 := src.PixOffset(sp.X, y-r.Min.Y+sp.Y)
			copy(dst.Pix[off0:][:r.Dx()], src.Pix[off1:])
		}
	default:
		drawImage(dst, r, src, sp)
	}
}

func drawGray16(dst *image.Gray16, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.Gray16:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			off0 := dst.PixOffset(r.Min.X, y)
			off1 := src.PixOffset(sp.X, y-r.Min.Y+sp.Y)
			copy(dst.Pix[off0:][:r.Dx()*2], src.Pix[off1:])
		}
	default:
		drawImage(dst, r, src, sp)
	}
}

func drawGray32f(dst *image_ext.Gray32f, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image_ext.Gray32f:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			off0 := dst.PixOffset(r.Min.X, y)
			off1 := src.PixOffset(sp.X, y-r.Min.Y+sp.Y)
			copy(dst.Pix[off0:][:r.Dx()*4], src.Pix[off1:])
		}
	default:
		drawImage(dst, r, src, sp)
	}
}

func drawRGBA(dst *image.RGBA, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.RGBA:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			off0 := dst.PixOffset(r.Min.X, y)
			off1 := src.PixOffset(sp.X, y-r.Min.Y+sp.Y)
			copy(dst.Pix[off0:][:r.Dx()*4], src.Pix[off1:])
		}
	default:
		drawImage(dst, r, src, sp)
	}
}

func drawRGBA64(dst *image.RGBA64, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image.RGBA64:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			off0 := dst.PixOffset(r.Min.X, y)
			off1 := src.PixOffset(sp.X, y-r.Min.Y+sp.Y)
			copy(dst.Pix[off0:][:r.Dx()*8], src.Pix[off1:])
		}
	default:
		drawImage(dst, r, src, sp)
	}
}

func drawRGBA128f(dst *image_ext.RGBA128f, r image.Rectangle, src image.Image, sp image.Point) {
	switch src := src.(type) {
	case *image_ext.RGBA128f:
		for y := r.Min.Y; y < r.Max.Y; y++ {
			off0 := dst.PixOffset(r.Min.X, y)
			off1 := src.PixOffset(sp.X, y-r.Min.Y+sp.Y)
			copy(dst.Pix[off0:][:r.Dx()*16], src.Pix[off1:])
		}
	default:
		drawImage(dst, r, src, sp)
	}
}

func drawImage(dst Image, r image.Rectangle, src image.Image, sp image.Point) {
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			dst.Set(x, y, src.At(x-r.Min.X+sp.X, y-r.Min.Y+sp.Y))
		}
	}
}
