package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "51.255.43.50"
)

var (
	CONN_PORT = "3333"
)

func serve() {
	l, err := net.Listen("tcp", ":3333")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func lance() {
	go serve()
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
			fmt.Println(err.Error())
			break
		}
	}
}
