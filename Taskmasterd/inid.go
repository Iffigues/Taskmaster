package main

import (
	"github.com/go-ini/ini"
	"strconv"
)

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
			UMASK, err := getumask(fd, ok)
			if err != nil && !NotFound(err) {
				return nil, err
			}
			CMD, PATH, err := make_cmd(fd, ok, UMASK)
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
			//if stime == 0 {
			//	stime = 1
			//}
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
			rf, err := getrun(fd, ok)
			if err != nil && !NotFound(err) {
				return nil, err
			}
			gr, err := getgroup(fd, ok, "group")
			if err != nil && !NotFound(err) {
				return nil, err
			}
			for y := 0; y < numprocs; y++ {
				name := namer(ok, vvv, y)
				a[name] = task{
					name:         ok,
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
					group:        gr,
					grap:         rf,
				}
			}
		}
	}
	err = nil
	return
}
