// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows && go1.9
<<<<<<< HEAD
=======
// +build windows,go1.9
>>>>>>> deathstrox/main

package windows

import "syscall"

type Errno = syscall.Errno
type SysProcAttr = syscall.SysProcAttr
