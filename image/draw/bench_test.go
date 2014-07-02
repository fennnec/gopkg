// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

import (
	"image"
	"testing"
)

// ----------------------------------------------------------------------------
// DrawPyrDownGray: 5x5
// ----------------------------------------------------------------------------

func BenchmarkDrawPyrDown_Average_gray_5x5(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 10, 10))
		src = image.NewGray(image.Rect(0, 0, 10, 10))
		r   = image.Rect(0, 0, 5, 5)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownGray_Average_slow_5x5(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 10, 10))
		src = image.NewGray(image.Rect(0, 0, 10, 10))
		r   = image.Rect(0, 0, 5, 5)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownGray_Average_slow(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownGray_Average_fast_5x5(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 10, 10))
		src = image.NewGray(image.Rect(0, 0, 10, 10))
		r   = image.Rect(0, 0, 5, 5)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownGray_Average_fast(dst, r, src, sp)
	}
}

// ----------------------------------------------------------------------------
// DrawPyrDownGray: 16x16
// ----------------------------------------------------------------------------

func BenchmarkDrawPyrDown_Average_gray_16x16(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 32, 32))
		src = image.NewGray(image.Rect(0, 0, 32, 32))
		r   = image.Rect(0, 0, 16, 16)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownGray_Average_slow_16x16(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 32, 32))
		src = image.NewGray(image.Rect(0, 0, 32, 32))
		r   = image.Rect(0, 0, 16, 16)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownGray_Average_slow(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownGray_Average_fast_16x16(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 32, 32))
		src = image.NewGray(image.Rect(0, 0, 32, 32))
		r   = image.Rect(0, 0, 16, 16)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownGray_Average_fast(dst, r, src, sp)
	}
}

// ----------------------------------------------------------------------------
// DrawPyrDownGray: 32x32
// ----------------------------------------------------------------------------

func BenchmarkDrawPyrDown_Average_gray_32x32(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 64, 64))
		src = image.NewGray(image.Rect(0, 0, 64, 64))
		r   = image.Rect(0, 0, 32, 32)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownGray_Average_slow_32x32(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 64, 64))
		src = image.NewGray(image.Rect(0, 0, 64, 64))
		r   = image.Rect(0, 0, 32, 32)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownGray_Average_slow(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownGray_Average_fast_32x32(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 64, 64))
		src = image.NewGray(image.Rect(0, 0, 64, 64))
		r   = image.Rect(0, 0, 32, 32)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownGray_Average_fast(dst, r, src, sp)
	}
}

// ----------------------------------------------------------------------------
// DrawPyrDownGray: 64x64
// ----------------------------------------------------------------------------

func BenchmarkDrawPyrDown_Average_gray_64x64(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 128, 128))
		src = image.NewGray(image.Rect(0, 0, 128, 128))
		r   = image.Rect(0, 0, 64, 64)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownGray_Average_slow_64x64(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 128, 128))
		src = image.NewGray(image.Rect(0, 0, 128, 128))
		r   = image.Rect(0, 0, 64, 64)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownGray_Average_slow(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownGray_Average_fast_64x64(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 128, 128))
		src = image.NewGray(image.Rect(0, 0, 128, 128))
		r   = image.Rect(0, 0, 64, 64)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownGray_Average_fast(dst, r, src, sp)
	}
}

// ----------------------------------------------------------------------------
// DrawPyrDownGray: 128x128
// ----------------------------------------------------------------------------

func BenchmarkDrawPyrDown_Average_gray_128x128(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 256, 256))
		src = image.NewGray(image.Rect(0, 0, 256, 256))
		r   = image.Rect(0, 0, 128, 128)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownGray_Average_slow_128x128(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 128, 128))
		src = image.NewGray(image.Rect(0, 0, 128, 128))
		r   = image.Rect(0, 0, 64, 64)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownGray_Average_slow(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownGray_Average_fast_128x128(b *testing.B) {
	var (
		dst = image.NewGray(image.Rect(0, 0, 128, 128))
		src = image.NewGray(image.Rect(0, 0, 128, 128))
		r   = image.Rect(0, 0, 64, 64)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownGray_Average_fast(dst, r, src, sp)
	}
}

// ----------------------------------------------------------------------------
// DrawPyrDownRGBA: 5x5
// ----------------------------------------------------------------------------

func BenchmarkDrawPyrDown_Average_rgba_5x5(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 10, 10))
		src = image.NewRGBA(image.Rect(0, 0, 10, 10))
		r   = image.Rect(0, 0, 5, 5)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownRGBA_Average_slow_5x5(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 10, 10))
		src = image.NewRGBA(image.Rect(0, 0, 10, 10))
		r   = image.Rect(0, 0, 5, 5)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownRGBA_Average_slow(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownRGBA_Average_fast_5x5(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 10, 10))
		src = image.NewRGBA(image.Rect(0, 0, 10, 10))
		r   = image.Rect(0, 0, 5, 5)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownRGBA_Average_fast(dst, r, src, sp)
	}
}

// ----------------------------------------------------------------------------
// DrawPyrDownRGBA: 16x16
// ----------------------------------------------------------------------------

func BenchmarkDrawPyrDown_Average_rgba_16x16(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 32, 32))
		src = image.NewRGBA(image.Rect(0, 0, 32, 32))
		r   = image.Rect(0, 0, 16, 16)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownRGBA_Average_slow_16x16(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 32, 32))
		src = image.NewRGBA(image.Rect(0, 0, 32, 32))
		r   = image.Rect(0, 0, 16, 16)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownRGBA_Average_slow(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownRGBA_Average_fast_16x16(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 32, 32))
		src = image.NewRGBA(image.Rect(0, 0, 32, 32))
		r   = image.Rect(0, 0, 16, 16)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownRGBA_Average_fast(dst, r, src, sp)
	}
}

// ----------------------------------------------------------------------------
// DrawPyrDownRGBA: 32x32
// ----------------------------------------------------------------------------

func BenchmarkDrawPyrDown_Average_rgba_32x32(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 64, 64))
		src = image.NewRGBA(image.Rect(0, 0, 64, 64))
		r   = image.Rect(0, 0, 32, 32)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownRGBA_Average_slow_32x32(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 64, 64))
		src = image.NewRGBA(image.Rect(0, 0, 64, 64))
		r   = image.Rect(0, 0, 32, 32)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownRGBA_Average_slow(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownRGBA_Average_fast_32x32(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 64, 64))
		src = image.NewRGBA(image.Rect(0, 0, 64, 64))
		r   = image.Rect(0, 0, 32, 32)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownRGBA_Average_fast(dst, r, src, sp)
	}
}

// ----------------------------------------------------------------------------
// DrawPyrDownRGBA: 64x64
// ----------------------------------------------------------------------------

func BenchmarkDrawPyrDown_Average_rgba_64x64(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 128, 128))
		src = image.NewRGBA(image.Rect(0, 0, 128, 128))
		r   = image.Rect(0, 0, 64, 64)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownRGBA_Average_slow_64x64(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 128, 128))
		src = image.NewRGBA(image.Rect(0, 0, 128, 128))
		r   = image.Rect(0, 0, 64, 64)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownRGBA_Average_slow(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownRGBA_Average_fast_64x64(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 128, 128))
		src = image.NewRGBA(image.Rect(0, 0, 128, 128))
		r   = image.Rect(0, 0, 64, 64)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownRGBA_Average_fast(dst, r, src, sp)
	}
}

// ----------------------------------------------------------------------------
// DrawPyrDownRGBA: 128x128
// ----------------------------------------------------------------------------

func BenchmarkDrawPyrDown_Average_rgba_128x128(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 256, 256))
		src = image.NewRGBA(image.Rect(0, 0, 256, 256))
		r   = image.Rect(0, 0, 128, 128)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDown_Average(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownRGBA_Average_slow_128x128(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 256, 256))
		src = image.NewRGBA(image.Rect(0, 0, 256, 256))
		r   = image.Rect(0, 0, 128, 128)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownRGBA_Average_slow(dst, r, src, sp)
	}
}

func BenchmarkDrawPyrDownRGBA_Average_fast_128x128(b *testing.B) {
	var (
		dst = image.NewRGBA(image.Rect(0, 0, 256, 256))
		src = image.NewRGBA(image.Rect(0, 0, 256, 256))
		r   = image.Rect(0, 0, 128, 128)
		sp  = image.Pt(0, 0)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawPyrDownRGBA_Average_fast(dst, r, src, sp)
	}
}

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------
