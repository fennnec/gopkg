// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "sum.h"

#include <stdio.h>

struct Arith {
	int Sum(int array[], int len) const {
		int sum = 0;
		for(int i = 0; i < len; ++i) {
			sum += array[i];
		}
		return sum;
	}
};

int sum(int array[], int len) {
	Arith arith;
	int v = arith.Sum(array, len);
	return v;
}

int add(int a, int b) {
	return a+b;
}
