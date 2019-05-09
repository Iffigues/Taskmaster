package main

import (
	"os"
)

var (
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
	if len(opt) == 0 {
	} else {
		if com, ok := option[opt[1]]; ok {
			com()
		}
	}
}
