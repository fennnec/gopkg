# Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

project(JXR_LIB)

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

add_definitions(
  -DDISABLE_PERF_MEASUREMENT
)
include_directories(AFTER
  ./include
  ./common/include
  ./image/decode
  ./image/encode
  ./image/sys
  ./jxrgluelib
  .
)

set(JXR_SRC
  ./jxrgluelib/JXRGlue.c
  ./jxrgluelib/JXRMeta.c
  ./jxrgluelib/JXRGluePFC.c
  ./jxrgluelib/JXRGlueJxr.c

  ./image/encode/encode.c
  ./image/encode/segenc.c
  ./image/encode/strenc.c
  ./image/encode/strFwdTransform.c
  ./image/encode/strPredQuantEnc.c

  ./image/decode/decode.c
  ./image/decode/postprocess.c
  ./image/decode/segdec.c
  ./image/decode/strdec.c
  ./image/decode/strInvTransform.c
  ./image/decode/strPredQuantDec.c
  ./image/decode/JXRTranscode.c

  ./image/sys/adapthuff.c
  ./image/sys/image.c
  ./image/sys/strcodec.c
  ./image/sys/strPredQuant.c
  ./image/sys/strTransform.c
  ./image/sys/perfTimerANSI.c

  ./src/jxr.c
)

set(JXR_TEST_SRC
  ./jxrtestlib/JXRTest.c
  ./jxrtestlib/JXRTestBmp.c
  ./jxrtestlib/JXRTestHdr.c
  ./jxrtestlib/JXRTestPnm.c
  ./jxrtestlib/JXRTestTif.c
  ./jxrtestlib/JXRTestYUV.c
)

#------------------------------------------------------------------------------

add_library(jxrlib STATIC
  ${JXR_SRC}
)
if(CMAKE_BUILD_TYPE STREQUAL "debug")
  set_target_properties(jxrlib
    PROPERTIES OUTPUT_NAME "jxrlib-${OS}-debug"
  )
else()
  set_target_properties(jxrlib
    PROPERTIES OUTPUT_NAME "jxrlib-${OS}"
  )
endif()

install(TARGETS jxrlib
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------
# JxrEncApp/JxrDecApp

include_directories(AFTER
  ./jxrtestlib
)

# JxrEncApp
add_executable(JxrEncApp
  ./jxrencoderdecoder/JxrEncApp.c

  ${JXR_TEST_SRC}
  ${JXR_SRC}
)
install(TARGETS JxrEncApp
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

# JxrDecApp
add_executable(JxrDecApp
  ./jxrencoderdecoder/JxrDecApp.c

  ${JXR_TEST_SRC}
  ${JXR_SRC}
)
install(TARGETS JxrDecApp
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------
# test

include_directories(AFTER
  ./test
)

add_executable(jxrtest
  ./test/test.cc
  ./test/test_util.cc
  ./test/test_util_jpg.cc

  ./test/jxr_test.cc
  ./test/jxr_bench_test.cc

  ${JXR_SRC}
)

install(TARGETS jxrtest
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------

