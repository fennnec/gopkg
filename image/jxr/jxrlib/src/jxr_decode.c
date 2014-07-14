// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "jxr.h"
#include "jxr_private.h"

struct jxr_decoder_t {
	const void*       pType;
	PKFactory*        pFactory;
	PKImageDecode*    pDecoder;
	struct WMPStream* pStream;
	int               width;
	int               height;
	int               channels;
	int               depth;
	jxr_data_type_t   dataType;
};

static const char* jxr_decoder_type = "jxr_decoder_t";

jxr_decoder_t* jxr_decoder_new() {
	jxr_decoder_t* p = (jxr_decoder_t*)calloc(1, sizeof(*p));
	if(!p) return NULL;

	p->pType = jxr_decoder_type;
	if(Failed(PKCreateFactory(&p->pFactory, PK_SDK_VERSION))) {
		jxr_decoder_delete(p);
		return NULL;
	}
	if(Failed(PKImageDecode_Create_WMP(&p->pDecoder))) {
		jxr_decoder_delete(p);
		return NULL;
	}
	return p;
}

void jxr_decoder_delete(jxr_decoder_t* p) {
	if(p == NULL || p->pType != jxr_decoder_type) {
		fprintf(stderr, "jxr: jxr_decoder_delete, invalid jxr_decoder_t type!");
		abort();
	}

	if(p->pDecoder != NULL) {
		p->pDecoder->Release(&p->pDecoder);
		p->pDecoder = NULL;
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

jxr_bool_t jxr_decoder_init(jxr_decoder_t* p, const char* data, int size) {
	if(p == NULL || p->pType != jxr_decoder_type) {
		fprintf(stderr, "jxr: jxr_decoder_init, invalid jxr_decoder_t type!");
		abort();
	}

	// close old stream
	if(p->pStream != NULL) {
		p->pStream->Close(&p->pStream);
		p->pStream = NULL;
	}
	p->width = 0;
	p->height = 0;
	p->channels = 0;
	p->depth = 0;
	p->dataType = jxr_unsigned;

	// create new stream
	if(Failed(p->pFactory->CreateStreamFromMemory(&p->pStream, (void*)data, size))) {
		return jxr_false;
	}
	if(Failed(p->pDecoder->Initialize(p->pDecoder, p->pStream))) {
		return jxr_false;
	}

	// parse image format
	if(!jxr_parse_format_guid(&(p->pDecoder->guidPixFormat), &p->channels, &p->depth, &p->dataType)) {
		return jxr_false;
	}
	p->width = (int)(p->pDecoder->uWidth);
	p->height = (int)(p->pDecoder->uHeight);
	return jxr_true;
}

int jxr_decoder_width(jxr_decoder_t* p) {
	if(p == NULL || p->pType != jxr_decoder_type) {
		fprintf(stderr, "jxr: jxr_decoder_width, invalid jxr_decoder_t type!");
		abort();
	}
	return p->width;
}

int jxr_decoder_height(jxr_decoder_t* p) {
	if(p == NULL || p->pType != jxr_decoder_type) {
		fprintf(stderr, "jxr: jxr_decoder_height, invalid jxr_decoder_t type!");
		abort();
	}
	return p->height;
}

int jxr_decoder_channels(jxr_decoder_t* p) {
	if(p == NULL || p->pType != jxr_decoder_type) {
		fprintf(stderr, "jxr: jxr_decoder_channels, invalid jxr_decoder_t type!");
		abort();
	}
	return p->channels;
}

int jxr_decoder_depth(jxr_decoder_t* p) {
	if(p == NULL || p->pType != jxr_decoder_type) {
		fprintf(stderr, "jxr: jxr_decoder_depth, invalid jxr_decoder_t type!");
		abort();
	}
	return p->depth;
}

int jxr_decoder_data_type(jxr_decoder_t* p) {
	if(p == NULL || p->pType != jxr_decoder_type) {
		fprintf(stderr, "jxr: jxr_decoder_data_type, invalid jxr_decoder_t type!");
		abort();
	}
	return p->dataType;
}

jxr_bool_t jxr_decoder_decode(jxr_decoder_t* p, const jxr_rect_t* r, char* buf, int stride) {
	PKRect rect;
	ERR err = WMP_errSuccess;

	if(p == NULL || p->pType != jxr_decoder_type) {
		fprintf(stderr, "jxr: jxr_decoder_decode, invalid jxr_decoder_t type!");
		abort();
	}

	if(buf == NULL) {
		return jxr_false;
	}
	if(p->width <= 0 || p->height <= 0 || p->channels <= 0 || p->depth <= 0) {
		return jxr_false;
	}

	// set stride size
	if(stride <= 0) {
		stride = (p->width)*(p->channels)*(p->depth)/8;
	}
	if(stride < ((p->width)*(p->channels)*(p->depth)/8)) {
		return jxr_false;
	}

	if(r != NULL) {
		rect.X = r->x;
		rect.Y = r->y;
		rect.Width = r->width;
		rect.Height = r->height;
	} else {
		rect.X = 0;
		rect.Y = 0;
		rect.Width = p->width;
		rect.Height = p->height;
	}

	// decode image data
	p->pDecoder->WMP.wmiI.bRGB = 1; // use RGB order
	err = p->pDecoder->Copy(p->pDecoder, &rect, buf, stride);
	return Failed(err)? jxr_false: jxr_true;
}
