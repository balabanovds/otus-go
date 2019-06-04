package stringutil

import (
	"testing"
)

func TestOne(t *testing.T) {
	in := "a4bc2d5e"
	expected := "aaaabccddddde"

	got := Reformat(in)

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}
