// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

//
// build jxr.dll on windows
//
// go run builder.go -win64 -dlldir=abcd
// go run builder.go -win32 -dlldir=${ENV}\bin
//
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	flagWin64  = flag.Bool("win64", false, "build win64 target")
	flagWin32  = flag.Bool("win32", false, "build win32 target")
	flagDllDir = flag.String("dlldir", "", "copy dll to this dir")
	flagClean  = flag.Bool("clean", false, "make clean")
)

func main() {
	flag.Parse()

	// check depend tools
	if _, err := exec.LookPath("devenv"); err != nil {
		log.Fatal("checkDependTools: can't find ", "devenv")
	}

	if *flagWin64 || *flagWin32 {
		buildTarget()
	}
	if *flagDllDir != "" {
		installDll(*flagDllDir)
	}
	if *flagClean {
		makeClean()
	}

	fmt.Println("Done.")
}

func buildTarget() {
	if *flagWin64 {
		fmt.Printf("build win64 jxr.cgo ...\n")
	} else {
		fmt.Printf("build win32 jxr.cgo ...\n")
	}

	var buildDir, logName string
	if *flagWin64 {
		buildDir = "zz_build_win64_proj_mt_tmp"
		logName = "builder-win64.log"
	} else {
		buildDir = "zz_build_win32_proj_mt_tmp"
		logName = "builder-win32.log"
	}
	os.Mkdir(buildDir, 0666)

	// generate vc2012 project
	var cmdGenXpProj *exec.Cmd
	if *flagWin64 {
		cmdGenXpProj = exec.Command(
			`cmake`, `..`,
			`-G`, `Visual Studio 11 Win64`,
			`-DCMAKE_BUILD_TYPE=release`,
			`-DCMAKE_INSTALL_PREFIX=..`,
			`-DCMAKE_C_FLAGS_DEBUG=/MTd /Zi /Od /Ob0 /RTC1`,
			`-DCMAKE_CXX_FLAGS_DEBUG=/MTd /Zi /Od /Ob0 /RTC1`,
			`-DCMAKE_C_FLAGS_RELEASE=/MT /O2 /Ob2 /DNDEBUG`,
			`-DCMAKE_CXX_FLAGS_RELEASE=/MT /O2 /Ob2 /DNDEBUG`,
			`-DCMAKE_EXE_LINKER_FLAGS=/MANIFEST:NO`,
		)
	} else {
		cmdGenXpProj = exec.Command(
			`cmake`, `..`,
			`-G`, `Visual Studio 11`,
			`-DCMAKE_GENERATOR_TOOLSET=v110_xp`, // VC2012, support xp
			`-DCMAKE_BUILD_TYPE=release`,
			`-DCMAKE_INSTALL_PREFIX=..`,
			`-DCMAKE_C_FLAGS_DEBUG=/MTd /Zi /Od /Ob0 /RTC1`,
			`-DCMAKE_CXX_FLAGS_DEBUG=/MTd /Zi /Od /Ob0 /RTC1`,
			`-DCMAKE_C_FLAGS_RELEASE=/MT /O2 /Ob2 /DNDEBUG`,
			`-DCMAKE_CXX_FLAGS_RELEASE=/MT /O2 /Ob2 /DNDEBUG`,
			`-DCMAKE_EXE_LINKER_FLAGS=/MANIFEST:NO`,
		)
	}
	cmdGenXpProj.Dir = buildDir

	// build and install
	var cmdBuildInstall *exec.Cmd
	if *flagWin64 {
		cmdBuildInstall = exec.Command(
			`devenv`, `JXR_LIB.sln`, `/build`, `Release|x64`, `/project`, `INSTALL.vcxproj`,
		)
	} else {
		cmdBuildInstall = exec.Command(
			`devenv`, `JXR_LIB.sln`, `/build`, `Release|Win32`, `/project`, `INSTALL.vcxproj`,
		)
	}
	cmdBuildInstall.Dir = buildDir

	// dlltool -dllname jxr.dll --def jxr.def --output-lib libjxr.a
	var cmdDll2A *exec.Cmd
	if *flagWin64 {
		cmdDll2A = exec.Command(
			`dlltool`,
			`-dllname`, `jxr-cgo-win64.dll`,
			`--def`, `jxr-cgo-win64.def`,
			`--output-lib`, `libjxr-cgo-win64.a`,
		)
	} else {
		cmdDll2A = exec.Command(
			`dlltool`,
			`-dllname`, `jxr-cgo-win32.dll`,
			`--def`, `jxr-cgo-win32.def`,
			`--output-lib`, `libjxr-cgo-win32.a`,
		)
	}
	cmdDll2A.Dir = "."

	// run ...
	out, err := cmdGenXpProj.CombinedOutput()
	if err != nil {
		ioutil.WriteFile(logName, out, 0666)
		log.Fatalf("buildTarget: %v, see %s", err, logName)
	}
	out, err = cmdBuildInstall.CombinedOutput()
	if err != nil {
		ioutil.WriteFile(logName, out, 0666)
		log.Fatalf("buildTarget: %v, see %s", err, logName)
	}
	out, err = cmdDll2A.CombinedOutput()
	if err != nil {
		ioutil.WriteFile(logName, out, 0666)
		log.Fatalf("buildTarget: %v, see %s", err, logName)
	}

	os.Remove(logName)
}

func installDll(dir_or_env string) {
	if dir := parseDir(dir_or_env); dir != "" {
		fmt.Printf("install dir: %s\n", dir)
		os.MkdirAll(dir, 0666)
		if *flagWin64 {
			cpFile(dir+"/jxr-cgo-win64.dll", "jxr-cgo-win64.dll")
		} else {
			cpFile(dir+"/jxr-cgo-win32.dll", "jxr-cgo-win32.dll")
		}
	}
}

func makeClean() {
	os.RemoveAll("zz_build_win64_proj_mt_tmp")
	os.RemoveAll("zz_build_win32_proj_mt_tmp")
}

func parseDir(dir_or_env string) string {
	if dir_or_env == "" {
		return ""
	}
	if !strings.HasPrefix(dir_or_env, "${") {
		return dir_or_env
	}
	if idx := strings.Index(dir_or_env, "}"); idx >= 0 {
		return os.Getenv(dir_or_env[2:idx]) + dir_or_env[idx+1:]
	} else {
		log.Fatalf("parseDir: bad dir %q", dir_or_env)
	}
	return ""
}

func cpFile(dst, src string) {
	fsrc, err := os.Open(src)
	if err != nil {
		log.Fatal("cpFile: ", err)
	}
	defer fsrc.Close()

	fdst, err := os.Create(dst)
	if err != nil {
		log.Fatal("cpFile: ", err)
	}
	defer fdst.Close()
	if _, err = io.Copy(fdst, fsrc); err != nil {
		log.Fatal("cpFile: ", err)
	}
}
