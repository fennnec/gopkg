// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	image_ext "github.com/chai2010/gopkg/image"
	color_ext "github.com/chai2010/gopkg/image/color"
)

func makeImageTileMap(r image.Rectangle, tileSize image.Point) (tileMap [][][]draw.Image) {
	xLevels := 0
	for i := 0; ; i++ {
		if x := (tileSize.X << uint8(i)); x >= r.Dx() {
			xLevels = i + 1
			break
		}
	}
	yLevels := 0
	for i := 0; ; i++ {
		if y := (tileSize.Y << uint8(i)); y >= r.Dy() {
			yLevels = i + 1
			break
		}
	}
	tileMap = make([][][]draw.Image, maxInt(xLevels, yLevels))
	for i := 0; i < len(tileMap); i++ {
		xTileSize := tileSize.X << uint8(len(tileMap)-i-1)
		yTileSize := tileSize.Y << uint8(len(tileMap)-i-1)
		xTilesNum := (r.Dx() + xTileSize - 1) / xTileSize
		yTilesNum := (r.Dy() + yTileSize - 1) / yTileSize

		tileMap[i] = make([][]draw.Image, xTilesNum)
		for x := 0; x < xTilesNum; x++ {
			tileMap[i][x] = make([]draw.Image, yTilesNum)
		}
	}
	return
}

func makeDemTileMap(r image.Rectangle, tileSize image.Point) (tileMap [][][]*image_ext.Gray32f) {
	xLevels := 0
	for i := 0; ; i++ {
		if x := (tileSize.X << uint8(i)); x >= r.Dx() {
			xLevels = i + 1
			break
		}
	}
	yLevels := 0
	for i := 0; ; i++ {
		if y := (tileSize.Y << uint8(i)); y >= r.Dy() {
			xLevels = i + 1
			break
		}
	}
	tileMap = make([][][]*image_ext.Gray32f, maxInt(xLevels, yLevels))
	for i := 0; i < len(tileMap); i++ {
		xTileSize := tileSize.X << uint8(len(tileMap)-i-1)
		yTileSize := tileSize.Y << uint8(len(tileMap)-i-1)
		xTilesNum := (r.Dx() + xTileSize - 1) / xTileSize
		yTilesNum := (r.Dy() + yTileSize - 1) / yTileSize

		tileMap[i] = make([][]*image_ext.Gray32f, xTilesNum)
		for x := 0; x < xTilesNum; x++ {
			tileMap[i][x] = make([]*image_ext.Gray32f, yTilesNum)
		}
	}
	return
}

func newImageTile(tileSize image.Point, model color.Model) draw.Image {
	switch model {
	case color.GrayModel:
		return image.NewGray(image.Rect(0, 0, tileSize.X, tileSize.Y))
	case color.Gray16Model:
		return image.NewGray16(image.Rect(0, 0, tileSize.X, tileSize.Y))
	case color_ext.Gray32fModel:
		return image_ext.NewGray32f(image.Rect(0, 0, tileSize.X, tileSize.Y))
	case color.RGBAModel:
		return image.NewRGBA(image.Rect(0, 0, tileSize.X, tileSize.Y))
	case color.RGBA64Model:
		return image.NewRGBA64(image.Rect(0, 0, tileSize.X, tileSize.Y))
	case color_ext.RGBA128fModel:
		return image_ext.NewRGBA128f(image.Rect(0, 0, tileSize.X, tileSize.Y))
	}
	panic(fmt.Sprintf("image/big: newImageTile, bad color model: %T", model))
}

func newDemTile(tileSize image.Point, zeroValue color_ext.Gray32f) *image_ext.Gray32f {
	m := image_ext.NewGray32f(image.Rect(0, 0, tileSize.X, tileSize.Y))
	if zeroValue.Y != 0 {
		b := m.Bounds()
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				m.SetGray32f(x, y, zeroValue)
			}
		}
	}
	return m
}

func isValidImageColorModel(model color.Model) bool {
	if model == nil {
		return false
	}
	switch model {
	case color.GrayModel:
		return true
	case color.Gray16Model:
		return true
	case color_ext.Gray32fModel:
		return true
	case color.RGBAModel:
		return true
	case color.RGBA64Model:
		return true
	case color_ext.RGBA128fModel:
		return true
	}
	return false
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
