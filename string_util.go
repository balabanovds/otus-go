package stringutil

import (
	"strings"
	"unicode"
)

type node struct {
	value rune
	next  *node
	prev  *node
	mult  int32
	hidden   bool
}

func (n *node) eval() string {
	// check
	if (n.prev.value != '\\') &&
		(unicode.IsDigit(n.value)) {
		n.mult = n.value
		for (unicode.IsDigit(n.next.value)) ||
			(n.next != nil) {
				if !n.next.hidden {
					n.mult = n.mult * 10 + n.next.value
				}
			n.next.hidden = true
		}
	}
	// if previous value is '\' so return current no matter what it is
	// or if current value is letter
	if (n.prev.value == '\\') ||
		(unicode.IsLetter(n.value)) {
		return string(n.value)
	}
	if unicode.IsDigit(n.value) {
		// n.previous.multiplier =
	}
	return ""
}

// Reformat string by key chars
func Reformat(s string) string {
	var nodes []*node

	for i, r := range s {
		n := &node{
			value: r,
		}
		if i > 0 {
			n.prev = nodes[i-1]
		}
		nodes = append(nodes, n)
	}

	var strb strings.Builder

	// create next branch
	for i, n := range nodes {
		if i < len(nodes)-1 {
			n.next = nodes[i+1]
		}
	}

	return strb.String()
}
