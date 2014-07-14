// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "jxr.h"
#include "jxr_private.h"

// TODO(chai2010):

jxr_encoder_t* jxr_encoder_new() {
	return NULL;
}
void jxr_encoder_delete(jxr_encoder_t* p) {
	//
}

jxr_bool_t jxr_encoder_init(jxr_encoder_t* p,
	const char* data, int stride,
	int width, int height, int channels, int depth,
	int quality, jxr_data_type_t type
) {
	return jxr_false;
}


jxr_bool_t jxr_encoder_need_buffer_size(jxr_encoder_t* p, int* size) {
	return jxr_false;
}
jxr_bool_t jxr_encoder_encode(jxr_encoder_t* p, char* buf, int buf_len, int* size) {
	return jxr_false;
}
