# Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

project(RAWP)

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

set(RAWP_SRC
  ./include/rawp.h

  ./src/rawp.cc
  ./src/rawp_crc32.cc

  ./src/snappy/rawp-snappy.cc
  ./src/snappy/rawp-snappy-sinksource.cc
  ./src/snappy/rawp-snappy-stubs-internal.cc
)

#------------------------------------------------------------------------------

add_library(rawplib STATIC
  ${RAWP_SRC}
)
if(CMAKE_BUILD_TYPE STREQUAL "debug")
  set_target_properties(rawplib
    PROPERTIES OUTPUT_NAME "rawplib-${OS}-debug"
  )
else()
  set_target_properties(rawplib
    PROPERTIES OUTPUT_NAME "rawplib-${OS}"
  )
endif()

install(TARGETS rawplib
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------
# test

include_directories(AFTER
  ./test
)

add_executable(rawptest
  ./test/test.cc
  ./test/test_util.cc
  ./test/test_util_jpg.cc

  ./src/rawp_test.cc

  ${RAWP_SRC}
)

install(TARGETS rawptest
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------

