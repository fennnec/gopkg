@rem Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
@rem Use of this source code is governed by a BSD-style
@rem license that can be found in the LICENSE file.

@rem --------------------------------------------------------------------------
@rem build

mkdir zz_build_tmp
cd    zz_build_tmp

cmake ..^
  -G "NMake Makefiles"^
  -DCMAKE_BUILD_TYPE=release^
  -DCMAKE_INSTALL_PREFIX=..^
  ^
  -DCMAKE_C_FLAGS_DEBUG="/MTd /Zi /Od /Ob0 /RTC1"^
  -DCMAKE_CXX_FLAGS_DEBUG="/MTd /Zi /Od /Ob0 /RTC1"^
  ^
  -DCMAKE_C_FLAGS_RELEASE="/MT /O2 /Ob2 /DNDEBUG"^
  -DCMAKE_CXX_FLAGS_RELEASE="/MT /O2 /Ob2 /DNDEBUG"^
  ^
  -DCMAKE_EXE_LINKER_FLAGS="/MANIFEST:NO"

nmake VERBOSE=1
nmake install

cd ..

@rem --------------------------------------------------------------------------
@rem MinGW: generate libleveldb.a

dlltool -dllname leveldb_c.dll --def leveldb_c.def --output-lib libleveldb_c.a

@rem --------------------------------------------------------------------------
@rem move leveldb.dll $GOPATH/bin

copy leveldb_c.dll ..

@rem --------------------------------------------------------------------------
@rem PAUSE
