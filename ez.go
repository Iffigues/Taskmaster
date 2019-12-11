package main

import (
	"fmt"
	"os"
	"time"
)

func change(){
	os.Chdir("..")
	dir, err := os.Getwd()
	fmt.Println(dir, err)
}

func main() {
	go change()
	dir, err := os.Getwd()
	fmt.Println(dir, err)
	time.Sleep(20*time.Second)
	fmt.Println(dir, err)
}
