package words

import (
	"regexp"
	"sort"
	"strings"
)

type node struct {
	count int
	word  string
}

type sortByCounter []node

func (a sortByCounter) Len() int {
	return len(a)
}

func (a sortByCounter) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a sortByCounter) Less(i, j int) bool {
	if a[i].count == a[j].count {
		return a[i].word < a[j].word
	}
	return a[i].count > a[j].count
}

type parser struct {
	data   map[string]int
	rating sortByCounter
}

func newTextParser(str string) *parser {
	re := regexp.MustCompile(`\.|,`)
	str = re.ReplaceAllString(str, "")
	chunks := strings.Split(str, " ")
	t := &parser{
		data: make(map[string]int),
	}
	for _, c := range chunks {
		t.data[c]++
	}

	return t
}

func (t *parser) rate() {
	t.rating = []node{}

	for word, count := range t.data {
		t.rating = append(t.rating, node{count, word})
	}

	sort.Sort(t.rating)
}

// Count words in straing and returns first 10 more frequent
func Count(str string) []string {

	t := newTextParser(str)
	t.rate()

	end := 10
	if len(t.rating) < end {
		end = len(t.rating)
	}
	res := make([]string, end)
	for i, n := range t.rating[:end] {
		res[i] = n.word
	}
	return res
}
