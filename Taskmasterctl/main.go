package main

import (
	_ "github.com/chzyer/readline"
	"supervisorctl/helper/str"
	"os"
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
