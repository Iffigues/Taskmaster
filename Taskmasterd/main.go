package main

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
	"os"
	"sync"
	"syscall"
)

var (
	mut             = sync.RWMutex{}
	mode            = false
	mypid           = syscall.Getpid()
	jobs, errorJobs = get("./ini/ini.ini")
)

func init() {
	if errorJobs != nil {
		fmt.Println(errorJobs)
		os.Exit(0)
	}
	begin()
	b := os.Args
	if len(b) == 3 {
		if b[1] == "mod" && b[2] == "daemon" {
			mode = true
		}
	}
}

func main() {
	if mode {
		cntxt := &daemon.Context{
			PidFileName: "./log/taskmaster.pid",
			PidFilePerm: 0777,
			LogFileName: "./log/sample.log",
			LogFilePerm: 0777,
			WorkDir:     "./",
			Umask:       022,
			Args:        []string{"l"},
		}
		d, err := cntxt.Reborn()
		if err != nil {
		}
		if d != nil {
			return
		}
		defer cntxt.Release()
	}
	serve()
}
