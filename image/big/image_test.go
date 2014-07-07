// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"testing"
)

type tImageTester struct {
	BgdImage    *Image
	BgdColor    color.Color
	FgdImage    draw.Image
	FgdColor    color.Color
	DrawRect    image.Rectangle
	DrawLevel   int
	FgdRects    []image.Rectangle
	FgdLevels   []int
	TilesAcross []int
	TilesDown   []int
}

func TestImage_pyramid(t *testing.T) {
	for i, v := range tDrawTesterList {
		if levels := v.BgdImage.Levels(); levels != len(v.TilesAcross) {
			t.Fatalf("%d: bad levels: %v", i, levels)
		}
		if levels := v.BgdImage.Levels(); levels != len(v.TilesDown) {
			t.Fatalf("%d: bad levels: %v", i, levels)
		}
		for level := 0; level < v.BgdImage.Levels(); level++ {
			if x := v.BgdImage.TilesAcross(level); x != v.TilesAcross[level] {
				t.Fatalf("%d: bad TilesAcross: %v, %v", i, level, x)
			}
			if x := v.BgdImage.TilesDown(level); x != v.TilesDown[level] {
				t.Fatalf("%d: bad TilesDown: %v, %v", i, level, x)
			}
		}

		// for adjustLevel
		if x := v.BgdImage.TilesAcross(-1); x != v.TilesAcross[v.BgdImage.Levels()-1] {
			t.Fatalf("%d: bad TilesAcross: %v", -1, x)
		}
		if x := v.BgdImage.TilesDown(-1); x != v.TilesDown[v.BgdImage.Levels()-1] {
			t.Fatalf("%d: bad TilesDown: %v", -1, x)
		}
		if x := v.BgdImage.TilesAcross(-v.BgdImage.Levels()); x != v.TilesAcross[0] {
			t.Fatalf("%d: bad TilesAcross: %v", -v.BgdImage.Levels(), x)
		}
		if x := v.BgdImage.TilesDown(-v.BgdImage.Levels()); x != v.TilesDown[0] {
			t.Fatalf("%d: bad TilesDown: %v", -v.BgdImage.Levels(), x)
		}
	}
}

func TestImage_readAndWrite(t *testing.T) {
	var err error
	for i, v := range tDrawTesterList {
		tClearImage(v.BgdImage, v.BgdColor)
		tClearImage(v.FgdImage, v.FgdColor)

		err = tCheckImageColor(v.BgdImage, v.BgdImage.Bounds(), v.BgdColor, -1)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}

		v.BgdImage.WriteRect(v.DrawLevel, v.DrawRect, v.FgdImage)
		for j := 0; j < len(v.FgdRects); j++ {
			err = tCheckImageColor(v.BgdImage, v.FgdRects[j], v.FgdColor, v.FgdLevels[j])
			if err != nil {
				t.Fatalf("%d,%d: %v", i, j, err)
			}
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

func tCheckImageColor(m *Image, r image.Rectangle, c color.Color, level int) error {
	level = m.adjustLevel(level)
	m = m.SubLevels(level + 1)
	r0, g0, b0, a0 := c.RGBA()
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			c1 := m.At(x, y)
			r1, g1, b1, a1 := c1.RGBA()
			if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
				return fmt.Errorf("level(%d), pixel(%d, %d): want %v, got %v", level, x, y, c, c1)
			}
		}
	}
	return nil
}

var tDrawTesterList = []tImageTester{
	tImageTester{
		BgdImage:    NewImage(image.Rect(0, 0, 10, 10), image.Pt(4, 4), color.GrayModel),
		TilesAcross: []int{1, 2, 3},
		TilesDown:   []int{1, 2, 3},
		BgdColor:    color.Gray{100},
		FgdImage:    image.NewGray(image.Rect(0, 0, 10, 10)),
		FgdColor:    color.Gray{150},
		DrawRect:    image.Rect(0, 0, 5, 5),
		DrawLevel:   -1,
		FgdRects: []image.Rectangle{
			image.Rect(0, 0, 5, 5), // -1
			image.Rect(0, 0, 2, 2), // -2
			image.Rect(0, 0, 1, 1), // -3
			image.Rect(0, 0, 1, 1), // +0
			image.Rect(0, 0, 2, 2), // +1
			image.Rect(0, 0, 5, 5), // +2
		},
		FgdLevels: []int{
			-1,
			-2,
			-3,
			+0,
			+1,
			+2,
		},
	},
}
