package main

import (
	"os"
	"os/exec"
	"syscall"
	"time"
)

type Cmd struct {
	Path   string
	Args   []string
	Env    []string
	Dir    string
	Stdout string
	Stderr string
}

type runner struct {
	runatfailed []string
	runatsucced []string
	runwhatever []string
}

type task struct {
	name         string
	Stderr       *os.File
	Stdout       *os.File
	status       bool
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
	umask        int64
	cmdl         *exec.Cmd
	verif        chan bool
	exectime     float64
	nbexec       int
	start        time.Time
	end          time.Time
	succed       bool
	failed       bool
	abort        bool
	group        []string
	grap         runner
}

type ret struct {
	end bool
}

type consolle struct {
	retrie int
	abort  int
	f      bool
}

type enqued map[string]*task
