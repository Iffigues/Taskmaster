package main

import (
	"net"
	"os"
	"taskmasterd/helper/str"
)

const (
	CONN_HOST = "192.168.1.255"
)

var (
	CONN_PORT = "3333"
)

func serve() {
	l, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		os.Exit(1)
	}
	defer l.Close()
ff:
	for {
		conn, err := l.Accept()
		if err != nil {
			goto ff
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			if err.Error() != "EOF" {
			}
			break
		}
		if err != nil {
			break
		}
		b, err := consoles(conn, str.StrToStrArray(string(buf))...)
		if b.end {
			_, err = conn.Write([]byte("EOF"))
			break
		}
	}
}
