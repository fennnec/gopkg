// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jxr

import (
	"bytes"
	"image"
	"io/ioutil"
	"os"
	"testing"
)

func openImage(filename string) (image.Image, error) {
	f, err := os.Open(testdataDir + filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Decode(f, nil)
}

func TestEncode(t *testing.T) {
	img0, err := openImage("video-001.wdp")
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = Encode(buf, img0, nil)
	if err != nil {
		t.Fatal(err)
	}

	img1, err := Decode(buf, nil)
	if err != nil {
		t.Fatal(err)
	}

	compare(t, img0, img1)
}

// BenchmarkEncode benchmarks the encoding of an image.
func BenchmarkEncode(b *testing.B) {
	img, err := openImage("video-001.wdp")
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
