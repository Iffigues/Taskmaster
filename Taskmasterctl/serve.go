package main

import (
	"bufio"
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
