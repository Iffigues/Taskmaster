package main

import (
	"os/exec"
)

func (Task *task) cmd() {
	var cmds exec.Cmd
	cmds.Path = Task.com
}
