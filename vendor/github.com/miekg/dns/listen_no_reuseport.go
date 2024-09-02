<<<<<<< HEAD
//go:build !aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd
=======
// +build !go1.11 !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd
>>>>>>> deathstrox/main

package dns

import "net"

const supportsReusePort = false

<<<<<<< HEAD
func listenTCP(network, addr string, reuseport, reuseaddr bool) (net.Listener, error) {
	if reuseport || reuseaddr {
=======
func listenTCP(network, addr string, reuseport bool) (net.Listener, error) {
	if reuseport {
>>>>>>> deathstrox/main
		// TODO(tmthrgd): return an error?
	}

	return net.Listen(network, addr)
}

<<<<<<< HEAD
const supportsReuseAddr = false

func listenUDP(network, addr string, reuseport, reuseaddr bool) (net.PacketConn, error) {
	if reuseport || reuseaddr {
=======
func listenUDP(network, addr string, reuseport bool) (net.PacketConn, error) {
	if reuseport {
>>>>>>> deathstrox/main
		// TODO(tmthrgd): return an error?
	}

	return net.ListenPacket(network, addr)
}
