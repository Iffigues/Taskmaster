package main

import (
	"fmt"
	"net"
	"taskmasterd/helper/str"
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
		c := make(chan bool, 1)
		go lance(c, a)
		go getbol(c)
	}
}

func getbol(c chan bool) {
	<-c
}

func start(conn net.Conn, a ...string) (ce ret, err error) {
	strs := ""
	if len(a) > 0 {
		c := make(chan bool, 1)
		if a[0] == "all" {
			pad := padding()
			for key, _ := range jobs {
				go lance(c, key)
				e := meme(c, key, "jobs started\n", "jobs not found\n", "already start\n")
				strs = str.StrConcat(strs, key, ":", getWidth(len(key), pad), e)
			}
		} else {
			pad := getPadding(a)
			for _, key := range a {
				go lance(c, key)
				e := meme(c, key, "jobs started\n", "jobs not found\n", "already start\n")
				strs = str.StrConcat(strs, key, ":", getWidth(len(key), pad), e)
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
			pad := padding()
			for key, _ := range jobs {
				ok, g := stop_command(key)
				oui, d := is_stopped(ok, g)
				e := mami(oui, d, "jobs don't stop\n")
				strs = str.StrConcat(strs, key, ":", getWidth(len(key), pad), e)
			}
		} else {
			pad := getPadding(a)
			for _, key := range a {
				ok, g := stop_command(key)
				oui, d := is_stopped(ok, g)
				e := mami(oui, d, "jobs don't stop\n")
				strs = str.StrConcat(strs, key, ":", getWidth(len(key), pad), e)
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
		c := make(chan bool, 1)
		if a[0] == "all" {
			pad := padding()
			for key, _ := range jobs {
				fmt.Println(stop_command(key))
				go lance(c, key)
				e := meme(c, key, "started command\n", "jobs not found\n", "alredie start\n")
				strs = str.StrConcat(strs, key, ":", getWidth(len(key), pad), e)
			}
		} else {
			pad := getPadding(a)
			for _, key := range a {
				fmt.Println(stop_command(key))
				go lance(c, key)
				e := meme(c, key, "started command\n", "jobs not found\n", "alraidi start\n&")
				strs = str.StrConcat(strs, key, ":", getWidth(len(key), pad), e)
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
