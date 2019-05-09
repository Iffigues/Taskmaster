package main

import (
	"bufio"
	"fmt"
	"os"
)

func prompt() {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(text)
	}
}
