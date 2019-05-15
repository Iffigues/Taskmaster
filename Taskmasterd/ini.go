package main

import (
	"github.com/go-ini/ini"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const ()

var (
	cfg, cfgErr = ini.Load("./conf/Taskmaster.conf")
)

func getK(ar *ini.File, section, key string) (a string) {
	bb, err := ar.Section(section).GetKey(key)
	if err != nil {
		return
	}
	a = bb.String()
	return
}

func getA(ar *ini.File, section, key string) (a []string) {
	bb, err := ar.Section(section).GetKey(key)
	if err != nil {
		return
	}
	a = strings.Fields(bb.String())
	return
}

func getStd(ar *ini.File, section, key string) (a string) {
	bb, err := ar.Section(section).GetKey(key)
	if err != nil {
		return
	}
	a = bb.String()
	return
}

func getumask(ar *ini.File, section string) (a int) {
	oc, _ := strconv.ParseInt("0666", 8, 64)
	bb, err := ar.Section(section).GetKey("umsak")
	if err != nil {
		oc, _ = strconv.ParseInt("022", 8, 64)
		return int(oc)
	}
	v, _ := strconv.Atoi(bb.String())
	aa := oc &^ int64(v)
	return int(aa)
}

func make_cmd(fd *ini.File, ok string) (ar exec.Cmd) {
	ar.Path = getK(fd, ok, "com")
	ar.Args = getA(fd, ok, "args")
	ar.Env = getA(fd, ok, "env")
	ar.Dir = getK(fd, ok, "dir")
	a := getStd(fd, ok, "stdout")
	if a != "" {
	}
	a = getStd(fd, ok, "stderr")
	if a != "" {
	}
	return
}

func get(st string) (a map[string]*task, err error) {
	fd, err := ini.Load(st)
	if err != nil {
		return nil, err
	}
	a = make(map[string]*task)
	ar := fd.SectionStrings()
	for _, ok := range ar {
		if ok != "DEFAULT" {
			a[ok] = &task{
				cmds:  make_cmd(fd, ok),
				umask: getumask(fd, ok),
			}
		}
	}
	return
}

func getKey(section, key string) (inu string) {
	ar, err := cfg.Section(section).GetKey(key)
	if err != nil {
		log.Panic(err)
	}
	return ar.String()
}
