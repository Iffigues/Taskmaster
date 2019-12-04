package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func ps(i int) (ok bool) {
	ii := strconv.Itoa(i)
	cmd := exec.Command("/usr/bin/ps", "-p", ii)
	out, _ := cmd.CombinedOutput()
	st := strings.Split(string(out), "\n")
	if len(st) > 2 {
		return true
	}
	return false
}
