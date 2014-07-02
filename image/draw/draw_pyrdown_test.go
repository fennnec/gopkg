// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	image_ext "github.com/chai2010/gopkg/image"
	color_ext "github.com/chai2010/gopkg/image/color"
)

type tDrawPyrDownTester struct {
	BgdImage draw.Image
	BgdColor color.Color
	FgdImage draw.Image
	FgdColor color.Color
	DrawRect image.Rectangle
	DrawSp   image.Point
	FgdRect  image.Rectangle
}

func TestDrawPyrDown_Average(t *testing.T) {
	for i, v := range tDrawPyrDownTesterList {
		tClearImage(v.BgdImage, v.BgdColor)
		tClearImage(v.FgdImage, v.FgdColor)
		DrawPyrDown(v.BgdImage, v.DrawRect, v.FgdImage, v.DrawSp, Filter_Average)
		err := tCheckImageColor(v.BgdImage, v.FgdRect, v.FgdColor, v.BgdColor)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
	}
}

func TestDrawPyrDown_Interlace(t *testing.T) {
	for i, v := range tDrawPyrDownTesterList {
		tClearImage(v.BgdImage, v.BgdColor)
		tClearImage(v.FgdImage, v.FgdColor)
		DrawPyrDown(v.BgdImage, v.DrawRect, v.FgdImage, v.DrawSp, Filter_Interlace)
		err := tCheckImageColor(v.BgdImage, v.FgdRect, v.FgdColor, v.BgdColor)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
	}
}

var tDrawPyrDownTesterList = []tDrawPyrDownTester{
	// Gray
	tDrawPyrDownTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(4, 4),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 8, 8)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(6, 6),           // +overflow
		FgdRect:  image.Rect(0, 0, 1, 1),
	},

	// Gray16
	tDrawPyrDownTester{
		BgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray16{100 << 8},
		FgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray16{250 << 8},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray16{100 << 8},
		FgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray16{250 << 8},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray16{100 << 8},
		FgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray16{250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray16{100 << 8},
		FgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray16{250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(4, 4),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewGray16(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray16{100 << 8},
		FgdImage: image.NewGray16(image.Rect(0, 0, 8, 8)),
		FgdColor: color.Gray16{250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(6, 6),           // +overflow
		FgdRect:  image.Rect(0, 0, 1, 1),
	},

	// Gray32f
	tDrawPyrDownTester{
		BgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.Gray32f{100 << 8},
		FgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.Gray32f{250 << 8},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.Gray32f{100 << 8},
		FgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.Gray32f{250 << 8},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.Gray32f{100 << 8},
		FgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.Gray32f{250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.Gray32f{100 << 8},
		FgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.Gray32f{250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(4, 4),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewGray32f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.Gray32f{100 << 8},
		FgdImage: image_ext.NewGray32f(image.Rect(0, 0, 8, 8)),
		FgdColor: color_ext.Gray32f{250 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(6, 6),           // +overflow
		FgdRect:  image.Rect(0, 0, 1, 1),
	},

	// RGBA, drawPyrDownRGBA_Average_fast only has 7bit for uint8
	tDrawPyrDownTester{
		BgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA{100, 101, 102, 103},
		FgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA{150, 152, 154, 156},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA{100, 101, 102, 103},
		FgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA{150, 152, 154, 156},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA{100, 101, 102, 103},
		FgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA{150, 152, 154, 156},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA{100, 101, 102, 103},
		FgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA{150, 152, 154, 156},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(4, 4),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewRGBA(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA{100, 101, 102, 103},
		FgdImage: image.NewRGBA(image.Rect(0, 0, 8, 8)),
		FgdColor: color.RGBA{150, 152, 154, 156},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(6, 6),           // +overflow
		FgdRect:  image.Rect(0, 0, 1, 1),
	},

	// RGBA64
	tDrawPyrDownTester{
		BgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA64{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA64{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA64{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA64{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA64{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA64{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA64{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		FgdColor: color.RGBA64{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(4, 4),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},
	tDrawPyrDownTester{
		BgdImage: image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		BgdColor: color.RGBA64{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image.NewRGBA64(image.Rect(0, 0, 8, 8)),
		FgdColor: color.RGBA64{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(6, 6),           // +overflow
		FgdRect:  image.Rect(0, 0, 1, 1),
	},

	// RGBA128f
	tDrawPyrDownTester{
		BgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.RGBA128f{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.RGBA128f{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.RGBA128f{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.RGBA128f{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.RGBA128f{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.RGBA128f{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.RGBA128f{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		FgdColor: color_ext.RGBA128f{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(4, 4),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		BgdColor: color_ext.RGBA128f{100 << 8, 101 << 8, 102 << 8, 103 << 8},
		FgdImage: image_ext.NewRGBA128f(image.Rect(0, 0, 8, 8)),
		FgdColor: color_ext.RGBA128f{250 << 8, 251 << 8, 252 << 8, 253 << 8},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(6, 6),           // +overflow
		FgdRect:  image.Rect(0, 0, 1, 1),
	},

	// YCbCr
	tDrawPyrDownTester{
		BgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		BgdColor: color.YCbCr{100, 101, 102},
		FgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		FgdColor: color.YCbCr{150, 152, 154},
		DrawRect: image.Rect(0, 0, 5, 5),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		BgdColor: color.YCbCr{100, 101, 102},
		FgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		FgdColor: color.YCbCr{150, 152, 154},
		DrawRect: image.Rect(0, 0, 10, 10),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		BgdColor: color.YCbCr{100, 101, 102},
		FgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		FgdColor: color.YCbCr{150, 152, 154},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 5, 5),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		BgdColor: color.YCbCr{100, 101, 102},
		FgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		FgdColor: color.YCbCr{150, 152, 154},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(4, 4),           // +overflow
		FgdRect:  image.Rect(0, 0, 3, 3),
	},
	tDrawPyrDownTester{
		BgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		BgdColor: color.YCbCr{100, 101, 102},
		FgdImage: image_ext.NewYCbCr(image.Rect(0, 0, 8, 8), image.YCbCrSubsampleRatio444),
		FgdColor: color.YCbCr{150, 152, 154},
		DrawRect: image.Rect(0, 0, 15, 15), // +overflow
		DrawSp:   image.Pt(6, 6),           // +overflow
		FgdRect:  image.Rect(0, 0, 1, 1),
	},
}
