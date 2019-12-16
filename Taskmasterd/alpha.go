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

func getPadding(a []string) (i int) {
	for _, val := range a {
		b := len(val)
		if b > i {
			i = b
		}
	}
	return i + 1
}

func getAlpha(a ...string) (keys []string) {
	keys = make([]string, 0, len(a))
	for _, k := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}

func alpha() (keys []string) {
	keys = make([]string, 0, len(jobs))
	for k, _ := range jobs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}
