// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "jxr.h"

int jxr_decode(
	char* buf, int buf_len, const char* data, int size,
	int* width, int* height, int* channels, int* depth,
	jxr_data_type_t* type
) {
	return -1;
}

int jxr_encode(
	char* buf, int buf_len, const char* data, int size,
	int width, int height, int channels, int depth,
	int quality, int width_step,
	jxr_data_type_t type
) {
	return -1;
}
