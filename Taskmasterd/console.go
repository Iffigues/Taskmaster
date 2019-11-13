package main

import (
	"fmt"
)

var (
	console = map[string]func(...string) error{
		"hello": hello,
	}
)

func hello(a ...string) (err error) {
	fmt.Println("hello world")
	return
}

func consoles(a ...string) (err error) {
	if len(a) > 0 {
		if e, d := console[a[0]]; d {
			if len(a) > 1 {
				e(a[1:]...)
			} else {
				e()
			}
		}
	}
	return
}
