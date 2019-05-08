package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_TYPE = "tcp"
)

var (
	CONN_PORT = "3333"
)

func serve() {
	i := true
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	for i {
		conn, err := l.Accept()
		if err != nil {
			i = false
		} else {
			go handleRequest(conn)
		}
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	fmt.Println(reqLen)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	conn.Write([]byte("Message received."))
	conn.Close()
}

func client() {
	conn, err := net.Dial("tcp", CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("not\n")
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("nat\n")
		}
		fmt.Fprintf(conn, text+"\n")
		messages, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("nit\n")
		}
		fmt.Print("Message from server: " + messages)
	}
}
