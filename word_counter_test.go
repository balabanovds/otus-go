package words

import (
	"testing"
)

func TestCounter(t *testing.T) {
	text := `a a a b b b b b b c c c c`
	Count(text)
}