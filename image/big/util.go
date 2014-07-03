// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"image"
	"image/color"

	image_ext "github.com/chai2010/gopkg/image"
	color_ext "github.com/chai2010/gopkg/image/color"
)

func makeImageTileMap(r image.Rectangle, tileSize image.Point) (tileMap [][][]image.Image) {
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
	tileMap = make([][][]image.Image, maxInt(xLevels, yLevels))
	for i := 0; i < len(tileMap); i++ {
		xTileSize := tileSize.X << uint8(len(tileMap)-i-1)
		yTileSize := tileSize.Y << uint8(len(tileMap)-i-1)
		xTilesNum := (r.Dx() + xTileSize - 1) / xTileSize
		yTilesNum := (r.Dy() + yTileSize - 1) / yTileSize

		tileMap[i] = make([][]image.Image, xTilesNum)
		for x := 0; x < xTilesNum; x++ {
			tileMap[i][x] = make([]image.Image, yTilesNum)
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

func newImageTile(tileSize image.Point, model color.Model) image.Image {
	return nil
}

func newDemTile(tileSize image.Point, zeroValue color_ext.Gray32f) *image_ext.Gray32f {
	return nil
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
