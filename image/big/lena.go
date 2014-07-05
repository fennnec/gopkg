// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"github.com/chai2010/gopkg/image/big"
)

func loadImage(name string) image.Image {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	m, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func saveImage(name string, m image.Image) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 90})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	lena := loadImage("../testdata/lena.jpg") // 512x512
	bigImg := big.NewImage(image.Rect(0, 0, 512, 512), image.Pt(128, 128), color.RGBAModel)

	if err := bigImg.WriteRect(lena.Bounds(), lena, -1); err != nil {
		log.Fatal(err)
	}

	os.Mkdir("output", 0666)
	saveImage("output/lena.jpg", bigImg)
	for level := 0; level < bigImg.Levels(); level++ {
		sub := bigImg.SubLevels(level + 1)
		m, _ := sub.ReadRect(sub.Bounds(), level)
		saveImage(
			fmt.Sprintf("output/lena-levels-%d-by-ReadRect.jpg", level+1),
			m,
		)
		saveImage(
			fmt.Sprintf("output/lena-levels-%d.jpg", level+1),
			sub,
		)
		for col := 0; col < bigImg.TilesAcross(level); col++ {
			for row := 0; row < bigImg.TilesDown(level); row++ {
				saveImage(
					fmt.Sprintf(
						"output/lena-tile-%01d-%02d-%02d.jpg",
						level, col, row,
					),
					bigImg.GetTile(level, col, row),
				)
			}
		}
	}

	fmt.Println("Done")
}
