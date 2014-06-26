// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "test_util.h"

#include <stdio.h>

bool loadImageData(const char* name, std::string* data) {
	FILE* fp = fopen(name, "rb");
	if(!fp) return false;

	fseek(fp, 0, SEEK_END);
	data->resize(ftell(fp));

	rewind(fp);
	int n = fread((void*)data->data(), 1, data->size(), fp);
	fclose(fp);
	return (n == data->size());
}

