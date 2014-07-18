// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"bytes"
	"io/ioutil"
	"testing"

	image_ext "github.com/chai2010/gopkg/image"
	_ "github.com/chai2010/gopkg/image/png"
)

// 做成列表测试
// 知道压缩的参数

func TestEncode(t *testing.T) {
	img0, _, err := image_ext.Load(testdataDir + "video-001.png")
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = Encode(buf, img0, nil)
	if err != nil {
		t.Fatal(err)
	}

	img1, err := Decode(buf)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the average delta to the tolerance level.
	want := int64(12 << 8)
	if got := averageDelta(img0, img1); got > want {
		t.Fatalf("average delta too high; got %d, want <= %d", got, want)
	}
}

func TestEncodeLossless(t *testing.T) {
	img0, _, err := image_ext.Load(testdataDir + "video-001.png")
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = Encode(buf, img0, &Options{Lossless: true})
	if err != nil {
		t.Fatal(err)
	}

	img1, err := Decode(buf)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the average delta to the tolerance level.
	want := int64(0)
	if got := averageDelta(img0, img1); got > want {
		t.Fatalf("average delta too high; got %d, want <= %d", got, want)
	}
}

// BenchmarkEncode benchmarks the encoding of an image.
func BenchmarkEncode(b *testing.B) {
	img, _, err := image_ext.Load(testdataDir + "video-001.png")
	if err != nil {
		b.Fatal(err)
	}
	s := img.Bounds().Size()
	b.SetBytes(int64(s.X * s.Y * 4))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Encode(ioutil.Discard, img, nil)
	}
}
