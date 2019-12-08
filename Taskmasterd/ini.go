package main

import (
	"errors"
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

func get_args(ar *ini.File, section, key string) (a []string, err error) {
	bb, err := getK(ar, section, key)
	if err != nil {
		return
	}
	a = str.StrToStrArray(bb)
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

func make_cmd(fd *ini.File, ok, path string) (ar Cmd, err error) {
	ar.Path = path
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
		return syscall.SIGKILL, err
	}
	if val, ik := oo[ll]; ik {
		return val, nil
	}
	return syscall.SIGKILL, errors.New("can't work with 42")
}

func getint(fd *ini.File, ok, path string) (b int, err error) {
	b = 0
	ll, err := getK(fd, ok, path)
	if err != nil {
		return 0, err
	}
	b, err = strconv.Atoi(ll)
	return
}

func getbool(fd *ini.File, ik, path string) (ok bool, err error) {
	ok = false
	fg, err := getK(fd, ik, path)
	if err != nil {
		return ok, err
	}
	if fg == "true" {
		ok = true
	}
	return
}

func getauto(fd *ini.File, ok, path string) (i int, err error) {
	ff, err := getK(fd, ok, path)
	if err != nil {
		return 0, err
	}
	if ff == "unexpected" {
		return 2, nil
	}
	if ff == "ever" {
		return 1, nil
	}
	return 0, nil
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
			stop, err := get_int_array(fd, ok, "exitcodes")
			if err != nil && !NotFound(err) {
				return nil, err
			}
			numprocs, vvv, err := nprocess(fd, ok, "numprocs")
			if err != nil {
				return nil, err
			}
			sig, err := getsignal(fd, ok, "stopsignal")
			if err != nil && !NotFound(err) {
				return nil, err
			}
			stime, err := getint(fd, ok, "stoptime")
			if err != nil && !NotFound(err) {
				return nil, err
			}
			btime, err := getint(fd, ok, "starttime")
			if err != nil && !NotFound(err) {
				return nil, err
			}
			as, err := getbool(fd, ok, "autostart")
			if err != nil && !NotFound(err) {
				return nil, err
			}
			st, err := getint(fd, ok, "startretries")
			if err != nil && !NotFound(err) {
				return nil, err
			}
			aar, err := getauto(fd, ok, "autorestart")
			if err != nil && !NotFound(err) {
				return nil, err
			}
			for y := 0; y < numprocs; y++ {
				name := namer(ok, vvv, y)
				a[name] = task{
					lp:           PATH,
					lancer:       false,
					finish:       false,
					cmds:         CMD,
					umask:        UMASK,
					exitcodes:    stop,
					numprocs:     numprocs,
					stopsignal:   sig,
					stoptime:     stime,
					autostart:    as,
					autorestart:  aar,
					starttime:    btime,
					startretries: st,
					memretries:   st,
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
