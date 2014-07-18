// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package webp implements a decoder and encoder for WEBP images.
//
// WEBP is defined at:
// https://developers.google.com/speed/webp/docs/riff_container
package webp

import (
	"image"

	image_ext "github.com/chai2010/gopkg/image"
)

func GetInfo(data []byte) (width, height int, hasAlpha bool, err error) {
	return webpGetInfo(data)
}

func DecodeGray(data []byte) (m *image.Gray, err error) {
	pix, w, h, err := webpDecodeGray(data)
	if err != nil {
		return
	}
	m = &image.Gray{pix, 1 * w, image.Rect(0, 0, w, h)}
	return
}

func DecodeRGB(data []byte) (m *image_ext.RGB, err error) {
	pix, w, h, err := webpDecodeRGB(data)
	if err != nil {
		return
	}
	m = &image_ext.RGB{pix, 3 * w, image.Rect(0, 0, w, h)}
	return
}

func DecodeRGBA(data []byte) (m *image.RGBA, err error) {
	pix, w, h, err := webpDecodeRGBA(data)
	if err != nil {
		return
	}
	m = &image.RGBA{pix, 4 * w, image.Rect(0, 0, w, h)}
	return
}

func EncodeGray(m *image.Gray, quality float32) (data []byte, err error) {
	return webpEncodeGray(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride, quality)
}

func EncodeRGB(m *image_ext.RGB, quality float32) (data []byte, err error) {
	return webpEncodeRGB(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride, quality)
}

func EncodeRGBA(m *image.RGBA, quality float32) (data []byte, err error) {
	return webpEncodeRGBA(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride, quality)
}

func EncodeLosslessGray(m *image.Gray) (data []byte, err error) {
	return webpEncodeLosslessGray(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride)
}

func EncodeLosslessRGB(m *image_ext.RGB) (data []byte, err error) {
	return webpEncodeLosslessRGB(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride)
}

func EncodeLosslessRGBA(m *image.RGBA) (data []byte, err error) {
	return webpEncodeLosslessRGBA(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride)
}
