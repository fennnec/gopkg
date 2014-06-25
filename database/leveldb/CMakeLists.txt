# Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

project(LEVELDB)

IF(WIN32)
  if(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(LEVELDB_DEF leveldb-cgo-win64.def)
    set(OS win64)
  else()
    set(LEVELDB_DEF leveldb-cgo-win32.def)
    set(OS win32)
  endif()
else()
  if(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(LEVELDB_DEF leveldb-cgo-posix64.def)
    set(OS posix64)
  else()
    set(LEVELDB_DEF leveldb-cgo-posix32.def)
    set(OS posix32)
  endif()
endif()

include_directories(AFTER
  ./include
)

set(LEVELDB_SRC
  ./src/arena.cc
  ./src/block.cc
  ./src/block_builder.cc
  ./src/bloom.cc
  ./src/builder.cc
  ./src/c.cc
  ./src/cache.cc
  ./src/coding.cc
  ./src/comparator.cc
  ./src/crc32c.cc
  ./src/db_impl.cc
  ./src/db_iter.cc
  ./src/dbformat.cc
  ./src/env.cc
  ./src/filename.cc
  ./src/filter_block.cc
  ./src/filter_policy.cc
  ./src/format.cc
  ./src/hash.cc
  ./src/histogram.cc
  ./src/iterator.cc
  ./src/leveldb_env.cc
  ./src/leveldb_port.cc
  ./src/log_reader.cc
  ./src/log_writer.cc
  ./src/logging.cc
  ./src/memtable.cc
  ./src/merger.cc
  ./src/options.cc
  ./src/repair.cc
  ./src/snappy-c.cc
  ./src/snappy-sinksource.cc
  ./src/snappy-stubs-internal.cc
  ./src/snappy.cc
  ./src/status.cc
  ./src/table.cc
  ./src/table_builder.cc
  ./src/table_cache.cc
  ./src/two_level_iterator.cc
  ./src/version_edit.cc
  ./src/version_set.cc
  ./src/write_batch.cc
)

add_library(leveldblib SHARED
  ${LEVELDB_SRC}
  ${LEVELDB_DEF}
)
set_target_properties(leveldblib
  PROPERTIES OUTPUT_NAME "leveldb-cgo-${OS}"
)

install(TARGETS leveldblib
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)
