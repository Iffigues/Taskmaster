package main

import (
	"os/exec"
)

func is_started(a string) (ok bool) {
	var key *task
	ok = false
	if key, ok = queued[a]; ok {
		ok = key.finish
	}
	return false
}

func start_command(a string) (key task, ok bool) {
	ok = false
	if key, ok := jobs[a]; ok {
		if !is_started(a) {
			cmd := exec.Command(key.cmds.Path, key.cmds.Args...)
			if len(key.cmds.Dir) > 0 {
				cmd.Dir = key.cmds.Dir
			}
			if len(key.cmds.Env) > 0 {
				cmd.Env = key.cmds.Env
			}
			key.cmdl = cmd
			queued[a] = &key
		}
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
		if is_started(key) {
		}
	}
}
