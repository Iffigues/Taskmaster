package main

import (
	"github.com/go-ini/ini"
	"net"
	"strconv"
)

func getStd(ar *ini.File, section, key string) (a string, err error) {
	bb, err := ar.Section(section).GetKey(key)
	if err != nil {
		return
	}
	a = bb.String()
	return
}

func pid(conn net.Conn, a ...string) (c ret, err error) {
	if len(a) == 0 {
		conn.Write([]byte(strconv.Itoa(mypid)))
	} else {
		if val, ok := queued[a[0]]; ok {
			conn.Write([]byte(strconv.Itoa(val.cmdl.Process.Pid)))
		} else {
			conn.Write([]byte("no process found"))
		}
	}
	return
}
