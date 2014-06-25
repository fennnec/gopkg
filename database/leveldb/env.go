// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

// #include <leveldb/c.h>
import "C"

// Env is a system call environment used by a database.
//
// Typically, NewDefaultEnv is all you need. Advanced users may create their
// own Env with a *C.leveldb_env_t of their own creation.
//
// To prevent memory leaks, an Env must have Close called on it when it is
// no longer needed by the program.
type Env struct {
	env *C.leveldb_env_t
}

// NewDefaultEnv creates a default environment for use in an Options.
//
// To prevent memory leaks, the Env returned should be deallocated with
// Close.
func NewDefaultEnv() *Env {
	return &Env{C.leveldb_create_default_env()}
}

// Close deallocates the Env, freeing the underlying struct.
func (env *Env) Close() {
	C.leveldb_env_destroy(env.env)
}
