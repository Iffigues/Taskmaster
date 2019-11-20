package main

import (
	"net"
	"taskmasterd/helper/str"
)

var (
	console = map[string]func(net.Conn, ...string) (ret, error){
		"exit":    exit,
		"reload":  reload,
		"start":   start,
		"stop":    stop,
		"status":  status,
		"restart": restart,
	}
	queued enqued
)

func start(conn net.Conn, a ...string) (c ret, err error) {
	if len(a) > 0 {
		_, ok := start_command(a[0])
		if ok {
			go func() {
				queued[a[0]].cmdl.Wait()
				queued[a[0]].finish = true
			}()
		}
	}
	conn.Write([]byte("bad init file"))
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

func begin() (err error) {
	return
}

func status(conn net.Conn, a ...string) (c ret, err error) {
	var t string
	for u, _ := range jobs {
		t = str.StrConcat(t, u, "\n")
	}
	conn.Write([]byte(t))
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
			_, err = conn.Write([]byte("bad command\n"))
			return
		}
	}
	_, err = conn.Write([]byte(""))
	return
}
