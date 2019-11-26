package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func fg(conn net.Conn, a ...string) (c ret, err error) {
	if val, ok := queued[a[0]]; ok {
		go func() {
			scanner := bufio.NewScanner(val.triade.StdOutPipe)
			for scanner.Scan() {
				_, err := conn.Write([]byte(scanner.Text()))
				fmt.Printf("\t > %s\n", scanner.Text())
				if err != nil {
				}
			}
		}()
		go func() {
			scanner := bufio.NewScanner(val.triade.StdErrPipe)
			for scanner.Scan() {
				fmt.Printf("\t > %s\n", scanner.Text())
			}
		}()
		for {
			buf := make([]byte, 1024)
			_, err := conn.Read(buf)
			if err != nil {
				break
			}
			io.WriteString(val.triade.StdInPipe, string(buf))
		}
	}
	_, err = conn.Write([]byte(""))
	return
}
