package main

import (
	_ "github.com/chzyer/readline"
	"os"
	"supervisorctl/helper/str"
)

func main() {
	arg := os.Args
	if len(arg) == 1 {
		go fanny()
		client(false, "")
	} else {
		client(true, str.ArrayToStr(arg[1:]))
	}
}
