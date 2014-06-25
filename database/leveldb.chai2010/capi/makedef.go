// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./leveldb_c.h")
	if err != nil {
		log.Fatal("ioutil.ReadFile: ", err)
	}

	var funcs []string
	var re = regexp.MustCompile(`leveldb_[a-z_]*\(`)
	for _, line := range strings.Split(string(data), "\n") {
		if s := re.FindString(line); s != "" {
			funcs = append(funcs, s[:len(s)-1])
		}
	}
	sort.Strings(funcs)

	var b bytes.Buffer
	fmt.Fprintf(&b, header[1:])
	for _, s := range funcs {
		fmt.Fprintf(&b, "\t%s\n", s)
	}

	err = ioutil.WriteFile("leveldb_c.def", b.Bytes(), 0666)
	if err != nil {
		log.Fatal("ioutil.WriteFile: ", err)
	}

	fmt.Printf("Done\n")
}

var header = `
; Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
; Use of this source code is governed by a BSD-style
; license that can be found in the LICENSE file.

; Auto Genrated by makedef.go; DO NOT EDIT!!

LIBRARY leveldb_c.dll

EXPORTS
`
