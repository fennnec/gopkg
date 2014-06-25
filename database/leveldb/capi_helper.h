// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef GO_LEVELDB_CAPI_HELPER_H_
#define GO_LEVELDB_CAPI_HELPER_H_

#include <stdlib.h>
#include <leveldb/c.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Write batch */

void go_leveldb_writebatch_iterate_helper(leveldb_writebatch_t* batch, void* state);

/* Comparator */

leveldb_comparator_t* leveldb_comparator_create_helper(void* state);

#ifdef __cplusplus
}
#endif

#endif	// GO_LEVELDB_CAPI_HELPER_H_
