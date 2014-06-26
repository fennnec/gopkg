:: Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

setlocal

cd %~dp0

rmdir /S/Q zz_build_win32_proj_mt_tmp_debug
rmdir /S/Q zz_build_win32_proj_mt_tmp_release
rmdir /S/Q zz_build_win64_proj_mt_tmp_debug
rmdir /S/Q zz_build_win64_proj_mt_tmp_release

del /Q JxrDecApp.exe
del /Q JxrEncApp.exe
del /Q jxrtest.exe

del /Q jxrlib-win32.lib
del /Q jxrlib-win32-mt.lib
del /Q jxrlib-win32-debug.lib
del /Q jxrlib-win32-debug-mt.lib

del /Q jxrlib-win64.lib
del /Q jxrlib-win64-mt.lib
del /Q jxrlib-win64-debug.lib
del /Q jxrlib-win64-debug-mt.lib

