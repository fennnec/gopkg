:: Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

setlocal

cd %~dp0

call %CHAI2010_GOPKG_ROOT%\src\register-vc2012-386.bat
go run builder.go -win32 -clean -dlldir=${CHAI2010_GOPKG_ROOT}/bin

call %CHAI2010_GOPKG_ROOT%\src\register-vc2012-x64.bat
go run builder.go -win64 -clean -dlldir=${CHAI2010_GOPKG_ROOT}/bin

