package main

import (
	"net"
)

func fg(conn net.Conn, a ...string) (c ret, err error) {
	_, err = conn.Write([]byte(""))
	return
}
