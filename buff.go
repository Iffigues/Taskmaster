package main

import (
	 term "github.com/nsf/termbox-go"
)

func reset() {
	term.Sync()
}

func readder() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	keyPressListenerLoop:
		for {
			switch ev := term.PollEvent(); ev.Type {
			case  term.EventKey:
				switch ev.Key {
				case term.KeyArrowDown:
					println("hahaha")
				case term.KeyEnter:
					break keyPressListenerLoop
				}
			}
		}
	reset()
}

func init() {
	readder()
}
