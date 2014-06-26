:: Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

setlocal

cd %~dp0

del /S/Q jxr-cgo-win32.dll
del /S/Q jxr-cgo-win32.lib

del /S/Q jxr-cgo-win64.dll
del /S/Q jxr-cgo-win64.lib

del /S/Q jxr-cgo-win32.a
del /S/Q jxr-cgo-win64.a

