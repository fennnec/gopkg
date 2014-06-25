/*
Package vedis provides the ability to create and access vedis databases.

Example:
	package main

	import (
		"chai2010.gopkg/database/vedis"
	)

	func main() {
		// connect
		store, _ := Open(":mem:")
		defer store.Close()

		// set key: x value: 123
		_ = store.Exec("SET x 123")

		// get x
		result, _ := store.ExecResult("GET x")

		// display it
		fmt.Println("x:", result.Int())
	}

See http://vedis.symisc.net/
*/
package vedis
