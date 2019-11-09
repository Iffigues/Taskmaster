package main

import (
//	"github.com/sevlyar/go-daemon"
//	"log"
	"fmt"
	"syscall"
)

var (
	mypid = syscall.Getpid()
)

func main() {
/*		cntxt := &daemon.Context{
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
	_, err = get("../ini/ini.ini")
	if err != nil {
		log.Fatal(err)
	}*/
	fmt.Println("yttyyt")
	serve()
}
