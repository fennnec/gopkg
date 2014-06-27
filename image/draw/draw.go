// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package draw provides image composition functions.
package draw

import (
	"fmt"
	"image"
)

type Filter int

const (
	Filter_Average Filter = iota
	Filter_Interlace
)

// Draw aligns r.Min in dst with sp in src and then replaces the rectangle r in dst with src.
func Draw(dst image.Image, r image.Rectangle, src image.Image, sp image.Point) error {
	return fmt.Errorf("image/draw: TODO")
}

// DrawPyrDown aligns r.Min in dst with sp in src and then replaces the rectangle r in dst with downsamples src.
// PyrDown downsamples Blurs an image and downsamples it.
func DrawPyrDown(
	dst image.Image, r image.Rectangle, src image.Image, sp image.Point,
	filter Filter,
) error {
	switch filter {
	case Filter_Average:
		return pyrDown_Average(dst, r, src, sp)
	case Filter_Interlace:
		return pyrDown_Interlace(dst, r, src, sp)
	}
	return fmt.Errorf("image/draw: unsupport filter(%d)", filter)
}

func pyrDown_Average(dst image.Image, r image.Rectangle, src image.Image, sp image.Point) error {
	return fmt.Errorf("image/draw: TODO(chai2010)")
}

func pyrDown_Interlace(dst image.Image, r image.Rectangle, src image.Image, sp image.Point) error {
	return fmt.Errorf("image/draw: TODO(chai2010)")
}
