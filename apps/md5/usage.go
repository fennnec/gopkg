// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var _ = `
md5sum --help
Usage: md5sum [OPTION] [FILE]...
Print or check MD5 (128-bit) checksums.
With no FILE, or when FILE is -, read standard input.

  -b, --binary            read in binary mode (default unless reading tty stdin)

  -c, --check             read MD5 sums from the FILEs and check them
  -t, --text              read in text mode (default if reading tty stdin)

The following two options are useful only when verifying checksums:
      --status            don't output anything, status code shows success
  -w, --warn              warn about improperly formatted checksum lines

      --help     display this help and exit
      --version  output version information and exit

The sums are computed as described in RFC 1321.  When checking, the input
should be a former output of this program.  The default mode is to print
a line with checksum, a character indicating type ('*' for binary, ' ' for
text), and name for each FILE.

Report bugs to <bug-coreutils@gnu.org>.
`

var _ = `
md5sum -h
md5sum: invalid option -- h
Try 'md5sum --help' for more information.
`

var _ = `
md5sum *
00000000000000000000000000000000 *a.go
00000000000000000000000000000000 *b.go
00000000000000000000000000000000 *c.go
`

// stdin
var _ = `
md5sum
00000000000000000000000000000000 *-
`
