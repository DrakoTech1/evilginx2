// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 && freebsd
<<<<<<< HEAD
=======
// +build 386,freebsd
>>>>>>> deathstrox/main

package unix

import (
	"syscall"
	"unsafe"
)

func setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: int32(sec), Nsec: int32(nsec)}
}

func setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: int32(sec), Usec: int32(usec)}
}

func SetKevent(k *Kevent_t, fd, mode, flags int) {
	k.Ident = uint32(fd)
	k.Filter = int16(mode)
	k.Flags = uint16(flags)
}

func (iov *Iovec) SetLen(length int) {
	iov.Len = uint32(length)
}

func (msghdr *Msghdr) SetControllen(length int) {
	msghdr.Controllen = uint32(length)
}

func (msghdr *Msghdr) SetIovlen(length int) {
	msghdr.Iovlen = int32(length)
}

func (cmsg *Cmsghdr) SetLen(length int) {
	cmsg.Len = uint32(length)
}

<<<<<<< HEAD
func (d *PtraceIoDesc) SetLen(length int) {
	d.Len = uint32(length)
}

=======
>>>>>>> deathstrox/main
func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	var writtenOut uint64 = 0
	_, _, e1 := Syscall9(SYS_SENDFILE, uintptr(infd), uintptr(outfd), uintptr(*offset), uintptr((*offset)>>32), uintptr(count), 0, uintptr(unsafe.Pointer(&writtenOut)), 0, 0)

	written = int(writtenOut)

	if e1 != 0 {
		err = e1
	}
	return
}

func Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno)

func PtraceGetFsBase(pid int, fsbase *int64) (err error) {
<<<<<<< HEAD
	return ptracePtr(PT_GETFSBASE, pid, unsafe.Pointer(fsbase), 0)
=======
	return ptrace(PT_GETFSBASE, pid, uintptr(unsafe.Pointer(fsbase)), 0)
}

func PtraceIO(req int, pid int, addr uintptr, out []byte, countin int) (count int, err error) {
	ioDesc := PtraceIoDesc{Op: int32(req), Offs: (*byte)(unsafe.Pointer(addr)), Addr: (*byte)(unsafe.Pointer(&out[0])), Len: uint32(countin)}
	err = ptrace(PT_IO, pid, uintptr(unsafe.Pointer(&ioDesc)), 0)
	return int(ioDesc.Len), err
>>>>>>> deathstrox/main
}
