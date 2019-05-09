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
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
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
		_, err = conn.Write([]byte("oui"))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}

func receive(conn net.Conn) {
	for {
		messages := make([]byte, 1024)
		lens, err := conn.Read(messages)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		if lens > 0 {
			fmt.Println("mess=" + string(messages))
		}
	}
}

func client() {
	conn, err := net.Dial("tcp", CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer conn.Close()
	go receive(conn)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Fprintf(conn, text+"lol\n")
	}
}
