// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package webp implements a decoder for WEBP images.
//
// WEBP is defined at:
// https://developers.google.com/speed/webp/docs/riff_container
package webp

import (
	"image"
	"io"
	"io/ioutil"

	image_ext "github.com/chai2010/gopkg/image"
)

func GetInfo(r io.Reader) (width, height int, hasAlpha bool, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return webpGetInfo(data)
}

func DecodeGray(r io.Reader) (m *image.Gray, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	pix, w, h, err := webpDecodeGray(data)
	if err != nil {
		return
	}
	m = &image.Gray{pix, 1 * w, image.Rect(0, 0, w, h)}
	return
}

func DecodeRGB(r io.Reader) (m *image_ext.RGB, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	pix, w, h, err := webpDecodeRGB(data)
	if err != nil {
		return
	}
	m = &image_ext.RGB{pix, 3 * w, image.Rect(0, 0, w, h)}
	return
}

func DecodeRGBA(r io.Reader) (m *image.RGBA, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	pix, w, h, err := webpDecodeRGBA(data)
	if err != nil {
		return
	}
	m = &image.RGBA{pix, 4 * w, image.Rect(0, 0, w, h)}
	return
}

func EncodeGray(w io.Writer, m *image.Gray, quality float32) (err error) {
	output, err := webpEncodeGray(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride, quality)
	if err != nil {
		return
	}
	_, err = w.Write(output)
	return
}

func EncodeRGB(w io.Writer, m *image_ext.RGB, quality float32) (err error) {
	output, err := webpEncodeRGB(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride, quality)
	if err != nil {
		return
	}
	_, err = w.Write(output)
	return
}

func EncodeRGBA(w io.Writer, m *image.RGBA, quality float32) (err error) {
	output, err := webpEncodeRGBA(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride, quality)
	if err != nil {
		return
	}
	_, err = w.Write(output)
	return
}

func EncodeLosslessGray(w io.Writer, m *image.Gray) (err error) {
	output, err := webpEncodeLosslessGray(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride)
	if err != nil {
		return
	}
	_, err = w.Write(output)
	return
}

func EncodeLosslessRGB(w io.Writer, m *image_ext.RGB) (err error) {
	output, err := webpEncodeLosslessRGB(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride)
	if err != nil {
		return
	}
	_, err = w.Write(output)
	return
}

func EncodeLosslessRGBA(w io.Writer, m *image.RGBA) (err error) {
	output, err := webpEncodeLosslessRGBA(m.Pix, m.Rect.Dx(), m.Rect.Dy(), m.Stride)
	if err != nil {
		return
	}
	_, err = w.Write(output)
	return
}
