package singlefile

import (
	"net"
	"path"
)

func lock(key string) (func() error, error) {
	addr := net.UnixAddr{
		Name: path.Join("@singlefile", key),
		Net:  "unix",
	}
	l, err := net.ListenUnix("unix", &addr)
	return l.Close, err
}
