// +build !plan9,!linux

package singlefile

import (
	"crypto/sha1"
	"encoding/binary"
	"net"
	"strings"
)

func lock(key string) (func() error, error) {
	port, err := hashkey(key)
	if err != nil {
		return nil, err
	}
	addr := net.TCPAddr{
		IP:   net.IP{127, 0, 0, 1},
		Port: port,
	}
	l, err := net.ListenTCP("tcp4", &addr)
	return l.Close, err
}

func hashkey(key string) (int, error) {
	const userport = 1024
	r := strings.NewReader(key)
	hash := sha1.New()
	r.WriteTo(hash)
	sum := hash.Sum(nil)
	port := binary.BigEndian.Uint16(sum)
	port %= 1<<16 - userport
	return int(port) + userport, nil
}
