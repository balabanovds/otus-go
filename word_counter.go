package words

import (
	"fmt"
	"strings"
)

type word struct {
	count int
	word string 
}

type text struct {
	data map[string]int
	rating []word
}

func newText() *text {
	return &text{
		data: make(map[string]int),
		rating: []word{},
	}
}

func (t *text) rate() {
	res := []word{}
	for word, count := range t.data {
		for i, w := range t.rating {
			if w.count < count {
				res = in
			}
		}
	}
	t.rating = res
}

// Count words in straing and returns first 10 more frequent
func Count(str string) {
	chunks := strings.Split(str, " ")
	t := newText()
	for _, c := range chunks {
		t.data[c]++
	}
	fmt.Printf("%+v\n", t)
}
