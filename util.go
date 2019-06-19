package event

import (
	"fmt"
	"time"
)

// this is just util function to shorten main code, and handy in tests
func getDateNow() string {
	t := time.Now()
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}
