package main

import ()

func verif_array(a, b []string) (ok bool) {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func verif_array_int(a, b []int) (ok bool) {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func verif_array_bool(a, b []bool) (ok bool) {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func verify_cmd(a, b Cmd) (ok bool) {
	ok = true
	if a.Path != b.Path || !verif_array(a.Args, b.Args) || !verif_array(a.Env, b.Env) || a.Dir != b.Dir || a.Stdout != b.Stdout || a.Stderr != a.Stderr {
		ok = false
	}
	return
}

func my_string_array(a task) (t []string, gg []int, hh []bool) {
	t = []string{
		a.lp,
	}
	gg = []int{
		a.autorestart,
		a.startretries,
		a.starttime,
		a.stoptime,
		a.numprocs,
	}
	hh = []bool{
		a.autostart,
	}
	return
}

func verify_change(a, b task) (ok bool) {
	ok = true
	if !verify_cmd(a.cmds, b.cmds) {
		return false
	}
	al, ag, ab := my_string_array(a)
	bl, bg, bb := my_string_array(b)
	if !verif_array(al, bl) {
		return false
	}
	if !verif_array_int(ag, bg) || !verif_array_int(a.exitcodes, b.exitcodes) {
		return false
	}
	if !verif_array_bool(ab, bb) {
		return false
	}
	if a.stopsignal != b.stopsignal {
		return false
	}
	if a.umask != b.umask {
		return false
	}
	return
}
