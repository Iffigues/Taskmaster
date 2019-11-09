package main

import (
	//"github.com/sevlyar/go-daemon"
	"fmt"
	"log"
	"syscall"
)

var (
	mypid = syscall.Getpid()
)

func main() {
	/*
		cntxt := &daemon.Context{
			PidFileName: "../log/taskmaster.pid",
			PidFilePerm: 0644,
			LogFileName: "../log/sample.log",
			LogFilePerm: 0640,
			WorkDir:     "/",
			Umask:       027,
			Args:        []string{"--server"},
		}
		d, err := cntxt.Reborn()
		if err != nil {
			log.Fatal("Unable to run: ", err)
		}
		if d != nil {
			return
		}
		defer cntxt.Release()
		log.Print("- - - - - - - - - - - - - - -")
		log.Print("daemon started")
	*/
	jobs, err := get("../ini/ini.ini")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jobs["yes"].cmds)
	serve()
}
