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
	for  {
		conn, err := l.Accept()
		if err != nil {
			os.Exit(1);
		}
		go handleRequest(conn)
	}
}

func lance() {
	go serve();
}

func handleRequest(conn net.Conn) {
	i := true
	for i {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		fmt.Println(buf)
		if err != nil {
			i = false
			fmt.Println("Error reading:", err.Error())

		}
		conn.Write([]byte("Message received.\n"))
	}
	defer conn.Close()
}

func client() {
	conn, err := net.Dial("tcp", CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("nat\n")
		}
		fmt.Fprintf(conn, text + "lol\n")
		messages, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("uyuyuy",err)
		}
		fmt.Println( "mess="+messages)
	}
}
