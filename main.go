package main

import (
	"os"
	"syscall"
)

var (
	mypid  = syscall.Getpid()
	option = map[string]func(){
		"--client": client,
		"--server": serve,
	}
)

func init() {
}

func main() {
	go fanny()
	opt := os.Args
	if len(opt) == 1 {
		prompt()
	}
	if len(opt) == 2 {
		if com, ok := option[opt[1]]; ok {
			com()
		}
	}
	os.Exit(0)
}
