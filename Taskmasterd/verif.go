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

func verif_arra_inty(a, b []int) (ok bool) {
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

func verify_change(a, b task) {

}
