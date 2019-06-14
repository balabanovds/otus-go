package hw04

import (
	"testing"
)

func TestFindMax(t *testing.T) {
	in := []string{"one", "two", "three"}
	predicate := func(i, j string) bool {
		return len(i) > len(j)
	}
	expected := "three"

	got := FindMax(slice, predicate)

	if got != expected {
		t.Errorf("got %v, expected: %v", got, expected)
	}

}
