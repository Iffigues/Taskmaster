package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	go fanny()
}

func map_array_string() (a []string) {
	for key, _ := range jobs {
		a = append(a, key)
	}
	return
}

func send_me() (err error) {
	yy, err := get("./ini/ini.ini")
	if err != nil {
		return
	}
	a := map_array_string()
	for _, val := range a {
		if _, ok := yy[val]; !ok {
			nn, nnn := is_started(val)
			if nn && !nnn {
				go queued[val].cmdl.Process.Kill()
			}
			delete(jobs, val)
			delete(queued, val)
		}
	}
	for key, val := range yy {
		if ta, ok := jobs[key]; ok {
			fmt.Println(verify_change(ta, val))
			if verify_change(ta, val) {
				if _, oi := queued[key]; oi {
					nn, nnn := is_started(key)
					if nn && !nnn {
						fmt.Println(queued[key])
						queued[key].cmdl.Process.Kill()
					}
					delete(queued, key)
				}
				jobs[key] = val
				veve(key)
			}
		} else {
			jobs[key] = val
			veve(key)
		}
	}
	return
}

func fanny() {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGCONT,
		syscall.SIGWINCH,
		syscall.SIGTSTP,
	)
	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGHUP:
				send_me()
			case syscall.SIGINT:
				exit_chan <- 0
			case syscall.SIGTERM:
				exit_chan <- 0
			case syscall.SIGQUIT:
				exit_chan <- 0
			default:
			}
		}
	}()
	code := <-exit_chan
	for key, _ := range jobs {
		stop_command(key)
	}
	os.Exit(code)
}
