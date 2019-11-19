package main

import (
	"strconv"
	"taskmasterd/helper/str"
)

func namer(a string, y, b int) (c string) {
	if y == 1 {
		return a
	}
	return str.StrConcat(a, "@", strconv.Itoa(b))
}
