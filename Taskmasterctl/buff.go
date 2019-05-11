package main

import (
	"bytes"
	term "github.com/nsf/termbox-go"
)

func reset() {
	term.Sync()
}

func readder(ls []bytes.Buffer) (next bytes.Buffer){
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyArrowDown:
				println("hahaha")
			case term.KeyArrowUp:
				println("hihihi")
			case term.KeyArrowRight:
				println("huhuhu")
			case term.KeyArrowLeft:
				println("hihihi")
			case term.KeyEnter:
				return
			case term.KeyCtrlC:
				exit(0)
			default:
				println(ev.Ch)
			}
		}
	}
	reset()
	return ;
}

func init() {

}
