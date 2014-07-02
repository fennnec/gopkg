// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

import (
	"image"
	"image/color"
	"testing"
)

func _TestDrawPyrDown_Interlace(t *testing.T) {
	for i, v := range tDrawPyrDownTesterList_Interlace {
		tClearImage(v.BgdImage, v.BgdColor)
		tClearImage(v.FgdImage, v.FgdColor)
		DrawPyrDown(v.BgdImage, v.DrawRect, v.FgdImage, v.DrawSp, Filter_Interlace)
		err := tCheckImageColor(v.BgdImage, v.FgdRect, v.FgdColor, v.BgdColor)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
	}
}

var tDrawPyrDownTesterList_Interlace = []tDrawPyrDownTester{
	// Gray
	tDrawPyrDownTester{
		BgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		BgdColor: color.Gray{100},
		FgdImage: image.NewGray(image.Rect(0, 0, 10, 10)),
		FgdColor: color.Gray{250},
		DrawRect: image.Rect(0, 0, 2, 2),
		DrawSp:   image.Pt(0, 0),
		FgdRect:  image.Rect(0, 0, 2, 2),
	},
}
