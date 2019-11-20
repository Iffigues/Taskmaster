package main

import (
	"os/exec"
)

type Cmd struct {
	Path   string
	Args   []string
	Env    []string
	Dir    string
	Stdout interface{}
	Stderr interface{}
}

type task struct {
	lp       string
	cmds     Cmd
	live     int
	lancer   bool
	finish   bool
	start    int
	restart  bool
	reboot   int
	time     int
	count    int
	signal   int
	numprocs int
	stop     []int
	umask    int
	cmdl     *exec.Cmd
}

type ret struct {
	end bool
}

type enqued map[string]*task
