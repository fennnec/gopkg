// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "raw.h"
#include "raw_crc32.h"
#include "raw-snappy.h"

#include <stdlib.h>
#include <string.h>

static bool rawpIsValidDataType(uint8_t data_type) {
	if(data_type == kRawPDataType_UInt) return true;
	if(data_type == kRawPDataType_Int) return true;
	if(data_type == kRawPDataType_Float) return true;
	return false;
}

static bool rawpIsValidDepth(uint8_t depth) {
	if(depth == 8) return true;
	if(depth == 16) return true;
	if(depth == 32) return true;
	if(depth == 64) return true;
	return false;
}

int rawpDecodeHeader(
	const uint8_t* data, int data_size,
	RawPHeader* hdr
) {
	if(data == NULL || data_size < kRawPHeaderSize || hdr == NULL) {
		return 0;
	}
	memset(hdr, 0, sizeof(*hdr));

	// reader header
	const uint8_t* p = data;
	memcpy(&hdr->Sig, p, sizeof(hdr->Sig));
	p += sizeof(hdr->Sig);
	memcpy(&hdr->UseRC32, p, sizeof(hdr->UseRC32));
	p += sizeof(hdr->UseRC32);
	memcpy(&hdr->UseSnappy, p, sizeof(hdr->UseSnappy));
	p += sizeof(hdr->UseSnappy);
	memcpy(&hdr->DataType, p, sizeof(hdr->DataType));
	p += sizeof(hdr->DataType);
	memcpy(&hdr->Depth, p, sizeof(hdr->Depth));
	p += sizeof(hdr->Depth);
	memcpy(&hdr->Channels, p, sizeof(hdr->Channels));
	p += sizeof(hdr->Channels);
	memcpy(&hdr->Width, p, sizeof(hdr->Width));
	p += sizeof(hdr->Width);
	memcpy(&hdr->Height, p, sizeof(hdr->Height));
	p += sizeof(hdr->Height);
	memcpy(&hdr->DataSize, p, sizeof(hdr->DataSize));
	p += sizeof(hdr->DataSize);
	hdr->Data = (uint8_t*)p;

	// check header
	if(hdr->Sig[0] != 'R' || hdr->Sig[1] != 'a' || hdr->Sig[2] != 'w' || hdr->Sig[3] != 'P') {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
	if(hdr->Magic != kRawPMagic) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
	if(hdr->UseRC32 != 0 && hdr->UseRC32 != 1) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
	if(hdr->UseSnappy != 0 && hdr->UseSnappy != 1) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
	if(!rawpIsValidDataType(hdr->DataType)) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
	if(!rawpIsValidDepth(hdr->Depth)) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
	if(hdr->Channels <= 0) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
	if(hdr->Width <= 0) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
	if(hdr->Height <= 0) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
	if(hdr->DataSize <= 0) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}

	// check type more ...
	if(hdr->Depth == 8 || hdr->Depth == 16) {
		if(hdr->DataType == kRawPDataType_Float) {
			memset(hdr, 0, sizeof(*hdr));
			return 0;
		}
	}
	// check data size more ...
	if(!hdr->UseSnappy) {
		size_t result;
		if(!raw::snappy::GetUncompressedLength((const char*)hdr->Data, hdr->DataSize, &result)) {
			memset(hdr, 0, sizeof(*hdr));
			return 0;
		}
		if(result != hdr->Width*hdr->Height*hdr->Channels*hdr->Depth/8) {
			memset(hdr, 0, sizeof(*hdr));
			return 0;
		}
	} else {
		if(hdr->DataSize != hdr->Width*hdr->Height*hdr->Channels*hdr->Depth/8) {
			memset(hdr, 0, sizeof(*hdr));
			return 0;
		}
	}

	// read checksum
	p += hdr->DataSize;
	memcpy(&hdr->CheckSum, p, sizeof(hdr->CheckSum));

	// OK
	return 1;
}

int rawpDecode(
	const uint8_t* data, int data_size,
	uint8_t* output, int output_size,
	RawPHeader* hdr
) {
	if(data == NULL || data_size < kRawPHeaderSize) {
		return 0;
	}
	if(output == NULL || output_size <= 0) {
		return 0;
	}
	if(hdr == NULL) {
		return 0;
	}

	if(!rawpDecodeHeader(data, data_size, hdr)) {
		return 0;
	}
	if(data_size < (kRawPHeaderSize+int(hdr->DataSize))) {
		return 0;
	}
	if(output_size < hdr->Width*hdr->Height*hdr->Channels*hdr->Depth/8) {
		return 0;
	}

	if(hdr->UseRC32) {
		if(rawHashCRC32((const char*)data, (kRawPHeaderSize+hdr->DataSize)-4) != hdr->CheckSum) {
			return NULL;
		}
	}
	if(hdr->UseSnappy) {
		if(!raw::snappy::RawUncompress((const char*)hdr->Data, (size_t)hdr->DataSize, (char*)output)) {
			return 0;
		}
	} else {
		memcpy(output, hdr->Data, hdr->Width*hdr->Height*hdr->Channels*hdr->Depth/8);
	}
	return 1;
}

int rawpEncodeInit(
	const uint8_t* pix, int width, int height, int stride,
	int channels, int depth, int data_type,
	const RawPEncodeOptions* opt,
	RawPEncodeContext* ctx
) {
	if(pix == NULL || width <= 0 || height <= 0) {
		return 0;
	}
	if(channels <= 0 || !rawpIsValidDepth(depth) || !rawpIsValidDataType(data_type)) {
		return 0;
	}
	if(ctx == NULL) {
		return 0;
	}

	memset(ctx, 0, sizeof(*ctx));

	ctx->Header.Sig[0] = 'R'; // RawP
	ctx->Header.Sig[1] = 'a';
	ctx->Header.Sig[2] = 'w';
	ctx->Header.Sig[3] = 'P';
	ctx->Header.Magic = kRawPMagic;
	ctx->Header.UseRC32 = (opt != NULL)? opt->UseRC32: false;
	ctx->Header.UseSnappy = (opt != NULL)? opt->UseSnappy: false;
	ctx->Header.DataType = data_type;
	ctx->Header.Depth = depth;
	ctx->Header.Channels = channels;
	ctx->Header.Width = width;
	ctx->Header.Height = height;
	ctx->Header.DataSize = width*height*channels*depth/8;
	ctx->Header.CheckSum = kRawPMagic;

	ctx->Pix = (uint8_t*)pix;
	ctx->MaxEncodedLength = kRawPHeaderSize + (
		ctx->Header.UseSnappy? raw::snappy::MaxCompressedLength(ctx->Header.DataSize):
		ctx->Header.DataSize
	);

	return 1;
}

size_t rawpEncode(
	RawPEncodeContext* ctx,
	uint8_t* output
) {
	if(ctx == NULL || output == NULL) {
		return 0;
	}
	RawPHeader* hdr = &(ctx->Header);

	// write header
	uint8_t* p = output;
	memcpy(p, &hdr->Sig, sizeof(hdr->Sig));
	p += sizeof(hdr->Sig);
	memcpy(p, &hdr->UseRC32, sizeof(hdr->UseRC32));
	p += sizeof(hdr->UseRC32);
	memcpy(p, &hdr->UseSnappy, sizeof(hdr->UseSnappy));
	p += sizeof(hdr->UseSnappy);
	memcpy(p, &hdr->DataType, sizeof(hdr->DataType));
	p += sizeof(hdr->DataType);
	memcpy(p, &hdr->Depth, sizeof(hdr->Depth));
	p += sizeof(hdr->Depth);
	memcpy(p, &hdr->Channels, sizeof(hdr->Channels));
	p += sizeof(hdr->Channels);
	memcpy(p, &hdr->Width, sizeof(hdr->Width));
	p += sizeof(hdr->Width);
	memcpy(p, &hdr->Height, sizeof(hdr->Height));
	p += sizeof(hdr->Height);

	// write data
	uint8_t* pDataBuf = p + sizeof(hdr->DataSize);
	if(hdr->UseSnappy) {
		size_t output_length;
		raw::snappy::RawCompress((const char*)ctx->Pix, hdr->DataSize, (char*)pDataBuf, &output_length);
		hdr->DataSize = output_length;
	} else {
		memcpy(pDataBuf, ctx->Pix, hdr->DataSize);
	}
	memcpy(p, &hdr->DataSize, sizeof(hdr->DataSize));
	p += sizeof(hdr->DataSize);
	hdr->Data = (uint8_t*)p;
	p += hdr->DataSize;

	// write crc32
	if(hdr->UseRC32) {
		hdr->CheckSum = rawHashCRC32((const char*)output, p-output);
	}
	memcpy(p, &hdr->CheckSum, sizeof(hdr->CheckSum));
	p += sizeof(hdr->CheckSum);

	// OK
	return size_t(p-output);
}

