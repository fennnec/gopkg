// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef RAWP_CRC32_H_
#define RAWP_CRC32_H_

#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

uint32_t rawpHashCRC32(const char* data, size_t n);

#ifdef __cplusplus
}
#endif
#endif // RAWP_CRC32_H_

