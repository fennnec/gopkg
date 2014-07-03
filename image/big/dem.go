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
		panic(fmt.Sprintf("image/big: NewDem, bad arguments: %v, %v", r, tileSize))
	}
	return &Dem{
		TileMap:   makeDemTileMap(r, tileSize),
		TileSize:  tileSize,
		Rect:      r,
		ZeroValue: zeroValue,
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

func (p *Dem) Levels() int {
	return len(p.TileMap)
}

func (p *Dem) TilesAcross(level int) int {
	return len(p.TileMap[level])
}

func (p *Dem) TilesDown(level int) int {
	return len(p.TileMap[level][0])
}

func (p *Dem) GetTile(level, col, row int) (m *image_ext.Gray32f, err error) {
	if m = p.TileMap[level][col][row]; m != nil {
		return
	}
	err = fmt.Errorf("image/big: Dem.GetTile, not found tile")
	return
}

func (p *Dem) SetTile(level, col, row int, m *image_ext.Gray32f) (err error) {
	if m.Bounds() != image.Rect(0, 0, p.TileSize.X, p.TileSize.Y) {
		err = fmt.Errorf("image/big: Dem.SetTile, bad bound size: %v", m.Bounds())
		return
	}
	p.TileMap[level][col][row] = m
	return
}

func (p *Dem) ReadRect(level, x, y, dx, dy int) (m *image_ext.Gray32f, err error) {
	if level < 0 || level >= p.Levels() {
		err = fmt.Errorf(
			"image/big: Dem.ReadRect, level = %d, x = %d, y = %d, dx = %d, dy = %d",
			level, x, y, dx, dy,
		)
		return
	}
	if !image.Rect(x, y, x+dx, y+dy).In(p.Bounds()) {
		err = fmt.Errorf(
			"image/big: Dem.ReadRect, level = %d, x = %d, y = %d, dx = %d, dy = %d",
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

	m = newDemTile(p.TileSize, p.ZeroValue)
	for col := tMinX; col < tMaxX; col++ {
		for row := tMinY; row < tMaxY; row++ {
			var tile *image_ext.Gray32f
			if tile, err = p.GetTile(level, col, row); err != nil {
				return
			}
			p.readRectFromTile(m, tile, x, y, dx, dy, col, row)
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

func (p *Dem) WriteRect(level, x, y, dx, dy int, m *image_ext.Gray32f) (err error) {
	if level < 0 || level >= p.Levels() {
		err = fmt.Errorf(
			"image/big: Dem.WriteRect, level = %d, x = %d, y = %d, dx = %d, dy = %d",
			level, x, y, dx, dy,
		)
		return
	}
	if !image.Rect(x, y, x+dx, y+dy).In(p.Bounds()) {
		err = fmt.Errorf(
			"image/big: Dem.WriteRect, level = %d, x = %d, y = %d, dx = %d, dy = %d",
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

	for col := tMinX; col < tMaxX; col++ {
		for row := tMinY; row < tMaxY; row++ {
			var tile *image_ext.Gray32f
			if tile, err = p.GetTile(level, col, row); err != nil {
				return
			}
			p.writeRectToTile(m, tile, x, y, dx, dy, col, row)
		}
	}

	err = p.updateRectPyramid(level, x, y, dx, dy)
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

func (p *Dem) updateTileAndParent(level, col, row int) (err error) {
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
