package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"strings"
	"os/exec"
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

func make_cmd(fd *ini.File, ok string) (ar exec.Cmd) {
	ar.Path = getK(fd, ok, "com")
	ar.Args = getA(fd, ok, "args")
	ar.Env = getA(fd, ok, "env")
	ar.Dir = getK(fd, ok, "dir")
	return
}

func get(st string) (a map[string]*task, err error) {
	fd, err := ini.Load(st)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	a = make(map[string]*task)
	ar := fd.SectionStrings()
	for _, ok := range ar {
		if ok != "DEFAULT" {
			a[ok] = &task{
				cmds: make_cmd(fd, ok),
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
