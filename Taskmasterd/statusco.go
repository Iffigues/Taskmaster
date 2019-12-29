package main

import (
	"net"
	"taskmasterd/helper/str"
)

func getbak() (t []string) {
	for key, _ := range queued {
		if _, ok := jobs[key]; !ok {
			t = append(t, key)
		}
	}
	return
}

func statusco(conn net.Conn, a ...string) (c ret, err error) {
	gg := getAlpha(getbak()...)
	pad := getPadding(gg)
	var t string
	for _, st := range gg {
		width := getWidth(len(st), pad)
		y := str.StrConcat(st, ":", width)
		i := queued[st]
		if i.abort {
			y = str.StrConcat(y, "abort")
		} else if i.stop {
			y = str.StrConcat(y, "stop")
		} else if i.failed {
			y = str.StrConcat(y, "failed")
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
