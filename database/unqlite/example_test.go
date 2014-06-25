// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package unqlite_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"chai2010.gopkg/database/unqlite"
)

func ExampleOpen() {
	f, err := ioutil.TempFile("./", "unqlite_")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())

	// Open file database.
	db, err := unqlite.Open(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	// Open in memory database.
	db, err = unqlite.Open(":mem:")
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func ExampleHandle_Close() {
	db, err := unqlite.Open(":mem:")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	fmt.Println(err)
	// Output: nil unqlite database
}

func ExampleHandle_Store() {
	db, err := unqlite.Open(":mem:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Store([]byte("key"), []byte("value"))
}

func ExampleHandle_Append() {
	db, err := unqlite.Open(":mem:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Append([]byte("key"), []byte{'a'})
	v, err := db.Fetch([]byte("key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	db.Append([]byte("key"), []byte{'b'})
	v, err = db.Fetch([]byte("key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	// Output: [97]
	// [97 98]
}

func ExampleHandle_Fetch() {
	db, err := unqlite.Open(":mem:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	v, err := db.Fetch([]byte("key"))
	fmt.Println(err)

	err = db.Store([]byte("key"), []byte{'a'})
	v, err = db.Fetch([]byte("key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	// Output: No such record
	// [97]
}

func ExampleHandle_Delete() {
	db, err := unqlite.Open(":mem:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Store([]byte("key"), []byte{'a'})
	v, err := db.Fetch([]byte("key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	err = db.Delete([]byte("key"))
	if err != nil {
		log.Fatal(err)
	}

	v, err = db.Fetch([]byte("key"))
	fmt.Println(err)

	// Output: [97]
	// No such record
}

func ExampleThreadsafe() {
	fmt.Println(unqlite.Threadsafe())
	// Output: true
}
