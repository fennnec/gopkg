// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// mingw can't use these ldflags:
// -static-libgcc -static-libstdc++ (or -static)
// https://code.google.com/p/go/issues/detail?id=6733
// https://code.google.com/p/go/issues/detail?id=6533

//#include "capi/leveldb_c.cc"
//#include "capi/leveldb_all.cc"

// Build with dll on windows:
// 0. #cgo windows LDFLAGS: -L./capi -lleveldb_c
// 1. install MinGW/MSVC/CMake
// 2. build leveldb_c.dll/libleveldb_c.a with build_msvc.bat
// 3. go run hello.go
