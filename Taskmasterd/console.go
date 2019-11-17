package main

import (
	"fmt"
	"net"
	"os/exec"
)

var (
	console = map[string]func(net.Conn, ...string) (ret, error){
		"exit":    exit,
		"reload":  reload,
		"start":   start,
		"stop":    stop,
		"restart": restart,
	}
	queued enqued
)

func start(conn net.Conn, a ...string) (c ret, err error) {
	if key, ok := jobs[a[0]]; ok {
		cmd := exec.Command(key.cmds.Path, key.cmds.Args...)
		if len(key.cmds.Dir) > 0 {
			cmd.Dir = key.cmds.Dir
		}
		if len(key.cmds.Env) > 0 {
			cmd.Env = key.cmds.Env
		}
	}
	return
}

func stop(conn net.Conn, a ...string) (c ret, err error) {
	return
}

func restart(conn net.Conn, a ...string) (c ret, err error) {
	return
}

func reload(conn net.Conn, a ...string) (c ret, err error) {
	tmp, err := get("../ini/ini.ini")
	if err == nil {
		jobs = tmp
		conn.Write([]byte("new configuration load"))
	} else {
		conn.Write([]byte("bad init file"))
	}
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
				return e(conn, a[1:len(a)-1]...)
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
