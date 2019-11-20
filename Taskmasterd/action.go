package main

import (
	"os/exec"
)

func start_command(a string) (key task, ok bool) {
	ok = false
	if key, ok = jobs[a]; ok {
		cmd := exec.Command(key.cmds.Path, key.cmds.Args...)
		if len(key.cmds.Dir) > 0 {
			cmd.Dir = key.cmds.Dir
		}
		if len(key.cmds.Env) > 0 {
			cmd.Env = key.cmds.Env
		}
		key.cmdl = cmd
		queued[a] = key
	}
	return
}
