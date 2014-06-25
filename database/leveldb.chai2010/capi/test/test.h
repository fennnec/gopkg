// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef TEST_H_
#define TEST_H_

#define TEST(x, y) \
	void x##y(void); \
	TestRegisterer r##x##y(x##y, # x "." # y); \
	void x##y(void)

void RegisterTest(void (*fn)(void), const char *tname);
void TestAssertTrue(bool condition, const char* fname, int lineno);
void TestAssertEQ(int a, int b, const char* fname, int lineno);
void TestAssertStrEQ(const char* a, const char* b, const char* fname, int lineno);
void TestAssertNear(float a, float b, float abs_error, const char* fname, int lineno);

struct TestRegisterer {
	TestRegisterer(void (*fn)(void), const char *s) {
		RegisterTest(fn, s);
	}
};

#define DIM(x) (sizeof(x)/sizeof((x)[0]))
#define ASSERT_TRUE(x) TestAssertTrue((x), __FILE__, __LINE__)
#define ASSERT_EQ(x, y) TestAssertEQ((x), (y), __FILE__, __LINE__)
#define ASSERT_STREQ(x, y)TestAssertStrEQ((x), (y), __FILE__, __LINE__)
#define ASSERT_NEAR(x, y, abs_error) TestAssertNear((x), (y), (abs_error), __FILE__, __LINE__)

#endif  // TEST_H_
