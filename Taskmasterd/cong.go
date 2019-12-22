package main

import (
	"fmt"
	"net"
	"syscall"
	"time"
)

func wait_finish(cc *task, errs error, ii bool, retrie int, a string) (vrai bool, i int) {
	err, fifi := errs, time.Since(cc.start)
	cc.lancer, cc.finish, cc.end, cc.exectime, ii, cc.nbexec = finish(fifi.Seconds(), cc.starttime, cc.nbexec)
	cc.succed, cc.status = isgood(err, cc.exitcodes, ii)
	return is_false(cc, retrie, a, cc.status)

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
		} else if cc.autorestart == 1 {
			return true, retrie
		} else {
			mut.Lock()
			cc.failed = true
			mut.Unlock()
			registre(a, "programme fail at: "+cc.end.String())
			return false, retrie
		}
	} else {
		retrie = cc.startretries
		if cc.autorestart == 1 {
			registre(a, "programme restart at:"+time.Now().String())
			return true, retrie
		}
		registre(a, "programme finish at:"+cc.end.String()+" during: "+fmt.Sprintf("%f", cc.exectime)+"begin at :"+cc.start.String(), 1, 2)
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
	strs := " "
	if len(a) > 0 && a[0] == "all" {
		for key, val := range queued {
			if val.lancer {
				err := val.cmdl.Process.Kill()
				if err != nil {
					strs = strs + key + ": " + err.Error() + "\n"
				} else {
					strs = strs + key + ": " + "process kill\n"
				}
			}
		}
	} else {
		for _, val := range a {
			if kk, ok := queued[val]; ok {
				if kk.lancer {
					err := queued[val].cmdl.Process.Kill()
					if err != nil {
						strs = strs + val + ": " + err.Error() + "\n"
					} else {
						strs = strs + val + ": " + "process kill\n"
					}
				}
			}
		}
	}
	conn.Write([]byte(strs))
	return
}

func meme(c chan bool, prog, a, b, y string) (strs string) {
	e := <-c
	if e {
		return a
	} else {
		if _, ok := jobs[prog]; ok {
			if hh, _ := is_started(prog); !hh {
				return "can't start\n"
			}
			return y
		}
		return b
	}
}
