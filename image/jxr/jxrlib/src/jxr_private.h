// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef JXR_PRIVATE_H_
#define JXR_PRIVATE_H_

#include "jxr.h"
#include <JXRGlue.h>

#ifdef  __cplusplus
extern "C" {
#endif

jxr_bool_t jxr_format_guid_valid(
	const PKPixelFormatGUID* fmt
);

jxr_bool_t jxr_parse_format_guid(
	const PKPixelFormatGUID* fmt,
	int* channels, int* depth,
	jxr_data_type_t* type
);

jxr_bool_t jxr_golden_format(
	int channels, int depth,
	jxr_data_type_t type,
	const PKPixelFormatGUID** fmt
);

ERR CreateWS_Discard(
	struct WMPStream** ppWS
);

#ifdef  __cplusplus
} // extern "C"
#endif
#endif // JXR_PRIVATE_H_
