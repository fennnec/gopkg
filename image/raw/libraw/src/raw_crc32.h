// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef RAW_CRC32_H_
#define RAW_CRC32_H_

#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

uint32_t rawHashCRC32(const char* data, size_t n);

#ifdef __cplusplus
}
#endif
#endif // RAW_CRC32_H_

