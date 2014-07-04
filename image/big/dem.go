// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"fmt"
	"image"
	"image/color"

	image_ext "github.com/chai2010/gopkg/image"
	color_ext "github.com/chai2010/gopkg/image/color"
	draw_ext "github.com/chai2010/gopkg/image/draw"
)

type Dem struct {
	TileMap   [][][]*image_ext.Gray32f // m.TileMap[level][col][row]
	TileSize  image.Point
	Rect      image.Rectangle
	ZeroValue color_ext.Gray32f
}

func NewDem(r image.Rectangle, tileSize image.Point, zeroValue color_ext.Gray32f) *Dem {
	if r.Empty() || tileSize.X <= 0 || tileSize.Y <= 0 {
		panic(fmt.Sprintf("image/big: NewDem, bad arguments: r = %v, tileSize = %v", r, tileSize))
	}
	return &Dem{
		TileMap:   makeDemTileMap(r, tileSize),
		TileSize:  tileSize,
		Rect:      r,
		ZeroValue: zeroValue,
	}
}

func (p *Dem) SubLevels(levels int) *Dem {
	r := p.Rect
	for i := levels; i < p.Levels(); i++ {
		r.Min.X /= 2
		r.Min.Y /= 2
		r.Max.X /= 2
		r.Max.Y /= 2
	}
	return &Dem{
		TileMap:   p.TileMap[:levels],
		TileSize:  p.TileSize,
		Rect:      r,
		ZeroValue: p.ZeroValue,
	}
}

func (p *Dem) ColorModel() color.Model { return color_ext.Gray32fModel }

func (p *Dem) Bounds() image.Rectangle { return p.Rect }

func (p *Dem) At(x, y int) color.Color {
	return p.Gray32fAt(x, y)
}

func (p *Dem) Gray32fAt(x, y int) color_ext.Gray32f {
	level, col, row := p.Levels()-1, x/p.TileSize.X, y/p.TileSize.Y
	if m := p.TileMap[level][col][row]; m != nil {
		return m.At(x%p.TileSize.X, y%p.TileSize.Y).(color_ext.Gray32f)
	}
	return color_ext.Gray32f{}
}

func (p *Dem) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	m := p.GetTile(p.Levels()-1, x/p.TileSize.X, y/p.TileSize.Y)
	m.Set(x%p.TileSize.X, y%p.TileSize.Y, c)
	return
}

func (p *Dem) SetGray32f(x, y int, c color_ext.Gray32f) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	m := p.GetTile(p.Levels()-1, x/p.TileSize.X, y/p.TileSize.Y)
	m.SetGray32f(x%p.TileSize.X, y%p.TileSize.Y, c)
	return
}

func (p *Dem) Levels() int {
	return len(p.TileMap)
}

func (p *Dem) adjustLevel(level int) int {
	if level < 0 {
		return p.Levels() + level
	}
	return level
}

func (p *Dem) TilesAcross(level int) int {
	level = p.adjustLevel(level)
	v := len(p.TileMap[level])
	return v
}

func (p *Dem) TilesDown(level int) int {
	level = p.adjustLevel(level)
	v := len(p.TileMap[level][0])
	return v
}

func (p *Dem) GetTile(level, col, row int) (m *image_ext.Gray32f) {
	level = p.adjustLevel(level)
	if m = p.TileMap[level][col][row]; m != nil {
		return
	}
	m = newDemTile(p.TileSize, p.ZeroValue)
	p.TileMap[level][col][row] = m
	return
}

func (p *Dem) SetTile(level, col, row int, m *image_ext.Gray32f) (err error) {
	level = p.adjustLevel(level)
	if m.Bounds() != image.Rect(0, 0, p.TileSize.X, p.TileSize.Y) {
		err = fmt.Errorf("image/big: Dem.SetTile, bad bound size: %v", m.Bounds())
		return
	}
	p.TileMap[level][col][row] = m
	return
}

func (p *Dem) ReadRect(r image.Rectangle, level int) (m *image_ext.Gray32f, err error) {
	level = p.adjustLevel(level)
	if !r.In(p.Bounds()) {
		err = fmt.Errorf("image/big: Dem.ReadRect, r = %v, level = %v", r, level)
		return
	}
	if level < 0 || level >= p.Levels() {
		err = fmt.Errorf("image/big: Dem.ReadRect, r = %v, level = %v", r, level)
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

	m = newDemTile(p.TileSize, p.ZeroValue)
	for col := tMinX; col < tMaxX; col++ {
		for row := tMinY; row < tMaxY; row++ {
			p.readRectFromTile(m, p.GetTile(level, col, row), r.Min.X, r.Min.Y, r.Dx(), r.Dy(), col, row)
		}
	}
	return
}

func (p *Dem) readRectFromTile(dst, tile *image_ext.Gray32f, x, y, dx, dy, col, row int) {
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
		dst, image.Rect(
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

func (p *Dem) WriteRect(r image.Rectangle, m *image_ext.Gray32f, level int) (err error) {
	level = p.adjustLevel(level)
	if !r.In(p.Bounds()) || level < 0 || level >= p.Levels() {
		err = fmt.Errorf("image/big: Dem.WriteRect, r = %v, level = %v", r, level)
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
			p.writeRectToTile(m, p.GetTile(level, col, row), r.Min.X, r.Min.Y, r.Dx(), r.Dy(), col, row)
		}
	}

	err = p.updateRectPyramid(level, r.Min.X, r.Min.Y, r.Dx(), r.Dy())
	return
}

func (p *Dem) writeRectToTile(src, tile *image_ext.Gray32f, x, y, dx, dy, col, row int) {
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

func (p *Dem) updateRectPyramid(level, x, y, dx, dy int) (err error) {
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

func (p *Dem) updateParentTile(level, col, row int) (err error) {
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
