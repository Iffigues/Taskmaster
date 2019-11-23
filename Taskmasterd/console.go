package main

import (
	"net"
	"syscall"
	"taskmasterd/helper/str"
	"time"
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
	queued = make(enqued)
)

func lance(a ...string) {
label:
	ok := start_command(a[0])
	if ok {
		cc := queued[a[0]]
		time.Sleep(time.Duration(cc.starttime) * time.Second)
		cc.cmdl.Start()
		b := cc.cmdl.Process.Pid
		cc.cmdl.Wait()
		cc.finish = true
		if cc.autorestart > 0 {
			goto label
		}
		if get_pid(b, a[0]) {
		}
	}
}

func start(conn net.Conn, a ...string) (c ret, err error) {
	if len(a) > 0 {
		go lance(a...)
	}
	conn.Write([]byte("bad init file"))
	return
}

func stop(conn net.Conn, a ...string) (c ret, err error) {
	stop_command(a[0])
	return
}

func restart(conn net.Conn, a ...string) (c ret, err error) {
	return
}

func reload(conn net.Conn, a ...string) (c ret, err error) {
	syscall.Kill(mypid, syscall.SIGHUP)
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
