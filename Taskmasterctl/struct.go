package main

import (
	"os/exec"
)

type task struct {
	cmds    exec.Cmd
	live    int
	com     string
	restart bool
	reboot  int
	code    int
	time    int
	count   int
	signal  int
	stop    int
	stdout  string
	stderr  string
	env     []string
	work    string
	umask   int
}

type Message struct {
	Types int
	Mess  string
}
