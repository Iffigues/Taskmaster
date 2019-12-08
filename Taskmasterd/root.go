package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func initee() {
	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	i, err := strconv.Atoi(string(output[:len(output)-1]))
	if err != nil {
		log.Fatal(err)
	}
	if i != 0 {
		os.Exit(1)
	}
}
