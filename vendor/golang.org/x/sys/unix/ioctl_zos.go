// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build zos && s390x
<<<<<<< HEAD
=======
// +build zos,s390x
>>>>>>> deathstrox/main

package unix

import (
	"runtime"
	"unsafe"
)

// ioctl itself should not be exposed directly, but additional get/set
// functions for specific types are permissible.

// IoctlSetInt performs an ioctl operation which sets an integer value
// on fd, using the specified request number.
<<<<<<< HEAD
func IoctlSetInt(fd int, req int, value int) error {
=======
func IoctlSetInt(fd int, req uint, value int) error {
>>>>>>> deathstrox/main
	return ioctl(fd, req, uintptr(value))
}

// IoctlSetWinsize performs an ioctl on fd with a *Winsize argument.
//
// To change fd's window size, the req argument should be TIOCSWINSZ.
<<<<<<< HEAD
func IoctlSetWinsize(fd int, req int, value *Winsize) error {
	// TODO: if we get the chance, remove the req parameter and
	// hardcode TIOCSWINSZ.
	return ioctlPtr(fd, req, unsafe.Pointer(value))
=======
func IoctlSetWinsize(fd int, req uint, value *Winsize) error {
	// TODO: if we get the chance, remove the req parameter and
	// hardcode TIOCSWINSZ.
	err := ioctl(fd, req, uintptr(unsafe.Pointer(value)))
	runtime.KeepAlive(value)
	return err
>>>>>>> deathstrox/main
}

// IoctlSetTermios performs an ioctl on fd with a *Termios.
//
// The req value is expected to be TCSETS, TCSETSW, or TCSETSF
<<<<<<< HEAD
func IoctlSetTermios(fd int, req int, value *Termios) error {
=======
func IoctlSetTermios(fd int, req uint, value *Termios) error {
>>>>>>> deathstrox/main
	if (req != TCSETS) && (req != TCSETSW) && (req != TCSETSF) {
		return ENOSYS
	}
	err := Tcsetattr(fd, int(req), value)
	runtime.KeepAlive(value)
	return err
}

// IoctlGetInt performs an ioctl operation which gets an integer value
// from fd, using the specified request number.
//
// A few ioctl requests use the return value as an output parameter;
// for those, IoctlRetInt should be used instead of this function.
<<<<<<< HEAD
func IoctlGetInt(fd int, req int) (int, error) {
	var value int
	err := ioctlPtr(fd, req, unsafe.Pointer(&value))
	return value, err
}

func IoctlGetWinsize(fd int, req int) (*Winsize, error) {
	var value Winsize
	err := ioctlPtr(fd, req, unsafe.Pointer(&value))
=======
func IoctlGetInt(fd int, req uint) (int, error) {
	var value int
	err := ioctl(fd, req, uintptr(unsafe.Pointer(&value)))
	return value, err
}

func IoctlGetWinsize(fd int, req uint) (*Winsize, error) {
	var value Winsize
	err := ioctl(fd, req, uintptr(unsafe.Pointer(&value)))
>>>>>>> deathstrox/main
	return &value, err
}

// IoctlGetTermios performs an ioctl on fd with a *Termios.
//
// The req value is expected to be TCGETS
<<<<<<< HEAD
func IoctlGetTermios(fd int, req int) (*Termios, error) {
=======
func IoctlGetTermios(fd int, req uint) (*Termios, error) {
>>>>>>> deathstrox/main
	var value Termios
	if req != TCGETS {
		return &value, ENOSYS
	}
	err := Tcgetattr(fd, &value)
	return &value, err
}
