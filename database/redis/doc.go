// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package redix provides a minimalistic Redis client library.

Example:

	import (
		"chai2010.gopkg/datebase/redis"
	)

	func main() {
		var c *redis.Client
		var r *redis.Reply
		var s string
		var err error

		if c, err = redis.Dial("tcp", "127.0.0.1:6379"); err != nil {
			log.Fatal(err)
		}
		defer c.Close()

		// select database
		if r = c.Cmd("select", 8); r.Err != nil {
			log.Fatal(err)
		}

		if r = c.Cmd("flushdb"); r.Err != nil {
			log.Fatal(err)
		}

		if r = c.Cmd("echo", "Hello world!"); r.Err != nil {
			log.Fatal(err)
		}
		fmt.Println("echo:", r)

		//* Strings
		if r = c.Cmd("set", "mykey0", "myval0"); r.Err != nil {
			log.Fatal(err)
		}

		if s, err = c.Cmd("get", "mykey0").Str(); r.Err != nil {
			log.Fatal(err)
		}
		fmt.Println("mykey0:", s)

		myhash := map[string]string{
			"mykey1": "myval1",
			"mykey2": "myval2",
			"mykey3": "myval3",
		}

		// Alternatively:
		// c.Cmd("mset", "mykey1", "myval1", "mykey2", "myval2", "mykey3", "myval3")
		if r = c.Cmd("mset", myhash); r.Err != nil {
			log.Fatal(err)
		}

		ls, err := c.Cmd("mget", "mykey1", "mykey2", "mykey3").List()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("mykeys values:", ls)

		//* List handling
		mylist := []string{"foo", "bar", "qux"}

		// Alternativaly:
		// c.Cmd("rpush", "mylist", "foo", "bar", "qux")
		if r = c.Cmd("rpush", "mylist", mylist); r.Err != nil {
			log.Fatal(err)
		}

		mylist, err = c.Cmd("lrange", "mylist", 0, -1).List()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("mylist:", mylist)

		//* Hash handling

		// Alternatively:
		// c.Cmd("hmset", "myhash", ""mykey1", "myval1", "mykey2", "myval2", "mykey3", "myval3")
		if r = c.Cmd("hmset", "myhash", myhash); r.Err != nil {
			log.Fatal(err)
		}

		myhash, err = c.Cmd("hgetall", "myhash").Hash()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("myhash:", myhash)

		//* Pipelining
		c.Append("set", "multikey", "multival")
		c.Append("get", "multikey")

		c.GetReply()     // set
		r = c.GetReply() // get
		if r.Err != nil {
			log.Fatal(err)
		}

		s, err = r.Str()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("multikey:", s)
	}
*/
package redis
