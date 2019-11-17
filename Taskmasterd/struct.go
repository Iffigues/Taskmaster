package main

import (
)


type Cmd struct {
	Path string
	Args []string
	Env  []string
	Dir  string
	Stdout interface{}
	Stderr interface{}
}

type task struct {
	lp      string
	cmds    Cmd
	live    int
	start   int
	restart bool
	reboot  int
	time    int
	count   int
	signal  int
	stop    []int
	umask   int
}
