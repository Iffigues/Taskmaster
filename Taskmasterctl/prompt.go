package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (a *task) test() {
	for {
	}
}

func prompt() {
	rr, err := get("./ini/ini.ini")
	if err != nil {
		fmt.Println(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		st := strings.Fields(text)
		rr["yes"].com = st[0]
	}
}
