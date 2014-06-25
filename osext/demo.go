// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	myos "."
)

func main() {
	fl, err := myos.NewFileLock("demo.lock")
	if err == nil {
		fmt.Printf("Lock Success!\n")
		defer fl.Release()
	} else {
		fmt.Printf("Lock Failed!\n")
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("Quit (%v)\n", <-ch)
}
