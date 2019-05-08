package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func fanny() {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGHUP:
				fmt.Println("hungup")
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
				exit_chan <- 1
			}
		}
	}()
	code := <-exit_chan
	os.Exit(code)
}
