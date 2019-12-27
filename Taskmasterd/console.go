package main

import (
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	console = map[string]func(net.Conn, ...string) (ret, error){
		"exit":    exit,
		"reload":  reload,
		"start":   start,
		"stop":    stop,
		"status":  status,
		"restart": restart,
		"kill":    kill,
		"signal":  send_signal,
		"pid":     pid,
	}
	queued = make(enqued)
)

func isgood(a error, b []int, i bool) (ok, status bool) {
	ok = true
	status = true
	if !i {
		ok = false
	}
	t := 0
	if a != nil {
		vv := strings.Split(a.Error(), " ")
		if len(vv) == 3 {
			t, _ = strconv.Atoi(vv[2])
		} else {
			return ok, false
		}
	}
	for _, ff := range b {
		if ff == t {
			status = true
			break
		} else {
			status = false
		}
	}
	return
}

func rerun(a string) (tt consolle) {
	if val, ok := jobs[a]; ok {
		return consolle{val.startretries, 50, false}
	}
	return consolle{0, 50, false}
}

func aborting(cc *task, abort int, a string) (ok bool) {
	if abort <= 0 {
		cc.finish, cc.abort = true, true
		registre(a, "programme abort at: "+time.Now().String())
		return true
	}
	return false
}

func add_bool(c chan bool, ok bool) {
	c <- ok
}

func lance(c chan bool, a ...string) {
	cons := rerun(a[0])
	first := true
	for {
		cons.abort = cons.abort - 1
		ok := start_command(a[0])
		if ok {
			if first {
				go add_bool(c, true)
				first = false
			}
			ii, cc := false, queued[a[0]]
			if ok := aborting(cc, cons.abort, a[0]); ok {
				return
			}
			cc.finish, cc.start, cc.lancer = false, time.Now(), true
			err, done := cc.cmdl.Start(), make(chan error, 1)
			if err != nil {
				cc.finish = true
				return
			}
			registre(a[0], "progam start at: "+cc.start.String())
			func() {
				done <- cc.cmdl.Wait()
			}()
			select {
			case errs := <-done:
				if cc.stop {
					cc.verif <- true
				}
				cc.Stdout.Close()
				cc.Stderr.Close()
				cc.lancer = false
				cons.f, cons.retrie = wait_finish(cc, errs, ii, cons.retrie, a[0])
				if !cons.f {
					print("ghghghg")
					return
				}
			}
		} else {
			if first {
				go add_bool(c, false)
			}
			return
		}
	}
}
