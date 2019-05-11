package main


// #cgo CFLAGS: -DPNG_DEBUG=1
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -lpng
// #include <png.h>
//	#cgo CFLAGS: -g -Wall
//	#include <readline/readline.h>
//	#include <readline/history.h>


import (
)

import "C"

var (
)

func init() {
	C.using_history()
}
