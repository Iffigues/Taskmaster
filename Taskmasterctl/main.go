package main

import (
	_ "github.com/chzyer/readline"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) == 1 {
		go fanny()
		client(false)
	} else  {
		
	}
}
