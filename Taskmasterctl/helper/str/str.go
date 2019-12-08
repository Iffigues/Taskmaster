package str

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

func StrToIntArray(a string) (tab []int, err error) {
	s := strings.TrimSpace(a)
	space := regexp.MustCompile(`\s+`)
	w := space.ReplaceAllString(s, " ")
	b := strings.Split(w, " ")
	for _, val := range b {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		tab = append(tab, i)
	}
	return
}

func StrToStrArray(a string) (b []string) {
	s := strings.TrimSpace(a)
	space := regexp.MustCompile(`\s+`)
	w := space.ReplaceAllString(s, " ")
	bb := strings.Split(w, " ")
	for _, v := range bb {
		b = append(b, v)
	}
	return
}

func StrConcat(a ...string) (b string) {
	var buff bytes.Buffer
	for _, ok := range a {
		buff.WriteString(ok)
	}
	return buff.String()
}

func ArrayToStr(a []string) (b string) {
	var buff bytes.Buffer
	for _, ok := range a {
		buff.WriteString(ok)
		buff.WriteString(" ")
	}
	b = strings.TrimSpace(buff.String())
	return
}
