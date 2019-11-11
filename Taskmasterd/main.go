package main

import (
	"github.com/sevlyar/go-daemon"
	"fmt"
	"os"
	"log"
	"syscall"
)

var (
	mode  = false
	mypid = syscall.Getpid()
	jobs, errorJobs = get("../ini/ini.ini")
)



func init(){
	b := os.Args
	if len(b) > 1 {

	}
}

func main() {
		if errorJobs != nil {
			return
		}
		if mode {
			cntxt := &daemon.Context{
				PidFileName: "../log/taskmaster.pid",
				PidFilePerm: 0644,
				LogFileName: "../log/sample.log",
				LogFilePerm: 0640,
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
