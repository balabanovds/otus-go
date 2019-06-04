package stringutil

import (
	"strconv"
	"strings"
)

type node struct {
	current  rune
	previous rune
}

func (n *node) eval() string {
	if num, err := strconv.Atoi(n.current); err == nil {
		// this is a number
		if !escapedChar && i-1 >= 0 {
			for n := 0; n < num; n++ {
				result = append(result, chars[i-1])
			}
		}
	}
	return n.current
}

// Reformat string by key chars
func Reformat(s string) string {
	var nodes []node

	for i, r := range s {

	}
	return strings.Join(result, "")
}
