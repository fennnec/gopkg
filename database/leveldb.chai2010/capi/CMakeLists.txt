# Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

project(LEVEL_DB)

include_directories(AFTER
  ../include
  ../capi
  ..
)

add_library(leveldblib SHARED
  ../capi/leveldb_c.def
  ../capi/leveldb_c.cc
  ../capi/leveldb_all.cc
)
set_target_properties(leveldblib
  PROPERTIES OUTPUT_NAME "leveldb_c"
)

add_executable(leveldbtest
  ./test/test.cc
  ./test/cache_test.cc
  ./test/all_test.cc

  ../capi/leveldb_c.cc
  ../capi/leveldb_all.cc
)
set_target_properties(leveldbtest
  PROPERTIES OUTPUT_NAME "leveldbtest"
)

install(TARGETS leveldblib leveldbtest
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)
