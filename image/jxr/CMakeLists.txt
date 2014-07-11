# Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

project(JXR_LIB)

#------------------------------------------------------------------------------

IF(WIN32)
  if(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(JXR_DEF jxr-cgo-win64.def)
    set(OS win64)
  else()
    set(JXR_DEF jxr-cgo-win32.def)
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
  ./jxrlib/include
  ./jxrlib/common/include
  ./jxrlib/image/decode
  ./jxrlib/image/encode
  ./jxrlib/image/sys
  ./jxrlib/jxrgluelib
)

set(JXR_SRC
  ./jxrlib/jxrgluelib/JXRGlue.c
  ./jxrlib/jxrgluelib/JXRMeta.c
  ./jxrlib/jxrgluelib/JXRGluePFC.c
  ./jxrlib/jxrgluelib/JXRGlueJxr.c

  ./jxrlib/image/encode/encode.c
  ./jxrlib/image/encode/segenc.c
  ./jxrlib/image/encode/strenc.c
  ./jxrlib/image/encode/strFwdTransform.c
  ./jxrlib/image/encode/strPredQuantEnc.c

  ./jxrlib/image/decode/decode.c
  ./jxrlib/image/decode/postprocess.c
  ./jxrlib/image/decode/segdec.c
  ./jxrlib/image/decode/strdec.c
  ./jxrlib/image/decode/strInvTransform.c
  ./jxrlib/image/decode/strPredQuantDec.c
  ./jxrlib/image/decode/JXRTranscode.c

  ./jxrlib/image/sys/adapthuff.c
  ./jxrlib/image/sys/image.c
  ./jxrlib/image/sys/strcodec.c
  ./jxrlib/image/sys/strPredQuant.c
  ./jxrlib/image/sys/strTransform.c
  ./jxrlib/image/sys/perfTimerANSI.c

  ./jxrlib/src/jxr.c
  ./jxrlib/src/jxr_private.c
)

add_library(jxrlib SHARED
  ${JXR_SRC}
  ${JXR_DEF}
)
set_target_properties(jxrlib
  PROPERTIES OUTPUT_NAME "jxr-cgo-${OS}"
)

install(TARGETS jxrlib
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)
