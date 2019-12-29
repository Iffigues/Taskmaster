package main

import (
	"net"
	"taskmasterd/helper/str"
)

func findgroup(a string) (b []string) {
	for key, val := range jobs {
		for _, hhh := range val.group {
			if hhh == a {
				b = append(b, key)
			}
		}
	}
	return
}

func startasgroup(conn net.Conn, a ...string) (c ret, err error) {
	d := make(chan bool, 1)
	strs := ""
	if len(a) > 0 {
		for _, v := range a {
			t := findgroup(v)
			pad := getPadding(t)
			for _, f := range t {
				go lance(d, f)
				i := <-d
				e := meme(i, f, "jobs started\n", "jobs not found\n", "already start\n")
				strs = str.StrConcat(strs, f, ":", getWidth(len(f), pad), e)
			}
		}
	}
	if strs == "" {
		conn.Write([]byte(" "))
	}
	conn.Write([]byte(strs))
	return
}

func stopasgroup(conn net.Conn, a ...string) (c ret, err error) {
	strs := ""
	if len(a) > 0 {
		for _, v := range a {
			t := findgroup(v)
			pad := getPadding(a)
			for _, f := range t {
				if _, ok := queued[f]; ok {
					ik, g := stop_command(f)
					oui, d := is_stopped(ik, g)
					e := mami(oui, d, "jobs don't stop\n")
					strs = str.StrConcat(strs, f, ":", getWidth(len(f), pad), e)
				}
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

func restartasgroup(conn net.Conn, a ...string) (c ret, err error) {
	strs := ""
	d := make(chan bool, 1)
	if len(a) > 0 {
		for _, v := range a {
			t := findgroup(v)
			pad := getPadding(a)
			for _, f := range t {
				stop_command(f)
				go lance(d, f)
				i := <-d
				e := meme(i, f, "started command\n", "jobs not found\n", "alraidi start\n")
				strs = str.StrConcat(strs, f, ":", getWidth(len(f), pad), e)
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

func killasgroup(conn net.Conn, a ...string) (c ret, err error) {
	strs := ""
	if len(a) > 0 {
		for _, v := range a {
			t := findgroup(v)
			for _, f := range t {
				if val, ok := queued[f]; ok {
					if val.lancer {
						err := val.cmdl.Process.Kill()
						if err != nil {
							strs = strs + f + ": " + err.Error() + "\n"
						} else {
							strs = strs + f + ": " + "process kill\n"
						}
					}
				}
			}
		}
		if strs == "" {
			conn.Write([]byte(" "))
			return
		}
	}
	if strs == "" {
		conn.Write([]byte(" "))
		return
	}
	conn.Write([]byte(strs))
	return
}
