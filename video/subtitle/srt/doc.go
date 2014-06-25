// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package srt privades a parser for .SRT files, a widely used subtitle text file format.

How To Use

Example:

	import (
		"chai2010.gopkg/video/subtitle/srt"
	)

	func main() {
		subtitles, err := srt.ParseData([]byte(srtText))
		if err != nil {
			log.Fatal(err)
		}
		...
	}

	const srtText = `
	1
	00:00:10,500 --> 00:00:13,000
	Elephant's Dream

	2
	00:00:15,000 --> 00:00:18,000
	At
	`

SubRip Text File Format

The SubRip file format, as reported on the Matroska multimedia container format website,
is "perhaps the most basic of all subtitle formats."
SubRip (SubRip Text) files are named with the extension .srt,
and contain formatted lines of plain text in groups separated by a blank line.
Subtitles are numbered sequentially, starting at 1.
The timecode format used is hours:minutes:seconds,milliseconds with time units fixed to two zero-padded digits and fractions fixed to three zero-padded digits (00:00:00,000).
The fractional separator used is the comma, since the program was written in France.
The subtitle separator, a blank line, is the double byte MS-DOS CR+LF pair,
though the POSIX single byte linefeed is also well supported.

	1. A numeric counter identifying each sequential subtitle
	2. The time that the subtitle should appear on the screen, followed by --> and the time it should disappear
	3. Subtitle text itself on one or more lines
	4. A blank line containing no text, indicating the end of this subtitle


SubRip (.srt) structure examples:

	1
	00:00:10,500 --> 00:00:13,000
	Elephant's Dream

	2
	00:00:15,000 --> 00:00:18,000
	At the left we can see...

With specific positioning and styling:

	1
	00:00:10,500 --> 00:00:13,000 X1:63 X2:223 Y1:43 Y2:58
	<i>Elephant's Dream</i>

	2
	00:00:15,000 --> 00:00:18,000 X1:53 X2:303 Y1:438 Y2:453
	<font color="cyan">At the left we can see...</font>

Formatting:

Unofficially the format has very basic text formatting,
which can be either interpreted or passed through for rendering depending on the processing application.
Formatting is derived from HTML tags for bold, italic, underline and color:

	- Bold – <b> ... </b> or {b} ... {/b}
	- Italic – <i> ... </i> or {i} ... {/i}
	- Underline – <u> ... </u> or {u} ... {/u}
	- Font color – <font color="color name or #code"> ... </font> (as in HTML)

Nested tags are allowed; some implementations prefer whole-line formatting only.

See http://en.wikipedia.org/wiki/SubRip for more infomation.
*/
package srt
