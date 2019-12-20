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
	key, ok := queued[a]
	if ok {
		return true, (key.finish)
	}
	return false, false
}

func start_command(a string) (ok bool) {
	ok = false
	var keys task
	if keys, ok = jobs[a]; ok {
		mut.Lock()
		gg, jj := is_started(a)
		mut.Unlock()
		if _, err := exec.LookPath(keys.cmds.Path); err != nil {
			return false
		}
		if (gg && jj) || (!gg) {
			mut.Lock()
			cmd := exec.Command(keys.cmds.Path, keys.cmds.Args...)
			if len(keys.cmds.Dir) > 0 {
				cmd.Dir = keys.cmds.Dir
			}
			if len(keys.cmds.Env) > 0 {
				cmd.Env = keys.cmds.Env
			}
			sout, zz := stdout(keys.cmds.Stdout, a, keys.umask)
			serr, zzz := stderr(keys.cmds.Stdout, a, keys.umask)
			if zz != nil || zzz != nil {
				return false
			}
			keys.Stderr = serr
			keys.Stdout = sout
			cmd.Stdout = sout
			cmd.Stderr = serr
			keys.cmdl = cmd
			keys.stop = false
			keys.finish = true
			keys.verif = make(chan bool)
			queued[a] = &keys
			mut.Unlock()
			return true
		}
		return false
	}
	return
}

func stop_command(a string) (ok, g bool) {
	mut.Lock()
	existe, ok := is_started(a)
	mut.Unlock()
	if existe && !ok {
		mut.Lock()
		queued[a].stop, queued[a].finish = true, true
		mut.Unlock()
		if err := queued[a].cmdl.Process.Signal(queued[a].stopsignal); err != nil {
		}
		select {
		case <-time.After(time.Duration(queued[a].stoptime) * time.Second):
			queued[a].cmdl.Process.Kill()
		case <-queued[a].verif:
			g = true
		}
		return !ok, g
	}
	return !ok, g
}
