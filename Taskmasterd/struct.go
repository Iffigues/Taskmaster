package main

import (
	"os/exec"
)

type task struct {
	lp      string
	cmds    exec.Cmd
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
