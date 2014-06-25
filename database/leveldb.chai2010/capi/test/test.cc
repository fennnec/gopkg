// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "test.h"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

struct Test {
	void (*fn)(void);
	const char *name;
};

static Test tests[10000];
static int ntests;
static int testid;

static const char* basename(const char* fname) {
	int len = strlen(fname);
	const char* s = fname + len;
	while(s > fname) {
		if(s[-1] == '/' || s[-1] == '\\') return s;
		s--;
	}
	return s;
}

void RegisterTest(void (*fn)(void), const char* tname) {
	tests[ntests].fn = fn;
	tests[ntests++].name = tname;
}

void TestAssertTrue(bool condition, const char* fname, int lineno) {
	if(!condition) {
		fname = basename(fname);
		printf("fail, %s, line %d: ASSERT_TRUE(false)\n", fname, lineno);
		exit(-1);
	}
}
void TestAssertEQ(int a, int b, const char* fname, int lineno) {
	if(a != b) {
		fname = basename(fname);
		printf("fail, %s, line %d: ASSERT_EQ(%d, %d)\n", fname, lineno, a, b);
		exit(-1);
	}
}
void TestAssertStrEQ(const char* a, const char* b, const char* fname, int lineno) {
	if(strcmp(a, b) != 0) {
		fname = basename(fname);
		printf("fail, %s, line %d: ASSERT_STREQ(\"%s\", \"%s\")\n", fname, lineno, a, b);
		exit(-1);
	}
}
void TestAssertNear(float a, float b, float abs_error, const char* fname, int lineno) {
	if(abs(a-b) > abs(abs_error)) {
		fname = basename(fname);
		printf("fail, %s, line %d: ASSERT_NEAR(%f, %f, %f)\n", fname, lineno, a, b, abs_error);
		exit(-1);
	}
}

int main(int argc, char* argv[]) {
	for(testid = 0; testid < ntests; testid++) {
		printf("%s ", tests[testid].name);
		tests[testid].fn();
		printf("ok\n");
	}
	printf("PASS\n");
	return 0;
}
