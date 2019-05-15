package main

import (
	"os/exec"
)

type task struct {
	cmds    exec.Cmd
	live    int
	restart bool
	reboot  int
	time    int
	count   int
	signal  int
	stop    int
	umask   int
}
