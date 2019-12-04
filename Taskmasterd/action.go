package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func get_pid(a int, c string) (ok bool) {
	ok = false
	var key *task
	if key, ok = queued[c]; ok {
		return (key.cmdl.Process.Pid == a)
	}
	return
}

func is_started(a string) (existe, ok bool) {
	var key *task
	ok = false
	if key, ok = queued[a]; ok {
		return true, key.finish
	}
	return false, false
}

func start_command(a string) (ok bool) {
	ok = false
	var keys task
	if keys, ok = jobs[a]; ok {
		gg, jj := is_started(a)
		if (gg && jj) || (!gg) {
			cmd := exec.Command(keys.cmds.Path, keys.cmds.Args...)
			if len(keys.cmds.Dir) > 0 {
				cmd.Dir = keys.cmds.Dir
			}
			if len(keys.cmds.Env) > 0 {
				cmd.Env = keys.cmds.Env
			}
			if keys.cmds.Stdout == "" {
				cmd.Stdout = os.Stdout
			}
			if keys.cmds.Stderr == "" {
				cmd.Stderr = os.Stderr
			}
			if keys.cmds.Stdout != "" {
				if err := ioutil.WriteFile(keys.cmds.Stdout, nil, os.FileMode(keys.umask)); err != nil {
					return false
				}
				f, err := os.OpenFile(keys.cmds.Stdout, os.O_WRONLY|os.O_APPEND, os.FileMode(keys.umask))
				if err != nil {
					return false
				}
				cmd.Stdout = f
			} else {
				f, _ := os.OpenFile("/dev/null", os.O_WRONLY|os.O_APPEND, os.FileMode(keys.umask))
				cmd.Stdout = f
				//keys.triade.StdOutPipe, _ = cmd.StdoutPipe()
			}
			if keys.cmds.Stderr != "" {
				if err := ioutil.WriteFile(keys.cmds.Stderr, nil, os.FileMode(keys.umask)); err != nil {
					return false
				}
				ff, err := os.OpenFile(keys.cmds.Stderr, os.O_WRONLY|os.O_APPEND, os.FileMode(keys.umask))
				if err != nil {
					return false
				}
				cmd.Stderr = ff
			} else {
				f, _ := os.OpenFile("/dev/null", os.O_WRONLY|os.O_APPEND, os.FileMode(keys.umask))
				cmd.Stderr = f
				//keys.triade.StdErrPipe, _ = cmd.StderrPipe()
			}
			keys.cmdl = cmd
			keys.stop = false
			keys.verif = make(chan bool)
			keys.triade.StdInPipe, _ = cmd.StdinPipe()
			queued[a] = &keys
			return true
		}
		return false
	}
	return
}

func stop_command(a string) (ok bool) {
	existe, ok := is_started(a)
	if existe && !ok {
		if err := queued[a].cmdl.Process.Signal(queued[a].stopsignal); err != nil {
		}
		select {
		case <-time.After(time.Duration(queued[a].stoptime) * time.Second):
			queued[a].cmdl.Process.Kill()
		case <-queued[a].verif:
			goto ha
		}
	ha:
		queued[a].stop = true
		return true
	}
	return
}

func start_all_command() {
	for key, _ := range jobs {
		start_command(key)
	}
}

func stop_all_command() {
	for key, _ := range jobs {
		is_started(key)
	}
}
