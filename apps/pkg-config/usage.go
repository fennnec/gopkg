// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var _ = `
pkg-config --help
Usage: pkg-config [OPTION...] [LIBRARIES]
  --version                             output version of pkg-config
  --modversion                          output version for package
  --atleast-pkgconfig-version=VERSION   require given version of pkg-config
  --libs                                output all linker flags
  --static                              output linker flags for static linking
  --short-errors                        print short errors
  --libs-only-l                         output -l flags
  --libs-only-other                     output other libs (e.g. -pthread)
  --libs-only-L                         output -L flags
  --cflags                              output all pre-processor and compiler
                                        flags
  --cflags-only-I                       output -I flags
  --cflags-only-other                   output cflags not covered by the
                                        cflags-only-I option
  --variable=NAME                       get the value of variable named NAME
  --define-variable=NAME=VALUE          set variable NAME to VALUE
  --exists                              return 0 if the module(s) exist
  --print-variables                     output list of variables defined by
                                        the module
  --uninstalled                         return 0 if the uninstalled version of
                                        one or more module(s) or their
                                        dependencies will be used
  --atleast-version=VERSION             return 0 if the module is at least
                                        version VERSION
  --exact-version=VERSION               return 0 if the module is at exactly
                                        version VERSION
  --max-version=VERSION                 return 0 if the module is at no newer
                                        than version VERSION
  --list-all                            list all known packages
  --debug                               show verbose debug information
  --print-errors                        show verbose information about missing
                                        or conflicting packages,default if
                                        --cflags or --libs given on the
                                        command line
  --silence-errors                      be silent about errors (default unless
                                        --cflags or --libsgiven on the command
                                        line)
  --errors-to-stdout                    print errors from --print-errors to
                                        stdout not stderr
  --print-provides                      print which packages the package
                                        provides
  --print-requires                      print which packages the package
                                        requires
  --print-requires-private              print which packages the package
                                        requires for static linking
  --dont-define-prefix                  don't try to override the value of
                                        prefix for each .pc file found with a
                                        guesstimated value based on the
                                        location of the .pc file
  --prefix-variable=PREFIX              set the name of the variable that
                                        pkg-config automatically sets
  --msvc-syntax                         output -l and -L flags for the
                                        Microsoft compiler (cl)

Help options
  -?, --help                            Show this help message
  --usage                               Display brief usage message

pkg-config --usage
Usage: pkg-config [-?] [--version] [--modversion]
        [--atleast-pkgconfig-version=VERSION] [--libs] [--static]
        [--short-errors] [--libs-only-l] [--libs-only-other] [--libs-only-L]
        [--cflags] [--cflags-only-I] [--cflags-only-other] [--variable=NAME]
        [--define-variable=NAME=VALUE] [--exists] [--print-variables]
        [--uninstalled] [--atleast-version=VERSION] [--exact-version=VERSION]
        [--max-version=VERSION] [--list-all] [--debug] [--print-errors]
        [--silence-errors] [--errors-to-stdout] [--print-provides]
        [--print-requires] [--print-requires-private] [--dont-define-prefix]
        [--prefix-variable=PREFIX] [--msvc-syntax] [--usage]
        [LIBRARIES]
`
