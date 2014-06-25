// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package srt

import (
	"strconv"
	"time"
)

// 00:00:15,000 --> 00:00:18,000
func parseDuration(s string) (time.Duration, error) {
	v0, err := strconv.Atoi(s[0:2])
	v1, err := strconv.Atoi(s[3:5])
	v2, err := strconv.Atoi(s[6:8])
	v3, err := strconv.Atoi(s[9:12])
	if err != nil {
		return 0, err
	}
	d := time.Duration(v0)*time.Hour +
		time.Duration(v1)*time.Minute +
		time.Duration(v2)*time.Second +
		time.Duration(v3)*time.Millisecond
	return d, nil
}

func makeDuration(hour, minute, sec, msec time.Duration) time.Duration {
	return hour*time.Hour + minute*time.Minute + sec*time.Second + msec*time.Millisecond
}
