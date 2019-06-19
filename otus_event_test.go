package event

import (
	"bytes"
	"testing"
)

func TestLogOtusEventSubmit(t *testing.T) {
	var b bytes.Buffer
	var h OtusEvent = HwSubmitted{
		ID:      3456,
		Comment: "please take a look at my homework",
	}
	LogOtusEvent(h, &b)
	got := b.String()
	want := getDateNow() + ` submitted 3456 "please take a look at my homework"`

	if got != want {
		t.Errorf("\nwant: >%v<\ngot: >%v<", want, got)
	}
}

func TestLogOtusEventAccept(t *testing.T) {
	var b bytes.Buffer
	var h OtusEvent = HwAccepted{
		ID:    3456,
		Grade: 4,
	}
	LogOtusEvent(h, &b)
	got := b.String()
	want := getDateNow() + ` accepted 3456 4`

	if got != want {
		t.Errorf("\nwant: >%v<\ngot: >%v<", want, got)
	}
}
