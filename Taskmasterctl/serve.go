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

func client(mod bool, str string) {
lab:
	conn, err := net.Dial("tcp", CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	c := make(chan Message)
	defer conn.Close()
	go receive(conn, c)
	reader := bufio.NewReader(os.Stdin)
	if mod {
		conn.Write([]byte(str + "\n"))
		fmt.Println("fggfgffg", <-c)
		return
	}
	for {
		text, _ := reader.ReadString('\n')
		if text != "\n" {
			conn.Write([]byte(text + "\n"))
			ddd := <-c
			fmt.Println(ddd)
			if ddd.Types == 1 {
				break
			}
			if text == "exit\n" {
				return
			}
		}
	}
	goto lab
}
