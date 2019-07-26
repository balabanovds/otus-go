package pool

import (
	"errors"
	"sync"
	"testing"
)

var m sync.Mutex

var cntr int

var maxWorkers = 3
var maxErrors = 3

func fnOk() error {
	m.Lock()
	defer m.Unlock()
	cntr++
	return nil
}

func fnErr() error {
	m.Lock()
	defer m.Unlock()
	cntr++
	return errors.New("")
}

type test struct {
	fns          []fn
	expectedDone int
	msg          string
}

var tests = []test{
	{
		fns: []fn{
			fnOk,
			fnOk,
			fnOk,
			fnErr,
			fnErr,
		},
		expectedDone: 5,
		msg:          "Test all funcs executed",
	},
	{
		fns: []fn{
			fnOk,
			fnOk,
			fnErr,
			fnErr,
			fnErr,
			fnErr,
			fnErr,
			fnErr,
		},
		expectedDone: maxErrors + 2,
		msg:          "Test max errors",
	},
}

func TestExec(t *testing.T) {
	for _, p := range tests {
		cntr = 0
		Exec(p.fns, maxWorkers, maxErrors)
		if p.expectedDone != cntr {
			t.Errorf("%s want %d got %d", p.msg, p.expectedDone, cntr)
		}
	}
}
