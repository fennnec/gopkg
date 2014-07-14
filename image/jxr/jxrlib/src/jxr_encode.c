// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "jxr.h"
#include "jxr_private.h"

struct jxr_encoder_t {
	const void*       pType;
	PKFactory*        pFactory;
	PKImageEncode*    pEncoder;
	struct WMPStream* pStream;
	CWMIStrCodecParam wmiSCP;
	const char*       data;
	int               dataSize;
	int               stride;
	int               width;
	int               height;
	int               channels;
	int               depth;
	int               quality;
	jxr_data_type_t   dataType;
	size_t            destSize;
};

static const char jxr_encoder_type[] = "jxr_encoder_t";

jxr_encoder_t* jxr_encoder_new() {
	jxr_encoder_t* p = (jxr_encoder_t*)calloc(1, sizeof(*p));
	if(!p) return NULL;

	p->pType = jxr_encoder_type;
	if(Failed(PKCreateFactory(&p->pFactory, PK_SDK_VERSION))) {
		jxr_encoder_delete(p);
		return NULL;
	}
	if(Failed(PKImageEncode_Create_WMP(&p->pEncoder))) {
		jxr_encoder_delete(p);
		return NULL;
	}
	return p;
}

void jxr_encoder_delete(jxr_encoder_t* p) {
	if(p == NULL || p->pType != jxr_encoder_type) {
		fprintf(stderr, "jxr: jxr_encoder_delete, invalid jxr_decoder_t type!");
		abort();
	}

	if(p->pEncoder != NULL) {
		p->pEncoder->Release(&p->pEncoder);
		p->pEncoder = NULL;
	}
	if(p->pStream != NULL) {
		p->pStream->Close(&p->pStream);
		p->pStream = NULL;
	}
	if(p->pFactory != NULL) {
		p->pFactory->Release(&p->pFactory);
		p->pFactory = NULL;
	}
	p->pType = NULL;
}

jxr_bool_t jxr_encoder_init(jxr_encoder_t* p,
	const char* data, int size, int stride,
	int width, int height, int channels, int depth,
	int quality, jxr_data_type_t type
) {
	const PKPixelFormatGUID* fmt = NULL;

	if(p == NULL || p->pType != jxr_encoder_type) {
		fprintf(stderr, "jxr: jxr_encoder_init, invalid jxr_decoder_t type!");
		abort();
	}

	// lookup best match format
	if(!jxr_golden_format(channels, depth, type, &fmt)) {
		return jxr_false;
	}
	if(Failed(p->pEncoder->SetPixelFormat(p->pEncoder, *fmt))) {
		return jxr_false;
	}

	// set image size
	if(Failed(p->pEncoder->SetSize(p->pEncoder, width, height))) {
		return jxr_false;
	}

	p->data = data;
	p->dataSize = size;
	p->stride = stride;

	p->width = width;
	p->height = height;
	p->channels = channels;
	p->depth = depth;
	p->quality = quality;
	p->dataType = type;
	p->destSize = 0;

	return jxr_true;
}

jxr_bool_t jxr_encoder_need_buffer_size(jxr_encoder_t* p, int* size) {
	struct WMPStream* pNilStream = NULL;

	if(p->destSize == 0) {
		// create discard stream
		if(Failed(CreateWS_Discard(&pNilStream))) {
			return jxr_false;
		}
		if(Failed(p->pEncoder->Initialize(p->pEncoder, pNilStream, &p->wmiSCP, sizeof(p->wmiSCP)))) {
			pNilStream->Close(&pNilStream);
			return jxr_false;
		}

		{
			// how to set quality ?
		}

		// try encode
		if(Failed(p->pEncoder->WritePixels(p->pEncoder, p->height, (U8*)p->data, p->stride))) {
			pNilStream->Close(&pNilStream);
			return jxr_false;
		}
		if(Failed(pNilStream->GetPos(pNilStream, &p->destSize))) {
			pNilStream->Close(&pNilStream);
			return jxr_false;
		}

		pNilStream->Close(&pNilStream);
	}
	if(size != NULL) {
		*size = p->destSize;
	}
	return jxr_true;
}

jxr_bool_t jxr_encoder_encode(jxr_encoder_t* p, char* buf, int buf_len, int* size) {
	// create new stream
	if(p->pStream == NULL) {
		if(Failed(p->pFactory->CreateStreamFromMemory(&p->pStream, (void*)buf, buf_len))) {
			return jxr_false;
		}
	}

	// set output steam
	if(Failed(p->pEncoder->Initialize(p->pEncoder, p->pStream, &p->wmiSCP, sizeof(p->wmiSCP)))) {
		return jxr_false;
	}

	// how to set quality ?
	{
		//
	}

	// encode
	if(Failed(p->pEncoder->WritePixels(p->pEncoder, p->height, (U8*)p->data, p->stride))) {
		return jxr_false;
	}
	if(Failed(p->pStream->GetPos(p->pStream, &p->destSize))) {
		return jxr_false;
	}
	if(size != NULL) {
		*size = p->destSize;
	}

	return jxr_true;
}
