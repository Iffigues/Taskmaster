package main

import (
	"github.com/sevlyar/go-daemon"
	"log"
	"syscall"
)

var (
	mypid  = syscall.Getpid()
)

func main() {
	cntxt := &daemon.Context{
		PidFileName: "taskmaster.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
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
	serve()
}
