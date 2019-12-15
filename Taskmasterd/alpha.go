package main

import (
	"sort"
)

func getWidth(a int, i int) (b string) {
	for a < i {
		b = b + " "
		a++
	}
	return
}

func padding() (i int) {
	for key, _ := range jobs {
		b := len(key)
		if b > i {
			i = b
		}
	}
	return i + 1
}

func alpha() (keys []string) {
	keys = make([]string, 0, len(jobs))
	for k := range jobs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}
