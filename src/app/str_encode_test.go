package main

import (
	"str"
	"testing"
)

type data struct {
	incoming string
	expected string
}

func TestEncode(t *testing.T) {
	datas := []data{
		{incoming: "a4bc2d5e", expected: "aaaabccddddde"},
		{incoming: "a12", expected: "aaaaaaaaaaaa"},
		{incoming: "abcd", expected: "abcd"},
		{incoming: "45", expected: ""},
		{incoming: `qwe\4\5`, expected: "qwe45"},
		{incoming: `qwe\45`, expected: "qwe44444"},
		{incoming: `qwe\\5`, expected: `qwe\\\\\`},
		{incoming: `qw0e`, expected: `qe`},
		{incoming: `<5`, expected: `<<<<<`},
		{incoming: `☀5`, expected: `☀☀☀☀☀`},
	}

	for _, d := range datas {
		got := str.Encode(d.incoming)

		if d.expected != got {
			t.Errorf("Expected: %s, got: %s", d.expected, got)
		}
	}
}
