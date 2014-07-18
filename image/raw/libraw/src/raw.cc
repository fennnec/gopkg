// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "raw.h"
#include "raw_crc32.h"
#include "raw-snappy.h"

#include <stdlib.h>
#include <string.h>

static bool rawIsValidDataType(uint8_t data_type) {
	if(data_type == kRawDataType_UInt) return true;
	if(data_type == kRawDataType_Int) return true;
	if(data_type == kRawDataType_Float) return true;
	return false;
}

static bool rawIsValidDepth(uint8_t depth) {
	if(depth == 8) return true;
	if(depth == 16) return true;
	if(depth == 32) return true;
	if(depth == 64) return true;
	return false;
}

int rawGetHeader(
	const uint8_t* data, size_t data_size,
	RawHeader* hdr
) {
	if(data == NULL || data_size < kRawHeaderSize || hdr == NULL) {
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
	if(hdr->Sig != kRawMagic) {
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
	if(!rawIsValidDataType(hdr->DataType)) {
		memset(hdr, 0, sizeof(*hdr));
		return 0;
	}
	if(!rawIsValidDepth(hdr->Depth)) {
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
		if(hdr->DataType == kRawDataType_Float) {
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

uint8_t* rawDecode(
	const uint8_t* data, size_t data_size,
	int* width, int* height, int* channels, int* depth, int* data_type
) {
	if(data == NULL || data_size <= 0) {
		return NULL;
	}

	RawHeader hdr;
	if(!rawGetHeader(data, data_size, &hdr)) {
		return NULL;
	}

	if(width != NULL) {
		*width = hdr.Width;
	}
	if(height != NULL) {
		*height = hdr.Height;
	}
	if(channels != NULL) {
		*channels = hdr.Channels;
	}
	if(data_type != NULL) {
		*data_type = hdr.DataType;
	}

	if(hdr.UseRC32) {
		if(rawHashCRC32((const char*)data, data_size) != hdr.CheckSum) {
			return NULL;
		}
	}

	int imageSize = hdr.Width*hdr.Height*hdr.Channels*hdr.Depth/8;
	uint8_t*pix = (uint8_t*)malloc(imageSize);
	if(!pix) {
		return NULL;
	}

	if(hdr.UseSnappy) {
		if(!raw::snappy::RawUncompress((const char*)hdr.Data, (size_t)hdr.DataSize, (char*)pix)) {
			free(pix);
			return NULL;
		}
		return pix;
	} else {
		memcpy(pix, hdr.Data, imageSize);
		return pix;
	}
}

size_t rawEncode(
	const uint8_t* pix, int width, int height, int stride,
	int channels, int depth, int data_type,
	const RawEncodeOptions* opt,
	uint8_t** output
) {
	if(pix == NULL || width <= 0 || height <= 0 || stride < 0) {
		return 0;
	}
	if(channels <= 0 || !rawIsValidDepth(depth) || !rawIsValidDataType(data_type)) {
		return 0;
	}
	if(output == NULL) {
		return 0;
	}

	if(stride == 0) {
		stride = width*channels*depth/8;
	}
	if(stride < width*channels*depth/8) {
		return 0;
	}

	RawHeader hdr[1];
	memset(&hdr, 0, sizeof(hdr));
	hdr->Sig = kRawMagic;
	hdr->UseRC32 = (opt != NULL)? opt->UseRC32: false;
	hdr->UseSnappy = (opt != NULL)? opt->UseSnappy: false;
	hdr->DataType = data_type;
	hdr->Depth = depth;
	hdr->Channels = channels;
	hdr->Width = width;
	hdr->Height = height;
	hdr->DataSize = width*height*channels*depth/8;
	hdr->CheckSum = kRawMagic;

	int bufferSize = kRawHeaderSize + (
		hdr->UseSnappy? raw::snappy::MaxCompressedLength(hdr->DataSize):
		hdr->DataSize
	);
	uint8_t* buffer = (uint8_t*)malloc(bufferSize);
	if(buffer == NULL) {
		return 0;
	}

	// write header
	uint8_t* p = buffer;
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
		raw::snappy::RawCompress((const char*)pix, hdr->DataSize, (char*)pDataBuf, &output_length);
		hdr->DataSize = output_length;
	} else {
		memcpy(pDataBuf, pix, hdr->DataSize);
	}
	memcpy(p, &hdr->DataSize, sizeof(hdr->DataSize));
	p += sizeof(hdr->DataSize);
	hdr->Data = (uint8_t*)p;
	p += hdr->DataSize;

	// write crc32
	if(hdr->UseRC32) {
		hdr->CheckSum = rawHashCRC32((const char*)buffer, p-buffer);
	}
	memcpy(p, &hdr->CheckSum, sizeof(hdr->CheckSum));
	p += sizeof(hdr->CheckSum);

	// OK
	*output = buffer;
	return p-buffer;
}

void rawFree(void* p) {
	free(p);
}
