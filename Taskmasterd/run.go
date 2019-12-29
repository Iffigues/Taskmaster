package main

import (
	"crypto/rand"
	"fmt"
)

func tok(size int) string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func endrun(cc task, ok bool) {
	c := make(chan bool, 1)
	if !ok {
		for _, val := range cc.grap.runatfailed {
			go lance(c, val, cc.name+":"+val+":"+tok(8))
			<-c
		}
	}
	if ok {
		for _, val := range cc.grap.runatsucced {
			go lance(c, val, cc.name+":"+val+":"+tok(8))
			<-c
		}
	}
	for _, val := range cc.grap.runwhatever {
		go lance(c, val, cc.name+":"+val+":"+tok(8))
		<-c
	}
}
