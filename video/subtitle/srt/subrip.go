// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package srt

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Subtitle represents a subtitle.
type Subtitle struct {
	Number    int
	Position  [4]int // X1,X2,Y1,Y2
	StartTime time.Duration
	EndTime   time.Duration
	Texts     []string
}

func ParseFile(name string) (subtitles []Subtitle, err error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return ParseData(data)
}

func ParseData(data []byte) (subtitles []Subtitle, err error) {
	r := newLineReader(string(data))
	for {
		var sub Subtitle
		if sub, err = readSubtitle(r); err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
		subtitles = append(subtitles, sub)
	}
}

func readSubtitle(r *lineReader) (sub Subtitle, err error) {
	r.SkipBlankLine()
	if err = sub.readNumber(r); err != nil {
		return
	}
	if err = sub.readTimeAndPosition(r); err != nil {
		return
	}
	if err = sub.readTexts(r); err != nil {
		return
	}
	return
}

func (p *Subtitle) readNumber(r *lineReader) (err error) {
	var s string
	if s, err = r.ReadLine(); err != nil {
		return err
	}
	if p.Number, err = strconv.Atoi(s); err != nil {
		return fmt.Errorf("srt: bad number, line = %d", r.CurrentPos())
	}
	return nil
}

// 00:00:15,000 --> 00:00:18,000
// 00:00:10,500 --> 00:00:13,000 X1:63 X2:223 Y1:43 Y2:58
func (p *Subtitle) readTimeAndPosition(r *lineReader) (err error) {
	var s string
	if s, err = r.ReadLine(); err != nil {
		return err
	}
	if err = p.readTimeValue(s, r.CurrentPos()); err != nil {
		return err
	}
	if err = p.readPositionValue(s, r.CurrentPos()); err != nil {
		return err
	}
	return nil
}

// 00:00:15,000 --> 00:00:18,000
func (p *Subtitle) readTimeValue(s string, pos int) (err error) {
	if len(s) < len(`00:00:15,000 --> 00:00:18,000`) {
		return fmt.Errorf("srt: bad time, line = %d", pos)
	}
	p.StartTime, err = parseDuration(s[:12])
	p.EndTime, err = parseDuration(s[17:])
	if err != nil {
		return fmt.Errorf("srt: bad time, line = %d", pos)
	}
	return nil
}

// 00:00:10,500 --> 00:00:13,000 X1:63 X2:223 Y1:43 Y2:58
func (p *Subtitle) readPositionValue(s string, pos int) (err error) {
	tags := [4]string{"X1", "X2", "Y1", "Y2"} // Need math p.Position's order
	for i := 0; i < len(tags); i++ {
		idx := strings.Index(s, tags[i])
		if idx < 0 {
			continue
		}
		t := s[idx+3:] // X1:??<spece><\t>...
		if idx = strings.IndexAny(t, " \n\t\r\n"); idx >= 0 {
			t = t[:idx]
		}
		if p.Position[i], err = strconv.Atoi(t); err != nil {
			return fmt.Errorf("srt: bad position, line = %d", pos)
		}
	}
	return nil
}

func (p *Subtitle) readTexts(r *lineReader) (err error) {
	var s string
	for {
		if s, err = r.ReadLine(); err != nil {
			return err
		}
		if strings.TrimSpace(s) == "" {
			return nil
		}
		p.Texts = append(p.Texts, s)
	}
}
