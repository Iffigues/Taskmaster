package main

import (
	"io/ioutil"
	"os"
)

func stdout(keys task, a string) (r *os.File, err error) {
	if keys.cmds.Stdout != "" && keys.cmds.Stdout != "none" {
		if err := ioutil.WriteFile(keys.cmds.Stdout, nil, os.FileMode(keys.umask)); err != nil {
			return nil, err
		}
		f, err := os.OpenFile(keys.cmds.Stdout, os.O_WRONLY|os.O_APPEND, os.FileMode(keys.umask))
		if err != nil {
			return f, err
		}
		return f, nil
	}
	if keys.cmds.Stdout == "none" {
		f, _ := os.OpenFile("/dev/null", os.O_WRONLY|os.O_APPEND, os.FileMode(keys.umask))
		return f, err
	}
	if err := ioutil.WriteFile("../log/stdout/"+a, nil, os.FileMode(keys.umask)); err != nil {
		return nil, err
	}
	ff, err := os.OpenFile("../log/stdout/"+a, os.O_WRONLY|os.O_APPEND, os.FileMode(keys.umask))
	if err != nil {
		return nil, err
	}
	return ff, nil
}

func stderr(keys task, a string) (r *os.File, err error) {
	if keys.cmds.Stderr != "" && keys.cmds.Stderr != "none" {
		if err := ioutil.WriteFile(keys.cmds.Stderr, nil, os.FileMode(keys.umask)); err != nil {
			return nil, err
		}
		f, err := os.OpenFile(keys.cmds.Stderr, os.O_WRONLY|os.O_APPEND, os.FileMode(keys.umask))
		if err != nil {
			return f, err
		}
		return f, nil
	}
	if keys.cmds.Stderr == "none" {
		f, _ := os.OpenFile("/dev/null", os.O_WRONLY|os.O_APPEND, os.FileMode(keys.umask))
		return f, err
	}
	if err := ioutil.WriteFile("../log/stderr/"+a, nil, os.FileMode(keys.umask)); err != nil {
		return nil, err
	}
	ff, err := os.OpenFile("../log/stderr/"+a, os.O_WRONLY|os.O_APPEND, os.FileMode(keys.umask))
	if err != nil {
		return nil, err
	}
	return ff, nil
}
