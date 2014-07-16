// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"bufio"
	"image"
	"io"
	"os"
	"strings"
)

// A format holds an image format's name, magic header and how to decode it.
type format struct {
	name, magic  string
	decode       func(io.Reader) (image.Image, error)
	decodeConfig func(io.Reader) (image.Config, error)
	encode       func(w io.Writer, m image.Image, opt interface{}) error
}

// Formats is the list of registered formats.
var formats []format

// RegisterFormat registers an image format for use by Decode.
// Name is the name of the format, like "jpeg" or "png".
// Magic is the magic prefix that identifies the format's encoding. The magic
// string can contain "?" wildcards that each match any one byte.
// Decode is the function that decodes the encoded image.
// DecodeConfig is the function that decodes just its configuration.
// Encode is the function that encodes just its configuration.
func RegisterFormat(
	name, magic string,
	decode func(io.Reader) (image.Image, error),
	decodeConfig func(io.Reader) (image.Config, error),
	encode func(w io.Writer, m image.Image, opt interface{}) error,
) {
	formats = append(formats, format{
		name, magic,
		decode,
		decodeConfig,
		encode,
	})
}

// A reader is an io.Reader that can also peek ahead.
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// asReader converts an io.Reader to a reader.
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of r's data.
func sniff(r reader) format {
	for _, f := range formats {
		b, err := r.Peek(len(f.magic))
		if err == nil && match(f.magic, b) {
			return f
		}
	}
	return format{}
}

// Decode decodes an image that has been encoded in a registered format.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
func Decode(r io.Reader) (image.Image, string, error) {
	rr := asReader(r)
	f := sniff(rr)
	if f.decode == nil {
		return nil, "", image.ErrFormat
	}
	m, err := f.decode(rr)
	return m, f.name, err
}

// DecodeConfig decodes the color model and dimensions of an image that has
// been encoded in a registered format. The string returned is the format name
// used during format registration. Format registration is typically done by
// an init function in the codec-specific package.
func DecodeConfig(r io.Reader) (image.Config, string, error) {
	rr := asReader(r)
	f := sniff(rr)
	if f.decodeConfig == nil {
		return image.Config{}, "", image.ErrFormat
	}
	c, err := f.decodeConfig(rr)
	return c, f.name, err
}

// Encode encodes an image as a registered format.
// The format is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
func Encode(format string, w io.Writer, m image.Image, opt interface{}) error {
	for _, f := range formats {
		if f.name == format {
			return f.encode(w, m, opt)
		}
	}
	return image.ErrFormat
}

func Load(filename string) (m image.Image, format string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	m, format, err = Decode(f)
	if err != nil {
		return
	}
	return
}

func Save(filename, format string, m image.Image) (err error) {
	if format == "" {
		if len(filename) != 0 {
			if idx := strings.LastIndex(filename, "."); idx >= 0 {
				format = strings.ToLower(filename[idx+1:])
			}
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	if err = Encode(format, f, m, nil); err != nil {
		return
	}
	return
}
