package main

import (
	"fmt"
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
	get()
}

func get() {
	ar := cfg.SectionStrings()
	for _, ok := range ar {
		if ok != "DEFAULT" {
			fmt.Println(ok)
		}
	}
}

func getKey(section, key string) (inu string) {
	ar, err := cfg.Section(section).GetKey(key)
	if err != nil {
		log.Panic(err)
	}
	return ar.String()
}
