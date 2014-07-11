// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef JXR_H_
#define JXR_H_

#ifdef  __cplusplus
extern "C" {
#endif

// ----------------------------------------------------------------------------
// Exported types
// ----------------------------------------------------------------------------

typedef struct jxr_decoder_t jxr_decoder_t;
typedef struct jxr_encoder_t jxr_encoder_t;

typedef enum jxr_bool_t {
	jxr_true  = 1,
	jxr_false = 0,
} jxr_bool_t;

typedef enum jxr_data_type_t {
	jxr_unsigned = 0,
	jxr_signed = 1,
	jxr_float = 2,
} jxr_data_type_t;

// ----------------------------------------------------------------------------
// decode/encode simple api
// ----------------------------------------------------------------------------

jxr_bool_t jxr_decode_config(
	const char* data, int size,
	int* width, int* height, int* channels, int* depth,
	jxr_data_type_t* type
);

jxr_bool_t jxr_decode(
	char* buf, int stride, const char* data, int size,
	int* width, int* height, int* channels, int* depth,
	jxr_data_type_t* type
);

int jxr_encode_len(
	const char* pix, int stride,
	int width, int height, int channels, int depth,
	int quality, jxr_data_type_t type
);

jxr_bool_t jxr_encode(
	char* buf, int buf_len, const char* pix, int stride,
	int width, int height, int channels, int depth,
	int quality, jxr_data_type_t type,
	int* size
);

// ----------------------------------------------------------------------------
// decoder
// ----------------------------------------------------------------------------

jxr_decoder_t* jxr_decoder_new();
void jxr_decoder_delete(jxr_decoder_t* p);

jxr_bool_t jxr_decoder_init(jxr_decoder_t* p, const char* data, int siz);

int jxr_decoder_width(jxr_decoder_t* p);
int jxr_decoder_height(jxr_decoder_t* p);
int jxr_decoder_channels(jxr_decoder_t* p);
int jxr_decoder_depth(jxr_decoder_t* p);
int jxr_decoder_data_type(jxr_decoder_t* p);

jxr_bool_t jxr_decoder_decode(jxr_decoder_t* p, char* buffer, int stride);

// ----------------------------------------------------------------------------
// encoder
// ----------------------------------------------------------------------------

jxr_encoder_t* jxr_encoder_new();
void jxr_encoder_delete(jxr_encoder_t* p);

jxr_bool_t jxr_encoder_init(jxr_encoder_t* p,
	const char* data, int stride,
	int width, int height, int channels, int depth,
	int quality, jxr_data_type_t type
);

int jxr_decoder_need_buffer_size(jxr_encoder_t* p);
jxr_bool_t jxr_decoder_encode(jxr_encoder_t* p, char* buffer);

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------

#ifdef  __cplusplus
} // extern "C"
#endif
#endif // JXR_H_
