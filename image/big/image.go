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
	TileMap  [][][]draw.Image // m.TileMap[level][col][row]
	TileSize image.Point
	Rect     image.Rectangle
	Model    color.Model // Gray/Gray16/Gray32f/RGBA/RGBA64/RGBA128f
}

func NewImage(r image.Rectangle, tileSize image.Point, model color.Model) *Image {
	if r.Empty() || tileSize.X <= 0 || tileSize.Y <= 0 {
		panic(fmt.Errorf("image/big: NewImage, bad arguments: r = %v, tileSize = %v", r, tileSize))
	}
	if !isValidImageColorModel(model) {
		panic(fmt.Errorf("image/big: NewImage, bad color model: %T", model))
	}
	return &Image{
		TileMap:  makeImageTileMap(r, tileSize),
		TileSize: tileSize,
		Rect:     r,
		Model:    model,
	}
}

func (p *Image) SubLevels(levels int) *Image {
	r := p.Rect
	for i := levels; i < p.Levels(); i++ {
		r.Min.X /= 2
		r.Min.Y /= 2
		r.Max.X /= 2
		r.Max.Y /= 2
	}
	return &Image{
		TileMap:  p.TileMap[:levels],
		TileSize: p.TileSize,
		Rect:     r,
		Model:    p.Model,
	}
}

func (p *Image) ColorModel() color.Model { return p.Model }

func (p *Image) Bounds() image.Rectangle { return p.Rect }

func (p *Image) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color.Gray{}
	}
	m := p.GetTile(p.Levels()-1, x/p.TileSize.X, y/p.TileSize.Y)
	c := m.At(x%p.TileSize.X, y%p.TileSize.Y)
	return c
}

func (p *Image) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	m := p.GetTile(p.Levels()-1, x/p.TileSize.X, y/p.TileSize.Y)
	m.Set(x%p.TileSize.X, y%p.TileSize.Y, c)
	return
}

func (p *Image) Levels() int {
	return len(p.TileMap)
}

func (p *Image) adjustLevel(level int) int {
	if level < 0 {
		return p.Levels() + level
	}
	return level
}

func (p *Image) TilesAcross(level int) int {
	level = p.adjustLevel(level)
	v := len(p.TileMap[level])
	return v
}

func (p *Image) TilesDown(level int) int {
	level = p.adjustLevel(level)
	v := len(p.TileMap[level][0])
	return v
}

func (p *Image) GetTile(level, col, row int) (m draw.Image) {
	level = p.adjustLevel(level)
	if m = p.TileMap[level][col][row]; m != nil {
		return
	}
	m = newImageTile(p.TileSize, p.Model)
	p.TileMap[level][col][row] = m
	return
}

func (p *Image) SetTile(level, col, row int, m draw.Image) (err error) {
	level = p.adjustLevel(level)
	if m.Bounds() != image.Rect(0, 0, p.TileSize.X, p.TileSize.Y) {
		err = fmt.Errorf("image/big: Image.SetTile, bad bound size: %v", m.Bounds())
		return
	}
	if m.ColorModel() != p.Model {
		err = fmt.Errorf("image/big: Image.SetTile, bad color model: %T", m.ColorModel())
		return
	}
	p.TileMap[level][col][row] = m
	return
}

func (p *Image) ReadRect(r image.Rectangle, level int) (m image.Image, err error) {
	level = p.adjustLevel(level)
	if !r.In(p.Bounds()) || level < 0 || level >= p.Levels() {
		err = fmt.Errorf("image/big: Image.ReadRect, rect = %v, level = %d", r, level)
		return
	}

	tMinX := r.Min.X / p.TileSize.X
	tMinY := r.Min.Y / p.TileSize.Y
	tMaxX := (r.Min.X + r.Dx() + p.TileSize.X - 1) / p.TileSize.X
	tMaxY := (r.Min.Y + r.Dy() + p.TileSize.Y - 1) / p.TileSize.Y

	if max := p.TilesAcross(level); tMaxX > max {
		tMaxX = max
	}
	if max := p.TilesDown(level); tMaxY > max {
		tMaxY = max
	}

	dst := newImageTile(p.TileSize, p.Model)
	for col := tMinX; col < tMaxX; col++ {
		for row := tMinY; row < tMaxY; row++ {
			p.readRectFromTile(dst, p.GetTile(level, col, row), r.Min.X, r.Min.Y, r.Dx(), r.Dy(), col, row)
		}
	}
	m = dst
	return
}

func (p *Image) readRectFromTile(dst, tile draw.Image, x, y, dx, dy, col, row int) {
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

func (p *Image) WriteRect(r image.Rectangle, m image.Image, level int) (err error) {
	level = p.adjustLevel(level)
	if !r.In(p.Bounds()) || level < 0 || level >= p.Levels() {
		err = fmt.Errorf("image/big: Image.WriteRect, rect = %v, level = %v", r, level)
		return
	}
	if m.ColorModel() != p.Model {
		err = fmt.Errorf("image/big: Image.WriteRect, bad color model: %T", m.ColorModel())
		return
	}

	tMinX := r.Min.X / p.TileSize.X
	tMinY := r.Min.Y / p.TileSize.Y
	tMaxX := (r.Min.X + r.Dx() + p.TileSize.X - 1) / p.TileSize.X
	tMaxY := (r.Min.Y + r.Dy() + p.TileSize.Y - 1) / p.TileSize.Y

	if max := p.TilesAcross(level); tMaxX > max {
		tMaxX = max
	}
	if max := p.TilesDown(level); tMaxY > max {
		tMaxY = max
	}

	for col := tMinX; col < tMaxX; col++ {
		for row := tMinY; row < tMaxY; row++ {
			p.writeRectToTile(p.GetTile(level, col, row), m, r.Min.X, r.Min.Y, r.Dx(), r.Dy(), col, row)
		}
	}

	err = p.updateRectPyramid(level, r.Min.X, r.Min.Y, r.Dx(), r.Dy())
	return
}

func (p *Image) writeRectToTile(tile draw.Image, src image.Image, x, y, dx, dy, col, row int) {
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
		tile, image.Rect(
			zMinX-tMinX,
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
				if err = p.updateParentTile(level, col, row); err != nil {
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

func (p *Image) updateParentTile(level, col, row int) (err error) {
	parent, child := p.GetTile(level-1, col/2, row/2), p.GetTile(level, col, row)
	switch {
	case col%2 == 0 && row%2 == 0:
		draw_ext.DrawPyrDown(
			parent, image.Rect((p.TileSize.X/2)*0, (p.TileSize.Y/2)*0, p.TileSize.X/2, p.TileSize.Y/2),
			child, image.Pt(0, 0),
			draw_ext.Filter_Average,
		)
	case col%2 == 0 && row%2 == 1:
		draw_ext.DrawPyrDown(
			parent, image.Rect((p.TileSize.X/2)*0, (p.TileSize.Y/2)*1, p.TileSize.X/2, p.TileSize.Y/2),
			child, image.Pt(0, 0),
			draw_ext.Filter_Average,
		)
	case col%2 == 1 && row%2 == 1:
		draw_ext.DrawPyrDown(
			parent, image.Rect((p.TileSize.X/2)*1, (p.TileSize.Y/2)*1, p.TileSize.X/2, p.TileSize.Y/2),
			child, image.Pt(0, 0),
			draw_ext.Filter_Average,
		)
	case col%2 == 1 && row%2 == 0:
		draw_ext.DrawPyrDown(
			parent, image.Rect((p.TileSize.X/2)*1, (p.TileSize.Y/2)*0, p.TileSize.X/2, p.TileSize.Y/2),
			child, image.Pt(0, 0),
			draw_ext.Filter_Average,
		)
	}
	return
}
