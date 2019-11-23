package main

import (
	"os/exec"
)

type Cmd struct {
	Path   string
	Args   []string
	Env    []string
	Dir    string
	Stdout string
	Stderr string
}

type task struct {
	lp           string
	cmds         Cmd
	live         int
	lancer       bool
	finish       bool
	autostart    bool
	autorestart  int
	startretries int
	starttime    int
	stoptime     int
	stopsignal   string
	numprocs     int
	exitcodes    []int
	umask        int
	cmdl         *exec.Cmd
}

type ret struct {
	end bool
}

type enqued map[string]*task
