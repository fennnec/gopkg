// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	draw_ext "github.com/chai2010/gopkg/image/draw"
)

type Image struct {
	TileMap  [][][]image.Image // m.TileMap[level][col][row]
	TileSize image.Point
	Rect     image.Rectangle
	Model    color.Model
}

func NewImage(r image.Rectangle, tileSize image.Point, model color.Model) *Image {
	if r.Empty() || tileSize.X <= 0 || tileSize.Y <= 0 || model == nil {
		panic(fmt.Sprintf("image/big: NewImage, bad arguments: %v, %v, %v", r, tileSize, model))
	}
	return &Image{
		TileMap:  makeImageTileMap(r, tileSize),
		TileSize: tileSize,
		Rect:     r,
		Model:    model,
	}
}

func (p *Image) ColorModel() color.Model { return p.Model }

func (p *Image) Bounds() image.Rectangle { return p.Rect }

func (p *Image) At(x, y int) color.Color {
	level, col, row := p.Levels()-1, x/p.TileSize.X, y/p.TileSize.Y
	if m := p.TileMap[level][col][row]; m != nil {
		return m.At(x%p.TileSize.X, y%p.TileSize.Y)
	}
	return color.Gray{}
}

func (p *Image) Levels() int {
	return len(p.TileMap)
}

func (p *Image) TilesAcross(level int) int {
	return len(p.TileMap[level])
}

func (p *Image) TilesDown(level int) int {
	return len(p.TileMap[level][0])
}

func (p *Image) GetTile(level, col, row int) (m image.Image, err error) {
	if m = p.TileMap[level][col][row]; m != nil {
		return
	}
	err = fmt.Errorf("image/big: Image.GetTile, not found tile")
	return
}

func (p *Image) SetTile(level, col, row int, m image.Image) (err error) {
	if m.ColorModel() != p.Model {
		err = fmt.Errorf("image/big: Image.SetTile, bad color model: %v", m.ColorModel())
		return
	}
	if m.Bounds() != image.Rect(0, 0, p.TileSize.X, p.TileSize.Y) {
		err = fmt.Errorf("image/big: Image.SetTile, bad bound size: %v", m.Bounds())
		return
	}
	p.TileMap[level][col][row] = m
	return
}

func (p *Image) ReadRect(level, x, y, dx, dy int) (m image.Image, err error) {
	if level < 0 || level >= p.Levels() {
		err = fmt.Errorf(
			"image/big: Image.ReadRect, level = %d, x = %d, y = %d, dx = %d, dy = %d",
			level, x, y, dx, dy,
		)
		return
	}
	if !image.Rect(x, y, x+dx, y+dy).In(p.Bounds()) {
		err = fmt.Errorf(
			"image/big: Image.ReadRect, level = %d, x = %d, y = %d, dx = %d, dy = %d",
			level, x, y, dx, dy,
		)
		return
	}

	tMinX := x / p.TileSize.X
	tMinY := y / p.TileSize.Y
	tMaxX := (x + dx + p.TileSize.X - 1) / p.TileSize.X
	tMaxY := (y + dy + p.TileSize.Y - 1) / p.TileSize.Y

	if max := p.TilesAcross(level); tMaxX > max {
		tMaxX = max
	}
	if max := p.TilesDown(level); tMaxY > max {
		tMaxY = max
	}

	m = newImageTile(p.TileSize, p.Model)
	for col := tMinX; col < tMaxX; col++ {
		for row := tMinY; row < tMaxY; row++ {
			var tile image.Image
			if tile, err = p.GetTile(level, col, row); err != nil {
				return
			}
			p.readRectFromTile(m, tile, x, y, dx, dy, col, row)
		}
	}
	return
}

func (p *Image) readRectFromTile(dst, tile image.Image, x, y, dx, dy, col, row int) {
	bMinX := x
	bMinY := y
	bMaxX := x + dx
	bMaxY := y + dy

	tMinX := col * p.TileSize.X
	tMinY := row * p.TileSize.Y
	tMaxX := tMinX + p.TileSize.X
	tMaxY := tMinY + p.TileSize.Y

	zMinX := maxInt(bMinX, tMinX)
	zMinY := maxInt(bMinY, tMinY)
	zMaxX := minInt(bMaxX, tMaxX)
	zMaxY := minInt(bMaxY, tMaxY)

	if zMinX >= zMaxX || zMinY >= zMaxY {
		return
	}

	draw_ext.Draw(
		dst.(draw.Image), image.Rect(
			zMinX-bMinX,
			zMinX-tMinX,
			zMaxX-zMinX,
			zMaxY-zMinY,
		),
		tile, image.Pt(
			zMinX-tMinX,
			zMinY-tMinY,
		),
	)
	return
}

func (p *Image) WriteRect(level, x, y, dx, dy int, m image.Image) (err error) {
	if level < 0 || level >= p.Levels() {
		err = fmt.Errorf(
			"image/big: Image.WriteRect, level = %d, x = %d, y = %d, dx = %d, dy = %d",
			level, x, y, dx, dy,
		)
		return
	}
	if !image.Rect(x, y, x+dx, y+dy).In(p.Bounds()) {
		err = fmt.Errorf(
			"image/big: Image.WriteRect, level = %d, x = %d, y = %d, dx = %d, dy = %d",
			level, x, y, dx, dy,
		)
		return
	}
	if m.ColorModel() != p.Model {
		err = fmt.Errorf("image/big: Image.WriteRect, bad color model: %v", m.ColorModel())
		return
	}

	tMinX := x / p.TileSize.X
	tMinY := y / p.TileSize.Y
	tMaxX := (x + dx + p.TileSize.X - 1) / p.TileSize.X
	tMaxY := (y + dy + p.TileSize.Y - 1) / p.TileSize.Y

	if max := p.TilesAcross(level); tMaxX > max {
		tMaxX = max
	}
	if max := p.TilesDown(level); tMaxY > max {
		tMaxY = max
	}

	for col := tMinX; col < tMaxX; col++ {
		for row := tMinY; row < tMaxY; row++ {
			var tile image.Image
			if tile, err = p.GetTile(level, col, row); err != nil {
				return
			}
			p.writeRectToTile(m, tile, x, y, dx, dy, col, row)
		}
	}

	err = p.updateRectPyramid(level, x, y, dx, dy)
	return
}

func (p *Image) writeRectToTile(src, tile image.Image, x, y, dx, dy, col, row int) {
	bMinX := x
	bMinY := y
	bMaxX := x + dx
	bMaxY := y + dy

	tMinX := col * p.TileSize.X
	tMinY := row * p.TileSize.Y
	tMaxX := tMinX + p.TileSize.X
	tMaxY := tMinY + p.TileSize.Y

	zMinX := maxInt(bMinX, tMinX)
	zMinY := maxInt(bMinY, tMinY)
	zMaxX := minInt(bMaxX, tMaxX)
	zMaxY := minInt(bMaxY, tMaxY)

	if zMinX >= zMaxX || zMinY >= zMaxY {
		return
	}

	draw_ext.Draw(
		tile.(draw.Image), image.Rect(
			zMinX-tMinY,
			zMinY-tMinY,
			zMaxX-tMinX,
			zMaxY-tMinY,
		),
		src, image.Pt(
			zMinX-bMinX,
			zMinY-bMinY,
		),
	)
	return
}

func (p *Image) updateRectPyramid(level, x, y, dx, dy int) (err error) {
	for level > 0 && dx > 0 && dy > 0 {
		minX, minY := x, y
		maxX, maxY := x+dx-1, y+dy-1

		tMinCol := minX / p.TileSize.X
		tMinRow := minY / p.TileSize.Y
		tMaxCol := maxX / p.TileSize.X
		tMaxRow := maxY / p.TileSize.Y

		for row := tMinRow; row <= tMaxRow; row++ {
			if row >= p.TilesDown(level) {
				continue
			}
			for col := tMinCol; col <= tMaxCol; col++ {
				if col >= p.TilesAcross(level) {
					continue
				}
				if err = p.updateTileAndParent(level, col, row); err != nil {
					return
				}
			}
		}

		x, dx = minX/2, maxX/2-minX/2
		y, dy = minY/2, maxY/2-minY/2
		level--
	}
	return
}

func (p *Image) updateTileAndParent(level, col, row int) (err error) {
	child, err := p.GetTile(level, col, row)
	if err != nil {
		return
	}
	parent, err := p.GetTile(level-1, col/2, row/2)
	if err != nil {
		return
	}
	switch {
	case col%2 == 0 && row%2 == 0:
		draw_ext.DrawPyrDown(
			parent.(draw.Image), image.Rect((p.TileSize.X/2)*0, (p.TileSize.Y/2)*0, p.TileSize.X/2, p.TileSize.Y/2),
			child, image.Pt(0, 0),
			draw_ext.Filter_Average,
		)
	case col%2 == 0 && row%2 == 1:
		draw_ext.DrawPyrDown(
			parent.(draw.Image), image.Rect((p.TileSize.X/2)*0, (p.TileSize.Y/2)*1, p.TileSize.X/2, p.TileSize.Y/2),
			child, image.Pt(0, 0),
			draw_ext.Filter_Average,
		)
	case col%2 == 1 && row%2 == 1:
		draw_ext.DrawPyrDown(
			parent.(draw.Image), image.Rect((p.TileSize.X/2)*1, (p.TileSize.Y/2)*1, p.TileSize.X/2, p.TileSize.Y/2),
			child, image.Pt(0, 0),
			draw_ext.Filter_Average,
		)
	case col%2 == 1 && row%2 == 0:
		draw_ext.DrawPyrDown(
			parent.(draw.Image), image.Rect((p.TileSize.X/2)*1, (p.TileSize.Y/2)*0, p.TileSize.X/2, p.TileSize.Y/2),
			child, image.Pt(0, 0),
			draw_ext.Filter_Average,
		)
	}
	return
}
