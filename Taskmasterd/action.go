package main

import (
	"os/exec"
)

func get_pid(a int, c string) (ok bool) {
	ok = false
	var key *task
	if key, ok = queued[c]; ok {
		return (key.cmdl.Process.Pid == a)
	}
	return
}

func is_started(a string) (ok bool) {
	var key *task
	ok = false
	if key, ok = queued[a]; ok {
		ok = key.finish
	}
	return
}

func start_command(a string) (ok bool) {
	ok = false
	var keys task
	if keys, ok = jobs[a]; ok {
		if !is_started(a) {
			cmd := exec.Command(keys.cmds.Path, keys.cmds.Args...)
			if len(keys.cmds.Dir) > 0 {
				cmd.Dir = keys.cmds.Dir
			}
			if len(keys.cmds.Env) > 0 {
				cmd.Env = keys.cmds.Env
			}
			keys.cmdl = cmd
			queued[a] = &keys
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
