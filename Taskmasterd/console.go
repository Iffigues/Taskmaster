package main

import (
	"fmt"
	"net"
)

type ret struct {
	end bool
}

var (
	console = map[string]func(net.Conn, ...string) (ret, error){
		"exit": exit,
	}
)

func reload(conn net.Conn, a ...string) (c ret, err error) {
	get("../ini/ini.ini")
	return
}

func exit(conn net.Conn, a ...string) (c ret, err error) {
	c.end = true
	return
}

func consoles(conn net.Conn, a ...string) (c ret, err error) {
	if len(a) > 0 && a[0] != "" {
		if e, d := console[a[0]]; d {
			if len(a) > 1 {
				return e(conn, a[1:]...)
			} else {
				return e(conn)
			}
		} else {
			fmt.Println(len(a[0]))
			_, err = conn.Write([]byte("bad command\n"))
			return
		}
	}
	_, err = conn.Write([]byte(""))
	return
}
