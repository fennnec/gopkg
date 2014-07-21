// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "rawp.h"
#include "rawp_crc32.h"
#include "rawp-snappy.h"

#include <stdlib.h>
#include <string.h>

static bool rawpIsValidChannels(uint8_t channels) {
	return channels == 1 || channels == 3 || channels == 4;
}

static bool rawpIsValidDepth(uint8_t depth) {
	return depth == 8 || depth == 16 || depth == 32 || depth == 64;
}

static bool rawpIsValidDataType(uint8_t x) {
	return x == kRawPDataType_UInt || x == kRawPDataType_Int || x == kRawPDataType_Float;
}

static bool rawpIsValidHeader(const RawPHeader* hdr) {
	if(hdr->Sig[0] != 'R' || hdr->Sig[1] != 'a' || hdr->Sig[2] != 'w' || hdr->Sig[3] != 'P') {
		return 0;
	}
	if(hdr->Magic != kRawPMagic) {
		return 0;
	}

	if(hdr->Width <= 0 || hdr->Height <= 0) {
		return 0;
	}
	if(!rawpIsValidChannels(hdr->Channels)) {
		return 0;
	}
	if(!rawpIsValidDepth(hdr->Depth)) {
		return 0;
	}
	if(!rawpIsValidDataType(hdr->DataType)) {
		return 0;
	}

	if(hdr->UseSnappy != 0 && hdr->UseSnappy != 1) {
		return 0;
	}
	if(hdr->DataSize <= 0) {
		return 0;
	}

	// check type more ...
	if(hdr->Depth == 8 || hdr->Depth == 16) {
		if(hdr->DataType != kRawPDataType_UInt && hdr->DataType != kRawPDataType_Int) {
			return 0;
		}
	}

	// check data size more ...
	if(hdr->UseSnappy) {
		size_t result;
		if(!rawp::snappy::GetUncompressedLength((const char*)hdr->Data, hdr->DataSize, &result)) {
			return 0;
		}
		if(result != hdr->Width*hdr->Height*hdr->Channels*hdr->Depth/8) {
			return 0;
		}
	} else {
		if(hdr->DataSize != hdr->Width*hdr->Height*hdr->Channels*hdr->Depth/8) {
			return 0;
		}
	}

	// Avoid check crc32
	return true;
}

static bool rwapCheckCRC32(const RawPHeader* hdr) {
	return rawpHashCRC32((const char*)hdr->Data, hdr->DataSize) == hdr->DataCheckSum;
}

int rawpDecodeHeader(
	const uint8_t* data, int data_size,
	RawPHeader* hdr
) {
	if(data == NULL || data_size < kRawPHeaderSize || hdr == NULL) {
		return 0;
	}
	// reader header
	memcpy(hdr, data, kRawPHeaderSize);
	hdr->Data = (uint8_t*)data + kRawPHeaderSize;

	// check header
	if(!rawpIsValidHeader(hdr)) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
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
	if(!rwapCheckCRC32(hdr)) {
		return 0;
	}

	if(hdr->UseSnappy) {
		if(!rawp::snappy::RawUncompress((const char*)hdr->Data, (size_t)hdr->DataSize, (char*)output)) {
			return 0;
		}
	} else {
		memcpy(output, hdr->Data, hdr->Width*hdr->Height*hdr->Channels*hdr->Depth/8);
	}
	return 1;
}

int rawpEncodeInit(
	const uint8_t* pix, int width, int height,
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

	if(opt != NULL) {
		if(opt->UseSnappy != 0 && opt->UseSnappy != 1) {
			return 0;
		}
	}

	memset(ctx, 0, sizeof(*ctx));

	ctx->Header.Sig[0] = 'R'; // RawP
	ctx->Header.Sig[1] = 'a';
	ctx->Header.Sig[2] = 'w';
	ctx->Header.Sig[3] = 'P';
	ctx->Header.Magic = kRawPMagic;
	ctx->Header.Width = width;
	ctx->Header.Height = height;
	ctx->Header.Channels = channels;
	ctx->Header.Depth = depth;
	ctx->Header.DataType = data_type;
	ctx->Header.UseSnappy = (opt != NULL)? opt->UseSnappy: 0;
	ctx->Header.DataSize = width*height*channels*depth/8;

	ctx->Pix = (uint8_t*)pix;
	ctx->MaxEncodedLength = kRawPHeaderSize + (
		ctx->Header.UseSnappy? rawp::snappy::MaxCompressedLength(ctx->Header.DataSize):
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

	// write data
	uint8_t* pData = output + kRawPHeaderSize;
	if(hdr->UseSnappy) {
		size_t output_length;
		rawp::snappy::RawCompress(
			(const char*)ctx->Pix, hdr->Width*hdr->Height*hdr->Channels*hdr->Depth/8,
			(char*)pData, &output_length
		);
		hdr->DataSize = output_length;
	} else {
		ctx->Header.DataSize = hdr->Width*hdr->Height*hdr->Channels*hdr->Depth/8;
		memcpy(pData, ctx->Pix, hdr->DataSize);
	}

	// write crc32
	hdr->DataCheckSum = rawpHashCRC32((const char*)pData, hdr->DataSize);

	// write header
	memcpy(output, hdr, kRawPHeaderSize);

	// OK
	return size_t(kRawPHeaderSize+hdr->DataSize);
}

