// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef RAWP_H_
#define RAWP_H_

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

// header size
static const int kRawPHeaderSize = 24;

// magic
static const char*    kRawPSig   = "RawP";
static const uint32_t kRawPMagic = 0x1BF2380A;

// data type
static const uint8_t kRawPDataType_UInt  = 1;
static const uint8_t kRawPDataType_Int   = 2;
static const uint8_t kRawPDataType_Float = 3;

// RawP Image Spec (Little Endian), 24Bytes.
typedef struct RawPHeader {
	char     Sig[4];       // 4Bytes, RawP
	uint32_t Magic;        // 4Bytes, 0x1BF2380A
	uint16_t Width;        // 2Bytes, image Width
	uint16_t Height;       // 2Bytes, image Height
	uint8_t  Channels;     // 1Bytes, 1=Gray, 3=RGB, 4=RGBA
	uint8_t  Depth;        // 1Bytes, 8/16/32/64 bits
	uint8_t  DataType;     // 1Bytes, 1=Uint, 2=Int, 3=Float
	uint8_t  UseSnappy;    // 1Bytes, 0=disabled, 1=enabled (RawPHeader.Data)
	uint32_t DataSize;     // 4Bytes, image data size (RawPHeader.Data)
	uint32_t DataCheckSum; // 4Bytes, CRC32(RawPHeader.Data[RawPHeader.DataSize])
	uint8_t* Data;         // ?Bytes, image data (RawPHeader.DataSize)
} RawPHeader;

typedef struct RawPEncodeOptions {
	uint8_t  UseSnappy;    // 0=disabled, 1=enabled (RawPHeader.Data)
} RawPEncodeOptions;

typedef struct RawPEncodeContext {
	RawPHeader Header;
	uint8_t*   Pix;
	uint32_t   MaxEncodedLength;
} RawPEncodeContext;

int rawpDecodeHeader(
	const uint8_t* data, int data_size,
	RawPHeader* hdr
);

int rawpDecode(
	const uint8_t* data, int data_size,
	uint8_t* output, int output_size,
	RawPHeader* hdr
);

int rawpEncodeInit(
	const uint8_t* pix, int width, int height,
	int channels, int depth, int data_type,
	const RawPEncodeOptions* opt,
	RawPEncodeContext* ctx
);

size_t rawpEncode(
	RawPEncodeContext* ctx,
	uint8_t* output
);

#ifdef __cplusplus
}
#endif
#endif // RAWP_H_
