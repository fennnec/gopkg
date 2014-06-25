// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type PackageConfig struct {
	Name            string
	Description     string
	URL             string
	Version         string
	Requires        string
	RequiresPrivate string
	Conflicts       string
	Cflags          string
	Libs            string
	LibsPrivate     string
	Variables       map[string]string
}

func ParsePackageConfigFile(name string) (*PackageConfig, error) {
	return nil, nil
}

func ParsePackageConfigData(data []byte) (*PackageConfig, error) {
	return nil, nil
}
