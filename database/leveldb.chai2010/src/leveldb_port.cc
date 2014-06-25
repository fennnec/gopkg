// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "port/port.h"

#if defined(LEVELDB_PLATFORM_POSIX)
#  include "./port/port_posix.cc"
#elif defined(LEVELDB_PLATFORM_WINDOWS)
#  include "./port/port_windows.cc"
#endif
