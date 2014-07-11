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

type jxr_bool_t C.jxr_bool_t

const (
	jxr_true  = jxr_bool_t(C.jxr_true)
	jxr_false = jxr_bool_t(C.jxr_false)
)

type jxr_data_type_t C.jxr_data_type_t

const (
	jxr_unsigned = jxr_data_type_t(C.jxr_unsigned)
	jxr_signed   = jxr_data_type_t(C.jxr_signed)
	jxr_float    = jxr_data_type_t(C.jxr_float)
)

func jxr_decode_config(data []byte) (
	width, height, channels, depth C.int,
	data_type jxr_data_type_t,
	err error,
) {
	if len(data) == 0 {
		err = fmt.Errorf("jxr_decode_config: bad arguments")
		return
	}
	rv := jxr_bool_t(C.jxr_decode_config(
		(*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)),
		(*C.int)(&width), (*C.int)(&height), (*C.int)(&channels), (*C.int)(&depth),
		(*C.jxr_data_type_t)(&data_type),
	))
	if rv != jxr_true {
		err = fmt.Errorf("jxr_decode_config: failed")
		return
	}
	return
}

func jxr_decode(data, pix []byte, stride int) (
	width, height, channels, depth C.int,
	data_type jxr_data_type_t,
	err error,
) {
	if len(data) == 0 {
		err = fmt.Errorf("jxr_decode: bad arguments")
		return
	}
	rv := jxr_bool_t(C.jxr_decode(
		(*C.char)(unsafe.Pointer(&pix[0])), C.int(stride),
		(*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)),
		(*C.int)(&width), (*C.int)(&height), (*C.int)(&channels), (*C.int)(&depth),
		(*C.jxr_data_type_t)(&data_type),
	))
	if rv != jxr_true {
		err = fmt.Errorf("jxr_decode: failed")
		return
	}
	return
}

func jxr_encode_len(
	pix []byte, stride int,
	width, height, channels, depth, quality int,
	data_type jxr_data_type_t,
) (err error) {
	if len(pix) == 0 {
		err = fmt.Errorf("jxr_encode_len: bad arguments")
		return
	}
	rv := jxr_bool_t(C.jxr_encode_len(
		(*C.char)(unsafe.Pointer(&pix[0])), C.int(stride),
		(C.int)(width), (C.int)(height), (C.int)(channels), (C.int)(depth), C.int(quality),
		(C.jxr_data_type_t)(data_type),
	))
	if rv != jxr_true {
		err = fmt.Errorf("jxr_encode_len: failed")
		return
	}
	return
}

func jxr_encode(
	buf, pix []byte, stride int,
	width, height, channels, depth, quality int,
	data_type jxr_data_type_t,
) (newSize C.int, err error) {
	if len(buf) == 0 || len(pix) == 0 {
		err = fmt.Errorf("jxr_encode: bad arguments")
		return
	}
	rv := jxr_bool_t(C.jxr_encode(
		(*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)),
		(*C.char)(unsafe.Pointer(&pix[0])), C.int(stride),
		C.int(width), C.int(height), C.int(channels), C.int(depth), C.int(quality),
		C.jxr_data_type_t(data_type),
		&newSize,
	))
	if rv != jxr_true {
		err = fmt.Errorf("jxr_encode_len: failed")
		return
	}
	return
}
