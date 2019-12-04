package main

import (
	"fmt"
	"os"
	"os/signal"
)

func fanny(b []string, ii int, f ...os.Signal) {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, f...)
	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			default:
				t := true
				fmt.Println(s, b)
				for _, val := range f {
					if val == s {
						fmt.Println("hjjhhj")
						t = false
					}
				}
				if t {
					exit_chan <- 0
				}
			}
		}
	}()
	<-exit_chan
	os.Exit(ii)
}
