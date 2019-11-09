package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
)

var (
	CONN_PORT = "3333"
)

func readerr(err error) {
	if err != nil {
		st := err.Error()
		if st == "Interrupt" {
			os.Exit(0)
		}
		if st == "EOF" {
			os.Exit(0)
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
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text + "\n"))
	}
}
