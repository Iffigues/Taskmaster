package main

import (
	"github.com/go-ini/ini"
	"log"
	"os/exec"
	"strconv"
	"syscall"
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

func look_path(ar *ini.File, section string) (f string, err error) {
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

func getumask(ar *ini.File, section string) (a int, err error) {
	oc, err := strconv.ParseInt("0666", 8, 64)
	bb, err := ar.Section(section).GetKey("umask")
	if err != nil && !NotFound(err) {
		return
	}
	if err != nil && NotFound(err) {
		err = nil
		oc, err = strconv.ParseInt("022", 8, 64)
		return int(oc), err
	}
	v, err := strconv.Atoi(bb.String())
	if err != nil {
		return
	}
	aa := oc &^ int64(v)
	return int(aa), nil
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

func make_cmd(fd *ini.File, ok, path string) (ar Cmd, err error) {
	ar.Path = path
	ll, err := getA(fd, ok, "args")
	if err != nil && !NotFound(err) {
		return
	}
	ar.Args = ll
	lll, err := getA(fd, ok, "env")
	if err != nil && !NotFound(err) {
		return
	}
	ar.Env = lll
	llll, err := getK(fd, ok, "dir")
	if err != nil && !NotFound(err) {
		return
	}
	ar.Dir = llll
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

func get(st string) (a map[string]task, err error) {
	fd, err := ini.Load(st)
	if err != nil {
		return nil, err
	}
	a = make(map[string]task)
	ar := fd.SectionStrings()
	for _, ok := range ar {
		if ok != "DEFAULT" {
			PATH, err := look_path(fd, ok)
			if err != nil {
				return nil, err
			}
			CMD, err := make_cmd(fd, ok, PATH)
			UMASK, err := getumask(fd, ok)
			if err != nil && !NotFound(err) {
				return nil, err
			}
			stop, err := get_int_array(fd, ok, "stop")
			if err != nil && !NotFound(err) {
				return nil, err
			}
			numprocs, vvv, err := nprocess(fd, ok, "numprocs")
			if err != nil {
				return nil, err
			}
			sig := syscall.SIGKILL
			for y := 0; y < numprocs; y++ {
				name := namer(ok, vvv, y)
				a[name] = task{
					lp:         PATH,
					lancer:     false,
					finish:     false,
					cmds:       CMD,
					umask:      UMASK,
					exitcodes:  stop,
					numprocs:   numprocs,
					stopsignal: sig,
				}
			}
		}
	}
	err = nil
	return
}

func getKey(section, key string) (inu string) {
	ar, err := cfg.Section(section).GetKey(key)
	if err != nil {
		log.Panic(err)
	}
	return ar.String()
}
