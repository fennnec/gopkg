// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
)

func _main() {
	var c int

	OptErr = 0
	for {
		if c = Getopt("a:bh"); c == EOF {
			break
		}
		switch c {
		case 'a':
			println("a=", OptArg)
		case 'b':
			println("i:on")
		case 'h':
			println("usage: example [-a foo|-b|-h]")
			os.Exit(1)
		}
	}

	for n := OptInd; n < len(os.Args); n++ {
		println(os.Args[n])
	}
}
