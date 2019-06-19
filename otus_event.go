package event

import (
	"fmt"
	"io"
	"log"
	"time"
)

type OtusEvent interface {
	fmt.Stringer
}

type HwSubmitted struct {
	ID      int
	Code    string
	Comment string
}

type HwAccepted struct {
	ID    int
	Grade int
}

func (h HwSubmitted) String() string {
	return fmt.Sprintf("%s submitted %d \"%s\"", getDateNow(), h.ID, h.Comment)
}

func (h HwAccepted) String() string {
	return fmt.Sprintf("%s accepted %d %d", getDateNow(), h.ID, h.Grade)
}

// LogOtusEvent logs events
func LogOtusEvent(e OtusEvent, w io.Writer) {
	_, err := fmt.Fprint(w, e)
	if err != nil {
		log.Fatal(err)
	}
}

func getDateNow() string {
	t := time.Now()
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}
