package stringutil

import (
	"testing"
)

func TestOne(t *testing.T) {
	in := "a12"
	expected := "aaaabccddddde"

	got := Reformat(in)

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}
