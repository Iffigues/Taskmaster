package main

import (
	"github.com/go-ini/ini"
	"log"
	"os/exec"
	"strconv"
	"supervisord/helper/str"
)

const ()

var (
	cfg, cfgErr = ini.Load("./conf/Taskmaster.conf")
)

func getK(ar *ini.File, section, key string) (a string, err error) {
	bb, err := ar.Section(section).GetKey(key)
	if err != nil {
		return
	}
	a = bb.String()
	return
}

func look_path(ar *ini.File, section string, m map[string]int) (f string, err error) {
	f, err = getK(ar, section, "commande")
	if f == "" || err != nil {
		return
	}
	f, err = exec.LookPath(f)
	return
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

func getStd(ar *ini.File, section, key string) (a string, err error) {
	bb, err := ar.Section(section).GetKey(key)
	if err != nil {
		return
	}
	a = bb.String()
	return
}

func getumask(ar *ini.File, section string, m map[string]int) (a int) {
	oc, _ := strconv.ParseInt("0666", 8, 64)
	bb, err := ar.Section(section).GetKey("umask")
	if err != nil {
		oc, _ = strconv.ParseInt("022", 8, 64)
		return int(oc)
	}
	v, _ := strconv.Atoi(bb.String())
	aa := oc &^ int64(v)
	return int(aa)
}

func get_int_array(ar *ini.File, section, key string) (d []int, err error) {
	strs, err := getK(ar, section, key)
	if err != nil {
		return nil, err
	}
	d, err = str.StrToIntArray(strs)
	return
}

func make_cmd(fd *ini.File, ok string, m map[string]int) (ar exec.Cmd, err error) {
	ar.Path, err = getK(fd, ok, "com")
	ar.Args, err = getA(fd, ok, "args")
	ar.Env, err = getA(fd, ok, "env")
	ar.Dir, err = getK(fd, ok, "dir")
	a, err := getStd(fd, ok, "stdout")
	if a != "" {
	}
	a, err = getStd(fd, ok, "stderr")
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
			bb := str.ArrayToMap(fd.Section(ok).KeyStrings())
			PATH, err := look_path(fd, ok, bb)
			CMD, err := make_cmd(fd, ok, bb)
			UMASK := getumask(fd, ok, bb)
			stop, err := get_int_array(fd, ok, "stop")
			if err != nil {
			}
			a[ok] = &task{
				lp:    PATH,
				cmds:  CMD,
				umask: UMASK,
				stop:  stop,
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
