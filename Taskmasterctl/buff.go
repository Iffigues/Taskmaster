package main

import (
	term "github.com/nsf/termbox-go"
)

func reset() {
	term.Sync()
}

func readder() {
	var gg []byte = []byte{'\t'}
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	for {
		switch ev := term.PollRawEvent(gg); ev.Type {
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
				println("huhuhu")
			default:

				println(ev.Ch)
			}
		}
	}
	reset()
}

func init() {

}
