// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"image"
	"image/color"
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func BenchmarkImage_ReadRect(b *testing.B) {
	m := image.NewRGBA(image.Rect(0, 0, 1000, 2000))
	bigImg := NewImage(m.Bounds(), image.Pt(128, 128), color.RGBAModel)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bigImg.ReadRect(-1, m.Bounds(), m)
	}
}

func BenchmarkImage_WriteRect(b *testing.B) {
	m := image.NewRGBA(image.Rect(0, 0, 1000, 2000))
	bigImg := NewImage(m.Bounds(), image.Pt(128, 128), color.RGBAModel)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bigImg.WriteRect(-1, m.Bounds(), m)
	}
}
