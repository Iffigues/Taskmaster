package main

import (
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"os/exec"
	"path/filepath"
	"taskmasterd/helper/str"
)

const ()

var (
	cfg, cfgErr = ini.Load("./conf/Taskmaster.conf")
)

func NotFound(err error) (vrai bool) {
	i := err.Error()
	v := "error when getting key of section"
	if i[:len(v)] == v {
		vrai = true
	}
	return
}

func getK(ar *ini.File, section, key string) (a string, err error) {
	bb, err := ar.Section(section).GetKey(key)
	if err != nil {
		return
	}
	a = bb.String()
	return
}

func get_args(ar *ini.File, section, key string) (a []string, err error) {
	bb, err := getK(ar, section, key)
	if err != nil {
		return
	}
	a = str.StrToStrArray(bb)
	return
}

func look_path(ar *ini.File, section, dir string) (f string, err error) {
	f, err = getK(ar, section, "commande")
	if f == "" {
		return "", errors.New("command can't be empty")
	}
	ff, err := exec.LookPath(f)
	if err != nil && dir != "" {
		info, err := os.Stat(dir + "/" + f)
		if os.IsNotExist(err) {
			fmt.Println(err)
			return "", err
		}
		if info.IsDir() {
			return "", err
		}
		f = dir + "/" + f
		f, err = filepath.Abs(f)
		if err != nil {
			return "", err
		}
		ff, err := exec.LookPath(f)
		return ff, err
	}
	ff, err = filepath.Abs(ff)
	return ff, err
}

func getA(ar *ini.File, section, key string) (a []string, err error) {
	bb, err := ar.Section(section).GetKey(key)
	if err != nil {
		return
	}
	jj := str.StrToStrArray(bb.String())
	for _, ok := range jj {
		hh, err := getK(ar, section, ok)
		if err != nil {
		}
		a = append(a, hh)
	}
	return
}
