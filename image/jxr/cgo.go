// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jxr

/*
#cgo windows,amd64 LDFLAGS: -L. -l"jxr-cgo-win64"
#cgo windows,386 LDFLAGS: -L. -l"jxr-cgo-win32"
#cgo linux,amd64 LDFLAGS: -L. -l"jxr-cgo-posix64"
#cgo linux,386 LDFLAGS: -L. -l"jxr-cgo-posix32"

#cgo windows CFLAGS: -I./jxrlib/include -fno-stack-check -fno-stack-protector -mno-stack-arg-probe
#cgo linux   CFLAGS: -I./jxrlib/include

#include "jxr.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type jxr_data_type_t C.jxr_data_type_t

const (
	jxr_unsigned = jxr_data_type_t(C.jxr_unsigned)
	jxr_signed   = jxr_data_type_t(C.jxr_signed)
	jxr_float    = jxr_data_type_t(C.jxr_float)
)

func jxr_encode(
	data []byte, width, height, channels, depth, quality, width_step int,
	data_type jxr_data_type_t,
	buf []byte,
) (n int, err error) {
	if len(data) == 0 {
		err = fmt.Errorf("jxr_encode: bad arguments")
		return
	}
	// check size
	if len(buf) == 0 {
		n = int(C.jxr_encode(
			(*C.char)(unsafe.Pointer(nil)), C.int(0),
			(*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)),
			C.int(width), C.int(height), C.int(channels), C.int(depth),
			C.int(quality), C.int(width_step),
			C.jxr_data_type_t(data_type),
		))
		if n < 0 {
			err = fmt.Errorf("jxr_encode: n = %d", n)
			return
		}
	} else {
		n = int(C.jxr_encode(
			(*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)),
			(*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)),
			C.int(width), C.int(height), C.int(channels), C.int(depth),
			C.int(quality), C.int(width_step),
			C.jxr_data_type_t(data_type),
		))
		if n < 0 {
			err = fmt.Errorf("jxr_encode: n = %d", n)
			return
		}
	}
	return
}

func jxr_decode(data, buf []byte) (
	width, height, channels, depth C.int,
	data_type jxr_data_type_t,
	n int, err error,
) {
	if len(data) == 0 {
		err = fmt.Errorf("jxr_decode: bad arguments")
		return
	}
	// check size
	if len(buf) == 0 {
		n = int(C.jxr_decode(
			(*C.char)(unsafe.Pointer(nil)), C.int(0),
			(*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)),
			(*C.int)(&width), (*C.int)(&height), (*C.int)(&channels), (*C.int)(&depth),
			(*C.jxr_data_type_t)(&data_type),
		))
		if n < 0 {
			err = fmt.Errorf("jxr_decode: n = %d", n)
			return
		}
	} else {
		n = int(C.jxr_decode(
			(*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)),
			(*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)),
			(*C.int)(&width), (*C.int)(&height), (*C.int)(&channels), (*C.int)(&depth),
			(*C.jxr_data_type_t)(&data_type),
		))
		if n < 0 {
			err = fmt.Errorf("jxr_decode: n = %d", n)
			return
		}
	}
	return
}
