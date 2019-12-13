package main

import (
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
		if _, err := exec.LookPath(keys.cmds.Path); err != nil {
			return false
		}
		if (gg && jj) || (!gg) {
			cmd := exec.Command(keys.cmds.Path, keys.cmds.Args...)
			if len(keys.cmds.Dir) > 0 {
				cmd.Dir = keys.cmds.Dir
			}
			if len(keys.cmds.Env) > 0 {
				cmd.Env = keys.cmds.Env
			}
			cmd.Stdout = keys.cmds.Stdout
			cmd.Stderr = keys.cmds.Stderr
			keys.cmdl = cmd
			keys.stop = false
			keys.verif = make(chan bool)
			queued[a] = &keys
			return true
		}
		return false
	}
	return
}

func stop_command(a string) (ok, g bool) {
	existe, ok := is_started(a)
	if existe && !ok {
		if err := queued[a].cmdl.Process.Signal(queued[a].stopsignal); err != nil {
		}
		select {
		case <-time.After(time.Duration(queued[a].stoptime) * time.Second):
			queued[a].cmdl.Process.Kill()
		case <-queued[a].verif:
			g = true
			break
		}
		queued[a].stop = true
		return !ok, g
	}
	return !ok, g
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
