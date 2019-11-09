package str

import (
	"regexp"
	"strings"
	"strconv"
)

func StrToIntArray(ar string) (tab []int, err error) {
	s := strings.TrimSpace(ar)
	space := regexp.MustCompile(`\s+`)
	w := space.ReplaceAllString(s, " ")
	b := strings.Split(w, " ")
	for _, val := range b {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		tab = append(tab, i);
	}
	return
}
