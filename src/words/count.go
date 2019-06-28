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

func newParser(str string) *parser {
	// remove all non word , digit or space
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]`)
	str = re.ReplaceAllString(str, "")

	// remove all double spaces
	space := regexp.MustCompile(`\s+`)
	str = space.ReplaceAllString(str, " ")

	chunks := strings.Split(str, " ")

	t := &parser{
		data: make(map[string]int),
	}
	for _, c := range chunks {
		t.data[strings.ToLower(c)]++
	}

	return t
}

func (t *parser) rate() {
	t.rating = make([]node, 0, len(t.data))

	for word, count := range t.data {
		t.rating = append(t.rating, node{count, word})
	}

	sort.Sort(t.rating)
}

// Count words in string and returns first 10 more frequent
func Count(str string) []string {

	t := newParser(str)
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
