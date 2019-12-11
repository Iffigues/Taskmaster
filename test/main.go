package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func perror(msg string) {
	fmt.Fprintln(os.Stderr, "%s", msg)
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func tab(a string) (ff []os.Signal, c []string) {
	oo := map[string]os.Signal{
		"TERM": syscall.SIGTERM,
		"HUP":  syscall.SIGHUP,
		"INT":  syscall.SIGINT,
		"KILL": syscall.SIGKILL,
		"USR1": syscall.SIGUSR1,
		"USR2": syscall.SIGUSR2,
	}
	c = strings.Split(a, " ")
	for _, h := range c {
		if p, ok := oo[h]; ok {
			ff = append(ff, os.Signal(p))
		}
	}
	return
}

func main() {
	for _, pair := range os.Environ() {
    fmt.Println(pair)
  }
	dir, err := os.Getwd()
	fmt.Println(dir, err)
	i := os.Args
	bbb := false
	exit := 0
	if len(i) == 1 {
		return
	}
	var g []os.Signal
	var haha []string
	t := false
	b := i[1:]
	if b[0] == "loop" {
		t = true
		if len(b) > 1 {
			b = b[1:]
		}
	}
	if b[0] == "signal" {
		if len(b) > 1 {
			g, haha = tab(b[1])
			b = b[1:]
			fmt.Println(g, haha)
			t = true
			if len(b) > 1 {
				b = b[1:]
			} else {
				b = nil
			}
		}
	}
	if b != nil {
		if b[0] == "time" {
			if len(b) == 1 {
				return
			}
			if tt, err := strconv.Atoi(b[1]); err == nil {
				println(tt)
				time.Sleep(time.Duration(tt) * time.Second)
				fmt.Println(len(b[1:]))
				if len(b[1:]) > 1 {
					fmt.Println("hghg")
					b = b[2:]
				}
			} else {
				return
			}
		}
		c := len(b)
		if c > 0 {
			s1 := rand.NewSource(time.Now().UnixNano())
			r2 := rand.New(s1)
			r3 := rand.New(s1)
			r4 := rand.New(s1)
			if b[0] == "stdout" {
				if len(b) > 1 && b[1] == "loop" {
					go func() {
						for {
							fmt.Println(String(r3.Intn(1000)))
						}
					}()
					b = b[1:]
				} else {
					tt := r3.Intn(1000)
					for ; tt > 0; tt-- {
						fmt.Println(String(r3.Intn(1000)))
					}
					if len(b) > 1 {
						b = b[1:]
					}
				}
			}
			if b[0] == "stderr" {
				if len(b) > 1 && b[1] == "loop" {
					go func() {
						for {
							perror(String(r3.Intn(1000)))
						}
					}()
					b = b[1:]
				} else {
					tt := r4.Intn(1000)
					for ; tt > 0; tt-- {
						perror(String(r3.Intn(1000)))
					}
					if len(b) > 1 {
						b = b[1:]
					}
				}
			}
			gg := r2.Intn(len(b))
			if tt, err := strconv.Atoi(b[gg]); err == nil {
				bbb = true
				exit = tt
			}
		}
	}
	if t {
		go func() {
			for {
			}
		}()
		fanny(haha, exit, g...)
	}
	if bbb {
		os.Exit(exit)
	}
	return
}
