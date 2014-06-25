// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package web is a lightweight web framework for Go.

It's ideal for writing simple, performant backend web services.

Example

	package main

	import (
		"code.google.com/p/chai2010.gopkg/web"
	)

	func hello(val string) string { return "hello " + val }

	func main() {
		web.Get("/(.*)", hello)
		web.Run("0.0.0.0:9999")
	}

Getting parameters

	package main

	import (
		"code.google.com/p/chai2010.gopkg/web"
	)

	func hello(ctx *web.Context, val string) {
		for k,v := range ctx.Params {
			println(k, v)
		}
	}

	func main() {
		web.Get("/(.*)", hello)
		web.Run("0.0.0.0:9999")
	}

In this example, if you visit `http://localhost:9999/?a=1&b=2`, you'll see the following printed out in the terminal:

	a 1
	b 2
*/
package web
