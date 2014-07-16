// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/chai2010/gopkg/builtin"
	image_ext "github.com/chai2010/gopkg/image"
	color_ext "github.com/chai2010/gopkg/image/color"
)

type tDrawTester struct {
	BgdImage draw.Image
	BgdColor color.Color
	FgdImage draw.Image
	FgdColor color.Color
	DrawRect image.Rectangle
	DrawSp   image.Point
	FgdRect  image.Rectangle
}

func TestDraw(t *testing.T) {
	for i, v := range tDrawTesterList {
		tClearImage(v.BgdImage, v.BgdColor)
		tClearImage(v.FgdImage, v.FgdColor)
		Draw(v.BgdImage, v.DrawRect, v.FgdImage, v.DrawSp)
		err := tCheckImageColor(v.BgdImage, v.FgdRect, v.FgdColor, v.BgdColor)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
	}
}

func tClearImage(m draw.Image, c color.Color) {
	b := m.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			m.Set(x, y, c)
		}
	}
}

func tCheckImageColor(m draw.Image, fgdRect image.Rectangle, fgdColor, bgdColor color.Color) error {
	b := m.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c0 := builtin.If(image.Pt(x, y).In(fgdRect), fgdColor, bgdColor).(color.Color)
			c1 := m.At(x, y)
			r0, g0, b0, a0 := c0.RGBA()
			r1, g1, b1, a1 := c1.RGBA()
			if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
				return fmt.Errorf("pixel(%d, %d): want %v, got %v", x, y, c0, c1)
			}
		}
	}
	return nil
}

var tDrawTesterList = []tDrawTester{
	// Gray
	tDrawTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 8, 8)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},
	// Gray16
	tDrawTester{
		BgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray16{100 << 8},
		FgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray16{250 << 8},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray16{100 << 8},
		FgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray16{250 << 8},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray16{100 << 8},
		FgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray16{250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray16{100 << 8},
		FgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray16{250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray16{100 << 8},
		FgdImage: image.NewGray16(image.Rect(0, 0, 8, 8)),
		FgdColor: color.Gray16{250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},
	// Gray32f
	tDrawTester{
		BgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.Gray32f{Y: 100 << 8},
		FgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.Gray32f{Y: 250 << 8},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.Gray32f{Y: 100 << 8},
		FgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.Gray32f{Y: 250 << 8},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.Gray32f{Y: 100 << 8},
		FgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.Gray32f{Y: 250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.Gray32f{Y: 100 << 8},
		FgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.Gray32f{Y: 250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.Gray32f{Y: 100 << 8},
		FgdImage: image_ext.NewGray32f(image.Rect(0, 0, 8, 8)),
		FgdColor: color_ext.Gray32f{Y: 250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},

	// RGBA
	tDrawTester{
		BgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA{100, 101, 102, 103},
		FgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA{250, 251, 252, 253},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA{100, 101, 102, 103},
		FgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA{250, 251, 252, 253},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA{100, 101, 102, 103},
		FgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA{250, 251, 252, 253},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA{100, 101, 102, 103},
		FgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA{250, 251, 252, 253},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA{100, 101, 102, 103},
		FgdImage: image.NewRGBA(image.Rect(0, 0, 8, 8)),
		FgdColor: color.RGBA{250, 251, 252, 253},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},

	// RGBA64
	tDrawTester{
		BgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA64{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA64{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA64{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA64{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA64{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA64{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA64{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA64{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA64{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image.NewRGBA64(image.Rect(0, 0, 8, 8)),
		FgdColor: color.RGBA64{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},

	// RGBA128f
	tDrawTester{
		BgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.RGBA128f{R: 100 << 8, G: 101 << 8, B: 102 << 8, A: 103 << 8},
		FgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.RGBA128f{R: 250 << 8, G: 251 << 8, B: 252 << 8, A: 253 << 8},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.RGBA128f{R: 100 << 8, G: 101 << 8, B: 102 << 8, A: 103 << 8},
		FgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.RGBA128f{R: 250 << 8, G: 251 << 8, B: 252 << 8, A: 253 << 8},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.RGBA128f{R: 100 << 8, G: 101 << 8, B: 102 << 8, A: 103 << 8},
		FgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.RGBA128f{R: 250 << 8, G: 251 << 8, B: 252 << 8, A: 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.RGBA128f{R: 100 << 8, G: 101 << 8, B: 102 << 8, A: 103 << 8},
		FgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.RGBA128f{R: 250 << 8, G: 251 << 8, B: 252 << 8, A: 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.RGBA128f{R: 100 << 8, G: 101 << 8, B: 102 << 8, A: 103 << 8},
		FgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 8, 8)),
		FgdColor: color_ext.RGBA128f{R: 250 << 8, G: 251 << 8, B: 252 << 8, A: 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},

	// YCbCr
	tDrawTester{
		BgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		BgdColor: color.YCbCr{100, 101, 102},
		FgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		FgdColor: color.YCbCr{250, 251, 252},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		BgdColor: color.YCbCr{100, 101, 102},
		FgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		FgdColor: color.YCbCr{250, 251, 252},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		BgdColor: color.YCbCr{100, 101, 102},
		FgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		FgdColor: color.YCbCr{250, 251, 252},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 10, 10),
	},
	tDrawTester{
		BgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		BgdColor: color.YCbCr{100, 101, 102},
		FgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		FgdColor: color.YCbCr{250, 251, 252},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawTester{
		BgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		BgdColor: color.YCbCr{100, 101, 102},
		FgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 8, 8), image.YCbCrSubsampleRatio444),
		FgdColor: color.YCbCr{250, 251, 252},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(5, 5),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},
}
