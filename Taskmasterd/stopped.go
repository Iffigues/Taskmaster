package main

import (
	"fmt"
)

func is_stopped(ok, g bool) (yes bool, str string) {
	fmt.Println(ok, g)
	if !ok {
		return false, "jobs don't stop\n"
	}
	if g {
		return true, "job stopped\n"
	}
	return true, "job force to stop\n"
}
