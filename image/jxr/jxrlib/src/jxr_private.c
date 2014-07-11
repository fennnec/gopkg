// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "jxr_private.h"


// fotmat info list
static const struct {
	const char* name;
	const PKPixelFormatGUID* fmt;
	int channels;
	int depth;
	jxr_data_type_t type;
	jxr_bool_t golden;
} fmtInfoList[] = {
#define _JXR_FMT_(fmt) #fmt, &fmt
	{ _JXR_FMT_(GUID_PKPixelFormatBlackWhite), 1, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat8bppGray), 1, 8, jxr_unsigned, jxr_true },            // golden Gray

	/* 16bpp formats */
	{ _JXR_FMT_(GUID_PKPixelFormat16bppRGB555), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat16bppRGB565), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat16bppGray), 1, 16, jxr_unsigned, jxr_true },          // golden Gray16

	/* 24bpp formats */
	{ _JXR_FMT_(GUID_PKPixelFormat24bppBGR), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat24bppRGB), 3, 8, jxr_unsigned, jxr_false },           // golden RGB

	/* 32bpp format */
	{ _JXR_FMT_(GUID_PKPixelFormat32bppBGR), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat32bppBGRA), 4, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat32bppPBGRA), 4, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat32bppGrayFloat), 1, 32, jxr_float, jxr_true },        // golden Gray32f
	{ _JXR_FMT_(GUID_PKPixelFormat32bppRGB), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat32bppRGBA), 4, 8, jxr_unsigned, jxr_true  },          // golden RGBA
	{ _JXR_FMT_(GUID_PKPixelFormat32bppPRGBA), 4, 8, jxr_unsigned, jxr_false },

	/* 48bpp format */
	{ _JXR_FMT_(GUID_PKPixelFormat48bppRGBFixedPoint), 3, 16, jxr_signed, jxr_true },   // golden RGB48i

	/* 16bpp format */
	{ _JXR_FMT_(GUID_PKPixelFormat16bppGrayFixedPoint), 1, 16, jxr_signed, jxr_true },  // golden Gray16i

	/* 32bpp format */
	{ _JXR_FMT_(GUID_PKPixelFormat32bppRGB101010), 3, 8, jxr_unsigned, jxr_false },

	/* 48bpp format */
	{ _JXR_FMT_(GUID_PKPixelFormat48bppRGB), 3, 8, jxr_unsigned, jxr_true },            // golden RGB48

	/* 64bpp format */
	{ _JXR_FMT_(GUID_PKPixelFormat64bppRGBA), 4, 16, jxr_unsigned, jxr_true },          // golden RGBA64
	{ _JXR_FMT_(GUID_PKPixelFormat64bppPRGBA), 4, 16, jxr_unsigned, jxr_false },

	/* 96bpp format */
	{ _JXR_FMT_(GUID_PKPixelFormat96bppRGBFixedPoint), 3, 32, jxr_signed, jxr_false },  // golden RGB96i
	{ _JXR_FMT_(GUID_PKPixelFormat96bppRGBFloat), 3, 32, jxr_float, jxr_false },        // golden RGB96f

	/* Floating point scRGB formats */
	{ _JXR_FMT_(GUID_PKPixelFormat128bppRGBAFloat), 4, 32, jxr_float, jxr_true },       // golden RGBA128f
	{ _JXR_FMT_(GUID_PKPixelFormat128bppPRGBAFloat), 4, 32, jxr_float, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat128bppRGBFloat), 4, 32, jxr_float, jxr_false },

	/* CMYK formats. */
	{ _JXR_FMT_(GUID_PKPixelFormat32bppCMYK), 3, 8, jxr_unsigned, jxr_false },

	/* Photon formats */
	{ _JXR_FMT_(GUID_PKPixelFormat64bppRGBAFixedPoint), 4, 16, jxr_signed, jxr_true },  // golden RGBA64i
	{ _JXR_FMT_(GUID_PKPixelFormat64bppRGBFixedPoint), 3, 16, jxr_signed, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat128bppRGBAFixedPoint), 4, 32, jxr_signed, jxr_true }, // golden RGBA128f
	{ _JXR_FMT_(GUID_PKPixelFormat128bppRGBFixedPoint), 3, 32, jxr_signed, jxr_false },

	{ _JXR_FMT_(GUID_PKPixelFormat64bppRGBAHalf), 4, 16, jxr_float, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat64bppRGBHalf), 3, 16, jxr_float, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat48bppRGBHalf), 3, 16, jxr_float, jxr_false },

	{ _JXR_FMT_(GUID_PKPixelFormat32bppRGBE), 3, 8, jxr_unsigned, jxr_false },

	{ _JXR_FMT_(GUID_PKPixelFormat16bppGrayHalf), 1, 16, jxr_float, jxr_true },         // golden Gray16f
	{ _JXR_FMT_(GUID_PKPixelFormat32bppGrayFixedPoint), 1, 32, jxr_signed, jxr_true },  // golden Gray32i

	/* YCrCb  from Advanced Profile */
	{ _JXR_FMT_(GUID_PKPixelFormat12bppYCC420), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat16bppYCC422), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat20bppYCC422), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat32bppYCC422), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat24bppYCC444), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat30bppYCC444), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat48bppYCC444), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat16bpp48bppYCC444FixedPoint), 3, 8 , jxr_signed, jxr_false },

	{ _JXR_FMT_(GUID_PKPixelFormat20bppYCC420Alpha), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat24bppYCC422Alpha), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat30bppYCC422Alpha), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat48bppYCC422Alpha), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat32bppYCC444Alpha), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat40bppYCC444Alpha), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat64bppYCC444Alpha), 3, 8, jxr_unsigned, jxr_false },
	{ _JXR_FMT_(GUID_PKPixelFormat64bppYCC444AlphaFixedPoint), 3, 8, jxr_signed, jxr_false  },
#undef _JXR_FMT_
};

jxr_bool_t jxr_format_guid_valid(const PKPixelFormatGUID* fmt) {
	int i;
	if(IsEqualGUID(fmt, &GUID_PKPixelFormatDontCare)) {
		return jxr_true;
	}
	for(i = 0; i < sizeof(fmtInfoList)/sizeof(fmtInfoList[0]); ++i) {
		if(IsEqualGUID(fmt, fmtInfoList[i].fmt)) {
			return jxr_true;
		}
	}
	return jxr_false;
}

jxr_bool_t jxr_parse_format_guid(
	const PKPixelFormatGUID* fmt,
	int* channels, int* depth,
	jxr_data_type_t* type
) {
	int i;
	for(i = 0; i < sizeof(fmtInfoList)/sizeof(fmtInfoList[0]); ++i) {
		if(IsEqualGUID(fmt, fmtInfoList[i].fmt)) {
			if(channels != NULL) *channels = fmtInfoList[i].channels;
			if(depth != NULL) *depth = fmtInfoList[i].depth;
			if(type != NULL) *type = fmtInfoList[i].type;
			return jxr_true;
		}
	}
	return jxr_false;
}

jxr_bool_t jxr_golden_format(
	int channels, int depth,
	jxr_data_type_t type,
	const PKPixelFormatGUID** fmt
) {
	int i;
	for(i = 0; i < sizeof(fmtInfoList)/sizeof(fmtInfoList[0]); ++i) {
		if(!fmtInfoList[i].golden) continue;
		if(fmtInfoList[i].channels != channels) continue;
		if(fmtInfoList[i].depth != depth) continue;
		if(fmtInfoList[i].type != type) continue;

		*fmt = fmtInfoList[i].fmt;
		return jxr_true;
	}
	return jxr_false;

}
