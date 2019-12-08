package main

import (
	"github.com/go-ini/ini"
	"log"
)

const ()

var (
	cfg, cfgErr = ini.Load("../conf/Taskmaster.conf")
)

func getK(ar *ini.File, section, key string) (a string) {
	bb, _ := ar.Section(section).GetKey(key)
	a = bb.String()
	return
}

func get(st string) (a map[string]task, err error) {
	fd, err := ini.Load(st)
	if err != nil {
		return nil, err
	}
	a = make(map[string]task)
	ar := fd.SectionStrings()
	for _, ok := range ar {
		if ok != "DEFAULT" {
			a[ok] = task{
				com: getK(fd, ok, "com"),
			}
		}
	}
	return
}

func getKey(section, key string) (inu string) {
	ar, err := cfg.Section(section).GetKey(key)
	if err != nil {
		log.Panic(err)
	}
	return ar.String()
}
