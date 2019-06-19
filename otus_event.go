package event

import (
	"fmt"
	"io"
	"log"
)

// OtusEvent interface
type OtusEvent interface {
	fmt.Stringer
}

// HwSubmitted used by student to submit homework
type HwSubmitted struct {
	ID      int
	Code    string
	Comment string
}

// HwAccepted used by couch to grade homework
type HwAccepted struct {
	ID    int
	Grade int
}

// HwSubmitted implements OtusEvent interface
func (h HwSubmitted) String() string {
	return fmt.Sprintf("%s submitted %d \"%s\"", getDateNow(), h.ID, h.Comment)
}

// HwAccepted implements OtusEvent interface
func (h HwAccepted) String() string {
	return fmt.Sprintf("%s accepted %d %d", getDateNow(), h.ID, h.Grade)
}

// LogOtusEvent log events using OtusEvent interface
func LogOtusEvent(e OtusEvent, w io.Writer) {
	_, err := fmt.Fprint(w, e)
	if err != nil {
		log.Fatal(err)
	}
}
