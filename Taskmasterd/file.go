package main

import (
	"io/ioutil"
	"os"
)

func stdout(keys, a string, umask int64) (r *os.File, err error) {
	if keys != "" && keys != "none" {
		if err := ioutil.WriteFile(keys, nil, os.FileMode(umask)); err != nil {
			return nil, err
		}
		f, err := os.OpenFile(keys, os.O_WRONLY|os.O_APPEND, os.FileMode(umask))
		if err != nil {
			return f, err
		}
		return f, nil
	}
	if keys == "none" {
		f, _ := os.OpenFile("/dev/null", os.O_WRONLY|os.O_APPEND, os.FileMode(umask))
		return f, err
	}
	if err := ioutil.WriteFile("./log/stdout/"+a, nil, os.FileMode(umask)); err != nil {
		return nil, err
	}
	ff, err := os.OpenFile("./log/stdout/"+a, os.O_WRONLY|os.O_APPEND, os.FileMode(umask))
	if err != nil {
		return nil, err
	}
	return ff, nil
}

func stderr(keys, a string, umask int64) (r *os.File, err error) {
	if keys != "" && keys != "none" {
		if err := ioutil.WriteFile(keys, nil, os.FileMode(umask)); err != nil {
			return nil, err
		}
		f, err := os.OpenFile(keys, os.O_WRONLY|os.O_APPEND, os.FileMode(umask))
		if err != nil {
			return f, err
		}
		return f, nil
	}
	if keys == "none" {
		f, _ := os.OpenFile("/dev/null", os.O_WRONLY|os.O_APPEND, os.FileMode(umask))
		return f, err
	}
	if err := ioutil.WriteFile("./log/stderr/"+a, nil, os.FileMode(umask)); err != nil {
		return nil, err
	}
	ff, err := os.OpenFile("./log/stderr/"+a, os.O_WRONLY|os.O_APPEND, os.FileMode(umask))
	if err != nil {
		return nil, err
	}
	return ff, nil
}
