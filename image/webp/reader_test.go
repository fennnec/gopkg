// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"testing"

	image_ext "github.com/chai2010/gopkg/image"
	_ "github.com/chai2010/gopkg/image/png"
)

const testdataDir = "../testdata/"

// TestDecode tests that decoding a PNG image and a BMP image result in the
// same pixel data.
func TestDecode(t *testing.T) {
	img0, _, err := image_ext.Load(testdataDir+"video-001.png", nil)
	if err != nil {
		t.Fatal(err)
	}

	img1, _, err := image_ext.Load(testdataDir+"video-001.webp", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the average delta to the tolerance level.
	want := int64(12 << 8)
	if got := averageDelta(img0, img1); got > want {
		t.Fatalf("average delta too high; got %d, want <= %d", got, want)
	}
}

// averageDelta returns the average delta in RGB space. The two images must
// have the same bounds.
func averageDelta(m0, m1 image.Image) int64 {
	b := m0.Bounds()
	var sum, n int64
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c0 := m0.At(x, y)
			c1 := m1.At(x, y)
			r0, g0, b0, _ := c0.RGBA()
			r1, g1, b1, _ := c1.RGBA()
			sum += delta(r0, r1)
			sum += delta(g0, g1)
			sum += delta(b0, b1)
			n += 3
		}
	}
	return sum / n
}

func delta(u0, u1 uint32) int64 {
	d := int64(u0) - int64(u1)
	if d < 0 {
		return -d
	}
	return d
}
