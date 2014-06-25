// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package srt

import (
	"reflect"
	"testing"
)

func TestSubtitle(t *testing.T) {
	for i := 0; i < len(testSubtitleSuits); i++ {
		golden := testSubtitleSuits[i].subtitles
		s, err := ParseData([]byte(testSubtitleSuits[i].text))
		if err != nil {
			t.Fatalf("%d: parse failed, err = %v", i, err)
		}
		if !reflect.DeepEqual(s, golden) {
			t.Fatalf("%d: expected = %q,\ngot = %q", i, golden, s)
		}
	}
}

var testSubtitleSuits = []struct {
	text      string
	subtitles []Subtitle
}{
	{
		text: `
1
00:00:10,500 --> 00:00:13,000
Elephant's Dream

2
00:00:15,000 --> 00:00:18,000
At the left we can see...
`,
		subtitles: []Subtitle{
			{
				Number:    1,
				Position:  [4]int{},
				StartTime: makeDuration(0, 0, 10, 500),
				EndTime:   makeDuration(0, 0, 13, 0),
				Texts:     []string{`Elephant's Dream`},
			},
			{
				Number:    2,
				Position:  [4]int{},
				StartTime: makeDuration(0, 0, 15, 0),
				EndTime:   makeDuration(0, 0, 18, 0),
				Texts:     []string{`At the left we can see...`},
			},
		},
	},

	{
		text: `
1
00:00:10,500 --> 00:00:13,000 X1:63 X2:223 Y1:43 Y2:58
<i>Elephant's Dream</i>

2
00:00:15,000 --> 00:00:18,000 X1:53 X2:303 Y1:438 Y2:453
<font color="cyan">At the left we can see...</font>
`,
		subtitles: []Subtitle{
			{
				Number:    1,
				Position:  [4]int{63, 223, 43, 58},
				StartTime: makeDuration(0, 0, 10, 500),
				EndTime:   makeDuration(0, 0, 13, 0),
				Texts:     []string{`<i>Elephant's Dream</i>`},
			},
			{
				Number:    2,
				Position:  [4]int{53, 303, 438, 453},
				StartTime: makeDuration(0, 0, 15, 0),
				EndTime:   makeDuration(0, 0, 18, 0),
				Texts:     []string{`<font color="cyan">At the left we can see...</font>`},
			},
		},
	},

	{
		text: `
1
00:00:16,240 --> 00:00:18,020
你干什么
Hey, what are you doing?

2
00:00:22,140 --> 00:00:23,380
天哪
Jesus!

3
00:00:23,730 --> 00:00:24,560
你看清肇事的车了吗
Did you get a good look?
`,
		subtitles: []Subtitle{
			{
				Number:    1,
				Position:  [4]int{},
				StartTime: makeDuration(0, 0, 16, 240),
				EndTime:   makeDuration(0, 0, 18, 20),
				Texts:     []string{`你干什么`, `Hey, what are you doing?`},
			},
			{
				Number:    2,
				Position:  [4]int{},
				StartTime: makeDuration(0, 0, 22, 140),
				EndTime:   makeDuration(0, 0, 23, 380),
				Texts:     []string{`天哪`, `Jesus!`},
			},
			{
				Number:    3,
				Position:  [4]int{},
				StartTime: makeDuration(0, 0, 23, 730),
				EndTime:   makeDuration(0, 0, 24, 560),
				Texts:     []string{`你看清肇事的车了吗`, `Did you get a good look?`},
			},
		},
	},
}
