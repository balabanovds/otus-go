package pool

import (
	"fmt"
	"sync"
	"time"
)

type logg struct {
	sync.Mutex
	data []string
}

func newLog() logg {
	return logg{data: []string{}}
}

func (l *logg) add(s string) {
	l.Lock()
	l.data = append(l.data, time.Now().String()+"   "+s)
	l.Unlock()
}
func (l *logg) addf(format string, a ...interface{}) {
	sf := fmt.Sprintf(format, a...)
	l.add(sf)
}

func (l logg) print() {
	for _, v := range l.data {
		fmt.Println(v)
	}
}
