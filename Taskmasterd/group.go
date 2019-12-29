package main

import (
	"github.com/go-ini/ini"
)

func getgroup(fd *ini.File, ok, gr string) (b []string, err error) {
	b, err = get_args(fd, ok, "group")
	if err != nil && !NotFound(err) {
		return
	}
	return
}

func getrun(fd *ini.File, ok string) (a runner, err error) {
	failed, err := get_args(fd, ok, "runatfailed")
	if err != nil && !NotFound(err) {
		return
	}
	a.runatfailed = failed
	succed, err := get_args(fd, ok, "runatsucced")
	if err != nil && !NotFound(err) {
		return
	}
	a.runatsucced = succed
	ever, err := get_args(fd, ok, "runwhatever")
	if err != nil && !NotFound(err) {
		return
	}
	a.runwhatever = ever
	return
}
