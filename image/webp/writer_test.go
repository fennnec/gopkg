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

type tTester struct {
	Filename string
	Lossless bool
	Quality  float32 // 0 ~ 100
	MaxDelta int64
}

var tTesterList = []tTester{
	tTester{
		Filename: "video-001.png",
		Lossless: false,
		Quality:  90,
		MaxDelta: 5,
	},
	tTester{
		Filename: "video-001.png",
		Lossless: true,
		Quality:  90,
		MaxDelta: 0,
	},
	tTester{
		Filename: "video-005.gray.png",
		Lossless: false,
		Quality:  90,
		MaxDelta: 5,
	},
	tTester{
		Filename: "video-005.gray.png",
		Lossless: true,
		Quality:  90,
		MaxDelta: 0,
	},
}

func TestEncode(t *testing.T) {
	for i, v := range tTesterList {
		img0, _, err := image_ext.Load(testdataDir + v.Filename)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}

		buf := new(bytes.Buffer)
		err = Encode(buf, img0, &Options{
			Lossless: v.Lossless,
			Quality:  v.Quality,
		})
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}

		img1, err := Decode(buf)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}

		// Compare the average delta to the tolerance level.
		want := int64(v.MaxDelta << 8)
		if got := averageDelta(img0, img1); got > want {
			t.Fatalf("%d: average delta too high; got %d, want <= %d", i, got, want)
		}
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
