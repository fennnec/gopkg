// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "jxr.h"
#include "jxr_private.h"

jxr_bool_t jxr_decode_config(
	const char* data, int size,
	int* width, int* height, int* channels, int* depth,
	jxr_data_type_t* type
) {
	jxr_decoder_t* p = jxr_decoder_new();
	if(!p) return jxr_false;

	if(!jxr_decoder_init(p, data, size)) {
		jxr_decoder_delete(p);
		return jxr_false;
	}

	if(width != NULL) *width = jxr_decoder_width(p);
	if(height != NULL) *height = jxr_decoder_height(p);
	if(channels != NULL) *channels = jxr_decoder_channels(p);
	if(depth != NULL) *depth = jxr_decoder_depth(p);
	if(type != NULL) *type = jxr_decoder_data_type(p);

	jxr_decoder_delete(p);
	return jxr_true;
}

jxr_bool_t jxr_decode(
	char* buf, int buf_len, int stride, const char* data, int size,
	int* width, int* height, int* channels, int* depth,
	jxr_data_type_t* type
) {
	jxr_decoder_t* p = jxr_decoder_new();
	if(!p) return jxr_false;

	if(!jxr_decoder_init(p, data, size)) {
		jxr_decoder_delete(p);
		return jxr_false;
	}

	// set stride size
	if(stride <= 0) {
		stride = jxr_decoder_width(p)*jxr_decoder_channels(p)*jxr_decoder_depth(p)/8;
	}
	if(stride < (jxr_decoder_width(p)*jxr_decoder_channels(p)*jxr_decoder_depth(p)/8)) {
		jxr_decoder_delete(p);
		return jxr_false;
	}

	// check buffer size
	if(buf_len < stride*jxr_decoder_height(p)) {
		jxr_decoder_delete(p);
		return jxr_false;
	}

	// decode all
	if(!jxr_decoder_decode(p, NULL, buf, stride)) {
		jxr_decoder_delete(p);
		return jxr_false;
	}

	if(width != NULL) *width = jxr_decoder_width(p);
	if(height != NULL) *height = jxr_decoder_height(p);
	if(channels != NULL) *channels = jxr_decoder_channels(p);
	if(depth != NULL) *depth = jxr_decoder_depth(p);
	if(type != NULL) *type = jxr_decoder_data_type(p);

	jxr_decoder_delete(p);
	return jxr_true;
}

jxr_bool_t jxr_encode_len(
	const char* data, int data_size, int stride,
	int width, int height, int channels, int depth,
	int quality, jxr_data_type_t type,
	int* size
) {
	jxr_encoder_t* p = jxr_encoder_new();
	if(!p) return jxr_false;

	if(!jxr_encoder_init(
		p, data, data_size, stride,
		width, height, channels, depth,
		quality, type
	)) {
		jxr_encoder_delete(p);
		return jxr_false;
	}
	if(!jxr_encoder_need_buffer_size(p, size)) {
		jxr_encoder_delete(p);
		return jxr_false;
	}

	jxr_encoder_delete(p);
	return jxr_true;
}

jxr_bool_t jxr_encode(
	char* buf, int buf_len,
	const char* data, int data_size, int stride,
	int width, int height, int channels, int depth,
	int quality, jxr_data_type_t type,
	int* size
) {
	int need_buf_size;

	jxr_encoder_t* p = jxr_encoder_new();
	if(!p) return jxr_false;

	if(!jxr_encoder_init(
		p, data, data_size, stride,
		width, height, channels, depth,
		quality, type
	)) {
		jxr_encoder_delete(p);
		return jxr_false;
	}
	if(!jxr_encoder_need_buffer_size(p, &need_buf_size)) {
		jxr_encoder_delete(p);
		return jxr_false;
	}

	if(buf_len < need_buf_size) {
		jxr_encoder_delete(p);
		return jxr_false;
	}
	if(!jxr_encoder_encode(p, buf, buf_len, size)) {
		jxr_encoder_delete(p);
		return jxr_false;
	}

	jxr_encoder_delete(p);
	return jxr_true;
}

