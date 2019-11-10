package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"net"
	"os"
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
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handle(rl *readline.Instance) {
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		fmt.Fprintln(rl.Stdout(), "receive:"+line)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	go fanny()
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
