package main

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
	"log"
	"os"
	"syscall"
)

var (
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
			PidFileName: "../log/taskmaster.pid",
			PidFilePerm: 0777,
			LogFileName: "../log/sample.log",
			LogFilePerm: 0777,
			WorkDir:     "./",
			Umask:       027,
			Args:        []string{"l"},
		}
		d, err := cntxt.Reborn()
		if err != nil {
			log.Fatal("Unable to run: ", err)
		}
		if d != nil {
			fmt.Println(d)
			return
		}
		defer cntxt.Release()
		log.Print("- - - - - - - - - - - - - - -")
		log.Print("daemon started")
	}
	serve()
}
