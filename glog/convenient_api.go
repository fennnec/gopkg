// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glog

const (
	LogThreshold_info    = int(infoLog)
	LogThreshold_warning = int(warningLog)
	LogThreshold_error   = int(errorLog)
	LogThreshold_fatal   = int(fatalLog)
	LogThreshold_Max     = int(numSeverity)
)

func GetToStderr() bool {
	return logging.toStderr
}
func SetToStderr(enabled bool) {
	logging.toStderr = enabled
}

func GetAlsoToStderr() bool {
	return logging.alsoToStderr
}
func SetAlsoToStderr(enabled bool) {
	logging.alsoToStderr = enabled
}

func GetStderrThreshold() int {
	return int(logging.stderrThreshold)
}
func SetStderrThreshold(threshold int) {
	if threshold >= 0 && threshold < int(numSeverity) {
		logging.stderrThreshold = severity(threshold)
	}
}

func GetLogDir() string {
	if len(logDirs) > 0 {
		return logDirs[0]
	} else {
		return ""
	}
}
func SetLogDir(dir string) {
	if len(logDirs) > 0 {
		logDirs[0] = dir
	} else {
		logDirs = []string{dir}
	}
}
