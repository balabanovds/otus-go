package str

import (
	"strings"
	"unicode"
)

type node struct {
	value     rune
	nextNode  *node
	mult      int
	printable bool
}

func (n *node) eval() {
	if (n.printable) || (unicode.IsLetter(n.value)) {
		/*
			we want to get all following digits to compute multiplier for current node
			for example in a12bc0 -> multiplier for 'a' should be 12, for 'b' multiplier should be 1
			for 'c' miltiplier should be 0, so 'c' should be absent in result
		*/
		if (n.nextNode != nil) && (unicode.IsDigit(n.nextNode.value)) {
			next := n.nextNode
			n.mult = int(next.value - '0')
			for (next.nextNode != nil) && (unicode.IsDigit(next.nextNode.value)) {
				n.mult = n.mult*10 + int(next.nextNode.value-'0')
				next = next.nextNode
			}
		} else {
			n.mult = 1
		}
		return
	}

	if n.value == '\\' {
		n.nextNode.printable = true
		return
	}
	if unicode.IsDigit(n.value) {
		return
	}
	n.printable = true
	n.eval()

}

func (n *node) string() string {
	var b strings.Builder
	for i := 0; i < n.mult; i++ {
		b.WriteRune(n.value)
	}
	return b.String()
}

// Encode string by key chars
func Encode(s string) string {
	nodes := make([]*node, 0)

	r := []rune(s)

	// we run through []rune in backward order to get next reference for each node
	for i, j := len(r)-1, 0; i >= 0; i, j = i-1, j+1 {
		n := &node{
			value: r[i],
		}
		if i < len(r)-1 {
			n.nextNode = nodes[j-1]
		}
		nodes = append(nodes, n)
	}

	// as we got []node in backward order so we reverse it just for convinience
	for left, right := 0, len(nodes)-1; left < right; left, right = left+1, right-1 {
		nodes[left], nodes[right] = nodes[right], nodes[left]
	}

	var strb strings.Builder

	// we evaluate each node
	for _, n := range nodes {
		n.eval()
	}

	// we write each node to strings buffer
	for _, n := range nodes {
		strb.WriteString(n.string())
	}

	return strb.String()
}
