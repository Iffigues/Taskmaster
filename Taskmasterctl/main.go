package main

import (
	_ "github.com/chzyer/readline"
)

func main() {
	go fanny()
	client()
}
