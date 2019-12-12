package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"syscall"
	"taskmasterd/helper/str"
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
	status = false
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
		}
	}
	return
}

func rerun(a string) (retrie, abort int) {
	if val, ok := jobs[a]; ok {
		return val.startretries, 50
	}
	return 0, 50
}

func delque(a string) {
	if _, ok := queued[a]; ok {
		delete(queued, a)
	}
}

func aborting(cc *task, abort int, a string) (ok bool) {
	if abort <= 0 {
		cc.finish, cc.abort = true, true
		registre(a, "programme abort at: "+time.Now().String())
		return true
	}
	return false
}

func lance(c chan bool, a ...string) {
	retrie, abort := rerun(a[0])
	var f bool
label:
	abort = abort - 1
	ok := start_command(a[0])
	if ok {
		go func() {
			c <- true
		}()
		ii, cc := false, queued[a[0]]
		if ok := aborting(cc, abort, a[0]); ok {
			return
		}
		cc.finish, cc.start, cc.lancer = false, time.Now(), true
		_, done := cc.cmdl.Start(), make(chan error, 1)
		registre(a[0], "progam start at: "+cc.start.String())
		go func() {
			done <- cc.cmdl.Wait()
		}()
		select {
		case errs := <-done:
			go func() {
				cc.verif <- true
			}()
			err, fifi := errs, time.Since(cc.start)
			cc.lancer, cc.finish, cc.end, cc.exectime, ii, cc.nbexec = finish(fifi.Seconds(), cc.starttime, cc.nbexec)
			cc.succed, cc.status = isgood(err, cc.exitcodes, ii)
			f, retrie = is_false(cc, retrie, a[0], cc.status)
			if f {
				goto label
			}
			registre(a[0], "programme finish at:"+cc.end.String()+" during: "+fmt.Sprintf("%f", cc.exectime)+"begin at :"+cc.start.String(), 1, 2)
		}
	} else {
		c <- false
	}
}

func is_false(cc *task, retrie int, a string, rrr bool) (vrai bool, i int) {
	if cc.stop {
		registre(a, "programme stop at:"+cc.end.String())
		return false, retrie
	}
	if !cc.succed || !cc.status {
		if retrie > 0 && !cc.succed {
			registre(a, "programme retrie process at: "+cc.end.String())
			return true, retrie - 1
		} else if cc.autorestart == 2 && !rrr {
			return true, retrie
		} else {
			registre(a, "programme fail at: "+cc.end.String())
			return false, retrie
		}
	} else {
		retrie = cc.startretries
		if cc.autorestart == 1 {
			registre(a, "programme restart at:"+time.Now().String())
			return true, retrie
		}
		return false, retrie
	}
}

func send_signal(conn net.Conn, a ...string) (ce ret, err error) {
	oo := map[string]syscall.Signal{
		"TERM": syscall.SIGTERM,
		"HUP":  syscall.SIGHUP,
		"INT":  syscall.SIGINT,
		"KILL": syscall.SIGKILL,
		"USR1": syscall.SIGUSR1,
		"USR2": syscall.SIGUSR2,
	}
	if len(a) > 1 {
		if val, ok := oo[a[0]]; ok {
			eee, bbb := is_started(a[1])
			if !eee {
				conn.Write([]byte("process not enqueued"))
				return
			}
			if bbb {
				conn.Write([]byte("process is finish"))
				return
			}
			err := queued[a[1]].cmdl.Process.Signal(val)
			if err != nil {
				conn.Write([]byte(err.Error()))
			} else {
				conn.Write([]byte("Signal envoyer"))
			}
		} else {
			conn.Write([]byte("it's not a good signal"))
		}
	} else {
		conn.Write([]byte("pls, specifie a process name"))
	}
	return
}

func kill(conn net.Conn, a ...string) (ce ret, err error) {
	if len(a) > 0 && a[0] == "all" {
		for _, val := range queued {
			if val.lancer {
				val.cmdl.Process.Kill()
			}
		}
	} else {
		for _, val := range a {
			if kk, ok := queued[val]; ok {
				if kk.lancer {
					queued[val].cmdl.Process.Kill()
				}
			}
		}
	}
	conn.Write([]byte("commande lancer"))
	return
}

func meme(c chan bool, a, b string) (strs string) {
	e := <-c
	if e {
		return a
	} else {
		return b
	}
}

func mami(c bool, a, b string) (strs string) {
	if c {
		return a
	} else {
		return b
	}
}

func veve(a string) {
	if jobs[a].autostart {
		c := make(chan bool)
		go lance(c, a)
	}
}

func start(conn net.Conn, a ...string) (ce ret, err error) {
	strs := ""
	if len(a) > 0 {
		c := make(chan bool)
		if a[0] == "all" {
			for key, _ := range jobs {
				go lance(c, key)
				e := meme(c, "jobs started\n", "jobs not found\n")
				strs = str.StrConcat(strs, e)
			}
		} else {
			for _, key := range a {
				go lance(c, key)
				e := meme(c, "jobs started\n", "jobs not found\n")
				strs = str.StrConcat(strs, e)
			}
		}
	} else {
		conn.Write([]byte("bad init file"))
	}
	if strs == "" {
		conn.Write([]byte(" "))
	}
	conn.Write([]byte(strs))
	return
}

func stop(conn net.Conn, a ...string) (c ret, err error) {
	strs := ""
	if len(a) > 0 {
		if a[0] == "all" {
			for key, _ := range jobs {
				ok, g := stop_command(key)
				oui, d := is_stopped(ok, g)
				e := mami(oui, d, "jobs don't stop\n")
				strs = str.StrConcat(strs, e)
			}
		} else {
			for _, key := range a {
				ok, g := stop_command(key)
				oui, d := is_stopped(ok, g)
				e := mami(oui, d, "jobs don't stop\n")
				strs = str.StrConcat(strs, e)
			}
		}
	}
	if strs == "" {
		conn.Write([]byte(" "))
	}
	conn.Write([]byte(strs))
	return
}

func restart(conn net.Conn, a ...string) (c ret, err error) {
	strs := ""
	if len(a) > 0 {
		c := make(chan bool)
		if a[0] == "all" {
			for key, _ := range jobs {
				stop_command(key)
				time.Sleep(1 * time.Second)
				go lance(c, key)
				e := meme(c, "started command\n", "jobs not found\n")
				strs = str.StrConcat(strs, e)
			}
		} else {
			for _, key := range a {
				stop_command(key)
				time.Sleep(1 * time.Second)
				go lance(c, key)
				e := meme(c, "started command\n", "jobs not found\n")
				strs = str.StrConcat(strs, e)
			}
		}
	}
	if strs == "" {
		conn.Write([]byte(" "))
		return
	}
	conn.Write([]byte(strs))
	return
}

func reload(conn net.Conn, a ...string) (c ret, err error) {
	syscall.Kill(mypid, syscall.SIGHUP)
	registre("reload", "taskmaster reload at:"+time.Now().String())
	conn.Write([]byte("reload"))
	return
}

func exit(conn net.Conn, a ...string) (c ret, err error) {
	c.end = true
	return
}

func begin() (err error) {
	fmt.Println(jobs["autostart"].autostart)
	for key, val := range jobs {
		if val.autostart {
			println("eee")
			c := make(chan bool)
			go lance(c, key)
		}
	}
	return
}

func status(conn net.Conn, a ...string) (c ret, err error) {
	var t string
	for u, _ := range jobs {
		y := str.StrConcat(u, ":  ")
		if i, ok := queued[u]; ok {
			if u == "abort" {
				fmt.Println(i.abort)
			}
			if i.abort {
				y = str.StrConcat(y, "abort")
				goto lab
			}
			if i.stop {
				y = str.StrConcat(y, " stop")
				goto lab
			}
			if i.finish {
				y = str.StrConcat(y, "finish")
				goto lab
			}
			if i.lancer {
				y = str.StrConcat(y, "start")
				goto lab
			}
		}
		y = str.StrConcat(y, "not started")
	lab:
		t = str.StrConcat(t, y, "\n")
	}
	if t == "" {
		conn.Write([]byte(" "))
		return
	}
	conn.Write([]byte(t))
	return
}

func consoles(conn net.Conn, a ...string) (c ret, err error) {
	if len(a) > 0 && a[0] != "" {
		if e, d := console[a[0]]; d {
			if len(a) > 1 {
				return e(conn, a[1:len(a)-1]...)
			} else {
				return e(conn)
			}
		} else {
			_, err = conn.Write([]byte("bad command\n"))
			return
		}
	}
	_, err = conn.Write([]byte(""))
	return
}
