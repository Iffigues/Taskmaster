package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func fanny(t map[string]*task) {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGCONT,
		syscall.SIGWINCH,
		//syscall.SIGTSTP,
	)
	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGHUP:
				println("errerere")
			case syscall.SIGINT:
				fmt.Println("Warikomi")
				exit_chan <- 0
			case syscall.SIGTERM:
				fmt.Println("force stop")
				exit_chan <- 0
			case syscall.SIGQUIT:
				fmt.Println("stop and core dump")
				exit_chan <- 0
			default:
				fmt.Println("Unknown signal.")
			}
		}
	}()
	code := <-exit_chan
	os.Exit(code)
}
