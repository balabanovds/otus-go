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

func TestTwo(t *testing.T) {
	in := "a12"
	expected := "aaaaaaaaaaaa"

	got := Reformat(in)

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func TestThree(t *testing.T) {
	in := "abcd"
	expected := "abcd"

	got := Reformat(in)

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func TestFour(t *testing.T) {
	in := "45"
	expected := ""

	got := Reformat(in)

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func TestFive(t *testing.T) {
	in := `qwe\4\5`
	expected := "qwe45"

	got := Reformat(in)

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func TestSix(t *testing.T) {
	in := `qwe\45`
	expected := "qwe44444"

	got := Reformat(in)

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func TestSeven(t *testing.T) {
	in := `qwe\\5`
	expected := `qwe\\\\\`

	got := Reformat(in)

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func TestEight(t *testing.T) {
	in := `qw0e`
	expected := `qe`

	got := Reformat(in)

	if expected != got {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}
