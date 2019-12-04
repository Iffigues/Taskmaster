package main

import (
	"io"
	"os/exec"
	"syscall"
	"time"
)

type Triade struct {
	StdErrPipe io.ReadCloser
	StdOutPipe io.ReadCloser
	StdInPipe  io.WriteCloser
}

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
	stop         bool
	autorestart  int
	startretries int
	memretries   int
	starttime    int
	stoptime     int
	stopsignal   syscall.Signal
	numprocs     int
	exitcodes    []int
	umask        int
	cmdl         *exec.Cmd
	triade       Triade
	verif        chan bool
	exectime     float64
	nbexec       int
	start        time.Time
	end          time.Time
	succed       bool
}

type ret struct {
	end bool
}

type enqued map[string]*task
