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
		}
	}
	for _, ff := range b {
		if ff == t {
			status = true
		}
	}
	return
}

func lance(c chan bool, a ...string) {
label:
	var err error
	ok := start_command(a[0])
	if ok {
		c <- true
		ii := false
		cc := queued[a[0]]
		cc.finish = false
		cc.lancer = false
		cc.cmdl.Start()
		cc.start = time.Now()
		registre(a[0], "progam start at: "+cc.start.String())
		cc.lancer = true
		done := make(chan error, 1)
		go func() {
			done <- cc.cmdl.Wait()
		}()
		select {
		case errs := <-done:
			go func() {
				cc.verif <- true
			}()
			fifi := time.Since(cc.start)
			cc.exectime = fifi.Seconds()
			cc.nbexec = cc.nbexec + 1
			cc.lancer = false
			cc.end = time.Now()
			ii = fifi.Seconds() >= float64(time.Second*time.Duration(cc.starttime))
			err = errs
			cc.finish = true
			cccc, rrr := isgood(err, cc.exitcodes, ii)
			cc.succed = cccc
			if cc.stop {
				registre(a[0], "programme stop at:"+cc.end.String())
			} else {
				if !cccc || !rrr {
					if cc.startretries > 0 && !cccc {
						registre(a[0], "programme retrie process at: "+cc.end.String())
						goto label
					} else if cc.autorestart == 2 && !rrr {
						goto label
					} else {
						registre(a[0], "programme fail at: "+cc.end.String())
					}
				} else {
					cc.startretries = cc.memretries
					if cc.autorestart == 1 {
						registre(a[0], "programme restart at:"+time.Now().String())
						goto label
					}
				}
			}
			registre(a[0], "programme finish at:"+cc.end.String()+" during: "+fmt.Sprintf("%f", cc.exectime)+"begin at :"+cc.start.String(), 1, 2)
		}
	} else {
		c <- false
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
	if a[0] == "all" {
		for _, val := range queued {
			if val.lancer {
				val.cmdl.Process.Kill()
			}
		}
		return
	}
	for _, val := range a {
		if kk, ok := queued[val]; ok {
			if kk.lancer {
				queued[val].cmdl.Process.Kill()
			}
		}
	}
	conn.Write([]byte("commande lancer"))
	return
}

func meme(c chan bool, a, b string, conn net.Conn) {
	if <-c {
		conn.Write([]byte(a))
	} else {
		conn.Write([]byte(b))
	}
}

func mami(c bool, a, b string, conn net.Conn) {
	if c {
		conn.Write([]byte(a))
	} else {
		conn.Write([]byte(b))
	}
}

func veve(a string) {
	if jobs[a].autostart {
		c := make(chan bool)
		go lance(c, a)
	}
}

func start(conn net.Conn, a ...string) (ce ret, err error) {
	if len(a) > 0 {
		c := make(chan bool)
		if a[0] == "all" {
			for key, _ := range jobs {
				go lance(c, key)
				meme(c, "jobs started", "jobs not found", conn)
			}
			return
		} else {
			for _, key := range a {
				go lance(c, key)
				meme(c, "jobs started", "jobs not found", conn)
			}
		}
	} else {
		conn.Write([]byte("bad init file"))
	}
	return
}

func stop(conn net.Conn, a ...string) (c ret, err error) {
	if len(a) > 0 {
		if a[0] == "all" {
			for key, _ := range jobs {
				mami(stop_command(key), "jobs stoped", "jobs don't stop", conn)
			}
		} else {
			for _, key := range a {
				mami(stop_command(key), "jobs stoped", "jobs don't stop", conn)
			}
		}
	}
	return
}

func restart(conn net.Conn, a ...string) (c ret, err error) {
	if len(a) > 0 {
		c := make(chan bool)
		if a[0] == "all" {
			for key, _ := range jobs {
				stop_command(key)
				go lance(c, key)
				meme(c, "started command", "jobs not found", conn)
			}
		} else {
			for _, key := range a {
				stop_command(key)
				go lance(c, key)
				meme(c, "started command", "jobs not found", conn)
			}
		}
	}
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
	for key, val := range jobs {
		if val.autostart {
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
