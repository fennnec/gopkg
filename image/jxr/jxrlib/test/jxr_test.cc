// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "test.h"
#include "test_util.h"

#include "jxr.h"

struct tImgInfo {
	int width;
	int height;
	int channels;
	int depth;
	const char* name;
};

static tImgInfo testCaseJpg[] = {
	{ 512 , 512 , 3, 8, "testdata/lena.jpg" },
	{ 512 , 512 , 3, 8, "testdata/lena-gray.jpg" },
	{ 2592, 3904, 3, 8, "testdata/FLOWER.jpg" },
	{ 3888, 2592, 3, 8, "testdata/SAKURA.jpg" },
	{ 3888, 2592, 3, 8, "testdata/SMALLTOMATO.jpg" },
};

static tImgInfo testCaseJxr[] = {
	{ 512 , 512 , 3, 8, "testdata/lena.wdp" },
	{ 512 , 512 , 3, 8, "testdata/lena-gray.wdp" },
	{ 2592, 3904, 3, 8, "testdata/FLOWER.wdp" },
	{ 3888, 2592, 3, 8, "testdata/SAKURA.wdp" },
	{ 3888, 2592, 3, 8, "testdata/SMALLTOMATO.wdp" },
};

TEST(jxr, JpgHelper) {
	auto buf = new std::string;
	auto src = new std::string;
	auto dst = new std::string;

	for(int i = 0; i < TEST_DIM(testCaseJpg); ++i) {
		bool rv = loadImageData(testCaseJpg[i].name, buf);
		ASSERT_TRUE(rv);

		// decode raw file data
		int width, height, channels;
		rv = jpegDecode(src, buf->data(), buf->size(), &width, &height, &channels);
		ASSERT_TRUE(rv);
		ASSERT_TRUE(width == testCaseJpg[i].width);
		ASSERT_TRUE(height == testCaseJpg[i].height);
		ASSERT_TRUE(channels == testCaseJpg[i].channels);

		// encode as jpg
		buf->clear();
		rv = jpegEncode(buf, src->data(), src->size(), width, height, channels, 90, 0);
		ASSERT_TRUE(rv);

		// decode again
		rv = jpegDecode(dst, buf->data(), buf->size(), &width, &height, &channels);
		ASSERT_TRUE(rv);
		ASSERT_TRUE(width == testCaseJpg[i].width);
		ASSERT_TRUE(height == testCaseJpg[i].height);
		ASSERT_TRUE(channels == testCaseJpg[i].channels);

		// compare
		double diff = diffImageData(
			(const unsigned char*)src->data(), (const unsigned char*)dst->data(),
			width, height, channels
		);
		ASSERT_TRUE(diff < 20);
	}

	delete buf;
	delete src;
	delete dst;
}

TEST(jxr, DecodeConfig) {
	auto buf = new std::string;
	for(int i = 0; i < TEST_DIM(testCaseJxr); ++i) {
		bool rv = loadImageData(testCaseJxr[i].name, buf);
		ASSERT_TRUE(rv);

		// decode jxr data
		int width, height, channels, depth;
		jxr_data_type_t type;
		jxr_bool_t ret = jxr_decode_config(buf->data(), buf->size(),
			&width, &height, &channels, &depth, &type
		);
		ASSERT_TRUE(ret == jxr_true);
		ASSERT_TRUE(width == testCaseJxr[i].width);
		ASSERT_TRUE(height == testCaseJxr[i].height);
		ASSERT_TRUE(channels == testCaseJxr[i].channels);
		ASSERT_TRUE(depth == testCaseJxr[i].depth);
	}
	delete buf;
}

TEST(jxr, Decode) {
	auto buf = new std::string;
	auto src = new std::string;
	auto dst = new std::string;

	for(int i = 0; i < TEST_DIM(testCaseJxr); ++i) {
		int width, height, channels, depth;
		jxr_data_type_t type;

		// decode jxr data
		bool rv = loadImageData(testCaseJxr[i].name, buf);
		ASSERT_TRUE(rv);
		src->resize(testCaseJxr[i].width*testCaseJxr[i].height*testCaseJxr[i].channels);
		int n = jxr_decode(
			(char*)src->data(), 0, buf->data(), buf->size(),
			&width, &height, &channels, &depth, &type
		);
		ASSERT_TRUE(n == jxr_true);
		ASSERT_TRUE(width == testCaseJxr[i].width);
		ASSERT_TRUE(height == testCaseJxr[i].height);
		ASSERT_TRUE(channels == testCaseJxr[i].channels);
		ASSERT_TRUE(depth == testCaseJxr[i].depth);

		// decode jpg data
		rv = loadImageData(testCaseJpg[i].name, buf);
		ASSERT_TRUE(rv);
		rv = jpegDecode(dst, buf->data(), buf->size(), &width, &height, &channels);
		ASSERT_TRUE(rv);
		ASSERT_TRUE(width == testCaseJpg[i].width);
		ASSERT_TRUE(height == testCaseJpg[i].height);
		ASSERT_TRUE(channels == testCaseJpg[i].channels);

		// compare
		double diff = diffImageData(
			(const unsigned char*)src->data(), (const unsigned char*)dst->data(),
			width, height, channels
		);
		ASSERT_TRUE(diff < 20);
	}

	delete buf;
	delete src;
	delete dst;
}

TEST(jxr, Encode) {
	//
}

TEST(jxr, DecodeAndEncode) {
	return; // skip

	auto buf = new std::string;
	auto src = new std::string;
	auto dst = new std::string;

	for(int i = 0; i < TEST_DIM(testCaseJxr); ++i) {
		bool rv = loadImageData(testCaseJxr[i].name, buf);
		ASSERT_TRUE(rv);

		// decode raw file data
		int width, height, channels, depth;
		jxr_data_type_t type;
		src->resize(testCaseJxr[i].width*testCaseJxr[i].height*testCaseJxr[i].channels);
		int n = jxr_decode(
			(char*)src->data(), 0, buf->data(), buf->size(),
			&width, &height, &channels, &depth, &type
		);
		ASSERT_TRUE(n == jxr_true);
		ASSERT_TRUE(width == testCaseJxr[i].width);
		ASSERT_TRUE(height == testCaseJxr[i].height);
		ASSERT_TRUE(channels == testCaseJxr[i].channels);
		ASSERT_TRUE(depth == testCaseJxr[i].depth);

		// encode as jxr
		buf->clear();
		buf->resize(src->size());
		n = jxr_encode(
			(char*)buf->data(), buf->size(), src->data(), 0,
			width, height, channels, depth,
			90, jxr_unsigned
		);
		ASSERT_TRUE(n == jxr_true);

		// decode again
		dst->resize(testCaseJxr[i].width*testCaseJxr[i].height*testCaseJxr[i].channels);
		n = jxr_decode(
			(char*)dst->data(), dst->size(), buf->data(), buf->size(),
			&width, &height, &channels, &depth, &type
		);
		ASSERT_TRUE(n == jxr_true);
		ASSERT_TRUE(width == testCaseJxr[i].width);
		ASSERT_TRUE(height == testCaseJxr[i].height);
		ASSERT_TRUE(channels == testCaseJxr[i].channels);
		ASSERT_TRUE(depth == testCaseJxr[i].depth);

		// compare
		if(depth == 8) {
			double diff = diffImageData(
				(const unsigned char*)src->data(), (const unsigned char*)dst->data(),
				width, height, channels
			);
			ASSERT_TRUE(diff < 20);
		} else if(depth == 16) {
			double diff = diffImageData(
				(const unsigned short*)src->data(), (const unsigned short*)dst->data(),
				width, height, channels
			);
			ASSERT_TRUE(diff < 20);
		} else {
			ASSERT_TRUE(false);
		}
	}

	delete buf;
	delete src;
	delete dst;
}

TEST(jxr, CompareJxrJpg) {
	// diff(jxr, jpg) < 20
}

// benchmark
