package main

import (
	"errors"
	"github.com/go-ini/ini"
	"strconv"
	"syscall"
	"taskmasterd/helper/str"
)

func getumask(ar *ini.File, section string) (a int64, err error) {
	oc, err := strconv.ParseInt("0666", 8, 64)
	bb, err := ar.Section(section).GetKey("umask")
	if err != nil {
		err = nil
		return oc, err
	}
	aa, err := strconv.ParseInt(bb.String(), 8, 64)
	return aa, err
}

func get_int_array(ar *ini.File, section, key string) (d []int, err error) {
	strs, err := getK(ar, section, key)
	if err != nil && !NotFound(err) {
		return nil, err
	}
	if err != nil && NotFound(err) {
		return nil, nil
	}
	d, err = str.StrToIntArray(strs)
	return
}

func nprocess(ar *ini.File, section, key string) (d, yy int, err error) {
	strs, err := getK(ar, section, key)
	if err != nil && NotFound(err) {
		return 1, 0, nil
	}
	if err != nil {
		return
	}
	d, err = strconv.Atoi(strs)
	return d, 1, err
}

func make_cmd(fd *ini.File, ok string, umask int64) (ar Cmd, PATH string, err error) {
	ll, err := get_args(fd, ok, "args")
	if err != nil && !NotFound(err) {
		return
	}
	ar.Args = ll
	lll, err := getA(fd, ok, "env")
	if err != nil && !NotFound(err) {
		return
	}
	ar.Env = lll
	llll, err := getK(fd, ok, "workingdir")
	if err != nil && !NotFound(err) {
		return
	}
	ar.Dir = llll
	PATH, err = look_path(fd, ok, llll)
	if err != nil {
		return
	}
	ar.Path = PATH
	dd, err := getK(fd, ok, "stdout")
	if err != nil && !NotFound(err) {
		return
	}
	ar.Stdout = dd
	ddd, err := getK(fd, ok, "stderr")
	if err != nil && !NotFound(err) {
		return
	}
	ar.Stderr = ddd
	return
}

func getsignal(fd *ini.File, ok, path string) (ff syscall.Signal, err error) {
	oo := map[string]syscall.Signal{
		"TERM": syscall.SIGTERM,
		"HUP":  syscall.SIGHUP,
		"INT":  syscall.SIGINT,
		"KILL": syscall.SIGKILL,
		"USR1": syscall.SIGUSR1,
		"USR2": syscall.SIGUSR2,
	}
	ll, err := getK(fd, ok, path)
	if err != nil {
		if NotFound(err) {
			return oo["TERM"], nil
		}
		return syscall.SIGKILL, err
	}
	if val, ik := oo[ll]; ik {
		return val, nil
	}
	return syscall.SIGKILL, errors.New("can't work with 42")
}
