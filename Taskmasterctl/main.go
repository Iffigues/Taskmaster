package main

import (
	_ "github.com/chzyer/readline"
	"os"
	"taskmasterctl/helper/str"
)

func main() {
	arg := os.Args
	if len(arg) == 1 {
		client(false, "")
	} else {
		client(true, str.ArrayToStr(arg[1:]))
	}
}
