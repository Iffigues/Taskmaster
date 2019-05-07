package main

import (
	"github.com/go-ini/ini"
	"log"
)

const (
	bdd = "bdd"
	htp = "htp"
)

var (
	cfg, cfgErr = ini.Load("ini.ini")
)

func init() {
}

func get() (a map[string]task) {
	a = make(map[string]task)
	ar := cfg.SectionStrings()
	for _, ok := range ar {
		if ok != "DEFAULT" {
			a[ok] = task{}
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
