# Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

project(RAW_LIB)

#------------------------------------------------------------------------------

IF(WIN32)
  if(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(OS win64)
  else()
    set(OS win32)
  endif()
else()
  if(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(OS posix64)
  else()
    set(OS posix64)
  endif()
endif()

#------------------------------------------------------------------------------

include_directories(AFTER
  ./include
  ./src/snappy
  ./src
)

set(RAW_SRC
  ./include/raw.h
  ./src/raw_decode.h
  ./src/raw_encode.h

  ./src/raw.cc
  ./src/raw_decode.cc
  ./src/raw_encode.cc
  ./src/raw_crc32.cc

  ./src/snappy/raw-snappy.cc
  ./src/snappy/raw-snappy-sinksource.cc
  ./src/snappy/raw-snappy-stubs-internal.cc
)

#------------------------------------------------------------------------------

add_library(rawlib STATIC
  ${RAW_SRC}
)
if(CMAKE_BUILD_TYPE STREQUAL "debug")
  set_target_properties(rawlib
    PROPERTIES OUTPUT_NAME "rawlib-${OS}-debug"
  )
else()
  set_target_properties(rawlib
    PROPERTIES OUTPUT_NAME "rawlib-${OS}"
  )
endif()

install(TARGETS rawlib
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------
# test

include_directories(AFTER
  ./test
)

add_executable(rawtest
  ./test/test.cc
  ./test/test_util.cc
  ./test/test_util_jpg.cc

  ${RAW_SRC}
)

install(TARGETS rawtest
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------

