// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package srt

import (
	"io"
	"strings"
)

type lineReader struct {
	lines []string
	pos   int
}

func newLineReader(data string) *lineReader {
	data = strings.Replace(data, "\r", "", -1)
	lines := strings.Split(data, "\n")
	return &lineReader{lines: lines}
}

func (r *lineReader) AtEOF() bool {
	return r.pos >= len(r.lines)
}

func (r *lineReader) CurrentPos() int {
	return r.pos
}

func (r *lineReader) CurrentLine() string {
	if !r.AtEOF() {
		return r.lines[r.pos]
	}
	return ""
}

func (r *lineReader) Next() {
	r.pos++
}

func (r *lineReader) ReadLine() (s string, err error) {
	if r.pos >= len(r.lines) {
		err = io.EOF
		return
	}
	s = r.lines[r.pos]
	r.pos++
	return
}

func (r *lineReader) UnreadLine() {
	if r.pos >= 0 {
		r.pos--
	}
}

func (r *lineReader) SkipBlankLine() {
	for ; r.pos < len(r.lines); r.pos++ {
		if strings.TrimSpace(r.lines[r.pos]) != "" {
			break
		}
	}
}

func (r *lineReader) Reset() {
	r.pos = 0
}
