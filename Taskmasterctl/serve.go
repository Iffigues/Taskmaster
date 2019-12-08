package main

import (
	"fmt"
	"github.com/chzyer/readline"
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

func receive(conn net.Conn, c chan Message) {
	for {
		messages := make([]byte, 1024)
		lens, err := conn.Read(messages)
		if err != nil {
			c <- Message{1, err.Error()}
			return
		}
		if lens > 0 {
			c <- Message{0, string(messages)}
		}
	}
}

func sendy(con net.Conn, y string, c chan Message) (b Message) {
	if len(y) > 1 {
		con.Write([]byte(y + "\n"))
		b = <-c
		fmt.Println(b.Mess)
	}
	return
}

func client(mod bool, str string) {
	c := make(chan Message)
lab:
	conn, err := net.Dial("tcp", CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer conn.Close()
	go receive(conn, c)
	if mod {
		sendy(conn, str, c)
		return
	}
	for {
		text, err := term()
		if err == readline.ErrInterrupt {
			text = "exit"
		}
		ddd := sendy(conn, text, c)
		if ddd.Types == 1 {
			break
		}
		if len(text) >= 4 && text[0:4] == "exit" {
			return
		}
	}
	goto lab
}
