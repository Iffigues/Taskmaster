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
	l, err := set_read()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer l.Close()
	for {
		next, err := l.Readline()
		readerr(err)
		if len(next) > 0 {
			fmt.Fprintf(conn, next)
		}
	}
}
