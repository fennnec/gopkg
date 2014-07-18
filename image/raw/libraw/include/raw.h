// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef RAW_H_
#define RAW_H_

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

// header size
static const int kRawHeaderSize = 21;

// magic
static const uint32_t kRawMagic = 0x1BF2380A;

// data type
static const uint8_t kRawDataType_UInt  = 1;
static const uint8_t kRawDataType_Int   = 2;
static const uint8_t kRawDataType_Float = 3;

// Raw Image Spec (Little Endian), 21Bytes(+CheckSum).
typedef struct RawHeader {
	uint32_t Sig;       // 4Bytes, 1BF2380A, CRC32(chaishushan@gmail.com)
	uint8_t  UseRC32;   // 1Bytes, 0=disabled, 1=enabled (RawHeader.CheckSum)
	uint8_t  UseSnappy; // 1Bytes, 0=disabled, 1=enabled (RawHeader.Data)
	uint8_t  DataType;  // 1Bytes, 1=Uint, 2=Int, 3=Float
	uint8_t  Depth;     // 1Bytes, 8/16/32/64 bits
	uint8_t  Channels;  // 1Bytes, 1=Gray, 3=RGB, 4=RGBA
	uint16_t Width;     // 2Bytes, image Width
	uint16_t Height;    // 2Bytes, image Height
	uint32_t DataSize;  // 4Bytes, image data size (RawHeader.Data)
	uint8_t* Data;      // ?Bytes, image data (RawHeader.DataSize)
	uint32_t CheckSum;  // 4Bytes, CRC32(RawHeader[:len(RawHeader)-len(CheckSum)]) or Sig
} RawHeader;

typedef struct RawEncodeOptions {
	uint8_t  UseRC32;   // 0=disabled, 1=enabled (RawHeader.CheckSum)
	uint8_t  UseSnappy; // 0=disabled, 1=enabled (RawHeader.Data)
} RawEncodeOptions;

int rawGetHeader(
	const uint8_t* data, size_t data_size,
	RawHeader* hdr
);

uint8_t* rawDecode(
	const uint8_t* data, size_t data_size,
	int* width, int* height, int* channels, int* depth, int* data_type
);

size_t rawEncode(
	const uint8_t* pix, int width, int height, int stride,
	int channels, int depth, int data_type,
	const RawEncodeOptions* opt,
	uint8_t** output
);

void rawFree(void* p);

#ifdef __cplusplus
}
#endif
#endif // WEBP_H_
