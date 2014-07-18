// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raw

// Memory Picture -> MemP -> WewP (Not WebP)

import (
	"image"
	"io"

	image_ext "github.com/chai2010/gopkg/image"
)

const (
	HeaderSize = 21         // +CheckSum
	Magic      = 0x1BF2380A // CRC32(chaishushan@gmail.com)

	// data type
	DataType_UInt  = 1
	DataType_Int   = 2
	DataType_Float = 3
)

// Raw Image Spec (Little Endian), 21Bytes(+CheckSum).
type RawHeader struct {
	Sig       uint32 // 4Bytes, 1BF2380A, CRC32(chaishushan@gmail.com)
	UseRC32   byte   // 1Bytes, 0=disabled, 1=enabled (RawHeader.CheckSum)
	UseSnappy byte   // 1Bytes, 0=disabled, 1=enabled (RawHeader.Data)
	DataType  byte   // 1Bytes, 1=Uint, 2=Int, 3=Float
	Depth     byte   // 1Bytes, 8/16/32/64 bits
	Channels  byte   // 1Bytes, 1=Gray, 3=RGB, 4=RGBA
	Width     uint16 // 2Bytes, image Width
	Height    uint16 // 2Bytes, image Height
	DataSize  uint32 // 4Bytes, image data size (RawHeader.Data)
	Data      []byte // ?Bytes, image data (RawHeader.DataSize)
	CheckSum  uint32 // 4Bytes, CRC32(RawHeader[:len(RawHeader)-len(CheckSum)]) or Sig
}

type Options struct {
	UseRC32   bool // 0=disabled, 1=enabled (RawHeader.CheckSum)
	UseSnappy bool // 0=disabled, 1=enabled (RawHeader.Data)
}

func GetHeader(data []byte) (hdr *RawHeader, err error) {
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
		"raw", "\x1B\xF2\x38\x0A", // rawMagic
		Decode, DecodeConfig,
		encode,
	)
}
