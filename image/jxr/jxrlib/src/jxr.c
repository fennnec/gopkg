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
	jxr_bool_t rv = jxr_false;
	ERR err = WMP_errSuccess;
	PKFactory* pFactory = NULL;
	PKImageDecode* pDecoder = NULL;
	struct WMPStream* pStream = NULL;

	if(data == NULL || size <= 0) return jxr_false;
	if(!width || !height || !channels || !depth) return jxr_false;
	if(!type) return jxr_false;

	err = PKCreateFactory(&pFactory, PK_SDK_VERSION);
	if(Failed(err)) goto Cleanup;
	err = pFactory->CreateStreamFromMemory(&pStream, (void*)data, size);
	if(Failed(err)) goto Cleanup;

	err = PKImageDecode_Create_WMP(&pDecoder);
	if(Failed(err)) goto Cleanup;
	err = PKImageDecode_Initialize_WMP(pDecoder, pStream);
	if(Failed(err)) goto Cleanup;

	*width = (int)(pDecoder->uWidth);
	*height = (int)(pDecoder->uHeight);

	rv = jxr_parse_format_guid(&(pDecoder->guidPixFormat), channels, depth, type);
	if(!rv) goto Cleanup;

Cleanup:

	if(pDecoder != NULL) {
		PKImageDecode_Release(&pDecoder);
		pDecoder = NULL;
	}
	if(pStream != NULL) {
		pStream->Close(&pStream);
		pStream = NULL;
	}
	if(pFactory != NULL) {
		PKCreateFactory_Release(&pFactory);
		pFactory = NULL;
	}

	if(err != WMP_errSuccess) {
		return jxr_false;
	}
	return rv;
}

jxr_bool_t jxr_decode(
	char* buf, int stride, const char* data, int size,
	int* width, int* height, int* channels, int* depth,
	jxr_data_type_t* type
) {
	jxr_bool_t rv = jxr_false;
	ERR err = WMP_errSuccess;
	PKFactory* pFactory = NULL;
	PKImageDecode* pDecoder = NULL;
	struct WMPStream* pStream = NULL;
	const PKPixelFormatGUID* fmt;
	PKRect rect;

	if(buf == NULL) return jxr_false;
	if(data == NULL || size <= 0) return jxr_false;
	if(!width || !height || !channels || !depth) return jxr_false;
	if(!type) return jxr_false;

	err = PKCreateFactory(&pFactory, PK_SDK_VERSION);
	if(Failed(err)) goto Cleanup;
	err = pFactory->CreateStreamFromMemory(&pStream, (void*)data, size);
	if(Failed(err)) goto Cleanup;

	err = PKImageDecode_Create_WMP(&pDecoder);
	if(Failed(err)) goto Cleanup;

	// parse image info
	err = PKImageDecode_Initialize_WMP(pDecoder, pStream);
	if(Failed(err)) goto Cleanup;
	*width = (int)(pDecoder->uWidth);
	*height = (int)(pDecoder->uHeight);

	// lookup format info
	rv = jxr_parse_format_guid(&(pDecoder->guidPixFormat), channels, depth, type);
	if(!rv) goto Cleanup;

	// set stride size
	if(stride <= 0) {
		stride = (*width)*(*channels)*(*depth)/8;
	}
	if(stride < ((*width)*(*channels)*(*depth)/8)) {
		rv = jxr_false;
		goto Cleanup;
	}

	// decode image data
	rect.X = 0;
	rect.Y = 0;
	rect.Width = *width;
	rect.Height = *height;
	pDecoder->WMP.wmiI.bRGB = 1; // use RGB order
	err = pDecoder->Copy(pDecoder, &rect, buf, stride);
	if(Failed(err)) goto Cleanup;

Cleanup:

	if(pDecoder != NULL) {
		PKImageDecode_Release(&pDecoder);
		pDecoder = NULL;
	}
	if(pStream != NULL) {
		pStream->Close(&pStream);
		pStream = NULL;
	}
	if(pFactory != NULL) {
		PKCreateFactory_Release(&pFactory);
		pFactory = NULL;
	}

	if(err != WMP_errSuccess) {
		return jxr_false;
	}
	return rv;
}

int jxr_encode_len(
	const char* data, int stride,
	int width, int height, int channels, int depth,
	int quality, jxr_data_type_t type
) {
	return -1;
}

jxr_bool_t jxr_encode(
	char* buf, int buf_len, const char* data, int stride,
	int width, int height, int channels, int depth,
	int quality, int width_step,
	jxr_data_type_t type,
	int* size
) {
	return jxr_false;
}
