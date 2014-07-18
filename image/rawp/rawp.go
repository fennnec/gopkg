// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rawp implements a decoder and encoder for RawP images.
package rawp

import (
	"image"
	"io"

	image_ext "github.com/chai2010/gopkg/image"
)

const (
	rawpHeaderSize = 25
	rawpMagic      = 0x1BF2380A

	// data type
	DataType_UInt  = 1
	DataType_Int   = 2
	DataType_Float = 3
)

// RawP Image Spec (Little Endian), 25Bytes(+CheckSum).
type RawPHeader struct {
	Sig       [4]byte // 4Bytes, WEWP
	Magic     uint32  // 4Bytes, 0x1BF2380A
	UseCRC32  byte    // 1Bytes, 0=disabled, 1=enabled (RawPHeader.CheckSum)
	UseSnappy byte    // 1Bytes, 0=disabled, 1=enabled (RawPHeader.Data)
	DataType  byte    // 1Bytes, 1=Uint, 2=Int, 3=Float
	Depth     byte    // 1Bytes, 8/16/32/64 bits
	Channels  byte    // 1Bytes, 1=Gray, 3=RGB, 4=RGBA
	Width     uint16  // 2Bytes, image Width
	Height    uint16  // 2Bytes, image Height
	DataSize  uint32  // 4Bytes, image data size (RawPHeader.Data)
	Data      []byte  // ?Bytes, image data (RawPHeader.DataSize)
	CheckSum  uint32  // 4Bytes, CRC32(RawPHeader[:len(RawPHeader)-len(CheckSum)]) or Sig
}

type Options struct {
	UseCRC32  bool // 0=disabled, 1=enabled (RawPHeader.CheckSum)
	UseSnappy bool // 0=disabled, 1=enabled (RawPHeader.Data)
}

func DecodeHeader(data []byte) (hdr *RawPHeader, err error) {
	return
}

func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return
}

func Decode(r io.Reader) (m image.Image, err error) {
	return
}

func Encode(w io.Writer, m image.Image, opt *Options) (err error) {
	return
}

func encode(w io.Writer, m image.Image, opt interface{}) error {
	if opt, ok := opt.(*Options); ok {
		return Encode(w, m, opt)
	} else {
		return Encode(w, m, nil)
	}
}

func init() {
	image_ext.RegisterFormat(
		"rawp", "RAWP\x1B\xF2\x38\x0A", // rawpMagic
		Decode, DecodeConfig,
		encode,
	)
}
