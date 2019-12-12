package main

import (
	"net"
	"taskmasterd/helper/str"
	"time"
)

func mami(c bool, a, b string) (strs string) {
	if c {
		return a
	} else {
		return b
	}
}

func veve(a string) {
	if jobs[a].autostart {
		c := make(chan bool)
		go lance(c, a)
	}
}

func start(conn net.Conn, a ...string) (ce ret, err error) {
	strs := ""
	if len(a) > 0 {
		c := make(chan bool)
		if a[0] == "all" {
			for key, _ := range jobs {
				go lance(c, key)
				e := meme(c, "jobs started\n", "jobs not found\n")
				strs = str.StrConcat(strs, e)
			}
		} else {
			for _, key := range a {
				go lance(c, key)
				e := meme(c, "jobs started\n", "jobs not found\n")
				strs = str.StrConcat(strs, e)
			}
		}
	} else {
		conn.Write([]byte("bad init file"))
	}
	if strs == "" {
		conn.Write([]byte(" "))
	}
	conn.Write([]byte(strs))
	return
}

func stop(conn net.Conn, a ...string) (c ret, err error) {
	strs := ""
	if len(a) > 0 {
		if a[0] == "all" {
			for key, _ := range jobs {
				ok, g := stop_command(key)
				oui, d := is_stopped(ok, g)
				e := mami(oui, d, "jobs don't stop\n")
				strs = str.StrConcat(strs, e)
			}
		} else {
			for _, key := range a {
				ok, g := stop_command(key)
				oui, d := is_stopped(ok, g)
				e := mami(oui, d, "jobs don't stop\n")
				strs = str.StrConcat(strs, e)
			}
		}
	}
	if strs == "" {
		conn.Write([]byte(" "))
	}
	conn.Write([]byte(strs))
	return
}

func restart(conn net.Conn, a ...string) (c ret, err error) {
	strs := ""
	if len(a) > 0 {
		c := make(chan bool)
		if a[0] == "all" {
			for key, _ := range jobs {
				stop_command(key)
				time.Sleep(1 * time.Second)
				go lance(c, key)
				e := meme(c, "started command\n", "jobs not found\n")
				strs = str.StrConcat(strs, e)
			}
		} else {
			for _, key := range a {
				stop_command(key)
				time.Sleep(1 * time.Second)
				go lance(c, key)
				e := meme(c, "started command\n", "jobs not found\n")
				strs = str.StrConcat(strs, e)
			}
		}
	}
	if strs == "" {
		conn.Write([]byte(" "))
		return
	}
	conn.Write([]byte(strs))
	return
}
