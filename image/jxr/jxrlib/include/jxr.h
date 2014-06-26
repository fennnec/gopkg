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

typedef enum jxr_data_type_t {
	jxr_unsigned = 0,
	jxr_signed = 1,
	jxr_float = 2,
} jxr_data_type_t;

// ----------------------------------------------------------------------------
// decode/encode
// if(buf == NULL && buf_len == 0) return len(?);
// ----------------------------------------------------------------------------

int jxr_encode(
	char* buf, int buf_len, const char* data, int size,
	int width, int height, int channels, int depth,
	int quality, int width_step,
	jxr_data_type_t type
);

int jxr_decode(
	char* buf, int buf_len, const char* data, int size,
	int* width, int* height, int* channels, int* depth,
	jxr_data_type_t* type
);

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------

#ifdef  __cplusplus
} // extern "C"
#endif
#endif // JXR_H_
