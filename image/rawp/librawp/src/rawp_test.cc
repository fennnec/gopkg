// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "rawp.h"
#include "test.h"

TEST(RawP, Simple) {
	int width = 800;
	int height = 600;
	int channels = 3;
	int depth = 8;
	int data_type = kRawPDataType_UInt;
	RawPEncodeOptions* opt = NULL;

	uint8_t* rgb = new uint8_t[width*height*channels*depth/8];
	for(int i = 0; i < width*height*channels*depth/8; ++i) {
		rgb[i] = i&0xFF;
	}

	RawPEncodeContext ctx;
	int rv = rawpEncodeInit(
		rgb, width, height,
		channels, depth, data_type,
		opt,
		&ctx
	);
	ASSERT_TRUE(rv != 0);

	uint8_t* output = new uint8_t[ctx.MaxEncodedLength];
	size_t output_size = rawpEncode(&ctx, output);
	ASSERT_TRUE(output_size > 0);

	RawPHeader hdr;
	memset(&hdr, 0, sizeof(hdr));
	rv = rawpDecodeHeader(output, output_size, &hdr);
	ASSERT_TRUE(rv != 0);
	ASSERT_TRUE(hdr.Width == width);
	ASSERT_TRUE(hdr.Height == height);
	ASSERT_TRUE(hdr.Channels == channels);
	ASSERT_TRUE(hdr.Depth == depth);
	ASSERT_TRUE(hdr.DataType == data_type);

	uint8_t* rgb2 = new uint8_t[width*height*channels*depth/8];
	memset(rgb2, 0, width*height*channels*depth/8);

	
	memset(&hdr, 0, sizeof(hdr));
	rv = rawpDecode(
		output, output_size,
		rgb2, width*height*channels*depth/8,
		&hdr
	);
	ASSERT_TRUE(rv != 0);
	ASSERT_TRUE(hdr.Width == width);
	ASSERT_TRUE(hdr.Height == height);
	ASSERT_TRUE(hdr.Channels == channels);
	ASSERT_TRUE(hdr.Depth == depth);
	ASSERT_TRUE(hdr.DataType == data_type);

	for(int i = 0; i < width*height*channels*depth/8; ++i) {
		if(rgb[i] != rgb2[i]) {
			ASSERT_TRUE_MSG(false, "i = %d", i);
		}
	}
}

TEST(RawP, Snappy) {
	int width = 800;
	int height = 600;
	int channels = 3;
	int depth = 8;
	int data_type = kRawPDataType_UInt;
	RawPEncodeOptions* opt = new RawPEncodeOptions;
	opt->UseSnappy = 1;

	uint8_t* rgb = new uint8_t[width*height*channels*depth/8];
	for(int i = 0; i < width*height*channels*depth/8; ++i) {
		rgb[i] = i&0xFF;
	}

	RawPEncodeContext ctx;
	int rv = rawpEncodeInit(
		rgb, width, height,
		channels, depth, data_type,
		opt,
		&ctx
	);
	ASSERT_TRUE(rv != 0);

	uint8_t* output = new uint8_t[ctx.MaxEncodedLength];
	size_t output_size = rawpEncode(&ctx, output);
	ASSERT_TRUE(output_size > 0);

	RawPHeader hdr;
	memset(&hdr, 0, sizeof(hdr));
	rv = rawpDecodeHeader(output, output_size, &hdr);
	ASSERT_TRUE(rv != 0);
	ASSERT_TRUE(hdr.Width == width);
	ASSERT_TRUE(hdr.Height == height);
	ASSERT_TRUE(hdr.Channels == channels);
	ASSERT_TRUE(hdr.Depth == depth);
	ASSERT_TRUE(hdr.DataType == data_type);

	ASSERT_TRUE(hdr.UseSnappy == 1);

	uint8_t* rgb2 = new uint8_t[width*height*channels*depth/8];
	memset(rgb2, 0, width*height*channels*depth/8);

	
	memset(&hdr, 0, sizeof(hdr));
	rv = rawpDecode(
		output, output_size,
		rgb2, width*height*channels*depth/8,
		&hdr
	);
	ASSERT_TRUE(rv != 0);
	ASSERT_TRUE(hdr.Width == width);
	ASSERT_TRUE(hdr.Height == height);
	ASSERT_TRUE(hdr.Channels == channels);
	ASSERT_TRUE(hdr.Depth == depth);
	ASSERT_TRUE(hdr.DataType == data_type);

	for(int i = 0; i < width*height*channels*depth/8; ++i) {
		if(rgb[i] != rgb2[i]) {
			ASSERT_TRUE_MSG(false, "i = %d", i);
		}
	}
}

