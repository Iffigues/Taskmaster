package main

import (
	"net"
	"syscall"
	"taskmasterd/helper/str"
	"time"
)

func reload(conn net.Conn, a ...string) (c ret, err error) {
	syscall.Kill(mypid, syscall.SIGHUP)
	registre("reload", "taskmaster reload at:"+time.Now().String())
	conn.Write([]byte("reload"))
	return
}

func exit(conn net.Conn, a ...string) (c ret, err error) {
	c.end = true
	return
}

func begin() (err error) {
	for key, val := range jobs {
		if val.autostart {
			c := make(chan bool)
			go lance(c, key)
		}
	}
	return
}

func status(conn net.Conn, a ...string) (c ret, err error) {
	var t string
	keys := alpha()
	pad := padding()
	for _, u := range keys {
		width := getWidth(len(u), pad)
		y := str.StrConcat(u, ":", width)
		if i, ok := queued[u]; !ok {
			y = str.StrConcat(y, "not started")
		} else if i.abort {
			y = str.StrConcat(y, "abort")
		} else if i.stop {
			y = str.StrConcat(y, "stop")
		} else if i.finish {
			y = str.StrConcat(y, "finish")
		} else if i.lancer {
			y = str.StrConcat(y, percent(int32(i.cmdl.Process.Pid)))
		} else {
			y = str.StrConcat(y, "not started")
		}
		t = str.StrConcat(t, y, "\n")
	}
	if t == "" {
		conn.Write([]byte(" "))
		return
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
