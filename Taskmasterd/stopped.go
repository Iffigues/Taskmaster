package main

import ()

func is_stopped(ok, g bool) (yes bool, str string) {
	if !ok {
		return false, "jobs not stop\n"
	}
	if g {
		return true, "job stop\n"
	}
	return true, "job stopped\n"
}
