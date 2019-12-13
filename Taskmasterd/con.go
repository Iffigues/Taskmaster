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
	for u, _ := range jobs {
		y := str.StrConcat(u, ":  ")
		if i, ok := queued[u]; ok {
			if i.abort {
				y = str.StrConcat(y, "abort")
				goto lab
			}
			if i.stop {
				y = str.StrConcat(y, " stop")
				goto lab
			}
			if i.finish {
				y = str.StrConcat(y, "finish")
				goto lab
			}
			if i.lancer {
				y = str.StrConcat(y, "start")
				goto lab
			}
		}
		y = str.StrConcat(y, "not started")
	lab:
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
