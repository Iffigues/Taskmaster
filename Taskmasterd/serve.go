package main

import (
	"fmt"
	"net"
	"os"
	"supervisord/helper/str"
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
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	ff:
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
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
			fmt.Println(err.Error())
			break
		}
		_, err = conn.Write([]byte(buf))
		if err != nil {
			fmt.Println("err=",err.Error())
			break
		}
		fmt.Println(str.StrToStrArray(string(buf)))
	}
}
