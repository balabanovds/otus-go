package words

import (
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

// Count words in straing and returns first 10 more frequent
func Count(str string) map[int]string {
	chunks := strings.Split(str, " ")
	
}
