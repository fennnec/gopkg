// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stdio.h>
#include <stdlib.h>
#include <signal.h>

#include <windows.h>

bool LockFile(const char* name, HANDLE* handle) {
	*handle = INVALID_HANDLE_VALUE;
	HANDLE f = CreateFileA(name,
		GENERIC_READ | GENERIC_WRITE,
		0, NULL, OPEN_ALWAYS,
		0, NULL
	);
	if(f == INVALID_HANDLE_VALUE) return false;
	*handle = f;
	return true;
}

bool UnlockFile(HANDLE handle) {
	return handle == INVALID_HANDLE_VALUE || CloseHandle(handle) != FALSE;
}

HANDLE flock;
void sigHandle(int sig) {
	switch(sig) {
	case SIGINT:
		printf("sigHandle: %d, SIGINT\n", sig);
		break;
	default:
		printf("sigHandle: %d, OTHER\n", sig);
		break;
	}
	UnlockFile(flock);
	exit(1);
}

int main() {
	if(LockFile("demo.lock", &flock)) {
		printf("Lock Success!\n");
	} else {
		printf("Lock Failed!\n");
	}
	signal(SIGINT, sigHandle);
	for(;;) {}
	return 0;
}
