package parallel

import (
	"errors"
	"fmt"
	"testing"
)

func TestRunParallel(t *testing.T) {
	s := []func() error{
		func() error {
			fmt.Println("First func")
			return nil
		},
		func() error {
			fmt.Println("Second func")
			return errors.New("2nd")
		},
		func() error {
			fmt.Println("Third func")
			return errors.New("3rd")
		},
	}

	RunParallel(s, 2, 1)
}
